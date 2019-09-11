package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	http.HandleFunc("/", handler)

	logrus.Info("Starting server on port 0.0.0.0:8126")

	err := http.ListenAndServe(fmt.Sprintf(":%d", 8126), nil)
	if err != nil {
		logrus.Fatal(err.Error())
	}
}

func handler(writer http.ResponseWriter, request *http.Request) {
	logrus.WithFields(logrus.Fields{
		"remote_address": request.RemoteAddr,
		"method": request.Method,
		"url": request.URL.Path,
	}).Info("Received APM data from a Datadog client")

	_, _ = fmt.Fprint(writer, "{\"rate_by_service\": null}")
}
