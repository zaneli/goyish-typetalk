package typetalk

import . "gopkg.in/check.v1"

func (s *typetalkSuite) TestGetTopicMessages(c *C) {
	client := newTestClient()
	messages, err := client.GetTopicMessages(1111)

	c.Assert(len(messages.Posts), Equals, 2)
	{
		post := messages.Posts[0]
		c.Assert(post.Id, Equals, 121212)
		c.Assert(post.Message, Equals, "テストメッセージ")
		c.Assert(post.ReplyTo, IsNil)
		c.Assert(post.Url, Equals, "https://typetalk.in/topics/1111/posts/121212")
		c.Assert(len(post.Likes), Equals, 0)
		c.Assert(len(post.Links), Equals, 0)
		c.Assert(post.Mention, IsNil)
		c.Assert(post.CreatedAt, Equals, toTime(2014, 6, 16, 16, 16, 16))
		c.Assert(post.UpdatedAt, Equals, toTime(2014, 6, 16, 16, 16, 16))

		account := post.Account
		c.Assert(account.Id, Equals, 3333)
		c.Assert(account.Name, Equals, "goyish")
		c.Assert(account.FullName, Equals, "goyish account")
		c.Assert(account.Suggestion, Equals, "goyish")
		c.Assert(account.ImageUrl, Equals, "https://typetalk.in/accounts/3333/profile_image.png?t=9999999999999")
		c.Assert(account.CreatedAt, Equals, toTime(2014, 1, 11, 11, 11, 11))
		c.Assert(account.UpdatedAt, Equals, toTime(2014, 2, 22, 12, 22, 22))

		c.Assert(len(post.Attachments), Equals, 0)
	}
	{
		post := messages.Posts[1]
		c.Assert(post.Id, Equals, 121213)
		c.Assert(post.Message, Equals, "テストメッセージ2")
		c.Assert(post.ReplyTo, IsNil)
		c.Assert(post.Url, Equals, "https://typetalk.in/topics/1111/posts/121213")

		c.Assert(len(post.Likes), Equals, 1)
		{
			like := post.Likes[0]
			c.Assert(like.Id, Equals, 77777)
			c.Assert(like.PostId, Equals, 121213)
			c.Assert(like.TopicId, Equals, 1111)
			c.Assert(like.Comment, Equals, "")

			account := like.Account
			c.Assert(account.Id, Equals, 3333)
			c.Assert(account.Name, Equals, "goyish")
			c.Assert(account.FullName, Equals, "goyish account")
			c.Assert(account.Suggestion, Equals, "goyish")
			c.Assert(account.ImageUrl, Equals, "https://typetalk.in/accounts/3333/profile_image.png?t=9999999999999")
			c.Assert(account.CreatedAt, Equals, toTime(2014, 1, 11, 11, 11, 11))
			c.Assert(account.UpdatedAt, Equals, toTime(2014, 2, 22, 12, 22, 22))
		}

		c.Assert(len(post.Links), Equals, 0)
		c.Assert(post.Mention, IsNil)
		c.Assert(post.CreatedAt, Equals, toTime(2014, 6, 16, 16, 16, 16))
		c.Assert(post.UpdatedAt, Equals, toTime(2014, 6, 16, 16, 16, 16))

		account := post.Account
		c.Assert(account.Id, Equals, 3333)
		c.Assert(account.Name, Equals, "goyish")
		c.Assert(account.FullName, Equals, "goyish account")
		c.Assert(account.Suggestion, Equals, "goyish")
		c.Assert(account.ImageUrl, Equals, "https://typetalk.in/accounts/3333/profile_image.png?t=9999999999999")
		c.Assert(account.CreatedAt, Equals, toTime(2014, 1, 11, 11, 11, 11))
		c.Assert(account.UpdatedAt, Equals, toTime(2014, 2, 22, 12, 22, 22))

		c.Assert(len(post.Attachments), Equals, 1)
		{
			attachment := post.Attachments[0]
			c.Assert(attachment.FileName, Equals, "goyishtypetalk876965047")
			c.Assert(attachment.ApiUrl, Equals, "https://typetalk.in/api/v1/topics/1111/posts/121213/attachments/1/goyishtypetalk876965047")
			c.Assert(attachment.WebUrl, Equals, "https://typetalk.in/topics/1111/posts/121213/attachments/1/goyishtypetalk876965047")

			file := attachment.Attachment
			c.Assert(file.FileKey, Equals, "acc238a6d9317f816be327b629ebaaf41463fa0a")
			c.Assert(file.FileName, Equals, "goyishtypetalk876965047")
			c.Assert(file.FileSize, Equals, 0)
		}
	}

	bookmark := messages.Bookmark
	c.Assert(bookmark.PostId, Equals, 121212)
	c.Assert(bookmark.UpdatedAt, Equals, toTime(2014, 6, 6, 16, 16, 16))

	c.Assert(messages.HasNext, Equals, false)

	c.Assert(err, IsNil)
}

