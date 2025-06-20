package adapter

import (
	"design-patterns-go/structural/adapter/adapters"
	"fmt"
)

func Run() {

	providerOne := adapters.NewProviderOneAdapter()
	// providerTwo := adapters.NewProviderTwoAdapter()
	salesReport := adapters.NewSalesReport(providerOne)

	err := salesReport.GenerateReport("filename.pdf", "content of the report")
	if err != nil {
		fmt.Println("Error generating report:", err)
		return
	}

	fmt.Println("Generated report successfully")
}
