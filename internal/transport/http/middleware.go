package http

import (
	"github.com/julienschmidt/httprouter"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func (h *Handler) loggingMiddleware(next httprouter.Handle) httprouter.Handle {
	logFile, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Failed to create log file: %v", err)
	}
	defer logFile.Close()

	logger := log.New(io.MultiWriter(os.Stdout, logFile), "", log.Ldate|log.Ltime|log.Lshortfile)

	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		logger.Printf("%s: [%s] - %s ", time.Now().Format(time.RFC3339), r.Method, r.RequestURI)
		next(w, r, params)
	}
}
