package impl_test

import (
	"os"
	"path/filepath"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/tortillaproduction/go-microservices/apps/command-service/infra/sqlboiler/handler"
)

func TestSrvImplPackage(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "application/impl package tests")
}

var _ = BeforeSuite(func() {
	absPath, _ := filepath.Abs("../../infra/sqlboiler/config/database.toml")
	os.Setenv("DATABASE_TOML_PATH", absPath)
	err := handler.DBConnect()
	Expect(err).NotTo(HaveOccurred(), "aborting tests because DB connection failed")
})
