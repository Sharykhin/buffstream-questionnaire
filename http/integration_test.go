// +build integration !unit

package http

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/nu7hatch/gouuid"
	"github.com/stretchr/testify/assert"

	"Sharykhin/buffstream-questionnaire/http/controller/stream"
	"Sharykhin/buffstream-questionnaire/http/response"
)

type application struct {
	handler http.Handler
}

func (a *application) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.handler.ServeHTTP(w, r)
}

func TestGetHealthCheck(t *testing.T) {
	assert := assert.New(t)
	app := application{
		handler: router(),
	}

	req, err := http.NewRequest(http.MethodGet, "/_healthcheck", nil)
	if err != nil {
		t.Errorf("error creating request: %v", err)
	}

	w := httptest.NewRecorder()
	app.ServeHTTP(w, req)

	if want, got := http.StatusOK, w.Code; want != got {
		assert.Fail("expected status code: %v, got status code: %v", want, got)
	}

	bodyBytes, err := ioutil.ReadAll(w.Body)
	if err != nil {
		assert.Fail("failed to read body: %v", err)
	}
	bodyString := string(bodyBytes)
	assert.Equal("OK", bodyString)
}

func TestGetStreams(t *testing.T) {
	assert := assert.New(t)
	app := application{
		handler: router(),
	}

	req, err := http.NewRequest(http.MethodGet, "/v1/streams", nil)
	if err != nil {
		t.Errorf("error creating request: %v", err)
	}

	w := httptest.NewRecorder()
	app.ServeHTTP(w, req)

	if want, got := http.StatusOK, w.Code; want != got {
		assert.Fail("expected status code: %v, got status code: %v", want, got)
	}

	type expectedResponse struct {
		response.Response
		Data struct {
			Streams []stream.ListItem
		}
	}

	var resp expectedResponse

	err = json.NewDecoder(w.Body).Decode(&resp)
	if err != nil {
		assert.Fail("failed to read body: %v", err)
	}

	assert.Equal(http.StatusOK, resp.StatusCode)
	assert.False(resp.Error.Valid())
	assert.Contains(resp.Meta, "Limit")
	assert.NotNil(resp.Data.Streams)

	for _, streamItem := range resp.Data.Streams {
		_, err := uuid.ParseHex(streamItem.UUID)
		if err != nil {
			assert.Fail("expected valid uuid but got error: %v", err)
		}
		assert.NotEmpty(streamItem.Title)
		assert.NotEqual(time.Time{}, streamItem.CreatedAt)
		assert.NotEqual(time.Time{}, streamItem.UpdatedAt)
	}
}
