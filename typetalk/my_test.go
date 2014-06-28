package typetalk

import . "gopkg.in/check.v1"

func (s *typetalkSuite) TestGetMyProfile(c *C) {
	client := newTestClient()
	account, err := client.GetMyProfile()
	c.Assert(account.Id, Equals, 3333)
	c.Assert(account.Name, Equals, "goyish")
	c.Assert(account.FullName, Equals, "goyish account")
	c.Assert(account.Suggestion, Equals, "goyish")
	c.Assert(account.ImageUrl, Equals, "https://typetalk.in/accounts/3333/profile_image.png?t=9999999999999")
	c.Assert(account.CreatedAt, Equals, toTime(2014, 1, 11, 11, 11, 11))
	c.Assert(account.UpdatedAt, Equals, toTime(2014, 2, 22, 12, 22, 22))
	c.Assert(err, IsNil)
}

func (s *typetalkSuite) TestGetMyTopics(c *C) {
	client := newTestClient()
	topicInfos, err := client.GetMyTopics()
	c.Assert(len(topicInfos), Equals, 1)

	topic := topicInfos[0].Topic
	c.Assert(topic.Id, Equals, 1111)
	c.Assert(topic.Name, Equals, "テストトピック")
	c.Assert(topic.Suggestion, Equals, "テストトピック")
	c.Assert(topic.Description, IsNil)
	c.Assert(topic.CreatedAt, Equals, toTime(2014, 1, 2, 1, 23, 45))
	c.Assert(topic.UpdatedAt, Equals, toTime(2014, 1, 12, 2, 34, 56))
	c.Assert(*topic.LastPostedAt, Equals, toTime(2014, 01, 22, 02, 34, 56))

	c.Assert(topicInfos[0].Favorite, Equals, true)

	unread := topicInfos[0].Unread
	c.Assert(unread.TopicId, Equals, 1111)
	c.Assert(unread.PostId, Equals, 123456)
	c.Assert(unread.Count, Equals, 1)

	c.Assert(err, IsNil)
}

func (s *typetalkSuite) TestFavoriteTopic(c *C) {
	client := newTestClient()
	topic, err := client.FavoriteTopic(1111)
	c.Assert(topic.Id, Equals, 1111)
	c.Assert(topic.Name, Equals, "テストトピック")
	c.Assert(topic.Suggestion, Equals, "テストトピック")
	c.Assert(topic.Description, IsNil)
	c.Assert(topic.CreatedAt, Equals, toTime(2014, 1, 2, 1, 23, 45))
	c.Assert(topic.UpdatedAt, Equals, toTime(2014, 1, 12, 2, 34, 56))
	c.Assert(*topic.LastPostedAt, Equals, toTime(2014, 01, 22, 02, 34, 56))
	c.Assert(err, IsNil)
}

func (s *typetalkSuite) TestUnfavoriteTopic(c *C) {
	client := newTestClient()
	topic, err := client.UnfavoriteTopic(1111)
	c.Assert(topic.Id, Equals, 1111)
	c.Assert(topic.Name, Equals, "テストトピック")
	c.Assert(topic.Suggestion, Equals, "テストトピック")
	c.Assert(topic.Description, IsNil)
	c.Assert(topic.CreatedAt, Equals, toTime(2014, 1, 2, 1, 23, 45))
	c.Assert(topic.UpdatedAt, Equals, toTime(2014, 1, 12, 2, 34, 56))
	c.Assert(*topic.LastPostedAt, Equals, toTime(2014, 01, 22, 02, 34, 56))
	c.Assert(err, IsNil)
}

func (s *typetalkSuite) TestGetNotificationCount(c *C) {
	client := newTestClient()
	notifications, err := client.GetNotificationCount()

	mention := notifications.Mention
	c.Assert(mention.Unread, Equals, 1)

	access := notifications.Access
	c.Assert(access.Unopened, Equals, 2)

	invite := notifications.Invite
	c.Assert(invite.Team.Pending, Equals, 3)
	c.Assert(invite.Topic.Pending, Equals, 4)

	c.Assert(err, IsNil)
}

