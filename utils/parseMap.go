package utils

import (
	"encoding/json"
	"goWatchYourself/models"
)

func ParseFreeMap(body []byte) (videos map[float64]float64, err error) {
	var freeJSON models.FreeJSON
	err = json.Unmarshal(body, &freeJSON)
	if err != nil {
		return
	}
	videos = make(map[float64]float64)
	for _, video := range freeJSON.ResultObject.Videos {
		videos[video.ID] = video.Duration
	}

	return
}

func NewParsePreviewMap(body []byte) (points map[string]float64, err error) {
	var previewJSON models.PreviewJSON
	err = json.Unmarshal(body, &previewJSON)
	if err != nil {
		return
	}
	points = make(map[string]float64)
	for _, chapter := range previewJSON.ResultObject.Chapters {
		for _, point := range chapter.Points {
			points[point.PointID] = point.VideoDuration
		}
	}

	return
}
