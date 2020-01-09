package builder

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"

	common "github.com/ragecryx/bob/common"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/ssh"
)

// Clone uses go-git library to single-branch clone
// the repository of the provided Recipe.
//
// Returns the final path the the repo was checked out
func Clone(recipe *common.Recipe) (string, error) {
	config := common.GetConfig()

	title := recipe.Repository.Name

	tmpDir := common.GenerateTmpDirName()
	finalDir := path.Join(config.WorkspacePath, tmpDir)

	if !path.IsAbs(finalDir) {
		dir, errPath := filepath.Abs(finalDir)

		if errPath != nil {
			finalDir = dir
		}
	}

	common.BuilderLog.Infof("* Checking out %s in %s", title, finalDir)

	if recipe.Repository.VCS == "git" {
		// Clone
		refBranch := fmt.Sprintf("refs/heads/%s", recipe.Repository.Branch)

		options := git.CloneOptions{
			URL: recipe.Repository.URL,
			// Progress:      os.Stdout,
			ReferenceName: plumbing.ReferenceName(refBranch),
			SingleBranch:  true,
		}

		// Check if we have ssh key for auth
		// TODO: Check recipe.Repository.SSHKey
		// if it's NOT absolute path search for the file
		// in $HOME/.ssh/ like:
		// sshPath := os.Getenv("HOME") + "/.ssh/THE_SSHKey_PARAM"
		if recipe.Repository.SSH != nil {
			// Read file
			var publicKey *ssh.PublicKeys
			sshKey, _ := ioutil.ReadFile(recipe.Repository.SSH.KeyFile)

			var password = ""
			if recipe.Repository.SSH.PasswdEnv != nil {
				password = os.Getenv(*recipe.Repository.SSH.PasswdEnv)
			}

			// TODO: Maybe user param (1st param) needs to be
			// read from recipe URL?
			publicKey, keyError := ssh.NewPublicKeys("git", []byte(sshKey), password)

			if keyError != nil {
				common.BuilderLog.Errorf("Bad SSH key file. Error: %s", keyError)
			}

			options.Auth = publicKey
		}

		_, err := git.PlainClone(finalDir, false, &options)

		if err != nil {
			common.BuilderLog.Errorf("! Error checking out Git repo %s: %s", title, err)
			return finalDir, err
		}

		return finalDir, nil
	}

	return finalDir, errors.New("Unsupported VCS")
}

// Tries to unmarshal JSON
func getValidJSON(payload []byte) (bool, map[string]interface{}) {
	var data map[string]interface{}

	if earlyParseErr := json.Unmarshal(payload, &data); earlyParseErr != nil {
		common.BuilderLog.Errorf("! Error parsing json: %s", earlyParseErr)
		return false, nil
	}

	return true, data
}

// IsManualTrigger detects a dummy payload that forces
// build of the recipe.
func IsManualTrigger(r common.Recipe, payload []byte) bool {
	if valid, _ := getValidJSON(payload); !valid {
		return false
	}

	trigger := ManualTrigger{}

	if parseErr := json.Unmarshal(payload, &trigger); parseErr != nil {
		return false
	}

	if trigger.Who == "Bob" && trigger.ForceBuild == true {
		return true
	}

	return false
}

// IsGithubMerge checks whether the provided Recipe
// matches the Github payload of the hook request.
func IsGithubMerge(r common.Recipe, payload []byte) bool {
	if !r.IsHostedIn("github") {
		return false
	}

	if valid, data := getValidJSON(payload); !valid {
		return false
	} else if _, ok := data["ref"]; !ok {
		return false
	}

	event := GHPushEvent{}

	if parseErr := json.Unmarshal(payload, &event); parseErr != nil {
		return false
	}

	if fmt.Sprintf("refs/heads/%s", r.Repository.Branch) == event.Ref {
		return true
	}

	return false
}

// IsBitBucketMerge checks whether the provided Recipe
// matches the BitBucket payload of the hook request.
func IsBitBucketMerge(r common.Recipe, payload []byte) bool {
	if !r.IsHostedIn("bitbucket") {
		return false
	}

	if valid, _ := getValidJSON(payload); !valid {
		common.BuilderLog.Errorf("BB failed json validation")
		return false
	}

	event := BBPushEvent{}

	if parseErr := json.Unmarshal(payload, &event); parseErr != nil {
		common.BuilderLog.Errorf("BB failed json to struct unmarshalling")
		return false
	}

	if strings.Compare(r.Repository.Name, event.Repository.Name) != 0 {
		common.BuilderLog.Errorf("BitBucket repository name mismatch. Recipe: %s vs Payload: %s", r.Repository.Name, event.Repository.Name)
		return false
	}

	isAboutRecipeBranch := false

	for _, change := range event.Push.Changes {
		isAboutRecipeBranch = isAboutRecipeBranch || (change.New != nil && strings.Compare(change.New.Name, r.Repository.Branch) == 0)
	}

	return isAboutRecipeBranch
}
