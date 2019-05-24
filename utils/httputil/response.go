package httputil

import (
	"encoding/json"
	"github.com/risoll/tokosidia/constants"
	"net/http"
	"time"
)

type (
	Response struct {
		Data        interface{} `json:"data"`
		Errors      []string    `json:"errors"`
		startTime   time.Time
		ProcessTime float64 `json:"process_time"`
	}
)

func NewResponse() Response {
	return Response{
		startTime: time.Now(),
	}
}

func (r *Response) ToJSON() []byte {
	res, _ := json.Marshal(r)
	return res
}

func (r *Response) WriteOK(w http.ResponseWriter, data interface{}) {
	r.ProcessTime = time.Since(r.startTime).Seconds()
	r.Data = data
	w.Header().Set(constants.HeaderContentTypeKey, constants.HeaderContentTypeJSON)
	w.WriteHeader(http.StatusOK)
	w.Write(r.ToJSON())
}

func (r *Response) WriteError(w http.ResponseWriter, httpStatus int, errors []string) {
	r.ProcessTime = time.Since(r.startTime).Seconds()
	r.Errors = errors
	w.Header().Set(constants.HeaderContentTypeKey, constants.HeaderContentTypeJSON)
	if httpStatus == 0 {
		httpStatus = http.StatusBadRequest
	}
	w.WriteHeader(httpStatus)
	w.Write(r.ToJSON())
}
