package functions

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func RespondWithError(w http.ResponseWriter, code int, msg string) {
	errorStruct := make(map[string]string)
	errorStruct["error"] = msg

	RespondWithJSON(w, code, errorStruct)
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		RespondWithError(w, 500, fmt.Sprint(err))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}