func (s *typetalkSuite) TestGetMessage(c *C) {
	client := newTestClient()
	post, err := client.GetMessage(1111, 121213)
	c.Assert(post.Id, Equals, 121213)
	c.Assert(post.Message, Equals, "テストメッセージ2")
	c.Assert(post.ReplyTo, IsNil)
	c.Assert(post.Url, Equals, "https://typetalk.in/topics/1111/posts/121213")

	c.Assert(len(post.Likes), Equals, 1)
	{
		like := post.Likes[0]
		c.Assert(like.Id, Equals, 77777)
		c.Assert(like.PostId, Equals, 121213)
		c.Assert(like.TopicId, Equals, 1111)
		c.Assert(like.Comment, Equals, "")

		account := like.Account
		c.Assert(account.Id, Equals, 3333)
		c.Assert(account.Name, Equals, "goyish")
		c.Assert(account.FullName, Equals, "goyish account")
		c.Assert(account.Suggestion, Equals, "goyish")
		c.Assert(account.ImageUrl, Equals, "https://typetalk.in/accounts/3333/profile_image.png?t=9999999999999")
		c.Assert(account.CreatedAt, Equals, toTime(2014, 1, 11, 11, 11, 11))
		c.Assert(account.UpdatedAt, Equals, toTime(2014, 2, 22, 12, 22, 22))
	}

	c.Assert(len(post.Links), Equals, 0)
	c.Assert(post.Mention, IsNil)
	c.Assert(post.CreatedAt, Equals, toTime(2014, 6, 16, 16, 16, 16))
	c.Assert(post.UpdatedAt, Equals, toTime(2014, 6, 16, 16, 16, 16))

	account := post.Account
	c.Assert(account.Id, Equals, 3333)
	c.Assert(account.Name, Equals, "goyish")
	c.Assert(account.FullName, Equals, "goyish account")
	c.Assert(account.Suggestion, Equals, "goyish")
	c.Assert(account.ImageUrl, Equals, "https://typetalk.in/accounts/3333/profile_image.png?t=9999999999999")
	c.Assert(account.CreatedAt, Equals, toTime(2014, 1, 11, 11, 11, 11))
	c.Assert(account.UpdatedAt, Equals, toTime(2014, 2, 22, 12, 22, 22))

	c.Assert(len(post.Attachments), Equals, 1)
	{
		attachment := post.Attachments[0]
		c.Assert(attachment.FileName, Equals, "goyishtypetalk876965047")
		c.Assert(attachment.ApiUrl, Equals, "https://typetalk.in/api/v1/topics/1111/posts/121213/attachments/1/goyishtypetalk876965047")
		c.Assert(attachment.WebUrl, Equals, "https://typetalk.in/topics/1111/posts/121213/attachments/1/goyishtypetalk876965047")

		file := attachment.Attachment
		c.Assert(file.FileKey, Equals, "acc238a6d9317f816be327b629ebaaf41463fa0a")
		c.Assert(file.FileName, Equals, "goyishtypetalk876965047")
		c.Assert(file.FileSize, Equals, 0)
	}
	c.Assert(err, IsNil)
}
