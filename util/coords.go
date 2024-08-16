package util

import "math/rand/v2"

func GetRandomCoords(maxWidth, maxHeight int) (int, int) {
	x := rand.IntN(maxWidth-0) + 0
	y := rand.IntN(maxHeight-0) + 0

	return x, y
}
