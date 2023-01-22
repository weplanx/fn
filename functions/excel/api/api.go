package api

import (
	"encoding/json"
	"excel/common"
	"fmt"
	"net/http"
	"time"
)

type API struct {
	*common.Inject
}

func (x *API) EventInvoke(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		return
	}

	var data map[string]interface{}
	if err := json.NewDecoder(req.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println(data)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(`已触发: %s`, time.Now())))
}
