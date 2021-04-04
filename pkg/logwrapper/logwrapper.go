package logwrapper

import (
	"github.com/tracer0tong/kafkalogrus"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

// Event stores messages to log later, from our standard interface
type Event struct {
	id      int
	message string
}

// StandardLogger enforces specific log message formats
type StandardLogger struct {
	*logrus.Logger
}

// NewLogger initializes the standard logger
func NewLogger() *StandardLogger {
	var baseLogger = logrus.New()

	var standardLogger = &StandardLogger{baseLogger}

	standardLogger.Formatter = &logrus.JSONFormatter{}

	return standardLogger
}

func LogFields(requestType, respondTime string) *log.Fields {
	return &log.Fields{
		"requestType": requestType,
		"respondTime": respondTime,
	}
}

func NewHook() (*kafkalogrus.KafkaLogrusHook, error) {
	return kafkalogrus.NewKafkaLogrusHook(
		"kh",
		logrus.AllLevels,
		&logrus.JSONFormatter{},
		/*[]string{"127.0.0.1:9092"},*/
		[]string{"172.17.0.1:9092"},
		"test",
		true,
		nil)
}

// Declare variables to store log messages as new Events
var (
	successMessage = Event{4, "Success"}
)

func (l *StandardLogger) Success(requestType, respondTime string) {
	fields := LogFields(requestType, respondTime)

	l.WithFields(*fields).Info(successMessage.message)
}
