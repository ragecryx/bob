package builder

import (
	"log"
	"os"
	"os/exec"
	"strings"

	common "github.com/ragecryx/bob/common"
)

var (
	taskAvailability []bool
	taskChannels     []chan common.Recipe
)

func ConfigureTasks(amount int) {
	log.Printf("* Initializing %d tasks", amount)

	taskAvailability = make([]bool, amount)
	taskChannels = make([]chan common.Recipe, amount)

	for i, _ := range taskChannels {
		taskAvailability[i] = true
		taskChannels[i] = make(chan common.Recipe)
		// Warmup some tasks
		go RunTask(i)
	}
}

func RunTask(index int) {
	config := common.GetConfig()

	for {
		recipe := <-taskChannels[index]
		taskAvailability[index] = false
		title := recipe.Repository.Name

		// Clone
		log.Printf("[T#%d] Building '%s'", index, title)
		cloneDir, err := Clone(&recipe)

		if err != nil && config.CleanupBuilds {
			errCleanup := os.Remove(cloneDir)

			if errCleanup != nil {
				log.Panicf("! Could not cleanup %s after failed cloning", cloneDir)
			}
		}

		// Build
		args := strings.Fields(recipe.Command)
		cmd := exec.Command(args[0], args[1:]...)
		cmd.Dir = cloneDir
		result, errCmd := cmd.Output()

		if errCmd != nil && config.CleanupBuilds {
			log.Printf("[T#%d] ! Error running build cmd: %s", index, recipe.Command)
			errCleanup := os.Remove(cloneDir)

			if errCleanup != nil {
				log.Panicf("! Could not cleanup %s after failed building", cloneDir)
			}
		}

		log.Printf("[T#%d] Finished '%s' with output:\n > %s", index, title, string(result))

		taskAvailability[index] = true
	}
}

func RunRecipe(recipe common.Recipe) bool {
	// Find the first available task
	// to run the recipe
	for i := range taskAvailability {
		if taskAvailability[i] {
			taskChannels[i] <- recipe
			return true
		}
	}

	return false
}
