package typetalk

import (
	"fmt"
	. "gopkg.in/check.v1"
	"net/http"
	"testing"
	"time"
)

const port = "7890"

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type typetalkSuite struct{}

var _ = Suite(&typetalkSuite{})

func (s *typetalkSuite) SetUpSuite(c *C) {
	go startServer()
}

func newTestClient() Client {
	client := &client{}
	client.apiUrl = "http://localhost:" + port + "/%s/%s"
	return client
}

func toTime(year int, month time.Month, day, hour, min, sec int) time.Time {
	return time.Date(year, month, day, hour, min, sec, 0, time.UTC)
}

func startServer() {
	http.HandleFunc("/", dummyResponseHandler)
	http.ListenAndServe(":"+port, nil)
}

func dummyResponseHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	w.Header().Set("Content-Type", "application/json")
	switch path {
	case "/oauth2/access_token":
		fmt.Fprint(w, `{"access_token": "test_access_token",
		                "token_type": "Bearer",
		                "expires_in": 3600,
		                "refresh_token": "test_refresh_token"}`)
	case "/api/v1/profile":
		fmt.Fprint(w, `{"account":
		                {"id":3333,
		                 "name":"goyish",
		                 "fullName":"goyish account",
		                 "suggestion":"goyish",
		                 "imageUrl":"https://typetalk.in/accounts/3333/profile_image.png?t=9999999999999",
		                 "createdAt":"2014-01-11T11:11:11Z",
		                 "updatedAt":"2014-02-22T12:22:22Z"}}`)
	case "/api/v1/topics":
		fmt.Fprint(w, `{"topics":[{
		                "topic":
		                 {"id":1111,
		                  "name":"テストトピック",
		                  "suggestion":"テストトピック",
		                  "description":null,
		                  "createdAt":"2014-01-02T01:23:45Z",
		                  "updatedAt":"2014-01-12T02:34:56Z",
		                  "lastPostedAt":"2014-01-22T02:34:56Z"},
		                 "favorite":true,
		                 "unread":{"topicId":1111,"postId":123456,"count":1}}]}`)
	case "/api/v1/topics/1111/favorite":
		fmt.Fprint(w, `{"id":1111,
		                "name":"テストトピック",
		                "suggestion":"テストトピック",
		                "description":null,
		                "createdAt":"2014-01-02T01:23:45Z",
		                "updatedAt":"2014-01-12T02:34:56Z",
		                "lastPostedAt":"2014-01-22T02:34:56Z"}`)
	case "/api/v1/notifications/status":
		fmt.Fprint(w, `{"mention":{"unread":1},
		                "access":{"unopened":2},
		                "invite":{"team":{"pending":3},"topic":{"pending":4}}}`)
	case "/api/v1/notifications":
		fmt.Fprint(w, `{"access":{"unopened":5}}`)
	case "/api/v1/bookmarks":
		fmt.Fprint(w, `{"unread":{"topicId":1111,"postId":123456,"count":10}}`)
	case "/api/v1/mentions":
		fmt.Fprint(w, `{"mentions":[
		                {"id":54321,
		                 "readAt":"2014-05-05T00:00:00Z",
		                 "post":
		                 {"id":123456,
		                  "message":"@goyish test!!",
		                  "replyTo":null,
		                  "url":"https://typetalk.in/topics/1111/posts/123456",
		                  "createdAt":"2014-03-03T13:33:33Z",
		                  "updatedAt":"2014-04-04T14:44:44Z",
		                  "topic":
		                  {"id":1111,
		                   "name":"テストトピック",
		                   "suggestion":"テストトピック",
		                   "description":null,
		                   "createdAt":"2014-01-02T01:23:45Z",
		                   "updatedAt":"2014-01-12T02:34:56Z",
		                   "lastPostedAt":"2014-01-22T02:34:56Z"},
		                  "account":
		                  {"id":3333,
		                   "name":"goyish",
		                   "fullName":"goyish account",
		                   "suggestion":"goyish",
		                   "imageUrl":"https://typetalk.in/accounts/3333/profile_image.png?t=9999999999999",
		                   "createdAt":"2014-01-11T11:11:11Z",
		                   "updatedAt":"2014-02-22T12:22:22Z"},
		                  "attachments":[]}}]}`)
	case "/api/v1/mentions/54321":
		fmt.Fprint(w, `{"mention":
		                {"id":54321,
		                 "readAt":"2014-05-05T00:00:00Z",
		                 "post":
		                 {"id":123456,
		                  "message":"@goyish test!!",
		                  "replyTo":null,
		                  "url":"https://typetalk.in/topics/1111/posts/123456",
		                  "createdAt":"2014-03-03T13:33:33Z",
		                  "updatedAt":"2014-04-04T14:44:44Z",
		                  "topic":
		                  {"id":1111,
		                   "name":"テストトピック",
		                   "suggestion":"テストトピック",
		                   "description":null,
		                   "createdAt":"2014-01-02T01:23:45Z",
		                   "updatedAt":"2014-01-12T02:34:56Z",
		                   "lastPostedAt":"2014-01-22T02:34:56Z"},
		                  "account":
		                  {"id":3333,
		                   "name":"goyish",
		                   "fullName":"goyish account",
		                   "suggestion":"goyish",
		                   "imageUrl":"https://typetalk.in/accounts/3333/profile_image.png?t=9999999999999",
		                   "createdAt":"2014-01-11T11:11:11Z",
		                   "updatedAt":"2014-02-22T12:22:22Z"},
		                  "attachments":[]}}}`)
	default:
		panic(path)
	}
}
