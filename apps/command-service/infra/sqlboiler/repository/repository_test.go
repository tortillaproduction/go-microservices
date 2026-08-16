package repository

import (
	"os"
	"path/filepath"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/tortillaproduction/go-microservices/apps/command-service/infra/sqlboiler/handler"
)

func TestRepImplPackage(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "infra/sqlboiler/repository package tests")
}

var _ = BeforeSuite(func() {
	absPath, _ := filepath.Abs("../config/database.toml")
	os.Setenv("DATABASE_TOML_PATH", absPath)
	err := handler.DBConnect()
	Expect(err).NotTo(HaveOccurred(), "Failed to connect to the database; aborting the test suite.")
})
