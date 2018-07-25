package controller

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/maxsakharov/gatlingcomparison/front/views"
	"github.com/maxsakharov/gatlingcomparison/models"
	"github.com/maxsakharov/gatlingcomparison/storage"
)

type ListController struct {
	Storage storage.Storage
}

func (lc *ListController) ListReports(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	reports, err := lc.Storage.GetAllReports()
	if err != nil {
		log.Panic(err)
	} else {
		reports := models.ReportMetadataList{Reports: reports}
		views.ListReports(w, reports)
	}
}
