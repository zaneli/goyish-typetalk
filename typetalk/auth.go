package typetalk

import (
  "strings"
  "time"
)

type Scope string

const (
  My Scope = "my"
  TopicRead Scope = "topic.read"
  TopicPost Scope = "topic.post"
)

type AuthClient struct {
  clientId string
  clientSecret string
  scopes []Scope
  AccessToken string
  RefreshToken string
  ExpireAt time.Time
}

func NewAuthClient(clientId string, clientSecret string, scopes ...Scope) (*AuthClient) {
  client := &AuthClient{}
  client.clientId = clientId
  client.clientSecret = clientSecret
  client.scopes = scopes
  return client
}

func (c *AuthClient) GetAccessToken() (string, error) {
  var scopes []string
  scopes = make([]string, len(c.scopes))
  for i := 0; i < len(c.scopes); i++ {
    scopes[i] = string(c.scopes[i])
  }
  if err := c.authorize(
    "access_token",
    map[string] string{
      "client_id":     c.clientId,
      "client_secret": c.clientSecret,
      "grant_type":    "client_credentials",
      "scope":         strings.Join(scopes, ",")}); err != nil {
    return "", err
  }
  return c.AccessToken, nil
}

func (c *AuthClient) UpdateToken() (string, error) {
  if err := c.authorize(
    "access_token",
    map[string] string{
      "client_id":     c.clientId,
      "client_secret": c.clientSecret,
      "grant_type":    "refresh_token",
      "refresh_token": c.RefreshToken}); err != nil {
    return "", err
  }
  return c.AccessToken, nil
}

func (c *AuthClient) authorize(apiName string, params map[string] string) error {
  var result struct {
    AccessToken  string `json:"access_token"`
    TokenType    string `json:"token_type"`
    ExpiresIn    int    `json:"expires_in"`
    RefreshToken string `json:"refresh_token"`
  }

  client := &Client{"", "oauth2"}
  if err := client.post(apiName, params, nil, false, &result); err != nil {
    return err
  }

  c.AccessToken = result.AccessToken
  c.RefreshToken = result.RefreshToken
  c.ExpireAt = time.Now().Add(time.Duration(result.ExpiresIn) * time.Second)
  return nil
}
