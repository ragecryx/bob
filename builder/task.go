package builder

import (
	"container/list"
	"log"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"

	common "github.com/ragecryx/bob/common"
)

var (
	queue            = list.New()
	queueMutex       = &sync.Mutex{}
	taskAvailability []bool
	taskChannels     []chan common.Recipe
)

// ConfigureTasks prepares the builder
// to be ready for picking up work
func ConfigureTasks(amount int) {
	log.Printf("* Initializing %d tasks", amount)

	taskAvailability = make([]bool, amount)
	taskChannels = make([]chan common.Recipe, amount)

	for i := range taskChannels {
		taskAvailability[i] = true
		taskChannels[i] = make(chan common.Recipe)
		// Warmup some tasks
		go RunTask(i)
	}

	go queueHandler()
}

// RunTask is implementing the core
// building logic.
// A fixed number of RunTask goroutines are
// started initially based on the task_queue_size
// parameter of the config
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

		// TODO: Support running multiple commands
		// Build
		args := strings.Fields(recipe.Command)
		cmd := exec.Command(args[0], args[1:]...)
		cmd.Dir = cloneDir
		result, errCmd := cmd.Output()

		if errCmd != nil && config.CleanupBuilds {
			log.Printf("[T#%d] ! Error running build cmd: %s", index, recipe.Command)
			log.Printf(" > ! Error: %s", errCmd.Error())
			errCleanup := os.Remove(cloneDir)

			if errCleanup != nil {
				log.Panicf("! Could not cleanup %s after failed building", cloneDir)
			}
		}

		log.Printf("[T#%d] Finished '%s' with output:\n > %s", index, title, string(result))

		taskAvailability[index] = true
	}
}

// Enqueue adds a recipe to the queue
// for running when a task is available
func Enqueue(r common.Recipe) {
	queueMutex.Lock()
	queue.PushBack(r)
	queueMutex.Unlock()
}

// queueHandler checks the queue for items
// picks up the first and distributes to the
// first available task waiting for work.
func queueHandler() {
	for {
		queueMutex.Lock()
		if queue.Len() > 0 {
			e := queue.Front()
			recipe := e.Value.(common.Recipe)

			running := RunRecipe(recipe)

			if running {
				// We can remove it from the queue
				// else it will wait until picked-up
				queue.Remove(e)
			}
		}
		queueMutex.Unlock()
		time.Sleep(1 * time.Second)
	}
}

// RunRecipe tries to find an available task
// to run the given recipe instance
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
