package logger

import (
	"log"
	"os"
)

var (
	infoLogger  = log.New(os.Stdout, "[INFO] ", log.Ldate|log.Ltime)
	warnLogger  = log.New(os.Stdout, "[WARN] ", log.Ldate|log.Ltime)
	errorLogger = log.New(os.Stderr, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile)
)

func Info(format string, v ...interface{}) {
	infoLogger.Printf(format, v...)
}

func Warn(format string, v ...interface{}) {
	warnLogger.Printf(format, v...)
}

func Error(format string, v ...interface{}) {
	errorLogger.Printf(format, v...)
}

func ServerStart(port string) {
	Info("Server started and listening on port %s", port)
}

func DBConnected() {
	Info("Database connected successfully")
}

func LogTransaction(txType string, amount float64, gram float64) {
	Info("Transaction Successful | Type: %s | Amount: %.2f | Gram: %.4f", txType, amount, gram)
}
