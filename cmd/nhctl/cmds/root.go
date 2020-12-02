package cmds

import (
	"fmt"
	"nocalhost/internal/nhctl/app"
	nocalhost2 "nocalhost/internal/nhctl/nocalhost"
	"nocalhost/pkg/nhctl/log"
	"nocalhost/pkg/nhctl/utils"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var settings *EnvSettings
var nh *nocalhost2.NocalHost
var nocalhostApp *app.Application

func init() {

	settings = NewEnvSettings()

	rootCmd.PersistentFlags().BoolVar(&settings.Debug, "debug", settings.Debug, "enable debug level log")
	rootCmd.PersistentFlags().StringVar(&settings.KubeConfig, "kubeconfig", "", "the path to the kubeconfig file")

	cobra.OnInitialize(func() {
		var err error
		nh, err = nocalhost2.NewNocalHost()
		utils.Mush(err)
	})
}

var rootCmd = &cobra.Command{
	Use:   "nhctl",
	Short: "nhctl use to deploy coding project",
	Long:  `nhctl can deploy project on Kubernetes. `,
	Run: func(cmd *cobra.Command, args []string) {
		log.Debug("hello nhctl")
		//fmt.Printf("kubeconfig is %s", settings.KubeConfig)
	},
}

func Execute() {
	//if settings.Debug {
	//	log.SetLevel(logrus.DebugLevel)
	//}
	if len(os.Args) == 1 {
		args := append([]string{"help"}, os.Args[1:]...)
		rootCmd.SetArgs(args)
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func InitAppAndSvc(appName string, svcName string) {
	var err error
	if settings.Debug {
		log.SetLevel(logrus.DebugLevel)
	}
	if !nh.CheckIfApplicationExist(appName) {
		log.Fatalf("application \"%s\" not found", appName)
	}
	nocalhostApp, err = app.NewApplication(appName)
	if err != nil {
		log.Fatal("fail to get application info")
	}
	if svcName == "" {
		log.Fatal("please use -d to specify a k8s deployment")
	}

	exist, err := nocalhostApp.CheckIfSvcExist(svcName, app.Deployment)
	if err != nil {
		log.Fatalf("fail to check if svc exist : %v", err)
	} else if !exist {
		log.Fatalf("\"%s\" not found", svcName)
	}
}
