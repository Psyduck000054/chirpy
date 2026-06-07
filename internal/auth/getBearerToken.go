package auth

import (
	"fmt"
	"net/http"
	"strings"
)

func GetBearerToken(headers http.Header) (string, error) {
	bearerString := headers.Get("Authorization")

	if bearerString == "" {
		return "", fmt.Errorf("No Authentication header found\n")
	} else {
		if strings.HasPrefix(bearerString, "Bearer ") {
			bearerString = strings.TrimPrefix(bearerString, "Bearer ")
			return bearerString, nil
		} else {
			return "", fmt.Errorf("No token string found in Authentication header\n")
		}
	}
}
