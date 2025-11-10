package customErrors

import "net/http"

func GetContentError(w http.ResponseWriter) {
	http.Error(w, "Error to GET content", http.StatusNotFound)
}

func ReadBodyError(w http.ResponseWriter) {
	http.Error(w, "Error to read body", http.StatusNotFound)
}

func CreateError(w http.ResponseWriter) {
	http.Error(w, "Error to create content", http.StatusInternalServerError)
}

func UpdateError(w http.ResponseWriter) {
	http.Error(w, "Error to update content", http.StatusInternalServerError)
}

func NotFound(w http.ResponseWriter) {
	http.Error(w, "Error to not found content", http.StatusNotFound)
}

func ServerError(w http.ResponseWriter) {
	http.Error(w, "Error to server error", http.StatusInternalServerError)
}
