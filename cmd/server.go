package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	ccmd "ntsc.ac.cn/ta-registry/pkg/cmd"
	"ntsc.ac.cn/ta-snmp-agent/internal/server"
)

var serverEnvs struct {
	enableLXD bool
}
var serverCmd = &cobra.Command{
	Use:    "server",
	Short:  "general tas server agent",
	PreRun: _server_prerun,
	Run:    _server_run,
}

func init() {
	rootCmd.AddCommand(serverCmd)
	serverCmd.Flags().BoolVar(&serverEnvs.enableLXD, "enable-lxd", false,
		"enable lxd service monitor")
}

func _server_prerun(cmd *cobra.Command, args []string) {
	ccmd.InitGlobalVars()
	_initTASClient()
	go func() {
		ccmd.RunWithSysSignal(nil)
	}()
}
func _server_run(cmd *cobra.Command, args []string) {
	serv, err := server.NewGeneralServer(tasClient)
	if err != nil {
		logrus.WithField("prefix", "server").
			Fatalf("create general server monitor service failed: %v", err)
	}
	logrus.WithField("prefix", "server").
		Fatalf("run general server monitor service failed: %v", <-serv.Start())
}
