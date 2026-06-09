package functions

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Psyduck000054/chirpy/internal/auth"
	"github.com/Psyduck000054/chirpy/internal/database"
)

func (cfg *ApiConfig) HandlerUpdateUser(w http.ResponseWriter, r *http.Request) {
	tokenExtractedFromRequest, err := ExtractTokenFromRequest(w, r)
	if err != nil {
		RespondWithError(w, 401, "Couldn't extract Token from request")
		return
	}

	type Input0 struct {
		Password string `json:"password"`
		Email    string `json:"email"`
	}

	decoder := json.NewDecoder(r.Body)
	params := Input0{}

	err = decoder.Decode(&params)
	if err != nil {
		RespondWithError(w, 400, fmt.Sprintln("Couldn't decode parameters", err))
		return
	}

	hashedPassword, err := auth.HashPassword(params.Password)
	if err != nil {
		RespondWithError(w, 401, "Can't hash new password")
		return
	}

	userID, err := auth.ValidateJWT(tokenExtractedFromRequest, cfg.SecretJWT)
	if err != nil {
		RespondWithError(w, 401, "Can't validate JWT Token")
		return
	}

	queryParam := database.EditUserInfoParams{
		ID:             userID,
		HashedPassword: hashedPassword,
		Email:          params.Email,
	}

	user0, err := cfg.Queries.EditUserInfo(r.Context(), queryParam)
	if err != nil {
		RespondWithError(w, 401, "Can't edit user info")
		return
	}

	RespondWithJSON(w, 200, User{
		ID:          user0.ID,
		CreatedAt:   user0.CreatedAt,
		UpdatedAt:   user0.UpdatedAt,
		Email:       user0.Email,
		IsChirpyRed: user0.IsChirpyRed,
	})
}
