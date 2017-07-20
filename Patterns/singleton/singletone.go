package singleton

import (
	"sync"
	"sync/atomic"
)

//http://marcio.io/2015/07/singleton-pattern-in-go/
type singleton struct {
}

var instance *singleton

func GetInstance1() *singleton {
	if instance == nil {
		instance = &singleton{}   // <--- NOT THREAD SAFE
	}
	// In the above scenario, multiple go routines could evaluate the first check and they would all create an instance
	// of the singleton type and override each other. There is no guarantee which instance it will be returned here,
	// and other further operations on the instance can be come inconsistent with the expectations by the developer.
	return instance
}


var mu sync.Mutex

func GetInstance2() *singleton {
	mu.Lock()                    // <--- Unnecessary locking if instance already created
	defer mu.Unlock()

	if instance == nil {
		instance = &singleton{}
	}
	//In the code above, we can see that we solve the thread-safety issue by introducing the
	// Sync.Mutex and acquiring the Lock before creating the singleton instance. The problem is that here we are
	// performing excessive locking even when we wouldn’t be required to do so, in the case the instance has been
	// already created and we should simply have returned the cached singleton instance. On a highly concurrent code
	// base, this can generate a bottle-neck since only one go routine could get the singleton instance at a time.
	return instance
}

func GetInstance3() *singleton {
	if instance == nil {     // <-- Not yet perfect. since it's not fully atomic
		mu.Lock()
		defer mu.Unlock()

		if instance == nil {
			instance = &singleton{}
		}
	}
	return instance
}

var initialized uint32


func GetInstance4() *singleton {

	if atomic.LoadUint32(&initialized) == 1 { //LoadUInt32(&initialized)
		return instance
	}

	mu.Lock()
	defer mu.Unlock()

	if initialized == 0 {
		instance = &singleton{}
		atomic.StoreUint32(&initialized, 1)
	}
	//But using the sync/atomic package, we can atomically load and set a flag that will indicate if we have initialized
	// or not our instance.

	return instance
}