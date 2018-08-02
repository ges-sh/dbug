package dbugdb

// Logger is an type used for database debugging.
// It wraps basic log method, Debug, which should be implemented in most logging libraries (logrus, zap, go-logging)
type Logger interface {
	Debug(args ...interface{})
}
