package methodoverride

import "net/http"

type XHTTPMethodOverrideHandler struct {
	SubHandler http.Handler
}

func (x XHTTPMethodOverrideHandler) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request,
) {
	if method := r.Header.Get("X-HTTP-Method-Override"); method != "" {
		r.Method = method
	}
	x.SubHandler.ServeHTTP(w, r)
}
