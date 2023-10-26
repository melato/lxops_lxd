package main

import (
	_ "embed"
	"fmt"

	"melato.org/command"
	"melato.org/command/usage"
	"melato.org/lxops"
	"melato.org/lxops_lxd/lxdutil"
)

//go:embed usage.yaml
var usageData []byte

//go:embed version
var version string

func main() {
	lxops.InitOSTypes()
	lxops.InitConfigTypes()
	client := &lxdutil.LxdClient{}
	cmd := lxops.RootCommand(client)
	templateOps := &lxdutil.TemplateOps{Client: client}
	cmd.Command("template").Flags(templateOps).RunFunc(templateOps.Apply)
	cmd.Command("version").NoConfig().RunMethod(func() { fmt.Println(version) })
	usage.Apply(cmd, usageData)
	command.Main(cmd)
}
