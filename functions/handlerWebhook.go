package functions

import (
	"encoding/json"
	"net/http"

	"github.com/Psyduck000054/chirpy/internal/auth"
	"github.com/google/uuid"
)

func (cfg *ApiConfig) HandlerWebhook(w http.ResponseWriter, r *http.Request) {
	type data0 struct {
		UserID uuid.UUID `json:"user_id"`
	}

	type input0 struct {
		Event string `json:"event"`
		Data  data0  `json:"data"`
	}

	decoder := json.NewDecoder(r.Body)
	input := input0{}
	err := decoder.Decode(&input)

	if err != nil {
		RespondWithError(w, 401, "Can't decode user input")
		return
	}

	polkaAPIKey, err := auth.GetToken(r.Header, "ApiKey ")
	if err != nil {
		RespondWithError(w, 401, "Couldn't get Polka token")
		return
	}

	if polkaAPIKey != cfg.PolkaAPIKey {
		RespondWithError(w, 401, "Couldn't validate Polka token")
		return
	}

	if input.Event != "user.upgraded" {
		w.WriteHeader(204)
		return
	}

	_, err = cfg.Queries.AddChirpyRed(r.Context(), input.Data.UserID)
	if err != nil {
		RespondWithError(w, 404, "Can't find user in database")
		return
	}

	w.WriteHeader(204)
}
