package katapotter

import (
	"fmt"
	"testing"
)

func TestBuyHarryPotter(t *testing.T) {
	testCases := []struct {
		s1    int
		s2    int
		s3    int
		s4    int
		s5    int
		price float32
	}{
		{1, 0, 0, 0, 0, 8.0},
		{1, 1, 0, 0, 0, 15.2},
		{1, 1, 1, 0, 0, 21.6},
		{1, 1, 1, 1, 0, 25.6},
		{1, 1, 1, 1, 1, 30.0},
	}
	for _, c := range testCases {
		t.Run(fmt.Sprintf(
			"S1: %d, S2: %d, S3: %d, S4: %d, S5: %d",
			c.s1, c.s2, c.s3, c.s4, c.s5,
		), func(t *testing.T) {
			result := BuyHarryPotter(c.s1, c.s2, c.s3, c.s4, c.s5)

			if result != c.price {
				t.Errorf("Got %5.2f, want %5.2f", result, c.price)
			}
		})
	}
}
