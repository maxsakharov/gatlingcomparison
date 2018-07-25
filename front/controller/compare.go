package controller

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/maxsakharov/gatlingcomparison/storage"
)

type CompareController struct {
	Storage storage.Storage
}

func (cc *CompareController) CompareReports(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Println("test")
	http.Redirect(w, r, "/reports", http.StatusSeeOther)
}
