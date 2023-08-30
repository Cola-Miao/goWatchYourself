package utils

import (
	"errors"
	"goWatchYourself/global"
	"time"
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
		return
	}
	videos = make(map[float64]float64)
	for _, v := range videoList.([]any) {
		id := v.(map[string]any)["id"]
		duration := v.(map[string]any)["duration"]
		videos[id.(float64)] = duration.(float64)
	}

	return
}

func RandomSleep() {
	t := global.Rand.Intn(1000)
	time.Sleep(time.Duration(t) * time.Millisecond)
}
