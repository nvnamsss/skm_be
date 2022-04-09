package services

import (
	"context"
	"errors"
	"famous-quote/dtos"
	mocksRepo "famous-quote/mocks/repositories"
	"famous-quote/models"
	"famous-quote/repositories"
	"testing"

	"github.com/stretchr/testify/mock"
)

func TestQuotes_Get(t *testing.T) {
	type fields struct {
		quotesRepository repositories.QuotesRepository
	}

	type args struct {
		ctx context.Context
	}

	var (
		quotesRepository    = &mocksRepo.QuotesRepository{}
		errQuotesRepository = &mocksRepo.QuotesRepository{}
		quotes              = models.Quotes{}
	)

	quotesRepository.On("GetRandom", mock.Anything).Return(&quotes, nil)
	errQuotesRepository.On("GetRandom", mock.Anything).Return(nil, errors.New("error"))
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "good",
			fields: fields{quotesRepository: quotesRepository},
			args:   args{ctx: context.Background()},
		},
		{
			name:    "error",
			fields:  fields{quotesRepository: errQuotesRepository},
			args:    args{ctx: context.Background()},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &quotesService{
				quotesRepository: tt.fields.quotesRepository,
			}
			_, err := s.Get(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr {
				return
			}

		})
	}

}

func TestQuotes_Like(t *testing.T) {
	type fields struct {
		quotesRepository repositories.QuotesRepository
	}

	type args struct {
		ctx context.Context
		req *dtos.LikeQuotesRequest
	}

	var (
		quotesRepository    = &mocksRepo.QuotesRepository{}
		errQuotesRepository = &mocksRepo.QuotesRepository{}

		reqs map[string]*dtos.LikeQuotesRequest = map[string]*dtos.LikeQuotesRequest{
			"like":    {},
			"dislike": {Negative: true},
		}
	)

	quotesRepository.On("Like", mock.Anything, mock.Anything, mock.Anything).Return(int64(1), nil)
	errQuotesRepository.On("Like", mock.Anything, mock.Anything, mock.Anything).Return(int64(0), errors.New("error"))
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "like",
			fields: fields{quotesRepository: quotesRepository},
			args: args{
				ctx: context.Background(),
				req: reqs["like"],
			},
		},
		{
			name:   "dislike",
			fields: fields{quotesRepository: quotesRepository},
			args: args{
				ctx: context.Background(),
				req: reqs["dislike"],
			},
		},
		{
			name:   "error",
			fields: fields{quotesRepository: errQuotesRepository},
			args: args{
				ctx: context.Background(),
				req: &dtos.LikeQuotesRequest{},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &quotesService{
				quotesRepository: tt.fields.quotesRepository,
			}
			_, err := s.Like(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Like() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr {
				return
			}

		})
	}
}
