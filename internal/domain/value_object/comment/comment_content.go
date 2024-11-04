package comment_vo

type CommentContent string

func NewCommentContent(commentContent string) CommentContent {
	return CommentContent(commentContent)
}

func (cc CommentContent) Validate() bool {
	return len(cc) > 1 && len(cc) < 240
}
