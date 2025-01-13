package auth

import (
	"fmt"
	"os"

	"github.com/okta/okta-jwt-verifier-golang"
)

var JwtVerifier *jwtverifier.JwtVerifier

func InitVerifier()  {
	toValidate := map[string]string{}
	toValidate["aud"] = "api://default"

	jwtVerifierSetup := jwtverifier.JwtVerifier{
		Issuer:           os.Getenv("OKTA_OAUTH2_ISSUER"),
		ClaimsToValidate: toValidate,
	}

	verifier := jwtVerifierSetup.New()

	// Prefetch the JWKs and cache them
	_, err:= verifier.VerifyIdToken(os.Getenv("DUMMY_JWT"))

	if err != nil {
		fmt.Println("error: " + err.Error())
	}

	JwtVerifier = verifier


}
