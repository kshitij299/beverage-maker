package models

type IngredientContainer struct {
	Ingredient *Ingredient
	Quantity   int
}

func NewIngredientContainer(i *Ingredient, qty int) IngredientContainer {
	return IngredientContainer{
		Ingredient: i,
		Quantity:   qty,
	}
}

func (ic *IngredientContainer) GetIngredient() *Ingredient {
	return ic.Ingredient
}

func (ic *IngredientContainer) GetQuantity() int {
	return ic.Quantity
}

func (ic *IngredientContainer) SetQuantity(qty int) {
	ic.Quantity = qty
}
