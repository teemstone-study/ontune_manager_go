package app

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
)

func TcpProcessing(reqChan chan<- *DataKey, resChan chan []byte, apiserver ApiServer) {
	fmt.Printf("TCP Server %s port\n", apiserver.Port)
	ln, err := net.Listen("tcp", fmt.Sprintf(":%s", apiserver.Port))
	if err != nil {
		log.Println("TCP Connection Error")
		return
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("Connection Closed")
			continue
		}

		go func(c net.Conn) {
			req_recv := make([]byte, 4096)
			for {
				n, err := conn.Read(req_recv)
				if err != nil {
					if io.EOF == err {
						return
					}
					log.Printf("Failed Connection: %v\n", err)
					reqChan <- &DataKey{DATAKEY_CODE, 0}
					return
				}
				if 0 < n {
					var req_recv_data DataKey
					if err := json.Unmarshal(req_recv[:n], &req_recv_data); err != nil {
						log.Printf("JSON Error")
						continue
					}
					//fmt.Printf("recv %v\n", req_recv_data)
					reqChan <- &req_recv_data
				}
			}
		}(conn)

		go func(c net.Conn) {
			for {
				msg := <-resChan
				msglen := make([]byte, 4)
				binary.LittleEndian.PutUint32(msglen, uint32(len(msg)))

				_, err := c.Write(msglen)
				if err != nil {
					log.Println(err)
					return
				}
				_, err = c.Write(msg)
				if err != nil {
					log.Println(err)
					return
				}
			}
		}(conn)
	}
}
