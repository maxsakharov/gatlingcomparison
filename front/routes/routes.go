package routes

import (
	"net/http"

	"github.com/maxsakharov/gatlingcomparison/front/controller"
	"github.com/maxsakharov/gatlingcomparison/storage"

	"github.com/julienschmidt/httprouter"
)

func InitRoutes(s storage.Storage) http.Handler {
	router := httprouter.New()

	lc := controller.ListController{Storage: s}
	dc := controller.DeleeteController{Storage: s}
	cc := controller.CompareController{Storage: s}
	ac := controller.AddController{Storage: s}

	router.GET("/reports", lc.ListReports)

	router.POST("/reports/delete", dc.DeleteReports)
	router.POST("/reports/compare", cc.CompareReports)

	router.GET("/reports/add", ac.ShowAddReport)
	router.POST("/reports/add", ac.AddReport)

	return router
}
