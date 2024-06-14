package logger

import (
	"bytes"
	"log"
	"testing"
)

func TestLogger_Info(t *testing.T) {
	var buf bytes.Buffer
	logger := log.New(&buf, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	testLogger := &Logger{
		infoLogger: logger,
	}

	testLogger.Info("This is an info message")
	if !bytes.Contains(buf.Bytes(), []byte("This is an info message")) {
		t.Errorf("Expected 'This is an info message' to be in log but got %s", buf.String())
	}
}

func TestLogger_Error(t *testing.T) {
	var buf bytes.Buffer
	logger := log.New(&buf, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	testLogger := &Logger{
		errorLogger: logger,
	}

	testLogger.Error("This is an error message")
	if !bytes.Contains(buf.Bytes(), []byte("This is an error message")) {
		t.Errorf("Expected 'This is an error message' to be in log but got %s", buf.String())
	}
}

func TestLogger_Fatal(t *testing.T) {
	// Note: log.Fatal() will call os.Exit(1), so it will stop the test execution
	// We will not test this function in the usual way to avoid exiting the test runner
}
