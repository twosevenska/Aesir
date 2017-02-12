package dice

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

const (
	d20 = 20
	di  = -1
)

func TestRoll(t *testing.T) {
	Convey("Test the dice roller", t, func() {
		Convey("Give a positive number of sides", func() {
			r, err := Roll(d20)
			So(r, ShouldBeGreaterThan, 0)
			So(err, ShouldBeNil)
		})

		Convey("Give an invalid number of sides", func() {
			r, err := Roll(di)
			So(r, ShouldBeLessThan, 0)
			So(err, ShouldNotBeNil)
		})
	})
}
