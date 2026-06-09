package functions

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Psyduck000054/chirpy/internal/auth"
	"github.com/Psyduck000054/chirpy/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID `json:"id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Email        string    `json:"email"`
	Password     string    `json:"-"`
	AccessToken  string    `json:"token"`
	RefreshToken string    `json:"refresh_token"`
	IsChirpyRed  bool      `json:"is_chirpy_red"`
}

func (cfg *ApiConfig) HandlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type Input0 struct {
		Password string `json:"password"`
		Email    string `json:"email"`
	}

	type Output0 struct {
		User
	}

	decoder := json.NewDecoder(r.Body)
	params := Input0{}

	err := decoder.Decode(&params)
	if err != nil {
		RespondWithError(w, 400, fmt.Sprintln("Couldn't decode parameters", err))
		return
	}

	hashedPassword, err := auth.HashPassword(params.Password)
	if err != nil {
		RespondWithError(w, 400, fmt.Sprintln("Couldn't hash password", err))
		return
	}

	var usr database.User

	usr, err = cfg.Queries.CreateUser(r.Context(), database.CreateUserParams{
		Email:          params.Email,
		HashedPassword: hashedPassword,
	})

	if err != nil {
		RespondWithError(w, 400, fmt.Sprintln("Couldn't create user", err))
		return
	}

	RespondWithJSON(w, 201, Output0{
		User: User{
			ID:          usr.ID,
			CreatedAt:   usr.CreatedAt,
			UpdatedAt:   usr.UpdatedAt,
			Email:       usr.Email,
			IsChirpyRed: usr.IsChirpyRed,
		},
	})
}
