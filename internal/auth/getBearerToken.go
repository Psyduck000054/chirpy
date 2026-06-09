package auth

import (
	"fmt"
	"net/http"
	"strings"
)

func GetToken(headers http.Header, strippedPrefix string) (string, error) {
	bearerString := headers.Get("Authorization")

	if bearerString == "" {
		return "", fmt.Errorf("No Authentication header found\n")
	} else {
		if strings.HasPrefix(bearerString, strippedPrefix) {
			bearerString = strings.TrimPrefix(bearerString, strippedPrefix)
			return bearerString, nil
		} else {
			return "", fmt.Errorf("No token string found in Authentication header\n")
		}
	}
}
