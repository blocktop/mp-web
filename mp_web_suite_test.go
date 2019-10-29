package main

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMpWeb(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "MpWeb Suite")
}
