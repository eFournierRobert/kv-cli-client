package api

import (
	"errors"
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
	if err != nil {
		return err
	}
	defer conn.Close()

	if n, err := conn.Write(packet); err != nil {
		return err
	} else if n < len(packet) {
		return errors.New("packet wasn't fully sent")
	}

	return nil
}
