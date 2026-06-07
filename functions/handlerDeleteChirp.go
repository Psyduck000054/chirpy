package functions

import (
	"net/http"

	"github.com/Psyduck000054/chirpy/internal/auth"
	"github.com/google/uuid"
)

func (cfg *ApiConfig) HandlerDeleteChirp(w http.ResponseWriter, r *http.Request) {
	chirpIDString := r.PathValue("chirpID")
	chirpID, err := uuid.Parse(chirpIDString)
	if err != nil {
		RespondWithError(w, 400, "Couldn't extract chirp ID")
		return
	}

	chirp, err := cfg.Queries.GetChirp(r.Context(), chirpID)
	if err != nil {
		RespondWithError(w, 404, "Couldn't find chirp")
		return
	}

	extractedTokenFromRequest, err := ExtractTokenFromRequest(w, r)
	if err != nil {
		RespondWithError(w, 401, "Couldn't extract token")
		return
	}

	userID, err := auth.ValidateJWT(extractedTokenFromRequest, cfg.SecretJWT)
	if err != nil {
		RespondWithError(w, 401, "Couldn't validate JWT token")
		return
	}

	if chirp.UserID != userID {
		RespondWithError(w, 403, "You're not the logged in user")
		return
	}

	err = cfg.Queries.DeleteChirp(r.Context(), chirpID)
	if err != nil {
		RespondWithError(w, 500, "Couldn't delete chirp")
		return
	}

	w.WriteHeader(204)
}
