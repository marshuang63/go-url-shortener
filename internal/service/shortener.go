package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/go-redis/redis/v8"
	"github.com/marshuang63/go-url-shortener/proto/urlshortenerpb"
	"github.com/speps/go-hashids"
	"google.golang.org/grpc"
)

var ctx = context.Background()
var redisClient *redis.Client

type URLShortenerServer struct {
	urlshortenerpb.UnimplementedURLShortenerServer // 正確引用 urlshortenerpb
}

func init() {
	redisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}

func (s *URLShortenerServer) ShortenURL(ctx context.Context, req *urlshortenerpb.ShortenRequest) (*urlshortenerpb.ShortenResponse, error) {
	hd := hashids.NewData()
	hd.Salt = "your-secret-salt"
	h, _ := hashids.NewWithData(hd)

	id, _ := h.Encode([]int{len(req.LongUrl)})
	shortURL := fmt.Sprintf("http://short.ly/%s", id)

	redisClient.Set(ctx, shortURL, req.LongUrl, 0)
	return &urlshortenerpb.ShortenResponse{ShortUrl: shortURL}, nil
}

func (s *URLShortenerServer) ExpandURL(ctx context.Context, req *urlshortenerpb.ExpandRequest) (*urlshortenerpb.ExpandResponse, error) {
	longURL, err := redisClient.Get(ctx, req.ShortUrl).Result()
	if err != nil {
		return nil, err
	}
	return &urlshortenerpb.ExpandResponse{LongUrl: longURL}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	urlshortenerpb.RegisterURLShortenerServer(s, &URLShortenerServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	fmt.Println("gRPC Server Running...")
}
