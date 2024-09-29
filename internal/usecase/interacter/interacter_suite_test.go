package interacter_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestInteracter(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Interacter Suite")
}
