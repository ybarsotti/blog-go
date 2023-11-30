package comment_handler

type PostCommentRequest struct {
	Content string `json:"content"`
	Author  string `json:"author"`
}
