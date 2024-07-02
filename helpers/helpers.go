package helpers

import "math/rand"

type SomeType struct {
	TypeName        string
	TypeDescription string
	TypeNumber      int
}

func RandomInt(n int) int {
	value := rand.Intn(n)
	return value
}
