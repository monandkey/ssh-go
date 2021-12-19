package log

import (
	"sync"

	"github.com/fatih/color"
)

// Logger is a color logger with the pod and container name
type Logger struct {
	hostname string

	msgw  *color.Color
	hostw *color.Color
	m     *sync.Mutex
}

// Println prints the str as a single line with the hostname and container name
func (l *Logger) Println(str string) {
	l.m.Lock()
	defer l.m.Unlock()

	l.hostw.Add(color.Bold).Printf("%s", l.hostname)
	l.msgw.Println("|", str)
}
