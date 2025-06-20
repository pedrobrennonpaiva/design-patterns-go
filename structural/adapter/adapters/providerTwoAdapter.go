package adapters

type ProviderTwoAdapter struct{}

func NewProviderTwoAdapter() *ProviderTwoAdapter {
	return &ProviderTwoAdapter{}
}

func (p *ProviderTwoAdapter) Generate(filename, content string) error {
	// Simulate generating a PDF using ProviderTwo's API
	provider := NewProviderTwoPdf()
	if err := provider.AddPage(); err != nil {
		return err
	}

	if err := provider.SetFont("Arial", 12); err != nil {
		return err
	}

	return provider.Write(content)
}
