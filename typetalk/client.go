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

type Client interface {
	GetAccessToken(clientId string, clientSecret string, scope ...Scope) (*Auth, error)
	UpdateAccessToken(clientId string, clientSecret string, refreshToken string) (*Auth, error)

	GetMyProfile() (*Account, error)
	GetMyTopics() ([]TopicInfo, error)
	FavoriteTopic(topicId int) (*Topic, error)
	UnfavoriteTopic(topicId int) (*Topic, error)
	GetNotificationCount() (*Notifications, error)
	ReadNotification() (*OpenStatus, error)
	ReadMessagesInTopicApi(topicId int) *ReadMessagesInTopicApi
	ReadMessagesInTopic(topicId int) (*Unread, error)
	GetMentionListApi() *GetMentionListApi
	GetMentionList() ([]Mention, error)
	ReadMention(mentionId int) (*Mention, error)

	PostMessageApi(message string, topicId int) *PostMessageApi
	PostMessage(message string, topicId int) (*PostResult, error)
	UploadAttachmentFile(topicId int, filePath string) (*File, error)
	RemoveMessage(topicId int, postId int) (bool, error)
	LikeMessage(topicId int, postId int) (*Like, error)
	UnlikeMessage(topicId int, postId int) (*Like, error)

	GetTopicMessagesApi(topicId int) *GetTopicMessagesApi
	GetTopicMessages(topicId int) (*Messages, error)
	GetMessage(topicId int, postId int) (*Post, error)
}

type client struct {
	accessToken string
	apiUrl      string
}

type endPoint struct {
	kind    string
	apiName string
}

func NewClient() Client {
	return &client{"", "https://typetalk.in/%s/%s"}
}

func AuthedClient(accessToken string) Client {
	return &client{accessToken, "https://typetalk.in/%s/%s"}
}

func (c *client) get(e endPoint, params map[string]string, result interface{}) error {
	return c.callApi(c.endPoint(e), "GET", params, nil, true, result)
}

func (c *client) post(e endPoint, params map[string]string, filePath *string, auth bool, result interface{}) error {
	return c.callApi(c.endPoint(e), "POST", params, filePath, auth, result)
}

func (c *client) put(e endPoint, params map[string]string, result interface{}) error {
	return c.callApi(c.endPoint(e), "PUT", params, nil, true, result)
}

func (c *client) delete(e endPoint, params map[string]string, result interface{}) error {
	return c.callApi(c.endPoint(e), "DELETE", params, nil, true, result)
}

func (c *client) endPoint(e endPoint) string {
	if e.kind == "" {
		e.kind = "api/v1"
	}
	return fmt.Sprintf(c.apiUrl, e.kind, e.apiName)
}

func (c *client) callApi(endPoint string, method string, params map[string]string, filePath *string, auth bool, result interface{}) error {
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
		req.Header.Set("Authorization", "Bearer "+c.accessToken)
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

func createUploadReq(endPoint string, params map[string]string, filePath string, result interface{}) (*http.Request, error) {
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

func createValues(params map[string]string) url.Values {
	values := map[string][]string{}
	for k, v := range params {
		values[k] = []string{v}
	}
	return values
}
