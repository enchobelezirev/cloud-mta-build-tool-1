package commands

import (
	"fmt"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/SAP/cloud-mta-build-tool/internal/fs"
	"github.com/SAP/cloud-mta-build-tool/internal/platform"
)

var _ = Describe("Commands", func() {

	BeforeEach(func() {
		mtadCmdTrg = getTestPath("result")
		metaCmdTrg = getTestPath("result")
		mtarCmdTrg = getTestPath("result")
		err := os.Mkdir(mtadCmdTrg, os.ModePerm)
		if err != nil {
			fmt.Println(err)
		}
	})

	AfterEach(func() {
		os.RemoveAll(mtadCmdTrg)
	})

	var _ = Describe("Generate commands call", func() {

		var ep dir.Loc

		AfterEach(func() {
			mtadCmdDesc = ""
			os.RemoveAll(getTestPath("result"))
		})

		It("Generate Meta", func() {
			os.MkdirAll(getTestPath("result", "mtahtml5", "testapp"), os.ModePerm)
			os.MkdirAll(getTestPath("result", "mtahtml5", "ui5app2"), os.ModePerm)
			metaCmdSrc = getTestPath("mtahtml5")
			ep = dir.Loc{SourcePath: metaCmdSrc, TargetPath: metaCmdTrg}
			Ω(metaCmd.RunE(nil, []string{})).Should(Succeed())
			Ω(ep.GetMtadPath()).Should(BeAnExistingFile())
		})
		It("Generate Mtad - Sanity", func() {
			mtadCmdSrc = getTestPath("mtahtml5")
			mtadCmdPlatform = "cf"
			Ω(mtadCmd.RunE(nil, []string{})).Should(Succeed())
			Ω(ep.GetMtadPath()).Should(BeAnExistingFile())
		})
		It("Generate Mtad - Invalid deployment descriptor", func() {
			mtadCmdSrc = getTestPath("mtahtml5")
			mtadCmdDesc = "xx"
			Ω(mtadCmd.RunE(nil, []string{})).Should(HaveOccurred())
		})
		It("Generate Mtad - Invalid source", func() {
			mtadCmdSrc = getTestPath("mtahtml6")
			Ω(mtadCmd.RunE(nil, []string{})).Should(HaveOccurred())
		})
		It("Generate Mtad - Invalid platform configuration", func() {
			mtadCmdSrc = getTestPath("mtahtml5")
			config := platform.PlatformConfig
			platform.PlatformConfig = []byte("wrong config")
			mtadCmdPlatform = "cf"
			Ω(mtadCmd.RunE(nil, []string{})).Should(HaveOccurred())
			platform.PlatformConfig = config
		})
		It("Generate Mtar", func() {
			os.MkdirAll(getTestPath("result", "mtahtml5", "testapp"), os.ModePerm)
			os.MkdirAll(getTestPath("result", "mtahtml5", "ui5app2"), os.ModePerm)
			mtarCmdSrc = getTestPath("mtahtml5")
			Ω(metaCmd.RunE(nil, []string{})).Should(Succeed())
			Ω(mtarCmd.RunE(nil, []string{})).Should(Succeed())
			Ω(getTestPath("result", "mta_archives", "mtahtml5_0.0.1.mtar")).Should(BeAnExistingFile())
		})
	})

})
