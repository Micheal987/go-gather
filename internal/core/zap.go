package core

import (
	"fmt"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func Zap(url string) {
	InitLog("internal/logs/global", "Global")
}

// InitLog 初始化日志
func InitLog(logPath string, appName string) {
	//時間分割
	fileDate := time.Now().Format("2006-01-02")
	//创建目录
	err := os.MkdirAll(fmt.Sprintf("%s/%s", logPath, fileDate), os.ModePerm)
	if err != nil {
		zap.Error(err)
		return
	}
	//文件的位置
	filename := fmt.Sprintf("%s/%s/%s.log", logPath, fileDate, appName)
	//lumberjack的調用
	createLogs(filename, 500, 3, 28, true, true)
}

// 日志写入文件
// internal/core/zap/zap.log
func createLogs(Filename string, MaxSize, MaxBackups, MaxAge int, Compress, LocalTime bool) {
	//
	w := &lumberjack.Logger{
		Filename:   Filename,   //log文件的位置
		MaxSize:    MaxSize,    // M 为单位，达到这个设置数后就进行日志切割
		MaxBackups: MaxBackups, //保留旧文件最大份数
		MaxAge:     MaxAge,     // days 旧文件最大保存天数
		Compress:   Compress,   //是否压缩日志归档，默认不压缩
		LocalTime:  LocalTime,  // 時間格式化
	}
	//Close
	defer w.Close()
	//NewCore配置
	//
	config := zap.NewDevelopmentEncoderConfig()
	//時間
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	//
	fileEncoder := zapcore.NewJSONEncoder(config)
	//
	core := zapcore.NewCore(
		fileEncoder,        //编码设置
		zapcore.AddSync(w), //输出到文件
		zap.InfoLevel,      //日志等级
	)
	//logger 實例化
	//強類型
	//最快的
	logger := zap.New(core)
	defer logger.Sync()

	//Info等級
	logger.Info(
		"failed to fetch URL",
		// Structured context as strongly typed Field values.
		zap.String("url", "sssurl"),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
	logger.Info("測試信息")
	logger.With(
		zap.String("url", "prot"),
		zap.String("name", "jimmmyr"),
		zap.Int("age", 23),
		zap.String("agradege", "no111-000222"),
	).Info("test info")
}
