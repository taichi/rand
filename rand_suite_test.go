package rand_test

import (
	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/config"
	. "github.com/onsi/gomega"
	"os"
	"runtime"

	"testing"
)

func TestRand(t *testing.T) {
	RegisterFailHandler(Fail)
	Configure()
	RunSpecs(t, "Rand Suite")
}

func Configure() {
	config.DefaultReporterConfig.NoColor = runtime.GOOS == "windows" && len(os.Getenv("ANSICON")) < 1
}
