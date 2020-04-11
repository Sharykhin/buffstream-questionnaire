// +build !integration unit

package controller

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueryParamAsInt64(t *testing.T) {
	tt := []struct {
		name          string
		inKey         string
		inValue       string
		expectedKey   string
		expectedValue int64
		defaultValue  int64
	}{
		{
			name:          "Set limit",
			inKey:         "limit",
			inValue:       "20",
			expectedKey:   "limit",
			expectedValue: 20,
			defaultValue:  10,
		},
		{
			name:          "Return default value",
			inKey:         "",
			inValue:       "",
			expectedKey:   "limit",
			expectedValue: 10,
			defaultValue:  10,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			url := fmt.Sprintf("?%s=%s", tc.inKey, tc.inValue)
			req, err := http.NewRequest("GET", url, nil)
			if err != nil {
				t.Errorf("failed to create test http request: %v", err)
			}
			v, err := QueryParamAsInt64(req, tc.expectedKey, tc.defaultValue)
			assert.Equal(t, tc.expectedValue, v)
			assert.Nil(t, err)

		})
	}
}
