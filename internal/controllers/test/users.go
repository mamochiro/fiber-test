package test

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
)

func (suite *PackageTestSuite) TestGetUser() {
	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	resp, _ := suite.app.Test(req)

	if resp.StatusCode == fiber.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(body)) // => Hello, World!
	}
}
