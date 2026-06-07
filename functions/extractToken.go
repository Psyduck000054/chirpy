package functions

import (
	"fmt"
	"net/http"
	"strings"
)

func ExtractTokenFromRequest(w http.ResponseWriter, r *http.Request) (string, error) {
	bearerString := r.Header.Get("Authorization")
	if bearerString == "" {
		return bearerString, fmt.Errorf("No authentication header found")
	} else {
		trimmedString, ok := strings.CutPrefix(bearerString, "Bearer ")
		if !ok {
			return "", fmt.Errorf("Can't extract token")
		} else {
			return trimmedString, nil
		}
	}
}
