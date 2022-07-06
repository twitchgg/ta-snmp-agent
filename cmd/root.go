package cmd

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"ntsc.ac.cn/st-pcie-a/card/pkg/snmp"
	ccmd "ntsc.ac.cn/ta-registry/pkg/cmd"
	"ntsc.ac.cn/ta-snmp-agent/internal/common"
)

var rootEnvs struct {
	registryEndpoint   string
	certPath           string
	registryServerName string
}

var rootCmd = &cobra.Command{
	Use:   "ta-snmp-agent",
	Short: "tas snmp agent",
}

func init() {
	_flags()
	rootCmd.PersistentFlags().StringVar(
		&ccmd.GlobalEnvs.LoggerLevel, "logger-level", "DEBUG",
		"logger level")
	rootCmd.PersistentFlags().StringVar(
		&rootEnvs.registryEndpoint, "registry-endpoint", "tcp://localhost:1358",
		"TAS registry endpoint")
	rootCmd.PersistentFlags().StringVar(
		&rootEnvs.certPath, "cert-path", "/etc/tas/certs",
		"TAS certificate storage path")
	rootCmd.PersistentFlags().StringVar(
		&rootEnvs.registryServerName, "registry-servername", "s1.restry.ta.ntsc.ac.cn",
		"TAS registry server name")
}

func _flags() {
	cobra.OnInitialize(func() {})
	viper.AutomaticEnv()
	viper.SetEnvPrefix("TA")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	snmp.NewTrapServer(nil)
}

var tasClient *common.TASClient

func _initTASClient() {
	var err error
	if tasClient, err = common.NewTASClient(&common.TASConfig{
		CertPath:        rootEnvs.certPath,
		ServerName:      rootEnvs.registryServerName,
		ManagerEndpoint: rootEnvs.registryEndpoint,
	}); err != nil {
		logrus.New().WithField("prefix", "cmd.root").
			Fatalf("create TAS client failed: %v", err)
	}
}
