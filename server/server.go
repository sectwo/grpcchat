package main

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	pb "study/chatting/api" // 생성된 Protobuf 코드를 위한 패키지
	"sync"
)

type server struct {
	pb.UnimplementedChatServiceServer
	channels map[string]map[string]pb.ChatService_SendMessageServer // 채널 관리를 위한 맵 - ex) 해당하는 이름의 채널이 활성화 되어있으면 bool 형태로 동작 유무 표시
	mu       sync.Mutex                                             // 고 루틴을 위한 뮤텍스
}

func newServer() *server {
	return &server{
		channels: make(map[string]map[string]pb.ChatService_SendMessageServer),
	}
}

func (s *server) ManageChannel(ctx context.Context, in *pb.ChannelRequest) (*pb.ChannelResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	switch in.Action {
	case "create":
		// 채널 생성 로직
		if _, exists := s.channels[in.Name]; exists {
			return &pb.ChannelResponse{
				Name:    in.Name,
				Success: false,
				Message: "Channel already exists",
			}, nil
		}
		s.channels[in.Name] = make(map[string]pb.ChatService_SendMessageServer)
		return &pb.ChannelResponse{
			Name:    in.Name,
			Success: true,
			Message: "Channel created successfully",
		}, nil

	case "delete":
		// 채널 제거 로직
		if channel, exists := s.channels[in.Name]; exists {
			for _, clientStream := range channel {
				_ = clientStream.Send(&pb.ChatMessage{
					User:    "System",
					Message: "Channel has been deleted",
				})
			}
			delete(s.channels, in.Name) // 채널 삭제는 모든 클라이언트에게 메시지를 보낸 후에 수행
			return &pb.ChannelResponse{
				Name:    in.Name,
				Success: true,
				Message: "Channel deleted successfully",
			}, nil
		} else {
			return &pb.ChannelResponse{
				Name:    in.Name,
				Success: false,
				Message: "Channel does not exist",
			}, nil
		}

	default:
		return nil, fmt.Errorf("invalid action")
	}
}

func (s *server) SendMessage(stream pb.ChatService_SendMessageServer) error {
	// 초기 연결 설정 및 클라이언트 식별
	// 첫 번째 메시지를 사용하여 클라이언트와 채널을 식별
	initalMsg, err := stream.Recv()
	if err != nil {
		return err
	}
	channelName := initalMsg.GetChannel()

	s.mu.Lock()
	_, exists := s.channels[channelName]
	s.mu.Unlock()

	if !exists {
		// 채널이 존재하지 않는 경우, 클라이언트에게 알림
		return errors.New("channel does not exist")
	}

	userName := initalMsg.GetUser()

	s.mu.Lock()
	if s.channels[channelName] == nil {
		s.channels[channelName] = make(map[string]pb.ChatService_SendMessageServer)
	}
	s.channels[channelName][userName] = stream
	s.mu.Unlock()

	defer func() {
		s.mu.Lock()
		delete(s.channels[channelName], userName)
		s.mu.Unlock()
	}()

	// 메시지 수신 및 해당 채널의 모든 클라이언트에게 전송
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			// 클라이언트 연결 종료 처리
			return nil
		}
		if err != nil {
			// 에러 처리
			return err
		}
		s.mu.Lock()
		for _, clientStream := range s.channels[in.Channel] {
			if err := clientStream.Send(in); err != nil {
				// 에러 처리
			}
		}
		s.mu.Unlock()
	}
}

func (s *server) ListChannels(ctx context.Context, in *pb.ChannelListRequest) (*pb.ChannelListResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var channels []string
	for channel := range s.channels {
		channels = append(channels, channel)
	}

	return &pb.ChannelListResponse{
		Channels: channels,
	}, nil
}

func (s *server) ListChannelUsers(ctx context.Context, in *pb.ChannelUsersRequest) (*pb.ChannelUsersResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	users := make([]string, 0)
	if channel, exists := s.channels[in.Channel]; exists {
		for user := range channel {
			users = append(users, user)
		}
	}
	return &pb.ChannelUsersResponse{
		Users: users,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterChatServiceServer(grpcServer, newServer())

	log.Printf("server listening at %v", lis.Addr())
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("Failed to serve : %v", err)
	}
}
