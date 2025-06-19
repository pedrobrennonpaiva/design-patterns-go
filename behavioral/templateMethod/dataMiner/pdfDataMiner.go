package dataminer

type PdfDataMiner struct {
	DefaultDataMine
}

func (d *PdfDataMiner) OpenFile(path string) string {
	return "file content pdf"
}

func (d *PdfDataMiner) ExtractData(content string) []string {
	return []string{"raw data pdf"}
}

func (d *PdfDataMiner) ParseData(data []string) []string {
	return []string{"parse data pdf"}
}
