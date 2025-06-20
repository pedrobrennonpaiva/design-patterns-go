package proxy

import (
	"design-patterns-go/structural/proxy/proxies"
	"fmt"
	"log"
)

func Run() {
	reportCache := proxies.NewReportCache()
	reportRepo := proxies.NewReportRepository()
	reportGeneratorProxy := proxies.NewReportGeneratorProxy(reportCache)

	id, err := reportRepo.GetReport(1)
	if err != nil {
		fmt.Println("Error getting report:", err)
		return
	}

	log.Println("Generating report")
	report, err := reportGeneratorProxy.GenerateReport(id)
	if err != nil {
		fmt.Println("Error generating report:", err)
		return
	}
	log.Println("Generated Report:", *report)

	// Simulate a second request to demonstrate caching inside the proxy
	log.Println("Generating report again")
	report, err = reportGeneratorProxy.GenerateReport(id)
	if err != nil {
		fmt.Println("Error generating report:", err)
		return
	}
	log.Println("Generated Report from Cache:", *report)

	log.Println("Report generation completed successfully.")
}
