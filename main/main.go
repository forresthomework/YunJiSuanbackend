package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
)

var db *sql.DB

func init() {
	database, err := sql.Open("sqlite3", "questions.db")
	if err != nil {
		fmt.Println("无法连接数据库:", err)
		return
	}
	db = database
}

type Result struct {
	Title           string `json:"title"`
	ID              string `json:"id"`
	Algorithm_label string `json:"algorithm_label"`
	URL             string `json:"url"`
	Difficulty      string `json:"difficulty"`
}

// 接收GET请求
func Search(w http.ResponseWriter, request *http.Request) {

	query := request.URL.Query()

	// 第一种方式
	// id := query["id"][0]

	// 第二种方式

	search := query.Get("search")

	log.Printf("GET: search=%s\n", search)
	result := GetSpecificQuestionFromDataBase(search)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func GetSpecificQuestionFromDataBase(search string) Result {
	var res Result
	query := "SELECT * FROM questions WHERE questionId = " + search
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println("无法查询数据:", err)
		return Result{}
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&res.Title, &res.ID, &res.Difficulty, &res.Algorithm_label, &res.URL)
		if err != nil {
			fmt.Println("数据扫描错误:", err)
			return Result{}
		}
	}

	return res
}

// 设置 CORS 头信息的处理器函数
func allowCORS(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 设置允许跨域请求的头信息
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// 如果是预检请求（OPTIONS），直接返回成功状态码
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// 调用原始的处理器函数处理请求
		handler.ServeHTTP(w, r)
	})
}

func main() {
	// 创建 CORS 处理器
	corsHandler := allowCORS(http.DefaultServeMux)
	http.HandleFunc("/results.html", Search)
	log.Println("Running at port 9999 ...")
	err := http.ListenAndServe(":9999", corsHandler)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}

}
