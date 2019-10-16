package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsHostedInGithub(t *testing.T) {
	// Create a dummy recipe
	recipe := Recipe{
		Repository: Repository{
			Name:   "testo",
			Branch: "dev",
			VCS:    "git",
			URL:    "https://github.com/hardcoredev/some_repo.git",
		},
		Command: "foo",
	}

	isGithub := recipe.IsHostedIn("github")
	isBitBucket := recipe.IsHostedIn("bitbucket")
	isPrivate := recipe.IsHostedIn("batcave")

	assert.Equal(t, isGithub, true)
	assert.Equal(t, isBitBucket, false)
	assert.Equal(t, isPrivate, false)
}

func TestIsHostedInBitBucket(t *testing.T) {
	// Create a dummy recipe
	recipe := Recipe{
		Repository: Repository{
			Name:   "untesto",
			Branch: "dev",
			VCS:    "git",
			URL:    "git@bitbucket.org:megacorp/kewl-ui.git",
		},
		Command: "foo",
	}

	isGithub := recipe.IsHostedIn("github")
	isBitBucket := recipe.IsHostedIn("bitbucket")
	isPrivate := recipe.IsHostedIn("batcave")

	assert.Equal(t, isGithub, false)
	assert.Equal(t, isBitBucket, true)
	assert.Equal(t, isPrivate, false)
}
