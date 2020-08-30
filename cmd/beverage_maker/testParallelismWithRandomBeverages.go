package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	. "github.com/kshitij299/beverage-maker/internal/pkg/models"
)

func testParallelismWithRandomBeverages(bm *BeverageMachine, parallelThreads int) {

	fmt.Printf("\nstarting test: testParallelismWithRandomBeverages\n\n")

	beverages := bm.Beverages()

	wg := sync.WaitGroup{}
	for i := 0; i < parallelThreads; i++ {
		wg.Add(1)
		go func(i int) {
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

			//2. get a random beverage from the list of supported beverages
			rand.Seed(time.Now().Unix())
			randomIndex := rand.Intn(len(beverages))
			randomBeverageName := beverages[randomIndex].GetName()

			fmt.Printf("thread %d-> starting to prepare %s on dispenser %d...\n", i+1, randomBeverageName, dispenser.GetId())

			//3. request the random beverage from the dispenser
			beverage, err := bm.RequestBeverage(dispenser, randomBeverageName)
			if err != nil {
				fmt.Printf("thread %d-> dispenser %d says: %s\n", i+1, dispenser.GetId(), err.Error())
				fmt.Printf("thread %d-> end\n", i+1)
				return
			}

			fmt.Printf("thread %d-> successfully served %s on dispenser %d\n", i+1, beverage.GetName(), dispenser.GetId())
			fmt.Printf("thread %d-> end\n", i+1)
		}(i)
	}
	wg.Wait()

	fmt.Printf("\ncompleted test: testParallelismWithRandomBeverages\n")

}
