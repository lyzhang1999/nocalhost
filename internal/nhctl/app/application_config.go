package app

import (
	"strconv"
	"time"
)

const (
	DefaultSideCarImage                      = "codingcorp-docker.pkg.coding.net/nocalhost/public/nocalhost-sidecar:syncthing"
	DefaultDevImage                          = "codingcorp-docker.pkg.coding.net/nocalhost/public/minideb:latest"
	DefaultWorkDir                           = "/home/nocalhost-dev"
	DefaultLocalSyncDirName                  = "."
	DefaultPortForwardDir                    = "port-forward"
	DefaultResourcesDir                      = "resources"
	DefaultForwardRemoteSshPort              = 22
	DefaultForwardLocalSshPort               = 30002
	DefaultNhctlHomeDirName                  = ".nh/nhctl"
	DefaultSshKeyDirName                     = "key"
	DefaultBinDirName                        = "bin"
	DefaultLogDirName                        = "logs"
	DefaultSyncLogFileName                   = "sync-port-forward-child-process.log"
	DefaultApplicationSyncPortForwardPidFile = "sync-port-forward.pid"
	DefaultBinSyncThingDirName               = "syncthing"
	DefaultBackGroundPortForwardLogFileName  = "alone-port-forward-child-process.log"
	DefaultApplicationOnlyPortForwardPidFile = "alone-port-forward.pid"
	DefaultApplicationSyncPidFile            = "syncthing.pid"
	DefaultApplicationDirName                = "application"
	DefaultApplicationProfilePath            = ".profile.yaml"
	DefaultApplicationConfigDirName          = ".nocalhost"
	DefaultApplicationConfigName             = "config.yaml"
	DefaultNewFilePermission                 = 0700
	DefaultClientGoTimeOut                   = time.Minute * 5
)

type NocalHostAppConfig struct {
	PreInstall []*PreInstallItem    `json:"pre_install" yaml:"preInstalls"`
	SvcConfigs []*ServiceDevOptions `json:"svc_config" yaml:"svcConfigs"`
	AppConfig  *AppConfig           `json:"app_config" yaml:"appConfig"`
}

type PreInstallItem struct {
	Path   string `json:"path" yaml:"path"`
	Weight string `json:"weight" yaml:"weight"`
}

type AppConfig struct {
	Name         string  `json:"name" yaml:"name"`
	Type         AppType `json:"type" yaml:"type"`
	ResourcePath string  `json:"resource_path" yaml:"resourcePath"`
}

type ServiceDevOptions struct {
	Name         string                `json:"name" yaml:"name"`
	Type         string                `json:"type" yaml:"type"`
	GitUrl       string                `json:"git_url" yaml:"gitUrl"`
	DevLang      string                `json:"dev_env" yaml:"devLang"` // java|go|node
	DevImage     string                `json:"dev_image" yaml:"devImage"`
	SideCarImage string                `json:"side_car_image" yaml:"sideCarImage"`
	WorkDir      string                `json:"work_dir" yaml:"workDir"`
	LocalWorkDir string                `json:"local_work_dir" yaml:""`
	Sync         []string              `json:"sync" yaml:"sync"`
	Ignore       []string              `json:"ignore" yaml:"ignore"`
	SshPort      *SshPortForwardConfig `json:"ssh_port" yaml:"sshPort"`
	DevPort      []string              `json:"dev_port" yaml:"devPort"`
	Command      []string              `json:"command" yaml:"command"`
	Jobs         []string              `json:"jobs" yaml:"jobs,omitempty"`
	Pods         []string              `json:"pods" yaml:"pods,omitempty"`
}

type SshPortForwardConfig struct {
	LocalPort int `json:"local_port" yaml:"localPort"`
	SshPort   int `json:"ssh_port" yaml:"sshPort"`
}

type ComparableItems []*PreInstallItem

func (a ComparableItems) Len() int      { return len(a) }
func (a ComparableItems) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ComparableItems) Less(i, j int) bool {
	iW, err := strconv.Atoi(a[i].Weight)
	if err != nil {
		iW = 0
	}

	jW, err := strconv.Atoi(a[j].Weight)
	if err != nil {
		jW = 0
	}
	return iW < jW
}

func (n *NocalHostAppConfig) GetSvcConfig(name string) *ServiceDevOptions {
	if n.SvcConfigs == nil {
		return nil
	}
	for _, svc := range n.SvcConfigs {
		if svc.Name == name {
			return svc
		}
	}
	return nil
}
