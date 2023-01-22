package api

import (
	"excel/common"
	"fmt"
	"github.com/bytedance/sonic/decoder"
	"net/http"
	"time"
)

type API struct {
	*common.Inject
}

type M map[string]interface{}

type Data struct {
	Records []Record `json:"records"`
}

type Record struct {
	Cos   `json:"cos"`
	Event `json:"event"`
}

type Cos struct {
	CosSchemaVersion  string `json:"cosSchemaVersion"`
	CosObject         `json:"cosObject"`
	CosBucket         `json:"cosBucket"`
	CosNotificationId string `json:"cosNotificationId"`
}

type CosObject struct {
	Url  string `json:"url"`
	Meta M      `json:"meta"`
	Vid  string `json:"vid"`
	Key  string `json:"key"`
	Size int64  `json:"size"`
}

type CosBucket struct {
	Region string `json:"region"`
	Name   string `json:"name"`
	Appid  string `json:"appid"`
}

type Event struct {
	EventName         string `json:"eventName"`
	EventVersion      string `json:"eventVersion"`
	EventTime         int64  `json:"eventTime"`
	EventSource       string `json:"eventSource"`
	RequestParameters M      `json:"requestParameters"`
	EventQueue        string `json:"eventQueue"`
	ReservedInfo      string `json:"reservedInfo"`
	Reqid             int64  `json:"reqid"`
}

func (x *API) EventInvoke(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		return
	}

	var data Data
	if err := decoder.NewStreamDecoder(req.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for _, record := range data.Records {
		fmt.Println(record.Cos.Url)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(`已触发: %s`, time.Now())))
}
