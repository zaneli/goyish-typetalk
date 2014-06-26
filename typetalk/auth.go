package typetalk

import (
	"strings"
	"time"
)

type Scope string

const (
	My        Scope = "my"
	TopicRead Scope = "topic.read"
	TopicPost Scope = "topic.post"
)

func (c *client) GetAccessToken(clientId string, clientSecret string, scope ...Scope) (*Auth, error) {
	var scopes []string
	scopes = make([]string, len(scope))
	for i := 0; i < len(scope); i++ {
		scopes[i] = string(scope[i])
	}
	auth, err := c.authorize(
		map[string]string{
			"client_id":     clientId,
			"client_secret": clientSecret,
			"grant_type":    "client_credentials",
			"scope":         strings.Join(scopes, ",")})
	if err != nil {
		return nil, err
	}
	return auth, nil
}

func (c *client) UpdateAccessToken(clientId string, clientSecret string, refreshToken string) (*Auth, error) {
	auth, err := c.authorize(
		map[string]string{
			"client_id":     clientId,
			"client_secret": clientSecret,
			"grant_type":    "refresh_token",
			"refresh_token": refreshToken})
	if err != nil {
		return nil, err
	}
	return auth, nil
}

func (c *client) authorize(params map[string]string) (*Auth, error) {
	var result struct {
		AccessToken  string `json:"access_token"`
		TokenType    string `json:"token_type"`
		ExpiresIn    int    `json:"expires_in"`
		RefreshToken string `json:"refresh_token"`
	}

	if err := c.post(endPoint{apiName: "access_token", kind: "oauth2"}, params, nil, false, &result); err != nil {
		return nil, err
	}

	auth := &Auth{}
	auth.AccessToken = result.AccessToken
	auth.TokenType = result.TokenType
	auth.RefreshToken = result.RefreshToken
	auth.ExpireAt = time.Now().Add(time.Duration(result.ExpiresIn) * time.Second)
	c.accessToken = auth.AccessToken
	return auth, nil
}
