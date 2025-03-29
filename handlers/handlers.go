package handlers // ディレクトリ名と同名にする

import (
	"fmt"
	"io"
	"net/http"
)

/*** ハンドラ（helloHandler）***/
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

// ハンドラを追加
func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		io.WriteString(w, "Posting Article...\n")
	} else {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}
}

func ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		io.WriteString(w, "Article List\n")
	} else {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}
}

func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	articleID := 1
	resString := fmt.Sprintf("Article No.%d\n", articleID)
	if req.Method == http.MethodGet {
		io.WriteString(w, resString)
		} else {
			http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		}
	}
	
	func PostNiceHandler(w http.ResponseWriter, req *http.Request) {
		if req.Method == http.MethodPost {
			io.WriteString(w, "Posting Nice...\n")
		} else {
			http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		}
	}
	
	func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
		if req.Method == http.MethodPost {
			io.WriteString(w, "Posting Comment...\n")
		} else {
			http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		}
}