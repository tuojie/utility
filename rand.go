package utility

import rand "math/rand"
import "time"

var r *rand.Rand = nil

func Init() {
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func RandInt32(max int32) int32 {
	return int32(r.Intn(int(max)))
}
