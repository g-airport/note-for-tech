package TCP_Problem

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"net"
)

//tcp process unpacking && sticky bag problem

//head+body的方式
//send data format --> data_size:data
//+++++++++++++++++++++++++++++++++++++
//size (2 bytes)  | body (size bytes)
//+++++++++++++++++++++++++++++++++++++

var bufSize = 1024
var headSize = 4
var address = ":port"

//data tx
func doConn(conn net.Conn) {
	var (
		buffer    = bytes.NewBuffer(make([]byte, 0, bufSize)) //buffer as cache data that from readBytes
		readBytes = make([]byte, bufSize)                     //readBytes as accept data and cache to buffer
		isHead    = true                                      //identify current status : size or body
		bodyLen   = 0
	)

	for {
		//read data
		readByteNum, err := conn.Read(readBytes)
		if err != nil {
			log.Fatal(err)
			return
		}
		buffer.Write(readBytes[0:readByteNum])

		//deal data
		for {
			if isHead {
				if buffer.Len() >= headSize {
					isHead = false
					head := make([]byte, headSize)
					_, err = buffer.Read(head)
					if err != nil {
						log.Fatal(err)
						return
					}
					bodyLen = int(binary.BigEndian.Uint16(head))
				} else {
					break
				}
			}

			if !isHead {
				if buffer.Len() >= bodyLen {
					body := make([]byte, bodyLen)
					_, err = buffer.Read(body[:bodyLen])
					if err != nil {
						log.Fatal(err)
						return
					}
					fmt.Println("received body: " + string(body[:bodyLen]))
					isHead = true
				} else {
					break
				}
			}
		}
	}

}

//------------------------------------------------------------*
// server deal data from client
func Handle() {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println("start listening on 1234")
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
			return
		}
		go doConn(conn)
	}
}

//-------------------------------------------------------------*
// client send data to server
// use head :: body
func Send(c string) error {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatal(err)
		return err
	}

	head := make([]byte, headSize)
	content := []byte(c)
	headSize := len(content)
	binary.BigEndian.PutUint16(head, uint16(headSize))

	//first  write head
	//second write body
	_, err = conn.Write(head)
	if err != nil {
		log.Fatal(err)
		return err
	}
	_, err = conn.Write(content)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
