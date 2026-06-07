package functions

import (
	"encoding/json"
	"fmt"
	"net/http"
	"slices"
	"strings"
	"time"

	"github.com/Psyduck000054/chirpy/internal/auth"
	"github.com/Psyduck000054/chirpy/internal/database"

	"github.com/google/uuid"
)

func (cfg *ApiConfig) HandlerCreateChirp(w http.ResponseWriter, r *http.Request) {
	type ChirpStruct struct {
		Chirp  string    `json:"body"`
		UserID uuid.UUID `json:"user_id"`
	}

	type ChirpOutput struct {
		Id        uuid.UUID `json:"id"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		Body      string    `json:"body"`
		UserID    uuid.UUID `json:"user_id"`
	}

	bannedWords := []string{
		"kerfuffle",
		"sharbert",
		"fornax",
	}

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

		// jwt check
		token, err := auth.GetBearerToken(r.Header)
		if err != nil {
			RespondWithError(w, 401, "Couldn't get Bearer token")
			return
		}

		validatedID, err := auth.ValidateJWT(token, cfg.SecretJWT)
		if err != nil {
			RespondWithError(w, 401, "Couldn't validate JWT token")
			return
		}

		params := database.SaveChirpParams{
			Body:   censoredString,
			UserID: validatedID,
		}

		chirp, err := cfg.Queries.SaveChirp(r.Context(), params)
		if err != nil {
			RespondWithError(w, 400, fmt.Sprintf("something went wrong: %s", err))
			return
		}

		jsonOutput := ChirpOutput{
			Id:        chirp.ID,
			CreatedAt: chirp.CreatedAt,
			UpdatedAt: chirp.UpdatedAt,
			Body:      chirp.Body,
			UserID:    chirp.UserID,
		}

		RespondWithJSON(w, 201, jsonOutput)
	}
}
