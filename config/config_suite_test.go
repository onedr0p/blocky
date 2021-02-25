package config

import (
	. "blocky/log"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestConfig(t *testing.T) {
	ConfigureLogger("Warn", "Text")
	RegisterFailHandler(Fail)
	RunSpecs(t, "Config Suite")
}
