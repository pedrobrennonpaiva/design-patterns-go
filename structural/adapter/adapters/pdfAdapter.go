package adapters

type PdfAdapter interface {
	Generate(filename, content string) error
}
