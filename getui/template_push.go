package getui

type PushSingleParams struct {
	AppKey              string   `json:"appkey"`
	Text                string   `json:"text"`
	Title               string   `json:"title"`
	TransmissionContent string   `json:"transmission_content"`
	Cid                 []string `json:"cid"`
}

type PushSingleContent struct {
	Message      Message      `json:"message"`
	Notification Notification `json:"notification"`
	Cid          string       `json:"cid"`
	RequestID    string       `json:"requestid"`
}

type Message struct {
	AppKey            string `json:"appkey"`
	IsOffline         bool   `json:"is_offline"`
	OfflineExpireTime int    `json:"offline_expire_time"`
	MsgType           string `json:"msgtype"`
}

type Notification struct {
	Style               Style  `json:"style"`
	TransmissionType    bool   `json:"transmission_type"`
	TransmissionContent string `json:"transmission_content"`
}

func NewPushSingleContent(params PushSingleParams) interface{} {
	return PushSingleContent{
		Message: Message{params.AppKey, true, 10000000, "notification"},
		Notification: Notification{
			Style:               Style{0, params.Text, params.Title},
			TransmissionType:    true,
			TransmissionContent: params.TransmissionContent,
		},
		Cid:       params.Cid[0],
		RequestID: timeStamp(),
	}
}

// "push_info": {
// 	"aps": {
// 		"alert": {
// 			"title": "测试",
// 			"body": "测试"
// 		}
// 	},
// 	"payload": "payload"
// },
type PushInfo struct {
	Aps     *Aps   `json:"aps"`
	PayLoad string `json:"payload"`
}
type Aps struct {
	Alert *Alert `json:"alert"`
}
type Alert struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

//推送IOS内容体
type PushSingleIOSContent struct {
	Message      Message       `json:"message"`
	Transmission *Transmission `json:"transmission"`
	PushInfo     *PushInfo     `json:"push_info"`
	Cid          string        `json:"cid"`
	RequestID    string        `json:"requestid"`
}
type Transmission struct {
	TransmissionType    bool   `json:"transmission_type"`
	TransmissionContent string `json:"transmission_content"`
}

func NewPushSingleIosContent(params PushSingleParams) interface{} {
	return PushSingleIOSContent{
		Message: Message{params.AppKey, true, 10000000, "transmission"},
		Transmission: &Transmission{
			TransmissionType:    true,
			TransmissionContent: params.TransmissionContent,
		},
		Cid:       params.Cid[0],
		RequestID: timeStamp(),
		PushInfo: &PushInfo{
			PayLoad: params.TransmissionContent,
			Aps: &Aps{
				Alert: &Alert{
					Title: params.Title,
					Body:  params.Text,
				},
			},
		},
	}
}

type Style struct {
	Type  int    `json:"type"`
	Text  string `json:"text"`
	Title string `json:"title"`
}

var IosPushSingleTemplate = `
{
	"message": {
		"appkey": "tH9Q0kQK3a7Y7l4SW8Nou6",
		"is_offline": true,
		"msgtype": "transmission"
	},
	"transmission": {
		"transmission_type": false,
		"transmission_content": "测试"
	},
	"push_info": {
		"aps": {
			"alert": {
				"title": "测试",
				"body": "测试"
			}
		},
		"payload": "payload"
	},
	"cid": "3f3d6b0426b319b52268e6d9d8210a3c",
	"requestid": "1211111111111111111111113xxxx1q"
}`

var AndroidPushSingleTemplate = `
{
	"message": {
		"appkey": {{.AppKey}},
		"is_offline": true,
		"offline_expire_time":10000000,
		"msgtype": "notification"
	}
	"notification": {
		"style": {
        	"type": 0,
            "text": {{.Text}},
            "title": {{.Title}}
		},
        "transmission_type": true,
        "transmission_content": {{.TransmissionContent}}
	},
    "cid": {{.Cid}}
}
`

type SaveListBodyContent struct {
	Message      Message      `json:"message"`
	Notification Notification `json:"notification"`
}

func NewSaveListBodyContent(params PushSingleParams) SaveListBodyContent {
	return SaveListBodyContent{
		Message: Message{params.AppKey, true, 10000000, "notification"},
		Notification: Notification{
			Style:               Style{0, params.Text, params.Title},
			TransmissionType:    true,
			TransmissionContent: params.TransmissionContent,
		},
	}
}

func SaveListBodyContentByMap(params map[string]string) SaveListBodyContent {
	return SaveListBodyContent{
		Message: Message{params["appkey"], true, 10000000, "notification"},
		Notification: Notification{
			Style:               Style{0, params["text"], params["title"]},
			TransmissionType:    true,
			TransmissionContent: params["transmission_content"],
		},
	}
}

var SaveListBody = `
{
                   "message": {
                   "appkey": {{.AppKey}},
                   "is_offline": true,
                   "offline_expire_time":10000000,
                   "msgtype": "notification"
                },
                "notification": {
                    "style": {
                        "type": 0,
                        "text": {{.Text}},
                        "title": {{.Title}}
                    },
                    "transmission_type": true,
                    "transmission_content": {{.TransmissionContent}}
                }
           }
`
