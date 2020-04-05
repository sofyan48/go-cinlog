package v1

import (
	"encoding/json"

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

// Save ...
func (v1 V1Session) Save(uuid, status string, data map[string]interface{}) (*entity.LoggerEventHistory, error) {
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
	if err != nil {
		return nil, err
	}
	result := &entity.LoggerEventHistory{}
	err = json.Unmarshal(resultData, result)
	return result, nil
}

// Get ...
func (v1 V1Session) Get(uuid string) (*entity.LoggerEventHistory, error) {
	path := v1.URL + "/get"
	payload := &entity.GetLoggerRequest{}
	payload.Action = v1.Service
	payload.UUID = uuid
	bytePayload, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	resultData, err := v1.Requester.POST(path, "", bytePayload)
	if err != nil {
		return nil, err
	}
	result := &entity.LoggerEventHistory{}
	err = json.Unmarshal(resultData, result)
	return result, nil
}

// Find ...
func (v1 V1Session) Find(uuid, action string) (*entity.LoggerEventHistory, error) {
	path := v1.URL + "/get"
	payload := &entity.GetLoggerRequest{}
	payload.Action = action
	payload.UUID = uuid
	bytePayload, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	resultData, err := v1.Requester.POST(path, "", bytePayload)
	if err != nil {
		return nil, err
	}
	result := &entity.LoggerEventHistory{}
	err = json.Unmarshal(resultData, result)
	return result, nil
}

// All ...
func (v1 V1Session) All() ([]entity.LoggerEventHistory, error) {
	path := v1.URL + "/all"
	payload := &entity.GetAllLoggerRequest{}
	payload.Action = v1.Service
	bytePayload, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	resultData, err := v1.Requester.POST(path, "", bytePayload)
	if err != nil {
		return nil, err
	}
	results := make([]entity.LoggerEventHistory, 0)
	err = json.Unmarshal(resultData, &results)
	if err != nil {
		return nil, err
	}
	return results, nil
}
