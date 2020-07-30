package main

import (
	"bytes"
	"fmt"
	"net/http"
)

func main() {

	// "/" にアクセスされたリクエストを処理する
	http.HandleFunc("/", RootHandler)
	http.HandleFunc("/hoge", HogeHandler)
	http.HandleFunc("/hello", HelloHandler)
	http.HandleFunc("/todo", TodoHandler)
	// 8080ポートで起動
	http.ListenAndServe(":8080", nil)
	// nil = ない

}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Print("アクセスされたよ\n")
	fmt.Fprint(w, "アクセスされたよ")
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Print("アクセスされたよ\n")
	fmt.Fprint(w, "hello")
}

func HogeHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Print("アクセスされたよ\n")
	fmt.Fprint(w, "hoge")
}

var hashMap map[int]string

//イニット関数内でやらなくてはいけない処理だった！→ググってやる！！
func init() {
	hashMap = make(map[int]string)
}

func TodoHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Print("アクセスされたよ\n")

	switch r.Method {
	case "GET":
		for _, b := range hashMap {
			fmt.Fprintln(w, b)
		}
	case "POST":
		bufbody := new(bytes.Buffer)
		bufbody.ReadFrom(r.Body)
		body := bufbody.String()
		hashMap[len(hashMap)] = body
	default:
		fmt.Fprint(w, "todoリストだよ")
	}

}
