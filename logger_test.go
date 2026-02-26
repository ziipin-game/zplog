package zplog

import (
	"fmt"
	"os"
	"testing"
)

func TestSimpleLog(t *testing.T) {
	Debugf("1")
	Infof("2")
	Warnf("3")
	Errorf("4")
	Fatalf("5")
}

func TestSetLogLevel(t *testing.T) {
	SetLogLevel(LOG_WARN)
	Debugf("1")
	Infof("2")
	Warnf("3")
	Errorf("4")
	Fatalf("5")
}

func TestNewLogger(t *testing.T) {
	SetLogLevel(LOG_WARN)
	Debugf("1")
	Infof("2")
	Warnf("%d", 3)
	Errorf("4")
	Fatalf("5")
	logger := NewLogger(os.Stdout, "[logger]-")
	logger.SetLogLevel(LOG_DEBUG)
	logger.Debugf("x %d", 1)
	logger.Infof("x 2")
	logger.Warnf("x 3")
	logger.Errorf("x 4")
	logger.Fatalf("x 5")
}

type Room struct {
	Id int
}

func (r *Room) Work() {
	r.Errorf("something wrong here")
}

func (r *Room) Errorf(format string, args ...interface{}) {
	UpErrorf(1, fmt.Sprintf("room: %d %s", r.Id, fmt.Sprintf(format, args...)))
}

func TestLogUpper(t *testing.T) {
	room := new(Room)
	room.Id = 10000
	room.Work()
}
