package web

import (
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/maxsakharov/gatlingcomparison/cmd/gc/storage"
)

type Web struct {
	storage storage.Storage
	ctl     *Controller
	Router  *httprouter.Router
}

func NewWeb(
	s storage.Storage,
) *Web {
	web := Web{
		storage: s,
		ctl:     &Controller{s},
		Router:  httprouter.New(),
	}

	web.init()

	return &web
}

func (web *Web) init() {
	web.Router.GET("/reports", web.showListReports)

	web.Router.GET("/reports/add", web.showAddReport)
	web.Router.POST("/reports/add", web.addReport)

	web.Router.POST("/reports/compare", web.showCompareReports)
	web.Router.POST("/reports/delete", web.deleteReports)
}

func (web *Web) showListReports(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	rml := web.ctl.ListReports()
	render(w, "web/static/templates/report-list.html", rml)
}

func (web *Web) showAddReport(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	render(w, "web/static/templates/add.html", nil)
}

func (web *Web) addReport(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	r.ParseForm()
	// link := r.Form.Get("link")
	http.Redirect(w, r, "/reports", http.StatusSeeOther)
}

func (web *Web) showCompareReports(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	http.Redirect(w, r, "/reports", http.StatusSeeOther)
}

func (web *Web) deleteReports(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	names := strings.Split(r.URL.Query().Get("reports"), ",")
	web.ctl.DeleteReports(names)
	http.Redirect(w, r, "/reports", http.StatusSeeOther)
}

func render(w http.ResponseWriter, tpl string, data interface{}) {
	t, err := template.ParseFiles(tpl)
	if err != nil {
		log.Panic(err)
	} else {
		t.Execute(w, data)
	}
}
