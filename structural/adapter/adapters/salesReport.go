package adapters

type SalesReport struct {
	adapter PdfAdapter
}

func NewSalesReport(adapter PdfAdapter) *SalesReport {
	return &SalesReport{
		adapter: adapter,
	}
}

func (r *SalesReport) GenerateReport(filename, content string) error {
	// Simulate report generation logic
	if err := r.adapter.Generate(filename, content); err != nil {
		return err
	}
	return nil
}
