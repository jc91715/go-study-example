package model

type Userinfo struct {
	Uid        int `orm:"pk"`
	Username   string
	Department string
}
type RainlabBlogPosts struct {
	Id          int
	Title       string
	ContentHtml string
}
