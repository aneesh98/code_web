package main

import (
	"fmt"
	"net/http"
	"time"
	"webterm-emulator/internal/log"
	"webterm-emulator/pkg"

	"github.com/gorilla/mux"
)

func main() {
	//routing
	router := mux.NewRouter()
	command := conf.GetString("command")
	connectionErrorLimit := conf.GetInt("connection-error-limit")
	arguments := conf.GetStringSlice("arguments")
	allowedHostnames := conf.GetStringSlice("allowed-hostnames")
	keepalivePingTimeout := time.Duration(conf.GetInt("keepalive-ping-timeout")) * time.Second
	maxBufferSizeBytes := conf.GetInt("max-buffer-size-bytes")
	// pathLiveness := conf.GetString("path-liveness")
	// pathMetrics := conf.GetString("path-metrics")
	// pathReadiness := conf.GetString("path-readiness")
	// pathXTermJS := conf.GetString("path-xtermjs")
	serverAddress := conf.GetString("server-addr")
	serverPort := conf.GetInt("server-port")
	// workingDirectory := conf.GetString("workdir")
	xtermjsHandlerOptions := pkg.HandlerOpts{
		AllowedHostNames:     allowedHostnames,
		Arguments:            arguments,
		Command:              command,
		ConnectionErrorLimit: connectionErrorLimit,
		CreateLogger: func(connectionUUID string, r *http.Request) pkg.Logger {
			createRequestLog(r, map[string]interface{}{"connection_uuid": connectionUUID}).Infof("created logger for connection '%s'", connectionUUID)
			return createRequestLog(nil, map[string]interface{}{"connection_uuid": connectionUUID})
		},
		KeepalivePingTimeout: keepalivePingTimeout,
		MaxBufferSizeBytes:   maxBufferSizeBytes,
	}
	router.HandleFunc("/terminal", pkg.GetHandler(xtermjsHandlerOptions))
	// listen
	listenOnAddress := fmt.Sprintf("%s:%v", serverAddress, serverPort)
	server := http.Server{
		Addr:    listenOnAddress,
		Handler: addIncomingRequestLogging(router),
	}

	log.Infof("starting server on interface:port '%s'...", listenOnAddress)
	server.ListenAndServe()
}
