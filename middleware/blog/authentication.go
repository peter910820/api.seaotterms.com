package blog

import "os"

var (
	// 因為驗證身份會去不同站台的資料庫抓該站台的身分驗證表，所以將該站台的資料庫寫死在middleware這邊，就不需要再路由額外傳遞
	authDbName = os.Getenv("DATABASE_NAME3")
)

func CheckLogin() {

}
