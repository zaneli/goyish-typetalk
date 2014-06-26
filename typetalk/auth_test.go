package typetalk

import . "gopkg.in/check.v1"

func (s *typetalkSuite) TestGetAccessToken(c *C) {
	client := newTestClient()
	auth, err := client.GetAccessToken("id", "secret")
	c.Assert(auth.AccessToken, Equals, "test_access_token")
	c.Assert(auth.TokenType, Equals, "Bearer")
	c.Assert(auth.RefreshToken, Equals, "test_refresh_token")
	c.Assert(err, IsNil)
}

func (s *typetalkSuite) TestUpdateAccessToken(c *C) {
	client := newTestClient()
	auth, err := client.UpdateAccessToken("id", "secret", "token")
	c.Assert(auth.AccessToken, Equals, "test_access_token")
	c.Assert(auth.TokenType, Equals, "Bearer")
	c.Assert(auth.RefreshToken, Equals, "test_refresh_token")
	c.Assert(err, IsNil)
}
