package handlers

import (
	"encoding/json"
	"net/http"
	"table-ms/pkg/mocks"
)

func GetTablesByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(mocks.Table)
}
