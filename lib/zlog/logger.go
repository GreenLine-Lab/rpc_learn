package zlog

import (
	"fmt"
	"github.com/rs/zerolog"
	"os"
	"strings"
	"time"
)

const (
	_ = iota + 30
	colorRed
	colorGreen
	colorYellow
	_
	colorMagenta

	colorBold = 1
	_         = 90
)

type ZLogger struct {
	NoColor    bool
	JSONFormat bool
}

func (zlog *ZLogger) Init() zerolog.Logger {

	if zlog.JSONFormat {
		return zerolog.New(zlog.consoleWriterJSON())
	}

	return zerolog.New(zlog.consoleWriter())
}

func (zlog *ZLogger) consoleWriterJSON() zerolog.ConsoleWriter {
	return zerolog.ConsoleWriter{
		Out:        os.Stdout,
		NoColor:    zlog.NoColor,
		TimeFormat: time.StampNano,
	}
}

func (zlog *ZLogger) consoleWriter() zerolog.ConsoleWriter {
	cWriter := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		NoColor:    zlog.NoColor,
		TimeFormat: time.StampNano,
	}

	cWriter.FormatLevel = func(i interface{}) string {
		var l string
		if ll, ok := i.(string); ok {
			switch ll {
			case "trace":
				l = colorize("TRC", colorMagenta, zlog.NoColor)
			case "debug":
				l = colorize("DBG", colorYellow, zlog.NoColor)
			case "info":
				l = colorize("INF", colorGreen, zlog.NoColor)
			case "warn":
				l = colorize("WRN", colorRed, zlog.NoColor)
			case "error":
				l = colorize(colorize("ERR", colorRed, zlog.NoColor), colorBold, zlog.NoColor)
			case "fatal":
				l = colorize(colorize("FTL", colorRed, zlog.NoColor), colorBold, zlog.NoColor)
			case "panic":
				l = colorize(colorize("PNC", colorRed, zlog.NoColor), colorBold, zlog.NoColor)
			default:
				l = colorize("???", colorBold, zlog.NoColor)
			}
		} else {
			if i == nil {
				l = colorize("???", colorBold, zlog.NoColor)
			} else {
				l = strings.ToUpper(fmt.Sprintf("| %s |", i))[0:3]
			}
		}
		return fmt.Sprintf("| %s |", l)
	}

	return cWriter
}

func colorize(s interface{}, c int, disabled bool) string {
	if disabled {
		return fmt.Sprintf("%s", s)
	}
	return fmt.Sprintf("\x1b[%dm%v\x1b[0m", c, s)
}
