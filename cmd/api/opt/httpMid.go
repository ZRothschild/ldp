package opt

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func logRequestBody(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, fmt.Sprintf("grpc server read request body err %+v", err), http.StatusBadRequest)
			return
		}
		clonedR := r.Clone(r.Context())
		clonedR.Body = io.NopCloser(bytes.NewReader(body))
		h.ServeHTTP(w, clonedR)
	})
}
