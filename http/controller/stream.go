package controller

import (
	"net/http"

	"Sharykhin/buffstream-questionnaire/domains/stream/application/service"
	"Sharykhin/buffstream-questionnaire/http/response"
)

// ListStreams returns list of streams with questions
func ListStreams(streamSrv service.StreamService, w http.ResponseWriter, r *http.Request) {
	limit, err := queryParamAsInt64(r, "limit", 10)
	if err != nil {
		response.BadRequest(w, err)
		return
	}

	offset, err := queryParamAsInt64(r, "offset", 0)
	if err != nil {
		response.BadRequest(w, err)
		return
	}

	l, t, _ := streamSrv.List(r.Context(), limit, offset)
	response.OK(w, response.Data{"Streams": l}, response.Meta{
		"Limit":  limit,
		"Offset": offset,
		"Total":  t,
	})
}
