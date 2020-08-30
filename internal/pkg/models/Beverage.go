package models

// Beverage defines the core entity of the system i.e. a beverage
type Beverage struct {
	name string
}

func NewBeverage(name string) Beverage {
	return Beverage{
		name: name,
	}
}

func (b *Beverage) GetName() string {
	return b.name
}
