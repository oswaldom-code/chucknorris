package services

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const (
	// API_CONSUMED_URL is the url of the api consumed
	API_CONSUMED_URL = "https://api.chucknorris.io/jokes/random"
	TARGET_OBJECTS = 25
)

// UnmarshalCall unmarshals an instance of Call from the specified map of raw messages.
func UnmarshalCall(data []byte) (Call, error) {
	var r Call
	err := json.Unmarshal(data, &r)
	return r, err
}

// Marshal returns the JSON encoding of Call.
func (r *Call) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// Call is the struct that contains the information of the call.
type Call struct {
	Categories []interface{} `json:"categories"`
	CreatedAt  string        `json:"created_at"`
	IconURL    string        `json:"icon_url"`
	ID         string        `json:"id"`
	UpdatedAt  string        `json:"updated_at"`
	URL        string        `json:"url"`
	Value      string        `json:"value"`
}

// HealthService is the interface for the health service.
type CallServices interface {
	Call() ([]Call, error)
	checkId([]Call, Call) bool
	doHttpRequest()(*http.Response, error)
}

// CallServiceImpl is the implementation of the Call service.
type CallServiceImpl struct{}

// NewCallService creates a new Call service.
func NewCallService() CallServices {
	return &CallServiceImpl{}
}

// checkId checks if the id is already in the array
func (s *CallServiceImpl)checkId(current []Call, call Call) bool {
	for _, c := range current {
		if c.ID == call.ID {
			return true
		}
	}
	return false
}

// doHttpRequest does the http request
func (s *CallServiceImpl)doHttpRequest()(*http.Response, error) {
	req, err := http.NewRequest("GET", API_CONSUMED_URL, nil)
	if err != nil {
		return &http.Response{}, err 
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return &http.Response{}, err
	}
	return res, nil
}

// Ping returns the health of the service.
func (s *CallServiceImpl) Call() ([]Call, error) {
	callResponseArray := []Call{}
	for len(callResponseArray) < TARGET_OBJECTS {
		// doHttpRequest returns the response and the error
		res, err := s.doHttpRequest()
		if err != nil {
			return []Call{}, err
		}
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return []Call{}, err
		}
		call, err := UnmarshalCall(body)
		if err != nil {
			return []Call{}, err
		}
		if !s.checkId(callResponseArray, call) {
		callResponseArray = append(callResponseArray, call)
		}
	}
	return callResponseArray, nil
}
