package dice

import (
	"fmt"
	"math/rand"
	"time"
)

// Roll rolls a n sided regular dice
// This a temp thing however we should read this:
// https://www.reddit.com/r/golang/comments/47y2ml/dont_seed_the_global_random_in_your_package/
func Roll(n int) (int, error) {

	if n < 1 {
		err := fmt.Errorf("Roll: Got and invalid number of side: %s", n)
		return -1, err
	}

	seed := rand.NewSource(time.Now().UnixNano())
	r := rand.New(seed)

	return r.Intn(n), nil
}
