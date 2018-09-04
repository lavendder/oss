package server

import "time"

// server global variable
var SERVER = make(map[string]interface{})

func SetRequestId(id string) {
	SERVER["RequestId"] = id
	return
}

func GetDataByKey(key string) interface{} {
	return SERVER[key]
}

func GetRequestId() string {
	requestId, ok := SERVER["RequestId"]
	if (ok) {
		return requestId.(string)
	}
	return ""
}

// set time for beginning of one request
func SetRequestTime()  {
	currentTime := time.Now()
	SERVER["CurrentTime"] = currentTime
	return
}

func GetRequestTime() time.Time {
	currentTime, ok := SERVER["CurrentTime"]
	if (ok) {
		return currentTime.(time.Time)
	}
	return time.Time{}
}