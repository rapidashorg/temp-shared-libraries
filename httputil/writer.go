package httputil

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/rapidashorg/temp-shared-libraries/errors"
	"github.com/rapidashorg/temp-shared-libraries/log"
)

type response struct {
	Header headerResponse `json:"header"`
	Data   interface{}    `json:"data"`
}

type headerResponse struct {
	ProcessTime float64  `json:"process_time"`
	IsSuccess   bool     `json:"is_success"`
	ErrorCode   *int     `json:"error_code,omitempty"`
	Messages    []string `json:"messages,omitempty"`
}

func (h *httpUtil) marshalJSONResponse(processTime float64, messages []string, data interface{}, statusCode int) ([]byte, error) {
	header := headerResponse{
		ProcessTime: processTime,
		Messages:    messages,
	}

	// this assume 2xx status code is success, and the rest is not
	// so statusCode = 2xx will set header.is_success to true, otherwise false
	if statusCode/100 == 2 {
		header.IsSuccess = true
	} else {
		header.IsSuccess = false
		header.ErrorCode = &statusCode
	}

	resp := response{
		header,
		data,
	}

	bs, err := json.Marshal(&resp)
	if err != nil {
		return nil, err
	}

	return bs, nil
}

func (h *httpUtil) writeJSONResponse(w http.ResponseWriter, r *http.Request, messages []string, statusCode int, data interface{}, processTime float64) {
	resp, err := h.marshalJSONResponse(processTime, messages, data, statusCode)
	if err != nil {
		log.Error(r.Context(), err.Error())
		h.WriteInternalServerError(w, processTime)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(resp)
}

func (h *httpUtil) WriteResponse(w http.ResponseWriter, r *http.Request, startTime time.Time, data interface{}) {
	processTime := time.Since(startTime).Seconds()
	h.writeJSONResponse(w, r, nil, http.StatusOK, data, processTime)
}

// WriteErrorResponse writes error response from given error/ErrorResponse
func (h *httpUtil) WriteErrorResponse(w http.ResponseWriter, r *http.Request, err *errors.ErrorWrapper, startTime time.Time, errorContext string) {
	processTime := time.Since(startTime).Seconds()

	ctx := r.Context()
	log.ErrorWrapper(ctx, err, errorContext)
	requestID := log.GetRequestID(ctx)

	messages := []string{
		fmt.Sprintf("%s <%s>", err.Error(), requestID),
		err.ActualError(),
		fmt.Sprintf("RequestID: %s", requestID),
	}

	h.writeJSONResponse(w, r, messages, err.HTTPErrorCode, nil, processTime)
}

func (h *httpUtil) WriteInternalServerError(w http.ResponseWriter, processTime float64) {
	err := errInternalServer.New()
	resp, _ := h.marshalJSONResponse(processTime, []string{err.Error()}, nil, err.Code)
	w.WriteHeader(err.HTTPErrorCode)
	w.Write(resp)
}
