package flyweight

import "fmt"

func Run() {

	factory := NewDeliveryFactory()

	deliveryContext(factory, "Alice", "123", "Av Brasil", "RJ")
	deliveryContext(factory, "Bob", "456", "Av Paulista", "SP")
	deliveryContext(factory, "Charlie", "789", "Av Brasil", "RJ")
	deliveryContext(factory, "Daniel", "012", "Av Brasil", "RJ")

	fmt.Println("Locations in factory:")
	for key, location := range factory.GetLocations("Av Brasil", "RJ") {
		fmt.Printf("Key: %s, Location: %+v\n", key, location)
	}
}

func deliveryContext(factory *DeliveryFactory, name, number, street, city string) {
	intrinsicState := DeliveryLocationData{
		street: street,
		city:   city,
	}
	location := factory.MakeLocation(intrinsicState)
	location.deliver(name, number)
}

type DeliveryFlyweight interface {
	deliver(name string, number string)
}

type DeliveryLocationData struct {
	street string
	city   string
}

type DeliveryLocation struct {
	intrinsicState DeliveryLocationData
}

func NewDeliveryLocation(intrinsicState DeliveryLocationData) *DeliveryLocation {
	return &DeliveryLocation{
		intrinsicState: intrinsicState,
	}
}

func (d *DeliveryLocation) deliver(name string, number string) {
	fmt.Println("Delivering to:", name)
	fmt.Println("Street:", d.intrinsicState.street)
	fmt.Println("City:", d.intrinsicState.city)
	fmt.Println("Number:", number)
	fmt.Println("Delivery completed.")
}

type DeliveryDictionary = map[string]DeliveryFlyweight

type DeliveryFactory struct {
	locations DeliveryDictionary
}

func NewDeliveryFactory() *DeliveryFactory {
	return &DeliveryFactory{
		locations: make(DeliveryDictionary),
	}
}

func (f *DeliveryFactory) createId(intrinsicState DeliveryLocationData) string {
	return fmt.Sprintf("%s-%s", intrinsicState.street, intrinsicState.city)
}

func (f *DeliveryFactory) MakeLocation(intrinsicState DeliveryLocationData) DeliveryFlyweight {
	key := f.createId(intrinsicState)
	if location, exists := f.locations[key]; exists {
		return location
	}

	location := NewDeliveryLocation(intrinsicState)
	f.locations[key] = location
	return location
}

func (f *DeliveryFactory) GetLocations(street string, city string) DeliveryDictionary {
	return f.locations
}
