package binance

import (
	"net/url"

	"github.com/gorilla/websocket"
)

const binanceWsAddr = "stream.binance.com:9443"
const binanceWsScheme = "wss"

func connectWebsocket(path string) (*websocket.Conn, error) {
	socketURL := url.URL{
		Scheme: binanceWsScheme,
		Host:   binanceWsAddr,
		Path:   "ws/" + path,
	}
	socket, _, err := websocket.DefaultDialer.Dial(socketURL.String(), nil)
	return socket, err
}
