package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/marmotedu/errors"
	"go-ecm/internal/pkg/code"
	"go-ecm/internal/pkg/docker"
	"go-ecm/pkg/core"
	"io"
	"net/http"
	"unicode/utf8"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

const readBufferSize = 8192

func Terminal(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer ws.Close()
	containerId := c.Query("container")
	command := c.Query("command")
	if containerId == "" {
		core.WriteResponse(c, errors.WithCode(code.ErrBind, "无法识别容器id"), nil)
		return
	}
	cli, _ := docker.NewDockerClient()
	stream, err := cli.Exec(c, containerId, command)
	if err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrBind, "无法识别容器id"), nil)
		return
	}
	defer stream.Close()
	defer func() {
		stream.Conn.Write([]byte("exit\r"))
	}()

	go func() {
		wsWriter(stream.Conn, ws)
	}()
	wsReader(ws, stream.Conn)
}

func wsWriter(reader io.Reader, writer *websocket.Conn) {
	for {
		out := make([]byte, readBufferSize)
		_, err := reader.Read(out)
		if err != nil {
			break
		}
		processedOutput := validString(string(out[:]))
		err = writer.WriteMessage(websocket.TextMessage, []byte(processedOutput))
		if err != nil {
			break
		}
	}
}

func wsReader(reader *websocket.Conn, writer io.Writer) {
	for {
		_, in, err := reader.ReadMessage()
		if err != nil {
			break
		}
		_, err = writer.Write(in)
		if err != nil {
			break
		}
	}
}

func validString(s string) string {
	if !utf8.ValidString(s) {
		v := make([]rune, 0, len(s))
		for i, r := range s {
			if r == utf8.RuneError {
				_, size := utf8.DecodeRuneInString(s[i:])
				if size == 1 {
					continue
				}
			}
			v = append(v, r)
		}
		s = string(v)
	}
	return s
}
