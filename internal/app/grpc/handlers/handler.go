package ghandlers

import (
	"context"
	"errors"
	"github.com/nglmq/ozon-test/internal/app/service"
	"github.com/nglmq/ozon-test/internal/storage"
	"github.com/nglmq/ozon-test/pkg/proto"
)

type ShortenerServer struct {
	proto.UnimplementedShortenerServer
	service *service.URLService
}

func (s *ShortenerServer) Shorten(ctx context.Context, req *proto.URLRequest) (*proto.URLResponse, error) {
	urlResponse, err := s.service.ShortenURL(ctx, req.GetUrl())
	if err != nil {
		return nil, err
	}
	return &proto.URLResponse{Short: urlResponse.Short}, nil
}

func (s *ShortenerServer) GetOriginal(ctx context.Context, req *proto.ShortURLRequest) (*proto.OriginalURLResponse, error) {
	originalURL, err := s.service.GetOriginalURL(ctx, req.GetShort())
	if err != nil {
		return nil, err
	}
	return &proto.OriginalURLResponse{Original: originalURL.Original}, nil
}
