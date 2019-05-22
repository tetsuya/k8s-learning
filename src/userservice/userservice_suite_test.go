package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestUserservice(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Userservice Suite")
}
