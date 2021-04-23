// Package google provides you access to Google's OAuth2
// infrastructure. The implementation is based on this blog post:
// http://skarlso.github.io/2016/06/12/google-signin-with-go/
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// User is a retrieved and authenticated user.
type GoogleUser struct {
	Sub           string `json:"sub"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Profile       string `json:"profile"`
	Picture       string `json:"picture"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Gender        string `json:"gender"`
	Hd            string `json:"hd"`
}

type token struct {
	oauth2.Token
}

// Access returns the access token.
func (t *token) Access() string {
	return t.AccessToken
}

func (t *token) Refresh() string {
	return t.RefreshToken
}

func (t *token) Expired() bool {
	if t == nil {
		return true
	}
	return !t.Token.Valid()
}

func (t *token) ExpiryTime() time.Time {
	return t.Expiry
}

func unmarshallToken(s sessions.Session) *token {
	if s.Get("token") == nil {
		return nil
	}
	data := s.Get("token").([]byte)
	var tk oauth2.Token
	json.Unmarshal(data, &tk)
	return &token{tk}
}

func Auth(clientSecret string, clientID string, scopes []string) gin.HandlerFunc {
	conf := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  "https://detective.ninja/auth",
		Scopes:       scopes,
		Endpoint:     google.Endpoint,
	}

	return func(ctx *gin.Context) {
		url := ctx.Request.URL
		s := sessions.Default(ctx)
		t := unmarshallToken(s)
		state := url.Path
		if ctx.Query("state") != "" {
			state = ctx.Query("state")
		}

		if strings.HasPrefix(url.Path, "/static/") || strings.HasPrefix(url.Path, "/x/") || strings.HasPrefix(url.Path, "/d/") || url.Path == "/" {
			return
		}
		if ctx.Query("code") != "" {
			tok, err := conf.Exchange(oauth2.NoContext, ctx.Query("code"), oauth2.AccessTypeOffline, oauth2.ApprovalForce)
			if err != nil {
				ctx.AbortWithError(http.StatusBadRequest, err)
				return
			}

			client := conf.Client(oauth2.NoContext, tok)
			email, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
			if err != nil {
				ctx.AbortWithError(http.StatusBadRequest, err)
				return
			}

			defer email.Body.Close()
			data, err := ioutil.ReadAll(email.Body)
			if err != nil {
				log.Printf("[Gin-OAuth] Could not read Body: %s", err)
				ctx.AbortWithError(http.StatusInternalServerError, err)
				return
			}

			var user GoogleUser
			err = json.Unmarshal(data, &user)
			if err != nil {
				log.Printf("[Gin-OAuth] Unmarshal userinfo failed: %s", err)
				ctx.AbortWithError(http.StatusInternalServerError, err)
				return
			}

			val, _ := json.Marshal(tok)
			s.Set("token", val)
			s.Set("email", user.Email)
			s.Set("name", user.Name)
			s.Set("photo", user.Picture)

			s.Save()

			log.Printf("redirecting tok: %v, state: %v", tok, state)
			ctx.Redirect(http.StatusFound, fmt.Sprintf("https://detective.ninja%s", state))
			ctx.Abort()
			return
		}

		if t != nil && t.Expired() {
			src := conf.TokenSource(ctx, &t.Token)
			newToken, err := src.Token()
			if err != nil {
				log.Printf("failed to renew token: old: %v, %v", t, err)
				s.Delete("token")
				s.Save()

				ctx.Redirect(http.StatusFound, conf.AuthCodeURL(state, oauth2.AccessTypeOffline, oauth2.ApprovalForce))
				ctx.Abort()
				return
			} else {
				log.Printf("successfully renewed token old: %v, new: %+v", t, newToken)
			}

			val, _ := json.Marshal(newToken)
			s.Set("token", val)
			s.Save()
			// dont get the email, just make sure they still have access
		} else if t == nil {
			ctx.Redirect(http.StatusFound, conf.AuthCodeURL(state, oauth2.AccessTypeOffline, oauth2.ApprovalForce))
			ctx.Abort()
			return
		}

		email := s.Get("email").(string)
		name := s.Get("name").(string)
		photo := s.Get("photo").(string)
		ctx.Set("email", email)
		ctx.Set("photo", photo)
		ctx.Set("name", name)

	}
}

func FakeAuth(clientSecret string, clientID string, scopes []string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		s := sessions.Default(ctx)
		if s.Get("email") == nil {
			names := []string{"alice@gmail.com", "bob@gmail.com", "jack@gmail.com", "john@gmail.com"}
			s.Set("email", names[rand.Intn(len(names))])
			s.Set("name", "jack doe")
			s.Set("photo", "https://thumbs.dreamstime.com/b/funny-cartoon-monster-face-vector-square-avatar-halloween-175916751.jpg")
			s.Save()
		}
		email := s.Get("email").(string)
		name := s.Get("name").(string)
		photo := s.Get("photo").(string)
		ctx.Set("email", email)
		ctx.Set("photo", photo)
		ctx.Set("name", name)
	}
}
