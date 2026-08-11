package categories

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

// entry point for tests
func TestEntityPackage(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "domain/models/categories package tests") // run all registered tests (specs).
}
