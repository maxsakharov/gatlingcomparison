package storage

import (
	"github.com/maxsakharov/gatlingcomparison/cmd/gc/models"
)

type Storage interface {
	GetAllReports() ([]models.ReportMetadata, error)
	GetReports(n ...string) ([]models.Report, error)
	GetRport(n string) (models.Report, error)
	SaveReport(r models.ReportMetadata) error
	DeleteReport(n string) (bool, error)
}
