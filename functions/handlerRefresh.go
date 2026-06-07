package functions

import (
	"net/http"
	"time"

	"github.com/Psyduck000054/chirpy/internal/auth"
)

func (cfg *ApiConfig) HandlerRefresh(w http.ResponseWriter, r *http.Request) {
	type output0 struct {
		AccessToken string `json:"token"`
	}

	extractedTokenFromRequest, err := ExtractTokenFromRequest(w, r)
	if err != nil {
		RespondWithError(w, 401, "Can't extract token")
		return
	}

	token, err := cfg.Queries.GetRefreshTokenByToken(r.Context(), extractedTokenFromRequest)
	if err != nil {
		RespondWithError(w, 401, "No matching token found in database")
		return
	}

	if token.ExpiresAt.Before(time.Now()) {
		RespondWithError(w, 401, "Token expired")
		return
	}

	if token.RevokedAt.Valid == true {
		RespondWithError(w, 401, "Token revoked")
		return
	}

	newAT, err := auth.MakeJWT(token.UserID, cfg.SecretJWT, 3600*time.Second)
	if err != nil {
		RespondWithError(w, 401, "Failed to generate a new Access Token")
		return
	}

	RespondWithJSON(w, 200, output0{
		AccessToken: newAT,
	})
}
