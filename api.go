package zplog

func SetLogLevel(logLevel int) {
	defaultLogger.SetLogLevel(logLevel)
}

func Debugf(format string, args ...interface{}) bool {
	return defaultLogger.log(LOG_DEBUG, 3, format, args...)
}

func Infof(format string, args ...interface{}) bool {
	return defaultLogger.log(LOG_INFO, 3, format, args...)
}

func Warnf(format string, args ...interface{}) bool {
	return defaultLogger.log(LOG_WARN, 3, format, args...)
}

func Errorf(format string, args ...interface{}) bool {
	return defaultLogger.log(LOG_ERROR, 3, format, args...)
}

func Fatalf(format string, args ...interface{}) bool {
	return defaultLogger.log(LOG_FATAL, 3, format, args...)
}

func UpDebugf(up int, format string, args ...interface{}) bool {
	return defaultLogger.log(LOG_DEBUG, 3+up, format, args...)
}

func UpInfof(up int, format string, args ...interface{}) bool {
	return defaultLogger.log(LOG_INFO, 3+up, format, args...)
}

func UpWarnf(up int, format string, args ...interface{}) bool {
	return defaultLogger.log(LOG_WARN, 3+up, format, args...)
}

func UpErrorf(up int, format string, args ...interface{}) bool {
	return defaultLogger.log(LOG_ERROR, 3+up, format, args...)
}

func UpFatalf(up int, format string, args ...interface{}) bool {
	return defaultLogger.log(LOG_FATAL, 3+up, format, args...)
}
