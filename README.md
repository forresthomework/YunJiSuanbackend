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

包含一个搜索框，单击“搜索”或者回车进行搜索，并会将返回的数据显示出来，返回数据的顺序按照count数量从大到小排列，数据包括题目ID、题目名称、关键字出现次数、算法标签、难度标签。

![image-20230529152911545](https://github.com/forresthomework/YunJiSuanbackend/blob/main/img/image-20230529152911545.png)

**错误处理**

当查询为空或者返回数据不成功时则会弹窗并刷新页面

![image-20230529152953565](https://github.com/forresthomework/YunJiSuanbackend/blob/main/img/image-20230529152953565.png)

### 后端设计



### 调试和运行

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
