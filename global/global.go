package global

import (
	"math/rand"
	"net/http"
)

var (
	Client    *http.Client
	CourseID  int
	PreviewID string
	Model     int
	Rand      *rand.Rand
)
