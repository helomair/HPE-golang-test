package helper

import (
	"strings"
	"time"
)

func ParseDataToParams(datas string) map[string]string {
	params := map[string]string{}

	for _, v := range strings.Split(datas, "&") {
		data := strings.Split(v, "=")
		params[data[0]] = data[1]
	}

	return params
}

func JoinParamsToData(params map[string]string) string {
	data := ""

	for key, value := range params {
		data += key + "=" + value + "&"
	}

	return data
}

func ParseHtmlDateTime(datetime string) time.Time {
	timeZone, _ := time.LoadLocation("Asia/Taipei")
	timeLayout := "2006-01-02T15:04"
	ret, _ := time.ParseInLocation(timeLayout, datetime, timeZone)
	return ret
}
