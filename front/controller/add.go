package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/maxsakharov/gatlingcomparison/front/views"
	"github.com/maxsakharov/gatlingcomparison/storage"
)

type AddController struct {
	Storage storage.Storage
}

func (ac *AddController) ShowAddReport(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	views.AddReport(w)
}

func (ac *AddController) AddReport(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	r.ParseForm()
	link := r.Form.Get("link")
	http.Redirect(w, r, "/reports", http.StatusSeeOther)
}
