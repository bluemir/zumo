package server

import (
	"encoding/json"
	"io"

	"github.com/bluemir/zumo/datatype"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/websocket"
)

type translater struct {
	encoder *json.Encoder
	decoder *json.Decoder
}

func (t *translater) OnMessage(channelID string, msg datatype.Message) error {
	return t.encoder.Encode(struct {
		*datatype.Message
		Type      string
		ChannelID string
	}{&msg, "message", channelID})
}
func (t *translater) channelChnaged(channelID string) {
	t.encoder.Encode(struct {
		Type string
		Name string
		Data map[string]string
	}{"event", "channel", map[string]string{"ID": channelID}})
}
func (t *translater) OnJoinChannel(channelID string) {
	// send to client
	t.channelChnaged(channelID)
}
func (t *translater) OnLeaveChannel(channelID string) {
	// send to client
	t.channelChnaged(channelID)
}
func (t *translater) runDispatcher() {
	// Wait until close socket
	for {
		msg := map[string]interface{}{}
		if err := t.decoder.Decode(msg); err == io.EOF {
			logrus.Debugf("[websocket:translater:runDispatcher] Disconnected")
			return
		} else if err != nil {
			// skip;
		}
	}
}

func (server *Server) ws(c *gin.Context) {
	username := c.MustGet(keyUsername).(string)
	websocket.Handler(func(conn *websocket.Conn) {
		defer conn.Close()

		encoder := json.NewEncoder(conn)
		decoder := json.NewDecoder(conn)

		agent := &translater{encoder, decoder}

		server.backend.RegisterUserAgent(username, agent)

		agent.runDispatcher()

	}).ServeHTTP(c.Writer, c.Request)
}
