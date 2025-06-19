package dataminer

import "fmt"

type DataMiner interface {
	OpenFile(path string) string
	ExtractData(content string) []string
	ParseData(data []string) []string
	AnalyzeData(data []string) string
	SendReport(results string)
}

type DataMine struct {
	dataMiner DataMiner
}

func NewDataMine(dataMiner DataMiner) *DataMine {
	return &DataMine{dataMiner: dataMiner}
}

func (d *DataMine) Mine(path string) {
	fileContent := d.dataMiner.OpenFile(path)
	fmt.Println("File content:", fileContent)

	rawData := d.dataMiner.ExtractData(fileContent)
	fmt.Println("Raw data extracted:", rawData)

	data := d.dataMiner.ParseData(rawData)
	fmt.Println("Parsed data:", data)

	analysisResult := d.dataMiner.AnalyzeData(data)
	fmt.Println("Analysis result:", analysisResult)

	d.dataMiner.SendReport(analysisResult)
	fmt.Println("Data mining completed successfully.")
}

type DefaultDataMine struct{}

func (d *DefaultDataMine) AnalyzeData(data []string) string {
	return "analyzed default data"
}

func (d *DefaultDataMine) SendReport(results string) {
	fmt.Println("Sending report for default data:", results)
}
