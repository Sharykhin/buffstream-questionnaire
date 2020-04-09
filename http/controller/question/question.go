package question

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nu7hatch/gouuid"

	"Sharykhin/buffstream-questionnaire/domains/question/application/service"
	appErrors "Sharykhin/buffstream-questionnaire/errors"
	"Sharykhin/buffstream-questionnaire/http/response"
)

// GetByIdentifier returns question with associated answers
func GetByIdentifier(questionSrv service.QuestionService, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID := vars["ID"]

	err := validateID(ID)
	if err != nil {
		response.BadRequest(w, errors.New("provided id is not valid"))
		return
	}

	question, err := questionSrv.GetOneByID(r.Context(), ID)
	if err != nil {
		if errors.Is(err, appErrors.ResourceNotFound) {
			response.NotFound(w, errors.New("question with such id was not found"))
		} else {
			response.ServerError(w, err)
		}

		return
	}

	response.OK(w, response.Data{"Questions": question}, nil)
}

func validateID(ID string) error {
	_, err := uuid.ParseHex(ID)
	if err != nil {
		return err
	}

	return nil
}
