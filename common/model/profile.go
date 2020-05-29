package model

import "encoding/json"

type Profile struct {
	Name          string
	Gender        string
	Age           int
	Height        int
	Weight        int
	Income        string
	Marriage      string
	Education     string
	Occupation    string
	Hukou         string
	Constellation string
	House         string
	Car           string
	WorkPlace     string
}

func FromJsonObj(o interface{}) (profile Profile, err error) {
	s, err := json.Marshal(o)
	if err != nil {
		return
	}
	err = json.Unmarshal(s, &profile)
	return
}
