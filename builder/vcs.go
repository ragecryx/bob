package builder

import (
	"errors"
	"fmt"
	"log"
	"path"
	"path/filepath"

	common "github.com/ragecryx/bob/common"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

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

	log.Printf("* Checking out %s in %s", title, finalDir)

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
			log.Panicf("! Error checking out Git repo %s: %s", title, err)
		}

		return finalDir, nil
	}

	return finalDir, errors.New("Unsupported VCS")
}
