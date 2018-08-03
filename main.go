package main

import (
	"odin_tool_v3/cmd"
	"os"

	"github.com/urfave/cli"
)

const APP_VAR = "1.0.0"

func main() {
	app := cli.NewApp()
	app.Name = "odin_tool"
	app.Usage = "12odin后台工具集"
	app.Version = APP_VAR
	app.Commands = []cli.Command{
		cmd.Web,
	}
	app.Run(os.Args)
}
