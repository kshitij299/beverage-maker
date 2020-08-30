package main

import (
	"flag"

	"github.com/kshitij299/beverage-maker/internal/pkg/config"
	. "github.com/kshitij299/beverage-maker/internal/pkg/models"
)

func main() {
	//declare flag for the Config file
	configFileName := flag.String("config", "", "config file option")

	//declare flag for test file
	testConfigFileName := flag.String("test-config", "", "test config file option")

	//Read the config file flags
	flag.Parse()
	err := config.ReadConfigFromFile(*configFileName)
	if err != nil {
		panic(err)
	}
	err = config.ReadTestConfigFromFile(*testConfigFileName)
	if err != nil {
		panic(err)
	}

	// testing parallelism with random beverages
	beverageMachine := createNewBeverageMachineFromConfig()
	testParallelismWithRandomBeverages(&beverageMachine, config.TestConfig().Test.ParallelThreads)

	// testing parallelism with given beverages
	beverageMachine = createNewBeverageMachineFromConfig()
	testParallelismWithBeverages(&beverageMachine, config.TestConfig().Test.Beverages)
}

func createNewBeverageMachineFromConfig() (beverageMachine BeverageMachine) {
	//1. add all ingredients
	ingredientManager := NewIngredientManager()
	for name, qty := range config.Config().Machine.Ingredients {
		ingredientManager.AddOrUpdateIngredient(name, qty)
	}

	//2. add beverages to beverage list and beverage recipes to recipe manager
	beverageRecipeManager := NewBeverageRecipeManager()
	beverages := []*Beverage{}
	for name, ingredientToQtyMap := range config.Config().Machine.Beverages {
		beverage := NewBeverage(name)
		beverages = append(beverages, &beverage)
		beverageRecipe := NewBeverageRecipe(&beverage, ingredientToQtyMap)
		beverageRecipeManager.AddBeverageRecipe(&beverageRecipe)
	}

	//3. add dispensers
	beverageDispensers := []*BeverageDispenser{}
	for i := 1; i <= config.Config().Machine.Outlets.Count; i++ {
		beverageDispenser := NewBeverageDispenser(i)
		beverageDispensers = append(beverageDispensers, &beverageDispenser)
	}

	//4. initialize beverage machine
	beverageMachine = NewBeverageMachine(beverages, beverageDispensers, &beverageRecipeManager, &ingredientManager)

	//5. Return
	return beverageMachine
}
