/* APIで使うデータ構造たいを定義 */
package models

import "time"

type Comment struct {
	CommentID int       `json:"comment_id"` //タグ（jsonキーの名前を指定）
	ArticleID int       `json:"article_id"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at""` // 標準パッケージtimeに含まれるtime.Time型
}

type Article struct {
	ID          int       `json:""`
	Title       string    `json:"title"`
	Contents    string    `json:"contents"`
	UserName    string    `json:"user_name"`
	NiceNum     int       `json:"nice"`
	CommentList []Comment `json:"comments"`
	CreatedAt   time.Time `json:"created_at"`
}
