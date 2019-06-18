package coin

// Coin class, a boolean would be simplier but this class is for
// understanding classes and serialization in go

import (
	"fmt"
	"math/rand"
	"time"
)

// A coin represents an object for storing a boolean value
type coin struct {
	side bool // 0 = tails, 1 = heads
}

// New creates instance of coin struct
func New() coin {
	rand.Seed(time.Now().UnixNano())
	c := coin{rand.Float32() < 0.5}
	return c
}

// flip flips coin
func flip(c coin) {
	c.side = rand.Float32() < 0.5
}

// look checks if the coin was heads of tails.
func (c coin) look() {
	if c.side {
		fmt.Printf("Landed tails")
	} else {
		fmt.Printf("Landed heads")
	}
}
