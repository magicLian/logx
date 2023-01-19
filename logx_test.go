package logx

import "testing"

func TestInfo(t *testing.T) {
	logx := NewLogx("app_service", ErrorLevel)
	logx.Info("aaaaaaaa")
}

func TestError(t *testing.T) {
	logx := NewLogx("app_service", DebugLevel)
	k := 3
	logx.Infof("k is %d", k)
}

func TestInfof(t *testing.T) {
	logx := NewLogx("app_service", TraceLevel)
	k := 3
	logx.Infof("k is %d", k)
}

func TestDebugf(t *testing.T) {
	logx := NewLogx("app_service", DebugLevel)
	k := 3
	logx.Debugf("k is %d", k)
}
