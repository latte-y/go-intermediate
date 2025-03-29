package main

import (
	"log"
	"net/http"

	"myapi/handlers" // 自作パッケージ
)

func main() {
	// 定義したハンドラを、サーバーで使用できるように登録
	// 第二引数に先ほど作成した自作ハンドラhellohandlerを渡す
	http.HandleFunc("/hello", handlers.HelloHandler)
	http.HandleFunc("/article", handlers.PostArticleHandler)
	http.HandleFunc("/article/list", handlers.ArticleListHandler)
	http.HandleFunc("/article/1", handlers.ArticleDetailHandler)
	http.HandleFunc("/article/nice", handlers.PostNiceHandler)
	http.HandleFunc("/comment", handlers.PostCommentHandler)

	// サーバー起動時のログを出力
	log.Println("server start at port 8080") // 日時とともに引数に渡された内容がターミナルに出力される

	// localhost:8080にてサーバーを起動
	log.Fatal(http.ListenAndServe(":8080", nil))
}