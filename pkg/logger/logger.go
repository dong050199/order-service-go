package logger

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math"
	"net"
	"order-service/pkg/ginutils/constants"
	"order-service/pkg/ginutils/timeutils"
	"order-service/pkg/ginutils/tracking"
	"os"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

type StandardLogger struct {
	*logrus.Logger
}

var loggerSingleton *StandardLogger

func InitLogger() {
	var prettyLog bool

	if os.Getenv("ENVIRONMENT") == "local" {
		prettyLog = true
	}
	lg := logrus.New()
	lg.SetFormatter(&logrus.JSONFormatter{
		PrettyPrint: prettyLog,
	})
	lg.SetOutput(os.Stdout)

	// Set logger level base on environment
	// if config.ServerConfig().Env == constants.ProductionEnv {
	// 	lg.SetLevel(logrus.InfoLevel)
	// }
	lg.SetReportCaller(true)

	// hook, err := NewUDPSplunkHook(config.ServerConfig().LoggerSplunkURL, config.ServerConfig().LoggerSplunkLayout)
	// if err != nil {
	// 	lg.Fatalf(`can't not configure splunk logger hook for logger`)
	// 	fmt.Printf("Init error: %v", err)
	// 	panic(err)
	// } else {
	// 	lg.AddHook(hook)
	// }

	loggerSingleton = &StandardLogger{lg}
}

func NewLogger() *StandardLogger {
	if loggerSingleton == nil {
		log.Fatal("can't not configure splunk logger hook for logger")
	}

	return loggerSingleton
}

// Không check long line linter cho interface
// nolint:lll
type Logger interface {
	WithField(key string, value interface{}) *Entry
	WithFields(fields logrus.Fields) *Entry
	WithError(err error) *Entry
	WithErrorStr(errStr string) *Entry
	WithQueue(input interface{}) *Entry
	WithInput(input interface{}) *Entry
	WithOutput(output interface{}) *Entry
	WithKeyword(ctx context.Context, keyword string) *Entry
	WithResponseTime(responseTime float64) *Entry
	WithFullLogDefer(ctx context.Context, keyWord string, apiURL string, path string, method string, startTime time.Time, dataLog map[string]interface{}, timeConfig float64, isShowInfo bool, isSlowLog bool) *Entry
	CtxPrefix(ctx context.Context, prefix string) *Entry
}

type Entry struct {
	*logrus.Entry
}

func NewEntry(entry *logrus.Entry) *Entry {
	return &Entry{Entry: entry}
}

func (s *Entry) WithField(key string, value interface{}) *Entry {
	entry := s.Entry.WithField(key, value)
	return NewEntry(entry)
}

func (s *Entry) WithFields(fields logrus.Fields) *Entry {
	entry := s.Entry.WithFields(fields)
	return NewEntry(entry)
}

func (s *Entry) WithError(err error) *Entry {
	entry := s.Entry.WithError(err)
	return NewEntry(entry)
}

func (s *Entry) WithErrorStr(errStr string) *Entry {
	entry := s.Entry.WithError(errors.New(errStr))
	return NewEntry(entry)
}

func (s *Entry) WithContext(ctx context.Context) *Entry {
	entry := s.Entry.WithContext(ctx)
	return NewEntry(entry)
}

func (s *Entry) WithInput(input interface{}) *Entry {
	entry := s.Entry.WithField("input", input)
	return NewEntry(entry)
}

func (s *Entry) WithOutput(output interface{}) *Entry {
	entry := s.Entry.WithField("output", output)
	return NewEntry(entry)
}

func (s *Entry) WithResponseTime(responsetime float64) *Entry {
	resTime := math.Round(responsetime)
	fieldTime := "response_time (ms):"
	entry := s.Entry.WithField(fieldTime, resTime)
	return NewEntry(entry)
}

func (s *Entry) Withkeyword(keyword string) *Entry {
	entry := s.Entry.WithField("keyword", keyword)
	return NewEntry(entry)
}

func (s *Entry) WithURL(url string) *Entry {
	entry := s.Entry.WithField("url", url)
	return NewEntry(entry)
}

func (s *Entry) WithStatusCode(code int) *Entry {
	entry := s.Entry.WithField("status_code", code)
	return NewEntry(entry)
}

func (s *StandardLogger) WithFields(fields logrus.Fields) *Entry {
	entry := s.Logger.WithFields(fields)
	return NewEntry(entry)
}

func (s *StandardLogger) WithError(err error) *Entry {
	entry := s.Logger.WithError(err)
	return NewEntry(entry)
}

func (s *StandardLogger) WithErrorStr(errStr string) *Entry {
	entry := s.Logger.WithError(errors.New(errStr))
	return NewEntry(entry)
}

func (s *StandardLogger) WithField(key string, value interface{}) *Entry {
	entry := s.Logger.WithField(key, value)
	return NewEntry(entry)
}

