package builder

import (
	"log"

	common "github.com/ragecryx/bob/common"
)

var (
	taskAvailability []bool
	taskChannels     []chan common.Recipe
)

func ConfigureTasks(amount uint) {
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
	// config := common.GetConfig()
	for {
		recipe := <-taskChannels[index]
		taskAvailability[index] = false

		// Do Build
		log.Printf("[T#%d] Building '%s'", index, recipe.Repository.URL)
		Clone(&recipe)
		log.Printf("[T#%d] Finished", index)

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
