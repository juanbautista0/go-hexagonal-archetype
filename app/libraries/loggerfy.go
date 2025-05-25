package libraries

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type LoggerLevel string

const (
	Info    LoggerLevel = "INFO"
	Warning LoggerLevel = "WARNING"
	Error   LoggerLevel = "ERROR"

	nullUUID = "00000000-0000-0000-0000-000000000000"
)

type LogEntry struct {
	Timestamp   string                 `json:"timestamp,omitempty"`
	ID          string                 `json:"id,omitempty"`
	Code        string                 `json:"code,omitempty"`
	Message     string                 `json:"message,omitempty"`
	Detail      string                 `json:"detail,omitempty"`
	Payload     map[string]interface{} `json:"payload,omitempty"`
	Level       LoggerLevel            `json:"level,omitempty"`
	Severity    LoggerLevel            `json:"severity,omitempty"`
	Service     string                 `json:"service,omitempty"`
	Environment string                 `json:"environment,omitempty"`
}

type LoggerfyBase struct {
	id          string
	code        string
	message     string
	detail      string
	metadata    map[string]interface{}
	level       LoggerLevel
	service     string
	environment string
}

func NewLoggerfyBase(level LoggerLevel) *LoggerfyBase {
	return &LoggerfyBase{
		id:          nullUUID,
		level:       level,
		service:     getEnv("SERVICE_NAME", "default-service"),
		environment: getEnv("NODE_ENV", "development"),
	}
}

func (l *LoggerfyBase) SetCode(code string) *LoggerfyBase {
	l.code = code
	return l
}

func (l *LoggerfyBase) SetMessage(msg string) *LoggerfyBase {
	l.message = msg
	return l
}

func (l *LoggerfyBase) SetDetail(detail string) *LoggerfyBase {
	l.detail = detail
	return l
}

func (l *LoggerfyBase) SetMetadata(metadata map[string]interface{}) *LoggerfyBase {
	l.metadata = metadata
	return l
}

func (l *LoggerfyBase) Write(customID ...string) {
	if l.code == "" || l.message == "" || l.detail == "" {
		return
	}

	if len(customID) > 0 {
		l.id = customID[0]
	} else {
		l.id = NewUuid().String()
	}

	log := LogEntry{
		Timestamp:   time.Now().UTC().Format(time.RFC3339),
		ID:          l.id,
		Code:        l.code,
		Message:     l.message,
		Detail:      l.detail,
		Payload:     l.metadata,
		Level:       l.level,
		Severity:    l.level,
		Service:     l.service,
		Environment: l.environment,
	}

	go func(e LogEntry) {
		jsonLog, _ := json.Marshal(log)
		fmt.Println(string(jsonLog))
	}(log)

	l.reset()
}

func (l *LoggerfyBase) reset() {
	l.id = nullUUID
	l.code = ""
	l.message = ""
	l.detail = ""
	l.metadata = nil
}

// Loggerfy provides a simple factory
type Loggerfy struct {
}

func NewLoggerfy() *Loggerfy {

	return &Loggerfy{}
}

func (l *Loggerfy) Info() *LoggerfyBase  { return NewLoggerfyBase(Info) }
func (l *Loggerfy) Warn() *LoggerfyBase  { return NewLoggerfyBase(Warning) }
func (l *Loggerfy) Error() *LoggerfyBase { return NewLoggerfyBase(Error) }

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
