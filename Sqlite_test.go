package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"os"
	"strings"
	"testing"
)

//func TestPullLeetCode(t *testing.T) {
//	parts := strings.SplitN("1 Two Sum Easy array hash-table https://leetcode.com/problems/two-sum/", " ", 2)
//	fmt.Println(parts[0], "++", parts[1])
//}

func Test_Convet_TXT_2_Redis(t *testing.T) {
	//连接redis服务器
	client_t := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis服务器地址和端口
		Password: "",               // Redis密码，如果有的话
		DB:       0,                // Redis数据库索引 (0默认使用的数据库)
	})
	_, err := client_t.Ping(client_t.Context()).Result()
	if err != nil {
		fmt.Println("连接Redis出错:", err)
		return
	}
	//打开文件
	file, err := os.Open("main/result1.txt")
	if err != nil {
		fmt.Println("无法打开文件:", err)
		return
	}
	defer file.Close()

	// 创建一个Scanner来读取文件内容
	scanner := bufio.NewScanner(file)

	// 逐行读取文件内容,并存储到redis中
	count := 1 //记录行信息，方便找到错误行数
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		kvPairs := strings.SplitN(line, " ", 2)
		mapp := make(map[string]string)
		keyWords := kvPairs[0] //trie
		//redis先取数据，如果没有数据，就插入，如果有，那就全部取到mapp里，
		exists, _ := client_t.Exists(context.Background(), keyWords).Result()
		if exists == 1 {
			//存在，于是全部放到mapp里
			res, err := client_t.HGetAll(context.Background(), keyWords).Result()
			if err != nil || res == nil {
				fmt.Println("发生错误或map为空")
				return
			}
			for k, v := range res {
				mapp[k] = v
			}
		}

		list := kvPairs[1] //(1032,1),(1803,1),(440,1),(745,1),(792,1)
		values := strings.Split(list, "),(")
		for i := 0; i < len(values); i++ {
			values[i] = strings.Trim(values[i], "(),") //values[0] = 1032,1 values[1] = 1803,1
			kv := strings.Split(values[i], ",")
			mapp[kv[0]] = kv[1]
		}
		err = client_t.HSet(context.Background(), keyWords, mapp).Err()
		if err != nil {
			panic(err)
		}
		count++
	}
	// 检查是否有错误发生
	if err := scanner.Err(); err != nil {
		fmt.Println("读取文件时发生错误:", err)
	}
}

//func Test_Convert_DB_2_TXT(t *testing.T) {
//	var res main.Result
//	query := "SELECT * FROM questions"
//	rows, err := main.db.Query(query)
//	if err != nil {
//		fmt.Println("无法查询数据:", err)
//		return
//	}
//	defer rows.Close()
//
//	for rows.Next() {
//		err = rows.Scan(&res.Title, &res.ID, &res.Difficulty, &res.Algorithm_label, &res.URL)
//		WriteToFile("input.txt", res)
//		if err != nil {
//			fmt.Println("数据扫描错误:", err)
//			return
//		}
//	}
//
//	return
//}

//func WriteToFile(fileName string, res main.Result) {
//	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
//	if err != nil {
//		fmt.Println("无法打开文件:", err)
//		return
//	}
//	s := CheckSpace(res.ID + " " + res.Title + " " + res.Difficulty + " " + res.Algorithm_label)
//	file.WriteString(s + "\n")
//	defer file.Close()
//}

//func CheckSpace(s string) string {
//	if strings.HasSuffix(s, " ") {
//		// 去除最后一个字符
//		s = strings.TrimRight(s, " ")
//	}
//	return s
//}

//func Test_Go_With_Redis(t *testing.T) {
//	client := redis.NewClient(&redis.Options{
//		Addr:     "localhost:6379", // Redis服务器地址和端口
//		Password: "",               // 密码，如果没有密码则留空
//		DB:       0,                // Redis数据库索引
//	})
//
//	// 检查连接是否成功
//	_, err := client.Ping(context.Background()).Result()
//	if err != nil {
//		panic(err)
//	}

// 插入数据
//key := "mykey3"
//value := map[string]string{
//	"field1": "1",
//	"field2": "2",
//	"field3": "3",
//}
//
//res := client.HGetAll(context.Background(), "Elements")
//if err != nil {
//	panic(err)
//}
//main.SortByValueDescending(res.Val())
//fmt.Println("****/n")
//print(client.HGet(context.Background(), "mykey", "key1").Result())
//fmt.Println("****/n")
//print(client.HGet(context.Background(), "mykey", "key1").String())
//fmt.Println("****/n")
//print(client.HGet(context.Background(), "mykey", "key1").Name())
//fmt.Println("****/n")
//fmt.Println(res)
//}
