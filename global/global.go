package global

import (
	"math/rand"
	"net/http"
)

var (
	Client   *http.Client
	CourseID int
	Rand     *rand.Rand
)
