package dataminer

type DocDataMiner struct {
	DefaultDataMine
}

func (d *DocDataMiner) OpenFile(path string) string {
	return "file content doc"
}

func (d *DocDataMiner) ExtractData(content string) []string {
	return []string{"raw data doc"}
}

func (d *DocDataMiner) ParseData(data []string) []string {
	return []string{"parse data doc"}
}

func (d *DocDataMiner) AnalyzeData(data []string) string {
	return "analyzed doc data"
}
