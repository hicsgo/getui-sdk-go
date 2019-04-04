package getui

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func (c Account) PushSingle(value PushSingleParams, os string) map[string]interface{} {
	headers := http.Header{
		"authtoken":    []string{c.Auth},
		"Content-Type": []string{"application/json"},
	}
	value.AppKey = c.AppKey
	var body interface{}
	if os == "android" {
		body = NewPushSingleContent(value)
	} else {
		body = NewPushSingleIosContent(value)
	}

	request, err := request(http.MethodPost, fmt.Sprintf(PUSHSINGLE, c.AppID), headers, body)
	var result = make(map[string]interface{})
	if err != nil {
		log.Panic(err)
		return result
	}
	content, err := HttpClientDo(request)
	if err != nil {
		log.Panic(err)
		return result
	}
	json.Unmarshal(content, &result)
	return result

}
