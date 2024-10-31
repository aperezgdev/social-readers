package post_vo

type PostComment string

func NewPostComment(postComment string) PostComment {
	return PostComment(postComment)
}

func (pc PostComment) Validate() bool {
	return len(pc) < 1 && len(pc) > 240
}
