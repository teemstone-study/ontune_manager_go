package app

import (
	"time"

	//"encoding/binary"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
)

var dataSucc bool

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
		fmt.Println("Client Accept!!")
		dataSucc = true
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
					dataSucc = true
					log.Printf("Failed Connection: %v\n", err)
					reqChan <- &DataKey{DATAKEY_CODE, 0}
					return
				}
				if 0 < n {
					if n == 1 && req_recv[0] == 0xFF {
						dataSucc = true
						//fmt.Println("0xFF Join!!")
						continue
					} else {
						var req_recv_data DataKey
						if err := json.Unmarshal(req_recv[:n], &req_recv_data); err != nil {
							log.Printf("JSON Error")
							continue
						}
						reqChan <- &req_recv_data
					}
				}
			}
		}(conn)

		go func(c net.Conn) {
			var last_Sendtime int64
			last_Sendtime = time.Now().Unix()
			for {

				for {
					if dataSucc {
						break
					} else {
						if last_Sendtime+int64(30) <= time.Now().Unix() {
							c.Close()
							fmt.Println("Conntion Close!!")
							return
						}
					}
					time.Sleep(time.Millisecond * 1)
				}
				msg := <-resChan
				last_Sendtime = time.Now().Unix()
				dataSucc = false
				msglen := make([]byte, 4)
				binary.LittleEndian.PutUint32(msglen, uint32(len(msg)))

				_, err := c.Write(msglen)
				if err != nil {
					log.Println(err)
					return
				}
				_, err = c.Write(msg)
				//log_write(fmt.Sprintf("LASTPERFCODE %d %d %s %s ", len(msg), nn, bytes.NewBuffer(msg[:20]).String(), bytes.NewBuffer(msg[len(msg)-20:])))
				if err != nil {
					log.Println(err)
					return
				}

			}
		}(conn)
	}
}
