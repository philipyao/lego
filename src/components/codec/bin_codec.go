package codec

import (
	"net"
	"io"
	"errors"
	"binary"
)

const (
	HEAD_BUFF_LEN       = 4
)

var head [HEAD_BUFF_LEN]byte

// -------------------------------------------------------
// 二进制协议： | Len(2) | Version(1) | Encrypt(1) | Data |
// -------------------------------------------------------
// 统一采用大端字节序
type BinMsg struct {
	Len     Uint16      //长度
	Version Uint8       //版本号
	Encrypt Uint8       //加密方式
	Data    []byte      //包体数据
}


func ReadBinMsg(conn net.Conn) ([]byte, error) {
	hd := head[:]
	n, err := io.ReadFull(conn, hd)
	if err != nil {
		return nil, err
	}
	if n != len(head) {
		//todo
		return nil, errors.New("read bin msg head failed")
	}

	msgLen := binary.BigEndian.Uint16(head[:2])
	//todo
	msgVer := Uint8(head[2])
	//todo
	msgEnc := uint8(head[3])
	//todo

	bd := make([]byte, msgLen)
	n, err = io.ReadFull(conn, bd)
	if err != nil {
		return nil, err
	}
	if n != msgLen {
		//todo
	}

	//解密包体
	return bd, nil
}


func WriteBinMsg(conn net.Conn) error {

}