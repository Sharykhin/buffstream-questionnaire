package stream

import (
	"Sharykhin/buffstream-questionnaire/errors"
	"encoding/json"
	"fmt"
	"net/http"

	streamService "Sharykhin/buffstream-questionnaire/domains/stream/application/service"
	"Sharykhin/buffstream-questionnaire/http/response"
)

type (
	CreateStreamRequest struct {
		Title string `json:"Title"`
	}
)

// Create creates a new stream
func Create(
	streamSrv streamService.StreamService,
	w http.ResponseWriter,
	r *http.Request,
) {
	var req CreateStreamRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		response.BadRequest(w, fmt.Errorf("request payload is invalid: %v", err))
		return
	}

	defer errors.CheckDefferError(r.Body.Close)

	stream, err := streamSrv.Create(r.Context(), req.Title)
	if err != nil {
		response.ServerError(w, err)
		return
	}

	response.Created(w, response.Data{"Stream": stream}, nil)
}
