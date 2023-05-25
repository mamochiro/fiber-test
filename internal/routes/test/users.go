package test

import (
	"bytes"
	"encoding/json"
	"fiber-test/internal/models"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
)

func (suite *PackageTestSuite) TestGetUserRoute() {
	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	resp, _ := suite.app.Test(req)

	if resp.StatusCode == fiber.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		fmt.Println(string(body))
	}
}

func (suite *PackageTestSuite) TestCreateUserRoute() {
	create := models.CreateUserRequest{
		Name:  "mark",
		Email: "mark@email.com",
	}

	data, err := json.Marshal(create)
	if err != nil {
		log.Fatal(err)
	}
	reader := bytes.NewReader(data)
	req := httptest.NewRequest(http.MethodPost, "/users", reader)
	req.Header.Add("Content-Type", "application/json")
	resp, _ := suite.app.Test(req)

	if resp.StatusCode == fiber.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		fmt.Println(string(body))
	}
}
