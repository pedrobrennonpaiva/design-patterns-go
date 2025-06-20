package adapters

type ProviderOneAdapter struct{}

func NewProviderOneAdapter() *ProviderOneAdapter {
	return &ProviderOneAdapter{}
}

func (p *ProviderOneAdapter) Generate(filename, content string) error {
	// Simulate generating a PDF using ProviderOne's API
	provider := NewProviderOnePdf()
	if err := provider.LoadHtml(content); err != nil {
		return err
	}

	if err := provider.SetPaper("A4", "portrait"); err != nil {
		return err
	}

	return provider.Render()
}
