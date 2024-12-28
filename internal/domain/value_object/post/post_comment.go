package post_vo

import "github.com/aperezgdev/social-readers-api/internal/domain/errors"

type PostComment string

func NewPostComment(postComment string) (PostComment, error) {
	return PostComment(postComment), validateComment(postComment)
}

func validateComment(pc string) error {
	if len(pc) < 1 || len(pc) > 240 {
		return errors.OutRange("Comment", 1, 240)
	}
	return nil
}
