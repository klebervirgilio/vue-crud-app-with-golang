package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/klebervirgilio/vue-crud-app-with-golang/pkg/core"
	"github.com/klebervirgilio/vue-crud-app-with-golang/pkg/kudo"
	jwtverifier "github.com/okta/okta-jwt-verifier-golang"
	"github.com/rs/cors"
)

type Service struct {
	repo   core.Repository
	Router http.Handler
}

func (s Service) Index(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	accessToken := r.Header["Authorization"]
	parts := strings.Split(accessToken[0], " ")
	jwt, _ := validateAccessToken(parts[1])
	service := kudo.NewService(s.repo, jwt.Claims["sub"].(string))
	kudos, err := service.GetKudos()
	fmt.Println(kudos)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(kudos)
}

func NewService(repo core.Repository) Service {
	service := Service{
		repo: repo,
	}

	router := httprouter.New()
	router.GET("/kudos", service.Index)

	corsConfig := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"*"},
		Debug:            true,
	})
	service.Router = corsConfig.Handler(router)
	return service
	// router.POST("/kudos", Create)
	// router.PUT("/kudo/:id", Update)
	// router.DELETE("/kudo/:id", Delete)
}

func validateAccessToken(accessToken string) (*jwtverifier.Jwt, error) {
	jwtVerifierSetup := jwtverifier.JwtVerifier{
		Issuer:           "https://dev-509836.oktapreview.com/oauth2/default",
		ClaimsToValidate: map[string]string{"aud": "api://default", "cid": "0oagcbm1o6GTTB9Da0h7"},
	}
	verifier := jwtVerifierSetup.New()
	return verifier.VerifyIdToken(accessToken)
}
