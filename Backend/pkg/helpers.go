package pkg

import (
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
)

func getConnectionUpgrader(
	allowedHostnames []string,
	maxBufferSizeBytes int,
	logger Logger,
) websocket.Upgrader {
	return websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			requesterHostName := r.Host
			if strings.Index(requesterHostName, ":") != -1 {
				requesterHostName = strings.Split(requesterHostName, ":")[0]
			}
			for _, allowedHostName := range allowedHostnames {
				if requesterHostName == allowedHostName {
					return true
				}
			}
			logger.Warnf("failed to find '%s' in the list of allowed hostnames ('%s')", requesterHostName)
			return false

		},
		HandshakeTimeout: 0,
		ReadBufferSize:   maxBufferSizeBytes,
		WriteBufferSize:  maxBufferSizeBytes,
	}
}
