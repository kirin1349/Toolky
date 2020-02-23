package toolky

import "time"

func GetTimestamp() (ts int64) {
	return time.Now().Unix()
}

func GetTimestampMS() (ts int64) {
	return time.Now().UnixNano() / 1000000
}
