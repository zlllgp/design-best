package slog

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"
	"runtime"
	"strconv"
	"time"
)

// ShowLog https://blog.csdn.net/walkskyer/article/details/138300327
func ShowLog() {
	slog.Info("a", "name", "test")
	slog.Log(context.Background(), slog.LevelWarn, "this", "name", "test")
}

// 初始化和配置
func ShowLogWithFile() {
	// 创建一个日志文件
	file, err := os.OpenFile("../../slog/app.slog", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// 设置日志输出到文件
	log.SetOutput(file)

	// 设置日志前缀
	log.SetPrefix("INFO: ")
	// 设置日志的格式
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)

	log.Println("This is a test slog entry")
}

const (
	LevelError = iota
	LevelWarning
	LevelInfo
)

// 日志级别
func ShowLogWithFileWithLevel() {
	file, err := os.OpenFile("../../slog/app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	logger := log.New(file, "PREFIX: ", log.Ldate|log.Ltime|log.Llongfile)

	// 使用不同的日志级别记录信息
	logError(logger, "This is an error message")
	logWarning(logger, "This is a warning message")
	logInfo(logger, "This is an info message")
}
func logError(logger *log.Logger, msg string) {
	logger.SetPrefix("ERROR: ")
	logger.Println(msg)
}

func logWarning(logger *log.Logger, msg string) {
	logger.SetPrefix("WARNING: ")
	logger.Println(msg)
}

func logInfo(logger *log.Logger, msg string) {
	logger.SetPrefix("INFO: ")
	logger.Println(msg)
}

type customLogger struct {
	logger *log.Logger
}

func newCustomLogger(out *os.File) *customLogger {
	return &customLogger{
		logger: log.New(out, "", 0),
	}
}

func (c *customLogger) Printf(format string, v ...interface{}) {
	c.logger.SetPrefix(fmt.Sprintf("CUSTOM LOG [%s]: ", time.Now().Format(time.RFC3339)))
	c.logger.Printf(format, v...)
}

// 自定义日志格式器
func ShowLogWithCustom() {
	file, err := os.OpenFile("../../slog/custom_app.slog", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	cLogger := newCustomLogger(file)
	cLogger.Printf("This is a custom slog message with dynamic timestamp.")
}

// 条件日志处理
func ShowCondition() {
	file, err := os.OpenFile("../../slog/condition_app.slog", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	logger := log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

	if runtime.GOOS == "linux" {
		logger.Println("This slog is only written on Linux systems.")
	} else {
		logger.Println("This slog is not written on Linux systems.")
	}
}

// 日志消息队列
var logQueue chan string

func init() {
	logQueue = make(chan string, 1000) // 创建一个有缓冲的通道
	go func() {
		file, err := os.OpenFile("../../log/async_app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		logger := log.New(file, "ASYNC: ", log.Ldate|log.Ltime|log.Lshortfile)

		for msg := range logQueue {
			logger.Println(msg)
		}
	}()
}

func logMessage(message string) {
	logQueue <- message // 将消息发送到队列
}

func AsyncLog() {
	for i := 0; i < 100; i++ {
		logMessage("Log entry number: " + strconv.Itoa(i))
	}
	time.Sleep(1 * time.Second)
}

func someFunctionThatMightFail() (int, error) {
	return 0, fmt.Errorf("something went wrong")
}

func LogError() {
	file, err := os.OpenFile("../../log/handling.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	logger := log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Llongfile)

	// 模拟一个错误情况
	_, err = someFunctionThatMightFail()
	if err != nil {
		logger.Printf("Error occurred: %v", err)
	}
}

func LogJson() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Debug("Debug message")
	logger.Info("Info message")
	logger.Warn("Warning message")
	logger.Error("Error message")

}
