package adapters

type ProviderTwoInterface interface {
	AddPage() error
	SetFont(font string, size int) error
	Write(content string) error
}

type ProviderTwoPdf struct{}

func NewProviderTwoPdf() *ProviderTwoPdf {
	return &ProviderTwoPdf{}
}

func (p *ProviderTwoPdf) AddPage() error {
	// Simulate adding a page logic
	return nil
}

func (p *ProviderTwoPdf) SetFont(font string, size int) error {
	// Simulate setting font logic
	return nil
}

func (p *ProviderTwoPdf) Write(content string) error {
	// Simulate writing content logic
	return nil
}
