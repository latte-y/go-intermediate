/* -------------------------------------------------------------- */
/*	Unit Test 実行方法：repositoriesディレクトリにて、`go test`を実行   */
/* -------------------------------------------------------------- */

// テストコードは本来その対象である SelectArticleDetail 関数だけが使えれば動く、という形が理想的。
// そのため、今回はテスト対象とテストコードを疎結合にするために、articles_test.go ファイルで使うパッケージ名としてrepositories_testを用いる

package repositories_test

import (
	"database/sql"
	"fmt"
	"myapi/models"
	"myapi/repositories"
	"testing"

	// ドライバ：MySQLサーバーと直接通信し、クエリ実行するコンポーネント
	_ "github.com/go-sql-driver/mysql" // _をつけてコンパイルエラ-を回避
)

func TestSelectArticleDetail(t *testing.T) {
	dbUser := "docker"
	dbPassword := "docker"
	dbDatabase := "sampledb"
	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)

	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		t.Fatal(err) // 関数の実行に失敗した場合には、テストも失敗させる（それ以降の処理も行われない）
	}
	defer db.Close()

	// テスト結果として期待する値
	expected := models.Article{
		ID:       1,
		Title:    "firstPost",
		Contents: "This is my first blog",
		UserName: "yuki",
		NiceNum:  3,
	}

	// テスト対象となる関数を実行 -> 結果が got に格納される
	got, err := repositories.SelectArticleDetail(db, expected.ID)
	if err != nil {
		t.Fatal(err) // 関数の実行に失敗した場合には、テストも失敗させる（それ以降の処理も行われない）
	}

	// テスト結果を比較
	if got.ID != expected.ID {
		t.Errorf("ID: get %d but want %d\n", got.ID, expected.ID) // 不一致の場合、テスト失敗（それ以降の処理は継続される）
	}
	if got.Title != expected.Title {
		t.Errorf("Title: get %s but want %s\n", got.Title, expected.Title) // 不一致の場合、テスト失敗（それ以降の処理は継続される）
	}
	if got.Contents != expected.Contents {
		t.Errorf("Contents: get %s but want %s\n", got.Contents, expected.Contents) // 不一致の場合、テスト失敗（それ以降の処理は継続される）
	}
	if got.UserName != expected.UserName {
		t.Errorf("UserName: get %s but want %s\n", got.UserName, expected.UserName) // 不一致の場合、テスト失敗（それ以降の処理は継続される）
	}
	if got.NiceNum != expected.NiceNum {
		t.Errorf("NiceNum: get %d but want %d\n", got.NiceNum, expected.NiceNum) // 不一致の場合、テスト失敗（それ以降の処理は継続される）
	}
}
