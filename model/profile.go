package model

import "encoding/json"

type Profile struct {
	Name string
	Age string
	AvatarUrl string
	BasicInfo string
	DetailInfo string
	Education string
	Gender string
	Height string
	Salary string
	Marriage string
}

func FromJsonObj(o interface{}) (Profile, error) {
	var profile Profile
	bytes, err := json.Marshal(o)
	if err != nil {
		return profile, err
	}
	err = json.Unmarshal(bytes, &profile)
	return profile, err
}