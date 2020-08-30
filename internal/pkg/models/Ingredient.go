package models

// Ingredient represents the ingredients the machine contains, read from config machine.total_items_count
type Ingredient struct {
	name string
}

func NewIngredient(name string) Ingredient {
	return Ingredient{
		name: name,
	}
}

func (i *Ingredient) GetName() string {
	return i.name
}
