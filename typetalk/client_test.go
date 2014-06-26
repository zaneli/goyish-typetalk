package typetalk

import (
	"fmt"
	. "gopkg.in/check.v1"
	"net/http"
	"testing"
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
	default:
		panic(path)
	}
}
