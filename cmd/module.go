package commands

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/SAP/cloud-mta-build-tool/internal/artifacts"
)

// flags of pack command
var packCmdSrc string
var packCmdTrg string
var packCmdDesc string
var packCmdModule string
var packCmdPlatform string

// flags of build command
var buildCmdSrc string
var buildCmdTrg string
var buildCmdDesc string
var buildCmdModule string
var buildCmdPlatform string

func init() {

	// set flags of command pack Module
	packModuleCmd.Flags().StringVarP(&packCmdSrc, "source", "s", "",
		"the path to the MTA project; the current path is default")
	packModuleCmd.Flags().StringVarP(&packCmdTrg, "target", "t", "",
		"the path to the MBT results folder; the current path is default")
	packModuleCmd.Flags().StringVarP(&packCmdDesc, "desc", "d", "",
		"the MTA descriptor; supported values: dev (development descriptor, default value) and dep (deployment descriptor)")
	packModuleCmd.Flags().StringVarP(&packCmdModule, "module", "m", "",
		"the name of the module")
	packModuleCmd.Flags().StringVarP(&packCmdPlatform, "platform", "p", "",
		"the deployment platform; supported plaforms: cf, xsa")

	// set flags of command build Module
	buildModuleCmd.Flags().StringVarP(&buildCmdSrc, "source", "s", "",
		"the path to the MTA project; the current path is default")
	buildModuleCmd.Flags().StringVarP(&buildCmdTrg, "target", "t", "",
		"the path to the MBT results folder; the current path is default")
	buildModuleCmd.Flags().StringVarP(&buildCmdDesc, "desc", "d", "",
		"the MTA descriptor; supported values: dev (development descriptor, default value) and dep (deployment descriptor)")
	buildModuleCmd.Flags().StringVarP(&buildCmdModule, "module", "m", "",
		"the name of the module")
	buildModuleCmd.Flags().StringVarP(&buildCmdPlatform, "platform", "p", "",
		"the deployment platform; supported plaforms: cf, xsa")
}

// buildModuleCmd - Build module
var buildModuleCmd = &cobra.Command{
	Use:   "build",
	Short: "builds module",
	Long:  "builds module and archives its artifacts",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := artifacts.ExecuteBuild(buildCmdSrc, buildCmdTrg, buildCmdDesc, buildCmdModule, buildCmdPlatform, os.Getwd)
		logError(err)
		return err
	},
	SilenceUsage:  true,
	SilenceErrors: false,
}

// zip specific module and put the artifacts on the temp folder according
// to the mtar structure, i.e each module has new entry as folder in the mtar folder
// Note - even if the path of the module was changed in the mta.yaml in the mtar the
// the module folder will get the module name
var packModuleCmd = &cobra.Command{
	Use:   "pack",
	Short: "packs module artifacts",
	Long:  "packs the module artifacts after the build process",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := artifacts.ExecutePack(packCmdSrc, packCmdTrg, packCmdDesc, packCmdModule, packCmdPlatform, os.Getwd)
		logError(err)
		return err
	},
}
