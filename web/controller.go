package web

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/maxsakharov/gatlingcomparison/cmd/gc/models"
	"github.com/maxsakharov/gatlingcomparison/cmd/gc/storage"
)

type Controller struct {
	storage storage.Storage
}

func (c *Controller) ListReports() models.ReportMetadataList {
	reports, err := c.storage.GetAllReports()
	if err != nil {
		log.Panic(err)
	}
	return models.ReportMetadataList{Reports: reports}
}

func (c *Controller) AddReport(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	r.ParseForm()

}

func (c *Controller) DeleteReports(names []string) {
	for _, name := range names {
		del, err := c.storage.DeleteReport(name)

		if del == false {
			log.Printf("Report with name '%s' does not exists", name)
		}

		if err != nil {
			log.Panic(err)
		}
	}
}
