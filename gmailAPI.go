// package gmailAPI is used in applications and libraries that access the gmail
// api. Most of this code was lightly modified from code found on:
// https://developers.google.com/gmail/api/quickstart/go
package gmailAPI

import (
	"context"
	"log"

	"google.golang.org/api/option"

	cg "github.com/hartsfield/connectToGoogleAPI"
)

type service interface {
	NewService(context.Context, option.ClientOption) ([]interface{}, error)
}

func ConnectToService(a *cg.Access, s service) []interface{} {
	srv, err := s.NewService(a.Context, option.WithHTTPClient(a.GetClient()))
	if err != nil {
		log.Fatalf("Unable to retrieve gmail Client %v", err)
	}
	return srv
}

// func refreshToken(config *oauth2.Config, token *oauth2.Token) *oauth2.Token {
// 	c := config.Client(context.Background(), token)
// 	res, err := c.Get("https: //oauth2.googleapis.com/" + token.RefreshToken)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	b, err := io.ReadAll(res.Body)
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	return &oauth2.Token{RefreshToken: string(b)}
// }
