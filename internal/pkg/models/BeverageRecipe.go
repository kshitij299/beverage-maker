package models

// BeverageRecipe represents the recipe of a beverage, it is read from config file
type BeverageRecipe struct {
	beverage             *Beverage      // the beverage for which the following recipe is
	ingredientQuantities map[string]int // the ingredients required in the recipe, with their units
}

func NewBeverageRecipe(b *Beverage, iqm map[string]int) BeverageRecipe {
	return BeverageRecipe{
		beverage:             b,
		ingredientQuantities: iqm,
	}
}

func (bc *BeverageRecipe) GetBeverage() *Beverage {
	return bc.beverage
}

func (bc *BeverageRecipe) GetIngredientQuantities() map[string]int {
	return bc.ingredientQuantities
}
