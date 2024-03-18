a Go-zero quick start

# 文档v1  
 
## 后端api生成方法：  
1. 在go目录下创建xxx文件夹，并编写xxx.api （api写法：https://go-zero.dev/docs/tasks/dsl/api）
2. goctl生成 ```goctl api go -api xxx.api -dir .```(每次修改api文件后都要执行一次以修正自动生成的代码)
3. 在logic中书写后端逻辑。(new一个respose)。
4. 启动服务例：```go run web/go/weather-Beta/weather.go  ```
5. 测试例
```
curl -X GET -H "Content-Type: application/json" -d '{
  "longitude": "xxx",
  "latitude": "xxx",
  "city": "KM"
}' http://localhost:8888/v1/weather/daily
```
6. 编译：xxx文件夹下，```env GOOS=linux GOARCH=amd64 go build ```

7. 执行生成的运行文件
