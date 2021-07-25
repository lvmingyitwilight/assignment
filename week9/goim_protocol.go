package week9

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"net"
)

type Conn struct {
	conn    net.Conn
	scanner *bufio.Scanner
}

func NewConn(conn net.Conn) *Conn {
	scanner := bufio.NewScanner(conn)
	scanner.Split(split)
	return &Conn{conn: conn, scanner: scanner}
}

func (c *Conn) Read() (buf []byte, err error) {
	if c.scanner.Scan() {
		buf = c.scanner.Bytes()
	} else {
		err = c.scanner.Err()
	}
	return
}

func (c *Conn) Write(data []byte) (err error) {
	_, err = c.conn.Write(encoder(data))
	return
}

func (c *Conn) Close() error {
	return c.conn.Close()
}

const (
	headerLength = 16
)

var (
	protocolVersion = 1
	operation       = 1
	sequenceID      = 1
)

/*
	Package Length 	 4bytes
	Header Length 	 2bytes
	Protocol Version 2bytes
	Operation		 4bytes
	SequenceID		 4bytes
	Body			 (PackageLength - HeaderLength)bytes
*/

func encoder(body []byte) []byte {
	buf := bytes.Buffer{}
	packageLength := headerLength + len(body)
	binary.Write(&buf, binary.BigEndian, uint32(packageLength))
	binary.Write(&buf, binary.BigEndian, uint16(headerLength))
	binary.Write(&buf, binary.BigEndian, uint16(protocolVersion))
	binary.Write(&buf, binary.BigEndian, uint32(operation))
	binary.Write(&buf, binary.BigEndian, uint32(sequenceID))
	binary.Write(&buf, binary.BigEndian, body)
	fmt.Println("encode: ", string(buf.Bytes()), buf.Len())
	return buf.Bytes()
}

//解决粘包
func split(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if len(data) > 4 {
		packageLength := binary.BigEndian.Uint32(data[:4])
		//判断包长
		if int(packageLength) > len(data) {
			return
		}
		advance, token = int(packageLength), data[:packageLength]
	}
	if atEOF {
		err = io.EOF
	}
	return
}
