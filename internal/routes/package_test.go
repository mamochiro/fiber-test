package routes_test

import (
	"fiber-test/internal/routes/test"
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestPackageTestSuite(t *testing.T) {
	suite.Run(t, new(test.PackageTestSuite))
}
