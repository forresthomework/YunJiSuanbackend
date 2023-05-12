package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Result struct {
	Content         string `json:"content"`
	Title           string `json:"title"`
	ID              int    `json:"id"`
	Algorithm_label string `json:"algorithm_label"`
	URL             string `json:"url"`
}

// 接收GET请求
func Search(w http.ResponseWriter, request *http.Request) {

	query := request.URL.Query()

	// 第一种方式
	// id := query["id"][0]

	// 第二种方式

	search := query.Get("search")

	log.Printf("GET: search=%s\n", search)
	results := []Result{
		{Content: "Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.\n\nYou may assume that each input would have exactly one solution, and you may not use the same element twice.\n\nYou can return the answer in any order.", Title: "Two-Sum", ID: 1, Algorithm_label: "Array,Hash Table", URL: "https://leetcode.com/problems/two-sum/"},
		{Content: "Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.\n\nYou may assume that each input would have exactly one solution, and you may not use the same element twice.\n\nYou can return the answer in any order.", Title: "Two-Sum", ID: 1, Algorithm_label: "Array,Hash Table", URL: "https://leetcode.com/problems/two-sum/"},
		{Content: "Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.\n\nYou may assume that each input would have exactly one solution, and you may not use the same element twice.\n\nYou can return the answer in any order.", Title: "Two-Sum", ID: 1, Algorithm_label: "Array,Hash Table", URL: "https://leetcode.com/problems/two-sum/"},
		{Content: "Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.\n\nYou may assume that each input would have exactly one solution, and you may not use the same element twice.\n\nYou can return the answer in any order.", Title: "Two-Sum", ID: 1, Algorithm_label: "Array,Hash Table", URL: "https://leetcode.com/problems/two-sum/"},
		{Content: "Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.\n\nYou may assume that each input would have exactly one solution, and you may not use the same element twice.\n\nYou can return the answer in any order.", Title: "Two-Sum", ID: 1, Algorithm_label: "Array,Hash Table", URL: "https://leetcode.com/problems/two-sum/"},
		{Content: "Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.\n\nYou may assume that each input would have exactly one solution, and you may not use the same element twice.\n\nYou can return the answer in any order.", Title: "Two-Sum", ID: 1, Algorithm_label: "Array,Hash Table", URL: "https://leetcode.com/problems/two-sum/"},
		{Content: "Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.\n\nYou may assume that each input would have exactly one solution, and you may not use the same element twice.\n\nYou can return the answer in any order.", Title: "Two-Sum", ID: 1, Algorithm_label: "Array,Hash Table", URL: "https://leetcode.com/problems/two-sum/"},
		{Content: "Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.\n\nYou may assume that each input would have exactly one solution, and you may not use the same element twice.\n\nYou can return the answer in any order.", Title: "Two-Sum", ID: 1, Algorithm_label: "Array,Hash Table", URL: "https://leetcode.com/problems/two-sum/"},
		{Content: "Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.\n\nYou may assume that each input would have exactly one solution, and you may not use the same element twice.\n\nYou can return the answer in any order.", Title: "Two-Sum", ID: 1, Algorithm_label: "Array,Hash Table", URL: "https://leetcode.com/problems/two-sum/"},
		{Content: "Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.\n\nYou may assume that each input would have exactly one solution, and you may not use the same element twice.\n\nYou can return the answer in any order.", Title: "Two-Sum", ID: 1, Algorithm_label: "Array,Hash Table", URL: "https://leetcode.com/problems/two-sum/"},
		{Content: "Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.\n\nYou may assume that each input would have exactly one solution, and you may not use the same element twice.\n\nYou can return the answer in any order.", Title: "Two-Sum", ID: 1, Algorithm_label: "Array,Hash Table", URL: "https://leetcode.com/problems/two-sum/"},
		{Content: "Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.\n\nYou may assume that each input would have exactly one solution, and you may not use the same element twice.\n\nYou can return the answer in any order.", Title: "Two-Sum", ID: 1, Algorithm_label: "Array,Hash Table", URL: "https://leetcode.com/problems/two-sum/"},
		{Content: "Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.\n\nYou may assume that each input would have exactly one solution, and you may not use the same element twice.\n\nYou can return the answer in any order.", Title: "Two-Sum", ID: 1, Algorithm_label: "Array,Hash Table", URL: "https://leetcode.com/problems/two-sum/"},
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(results)
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

	http.HandleFunc("/results.html", Search)

	// 创建 CORS 处理器
	corsHandler := allowCORS(http.DefaultServeMux)

	log.Println("Running at port 9999 ...")
	err := http.ListenAndServe(":9999", corsHandler)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}

}
