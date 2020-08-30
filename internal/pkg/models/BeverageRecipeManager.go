package models

import (
	"errors"
)

// BeverageRecipeManager stores and manages recipes
type BeverageRecipeManager struct {
	beverageRecipes []*BeverageRecipe // all beverage recipes, stored in array
}

func NewBeverageRecipeManager() BeverageRecipeManager {
	return BeverageRecipeManager{
		beverageRecipes: []*BeverageRecipe{},
	}
}

func (brm *BeverageRecipeManager) AddBeverageRecipe(b *BeverageRecipe) {
	brm.beverageRecipes = append(brm.beverageRecipes, b)
}

// GetBeverageRecipe returns the recipe for a beverage
// or error if recipe is not supported
func (brm *BeverageRecipeManager) GetBeverageRecipe(beverageName string) (beverageRecipe *BeverageRecipe, err error) {
	for _, b := range brm.beverageRecipes {
		if b.GetBeverage().GetName() == beverageName {
			return b, nil
		}
	}
	return beverageRecipe, errors.New("beverage not supported")
}
