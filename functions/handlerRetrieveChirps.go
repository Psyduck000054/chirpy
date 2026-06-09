package functions

import (
	"net/http"
	"time"

	"github.com/Psyduck000054/chirpy/internal/database"
	"github.com/google/uuid"
)

func (cfg *ApiConfig) HandlerRetrieveChirps(w http.ResponseWriter, r *http.Request) {
	authorID := r.URL.Query().Get("author_id")
	var userID uuid.UUID
	var err error

	if authorID != "" {
		userID, err = uuid.Parse(authorID)
		if err != nil {
			RespondWithError(w, 401, "Invalid User ID")
		}
	} else {
		userID = uuid.Nil
	}

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

	var chirpList []database.Chirp
	if userID == uuid.Nil {
		chirpList, err = cfg.Queries.RetrieveAllChirps(r.Context())
	} else {
		chirpList, err = cfg.Queries.RetrieveChirpsByAuthorID(r.Context(), userID)
	}
	chirpStruct0 := ChirpStruct{}
	output0 := Output{}

	if err != nil {
		RespondWithError(w, 400, "Can't retrieve the list of chirps from database")
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
