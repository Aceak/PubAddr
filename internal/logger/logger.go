package logger

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

type Level int

const (
	DEBUG Level = iota
	INFO
	WARN
	ERROR
	FATAL
)

var levelNames = map[Level]string{
	DEBUG: "DEBUG",
	INFO:  "INFO",
	WARN:  "WARN",
	ERROR: "ERROR",
	FATAL: "FATAL",
}

type Logger struct {
	level Level
	mu    sync.Mutex
}

var (
	instance *Logger
	once     sync.Once
)

func InitLogger(levelStr string) {
	once.Do(func() {
		instance = &Logger{
			level: ParseLevel(levelStr),
		}
	})
}

func ParseLevel(s string) Level {
	switch strings.ToLower(strings.TrimSpace(s)) {
	case "debug":
		return DEBUG
	case "info":
		return INFO
	case "warn", "warning":
		return WARN
	case "error":
		return ERROR
	case "fatal":
		return FATAL
	default:
		return INFO
	}
}

// SetLevel 运行时修改日志等级
func SetLevel(levelStr string) {
	level := ParseLevel(levelStr)
	if instance != nil {
		instance.mu.Lock()
		instance.level = level
		instance.mu.Unlock()
	}
}

// 核心输出
func (l *Logger) log(level Level, msg string, args ...interface{}) {
	if level < l.level {
		return
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	now := time.Now()
	tzName, _ := now.Zone()
	timestamp := now.Format("2006-01-02 15:04:05") + " " + tzName

	content := msg
	if len(args) > 0 {
		content = fmt.Sprintf(msg, args...)
	}

	fmt.Fprintf(os.Stdout, "[%s] [%s] %s\n",
		timestamp,
		levelNames[level],
		strings.TrimSpace(content),
	)

	if level == FATAL {
		os.Exit(1)
	}
}

func Debug(msg string, args ...interface{}) {
	if instance != nil {
		instance.log(DEBUG, msg, args...)
	}
}

func Info(msg string, args ...interface{}) {
	if instance != nil {
		instance.log(INFO, msg, args...)
	}
}

func Warn(msg string, args ...interface{}) {
	if instance != nil {
		instance.log(WARN, msg, args...)
	}
}

func Error(msg string, args ...interface{}) {
	if instance != nil {
		instance.log(ERROR, msg, args...)
	}
}

func Fatal(msg string, args ...interface{}) {
	if instance != nil {
		instance.log(FATAL, msg, args...)
	}
}
