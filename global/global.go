package global

import (
	"math/rand"
	"net/http"
)

var (
	Client        *http.Client
	CourseID      int
	PreviewID     string
	Model         int
	Rand          *rand.Rand
	Version       = "0.5"
	DebugUserName = "15526811889"
	DebugPassword = "@123456ABcdqqqq"
)
