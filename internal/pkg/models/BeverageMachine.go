package models

import (
	"errors"
	"time"
)

// BeverageMachine represents the machine as visible to a user
type BeverageMachine struct {
	beverages          []*Beverage            // list of beverages the machine supports, taken from config
	beverageDispensers []*BeverageDispenser   // dispensers that the machine has
	recipeManager      *BeverageRecipeManager // recipe manager of the machine
	ingredientManager  *IngredientManager     // ingredient manager of the machine
}

func NewBeverageMachine(bs []*Beverage, bds []*BeverageDispenser, brm *BeverageRecipeManager, im *IngredientManager) BeverageMachine {
	return BeverageMachine{
		beverages:          bs,
		beverageDispensers: bds,
		recipeManager:      brm,
		ingredientManager:  im,
	}
}

// GetIdleDispenser checks all dispensers and tries to acquire the first idle dispenser
// if no dispenser is idle, it returns error
func (bm *BeverageMachine) GetIdleDispenser() (b *BeverageDispenser, err error) {
	for _, beverageDispenser := range bm.beverageDispensers {
		if beverageDispenser.Acquire() {
			return beverageDispenser, nil
		}
	}
	return b, errors.New("all dispensers are busy")
}

// Beverages gives the list of beverages the system supports
func (bm *BeverageMachine) Beverages() (beverages []*Beverage) {
	return bm.beverages
}

// RequestBeverage is called by user threads to request a beverage from the machine, over a dispenser
// it prepares and serves the beverage if all ingredients are available and sufficient, else it returns error
func (bm *BeverageMachine) RequestBeverage(beverageDispenser *BeverageDispenser, beverageName string) (beverage Beverage, err error) {

	//release the dispenser instance after the beverafe is served
	defer beverageDispenser.Release()

	//1. get beverage recipe
	recipe, err := bm.recipeManager.GetBeverageRecipe(beverageName)
	if err != nil {
		return
	}

	//2. get ingredients required
	_, err = bm.ingredientManager.GetIngredientsForRecipe(recipe)
	if err != nil {
		return beverage, errors.New(beverageName + " cannot be prepared because " + err.Error())
	}

	//3. prepare beverage from ingredients
	// taking 2 seconds as the beverage preparation time
	time.Sleep(2 * time.Second)

	//4. return the beverage
	return NewBeverage(beverageName), nil

}
