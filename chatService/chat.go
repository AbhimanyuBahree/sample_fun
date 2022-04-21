package chatService

import (
	"context"
	"log"
)

type Server struct {}

func(s * Server) SayHello(ctx context.Context, message * Message) (* Message, error){
	log.Printf("Received body from the client : %s", message.body)
	return &Message{Body: "Hello from the Server!"}, nil
	
}
