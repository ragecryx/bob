package builder

import (
	"log"
	"time"

	types "github.com/ragecryx/bob/common"
)

var (
	taskAvailability []bool
	taskChannels     []chan types.Recipe
)

func ConfigureTasks(amount uint) {
	taskAvailability = make([]bool, amount)
	taskChannels = make([]chan types.Recipe, amount)

	for i, _ := range taskChannels {
		taskAvailability[i] = true
		taskChannels[i] = make(chan types.Recipe)
		// Warmup some tasks
		go RunTask(i)
	}
}

func RunTask(index int) {
	for {
		recipe := <-taskChannels[index]
		taskAvailability[index] = false

		// Do Build
		log.Printf("[T#%d] Building '%s'", index, recipe.Repository.URL)
		time.Sleep(40 * time.Second)
		log.Printf("[T#%d] Finished", index)

		taskAvailability[index] = true
	}
}

func QueueRecipe(recipe types.Recipe) bool {
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