func (s *StandardLogger) WithInput(input interface{}) *Entry {
	entry := s.Logger.WithField("input", input)
	return NewEntry(entry)
}

func (s *StandardLogger) WithKeyword(ctx context.Context, keyword string) *Entry {
	trackID := tracking.GetTrackIDFromContext(ctx)
	entry := s.Logger.WithFields(logrus.Fields{"keyword": keyword, constants.TrackIDHeader: trackID})
	return NewEntry(entry)
}

func (s *StandardLogger) WithResponseTime(responseTime float64) *Entry {
	fieldTime := "ResponseTime"
	entry := s.Logger.WithField(fieldTime, math.Round(responseTime))
	return NewEntry(entry)
}

func (s *StandardLogger) WithOutput(output interface{}) *Entry {
	entry := s.Logger.WithField("output", output)
	return NewEntry(entry)
}

func (s *StandardLogger) WithFullLogDefer(
	ctx context.Context,
	keyWord string,
	apiURL string,
	path string,
	method string,
	startTime time.Time,
	dataLog map[string]interface{},
	timeConfig float64,
	isShowInfo bool,
	isSlowLog bool,
) *Entry {
	allTime := timeutils.Since(startTime).Seconds()
	dataLog["allTime"] = allTime
	fullURL := ""
	if len(apiURL) > 0 {
		fullURL = fmt.Sprintf("%s/%s", strings.TrimRight(apiURL, "/"), path)
	}
	trackID := tracking.GetTrackIDFromContext(ctx)
	entry := s.WithField(constants.TrackIDHeader, trackID).WithOutput(dataLog).WithResponseTime(allTime)
	_, ok := dataLog[constants.ErrorKey]
	if ok { //nolint: gocritic
		keyWord = fmt.Sprintf("%s defer error [%s]", keyWord, method)
		entry.WithURL(fullURL).WithField("keyword", keyWord).
			WithErrorStr(fmt.Sprintf("%+v", dataLog[constants.ErrorKey])).Error()
	} else if timeConfig > 0 && allTime >= 1 {
		if isSlowLog && allTime >= timeConfig {
			keyWord = fmt.Sprintf("Slow response from %s: time= %fs [%s] url=%s", apiURL, allTime, method, fullURL)
			entry.WithField("keyword", keyWord).Warning()
		} else {
			keyWord = fmt.Sprintf("%s defer warn [%s]", keyWord, method)
			entry.WithURL(fullURL).WithField("keyword", keyWord).Warning()
		}
	} else if isShowInfo {
		keyWord = fmt.Sprintf("%s defer info [%s]", keyWord, method)
		entry.WithURL(fullURL).WithField("keyword", keyWord).Info()
	}

	return entry
}

func (s *StandardLogger) CtxPrefix(ctx context.Context, prf string) *Entry {
	return s.WithField("request-id", fmt.Sprint(ctx.Value(constants.RequestIDKey))).WithField("prefix", prf)
}

func (s *StandardLogger) WithStatusCode(code int) *Entry {
	entry := s.Logger.WithField("status_code", code)
	return NewEntry(entry)
}

type udpSplunkHook struct {
	Conn   net.Conn
	url    string
	layout string
}

func NewUDPSplunkHook(url string, layout string) (*udpSplunkHook, error) {
	conn, err := net.Dial("udp", url)
	if err != nil {
		fmt.Printf("Cannot connect to Logger Server: %s\n error: %v", url, err)
	}
	return &udpSplunkHook{Conn: conn, layout: layout, url: url}, err
}

func (hook *udpSplunkHook) Fire(entry *logrus.Entry) error {
	line, err := entry.String()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to read entry, %v", err)
		return err
	}
	class := ""
	threadid := ""
	if entry.HasCaller() && entry.Caller != nil {
		class = entry.Caller.Function
		threadid = fmt.Sprintf("%s#%d", entry.Caller.File, entry.Caller.Line)
	}
	hostName, _ := os.Hostname()
	timeSt := strings.ReplaceAll(
		timeutils.ConvertTimeToString(
			timeutils.NowInGMT07(),
			constants.TimeLayoutyyyyMMddHHmmSSms),
		".", ",",
	)
	keyContent := []struct{ key, value string }{{"date", timeSt},
		{"hostname", hostName},
		{"class", class},
		{"threadid", threadid},
		{"message", line},
		{"level", strings.ToUpper(entry.Level.String())}}
	stringLog := hook.layout
	for i := 0; i < len(keyContent); i++ {
		stringLog = strings.ReplaceAll(stringLog, "%"+keyContent[i].key+"%", keyContent[i].value)
	}
	_, err = fmt.Fprint(hook.Conn, stringLog)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		return err
	}

	return err
}

func (hook *udpSplunkHook) Levels() []logrus.Level {
	return logrus.AllLevels
}
