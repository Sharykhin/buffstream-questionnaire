package controller

import (
	"Sharykhin/buffstream-questionnaire/domains/question/application/service"
	appErrors "Sharykhin/buffstream-questionnaire/errors"
	"Sharykhin/buffstream-questionnaire/http/response"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
)

func GetQuestionByIdentifier(questionSrv service.QuestionService, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID := vars["ID"]

	question, err := questionSrv.GetOneByUUID(r.Context(), ID)
	if err != nil {
		if errors.Is(err, appErrors.ResourceNotFound) {
			response.NotFound(w, errors.New("Question with such id was not found"))
		} else {
			response.ServerError(w, err)
		}

		return
	}

	response.OK(w, response.Data{"Questions": question}, nil)
}
