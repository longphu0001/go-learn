package tut

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGoconvey(t *testing.T) {

	// Only pass t into top-level Convey calls
	Convey("Given some integer with a starting value", t, func() {
		x := 1

		Convey("When the integer is incremented", func() {
			x++

			Convey("The value should be greater by one", func() {
				So(x, ShouldEqual, 2)
			})
		})
	})

	// https://github.com/smartystreets/goconvey/wiki/Assertions#strings
	Convey("Comparing two variables", t, func() {
		helloGoconvey := ""

		Convey(`helloGoconvey should be blank`, func() {
			So(helloGoconvey, ShouldBeBlank)
		})

		helloGoconvey = HelloGoconvey()

		Convey("helloGoconvey should not be blank", func() {
			So(helloGoconvey, ShouldNotBeBlank)
		})

		Convey(`"Hello, Goconvey!" should NOT equal "Hello, World!"`, func() {
			So(helloGoconvey, ShouldNotEqual, "Hello, World!")
		})

		Convey(`"Hello, Goconvey!" should NOT equal "Hello, Goconvey!"`, func() {
			So(helloGoconvey, ShouldEqual, "Hello, Goconvey!")
		})

		Convey(`helloGoconvey Should Start With "Hello"`, func() {
			So(helloGoconvey, ShouldStartWith, "Hello")
		})

		Convey(`helloGoconvey Should End With "Goconvey!"`, func() {
			So(helloGoconvey, ShouldEndWith, "Goconvey!")
		})

	})
}
