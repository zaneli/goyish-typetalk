package typetalk

import (
  "fmt"
  "strconv"
)

type GetTopicMessagesApi struct {
  c *Client
  topicId int
  count *int
  from *int
  forward bool
}
func (c *Client) GetTopicMessagesApi(topicId int) (*GetTopicMessagesApi) {
  a := &GetTopicMessagesApi{}
  a.c = c
  a.topicId = topicId
  return a
}
func (a *GetTopicMessagesApi) Count(count int) (*GetTopicMessagesApi) {
  a.count = &count
  return a
}
func (a *GetTopicMessagesApi) From(from int) (*GetTopicMessagesApi) {
  a.from = &from
  return a
}
func (a *GetTopicMessagesApi) Forward() (*GetTopicMessagesApi) {
  a.forward = true
  return a
}
func (a *GetTopicMessagesApi) Call() (*Messages, error) {
  var params = map[string] string{}
  if a.count != nil {
    params["count"] = strconv.Itoa(*a.count)
  }
  if a.from != nil {
    params["from"] = strconv.Itoa(*a.from)
  }
  if a.forward {
    params["direction"] = "forward"
  }
  var messages Messages
  err := a.c.get(fmt.Sprintf("topics/%d", a.topicId), params, &messages)
  if err != nil {
    return nil, err
  }
  return &messages, nil
}
func (c *Client) GetTopicMessages(topicId int) (*Messages, error) {
  return c.GetTopicMessagesApi(topicId).Call()
}

func (c *Client) GetMessage(topicId int, postId int) (*Post, error) {
  var post struct {
    Post Post `json:"post"`
  }
  err := c.get(fmt.Sprintf("topics/%d/posts/%d", topicId, postId), map[string] string {}, &post)
  if err != nil {
    return nil, err
  }
  return &post.Post, nil
}
