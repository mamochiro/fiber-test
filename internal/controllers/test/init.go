package test

import (
	"fiber-test/internal/app"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/suite"
)

type PackageTestSuite struct {
	suite.Suite
	app *fiber.App
}

func (suite *PackageTestSuite) SetupTest() {
	suite.app = app.SetupApp()
}
