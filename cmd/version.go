package cmd

import (
	"encoding/json"
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

var GitVersion = "" // GitVersion 是语义化的版本号.
var BuildDate = ""  // BuildDate 是构建时间

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of voacap",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(getInfo().ToJSON())
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

// info 包含了版本信息.
type info struct {
	GitVersion string `json:"gitVersion"`
	BuildDate  string `json:"buildDate"`
	GoVersion  string `json:"goVersion"`
	Compiler   string `json:"compiler"`
	Platform   string `json:"platform"`
}

// ToJSON 以 JSON 格式返回版本信息.
func (info info) ToJSON() string {
	s, _ := json.Marshal(info)

	return string(s)
}

func getInfo() info {
	return info{
		GitVersion: GitVersion,
		BuildDate:  BuildDate,
		GoVersion:  runtime.Version(),
		Compiler:   runtime.Compiler,
		Platform:   fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}
}
