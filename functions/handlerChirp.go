package functions

import (
	"encoding/json"
	"fmt"
	"net/http"
	"slices"
	"strings"
)

func (cfg *ApiConfig) HandlerChirp(w http.ResponseWriter, r *http.Request) {
	type ChirpStruct struct {
		Chirp string `json:"body"`
	}

	type ReturnCensoredChirp struct {
		CensoredChirp string `json:"cleaned_body"`
	}

	bannedWords := []string{
		"kerfuffle",
		"sharbert",
		"fornax",
	}

	censoredResp := new(ReturnCensoredChirp)

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
		wordList := strings.Split(Chirps.Chirp, " ")
		for index, word := range wordList {
			if slices.Contains(bannedWords, strings.ToLower(word)) {
				wordList[index] = "****"
			}
		}

		censoredString := strings.Join(wordList, " ")
		censoredResp.CensoredChirp = censoredString

		RespondWithJSON(w, 200, censoredResp)
	}
}
