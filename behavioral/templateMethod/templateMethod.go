package template_method

import (
	dataminer "design-patterns-go/behavioral/templateMethod/dataMiner"
	"fmt"
)

func Run() {

	// csvMiner := dataminer.CsvDataMiner{}
	// pdfMiner := dataminer.PdfDataMiner{}
	docMiner := dataminer.DocDataMiner{}
	dataMine := dataminer.NewDataMine(&docMiner)

	fmt.Println("Doc Data Miner:")
	dataMine.Mine("path/to/doc/file")
}
