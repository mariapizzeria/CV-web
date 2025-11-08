package customErrors

import "net/http"

func GetContentError(w http.ResponseWriter) {
	http.Error(w, "Error to GET content", http.StatusNotFound)
}
