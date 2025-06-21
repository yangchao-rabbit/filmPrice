package logger

import (
	"io"
	"log"
	"os"
)

type Option func(*option)

type option struct {
	to     io.Writer
	prefix string
	flag   int
}

func New(opts ...Option) *log.Logger {
	opt := new(option)

	for _, f := range opts {
		f(opt)
	}

	return log.New(opt.to, opt.prefix, opt.flag)
}

func WithDefault() Option {
	return func(opt *option) {
		opt.to = os.Stdout
		opt.prefix = ""
		opt.flag = log.Lshortfile | log.LstdFlags
	}
}

// WithFile 设置日志存储文件
func WithFile(file string) Option {
	if file == "stdout" {
		return func(opt *option) {
			opt.to = os.Stdout
		}
	} else {
		logFile, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			panic(err)
		}
		return func(opt *option) {
			opt.to = logFile
		}
	}
}

// WithFlag 设置日志的 flag
// f int:
//
// 例:
//
//	log.Lshortfile | log.LstdFlags
//	log.Llongfile | log.LstdFlags
func WithFlag(f int) Option {
	return func(opt *option) {
		opt.flag = f
	}
}

// WithPrefix 设置日志前缀
func WithPrefix(s string) Option {
	return func(opt *option) {
		opt.prefix = s
	}
}
