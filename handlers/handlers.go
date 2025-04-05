/* ハンドラ定義を記述 */
package handlers // ディレクトリ名と同名にする

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"myapi/models" //自作パッケージ

	"github.com/gorilla/mux"
)

/*** ハンドラ（helloHandler）***/
/***
// HTTPリクエストを受け取って、それに対するHTTPレスポンスの内容をコネクションに書き込む関数
// 他のパッケージからも参照可能な関数・変数・定数を作成するためには、その名前を大文字から始める
func HelloHandler(w http.ResponseWriter, req *http.Request) {
	// ハンドラの第一引数として渡されていたhttp.ResponseWriter型の変数wに
	// "Hello World!"と書き込む（wに書き込まれた内容が、そのままHTTPレスポンスになる）
	// reqの中のMethodフィールドで条件分岐
	if req.Method == http.MethodGet { // MethodGetには定数「GET」が入っている
	io.WriteString(w, "Hello, world!\n")
	} else {
		// Invalid methodというレスポンスを、405番のステータスコードと共に返す
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}
}
***/
// HTTPリクエストを受け取って、それに対するHTTPレスポンスの内容をコネクションに書き込む関数
// 他のパッケージからも参照可能な関数・変数・定数を作成するためには、その名前を大文字から始める
func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	var reqArticle models.Article

	/*** jsonデコード：json→Go構造体に変換すること ***/
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
	}

	// article := models.Article1 // モックデータ
	article := reqArticle

	/*** jsonエンコード：Go構造体→jsonに変換すること ***/
	json.NewEncoder(w).Encode(article)
}

func ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	queryMap := req.URL.Query()

	var page int
	// 最初の値(p)にキー"page"の下に保存された値が割り当てられる（そのキーが存在しない場合、pは値型のゼロ値(0)）
	// 2番目の値(ok)は、キーがマップに存在する場合はtrue、存在しない場合はfalseのbool型の値が入る。
	if p, ok := queryMap["page"]; ok && len(p) > 0 {
		var err error
		page, err = strconv.Atoi(p[0])

		if err != nil {
			http.Error(w, "Invalid query parameter", http.StatusBadRequest)
			return
		}
	} else {
		page = 1
	}

	log.Println(page) // コンパイルエラー回避のため、一時的に追加

	articleList := []models.Article{models.Article1, models.Article2}
	/*** jsonエンコード ***/
	json.NewEncoder(w).Encode(articleList)
}

func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	// Vars関数は、map形式で取り出す。今回は、idの値を取り出したいので、["id"]としている（もし、Vars(req)だと、map[id:1]が取れる）
	articleID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(w, "Invalid query parameter", http.StatusBadRequest)
		return
	}
	log.Println(articleID) //tmp

	article := models.Article1

	/*** jsonエンコード ***/
	json.NewEncoder(w).Encode(article)
}

func PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	var reqArticle models.Article
	/*** jsonデコード ***/
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
	}

	article := reqArticle
	/*** jsonエンコード ***/
	json.NewEncoder(w).Encode(article)
}

func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	var reqComment models.Comment
	/*** jsonデコード ***/
	if err := json.NewDecoder(req.Body).Decode(&reqComment); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
	}

	comment := models.Comment1
	/*** jsonエンコード ***/
	json.NewEncoder(w).Encode(comment)
}
