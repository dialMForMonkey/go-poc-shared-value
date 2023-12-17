package core

import "math/rand"

func RandomResult() bool {

	if rand.Intn(2) == 0 {
		return false
	}
	return true

}
