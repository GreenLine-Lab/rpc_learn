package zlog

import (
	"fmt"
	"github.com/rs/zerolog"
	"os"
	"time"
)

type ZLogger struct {
	Colorize   bool
	JSONFormat bool
}

func (zlog *ZLogger) Init() zerolog.Logger {
	return zerolog.Logger{}
}

func (zlog *ZLogger) consoleWriterJSON() zerolog.ConsoleWriter {
	return zerolog.ConsoleWriter{
		Out:        os.Stdout,
		NoColor:    zlog.Colorize,
		TimeFormat: time.StampNano,
	}
}

func (zlog *ZLogger) consoleWriter() zerolog.ConsoleWriter {
	cWriter := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		NoColor:    zlog.Colorize,
		TimeFormat: time.StampNano,
	}

	cWriter.FormatLevel = func(i interface{}) string {
		zlevel, ok := i.(zerolog.Level)
		if !ok {
			return fmt.Sprintf("| %-6s|", i)
		}

		var level string
		switch zlevel.String() {
		case zerolog.PanicLevel.String():
			level = "PNC"
		case zerolog.FatalLevel.String():
			level = "FTL"
		case zerolog.ErrorLevel.String():
			level = "ERR"
		case zerolog.WarnLevel.String():
			level = "WRN"
		case zerolog.InfoLevel.String():
			level = "INF"
		case zerolog.DebugLevel.String():
			level = "DBG"
		case zerolog.TraceLevel.String():
			level = "TRC"
		default:
			level = "???"
		}

		return fmt.Sprintf("| %-6s|", level)
	}

	return cWriter
}
