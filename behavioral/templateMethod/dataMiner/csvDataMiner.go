package dataminer

type CsvDataMiner struct{}

func (d *CsvDataMiner) OpenFile(path string) string {
	return "file content csv"
}

func (d *CsvDataMiner) ExtractData(content string) []string {
	return []string{"raw data csv"}
}

func (d *CsvDataMiner) ParseData(data []string) []string {
	return []string{"parse data csv"}
}
