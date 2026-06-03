package functions

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Psyduck000054/chirpy/internal/database"
	"github.com/google/uuid"
)

func (cfg *ApiConfig) HandlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type emailInput struct {
		Email string `json:"email"`
	}

	type userOutput struct {
		Id        uuid.UUID `json:"id"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		Email     string    `json:"email"`
	}

	decoder := json.NewDecoder(r.Body)
	emailStruct0 := emailInput{}
	user0 := userOutput{}

	err := decoder.Decode(&emailStruct0)
	if err != nil {
		RespondWithError(w, 400, fmt.Sprintf("something went wrong: %s", err))
		return
	}

	var usr database.User

	usr, err = cfg.Queries.CreateUser(r.Context(), emailStruct0.Email)
	if err != nil {
		RespondWithError(w, 400, fmt.Sprintf("something went wrong: %s", err))
		return
	}

	user0.Id = usr.ID
	user0.CreatedAt = usr.CreatedAt
	user0.UpdatedAt = usr.UpdatedAt
	user0.Email = usr.Email

	RespondWithJSON(w, 201, user0)
}
