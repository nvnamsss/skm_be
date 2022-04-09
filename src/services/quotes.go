package services

import (
	"context"
	"famous-quote/dtos"
	"famous-quote/errors"
	"famous-quote/logger"
	"famous-quote/models"
	"famous-quote/repositories"
	"net/http"

	"github.com/jinzhu/copier"
)

type QuotesService interface {
	Get(ctx context.Context) (*dtos.GetQuotesResponse, error)
	Create(ctx context.Context, req *dtos.CreateQuotesRequest) (*dtos.CreateQuotesResponse, error)
	Like(ctx context.Context, req *dtos.LikeQuotesRequest) (*dtos.LikeQuotesResponse, error)
}

type quotesService struct {
	quotesRepository repositories.QuotesRepository
}

func (s *quotesService) Get(ctx context.Context) (*dtos.GetQuotesResponse, error) {
	quotes, err := s.quotesRepository.GetRandom(ctx)
	if err != nil {
		logger.Context(ctx).Errorf("get random quotes error: %v", err)
		return nil, errors.New(errors.ErrInternalServer, err.Error())
	}
	var data dtos.QuotesData
	_ = copier.Copy(&data, quotes)
	logger.Infof("hi mom")

	return &dtos.GetQuotesResponse{
		Meta: dtos.Meta{
			Code:    http.StatusOK,
			Message: "OK",
		},
		Data: &data,
	}, nil
}

func (s *quotesService) Create(ctx context.Context, req *dtos.CreateQuotesRequest) (*dtos.CreateQuotesResponse, error) {
	var (
		quotes *models.Quotes = &models.Quotes{}
	)
	_ = copier.Copy(quotes, req)
	if err := s.quotesRepository.Create(ctx, quotes); err != nil {
		logger.Context(ctx).Errorf("create quotes error: %v", err)
		return nil, errors.New(errors.ErrInternalServer, err.Error())
	}

	return &dtos.CreateQuotesResponse{
		Meta: dtos.Meta{
			Code:    http.StatusOK,
			Message: "OK",
		},
	}, nil
}

func (s *quotesService) Like(ctx context.Context, req *dtos.LikeQuotesRequest) (*dtos.LikeQuotesResponse, error) {
	var (
		incr int64 = 1
	)

	if req.Negative {
		incr = -1
	}
	logger.Infof("req: %v", req)
	like, err := s.quotesRepository.Like(ctx, req.ID, incr)
	if err != nil {
		logger.Context(ctx).Errorf("like got error: %v", err)
		return nil, errors.New(errors.ErrInternalServer, err.Error())
	}

	return &dtos.LikeQuotesResponse{
		Meta: dtos.Meta{
			Code:    http.StatusOK,
			Message: "OK",
		},
		Data: &dtos.LikeQuotesData{
			Like: like,
		},
	}, nil
}

func NewQuotesService(quotesRepository repositories.QuotesRepository) QuotesService {
	return &quotesService{
		quotesRepository: quotesRepository,
	}
}
