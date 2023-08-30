package utils

import (
	"errors"
	"fmt"
)

func ParseFreeMap(m map[string]any) (videos map[float64]float64, err error) {
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

func ParsePreviewMap(m map[string]any) (chapters map[string]float64, err error) {
	resultObject, ok := m["resultObject"]
	if !ok {
		err = errors.New("can`t parse resultObject")
		return
	}
	chapterList, ok := resultObject.(map[string]any)["chapters"]
	if !ok {
		err = errors.New("can`t parse chapters")
		return
	}
	chapters = make(map[string]float64)
	for _, v := range chapterList.([]any) {
		points := v.(map[string]any)["points"]
		for _, v2 := range points.([]any) {
			var pointID string
			var videoDuration float64
			for key, val := range v2.(map[string]any) {
				if key == "point_id" {
					pointID = val.(string)
				}
				if key == "video_duration" {
					videoDuration = val.(float64)
				}
			}
			chapters[pointID] = videoDuration
			fmt.Println(pointID, videoDuration)
		}
	}

	return
}
