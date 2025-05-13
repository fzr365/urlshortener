package service

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/fzr365/urlshortener/internal/model"
	"github.com/fzr365/urlshortener/internal/repo"
)

// shortcode生成的接口
type ShortCodeGenerator interface {
	GenerateShortCode() string
}

// 抽象缓存接口
type Cacher interface {
	SetURL(ctx context.Context, url repo.Url) error
	GetURL(ctx context.Context, shortCode string) (*repo.Url, error)
}

type URLService struct {
	querier            repo.Querier
	shortCodeGenerator ShortCodeGenerator
	defaultDuration    time.Duration
	cache              Cacher
	baseURL            string
}

//对外服务层划分
// NewURLService 构造函数
func NewURLService(db *sql.DB, shortCodeGenerator ShortCodeGenerator, duration time.Duration, cache Cacher, baseURL string) *URLService {
	return &URLService{
		querier:            repo.New(db),
		shortCodeGenerator: shortCodeGenerator,
		defaultDuration:    duration,
		cache:              cache,
		baseURL:            baseURL,
	}	
}



func (s *URLService) CreateURL(ctx context.Context, req model.CreateURLRequest) (*model.CreateURLResponse, error) {
	var shortCode string
	var isCustom bool
	var expiredAt time.Time

	//判断是否是自定义的shortcode
	//存在的情况
	if req.CustomCode != "" {
		//查询数据库
		isAvailable, err := s.querier.IsShortCodeAvailable(ctx, req.CustomCode)
		if err != nil {
			return nil, err
		}
		//没有错误
		//已经存在于数据库中
		if !isAvailable {
			return nil, errors.New("别名已存在")
		}
		//赋值
		shortCode = req.CustomCode
		isCustom = true
	} else {
		code, err := s.getShortCode(ctx, 0)
		if err != nil {
			return nil, err
		}
		shortCode = code
	}


	
	if req.Duration == nil {
		expiredAt = time.Now().Add(s.defaultDuration)
	} else {
		//涉及到指针的解引用
		expiredAt = time.Now().Add(time.Hour * time.Duration(*req.Duration))
	}

	//存入数据库
	err := s.querier.CreateURL(ctx, repo.CreateURLParams{
		OriginalUrl: req.OriginalURL,
		ShortCode:   shortCode,
		IsCustom:    isCustom,
		ExpiredAt:   expiredAt,
	})

	if err != nil {
		return nil, err
	}

	url, err := s.querier.GetInsertedURL(ctx)
	if err != nil {
		return nil, err
	}

	//存入缓存
	if err := s.cache.SetURL(ctx, url); err != nil {
		return nil, err
	}
	//返回结果
	return &model.CreateURLResponse{
		ShortURL:  s.baseURL + "/" + url.ShortCode,
		ExpiredAt: url.ExpiredAt,
	}, nil
}

// GetURL(ctx context.Context, shortCode string) (string, error)
func (s *URLService) GetURL(ctx context.Context, shortCode string) (string, error) {
	//先访问缓存
	url, err := s.cache.GetURL(ctx, shortCode)
	if err != nil {
		return "", err
	}
	if url != nil {
		return url.OriginalUrl, nil
	}
	//访问数据库
	url2, err := s.querier.GetURLByShortCode(ctx, shortCode)
	if err != nil {
		return "", err
	}

	//存入缓存
	if err := s.cache.SetURL(ctx, url2); err != nil {
		return "", err
	}

	return url2.OriginalUrl, nil
}

// 重试机制，不存在shortcode，生成一次，插入数据库，再查询，直到成功
func (s *URLService) getShortCode(ctx context.Context, n int) (string, error) {
	if n > 5 {
		return "", errors.New("重试次数过多")
	}
	shortCode := s.shortCodeGenerator.GenerateShortCode()

	isAvailable, err := s.querier.IsShortCodeAvailable(ctx, shortCode)
	if err != nil {
		return "", err
	}

	if isAvailable {
		return shortCode, nil
	}

	return s.getShortCode(ctx, n+1)
}
