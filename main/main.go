package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	"strconv"
	"strings"
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
	Status_info     string `json:"status_info"`
}

// 接收GET请求
func Search(w http.ResponseWriter, request *http.Request) {

	query := request.URL.Query()

	search := query.Get("search")
	/*
			合理的search应该是这样的
		1.不能为空串 code:9
		2.如果是数字，代表着特定题目，所以应该直接返回,数字应该有一个限制 code:1
		3.如果是字符串分为
			a).two形式，代表一个关键字 code:4
			a).Two Sum形式，代表着指定题目，同2应该直接返回 code:2
			b).easy+sum形式，中间有加号代表多个关键字 code:3
		4.其他情况返回错误信息：未授权的搜索方式 code:9

	*/

	log.Printf("GET: search=%s\n", search)
	checkCode, err := CheckSearchString(search)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(Result{Status_info: err.Error()})
	} else {
		var result Result
		var errMsg error
		var results []Result
		switch checkCode {
		case 1:
			result, errMsg = GetQuestionFromQuestionId(search)
		case 2:
			result, errMsg = GetQuestionFromTitle(search)
		case 3:
			result, errMsg = GetQuestionFromQuestionId(search)
		case 4:
			result, errMsg = GetQuestionFromQuestionId(search)
		case 9:
			result, errMsg = GetQuestionFromQuestionId(search)
		}
		if errMsg != nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(Result{Status_info: err.Error()})
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			if len(results) == 0 {
				json.NewEncoder(w).Encode(result)
			} else {
				json.NewEncoder(w).Encode(results)
			}
		}
	}
}

func CheckSearchString(search string) (int, error) { // int代表
	if isNumberInRange(search) {
		return 1, nil
	} else if strings.Contains(search, " ") {
		return 2, nil
	} else if strings.Contains(search, "+") {
		return 3, nil
	} else if len(search) > 0 && search != "" {
		return 4, nil
	} else {
		return 9, fmt.Errorf("Error:Search formation is wrong!")
	}
}

func GetQuestionsFromKeyWords(search string) ([]Result, error) {

}

func GetQuestionFromTitle(title string) (Result, error) {
	var res Result
	query := "SELECT * FROM questions WHERE title = " + title
	rows, err := db.Query(query)
	if err != nil {
		return Result{}, fmt.Errorf("Error:Can't query data::", err)
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&res.Title, &res.ID, &res.Difficulty, &res.Algorithm_label, &res.URL)
		if err != nil {
			return Result{}, fmt.Errorf("Error:Can't scan data::", err)
		}
	}
	res.Status_info = "Success"
	return res, nil
}

func GetQuestionFromQuestionId(search string) (Result, error) {
	var res Result
	query := "SELECT * FROM questions WHERE questionId = " + search
	rows, err := db.Query(query)
	if err != nil {
		return Result{}, fmt.Errorf("Error:Can't query data::", err)
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&res.Title, &res.ID, &res.Difficulty, &res.Algorithm_label, &res.URL)
		if err != nil {
			return Result{}, fmt.Errorf("Error:Can't scan data::", err)
		}
	}
	res.Status_info = "Success!"
	return res, nil
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

func isNumberInRange(str string) bool {
	// 将字符串转换为整数
	num, err := strconv.Atoi(str)
	if err != nil {
		return false // 转换失败，不是一个有效的数字
	}

	// 判断数字是否在指定范围内
	if num >= 1 && num <= 2200 {
		return true // 在范围内
	} else {
		return false // 不在范围内
	}
}
