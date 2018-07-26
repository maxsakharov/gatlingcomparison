package web

import (
	"errors"
	"log"
	"testing"
	"time"

	"github.com/maxsakharov/gatlingcomparison/cmd/gc/models"
	"github.com/maxsakharov/gatlingcomparison/cmd/gc/storage/mocks"
	"github.com/stretchr/testify/assert"
)

func TestListReportsT(t *testing.T) {
	rep1 := models.ReportMetadata{"test1", time.Unix(10000000, 0), models.REPORT_TYPE_HTML}
	rep2 := models.ReportMetadata{"test2", time.Unix(20000000, 0), models.REPORT_TYPE_HTML}

	storage := &mocks.Storage{}
	storage.On("GetAllReports").Return([]models.ReportMetadata{rep1, rep2}, nil).Once()
	subject := Controller{storage}

	reports := subject.ListReports()

	expected := models.ReportMetadataList{[]models.ReportMetadata{rep1, rep2}}

	assert.Equal(t, expected, reports, "Expected reports to match")
}

func TestListReportsWhenStorageFail(t *testing.T) {
	storage := &mocks.Storage{}
	storage.On("GetAllReports").Return(nil, errors.New("Test failed report")).Once()
	subject := Controller{storage}

	defer func() {
		recover()
	}()

	subject.ListReports()

	log.Fatal("Storage error was expected, but wasn't thrown")
}

func TestDeleteReport(t *testing.T) {
	storage := &mocks.Storage{}
	storage.On("DeleteReport", "rep1").Return(true, nil).Once()
	storage.On("DeleteReport", "rep2").Return(true, nil).Once()

	subject := Controller{storage}

	subject.DeleteReports([]string{"rep1", "rep2"})

	storage.AssertExpectations(t)
}

func TestDeleteReportsWhenStorageFail(t *testing.T) {
	storage := &mocks.Storage{}
	storage.On("DeleteReport").Return(false, errors.New("Test failed report")).Once()
	subject := Controller{storage}

	defer func() {
		recover()
	}()

	subject.DeleteReports([]string{"rep1"})

	log.Fatal("Storage error was expected, but wasn't thrown")
}

// func (ms MockStorage) GetAllReports() ([]models.ReportMetadata, error) {
// 	rmd := []models.ReportMetadata{
// 		models.ReportMetadata{"test1", time.Unix(10000000, 0), models.REPORT_TYPE_HTML},
// 		models.ReportMetadata{"test2", time.Unix(20000000, 0), models.REPORT_TYPE_HTML},
// 	}
// 	return rmd, ms.Err
// }

// func (ms MockStorage) GetReports(n ...string) ([]models.Report, error) {
// 	repReq1 := []models.ReportRequest{
// 		models.ReportRequest{
// 			Name:        "test1",
// 			RequestName: "test_req_1",
// 			Ok:          1,
// 			Ko:          2,
// 			Rps:         3,
// 			Min:         4,
// 			Max:         5,
// 			Pct50:       1.11,
// 			Pct75:       2.22,
// 			Pct90:       3.33,
// 			Pct95:       4.44,
// 			Pct99:       5.55,
// 		},
// 	}

// 	repReq2 := []models.ReportRequest{
// 		models.ReportRequest{
// 			Name:        "test2",
// 			RequestName: "test_req_2",
// 			Ok:          10,
// 			Ko:          20,
// 			Rps:         30,
// 			Min:         40,
// 			Max:         50,
// 			Pct50:       10.11,
// 			Pct75:       20.22,
// 			Pct90:       30.33,
// 			Pct95:       40.44,
// 			Pct99:       50.55,
// 		},
// 	}

// 	report := []models.Report{
// 		models.Report{&repReq1},
// 		models.Report{&repReq2},
// 	}

// 	return report, ms.Err
// }

// func (ms MockStorage) GetRport(n string) (models.Report, error) {
// 	repReq1 := []models.ReportRequest{
// 		models.ReportRequest{
// 			Name:        "test1",
// 			RequestName: "test_req_1",
// 			Ok:          1,
// 			Ko:          2,
// 			Rps:         3,
// 			Min:         4,
// 			Max:         5,
// 			Pct50:       1.11,
// 			Pct75:       2.22,
// 			Pct90:       3.33,
// 			Pct95:       4.44,
// 			Pct99:       5.55,
// 		},
// 	}

// 	return models.Report{&repReq1}, ms.Err
// }

// func (ms MockStorage) SaveReport(r models.ReportMetadata) error {
// 	return ms.Err
// }

// func (ms MockStorage) DeleteReport(n string) (bool, error) {
// 	return true, ms.Err
// }
