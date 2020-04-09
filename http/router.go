package http

import (
	"net/http"

	"github.com/gorilla/mux"

	"Sharykhin/buffstream-questionnaire/di"
	questionController "Sharykhin/buffstream-questionnaire/http/controller/question"
	streamController "Sharykhin/buffstream-questionnaire/http/controller/stream"
	"Sharykhin/buffstream-questionnaire/http/middleware"
)

func router() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/_healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("OK"))
	}).Methods("GET")

	v1 := r.PathPrefix("/v1").Subrouter()
	v1.Use(middleware.JsonContentType)

	v1.HandleFunc("/streams", func(w http.ResponseWriter, r *http.Request) {
		streamController.List(di.StreamService, di.QuestionService, w, r)
	}).Methods("GET")

	v1.HandleFunc("/streams", func(w http.ResponseWriter, r *http.Request) {
		streamController.Create(di.StreamService, w, r)
	}).Methods("POST")

	v1.HandleFunc("/questions/{ID}", func(w http.ResponseWriter, r *http.Request) {
		questionController.GetByIdentifier(di.QuestionService, w, r)
	}).Methods("GET")

	return r
}
