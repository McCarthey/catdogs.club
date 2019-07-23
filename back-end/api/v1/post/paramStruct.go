package post

type SetPost struct {
	Title   string `form:"title"`
	Content string `form:"content"`
	Author  string `form:"author"`
}
