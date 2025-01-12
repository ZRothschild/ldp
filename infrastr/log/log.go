package log

import (
	"context"
	"errors"
	"fmt"
	"github.com/ZRothschild/ldp/infrastr/lib/tool"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
	"log/slog"
	"os"
	"time"
)

var (
	SLog = slog.New(slog.NewJSONHandler(os.Stdout, nil))
)

func NewDefault(l *slog.Logger, config logger.Config) *OrmLog {
	var (
		infoStr      = "%s\n[info] "
		warnStr      = "%s\n[warn] "
		errStr       = "%s\n[error] "
		traceStr     = "%s\n[%.3fms] [rows:%v] %s"
		traceWarnStr = "%s %s\n[%.3fms] [rows:%v] %s"
		traceErrStr  = "%s %s\n[%.3fms] [rows:%v] %s"
	)
	if config.Colorful {
		infoStr = logger.Green + "%s\n" + logger.Reset + logger.Green + "[info] " + logger.Reset
		warnStr = logger.BlueBold + "%s\n" + logger.Reset + logger.Magenta + "[warn] " + logger.Reset
		errStr = logger.Magenta + "%s\n" + logger.Reset + logger.Red + "[error] " + logger.Reset
		traceStr = logger.Green + "%s\n%s" + logger.Reset + logger.Yellow + "[%.3fms] " + logger.BlueBold + "[rows:%v]" + logger.Reset + " %s"
		traceWarnStr = logger.Green + "%s " + logger.Yellow + "%s\n" + logger.Reset + logger.RedBold + "[%.3fms] " + logger.Yellow + "[rows:%v]" + logger.Magenta + " %s" + logger.Reset
		traceErrStr = logger.RedBold + "%s " + logger.MagentaBold + "%s\n" + logger.Reset + logger.Yellow + "[%.3fms] " + logger.BlueBold + "[rows:%v]" + logger.Reset + " %s"
	}
	return &OrmLog{
		Logger:       l,
		Config:       config,
		infoStr:      infoStr,
		warnStr:      warnStr,
		errStr:       errStr,
		traceStr:     traceStr,
		traceWarnStr: traceWarnStr,
		traceErrStr:  traceErrStr,
	}

}

type (
	OrmLog struct {
		*slog.Logger
		Config       logger.Config
		infoStr      string
		warnStr      string
		errStr       string
		traceStr     string
		traceWarnStr string
		traceErrStr  string
	}
)

func (l *OrmLog) LogMode(level logger.LogLevel) logger.Interface {
	newL := *l
	newL.Config.LogLevel = level
	return &newL
}

func (l *OrmLog) Info(ctx context.Context, msg string, arg ...interface{}) {
	l.Logger.InfoContext(ctx, fmt.Sprintf(msg, arg...))
}

func (l *OrmLog) Warn(ctx context.Context, msg string, arg ...interface{}) {
	l.Logger.WarnContext(ctx, fmt.Sprintf(msg, arg...))
}

func (l *OrmLog) Error(ctx context.Context, msg string, arg ...interface{}) {
	l.Logger.ErrorContext(ctx, fmt.Sprintf(msg, arg...))
}

func (l *OrmLog) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	if l.Config.LogLevel <= logger.Silent {
		return
	}
	var (
		s = time.Since(begin)
		f = func(str string, i interface{}, level logger.LogLevel) {
			sql, rows := fc()
			ok := tool.Include(level, []logger.LogLevel{logger.Warn, logger.Error}) > -1
			if rows == -1 && ok {
				l.Info(ctx, str, utils.FileWithLineNum(), i, float64(s.Nanoseconds())/1e6, "-", sql)
			} else if ok {
				l.Info(ctx, str, utils.FileWithLineNum(), i, float64(s.Nanoseconds())/1e6, rows, sql)
			} else if rows == -1 {
				l.Info(ctx, str, utils.FileWithLineNum(), float64(s.Nanoseconds())/1e6, "-", sql)
			} else {
				l.Info(ctx, str, utils.FileWithLineNum(), float64(s.Nanoseconds())/1e6, rows, sql)
			}
		}
	)
	switch {
	case err != nil && l.Config.LogLevel >= logger.Error && (!errors.Is(err, logger.ErrRecordNotFound) || !l.Config.IgnoreRecordNotFoundError):
		f(l.traceErrStr, err, logger.Error)
	case s > l.Config.SlowThreshold && l.Config.SlowThreshold != 0 && l.Config.LogLevel >= logger.Warn:
		f(l.traceWarnStr, fmt.Sprintf("SLOW SQL >= %v", l.Config.SlowThreshold), logger.Error)
	case l.Config.LogLevel == logger.Info:
		f(l.traceStr, "", logger.Info)
	}
}