func (s *typetalkSuite) TestReadNotifications(c *C) {
	client := newTestClient()
	openStatus, err := client.ReadNotifications()
	c.Assert(openStatus.Unopened, Equals, 5)
	c.Assert(err, IsNil)
}

func (s *typetalkSuite) TestReadMessagesInTopic(c *C) {
	client := newTestClient()
	unread, err := client.ReadMessagesInTopic(1111)
	c.Assert(unread.Count, Equals, 10)
	c.Assert(unread.TopicId, Equals, 1111)
	c.Assert(unread.PostId, Equals, 123456)
	c.Assert(err, IsNil)
}

func (s *typetalkSuite) TestGetMentionList(c *C) {
	client := newTestClient()
	mentions, err := client.GetMentionList()
	c.Assert(len(mentions), Equals, 1)
	c.Assert(mentions[0].Id, Equals, 54321)
	c.Assert(*mentions[0].ReadAt, Equals, toTime(2014, 5, 5, 0, 0, 0))

	post := mentions[0].Post
	c.Assert(post.Id, Equals, 123456)
	c.Assert(post.Message, Equals, "@goyish test!!")
	c.Assert(post.ReplyTo, IsNil)
	c.Assert(post.Url, Equals, "https://typetalk.in/topics/1111/posts/123456")
	c.Assert(post.CreatedAt, Equals, toTime(2014, 3, 3, 13, 33, 33))
	c.Assert(post.UpdatedAt, Equals, toTime(2014, 4, 4, 14, 44, 44))
	c.Assert(len(post.Likes), Equals, 0)
	c.Assert(len(post.Links), Equals, 0)

	topic := post.Topic
	c.Assert(topic.Id, Equals, 1111)
	c.Assert(topic.Name, Equals, "テストトピック")
	c.Assert(topic.Suggestion, Equals, "テストトピック")
	c.Assert(topic.Description, IsNil)
	c.Assert(topic.CreatedAt, Equals, toTime(2014, 1, 2, 1, 23, 45))
	c.Assert(topic.UpdatedAt, Equals, toTime(2014, 1, 12, 2, 34, 56))
	c.Assert(*topic.LastPostedAt, Equals, toTime(2014, 01, 22, 02, 34, 56))

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

func (s *typetalkSuite) TestReadMention(c *C) {
	client := newTestClient()
	mention, err := client.ReadMention(54321)
	c.Assert(mention.Id, Equals, 54321)
	c.Assert(*mention.ReadAt, Equals, toTime(2014, 5, 5, 0, 0, 0))

	post := mention.Post
	c.Assert(post.Id, Equals, 123456)
	c.Assert(post.Message, Equals, "@goyish test!!")
	c.Assert(post.ReplyTo, IsNil)
	c.Assert(post.Url, Equals, "https://typetalk.in/topics/1111/posts/123456")
	c.Assert(post.CreatedAt, Equals, toTime(2014, 3, 3, 13, 33, 33))
	c.Assert(post.UpdatedAt, Equals, toTime(2014, 4, 4, 14, 44, 44))
	c.Assert(len(post.Likes), Equals, 0)
	c.Assert(len(post.Links), Equals, 0)

	topic := post.Topic
	c.Assert(topic.Id, Equals, 1111)
	c.Assert(topic.Name, Equals, "テストトピック")
	c.Assert(topic.Suggestion, Equals, "テストトピック")
	c.Assert(topic.Description, IsNil)
	c.Assert(topic.CreatedAt, Equals, toTime(2014, 1, 2, 1, 23, 45))
	c.Assert(topic.UpdatedAt, Equals, toTime(2014, 1, 12, 2, 34, 56))
	c.Assert(*topic.LastPostedAt, Equals, toTime(2014, 01, 22, 02, 34, 56))

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
