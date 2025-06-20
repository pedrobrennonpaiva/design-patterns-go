package proxies

import (
	"fmt"
	"time"
)

type ReportRepository struct{}

func NewReportRepository() *ReportRepository {
	return &ReportRepository{}
}

func (r *ReportRepository) GetReport(id int) (int, error) {
	// Simulate a database call
	return id, nil
}

type ReportGeneratorInterface interface {
	GenerateReport(id int) (*string, error)
}

type ReportGenerator struct{}

func NewReportGenerator() *ReportGenerator {
	return &ReportGenerator{}
}

func (g *ReportGenerator) GenerateReport(id int) (*string, error) {
	// Simulate report generation without cache
	time.Sleep(5 * time.Second)

	report := fmt.Sprintf("Report content for ID %d", id)
	return &report, nil
}

type ReportCacher interface {
	GetCachedReport(id int) (*string, error)
	SetCachedReport(report string)
}

type ReportCache struct {
	reportCached string
}

func NewReportCache() *ReportCache {
	return &ReportCache{}
}

func (c *ReportCache) GetCachedReport(id int) (*string, error) {
	// Simulate cache retrieval
	if c.reportCached == "" {
		return nil, fmt.Errorf("no cached report found for ID %d", id)
	}

	report := fmt.Sprintf("Cached report content for ID %d", id)
	return &report, nil
}

func (c *ReportCache) SetCachedReport(report string) {
	c.reportCached = report
}
