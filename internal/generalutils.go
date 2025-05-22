package internal

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateRandom4DigitString() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%04d", 1000+rand.Intn(9000))
}
