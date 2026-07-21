package logger

import (
	"bytes"
	"log"
	"os"
	"strings"
	"testing"
)

func TestLoggers(t *testing.T) {
	var buf bytes.Buffer

	infoLogger.SetOutput(&buf)
	warnLogger.SetOutput(&buf)
	errorLogger.SetOutput(&buf)

	defer func() {
		infoLogger.SetOutput(os.Stdout)
		warnLogger.SetOutput(os.Stdout)
		errorLogger.SetOutput(os.Stderr)
	}()

	t.Run("Test Info Log", func(t *testing.T) {
		buf.Reset()
		Info("Hello %s", "World")
		if !strings.Contains(buf.String(), "[INFO]") || !strings.Contains(buf.String(), "Hello World") {
			t.Errorf("Format Info log salah: %s", buf.String())
		}
	})

	t.Run("Test Warn Log", func(t *testing.T) {
		buf.Reset()
		Warn("Disk usage %d%%", 90)
		if !strings.Contains(buf.String(), "[WARN]") || !strings.Contains(buf.String(), "Disk usage 90%") {
			t.Errorf("Format Warn log salah: %s", buf.String())
		}
	})

	t.Run("Test Error Log", func(t *testing.T) {
		buf.Reset()
		errorLogger.SetFlags(0)
		Error("Failed to %s", "connect")
		if !strings.Contains(buf.String(), "[ERROR]") || !strings.Contains(buf.String(), "Failed to connect") {
			t.Errorf("Format Error log salah: %s", buf.String())
		}
		errorLogger.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	})

	t.Run("Test Domain Specific Logs", func(t *testing.T) {
		buf.Reset()
		ServerStart("8080")
		if !strings.Contains(buf.String(), "Server started and listening on port 8080") {
			t.Errorf("ServerStart log salah: %s", buf.String())
		}

		buf.Reset()
		DBConnected()
		if !strings.Contains(buf.String(), "Database connected successfully") {
			t.Errorf("DBConnected log salah: %s", buf.String())
		}

		buf.Reset()
		LogTransaction("BUY", 500000.0, 0.2571)
		if !strings.Contains(buf.String(), "Transaction Successful") || !strings.Contains(buf.String(), "BUY") {
			t.Errorf("LogTransaction log salah: %s", buf.String())
		}
	})
}
