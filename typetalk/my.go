package typetalk

import (
	"fmt"
	"strconv"
)

func (c *Client) GetMyProfile() (*Account, error) {
	var profile struct {
		Account Account `json:"account"`
	}
	err := c.get(endPoint{apiName: "profile"}, map[string]string{}, &profile)
	if err != nil {
		return nil, err
	}
	return &profile.Account, nil
}

func (c *Client) GetMyTopics() ([]TopicInfo, error) {
	var topics struct {
		TopicInfo []TopicInfo `json:"topics"`
	}
	err := c.get(endPoint{apiName: "topics"}, map[string]string{}, &topics)
	if err != nil {
		return []TopicInfo{}, err
	}
	return topics.TopicInfo, nil
}

func (c *Client) FavoriteTopic(topicId int) (*Topic, error) {
	var topic Topic
	err := c.post(endPoint{apiName: fmt.Sprintf("topics/%d/favorite", topicId)}, map[string]string{}, nil, true, &topic)
	if err != nil {
		return nil, err
	}
	return &topic, nil
}

func (c *Client) UnfavoriteTopic(topicId int) (*Topic, error) {
	var topic Topic
	err := c.delete(endPoint{apiName: fmt.Sprintf("topics/%d/favorite", topicId)}, map[string]string{}, &topic)
	if err != nil {
		return nil, err
	}
	return &topic, nil
}

func (c *Client) GetNotificationCount() (*Notifications, error) {
	var notifications Notifications
	err := c.get(endPoint{apiName: "notifications/status"}, map[string]string{}, &notifications)
	if err != nil {
		return nil, err
	}
	return &notifications, nil
}

func (c *Client) ReadNotification() (*OpenStatus, error) {
	var access struct {
		Access OpenStatus `json:"access"`
	}
	err := c.put(endPoint{apiName: "notifications/open"}, map[string]string{}, &access)
	if err != nil {
		return nil, err
	}
	return &access.Access, nil
}

type ReadMessagesInTopicApi struct {
	c       *Client
	topicId int
	postId  *int
}

func (c *Client) ReadMessagesInTopicApi(topicId int) *ReadMessagesInTopicApi {
	a := &ReadMessagesInTopicApi{}
	a.c = c
	a.topicId = topicId
	return a
}
func (a *ReadMessagesInTopicApi) PostId(postId int) *ReadMessagesInTopicApi {
	a.postId = &postId
	return a
}
func (a *ReadMessagesInTopicApi) Call() (*Unread, error) {
	var params = map[string]string{}
	params["topicId"] = strconv.Itoa(a.topicId)
	if a.postId != nil {
		params["postId"] = strconv.Itoa(*a.postId)
	}
	var unread struct {
		Unread Unread `json:"unread"`
	}
	err := a.c.post(endPoint{apiName: "bookmark/save"}, params, nil, true, &unread)
	if err != nil {
		return nil, err
	}
	return &unread.Unread, nil
}
func (c *Client) ReadMessagesInTopic(topicId int) (*Unread, error) {
	return c.ReadMessagesInTopicApi(topicId).Call()
}

type GetMentionListApi struct {
	c      *Client
	from   *int
	unread *bool
}

func (c *Client) GetMentionListApi() *GetMentionListApi {
	a := &GetMentionListApi{}
	a.c = c
	return a
}
func (a *GetMentionListApi) From(from int) *GetMentionListApi {
	a.from = &from
	return a
}
func (a *GetMentionListApi) Unread(unread bool) *GetMentionListApi {
	a.unread = &unread
	return a
}
func (a *GetMentionListApi) Call() ([]Mention, error) {
	var params = map[string]string{}
	if a.from != nil {
		params["from"] = strconv.Itoa(*a.from)
	}
	if a.unread != nil {
		params["unread"] = strconv.FormatBool(*a.unread)
	}
	var mentions struct {
		Mentions []Mention `json:"mentions"`
	}
	err := a.c.get(endPoint{apiName: "mentions"}, params, &mentions)
	if err != nil {
		return []Mention{}, err
	}
	return mentions.Mentions, nil
}

func (c *Client) GetMentionList() ([]Mention, error) {
	return c.GetMentionListApi().Call()
}

func (c *Client) ReadMention(mentionId int) (*Mention, error) {
	var mention struct {
		Mention Mention `json:"mention"`
	}
	err := c.put(endPoint{apiName: fmt.Sprintf("mentions/%d", mentionId)}, map[string]string{}, &mention)
	if err != nil {
		return nil, err
	}
	return &mention.Mention, nil
}
