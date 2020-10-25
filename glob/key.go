package glob

import (
	"strconv"
)

func CreateIPKey(ip string, nowTime int64) (key string) {

	key = ip + "_" + strconv.FormatInt(nowTime, 10)
	return key
}
