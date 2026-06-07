package functions

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/Psyduck000054/chirpy/internal/auth"
	"github.com/Psyduck000054/chirpy/internal/database"
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

	// access [jwt] token expire after 1 hour
	ATtimeoutInterval := time.Duration(1) * time.Hour
	RTTimeoutInterval := time.Duration(60*24) * time.Hour

	NewRefreshToken := auth.MakeRefreshToken()

	refreshTokenParams := database.CreateRefreshTokenParams{
		Token:     NewRefreshToken,
		UserID:    user.ID,
		ExpiresAt: time.Now().Add(RTTimeoutInterval),
		RevokedAt: sql.NullTime{},
	}

	refreshToken, err := cfg.Queries.CreateRefreshToken(r.Context(), refreshTokenParams)
	if err != nil {
		RespondWithError(w, 500, "Couldn't generate Refresh Token")
		return
	}

	token, err := auth.MakeJWT(user.ID, cfg.SecretJWT, ATtimeoutInterval)
	if err != nil {
		RespondWithError(w, 500, "Couldn't generate JWT Token")
		return
	}

	RespondWithJSON(w, http.StatusOK, response{
		User: User{
			Id:           user.ID,
			Email:        user.Email,
			CreatedAt:    user.CreatedAt,
			UpdatedAt:    user.UpdatedAt,
			AccessToken:  token,
			RefreshToken: refreshToken.Token,
		},
	})
}
