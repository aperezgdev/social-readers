package comment_vo

import "github.com/aperezgdev/social-readers-api/internal/domain/errors"

type CommentContent string

func NewCommentContent(commentContent string) (CommentContent, error) {
	return CommentContent(commentContent), validateCommentContent(commentContent)
}

func validateCommentContent(cc string) error {
	if len(cc) < 1 || len(cc) > 240 {
		return errors.OutRange("CommentContent", 1, 240)
	}
	return nil
}
