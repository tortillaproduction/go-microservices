package server

import (
	"os"
	"path/filepath"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/tortillaproduction/go-microservices/apps/command-service/infra/sqlboiler/handler"
)

func TestHelperPackage(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "presen/server package tests")
}

var _ = BeforeSuite(func() {
	absPath, err := filepath.Abs("../../infra/sqlboiler/config/database.toml")
	Expect(err).NotTo(HaveOccurred())

	os.Setenv("DATABASE_TOML_PATH", absPath)

	err = handler.DBConnect()
	Expect(err).NotTo(HaveOccurred(), "Failed to connect to the database; aborting the test suite.")
})
