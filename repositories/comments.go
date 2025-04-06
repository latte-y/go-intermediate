package repositories

import (
	"database/sql"
	"myapi/models"
)

// 新規投稿をDBにinsertする関数
// -> DBに保存したコメント内容と、発生したエラーを返り値にする
func InsertComment(db *sql.DB, comment models.Comment) (models.Comment, error) {
	const sqlStr = `
		insert into comments (article_id, message, created_at) values
		(?, ?, now());
	`

	// 構造体 models.Comment を受け取って、それをデータベースに挿入する処理
	var newComment models.Comment // DBに保存するための新しいmodels.Comment型の変数
	newComment.ArticleID, newComment.Message = comment.ArticleID, comment.Message

	result, err := db.Exec(sqlStr, comment.ArticleID, comment.Message)
	if err != nil {
		return models.Comment{}, err // InsertCommentの返り値1がmodels.Commentなので、空のmodels.Commentを返す
	}

	id, err := result.LastInsertId()
	if err != nil {
		return models.Comment{}, err
	}

	newComment.CommentID = int(id)

	return newComment, nil
}

// 指定IDの記事についたコメント一覧を取得する関数
// -> 取得したコメントデータと、発生したエラーを返り値にする
func SelectCommentList(db *sql.DB, articleID int) ([]models.Comment, error) {
	const sqlStr = `
		select *
		from comments
		where article_id = ?;
	`
	// 指定 ID の記事についたコメント一覧をデータベースから取得し、それを`models.Comment`構造体のスライス`[]models.Comment`に詰めて返す処理
	rows, err := db.Query(sqlStr, articleID)
	if err != nil {
		return nil, err // SelectCommentListの返り値1が[]models.Commentスライスなので、スライス型の0値という意味でnilを返す
	}
	defer rows.Close()

	commentArray := make([]models.Comment, 0) // 要素数0の[]models.Commentスライス用の新規変数commentArray
	for rows.Next() {                         // Next prepares the next result row for reading with the [Rows.Scan] method.
		var comment models.Comment
		var createdTime sql.NullTime // NullTime represents a [time.Time] that may be null.
		// timeに関しては、NULLが取得されるかもしれないので、一旦createdTimeに入れる
		rows.Scan(&comment.CommentID, &comment.ArticleID, &comment.Message, &createdTime)

		// Valid が false の場合、Time フィールドには「ゼロ値」と呼ばれる初期状態（0001-01-01 00:00:00）が入っている
		if createdTime.Valid {
			comment.CreatedAt = createdTime.Time // NULLでは無い場合、comment.CreatedAtに値を代入
		}

		commentArray = append(commentArray, comment)
	}

	return commentArray, nil
}
