package typingwars

import (
	"fmt"
	"log"

	"github.com/Jaacky/typingwars/pb"
	"github.com/gofrs/uuid"
	"github.com/golang/protobuf/proto"
)

type Room struct {
	ID          uuid.UUID
	clients     map[uuid.UUID]*Client
	players     map[uuid.UUID]*Player
	readyStatus map[uuid.UUID]bool
}

func NewRoom() *Room {
	id, err := uuid.NewV4()
	if err != nil {
		log.Fatalf("Failed to generate uuid: %v", err)
	}

	return &Room{
		ID:          id,
		clients:     make(map[uuid.UUID]*Client),
		players:     make(map[uuid.UUID]*Player),
		readyStatus: make(map[uuid.UUID]bool),
	}
}

func (room *Room) addClient(client *Client, username string) {
	currentPlayer := NewPlayer(client.ID, username)
	room.clients[client.ID] = client
	room.players[client.ID] = currentPlayer
	room.readyStatus[client.ID] = false

	joinRoomAck := &pb.JoinRoomAck{
		ClientId: fmt.Sprintf("%s", client.ID),
	}

	joinRoomAckMessage := &pb.UserMessage{
		Content: &pb.UserMessage_JoinRoomAck{
			JoinRoomAck: joinRoomAck,
		},
	}

	room.SendToClient(client.ID, joinRoomAckMessage)
	room.update()
}

func (room *Room) updatePlayerReady(clientID uuid.UUID, readyStatus bool) {
	log.Println("Updating player ready status for client %s", clientID)
	if _, ok := room.readyStatus[clientID]; ok {
		room.readyStatus[clientID] = readyStatus
		log.Printf("Updated client: %s ready status to: %t", clientID, readyStatus)
		room.update()
	}
}

func (room *Room) update() {
	pbPlayers := make(map[string]*pb.Player)
	pbReadyStatus := make(map[string]bool)
	for id, player := range room.players {
		idString := id.String()
		pbPlayer := &pb.Player{
			Id:       idString,
			Username: player.Username,
		}

		pbPlayers[idString] = pbPlayer
		pbReadyStatus[idString] = room.readyStatus[id]
	}

	updateRoom := &pb.UpdateRoom{
		RoomId:      fmt.Sprintf("%s", room.ID),
		Players:     pbPlayers,
		ReadyStatus: pbReadyStatus,
	}

	updateRoomMessage := &pb.UserMessage{
		Content: &pb.UserMessage_UpdateRoom{
			UpdateRoom: updateRoom,
		},
	}

	room.SendToAllClients(updateRoomMessage)
}

func (room *Room) SendToClient(clientID uuid.UUID, message proto.Message) {
	client, ok := room.clients[clientID]
	if ok {
		client.SendMessage(message)
	} else {
		log.Printf("Client %d not found\n", clientID)
		return
	}
}

func (room *Room) SendToAllClients(message proto.Message) {
	for _, client := range room.clients {
		client.SendMessage(message)
	}
}
