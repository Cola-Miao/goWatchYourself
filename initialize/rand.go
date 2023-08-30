package initialize

import (
	"goWatchYourself/global"
	"math/rand"
	"time"
)

func initRand() {
	source := rand.NewSource(time.Now().Unix())
	rd := rand.New(source)
	global.Rand = rd
}
