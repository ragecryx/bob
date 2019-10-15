package builder

import (
	"log"
	"os"

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

		// Do Build
		log.Printf("[T#%d] Building '%s'", index, title)
		cloneDir, err := Clone(&recipe)

		if err != nil && config.CleanupBuilds {
			errCleanup := os.Remove(cloneDir)

			if errCleanup != nil {
				log.Panicf("! Could not cleanup %s after failed cloning", cloneDir)
			}
		}

		log.Printf("[T#%d] Finished '%s'", index, title)

		taskAvailability[index] = true
	}
}

func QueueRecipe(recipe common.Recipe) bool {
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
