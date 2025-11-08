package response

import (
	"encoding/json"
	"net/http"
)

func JsonEncoder(w http.ResponseWriter, res any, status int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(res)
}
