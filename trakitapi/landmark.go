package trakitapi

import (
	"encoding/json"
	"reflect"
)

type Landmark struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Address     string  `json:"address"`
	Description string  `json:"desciption"`
	Icon        string  `json:"icon"`
	ShowOnMap   bool    `json:"showOnMap"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	Status      string  `json:"status"`
}

func (landmark Landmark) ToRow() []string {
	v := reflect.ValueOf(landmark)
	return getStrValues(v)
}

func GetAllLandmarks() ([]Landmark, error) {
	var landmarks []Landmark

	res := Get("landmark")

	defer res.Body.Close()

	err := json.NewDecoder(res.Body).Decode(&landmarks)
	if err != nil {
		return nil, err
	}

	return landmarks, nil
}
