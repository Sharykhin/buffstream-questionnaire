package controller

import (
	"fmt"
	"net/http"
	"strconv"
)

func queryParamAsInt64(r *http.Request, param string, def int64) (int64, error) {
	value := r.URL.Query().Get(param)
	if value == "" {
		return def, nil
	}

	i, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("could not convert query parameter %s into int: %v ", value, err)
	}

	return i, nil
}
