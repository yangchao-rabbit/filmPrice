package logger

import (
	"testing"
)

func TestNew(t *testing.T) {

	logger := New(WithDefault())
	logger.Println("test")
}
