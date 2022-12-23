package logic

import (
	"context"
	"google.golang.org/api/idtoken"
	"log"
)

func GoogleId(audience string) string {
	ctx := context.Background()

	ts, err := idtoken.NewTokenSource(ctx, audience)
	if err != nil {
		log.Fatal(err)
	}
	token, err := ts.Token()
	if err != nil {
		log.Fatal(err)
	}

	return "Bearer " + token.AccessToken
}
