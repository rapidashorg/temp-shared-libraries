package httputil

import (
	"net/http"
	"time"

	"github.com/rapidashorg/temp-shared-libraries/errors"
)

type HTTPUtil interface {
	GetURLParam(r *http.Request, name string) string
	GetQueryParam(r *http.Request, name string) string
	ReadBody(r *http.Request) ([]byte, error)

	WriteResponse(w http.ResponseWriter, r *http.Request, startTime time.Time, data interface{})
	WriteErrorResponse(w http.ResponseWriter, r *http.Request, err *errors.ErrorWrapper, startTime time.Time, errorContext string)
	WriteInternalServerError(w http.ResponseWriter, processTime float64)
}

type httpUtil struct {
}

func New() HTTPUtil {
	return &httpUtil{}
}
