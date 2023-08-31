package models

type Video struct {
	ID       float64 `json:"id"`
	Duration float64 `json:"duration"`
}

type Videos struct {
	Video map[string]Video
}

type FreeResultObject struct {
	Videos []Video `json:"videos"`
}

type FreeJSON struct {
	ResultObject FreeResultObject `json:"resultObject"`
}
