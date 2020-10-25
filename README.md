# Dcard_demo

## 需求
- 每個 IP 每分鐘僅能接受 60 個 requests
- 在首頁顯示目前的 request 量,超過限制的話則顯示 “Error”,例如在一分鐘內第 30 個 request 則顯示 30,第 61 個
- request 則顯示 Error
- 可以使用任意資料庫,也可以自行設計 in-memory 資料結構,並在文件中說明理由
- 請附上測試
- 請不要使用任何現成的 rate limit library

## 專案講解
**使用環境及工具**

1. golang
2. redis
3. gin
4. viper

**解說**
此專案用了 golang  的 gin 框架來開發 Restful API ,裡面有兩種做法來到此需求的目的,實作此兩種方法的地方為 gin 的 middleware

1. in-memory (service/demo) 這個可以使用在需求較沒有這麼大只需要單機就好的小型專案
2. redis (service/demo_redis) 這個可以使用在大型專案的開發當資料量大到 in-memory 承受不住時,我們可以透過 redis 的方式來分散對機器的負擔,當單個 redis 不行時可以透過水平擴張的方式來抒解此專案的壓力

**啟動專案**
修改專案中的 app.conf (預設為 local ) 修改完畢後直接在 terminal 輸入以下指令即可運行

`go run main.go`


**測試**

1. in-memory (service/demo) 前往test 的資料夾輸入以下指令即可測試

`go test -v -run=CheckRequestLimit`

- result

		 RUN   TestCheckRequestLimit
		 PASS: TestCheckRequestLimit (0.00s)
		 PASS
		 ok      github.com/1101137105/dcard/test        0.486s

1. redis (service/demo_redis) 前往test 的資料夾輸入以下指令即可測試

`go test -v -run=CheckRequestRedisLimit`

- result

		 RUN   TestCheckRequestRedisLimit
		 PASS: TestCheckRequestRedisLimit (0.02s)
		 PASS
		 ok      github.com/1101137105/dcard/test        0.491s
