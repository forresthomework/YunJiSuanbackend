package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

var db *sql.DB
var client *redis.Client

func init() {
	database, err := sql.Open("sqlite3", "questions.db")
	if err != nil {
		fmt.Println("无法连接数据库:", err)
		return
	}
	clt := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis服务器地址和端口
		Password: "",               // Redis密码，如果有的话
		DB:       0,                // Redis数据库索引 (0默认使用的数据库)
	})
	_, err = clt.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println("连接Redis出错:", err)
		return
	}
	db = database
	client = clt
}

type Result struct {
	Title           string `json:"title"`
	ID              string `json:"id"`
	Algorithm_label string `json:"algorithm_label"`
	URL             string `json:"url"`
	Difficulty      string `json:"difficulty"`
	Status_info     string `json:"status_info"`
	Times           string `json:"times"` //单词在文章中出现的次数
}

func main() {
	// 创建 CORS 处理器
	//corsHandler := allowCORS(http.DefaultServeMux)
	http.HandleFunc("/results.html", Search)
	log.Println("Running at port 9999 ...")
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}

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
		//case 3:
		//	result, errMsg = GetQuestionFromQuestionId(search)
		case 4:
			results, errMsg = GetQuestionsFromKeyWords(search)
		case 9:
			errMsg = errors.New("do not give a empty string or the search words are unauthorized")
		}
		if errMsg != nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(Result{Status_info: errMsg.Error()})
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
		return 9, errors.New("Error:Search formation is wrong!")
	}
}

func GetQuestionsFromKeyWords(search string) ([]Result, error) {
	exists, err := client.Exists(context.Background(), search).Result()
	if err != nil {
		return nil, err
	}
	if exists == 0 {
		return nil, errors.New("Error:This word in not in database")
	}
	res, err := client.HGetAll(context.Background(), search).Result()
	if err != nil || res == nil {
		return nil, errors.New("Error:can't retrieve data in DB")
	}
	pairs := SortByValueDescending(res)
	results := make([]Result, len(pairs))
	for i, pair := range pairs {
		results[i], err = GetQuestionFromQuestionId(pair.Key)
		if err != nil {
			return nil, err
		}
		results[i].Times = pair.Value
	}
	return results, nil
}

func GetQuestionFromTitle(title string) (Result, error) {
	var res Result
	query := "SELECT * FROM questions WHERE title = " + title
	rows, err := db.Query(query)
	if err != nil {
		return Result{}, fmt.Errorf("Error:Can't query data:%s", err)
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&res.Title, &res.ID, &res.Difficulty, &res.Algorithm_label, &res.URL)
		if err != nil {
			return Result{}, fmt.Errorf("Error:Can't scan data:%s", err)
		}
	}
	res.Status_info = "Success!"
	res.Times = "1"
	return res, nil
}

func GetQuestionFromQuestionId(search string) (Result, error) {
	fmt.Println(search)
	var res Result
	query := "SELECT * FROM QUESTIONS WHERE questionId = " + search
	rows, err := db.Query(query)
	if err != nil {
		return Result{}, fmt.Errorf("Error:Can't query data:%s", err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&res.Title, &res.ID, &res.Difficulty, &res.Algorithm_label, &res.URL)
		if err != nil {
			return Result{}, fmt.Errorf("Error:Can't scan data:%s", err)
		}
	}
	if res.ID == "" {
		return Result{}, errors.New("We can't get this question info because we need ")
	}
	res.Status_info = "Success!"
	res.Times = "1"
	return res, nil
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

func SortByValueDescending(myMap map[string]string) []Pair {
	// 创建一个切片来存储map的键值对
	// 创建一个切片来存储map的键值对
	pairs := make([]Pair, 0, len(myMap))

	// 将键值对存储到切片中
	for key, value := range myMap {
		pairs = append(pairs, Pair{key, value})
	}

	// 使用sort.Slice函数对切片进行排序
	sort.Slice(pairs, func(i, j int) bool {
		value1 := parseInt(pairs[i].Value)
		value2 := parseInt(pairs[j].Value)
		return value1 > value2
	})

	// 按照排序后的切片顺序遍历map

	return pairs
}

func parseInt(s string) int {
	// 这里可以根据实际情况添加错误处理逻辑
	// 这里简单地使用fmt包提供的函数将字符串转换为整数
	num := 0
	_, _ = fmt.Sscanf(s, "%d", &num)
	return num
}

type Pair struct {
	Key   string
	Value string
}
