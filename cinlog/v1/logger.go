package v1

import (
	"encoding/json"
	"fmt"

	"github.com/sofyan48/go-cinlog/entity"
	"github.com/sofyan48/go-cinlog/libs"
)

type V1Session struct {
	URL       string
	Service   string
	Requester libs.RequesterInterface
}

// VERSION ... v
const VERSION = "v1"

// V1SessionHandler ...
func V1SessionHandler(sess *entity.SessionGlobal) *V1Session {
	return &V1Session{
		URL:       "http://" + sess.URL + "/" + VERSION + "/logger",
		Service:   sess.Service,
		Requester: libs.GetRequesterLibs(),
	}
}

// SaveLogger ...
func (v1 V1Session) SaveLogger(uuid, status string, data map[string]interface{}) (*entity.LoggerEventHistory, error) {
	path := v1.URL
	payload := &entity.SaveLogger{}
	payload.Data = data
	payload.UUID = uuid
	payload.Action = v1.Service
	payload.Status = status
	bytePayload, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	resultData, err := v1.Requester.POST(path, "", bytePayload)
	fmt.Println(string(resultData))
	if err != nil {
		return nil, err
	}
	result := &entity.LoggerEventHistory{}
	err = json.Unmarshal(resultData, result)
	return result, nil
}
