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
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		dummyGetResponse(w, r.URL.Path)
	case "POST":
		dummyPostResponse(w, r.URL.Path)
	case "PUT":
		dummyPutResponse(w, r.URL.Path)
	case "DELETE":
		dummyDeleteResponse(w, r.URL.Path)
	default:
		panic(r.Method)
	}
}

func dummyGetResponse(w http.ResponseWriter, path string) {
	switch path {
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
	case "/api/v1/notifications":
		fmt.Fprint(w, `{"access":{"unopened":5}}`)
	case "/api/v1/notifications/status":
		fmt.Fprint(w, `{"mention":{"unread":1},
		                "access":{"unopened":2},
		                "invite":{"team":{"pending":3},"topic":{"pending":4}}}`)
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
	case "/api/v1/topics/1111":
		fmt.Fprint(w, `{"posts":[
		                {"id":121212,
		                 "message":"テストメッセージ",
		                 "replyTo":null,
		                 "url":"https://typetalk.in/topics/1111/posts/121212",
		                 "topicId":1111,
		                 "likes":[],
		                 "links":[],
		                 "talks":[],
		                 "mention":null,
		                 "createdAt":"2014-06-16T16:16:16Z",
		                 "updatedAt":"2014-06-16T16:16:16Z",
		                 "account":
		                 {"id":3333,
		                  "name":"goyish",
		                  "fullName":"goyish account",
		                  "suggestion":"goyish",
		                  "imageUrl":"https://typetalk.in/accounts/3333/profile_image.png?t=9999999999999",
		                  "createdAt":"2014-01-11T11:11:11Z",
		                  "updatedAt":"2014-02-22T12:22:22Z"},
		                 "attachments":[]},
		                {"id":121213,
		                 "message":"テストメッセージ2",
		                 "replyTo":null,
		                 "url":"https://typetalk.in/topics/1111/posts/121213",
		                 "topicId":1111,
		                 "likes":[
		                 {"id":77777,
		                  "topicId":1111,
		                  "postId":121213,
		                  "comment":"",
		                  "account":
		                  {"id":3333,
		                   "name":"goyish",
		                   "fullName":"goyish account",
		                   "suggestion":"goyish",
		                   "imageUrl":"https://typetalk.in/accounts/3333/profile_image.png?t=9999999999999",
		                   "createdAt":"2014-01-11T11:11:11Z",
		                   "updatedAt":"2014-02-22T12:22:22Z"}}],
		                 "links":[],
		                 "talks":[],
		                 "mention":null,
		                 "createdAt":"2014-06-16T16:16:16Z",
		                 "updatedAt":"2014-06-16T16:16:16Z",
		                 "account":
		                 {"id":3333,
		                  "name":"goyish",
		                  "fullName":"goyish account",
		                  "suggestion":"goyish",
		                  "imageUrl":"https://typetalk.in/accounts/3333/profile_image.png?t=9999999999999",
		                  "createdAt":"2014-01-11T11:11:11Z",
		                  "updatedAt":"2014-02-22T12:22:22Z"},
		                 "attachments":[
		                 {"fileName":"goyishtypetalk876965047",
		                  "attachment":{"fileKey":"acc238a6d9317f816be327b629ebaaf41463fa0a","fileName":"goyishtypetalk876965047","fileSize":0},
		                  "webUrl":"https://typetalk.in/topics/1111/posts/121213/attachments/1/goyishtypetalk876965047",
		                  "apiUrl":"https://typetalk.in/api/v1/topics/1111/posts/121213/attachments/1/goyishtypetalk876965047"}]}],
		                "team":null,
		                "topic":
		                 {"id":1111,
		                  "name":"テストトピック",
		                  "suggestion":"テストトピック",
		                  "description":null,
		                  "createdAt":"2014-01-02T01:23:45Z",
		                  "updatedAt":"2014-01-12T02:34:56Z",
		                  "lastPostedAt":"2014-01-22T02:34:56Z"},
		                "bookmark":{"postId":121212,"updatedAt":"2014-06-06T16:16:16Z"},
		                "hasNext":false}`)
	case "/api/v1/topics/1111/posts/121213":
		fmt.Fprint(w, `{"topic":
		                 {"id":1111,
		                  "name":"テストトピック",
		                  "suggestion":"テストトピック",
		                  "description":null,
		                  "createdAt":"2014-01-02T01:23:45Z",
		                  "updatedAt":"2014-01-12T02:34:56Z",
		                  "lastPostedAt":"2014-01-22T02:34:56Z"},
		                "post":
		                {"id":121213,
		                 "message":"テストメッセージ2",
		                 "replyTo":null,
		                 "url":"https://typetalk.in/topics/1111/posts/121213",
		                 "topicId":1111,
		                 "likes":[
		                 {"id":77777,
		                  "topicId":1111,
		                  "postId":121213,
		                  "comment":"",
		                  "account":
		                  {"id":3333,
		                   "name":"goyish",
		                   "fullName":"goyish account",
		                   "suggestion":"goyish",
		                   "imageUrl":"https://typetalk.in/accounts/3333/profile_image.png?t=9999999999999",
		                   "createdAt":"2014-01-11T11:11:11Z",
		                   "updatedAt":"2014-02-22T12:22:22Z"}}],
		                 "links":[],
		                 "talks":[],
		                 "mention":null,
		                 "createdAt":"2014-06-16T16:16:16Z",
		                 "updatedAt":"2014-06-16T16:16:16Z",
		                 "account":
		                 {"id":3333,
		                  "name":"goyish",
		                  "fullName":"goyish account",
		                  "suggestion":"goyish",
		                  "imageUrl":"https://typetalk.in/accounts/3333/profile_image.png?t=9999999999999",
		                  "createdAt":"2014-01-11T11:11:11Z",
		                  "updatedAt":"2014-02-22T12:22:22Z"},
		                 "attachments":[
		                 {"fileName":"goyishtypetalk876965047",
		                  "attachment":{"fileKey":"acc238a6d9317f816be327b629ebaaf41463fa0a","fileName":"goyishtypetalk876965047","fileSize":0},
		                  "webUrl":"https://typetalk.in/topics/1111/posts/121213/attachments/1/goyishtypetalk876965047",
		                  "apiUrl":"https://typetalk.in/api/v1/topics/1111/posts/121213/attachments/1/goyishtypetalk876965047"}]},
		                "team":null,
		                "replies":[]}`)
	default:
		panic(path)
	}
}

