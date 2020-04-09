package stream

import (
	"net/http"
	"time"

	questionModel "Sharykhin/buffstream-questionnaire/domains/question/application/model"
	questionService "Sharykhin/buffstream-questionnaire/domains/question/application/service"
	streamModel "Sharykhin/buffstream-questionnaire/domains/stream/application/model"
	streamService "Sharykhin/buffstream-questionnaire/domains/stream/application/service"
	"Sharykhin/buffstream-questionnaire/http/controller"
	"Sharykhin/buffstream-questionnaire/http/response"
)

type (
	// StreamListItem is a stream view model that will be sent to a http client when it requests streams list
	ListItem struct {
		UUID      string     `json:"UUID"`
		Title     string     `json:"Title"`
		CreatedAt time.Time  `json:"CreatedAt"`
		UpdatedAt time.Time  `json:"UpdatedAt"`
		Questions []Question `json:"Questions"`
	}

	// Question is a related view model of a stream for a list
	Question struct {
		UUID string `json:"UUID"`
		Text string `json:"Text"`
	}
)

// List returns list of streams with questions
func List(
	streamSrv streamService.StreamService,
	questionSrv questionService.QuestionService,
	w http.ResponseWriter,
	r *http.Request,
) {
	limit, err := controller.QueryParamAsInt64(r, "limit", 10)
	if err != nil {
		response.BadRequest(w, err)
		return
	}

	offset, err := controller.QueryParamAsInt64(r, "offset", 0)
	if err != nil {
		response.BadRequest(w, err)
		return
	}

	streamList, total, err := streamSrv.List(r.Context(), limit, offset)
	if err != nil {
		response.ServerError(w, err)
		return
	}
	UUIDs := make([]string, len(streamList))
	for i, stream := range streamList {
		UUIDs[i] = stream.UUID
	}

	streamQuestions, err := questionSrv.GetAllByStreamIDs(r.Context(), UUIDs)
	if err != nil {
		response.ServerError(w, err)
		return
	}

	streams := createStreamViewModels(streamList, streamQuestions)

	response.OK(w, response.Data{"Streams": streams}, response.Meta{
		"Limit":  limit,
		"Offset": offset,
		"Total":  total,
	})
}

func createStreamViewModels(streamList []streamModel.Stream, streamQuestions questionModel.Streams) []ListItem {
	streams := make([]ListItem, len(streamList))

	for i, stream := range streamList {
		questionsLen := 0
		if _, ok := streamQuestions[stream.UUID]; ok {
			questionsLen = len(streamQuestions[stream.UUID])
		}
		s := ListItem{
			UUID:      stream.UUID,
			Title:     stream.Title,
			CreatedAt: stream.CreatedAt,
			UpdatedAt: stream.UpdatedAt,
			Questions: make([]Question, questionsLen),
		}
		if _, ok := streamQuestions[stream.UUID]; ok {
			for j, question := range streamQuestions[stream.UUID] {
				s.Questions[j] = Question{
					UUID: question.UUID,
					Text: question.Text,
				}
			}
		}

		streams[i] = s
	}

	return streams
}
