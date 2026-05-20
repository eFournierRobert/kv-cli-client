package api

import (
	"kv-cli-client/internal/models"
	"net"
)

const server = "127.0.0.1:8080"

func SendSet(key, value string) {
	set := models.Set{Key: key, Value: value}
	packet := set.ToBytes()

}

func sendPacket(packet []byte) error {
	conn, err := net.Dial("tcp", server)
	defer conn.Close()
}
