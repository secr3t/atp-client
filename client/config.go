package client

import "fmt"

const (
	host  = "https://asia.atphosting24.com/taobao/index.php"
	route = "api_tester/call"
	lang  = "zh-CN"
	sort  = "sale"
)

func GetUri(query string) string {
	return fmt.Sprintf("%s?%s", host, query)
}
