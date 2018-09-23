package rest

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/klebervirgilio/vue-crud-app-with-golang/pkg/core"
	"github.com/klebervirgilio/vue-crud-app-with-golang/pkg/kudo"
)

func index(repo core.Repository, w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	userToken := w.Header().Get("user-token")
	service := kudo.NewService(repo, userToken)
	kudos, err := service.GetKudos()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(kudos)
}

func Handler(repo core.Repository) {
	router := httprouter.New()
	router.GET("/kudos", func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		index(repo, w, r, params)
	})
	// router.POST("/kudos", Create)
	// router.PUT("/kudo/:id", Update)
	// router.DELETE("/kudo/:id", Delete)
}
