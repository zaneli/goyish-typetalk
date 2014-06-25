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

type Client struct {
  accessToken string
}

type endPoint struct {
  apiUrl string
  kind string
  apiName string
}

func NewClient() (*Client) {
  return &Client{}
}

func AuthedClient(accessToken string) (*Client) {
  client := &Client{}
  client.accessToken = accessToken
  return client
}

func (c *Client) get(e endPoint, params map[string] string, result interface{}) error {
  return c.callApi(c.endPoint(e), "GET", params, nil, true, result)
}

func (c *Client) post(e endPoint, params map[string] string, filePath *string, auth bool, result interface{}) error {
  return c.callApi(c.endPoint(e), "POST", params, filePath, auth, result)
}

func (c *Client) put(e endPoint, params map[string] string, result interface{}) error {
  return c.callApi(c.endPoint(e), "PUT", params, nil, true, result)
}

func (c *Client) delete(e endPoint, params map[string] string, result interface{}) error {
  return c.callApi(c.endPoint(e), "DELETE", params, nil, true, result)
}

func (c *Client) endPoint(e endPoint) string {
  if e.apiUrl == "" {
    e.apiUrl = "https://typetalk.in/%s/%s"
  }
  if e.kind == "" {
    e.kind = "api/v1"
  }
  return fmt.Sprintf(e.apiUrl, e.kind, e.apiName)
}

func (c *Client) callApi(endPoint string, method string, params map[string] string, filePath *string, auth bool, result interface{}) error {
  var req *http.Request
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
    req, err = http.NewRequest(method, fmt.Sprint(u), nil)
    if err != nil {
      return err
    }
  } else if method == "POST" {
    if filePath != nil {
      req, err = createUploadReq(endPoint, params, *filePath, result)
      if err != nil {
        return err
      }
    } else {
      req, err = http.NewRequest(method, endPoint, strings.NewReader(createValues(params).Encode()))
      if err != nil {
        return err
      }
      req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
    }
  } else if method == "PUT" || method == "DELETE" {
    req, err = http.NewRequest(method, endPoint, strings.NewReader(createValues(params).Encode()))
    if err != nil {
      return err
    }
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
  } else {
    return fmt.Errorf("Invalid http method: %s", method)
  }
  if auth {
    req.Header.Set("Authorization", "Bearer " + c.accessToken)
  }
  res, err := http.DefaultClient.Do(req)
  if err != nil {
    return err
  }
  defer res.Body.Close()

  if res.StatusCode != http.StatusOK {
    authErrInfo := res.Header["Www-Authenticate"]
    if len(authErrInfo) > 0 {
      return fmt.Errorf("Invalid status: %s, Www-Authenticate: %s", res.Status, authErrInfo[0])
    }
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

func createUploadReq(endPoint string, params map[string] string, filePath string, result interface{}) (*http.Request, error) {
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
  req.Header.Set("Content-Type", writer.FormDataContentType())
  return req, nil
}

func createValues(params map[string] string) url.Values {
  values := map[string] []string{}
  for k, v := range params {
    values[k] = []string{v}
  }
  return values
}
