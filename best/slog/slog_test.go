package slog

import "testing"

func TestSlog(t *testing.T) {
	ShowLog()
}

func TestSlogWithFile(t *testing.T) {
	ShowLogWithFile()
}

func TestSlogWithFileWithLevel(t *testing.T) {
	ShowLogWithFileWithLevel()
}
func TestSlogWithCustom(t *testing.T) {
	ShowLogWithCustom()
}
func TestSlogWithCondition(t *testing.T) {
	ShowCondition()
}

func TestAsyncLog(t *testing.T) {
	AsyncLog()
}
func TestLogError(t *testing.T) {
	LogError()
}
func TestLogJson(t *testing.T) {
	LogJson()
}
