package main

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestPullLeetCode(t *testing.T) {
	parts := strings.SplitN("1 Two Sum Easy array hash-table https://leetcode.com/problems/two-sum/", " ", 2)
	fmt.Println(parts[0], "++", parts[1])
}

func Test_Convert_DB_2_TXT(t *testing.T) {
	var res Result
	query := "SELECT * FROM questions"
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println("无法查询数据:", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&res.Title, &res.ID, &res.Difficulty, &res.Algorithm_label, &res.URL)
		WriteToFile("input.txt", res)
		if err != nil {
			fmt.Println("数据扫描错误:", err)
			return
		}
	}

	return
}

func WriteToFile(fileName string, res Result) {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("无法打开文件:", err)
		return
	}
	s := CheckSpace(res.ID + " " + res.Title + " " + res.Difficulty + " " + res.Algorithm_label)
	file.WriteString(s + "\n")
	defer file.Close()
}

func CheckSpace(s string) string {
	if strings.HasSuffix(s, " ") {
		// 去除最后一个字符
		s = strings.TrimRight(s, " ")
	}
	return s

}
