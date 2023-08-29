package utils

import (
	"errors"
)

func ParseMap(m map[string]any) (videos map[float64]float64, err error) {
	resultObject, ok := m["resultObject"]
	if !ok {
		err = errors.New("can`t parse resultObject")
		return
	}
	videoList, ok := resultObject.(map[string]any)["videos"]
	if !ok {
		err = errors.New("can`t parse videos")
	}
	videos = make(map[float64]float64)
	for _, v := range videoList.([]any) {
		id := v.(map[string]any)["id"]
		duration := v.(map[string]any)["duration"]
		videos[id.(float64)] = duration.(float64)
	}

	return
}
