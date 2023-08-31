package models

type Point struct {
	PointID       string  `json:"point_id"`
	VideoDuration float64 `json:"video_duration"`
}

type Points struct {
	Point map[string]Point
}

type Chapter struct {
	Points []Point `json:"points"`
}

type Chapters struct {
	Chapter map[string]Chapter
}

type PreviewResultObject struct {
	Chapters []Chapter `json:"chapters"`
}

type PreviewJSON struct {
	ResultObject PreviewResultObject `json:"resultObject"`
}
