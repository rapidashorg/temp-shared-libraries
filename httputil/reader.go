package httputil

import (
	"io/ioutil"
	"net/http"

	"github.com/go-chi/chi"
)

func (h *httpUtil) GetURLParam(r *http.Request, name string) string {
	return chi.URLParam(r, name)
}

func (h *httpUtil) GetQueryParam(r *http.Request, name string) string {
	return r.URL.Query().Get(name)
}

func (h *httpUtil) ReadBody(r *http.Request) ([]byte, error) {
	return ioutil.ReadAll(r.Body)
}
