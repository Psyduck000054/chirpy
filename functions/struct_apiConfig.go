package functions

import (
	"sync/atomic"

	"github.com/Psyduck000054/chirpy/internal/database"
)

type ApiConfig struct {
	FileServerHits atomic.Int32
	Queries        *database.Queries
	Platform       string
	SecretJWT      string
	PolkaAPIKey    string
}
