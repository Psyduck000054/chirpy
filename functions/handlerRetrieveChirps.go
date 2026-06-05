package functions

import (
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func (cfg *ApiConfig) HandlerRetrieveChirps(w http.ResponseWriter, r *http.Request) {

	type ChirpStruct struct {
		Id        uuid.UUID `json:"id"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		Body      string    `json:"body"`
		UserID    uuid.UUID `json:"user_id"`
	}

	type Output struct {
		Body []ChirpStruct `json:"body"`
	}

	chirpList, err := cfg.Queries.RetrieveAllChirps(r.Context())
	chirpStruct0 := ChirpStruct{}
	output0 := Output{}

	if err != nil {
		RespondWithError(w, 400, fmt.Sprintf("something went wrong: %s\n", err))
		return
	}

	for _, chirp := range chirpList {
		chirpStruct0.Id = chirp.ID
		chirpStruct0.CreatedAt = chirp.CreatedAt
		chirpStruct0.UpdatedAt = chirp.UpdatedAt
		chirpStruct0.Body = chirp.Body
		chirpStruct0.UserID = chirp.UserID
		output0.Body = append(output0.Body, chirpStruct0)
	}

	RespondWithJSON(w, 200, output0.Body)
}
