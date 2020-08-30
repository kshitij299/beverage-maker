package main

import (
	"fmt"
	"sync"
	"time"

	. "github.com/kshitij299/beverage-maker/internal/pkg/models"
)

// This function tests parallelism for the given beverage machine,
// providing it with requests for the given bevergages, in parallel
// each provided beverage is requested in a parallel thread
func testParallelismWithBeverages(bm *BeverageMachine, beverageNames []string) {

	fmt.Printf("\nstarting test: testParallelismWithBeverages\n\n")

	wg := sync.WaitGroup{}
	for i, beverageName := range beverageNames {
		wg.Add(1)
		go func(i int, beverageName string) {
			defer wg.Done()

			fmt.Printf("thread %d-> start\n", i+1)

			//1. get an idle dispenser
			dispenser, err := bm.GetIdleDispenser()
			for err != nil {
				fmt.Printf("thread %d-> %s, retrying in 2 seconds...\n", i+1, err.Error())
				time.Sleep(2 * time.Second)
				dispenser, err = bm.GetIdleDispenser()
			}
			fmt.Printf("thread %d-> acquired dispenser %d\n", i+1, dispenser.GetId())

			fmt.Printf("thread %d-> starting to prepare %s on dispenser %d...\n", i+1, beverageName, dispenser.GetId())

			//2. request the beverage from the dispenser
			beverage, err := bm.RequestBeverage(dispenser, beverageName)
			if err != nil {
				fmt.Printf("thread %d-> dispenser %d says: %s\n", i+1, dispenser.GetId(), err.Error())
				fmt.Printf("thread %d-> end\n", i+1)
				return
			}

			fmt.Printf("thread %d-> successfully served %s on dispenser %d\n", i+1, beverage.GetName(), dispenser.GetId())
			fmt.Printf("thread %d-> end\n", i+1)
		}(i, beverageName)
	}
	wg.Wait()

	fmt.Println("\ncompleted test: testParallelismWithBeverages\n")
}
