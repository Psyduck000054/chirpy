package functions

import "net/http"

func (cfg *ApiConfig) HandlerRevoke(w http.ResponseWriter, r *http.Request) {
	extractedTokenFromRequest, err := ExtractTokenFromRequest(w, r)
	if err != nil {
		RespondWithError(w, 401, "Can't extract Token from request")
		return
	}

	token0, err := cfg.Queries.GetRefreshTokenByToken(r.Context(), extractedTokenFromRequest)
	if err != nil {
		RespondWithError(w, 401, "Can't find Token in database")
		return
	}

	cfg.Queries.RevokeToken(r.Context(), token0.Token)
	w.WriteHeader(204)
}
