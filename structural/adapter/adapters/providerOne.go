package adapters

type ProviderOneInterface interface {
	LoadHtml(string) error
	SetPaper(size string, orientation string) error
	Render() error
}

type ProviderOnePdf struct{}

func NewProviderOnePdf() *ProviderOnePdf {
	return &ProviderOnePdf{}
}

func (p *ProviderOnePdf) LoadHtml(html string) error {
	// Simulate loading HTML logic
	return nil
}

func (p *ProviderOnePdf) SetPaper(size string, orientation string) error {
	// Simulate setting paper size and orientation logic
	return nil
}

func (p *ProviderOnePdf) Render() error {
	// Simulate rendering logic
	return nil
}
