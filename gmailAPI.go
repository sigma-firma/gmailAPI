// package gmailAPI is used in applications and libraries that access the gmail
// api. Most of this code was lightly modified from code found on:
// https://developers.google.com/gmail/api/quickstart/go
package gmailAPI

import (
	"context"
	"encoding/json"
	"fmt"

	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	gmail "google.golang.org/api/gmail/v1"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func ConnectToService(ctx context.Context, scope ...string) *gmail.Service {
	b, err := os.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	config, err := google.ConfigFromJSON(b, scope...)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}

	cacheFile, err := newTokenizer("./credentials")
	if err != nil {
		log.Fatalf("Unable to get path to cached credential file. %v", err)
	}

	token, err := tokenFromFile(cacheFile)
	if err != nil {
		token = getTokenFromWeb(config)
		saveToken(cacheFile, token)
	}
	srv, err := gmail.New(config.Client(ctx, token))
	if err != nil {
		log.Fatalf("Unable to retrieve gmail Client %v", err)
	}
	return srv
}

// Retrieve a token, saves the token, then returns the generated client.
func getClient(config *oauth2.Config) *http.Client {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	tokFile := "token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}

// newTokenizer returns a new token and generates credential file path and
// returns the generated credential path/filename along with any errors.
func newTokenizer(tokenCacheDir string) (string, error) {
	err := os.MkdirAll(tokenCacheDir, 0700)
	return filepath.Join(tokenCacheDir, url.QueryEscape("gmail-go-quickstart.json")),
		err
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	err = json.NewEncoder(f).Encode(token)
	if err != nil {
		log.Println(err)
	}
}
