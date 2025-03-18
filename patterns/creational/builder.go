package creational

type IBuilder interface {
	setProtein()
	setCarbohydrates()
	setFat()
	GetBurger() Burger
}

func GetBuilder(builderType string) IBuilder {
	switch builderType {
	case "cheese":
		return newCheeseBurgerBuilder()
	case "veggie":
		return newVeggieBurgerBuilder()
	case "chicken":
		return newChickenBurgerBuilder()
	}

	return nil
}

type Chef struct {
	builder IBuilder
}

func NewChef(b IBuilder) *Chef {
	return &Chef{
		builder: b,
	}
}

func (c *Chef) SetBuilder(b IBuilder) {
	c.builder = b
}

func (c *Chef) BuildBurger() Burger {
	c.builder.setProtein()
	c.builder.setCarbohydrates()
	c.builder.setFat()
	return c.builder.GetBurger()
}

type Burger struct {
	Protein       string
	Carbohydrates string
	Fat           string
}

type CheeseBurgerBuilder struct {
	Burger
}

func newCheeseBurgerBuilder() *CheeseBurgerBuilder {
	return &CheeseBurgerBuilder{}
}

func (b *CheeseBurgerBuilder) setProtein() {
	b.Protein = "Cheese"
}

func (b *CheeseBurgerBuilder) setCarbohydrates() {
	b.Carbohydrates = "Bread"
}

func (b *CheeseBurgerBuilder) setFat() {
	b.Fat = "Cheese"
}

func (b *CheeseBurgerBuilder) GetBurger() Burger {
	return b.Burger
}

type VeggieBurgerBuilder struct {
	Burger
}

func newVeggieBurgerBuilder() *VeggieBurgerBuilder {
	return &VeggieBurgerBuilder{}
}

func (b *VeggieBurgerBuilder) setProtein() {
	b.Protein = "Veggie"
}

func (b *VeggieBurgerBuilder) setCarbohydrates() {
	b.Carbohydrates = "Bread"
}

func (b *VeggieBurgerBuilder) setFat() {
	b.Fat = "Lettuce"
}

func (b *VeggieBurgerBuilder) GetBurger() Burger {
	return b.Burger
}

type ChickenBurgerBuilder struct {
	Burger
}

func newChickenBurgerBuilder() *ChickenBurgerBuilder {
	return &ChickenBurgerBuilder{}
}

func (b *ChickenBurgerBuilder) setProtein() {
	b.Protein = "Chicken"
}

func (b *ChickenBurgerBuilder) setCarbohydrates() {
	b.Carbohydrates = "Bread"
}

func (b *ChickenBurgerBuilder) setFat() {
	b.Fat = "Chicken"
}

func (b *ChickenBurgerBuilder) GetBurger() Burger {
	return b.Burger
}
