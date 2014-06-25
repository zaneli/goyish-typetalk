package typetalk

import (
	"fmt"
	"strconv"
)

type PostMessageApi struct {
	c        *Client
	message  string
	topicId  int
	replyTo  *int
	fileKeys []string
}

func (c *Client) PostMessageApi(message string, topicId int) *PostMessageApi {
	a := &PostMessageApi{}
	a.c = c
	a.message = message
	a.topicId = topicId
	return a
}
func (a *PostMessageApi) ReplyTo(replyTo int) *PostMessageApi {
	a.replyTo = &replyTo
	return a
}
func (a *PostMessageApi) FileKeys(fileKeys ...string) *PostMessageApi {
	a.fileKeys = fileKeys
	return a
}
func (a *PostMessageApi) Call() (*PostResult, error) {
	var params = map[string]string{}
	params["message"] = a.message
	if a.replyTo != nil {
		params["replyTo"] = strconv.Itoa(*a.replyTo)
	}
	for i := 0; i < len(a.fileKeys); i++ {
		params[fmt.Sprintf("fileKeys[%d]", i)] = a.fileKeys[i]
	}
	var result PostResult
	err := a.c.post(endPoint{apiName: fmt.Sprintf("topics/%d", a.topicId)}, params, nil, true, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
func (c *Client) PostMessage(message string, topicId int) (*PostResult, error) {
	return c.PostMessageApi(message, topicId).Call()
}

func (c *Client) UploadAttachmentFile(topicId int, filePath string) (*File, error) {
	var result File
	err := c.post(endPoint{apiName: fmt.Sprintf("topics/%d/attachments", topicId)}, map[string]string{}, &filePath, true, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) RemoveMessage(topicId int, postId int) (bool, error) {
	err := c.delete(endPoint{apiName: fmt.Sprintf("topics/%d/posts/%d", topicId, postId)}, map[string]string{}, nil)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (c *Client) LikeMessage(topicId int, postId int) (*Like, error) {
	var like struct {
		Like Like `json:"like"`
	}
	err := c.post(endPoint{apiName: fmt.Sprintf("topics/%d/posts/%d/like", topicId, postId)}, map[string]string{}, nil, true, &like)
	if err != nil {
		return nil, err
	}
	return &like.Like, nil
}

func (c *Client) UnlikeMessage(topicId int, postId int) (*Like, error) {
	var like struct {
		Like Like `json:"like"`
	}
	err := c.delete(endPoint{apiName: fmt.Sprintf("topics/%d/posts/%d/like", topicId, postId)}, map[string]string{}, &like)
	if err != nil {
		return nil, err
	}
	return &like.Like, nil
}
