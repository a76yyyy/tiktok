package pack

import (
	"github.com/a76yyyy/tiktok/kitex_gen/comment"

	"github.com/a76yyyy/tiktok/dal/db"
)

func Comments(vs []*db.Comment) ([]*comment.Comment, error) {
	comments := make([]*comment.Comment, 0)
	for _, v := range vs {
		comments = append(comments, &comment.Comment{
			Id:         int64(v.ID),
			User:       User(&v.User),
			Content:    v.Content,
			CreateDate: v.CreatedAt.Format("01-02"),
		})
	}
	return comments, nil
}
