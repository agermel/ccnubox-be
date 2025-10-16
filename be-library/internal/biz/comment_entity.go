package biz

import (
	"time"
)

type Comment struct {
	ID        int       // 评论ID
	SeatID    string    // 关联座位
	Username  string    // 发表评论的用户
	Content   string    // 评论内容
	Rating    int       // 评分（1-5）
	CreatedAt time.Time // 创建时间
}

type CommentRepo interface {
	CreateComment(req *CreateCommentReq) (string, error)
	GetCommentsBySeatID(seatID int) ([]*Comment, error)
	DeleteComment(id int) (string, error)
}

type CreateCommentReq struct {
	SeatID   string
	Content  string
	Rating   int
	Username string
}
