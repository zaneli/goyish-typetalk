package typetalk

import (
  "bytes"
  "encoding/json"
  "fmt"
  "io"
  "mime/multipart"
  "net/http"
  "net/url"
  "os"
  "path/filepath"
  "strings"
)

const (
  apiUrl = "https://typetalk.in/%s/%s"
)

type Client struct {
  accessToken string
  kind string
}

func NewClient(auth *AuthClient) (*Client) {
  client := &Client{}
  client.kind = "api/v1"
  client.accessToken = auth.AccessToken
  return client
}

func (c *Client) get(apiName string, params map[string] string, result interface{}) error {
  return c.callApi(apiName, "GET", params, nil, true, result)
}

func (c *Client) post(apiName string, params map[string] string, filePath *string, auth bool, result interface{}) error {
  return c.callApi(apiName, "POST", params, filePath, auth, result)
}

func (c *Client) put(apiName string, params map[string] string, result interface{}) error {
  return c.callApi(apiName, "PUT", params, nil, true, result)
}

func (c *Client) delete(apiName string, params map[string] string, result interface{}) error {
  return c.callApi(apiName, "DELETE", params, nil, true, result)
}

func (c *Client) callApi(apiName string, method string, params map[string] string, filePath *string, auth bool, result interface{}) error {
  if auth {
    params["access_token"] = c.accessToken
  }
  endPoint := fmt.Sprintf(apiUrl, c.kind, apiName)
  var res *http.Response
  var err error
  if method == "GET" {
    u, err := url.Parse(endPoint)
    if err != nil {
      return err
    }
    q := u.Query()
    for k, v := range params {
      q.Add(k, v)
    }
    u.RawQuery = q.Encode()
    res, err = http.Get(fmt.Sprint(u))
    if err != nil {
      return err
    }
  } else if method == "POST" {
    if filePath != nil {
      res, err = upload(endPoint, params, *filePath, result)
      if err != nil {
        return err
      }
    }
    res, err = http.PostForm(endPoint, createValues(params))
    if err != nil {
      return err
    }
  } else if method == "PUT" || method == "DELETE" {
    req, err := http.NewRequest(method, endPoint, strings.NewReader(createValues(params).Encode()))
    if err != nil {
      return err
    }
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
    res, err = http.DefaultClient.Do(req)
    if err != nil {
      return err
    }
  }
  defer res.Body.Close()

  if res.StatusCode != http.StatusOK {
    return fmt.Errorf("Invalid status: %s", res.Status)
  }
  if result == nil {
    return nil
  }
  if err = json.NewDecoder(res.Body).Decode(result); err != nil {
    return err
  }
  return nil
}

func upload(endPoint string, params map[string] string, filePath string, result interface{}) (*http.Response, error) {
  file, err := os.Open(filePath)
  if err != nil {
    return nil, err
  }
  defer file.Close()
  body := &bytes.Buffer{}
  writer := multipart.NewWriter(body)
  part, err := writer.CreateFormFile("file", filepath.Base(filePath))
  if err != nil {
    return nil, err
  }
  _, err = io.Copy(part, file)
  for k, v := range params {
    err = writer.WriteField(k, v)
    if err != nil {
      return nil, err
    }
  }
  err = writer.Close()
  if err != nil {
    return nil, err
  }
  req, err := http.NewRequest("POST", endPoint, body)
  if err != nil {
    return nil, err
  }
  req.Header.Add("Content-Type", writer.FormDataContentType())
  res, err := http.DefaultClient.Do(req)
  if err != nil {
    return nil, err
  }
  defer res.Body.Close()
  return res, nil
}

func createValues(params map[string] string) url.Values {
  values := map[string] []string{}
  for k, v := range params {
    values[k] = []string{v}
  }
  return values
}
