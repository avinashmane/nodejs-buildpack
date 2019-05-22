package integration_test

import (
	"path/filepath"

	"github.com/cloudfoundry/libbuildpack/cutlass"

	. "github.com/onsi/ginkgo"
)

var _ = Describe("CF NodeJS Buildpack", func() {
	var app *cutlass.App
	AfterEach(func() {
		if app != nil {
			app.Destroy()
		}
		app = nil
	})

	Context("deploying a Node.js app with mysql", func() {
		BeforeEach(func() {
			app = cutlass.New(filepath.Join(bpDir, "v2b_integration", "testdata", "with_mysql"))
		})

		XIt("should push the app with mysql successfully", func() {
			PushAppAndConfirm(app)
		})
	})
})
