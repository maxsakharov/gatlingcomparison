package views

import (
	"html/template"
	"log"
	"net/http"

	"github.com/maxsakharov/gatlingcomparison/models"
)

func ListReports(w http.ResponseWriter, rl models.ReportMetadataList) {
	t, err := template.ParseFiles("front/views/template/report-list.html")
	if err != nil {
		log.Panic(err)
	} else {
		t.Execute(w, rl)
	}
}

func AddReport(w http.ResponseWriter) {
	t, err := template.ParseFiles("front/views/template/add.html")
	if err != nil {
		log.Panic(err)
	} else {
		t.Execute(w, nil)
	}
}
