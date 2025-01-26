package util

import (
    "fmt"
    "github.com/golang-jwt/jwt/v4"
)

func GetEmailFromToken(tokenString string) (string, error) {
    // Parse the token without validating the signature
    token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
    if err != nil {
        return "", fmt.Errorf("failed to parse token: %v", err)
    }

    // Extract claims
    if claims, ok := token.Claims.(jwt.MapClaims); ok {
        // Retrieve the 'sub' claim
        if sub, ok := claims["sub"].(string); ok {
            return sub, nil
        }
        return "", fmt.Errorf("'sub' claim is not present in token")
    }

    return "", fmt.Errorf("invalid token claims")
}
