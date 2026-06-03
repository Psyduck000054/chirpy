package functions

import (
	"fmt"
	"net/http"
)

func (cfg *ApiConfig) HandlerReset(w http.ResponseWriter, r *http.Request) {
	err := cfg.Queries.DeleteAllUsers(r.Context())
	if err != nil {
		fmt.Println("Error in deleting the users database")
		return
	}

	if cfg.Platform != "dev" {
		w.WriteHeader(403)
		return
	}

	cfg.FileServerHits.Store(0)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hits reset to 0"))
}
