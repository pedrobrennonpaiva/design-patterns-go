package proxies

import "log"

type ReportGeneratorProxy struct {
	report ReportGeneratorInterface
	cache  ReportCacher
}

func NewReportGeneratorProxy(cache ReportCacher) *ReportGeneratorProxy {
	return &ReportGeneratorProxy{
		report: NewReportGenerator(),
		cache:  cache,
	}
}

func (p *ReportGeneratorProxy) GenerateReport(id int) (*string, error) {
	// Here you can add logic to check permissions, log access, or cache results
	// For this example, we will focus on caching the report generation

	report, err := p.cache.GetCachedReport(id)
	if err != nil {
		log.Println("Cache miss, generating report...")

		report, err = p.report.GenerateReport(id)
		if err != nil {
			return nil, err
		}

		p.cache.SetCachedReport(*report)
	}

	return report, nil
}
