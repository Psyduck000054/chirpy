package functions

import (
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func (cfg *ApiConfig) HandlerRetrieveChirp(w http.ResponseWriter, r *http.Request) {
	type ChirpStruct struct {
		Id        uuid.UUID `json:"id"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		Body      string    `json:"body"`
		UserID    uuid.UUID `json:"user_id"`
	}

	chirpIDString := r.PathValue("chirpID")
	chirpID, err := uuid.Parse(chirpIDString)
	if err != nil {
		RespondWithError(w, 404, fmt.Sprintln("Invalid Chirp ID"))
		return
	}

	dbChirp, err := cfg.Queries.GetChirp(r.Context(), chirpID)
	if err != nil {
		RespondWithError(w, 404, fmt.Sprintln("Couldn't get Chirp"))
		return
	}

	RespondWithJSON(w, 200, ChirpStruct{
		Id:        dbChirp.ID,
		CreatedAt: dbChirp.CreatedAt,
		UpdatedAt: dbChirp.UpdatedAt,
		UserID:    dbChirp.UserID,
		Body:      dbChirp.Body,
	})
}
