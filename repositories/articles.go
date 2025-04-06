package repositories

import (
	"database/sql"
	"myapi/models"
)

const (
	articleNumPerPage = 5
)

// 新規投稿をDBにinsertする関数
// DBに保存した記事内容と、発生したエラーを返り値にする
/* func 関数名(引数1,引数2, ...) (返り値1, 返り値2, ...) */
func InsertArticle(db *sql.DB, article models.Article) (models.Article, error) {
	const sqlStr = `
		insert into articles (title, contents, username, nice, created_at) values
			(?, ?, ?, 0, now())
	`

	/* TODO: 構造体`models.Article`を受け取って、それをDBに挿入する処理 */
	var newArticle models.Article // DBに保存するための新しいmodels.Article型の変数
	newArticle.Title, newArticle.Contents, newArticle.UserName = article.Title, article.Contents, article.UserName

	result, err := db.Exec(sqlStr, article.Title, article.Contents, article.UserName)
	if err != nil {
		return models.Article{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return models.Article{}, err
	}

	newArticle.ID = int(id)

	return newArticle, nil
}

// 変数pageで指定されたページに表示する投稿一覧をDBから取得する関数
// 取得したデータと、発生したエラーを返り値にする
func SelectArticleList(db *sql.DB, page int) ([]models.Article, error) {
	const sqlStr = `
		select
			article_id, title, contents, username, nice
		from
			articles
		limit ? offset ?;
	`

	/* TODO: 指定された記事データをDBから取得し、それをmodels.Article構造体のスライス[]models.Articleに詰めて返す処理 */
	// db.Query(クエリ, プレースホルダに入れる引数1、プレースホルダに入れる引数2)
	// Query executes a query that returns rows, typically a SELECT.
	rows, err := db.Query(sqlStr, articleNumPerPage, (page-1)*articleNumPerPage) // limit, offsetに使う数値 //
	if err != nil {
		return nil, err // SelectArticleListの返り値1が[]models.Articleスライスなので、スライス型の0値という意味でnilを返す
	}
	defer rows.Close()

	articleArray := make([]models.Article, 0)
	for rows.Next() { // Next prepares the next result row for reading with the [Rows.Scan] method.
		var article models.Article
		rows.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum)

		articleArray = append(articleArray, article)
	}

	return articleArray, nil
}

// 投稿 ID を指定して、記事データを取得する関数
// -> 取得した記事データと、発生したエラーを返り値にする
func SelectArticleDetail(db *sql.DB, articleID int) (models.Article, error) {
	const sqlStr = `
		select *
		from articles
		where article_id = ?;
	`

	/* TODO: 指定 ID の記事データをデータベースから取得して、それを models.Article 構造体の形で返す処理 */
	row := db.QueryRow(sqlStr, articleID) // QueryRow executes a query that is expected to return at most one row.
	if err := row.Err(); err != nil {
		return models.Article{}, err // SelectArticleDetailの返り値1がmodels.Articleなので、空のmodels.Articleを返す
	}

	var article models.Article
	var createdTime sql.NullTime

	// CommentListが不要である理由：この項目はcommentsテーブルの項目なので`SELECT *`で取得されないため
	err := row.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum, &article.CreatedAt)
	if err != nil {
		return models.Article{}, err
	}

	if createdTime.Valid {
		article.CreatedAt = createdTime.Time
	}

	return article, nil
}

// いいねの数を update する関数
// -> 発生したエラーを返り値にする
func UpdateNiceNum(db *sql.DB, articleID int) error {
	/* TODO: 指定された ID の記事のいいね数を+1 するようにデータベースの中身を更新する処理 */
	// トランザクションの開始
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	const sqlGetNice = `
		select nice
		from articles
		where article_id = ?;
	`

	row := db.QueryRow(sqlGetNice, articleID) // QueryRow executes a query that is expected to return at most one row.
	if err := row.Err(); err != nil {
		tx.Rollback()
		return err
	}

	// sqlGetNiceで抽出したniceをrowに入れる
	var nicenum int
	if err := row.Scan(&nicenum); err != nil {
		tx.Rollback()
		return err
	}

	const sqlUpdateNice = `update articles set nice = ? where article_id = ?`
	_, err = tx.Exec(sqlUpdateNice, nicenum+1, articleID)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil // UpdateNiceNumの返り値はerrorなので、Commit()できた場合はerrが無いという意味でnilを返す
}
