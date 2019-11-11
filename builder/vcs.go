package builder

import (
	"encoding/json"
	"errors"
	"fmt"
	"path"
	"path/filepath"

	common "github.com/ragecryx/bob/common"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
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

		_, err := git.PlainClone(finalDir, false, &git.CloneOptions{
			URL: recipe.Repository.URL,
			// Progress:      os.Stdout,
			ReferenceName: plumbing.ReferenceName(refBranch),
			SingleBranch:  true,
		})

		if err != nil {
			common.BuilderLog.Errorf("! Error checking out Git repo %s: %s", title, err)
			return finalDir, err
		}

		return finalDir, nil
	}

	return finalDir, errors.New("Unsupported VCS")
}

// IsManualTrigger detects a dummy payload that forces
// build of the recipe.
func IsManualTrigger(r common.Recipe, payload []byte) bool {
	var data map[string]interface{}

	if earlyParseErr := json.Unmarshal(payload, &data); earlyParseErr != nil {
		common.BuilderLog.Errorf("! Error parsing json: %s", earlyParseErr)
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

	var data map[string]interface{}
	if earlyParseErr := json.Unmarshal(payload, &data); earlyParseErr != nil {
		common.BuilderLog.Errorf("! Error parsing json: %s", earlyParseErr)
		return false
	}

	if _, ok := data["ref"]; !ok {
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
