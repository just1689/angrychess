package controller

import (
	"encoding/json"
	"fmt"
	"github.com/team142/chessfor4/io/ws"
	"github.com/team142/chessfor4/model"
	"log"
)

func handleInMessageMove(server *model.Server, client *ws.Client, msg []byte) {
	message := &model.MessageMove{}
	if err := json.Unmarshal(msg, message); err != nil {
		log.Println(fmt.Sprintf("Error unmarshaling, %s", err))
		return
	}
	move(server, message, client)

}
