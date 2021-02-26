package tkinit

import (
	tkuitls "github.com/ikaiguang/go_srv_kit/utils"
	"github.com/pkg/errors"
	"os"
	"path/filepath"
)

// DeployEnvMode .
type DeployEnvMode string

// String .
func (env DeployEnvMode) String() string {
	return string(env)
}

// config
const (
	defaultFileMode os.FileMode = 0644

	// deploy env.
	DeployEnvDev  DeployEnvMode = "dev"
	DeployEnvPre  DeployEnvMode = "pre"
	DeployEnvProd DeployEnvMode = "prod"

	// 相对路径
	RelPathConfig      = "conf.d"
	RelPathStatic      = "static.d"
	RelPathRuntime     = "runtime"
	RelPathLogs        = "logs"
	RelPathAttachments = "attachments"
)

// path
var (
	appPath         string
	configPath      string
	runtimePath     string
	logPath         string
	staticPath      string
	attachmentsPath string
	version         string
	deployEnv       DeployEnvMode
)

// Config .
type Config struct {
	AppID     string
	Version   string
	DeployEnv string // dev/pre/prod
	//Region    string
	//Zone      string
	//Hostname  string
}

// Setup .
func Setup(appConf, section string) {
	var err error
	defer func() {
		if err != nil {
			tkuitls.ExitTracer(err)
		}
	}()

	// path
	err = InitPath()
	if err != nil {
		err = errors.WithMessage(err, "*** init project path fail : ")
		return
	}
}

// Close .
func Close() (err error) {
	return
}

// DeployEnv deploy env
func DeployEnv() DeployEnvMode {
	return deployEnv
}

// IsProd test
func IsProd() bool {
	switch deployEnv {
	case DeployEnvPre, DeployEnvProd:
		return true
	default:
		return false
	}
}

// ConfigPath config path
func ConfigPath() string {
	return configPath
}

// AppPath app path
func AppPath() string {
	return appPath
}

// Version .
func Version() string {
	return version
}

// RuntimePath runtime path
func RuntimePath() string {
	return runtimePath
}

// LogPath log path
func LogPath() string {
	return logPath
}

// StaticPath static path
func StaticPath() string {
	return staticPath
}

// AttachmentPath attachments path
func AttachmentPath() string {
	return attachmentsPath
}

// InitPath .
func InitPath() (err error) {
	// app
	appPath, err = os.Getwd()
	if err != nil {
		err = errors.WithStack(err)
		return
	}

	// 目录信息
	var dirInfo os.FileInfo

	// config
	configPath = filepath.Join(appPath, RelPathConfig)
	dirInfo, err = os.Stat(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			err = errors.Errorf("cannot find config path : ./%s;\n\terror : %s", RelPathConfig, err.Error())
		} else {
			err = errors.WithStack(err)
		}
		return
	}
	if !dirInfo.IsDir() {
		err = errors.Errorf("Not a directory : %s", configPath)
		return
	}

	// static
	staticPath = filepath.Join(appPath, RelPathStatic)
	dirInfo, err = os.Stat(staticPath)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.Mkdir(staticPath, defaultFileMode)
			if err != nil {
				err = errors.WithStack(err)
				return
			}
			dirInfo, err = os.Stat(staticPath)
			if err != nil {
				err = errors.WithStack(err)
				return
			}
		} else {
			err = errors.WithStack(err)
			return
		}
	}
	if !dirInfo.IsDir() {
		err = errors.Errorf("Not a directory : %s", staticPath)
		return
	}

	// runtime
	runtimePath = filepath.Join(appPath, RelPathRuntime)
	dirInfo, err = os.Stat(runtimePath)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.Mkdir(runtimePath, defaultFileMode)
			if err != nil {
				err = errors.WithStack(err)
				return
			}
			dirInfo, err = os.Stat(runtimePath)
			if err != nil {
				err = errors.WithStack(err)
				return
			}
		} else {
			err = errors.WithStack(err)
			return
		}
	}
	if !dirInfo.IsDir() {
		err = errors.Errorf("Not a directory : %s", runtimePath)
		return
	}

	// log
	logPath = filepath.Join(runtimePath, RelPathLogs)
	dirInfo, err = os.Stat(logPath)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.Mkdir(logPath, defaultFileMode)
			if err != nil {
				err = errors.WithStack(err)
				return
			}
			dirInfo, err = os.Stat(logPath)
			if err != nil {
				err = errors.WithStack(err)
				return
			}
		} else {
			err = errors.WithStack(err)
			return
		}
	}
	if !dirInfo.IsDir() {
		err = errors.Errorf("Not a directory : %s", logPath)
		return
	}

	// attachments
	attachmentsPath = filepath.Join(appPath, RelPathAttachments)
	dirInfo, err = os.Stat(attachmentsPath)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.Mkdir(attachmentsPath, defaultFileMode)
			if err != nil {
				err = errors.WithStack(err)
				return
			}
			dirInfo, err = os.Stat(attachmentsPath)
			if err != nil {
				err = errors.WithStack(err)
				return
			}
		} else {
			err = errors.WithStack(err)
			return
		}
	}
	if !dirInfo.IsDir() {
		err = errors.Errorf("Not a directory : %s", attachmentsPath)
		return
	}
	return
}
