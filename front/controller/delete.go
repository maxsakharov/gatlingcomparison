package controller

import (
	"log"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/maxsakharov/gatlingcomparison/storage"
)

type DeleeteController struct {
	Storage storage.Storage
}

func (dc *DeleeteController) DeleteReports(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	reps := strings.Split(r.URL.Query().Get("reports"), ",")
	for _, r := range reps {
		del, err := dc.Storage.DeleteReport(r)

		if del == false {
			log.Printf("Report with name '%s' does not exists", r)
		}

		if err != nil {
			log.Panic(err)
		}
	}

	http.Redirect(w, r, "/reports", http.StatusSeeOther)
}
