package zap

type Logger struct{}

func (l *Logger) Info(msg string, fields ...interface{})  {}
func (l *Logger) Error(msg string, fields ...interface{}) {}
func (l *Logger) Warn(msg string, fields ...interface{})  {}
func (l *Logger) Debug(msg string, fields ...interface{}) {}
func (l *Logger) Fatal(msg string, fields ...interface{}) {}
func (l *Logger) Panic(msg string, fields ...interface{}) {}

type SugaredLogger struct{}

func (s *SugaredLogger) Infof(template string, args ...interface{})  {}
func (s *SugaredLogger) Errorf(template string, args ...interface{}) {}
func (s *SugaredLogger) Warnf(template string, args ...interface{})  {}
func (s *SugaredLogger) Debugf(template string, args ...interface{}) {}
func (s *SugaredLogger) Fatalf(template string, args ...interface{}) {}
func (s *SugaredLogger) Panicf(template string, args ...interface{}) {}

func (s *SugaredLogger) Infow(msg string, keysAndValues ...interface{})  {}
func (s *SugaredLogger) Errorw(msg string, keysAndValues ...interface{}) {}
func (s *SugaredLogger) Warnw(msg string, keysAndValues ...interface{})  {}
func (s *SugaredLogger) Debugw(msg string, keysAndValues ...interface{}) {}
func (s *SugaredLogger) Fatalw(msg string, keysAndValues ...interface{}) {}
func (s *SugaredLogger) Panicw(msg string, keysAndValues ...interface{}) {}
