package helper

import "strings"

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
		data += key + "=" + value
	}

	return data
}
