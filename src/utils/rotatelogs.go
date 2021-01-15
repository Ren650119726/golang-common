package utils

import (
	zaprotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap/zapcore"
	"lottery-demo/src/core/global"
	"os"
	"path"
	"runtime"
	"time"
)

//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: GetWriteSyncer
//@description: zap logger中加入file-rotatelogs
//@return: zapcore.WriteSyncer, error

func GetWriteSyncer(fileName string) (zapcore.WriteSyncer, error) {
	var fileWriter *zaprotatelogs.RotateLogs
	var err error
	if runtime.GOOS != "windows" {
		fileWriter, err = zaprotatelogs.New(
			path.Join(global.GVA_CONFIG.Zap.Director, fileName+"%Y-%m-%d.log"),
			zaprotatelogs.WithLinkName(global.GVA_CONFIG.Zap.LinkName),
			zaprotatelogs.WithMaxAge(7*24*time.Hour),
			zaprotatelogs.WithRotationTime(24*time.Hour),
		)
	} else {
		fileWriter, err = zaprotatelogs.New(
			path.Join(global.GVA_CONFIG.Zap.Director, fileName+"%Y-%m-%d.log"),
			zaprotatelogs.WithMaxAge(7*24*time.Hour),
			zaprotatelogs.WithRotationTime(24*time.Hour),
		)
	}
	if global.GVA_CONFIG.Zap.LogInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter)), err
	}
	return zapcore.AddSync(fileWriter), err

}
