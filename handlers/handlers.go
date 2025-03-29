package handlers // ディレクトリ名と同名にする

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

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
func HelloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, world!\n")
}

// ハンドラを追加
func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Article...\n")
}

func ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	queryMap := req.URL.Query()

	var page int
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

	resString := fmt.Sprintf("Article List (page %d)\n", page)
	io.WriteString(w, resString)
}

func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	// Vars関数は、map形式で取り出す。今回は、idの値を取り出したいので、["id"]としている
	// （もし、Vars(req)だと、map[id:1]が取れる）
	// Atoi：string to int
	articleID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(w, "Invalid query parameter", http.StatusBadRequest)
		return
	}
	resString := fmt.Sprintf("Article No.%d\n", articleID)
	io.WriteString(w, resString)
}

func PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Nice...\n")
}

func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Comment...\n")
}
