package repositories

import (
	"context"
	"errors"
	"famous-quote/adapters/cache"
	"famous-quote/adapters/database"
	"famous-quote/logger"
	"famous-quote/models"
	"famous-quote/utils"
	"math/rand"
	"time"

	"github.com/vmihailenco/msgpack"
	"gorm.io/gorm"
)

type QuotesRepository interface {
	Get(ctx context.Context, id int64) (*models.Quotes, error)
	GetRandom(ctx context.Context) (*models.Quotes, error)
	Create(ctx context.Context, quotes *models.Quotes) error
	Like(ctx context.Context, id int64, incr int64) (int64, error)
}

type quotesRepository struct {
	db    database.DBAdapter
	cache cache.CacheAdapter
}

func (r *quotesRepository) Get(ctx context.Context, id int64) (*models.Quotes, error) {
	var (
		quotes *models.Quotes = &models.Quotes{}
	)

	v, err := r.cache.HGet(ctx, utils.QuotesKey(), utils.GetKey(id))
	if err != nil {
		if err := r.db.Gormer().WithContext(ctx).Find(quotes, id).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return nil, nil
			}
			return nil, err
		}
		r.cache.HSet(ctx, utils.QuotesKey(), utils.GetKey(id), quotes)
	}

	if err := msgpack.Unmarshal([]byte(v), quotes); err != nil {
		logger.Context(ctx).Errorf("unmarshal msgpack error: %", err)
	}

	return quotes, nil
}

func (r *quotesRepository) GetRandom(ctx context.Context) (*models.Quotes, error) {
	id, err := r.cache.GetInt64(ctx, "today")
	logger.Infof("%v", id)
	if err != nil {
		logger.Context(ctx).Errorf("get today quotes error: %v", err)
		q := &models.QuotesOfTheDay{}
		if err := r.db.Gormer().WithContext(ctx).Last(q).Error; err != nil {
			if err != gorm.ErrRecordNotFound {
				logger.Context(ctx).Errorf("aaaa")
				return nil, err
			}
		}

		now := time.Now()
		last := time.Date(q.CreatedAt.Year(), q.CreatedAt.Month(), q.CreatedAt.Day(), 0, 0, 0, 0, q.CreatedAt.Location())
		if now.Sub(last) < time.Hour*24 {
			id = q.QuotesID
		}
	}

	if id <= 0 {
		var count int64
		if err := r.db.Gormer().Model(&models.Quotes{}).Count(&count).Error; err != nil {
			return nil, err
		}

		if count == 0 {
			return nil, errors.New("there are no quotes")
		}
		id = rand.Int63n(count) + 1
		logger.Infof("count: %v, random: %v", count, id)
		now := time.Now()
		nextDay := time.Now().AddDate(0, 0, 1)
		nextDay = time.Date(nextDay.Year(), nextDay.Month(), nextDay.Day(), 0, 0, 0, 0, nextDay.Location())
		r.cache.Set(ctx, "today", id, nextDay.Sub(now))

		if err := r.db.Gormer().WithContext(ctx).Create(&models.QuotesOfTheDay{QuotesID: id}).Error; err != nil {
			return nil, err
		}
	}

	return r.Get(ctx, id)
}

func (r *quotesRepository) Create(ctx context.Context, quotes *models.Quotes) error {
	return r.db.Gormer().WithContext(ctx).Create(quotes).Error
}

func (r *quotesRepository) Like(ctx context.Context, id int64, incr int64) (int64, error) {
	if err := r.db.Gormer().WithContext(ctx).Model(&models.Quotes{ID: id}).Updates(map[string]interface{}{
		"like": gorm.Expr("`like` + ?", incr),
	}).Error; err != nil {
		return 0, err
	}

	q := models.Quotes{ID: id}
	if err := r.db.Gormer().WithContext(ctx).Find(&q).Error; err != nil {
		return 0, err
	}
	r.cache.HSet(ctx, utils.QuotesKey(), utils.GetKey(id), q)
	return q.Like, nil
}

func NewQuotesRepository(db database.DBAdapter, cache cache.CacheAdapter) QuotesRepository {
	return &quotesRepository{
		db:    db,
		cache: cache,
	}
}
