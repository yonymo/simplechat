package msg

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/yonymo/simplechat/pkg/code"
	gin2 "github.com/yonymo/simplechat/pkg/common/gin"
	"github.com/yonymo/simplechat/pkg/errors"
	"github.com/yonymo/simplechat/pkg/middleware"
	"net/http"
	"time"

	"github.com/yonymo/simplechat/pkg/log"
)

const (
	ServerMessageLen uint = 100
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var server Server

type Server struct {
	conns_map map[uint]*Conn
	conns     map[*Conn]struct{}

	message chan *Message
}

func (s *Server) DoMsg() {
	for {
		msg := <-s.message
		peer, ok := s.conns_map[msg.DstId]
		if !ok {
			log.Warnf("client<%d> not found.\n", msg.DstId)
			continue
		}
		err := peer.conn.WriteJSON(msg)
		if err != nil {
			log.Errorf("write to peer<%d> failed\n", msg.DstId)
		}
	}
}

type Conn struct {
	server *Server

	// The websocket connection.
	conn *websocket.Conn

	id uint
}

type Message struct {
	MsgId       int64     `json:"id,omitempty" form:"id"`         //消息ID
	FromId      uint      `json:"userid,omitempty" form:"userid"` //谁发的
	DstId       uint      `json:"dstid,omitempty" form:"dstid"`   //对端用户ID
	ClientMsgId int64     `json:"client_msg_id" form:"client_msg_id"`
	Content     string    `json:"content,omitempty" form:"content"` //消息的内容
	Time        time.Time `json:"time" form:"time"`
	IsSender    bool      `json:"is_sender" form:"is_sender"`
}

type MessageRsp struct {
	MsgId int64 `json:"msgid,omitempty" form:"id"`
}

func (c *Conn) read_loop() {
	log.Debug("read_loop")
	for {
		rmsg := Message{}
		err := c.conn.ReadJSON(&rmsg)
		if err != nil {
			log.Errorf("ReadJSON failed: %v\n", err)
			continue
		}

		c.server.message <- &rmsg

		c.conn.WriteJSON(&MessageRsp{
			MsgId: rmsg.ClientMsgId,
		})

		//fmt.Println(msgType, data, err)

	}
}

func Open(c *gin.Context) {
	val, ok := c.Get(middleware.KeyUserID)
	if !ok {
		gin2.WriteResponse(c, errors.WithCode(code.ErrServerInternal, "not found uid"), nil)
		return
	}
	uid := uint(val.(float64))

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Errorf("Open websocket failed: %v\n", err.Error())
		return
	}

	cli := &Conn{
		conn:   conn,
		server: &server,
		id:     uid,
	}

	server.conns_map[uid] = cli
	server.conns[cli] = struct{}{}

	go cli.read_loop()
}

func init() {
	server = Server{}
	server.conns_map = make(map[uint]*Conn)
	server.conns = make(map[*Conn]struct{})
	server.message = make(chan *Message, ServerMessageLen)
	go server.DoMsg()
}
