package main

import (
	"bufio"
	"context"
	"google.golang.org/grpc"
	"io"
	"log"
	"os"
	"strings"
	pb "study/chatting/api"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewChatServiceClient(conn)

	reader := bufio.NewReader(os.Stdin)

	for {
		showMainMenu()
		option, _ := reader.ReadString('\n')
		option = strings.TrimSpace(option)

		switch option {
		case "1":
			manageChannelMenu(c)
		case "2":
			sendMessageMenu(c)
		case "3":
			listChannelsMenu(c)
		case "4":
			log.Println("Exiting...")
			return
		default:
			log.Println("Invalid option, please try again.")
		}
	}
}

func showMainMenu() {
	log.Println("\nChoose an option:")
	log.Println("1) Manage channel")
	log.Println("2) Send message")
	log.Println("3) List channels")
	log.Println("4) Exit")
	log.Print("Enter option: ")
}

func manageChannelMenu(c pb.ChatServiceClient) {
	reader := bufio.NewReader(os.Stdin)
	log.Println("\nChannel Management")
	log.Println("Enter channel name:")
	channelName, _ := reader.ReadString('\n')
	channelName = strings.TrimSpace(channelName)

	log.Println("Choose action: 1) Create 2) Delete")
	action, _ := reader.ReadString('\n')
	action = strings.TrimSpace(action)

	var actionStr string
	if action == "1" {
		actionStr = "create"
	} else if action == "2" {
		actionStr = "delete"
	} else {
		log.Printf("Invalid action : %s", actionStr)
		return
	}

	log.Printf("action : %s", actionStr)

	res, err := c.ManageChannel(context.Background(), &pb.ChannelRequest{
		Name:   channelName,
		Action: actionStr,
	})
	if err != nil {
		log.Fatalf("could not manage channel : %v", err)
	}
	log.Printf("Channel response : %s", res.Message)
}

func listChannelsMenu(c pb.ChatServiceClient) {
	res, err := c.ListChannels(context.Background(), &pb.ChannelListRequest{})
	if err != nil {
		log.Fatalf("could not list channels : %v", err)
	}

	log.Println("Channels: ")
	for _, channel := range res.Channels {
		log.Println(channel)
	}
}

func sendMessageMenu(c pb.ChatServiceClient) {
	reader := bufio.NewReader(os.Stdin)
	log.Println("Enter channel name to join:")
	channelName, _ := reader.ReadString('\n')
	channelName = strings.TrimSpace(channelName)

	log.Println("Enter your user name:")
	userName, _ := reader.ReadString('\n')
	userName = strings.TrimSpace(userName)

	stream, err := c.SendMessage(context.Background())
	if err != nil {
		log.Fatalf("could not send message: %v", err)
		return
	}

	// 서버에 초기 메시지 전송하여 채널 접속 알림
	if err := stream.Send(&pb.ChatMessage{
		Channel: channelName,
		User:    userName,
		Message: "has joined the channel",
	}); err != nil {
		log.Fatalf("could not send message: %v", err)
		return
	}

	// 별도의 고루틴에서 메시지 수신
	go receiveMessages(stream)

	// 사용자로부터 메시지를 입력받아 서버로 전송
	for {
		log.Println("Enter your message (or type 'exit' to leave, 'users' to list users):")
		message, _ := reader.ReadString('\n')
		message = strings.TrimSpace(message)

		if message == "exit" {
			return
		} else if message == "users" {
			listChannelUsers(c, channelName)
			continue
		}

		if err := stream.Send(&pb.ChatMessage{
			Channel: channelName,
			User:    userName,
			Message: message,
		}); err != nil {
			log.Fatalf("could not send message to channel: %v", err)
		}
	}
}

func receiveMessages(stream pb.ChatService_SendMessageClient) {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			// 연결 종료 처리
			log.Println("Disconnected from the channel")
			return
		}
		if err != nil {
			log.Fatalf("Failed to receive a message: %v", err)
		}

		if in.User == "System" && in.Message == "Channel has been deleted" {
			log.Println("The channel has been deleted. Returning to main menu.")
			return
		}

		log.Printf("[%s] %s: %s", in.Channel, in.User, in.Message)
	}
}

func listChannelUsers(c pb.ChatServiceClient, channelName string) {
	res, err := c.ListChannelUsers(context.Background(), &pb.ChannelUsersRequest{
		Channel: channelName,
	})
	if err != nil {
		log.Fatalf("Error listing channel users: %v", err)
	}

	log.Println("Users in channel:", channelName)
	for _, user := range res.Users {
		log.Println(user)
	}
}
