package builders

import "design-patterns-go/creational/builder/models"

type VeggieBurgerBuilder struct {
	models.Burger
}

func NewVeggieBurgerBuilder() *VeggieBurgerBuilder {
	return &VeggieBurgerBuilder{}
}

func (b *VeggieBurgerBuilder) SetProtein() {
	b.Protein = "Veggie"
}

func (b *VeggieBurgerBuilder) SetCarbohydrates() {
	b.Carbohydrates = "Bread"
}

func (b *VeggieBurgerBuilder) SetFat() {
	b.Fat = "Lettuce"
}

func (b *VeggieBurgerBuilder) GetBurger() models.Burger {
	return b.Burger
}
