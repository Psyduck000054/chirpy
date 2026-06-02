package functions

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (cfg *ApiConfig) HandlerChirp(w http.ResponseWriter, r *http.Request) {
	type ChirpStruct struct {
		Chirp string `json:"body"`
	}

	type ReturnChirpValid struct {
		Check bool `json:"valid"`
	}

	validResp := new(ReturnChirpValid)

	decoder := json.NewDecoder(r.Body)
	Chirps := ChirpStruct{}
	err := decoder.Decode(&Chirps)

	if err != nil {
		RespondWithError(w, 400, fmt.Sprintf("something went wrong: %s", err))
		return
	}

	if len(Chirps.Chirp) > 140 {
		RespondWithError(w, 400, "Chirp is too long")
		return
	} else {
		validResp.Check = true
		RespondWithJSON(w, 200, validResp)
	}
}
