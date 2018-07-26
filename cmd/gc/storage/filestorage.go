package storage

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/maxsakharov/gatlingcomparison/cmd/gc/models"
)

type FileStorage struct {
	ReportPath string
}

func (fs FileStorage) GetAllReports() ([]models.ReportMetadata, error) {
	files, err := fs.listFiles()
	if err != nil {
		return nil, err
	}

	var reports []models.ReportMetadata

	for _, file := range *files {
		ext := strings.Replace(filepath.Ext(file.Name()), ".", "", 1)
		report := models.ReportMetadata{
			Name:       file.Name(),
			DateUpload: file.ModTime(),
			Type:       models.ReportTypeForName(ext),
		}
		reports = append(reports, report)
	}

	return reports, nil
}

func (fs FileStorage) GetReports(n ...string) ([]models.Report, error) {
	return nil, nil
}

func (fs FileStorage) GetRport(n string) (models.Report, error) {
	return models.Report{}, nil
}

func (fs FileStorage) SaveReport(r models.ReportMetadata) error {
	return nil
}

func (fs FileStorage) DeleteReport(n string) (bool, error) {

	if len(n) < 1 {
		panic("Empty file name")
	}

	if !isValidFilename(n) {
		panic("Invalid file name")
	}

	path := fs.ReportPath + "/" + n

	if _, err := os.Stat(path); err == nil {
		err = os.Remove(path)
		return true, err
	}

	return false, nil
}

func (fs FileStorage) listFiles() (*[]os.FileInfo, error) {
	var files []os.FileInfo

	err := filepath.Walk(fs.ReportPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Panic(err)
		}
		if !info.IsDir() {
			files = append(files, info)
		}
		return nil
	})

	return &files, err
}

// TODO make more clear validation for parent directory check
func isValidFilename(f string) bool {
	return !strings.Contains(f, "..")
}
