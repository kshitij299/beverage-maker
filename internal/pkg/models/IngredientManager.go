package models

import (
	"errors"
	"sync"
)

// IngredientManager manages all ingredients and ingredient containers and hides this complexity from the rest of the system
type IngredientManager struct {
	ingredientContainers []*IngredientContainer // ingredient containers the machine has
	ingredientLock       *sync.Mutex            //to synchronize access to ingredients
}

func NewIngredientManager() IngredientManager {
	return IngredientManager{
		ingredientContainers: []*IngredientContainer{},
		ingredientLock:       &sync.Mutex{},
	}
}

// AddOrUpdateIngredient adds new ingredient or updates existing ingredients quantity
// this can be used to refill old ingredients or add new ingredients to the system
func (im *IngredientManager) AddOrUpdateIngredient(ingredientName string, qty int) {

	im.ingredientLock.Lock()
	defer im.ingredientLock.Unlock()

	isAlreadyPresent := false
	for _, ingredientContainer := range im.ingredientContainers {
		if ingredientContainer.GetIngredient().GetName() == ingredientName {
			isAlreadyPresent = true
			ingredientContainer.SetQuantity(ingredientContainer.GetQuantity() + qty)
		}
	}

	if !isAlreadyPresent {
		ingredient := NewIngredient(ingredientName)
		ingredientContainer := NewIngredientContainer(&ingredient, qty)
		im.ingredientContainers = append(im.ingredientContainers, &ingredientContainer)
	}
}

// GetIngredientsForRecipe returns the ingredients required by the provided recipe
// returns error if some ingredient is unavailable or insufficient
func (im *IngredientManager) GetIngredientsForRecipe(recipe *BeverageRecipe) (ingredients []Ingredient, err error) {

	im.ingredientLock.Lock()
	defer im.ingredientLock.Unlock()

	//for each ingredient as per recipe, check if quantity is sufficient to create recipe
	for ingredient, qty := range recipe.GetIngredientQuantities() {
		isPresent, isSufficient := false, false
		for _, ingredientContainer := range im.ingredientContainers {
			if ingredient == ingredientContainer.GetIngredient().GetName() {
				isPresent = true
				if qty <= ingredientContainer.GetQuantity() {
					isSufficient = true
				}
			}
		}
		if isPresent && isSufficient {
			continue
		} else {
			if !isPresent {
				return ingredients, errors.New(ingredient + " is not available")
			}
			if !isSufficient {
				return ingredients, errors.New(ingredient + " is not sufficient")
			}
		}
	}

	//the code reached here means all ingredients are present and sufficient, so now return them
	ingredients = []Ingredient{}
	for ingredient, qty := range recipe.GetIngredientQuantities() {
		for _, ingredientContainer := range im.ingredientContainers {
			if ingredient == ingredientContainer.GetIngredient().GetName() {
				currentQty := ingredientContainer.GetQuantity()
				ingredientContainer.SetQuantity(currentQty - qty)
				ingredients = append(ingredients, NewIngredient(ingredient))
			}
		}
	}
	return ingredients, nil

}
