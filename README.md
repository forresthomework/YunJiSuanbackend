# 基于云和MapReduce的文本检索

[![](https://img.shields.io/badge/env-dockercompose2.17.3-blue.svg)](https://github.com/docker/compose)

## 项目介绍



## 开发环境

### 前端部分

#### 环境

[![](https://img.shields.io/badge/Node.js-18.5.0-green.svg)]() [![](https://img.shields.io/badge/npm-9.5.0-green.svg)]()

#### 语言

[![](https://img.shields.io/badge/vue.js-green.svg)]() [![](https://img.shields.io/badge/javascript-yellow.svg)]() [![](https://img.shields.io/badge/html-red.svg)]() [![](https://img.shields.io/badge/css-purple.svg)]()

### 后端部分



## 项目细节



### 倒排索引介绍

倒排索引（Inverted Index）是一种常用的文本索引数据结构，用于加快文本搜索和信息检索的速度。它是一种反转（Inverted）的索引结构，将文档中的每个单词映射到包含该单词的文档列表。

通常，倒排索引由两个主要组成部分构成：词项表（Term Dictionary）和倒排列表（Inverted List）。

词项表（Term Dictionary）：词项表是一个词项到倒排列表的映射，它记录了所有不重复的单词（或词项）以及它们对应的倒排列表的位置信息。

倒排列表（Inverted List）：倒排列表包含了一个单词在文档集合中的出现位置。对于每个单词，倒排列表记录了包含该单词的文档的标识符（例如文档ID）以及该单词在文档中的位置信息（例如单词出现的位置或出现的频率）。


### 前端设计

**文件目录**
```bash
│  .dockerignore
│  Dockerfile
│  index.html
│  package-lock.json
│  package.json
│  results.html
│  server.js
│  vite.config.js
│
├─public
│      favicon-16x16.png
│      icon_search.svg
│      leetcode.png
│      title.ttf
│
└─src
    └─assets
            base.css
            main.css
```

**文件介绍**

index.html和results.html为主要的交互页面

public目录存放一些公共的图片和字体文件

src/assets目录存放css样式

.dockerignore和Dockerfile文件用来生成镜像

server.js用来本地调试启动应用

package.json用来配置项目依赖，package-lock.json是由package.json自动生成

**交互设计**

1. index.html输入查询语句，使用POST发送请求给后端
2. 后端响应并将查询结果返回
3. index.html将后端返回的数据POST发送给result.html进行解析并展示

![image-20230529152714576](https://github.com/forresthomework/YunJiSuanbackend/blob/main/img/image-20230529152714576.png)

index.html

包含一个搜索框，单击“搜索”或者回车进行搜索

![image-20230529152828238](https://github.com/forresthomework/YunJiSuanbackend/blob/main/img/image-20230529152828238.png)

result.html

包含一个搜索框，单击“搜索”或者回车进行搜索，并会将返回的数据显示出来，返回数据的顺序按照count数量从大到小排列，数据包括题目ID、题目名称、关键字出现次数、算法标签、难度标签。通过单击显示的题目可以链接到对应的leetcode题目网页。

![image-20230529152911545](https://github.com/forresthomework/YunJiSuanbackend/blob/main/img/image-20230529152911545.png)

**错误处理**

当查询为空或者返回数据不成功时则会弹窗并刷新页面

![image-20230529152953565](https://github.com/forresthomework/YunJiSuanbackend/blob/main/img/image-20230529152953565.png)

### 后端设计



### 调试和运行
**前端本地调试**

切换到前端网页根目录
```bash
cd front
```
安装依赖
```bash
npm install
```
启动应用
```bash
node server.js
```
访问localhost:5173即可

**docker-compose一键部署**

本程序运行所需要的三个镜像已经上传至Docker Hub，直接拉取使用即可。

编写docker-compose.yml文件

```yaml
version: '3.8'
services:
  backend:
    image: merk11/search-engine:backend
    ports:
      - 9999:9999
    networks:
      - mynetwork
    depends_on:
      - db
  db:
    image: merk11/search-engine:redis
    ports:
      - 6379:6379
    networks:
      - mynetwork	
  front:
    image: merk11/search-engine:front
    ports:
      - 5173:5173
    networks:
      - mynetwork

networks:
  mynetwork:
    driver: bridge
```

在docker-compose.yml文件目录下执行命令

```bash
docker-compose up
```

即可拉取镜像，并启动服务，访问localhost:5173即可。
