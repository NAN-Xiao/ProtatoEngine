package connection

import (
	"encoding/binary"
	"io"
	"net"
	"potatoengine/src/netmessage"
)

type TcpConnect struct {
	conn []*net.TCPConn
}

func (this *TcpConnect) Read() (l int, err error) {
	go func() {

	}()
	return
}
func (this *TcpConnect) Write(data []byte) {

}
func (this *TcpConnect) Close() bool {
	return false
}

//整个server监听
func (this *TcpConnect) Listen() {
	go func() {

		addr, err := net.ResolveTCPAddr("tcp", "0.0.0.0:9000")
		if err != nil {
			println(err)
			return
		}
		lisener, err := net.ListenTCP("tcp", addr)
		for {
			c, err := lisener.AcceptTCP()
			if err != nil {
				println(err)
				return
			}
			println("new client")
			this.conn = append(this.conn, c)

			go func(conn *net.TCPConn) {
				println("tcp listening")
				for {
					var buf = make([]byte, 4)
					n, err := io.ReadFull(conn, buf)
					if err == io.EOF {
						break
					}
					if n < 4 {
						continue
					}
					len := binary.BigEndian.Uint32(buf) - 4
					buf = make([]byte, len)
					n, _ = io.ReadFull(conn, buf)
					if n < 4 {
						continue
					}
					id, obj := netmessage.UnPackNetMessage(buf)
					if id < 0 || obj == nil {
						//消息错误
						continue
					}
					//todo 接受数据分发消息
					println("message")
				}
				conn.Close()
				println("CLose tcp con")
			}(c)
		}

	}()

}
