package log

import (
	"sync"

	"github.com/fatih/color"
)

// LoggerFactory is a factory for Logger
type LoggerFactory struct {
	color color.Attribute
	m     sync.Mutex
}

// NewLoggerFactory creates new LoggerFactory
func NewLoggerFactory() *LoggerFactory {
	return &LoggerFactory{}
}

// NewLogger creates a logger with a new color for the container of the hostname
func (f *LoggerFactory) NewLogger(hostname string) *Logger {
	f.color = f.nextColor()
	return &Logger{
		hostname: hostname,
		hostw:    color.New(f.color, color.Bold),
		msgw:     color.New(f.color),
		m:        &f.m,
	}
}

func (f *LoggerFactory) nextColor() color.Attribute {
	switch f.color {
	case color.FgRed:
		return color.FgGreen
	case color.FgGreen:
		return color.FgYellow
	case color.FgYellow:
		return color.FgBlue
	case color.FgBlue:
		return color.FgMagenta
	case color.FgMagenta:
		return color.FgCyan
	case color.FgCyan:
		return color.FgHiRed
	case color.FgHiRed:
		return color.FgHiGreen
	case color.FgHiGreen:
		return color.FgHiYellow
	case color.FgHiYellow:
		return color.FgHiBlue
	case color.FgHiBlue:
		return color.FgHiMagenta
	case color.FgHiMagenta:
		return color.FgHiCyan
	case color.FgHiCyan:
		return color.FgRed
	}
	return color.FgRed
}
