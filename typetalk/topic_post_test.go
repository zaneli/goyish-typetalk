package typetalk

import (
	. "gopkg.in/check.v1"
	"io/ioutil"
	"os"
	"path/filepath"
)

func (s *typetalkSuite) TestPostMessage(c *C) {
	client := newTestClient()
	postResult, err := client.PostMessage("テストメッセージ", 1111)

	topic := postResult.Topic
	c.Assert(topic.Id, Equals, 1111)
	c.Assert(topic.Name, Equals, "テストトピック")
	c.Assert(topic.Suggestion, Equals, "テストトピック")
	c.Assert(topic.Description, IsNil)
	c.Assert(topic.CreatedAt, Equals, toTime(2014, 1, 2, 1, 23, 45))
	c.Assert(topic.UpdatedAt, Equals, toTime(2014, 1, 12, 2, 34, 56))
	c.Assert(*topic.LastPostedAt, Equals, toTime(2014, 01, 22, 02, 34, 56))

	post := postResult.Post
	c.Assert(post.Id, Equals, 121212)
	c.Assert(post.Message, Equals, "テストメッセージ")
	c.Assert(post.ReplyTo, IsNil)
	c.Assert(post.Url, Equals, "https://typetalk.in/topics/1111/posts/121212")
	c.Assert(post.CreatedAt, Equals, toTime(2014, 6, 16, 16, 16, 16))
	c.Assert(post.UpdatedAt, Equals, toTime(2014, 6, 16, 16, 16, 16))
	c.Assert(len(post.Likes), Equals, 0)
	c.Assert(len(post.Links), Equals, 0)
	c.Assert(post.Topic, Equals, postResult.Topic)

	account := post.Account
	c.Assert(account.Id, Equals, 3333)
	c.Assert(account.Name, Equals, "goyish")
	c.Assert(account.FullName, Equals, "goyish account")
	c.Assert(account.Suggestion, Equals, "goyish")
	c.Assert(account.ImageUrl, Equals, "https://typetalk.in/accounts/3333/profile_image.png?t=9999999999999")
	c.Assert(account.CreatedAt, Equals, toTime(2014, 1, 11, 11, 11, 11))
	c.Assert(account.UpdatedAt, Equals, toTime(2014, 2, 22, 12, 22, 22))

	c.Assert(len(post.Attachments), Equals, 0)

	c.Assert(err, IsNil)
}

func (s *typetalkSuite) TestUploadAttachmentFile(c *C) {
	tmp, err := ioutil.TempFile(os.TempDir(), "goyishtypetalk")
	c.Assert(err, IsNil)
	path, err := filepath.Abs(tmp.Name())
	c.Assert(err, IsNil)
	defer os.Remove(path)

	client := newTestClient()
	file, err := client.UploadAttachmentFile(1111, path)
	c.Assert(file.FileKey, Equals, "acc238a6d9317f816be327b629ebaaf41463fa0a")
	c.Assert(file.FileName, Equals, "goyishtypetalk876965047")
	c.Assert(file.FileSize, Equals, 0)
	c.Assert(err, IsNil)
}

func (s *typetalkSuite) TestRemoveMessage(c *C) {
	client := newTestClient()
	result, err := client.RemoveMessage(1111, 999999)
	c.Assert(result, Equals, true)
	c.Assert(err, IsNil)
}

func (s *typetalkSuite) TestLikeMessage(c *C) {
	client := newTestClient()
	like, err := client.LikeMessage(1111, 121212)
	c.Assert(like.Id, Equals, 66666)
	c.Assert(like.TopicId, Equals, 1111)
	c.Assert(like.PostId, Equals, 121212)
	c.Assert(like.Comment, Equals, "")

	account := like.Account
	c.Assert(account.Id, Equals, 3333)
	c.Assert(account.Name, Equals, "goyish")
	c.Assert(account.FullName, Equals, "goyish account")
	c.Assert(account.Suggestion, Equals, "goyish")
	c.Assert(account.ImageUrl, Equals, "https://typetalk.in/accounts/3333/profile_image.png?t=9999999999999")
	c.Assert(account.CreatedAt, Equals, toTime(2014, 1, 11, 11, 11, 11))
	c.Assert(account.UpdatedAt, Equals, toTime(2014, 2, 22, 12, 22, 22))

	c.Assert(err, IsNil)
}

func (s *typetalkSuite) TestUnlikeMessage(c *C) {
	client := newTestClient()
	like, err := client.UnlikeMessage(1111, 121212)
	c.Assert(like.Id, Equals, 66666)
	c.Assert(like.TopicId, Equals, 1111)
	c.Assert(like.PostId, Equals, 121212)
	c.Assert(like.Comment, Equals, "")

	account := like.Account
	c.Assert(account.Id, Equals, 3333)
	c.Assert(account.Name, Equals, "goyish")
	c.Assert(account.FullName, Equals, "goyish account")
	c.Assert(account.Suggestion, Equals, "goyish")
	c.Assert(account.ImageUrl, Equals, "https://typetalk.in/accounts/3333/profile_image.png?t=9999999999999")
	c.Assert(account.CreatedAt, Equals, toTime(2014, 1, 11, 11, 11, 11))
	c.Assert(account.UpdatedAt, Equals, toTime(2014, 2, 22, 12, 22, 22))

	c.Assert(err, IsNil)
}
