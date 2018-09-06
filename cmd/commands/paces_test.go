package commands

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"cloud-mta-build-tool/cmd/fsys"
	"cloud-mta-build-tool/cmd/logs"
	"github.com/stretchr/testify/assert"
)

func Test_copyModuleAndClean(t *testing.T) {
	srcFilePath := filepath.Join("testdata", "mtahtml5", "testapp", "webapp", "controller")
	dstFilePath := filepath.Join("testdata", "result")
	copyModule.Run(nil, []string{srcFilePath, dstFilePath, "controller"})
	resultPath := filepath.Join(dir.GetPath(), "testdata", "result", "testdata", "mtahtml5", "testapp", "webapp", "controller", "View1.controller.js")
	fileInfo, _ := os.Stat(resultPath)
	assert.Equal(t, fileInfo.IsDir(), false)

	cleanup.Run(nil, []string{filepath.Join("testdata", "result")})
	fileInfo, _ = os.Stat(resultPath)
	assert.Nil(t, fileInfo)

}

func Test_genMetaFunction(t *testing.T) {
	args := []string{filepath.Join("testdata", "result"), "testapp"}
	generateMeta(filepath.Join("testdata", "mtahtml5"), args)
	actualContent, _ := ioutil.ReadFile(filepath.Join(dir.GetPath(), "testdata", "result", "META-INF", "mtad.yaml"))
	expectedContent, _ := ioutil.ReadFile(filepath.Join(dir.GetPath(), "testdata", "golden", "mtad.yaml"))
	assert.Equal(t, actualContent, expectedContent)
	os.RemoveAll(filepath.Join(dir.GetPath(), "testdata", "result"))
}

func Test_genMetaCommand(t *testing.T) {
	args := []string{filepath.Join("testdata", "result"), "testapp"}
	genMeta.Run(nil, args)
	actualContent, _ := ioutil.ReadFile(filepath.Join(dir.GetPath(), "testdata", "result", "META-INF", "mtad.yaml"))
	assert.Nil(t, actualContent)
}

func Test_genMtarFunction(t *testing.T) {
	args := []string{filepath.Join(dir.GetPath(), "testdata", "mtahtml5"), filepath.Join(dir.GetPath(), "testdata")}
	generateMtar(filepath.Join("testdata", "mtahtml5"), args)
	actualContent, _ := ioutil.ReadFile(filepath.Join(dir.GetPath(), "testdata", "mtahtml5.mtar"))
	expectedContent, _ := ioutil.ReadFile(filepath.Join(dir.GetPath(), "testdata", "golden", "mtahtml5.mtar"))
	assert.Equal(t, actualContent, expectedContent)
	os.RemoveAll(filepath.Join(dir.GetPath(), "testdata", "mtahtml5.mtar"))
}

func Test_genMtarCommand(t *testing.T) {
	args := []string{filepath.Join(dir.GetPath(), "testdata", "mtahtml5"), filepath.Join(dir.GetPath(), "testdata")}
	genMtar.Run(nil, args)
	actualContent, _ := ioutil.ReadFile(filepath.Join(dir.GetPath(), "testdata", "mtahtml5.mtar"))
	assert.Nil(t, actualContent)
}

func Test_pack(t *testing.T) {
	tests := []struct {
		name      string
		args      []string
		validator func(t *testing.T, args []string)
	}{
		{
			name: "SanityTest",
			args: []string{filepath.Join(dir.GetPath(), "testdata", "result"),
				filepath.Join("testdata", "mtahtml5", "testapp"),
				"ui5app"},
			validator: func(t *testing.T, args []string) {
				resultPath := filepath.Join(args[0], "ui5app", "data.zip")
				fileInfo, _ := os.Stat(resultPath)
				assert.NotNil(t, fileInfo)
				assert.Equal(t, fileInfo.IsDir(), false)
				os.RemoveAll(args[0])
			},
		},
		{
			name: "Wrong relative path to module",
			args: []string{filepath.Join(dir.GetPath(), "testdata", "result"),
				filepath.Join("testdata", "mtahtml5", "ui5app"),
				"ui5app"},
			validator: func(t *testing.T, args []string) {
				resultPath := filepath.Join(args[0], "ui5app", "data.zip")
				fileInfo, _ := os.Stat(resultPath)
				assert.Nil(t, fileInfo)
			},
		},
		{
			name: "Missing arguments",
			args: []string{filepath.Join(dir.GetPath(), "testdata", "result"),
				"ui5app"},
			validator: func(t *testing.T, args []string) {
				resultPath := filepath.Join(args[0], "ui5app", "data.zip")
				fileInfo, _ := os.Stat(resultPath)
				assert.Nil(t, fileInfo)
			},
		},
	}
	logs.NewLogger()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pack.Run(nil, tt.args)
			tt.validator(t, tt.args)

		})
	}
}
