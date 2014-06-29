package typetalk

import "time"

type Auth struct {
	AccessToken  string
	TokenType    string
	RefreshToken string
	ExpireAt     time.Time
}

type Account struct {
	Id         int       `json:"id"`
	Name       string    `json:"name"`
	FullName   string    `json:"fullName"`
	Suggestion string    `json:"suggestion"`
	ImageUrl   string    `json:"imageUrl"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

type AccountInfo struct {
	Account Account `json:"account"`
	Online  bool    `json:"online"`
}

type TopicInfo struct {
	Topic    Topic  `json:"topic"`
	Favorite bool   `json:"favorite"`
	Unread   Unread `json:"unread"`
}

type Topic struct {
	Id           int        `json:"id"`
	Name         string     `json:"name"`
	Suggestion   string     `json:"suggestion"`
	Description  *string    `json:"description"`
	CreatedAt    time.Time  `json:"createdAt"`
	UpdatedAt    time.Time  `json:"updatedAt"`
	LastPostedAt *time.Time `json:"lastPostedAt"`
}

type Unread struct {
	PostId  int `json:"postId"`
	TopicId int `json:"topicId"`
	Count   int `json:"count"`
}

type Messages struct {
	Posts []Post `json:"posts"`
	// Team
	Topic    Topic    `json:"topic"`
	Bookmark Bookmark `json:"bookmark"`
	HasNext  bool     `json:"hasNext"`
}

type PostResult struct {
	Post  Post  `json:"post"`
	Topic Topic `json:"topic"`
}

type Message struct {
	Post     Post          `json:"post"`
	Topic    Topic         `json:"topic"`
	Replies  []Post        `json:"replies"`
	Accounts []AccountInfo `json:"accounts"`
}

type Post struct {
	Id      int          `json:"id"`
	Message string       `json:"message"`
	Url     string       `json:"url"`
	Account Account      `json:"account"`
	Topic   Topic        `json:"topic"`
	Mention *IdAndReadAt `json:"mention"`
	ReplyTo *int         `json:"replyTo"`
	//Talks
	Links       []Link       `json:"links"`
	Likes       []Like       `json:"likes"`
	Attachments []Attachment `json:"attachments"`
	CreatedAt   time.Time    `json:"createdAt"`
	UpdatedAt   time.Time    `json:"updatedAt"`
}

type File struct {
	FileName string `json:"fileName"`
	FileKey  string `json:"fileKey"`
	FileSize int    `json:"fileSize"`
}

type Link struct {
	Id          int       `json:"id"`
	Url         string    `json:"url"`
	ImageUrl    *string   `json:"imageUrl"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	ContentType string    `json:"contentType"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type Like struct {
	Id      int     `json:"id"`
	TopicId int     `json:"topicId"`
	PostId  int     `json:"postId"`
	Account Account `json:"account"`
	Comment string  `json:"comment"`
}

type Notifications struct {
	Access  OpenStatus `json:"access"`
	Invite  Invite     `json:"invite"`
	Mention ReadStatus `json:"mention"`
}

type Mention struct {
	Id     int        `json:"id"`
	Post   Post       `json:"post"`
	ReadAt *time.Time `json:"readAt"`
}

type Bookmark struct {
	PostId    int       `json:"postId"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Invite struct {
	Team  Process `json:"team"`
	Topic Process `json:"topic"`
}

type Process struct {
	Pending int `json:"pending"`
}

type OpenStatus struct {
	Unopened int `json:"unopened"`
}

type ReadStatus struct {
	Unread int `json:"unread"`
}

type Attachment struct {
	FileName   string `json:"fileName"`
	Attachment File   `json:"attachment"`
	WebUrl     string `json:"webUrl"`
	ApiUrl     string `json:"apiUrl"`
}

type IdAndReadAt struct {
	Id     int        `json:"id"`
	ReadAt *time.Time `json:"readAt"`
}
