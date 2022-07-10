package log

import (
	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs" // 日志分割
	"github.com/rifflock/lfshook"                       //日志钩子
	"github.com/sirupsen/logrus"
	"haolinju.xianhaohan.com/common"
	"haolinju.xianhaohan.com/internal/conf"
	"time"
)

// Create a new instance of the logger. You can have any number of instances.
var (
	logger = logrus.New()
)

type Fields logrus.Fields

func Init() {
	// 日志存放路径
	logPath := conf.Conf.App.Log.Path
	// 初始化日志文件对象
	//_, err = os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	//if err != nil {
	//	return nil, err
	//}

	// 设置日志级别
	logger.SetLevel(logrus.DebugLevel)

	// 日志分隔：1. 每天产生的日志写在不同的文件；2. 只保留一定时间的日志（例如：一星期）
	logWriter, _ := rotatelogs.New(
		logPath+"_%Y%m%d.log", // 日志文件名格式
		rotatelogs.WithMaxAge(time.Duration(conf.Conf.App.Log.MaxAge)*time.Hour),             // 最多保留7天之内的日志
		rotatelogs.WithRotationTime(time.Duration(conf.Conf.App.Log.RotationTime)*time.Hour), // 一天保存一个日志文件
		rotatelogs.WithLinkName(logPath),                                                     // 为最新日志建立软连接
		//rotatelogs.WithRotationSize(64),           // 根据文件大小切割
	)
	writeMap := lfshook.WriterMap{
		logrus.DebugLevel: logWriter,
		logrus.InfoLevel:  logWriter, // info级别使用logWriter写日志
		logrus.ErrorLevel: logWriter,
		logrus.FatalLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}

	// 设置当前应用环境变量
	gin.SetMode(conf.Conf.App.Mode)
	if gin.Mode() == gin.ReleaseMode {
		// 关闭日志的颜色输出
		//gin.DisableConsoleColor()
		// 把产生的日志内容写进日志文件中
		//logger.Out = file
	} else {
		//同时写文件和屏幕
		//同时写文件和屏幕
		//同时写文件和屏幕
		//writers := []io.Writer{file}
		//fileAndStdoutWriter := io.MultiWriter(writers...)
		//logger.SetOutput(file)
		//logger.Out = os.Stdout
		// 设置显示文件、行号
		//logger.SetReportCaller(true)
	}

	Hook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	logger.AddHook(Hook)

	return
}

func Debug(ctx *gin.Context, message interface{}, fields Fields) {
	log(logrus.DebugLevel, ctx, message, fields)
}

func Info(ctx *gin.Context, message interface{}, fields Fields) {
	log(logrus.InfoLevel, ctx, message, fields)
}

func Warn(ctx *gin.Context, message interface{}, fields Fields) {
	log(logrus.WarnLevel, ctx, message, fields)
}

func Error(ctx *gin.Context, message interface{}, fields Fields) {
	log(logrus.ErrorLevel, ctx, message, fields)
}

//  log 之后会调用 os.Exit(1)
func Fatal(ctx *gin.Context, message interface{}, fields Fields) {
	log(logrus.FatalLevel, ctx, message, fields)
}

// log 之后会 panic()
func Panic(ctx *gin.Context, message interface{}, fields Fields) {
	log(logrus.PanicLevel, ctx, message, fields)
}

func log(level logrus.Level, ctx *gin.Context, message interface{}, f Fields) {
	fields := logrus.Fields(f)
	if fields == nil {
		fields = logrus.Fields{}
	}
	if ctx == nil {
		ctx = &gin.Context{}
	} else {
		fields["clientIp"] = ctx.ClientIP()
	}

	// 添加链路追踪requestId
	fields["requestId"] = common.Trace(ctx)

	logger.WithFields(fields).Log(level, message)
}
