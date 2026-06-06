package functions

import (
	"encoding/json"
	"net/http"

	"github.com/Psyduck000054/chirpy/internal/auth"
)

func (cfg *ApiConfig) HandlerLogin(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Password string `json:"password"`
		Email    string `json:"email"`
	}

	type response struct {
		User
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		RespondWithError(w, 400, "Couldn't decode parameters")
		return
	}

	user, err := cfg.Queries.GetUserByEmail(r.Context(), params.Email)
	if err != nil {
		RespondWithError(w, 401, "Incorrect email")
		return
	}

	match, err := auth.CheckPasswordHash(params.Password, user.HashedPassword)
	if err != nil || !match {
		RespondWithError(w, 401, "Incorrect password")
		return
	}

	RespondWithJSON(w, http.StatusOK, response{
		User: User{
			Id:        user.ID,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		},
	})
}
