package settingService

import "encoding/json"

// ChangeNickNameRequest is a structure for request to change nick-name
type ChangeNickNameRequest struct {
	Header RequestHeader `json:"header"`
	Body   ChangeNickNameBody `json:"body"`
}

// ChangeNickNameBody is a structure for passing new nick-name
type ChangeNickNameBody struct {
	NickName string
}

// UnmarshalChangePassRequest is a function for unmarshaling request (from [] byte JSON to ChangePassRequest structure)
func UnmarshalChangeNickNameRequest(changeNickNameRequest [] byte) (*ChangeNickNameRequest, error) {
	var RequestStructure *ChangeNickNameRequest
	err := json.Unmarshal(changeNickNameRequest, &RequestStructure)
	if err != nil {
		loger.Log.Errorf("Error has occured: ", err)
		return RequestStructure, err
	}
	return RequestStructure, nil

}

