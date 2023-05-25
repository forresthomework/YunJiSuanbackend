# YunJiSuanbackend
云计算安全期末后端作业

```bash
git clone https://github.com/forresthomework/YunJiSuanbackend.git
cd YunJiSuanbackend
docker run -d --name my-redis -p 6379:6379 redis
go test -run Test_Convet_TXT_2_Redis
//1.从main函数启动
cd main
go run main.go
//2.从docker启动
docker build -t backend .
docker run -d --name yunjisuanbackend -p 9999:9999 backend
```