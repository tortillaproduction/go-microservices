package handler

import (
	"os"
	"path/filepath"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestConn(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "infra/sqlboiler/handler package tests")
}

var _ = Describe("Database connection", func() {
	It("returns nil if the connection is successful", Label("db-connection"), func() {
		absPath, _ := filepath.Abs("../config/database.toml")

		os.Setenv("DATABASE_TOML_PATH", absPath)
		result := DBConnect()
		Expect(result).To(BeNil())
	})
})
