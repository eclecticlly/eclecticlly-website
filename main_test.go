package main

import (
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	. "github.com/smartystreets/goconvey/convey"
)

func TestHttpRoute(t *testing.T) {
	// Define a structure for specifying input and output
	// data of a single test case. This structure is then used
	// to create a so called test map, which contains all test
	// cases, that should be run for testing this function
	tests := []struct {
		// Test input
		route string

		// Expected output
		expectedError bool
		expectedCode  int
		expectedBody  string
	}{
		{
			route:         "/",
			expectedError: false,
			expectedCode:  200,
			expectedBody:  "Hello From Fiber Jet Engine",
		},
		{
			route:         "/i-dont-exist",
			expectedError: false,
			expectedCode:  404,
			expectedBody:  "Not found",
		},
	}

	// Setup the app as it is done in the main function
	app := Setup()

	Convey("Given the user want to navigate the site", t, func() {

		Convey("When the user go to Home page", func() {
			res, body, err := route(app, tests[0].route)

			Convey("Then he can read \"Hello From Fiber Jet Engine\"", func() {
				So(tests[0].expectedError, ShouldEqual, err != nil)
				So(tests[0].expectedCode, ShouldEqual, res.StatusCode)
				So(string(body), ShouldContainSubstring, tests[0].expectedBody)
			})

		})

		Convey("When the use try to go to a non existent page", func() {
			res, body, err := route(app, tests[1].route)

			Convey("Then he can see the 404 page", nil)
			So(tests[1].expectedError, ShouldEqual, err != nil)
			So(tests[1].expectedCode, ShouldEqual, res.StatusCode)
			So(string(body), ShouldContainSubstring, tests[1].expectedBody)
		})

	})

}

func route(app *fiber.App, url string) (*http.Response, string, error) {
	req, _ := http.NewRequest(
		"GET",
		url,
		nil,
	)

	res, err := app.Test(req, -1)

	body, _ := ioutil.ReadAll(res.Body)

	return res, string(body), err
}