func dummyPostResponse(w http.ResponseWriter, path string) {
	switch path {
	case "/oauth2/access_token":
		fmt.Fprint(w, `{"access_token": "test_access_token",
		                "token_type": "Bearer",
		                "expires_in": 3600,
		                "refresh_token": "test_refresh_token"}`)
	case "/api/v1/topics/1111/favorite":
		fmt.Fprint(w, `{"id":1111,
		                "name":"テストトピック",
		                "suggestion":"テストトピック",
		                "description":null,
		                "createdAt":"2014-01-02T01:23:45Z",
		                "updatedAt":"2014-01-12T02:34:56Z",
		                "lastPostedAt":"2014-01-22T02:34:56Z"}`)
	case "/api/v1/topics/1111":
		fmt.Fprint(w, `{"topic":
		                 {"id":1111,
		                  "name":"テストトピック",
		                  "suggestion":"テストトピック",
		                  "description":null,
		                  "createdAt":"2014-01-02T01:23:45Z",
		                  "updatedAt":"2014-01-12T02:34:56Z",
		                  "lastPostedAt":"2014-01-22T02:34:56Z"},
		                 "post":
		                 {"id":121212,
		                  "message":"テストメッセージ",
		                  "replyTo":null,
		                  "url":"https://typetalk.in/topics/1111/posts/121212",
		                  "topicId":1111,
		                  "likes":[],
		                  "links":[],
		                  "talks":[],
		                  "mention":null,
		                  "createdAt":"2014-06-16T16:16:16Z",
		                  "updatedAt":"2014-06-16T16:16:16Z",
		                  "account":
		                  {"id":3333,
		                   "name":"goyish",
		                   "fullName":"goyish account",
		                   "suggestion":"goyish",
		                   "imageUrl":"https://typetalk.in/accounts/3333/profile_image.png?t=9999999999999",
		                   "createdAt":"2014-01-11T11:11:11Z",
		                   "updatedAt":"2014-02-22T12:22:22Z"},
		                  "attachments":[]}}`)
	case "/api/v1/topics/1111/attachments":
		fmt.Fprint(w, `{"fileKey":"acc238a6d9317f816be327b629ebaaf41463fa0a","fileName":"goyishtypetalk876965047","fileSize":0}`)
	case "/api/v1/topics/1111/posts/121212/like":
		fmt.Fprint(w, `{"like":
		                {"id":66666,
		                 "topicId":1111,
		                 "postId":121212,
		                 "comment":"",
		                 "account":
		                 {"id":3333,
		                  "name":"goyish",
		                  "fullName":"goyish account",
		                  "suggestion":"goyish",
		                  "imageUrl":"https://typetalk.in/accounts/3333/profile_image.png?t=9999999999999",
		                  "createdAt":"2014-01-11T11:11:11Z",
		                  "updatedAt":"2014-02-22T12:22:22Z"}}}`)
	default:
		panic(path)
	}
}

func dummyPutResponse(w http.ResponseWriter, path string) {
	switch path {
	case "/api/v1/notifications":
		fmt.Fprint(w, `{"access":{"unopened":5}}`)
	case "/api/v1/bookmarks":
		fmt.Fprint(w, `{"unread":{"topicId":1111,"postId":123456,"count":10}}`)
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

func dummyDeleteResponse(w http.ResponseWriter, path string) {
	switch path {
	case "/api/v1/topics/1111/favorite":
		fmt.Fprint(w, `{"id":1111,
		                "name":"テストトピック",
		                "suggestion":"テストトピック",
		                "description":null,
		                "createdAt":"2014-01-02T01:23:45Z",
		                "updatedAt":"2014-01-12T02:34:56Z",
		                "lastPostedAt":"2014-01-22T02:34:56Z"}`)
	case "/api/v1/topics/1111/posts/999999":
	case "/api/v1/topics/1111/posts/121212/like":
		fmt.Fprint(w, `{"like":
		                {"id":66666,
		                 "topicId":1111,
		                 "postId":121212,
		                 "comment":"",
		                 "account":
		                 {"id":3333,
		                  "name":"goyish",
		                  "fullName":"goyish account",
		                  "suggestion":"goyish",
		                  "imageUrl":"https://typetalk.in/accounts/3333/profile_image.png?t=9999999999999",
		                  "createdAt":"2014-01-11T11:11:11Z",
		                  "updatedAt":"2014-02-22T12:22:22Z"}}}`)
	default:
		panic(path)
	}
}
