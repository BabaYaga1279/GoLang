package httpserver

import "net/http"

func Test() {
	http.ListenAndServe(":0808", nil)
}
