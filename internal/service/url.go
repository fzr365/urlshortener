package service

import (
	"context"
	"fmt"
	"time"
	"errors"
)	

//shortcode生成的接口
type ShortCodeGenerator interface {
	GenerateShortCode() string
}

//抽象缓存接口
type Cacher interface {
	SetURL(ctx context.Context, url repo.Url) error
}



type URLService struct {
    querier repo.Querier
	shortCodeGenerator ShortCodeGenerator
	defaultDuration time.Duration
	cache Cacher
	baseURL string
}

func(s *URLService) CreateURL(ctx context.Context, req model.CreateURLRequest) (*model.CreateURLResponse, error) {
	var shortCode string
	var isCustom bool
	var expiredAt time.Time

	//判断是否是自定义的shortcode
	//存在的情况
	if req.CustomCode != "" {
		//查询数据库
	    isAvailable,err:= s.querier.IsShortCodeAvailable(ctx, req.CustomCode)
		if err!=nil {
			return nil, err
		}
		//没有错误
		//已经存在于数据库中
		if !isAvailable {
			return nil, fmt.Error("别名已存在")
		}
		//赋值
		shortCode = req.CustomCode
		isCustom = true
	} else {
		code,err:=s.getShortCode(ctx,0)
		if err!=nil {
			return nil, err
		}
		shortCode = code 
	}

    if req.Expiration == nil{
		expiredAt = time.Now().Add(s.defaultDuration)
	} else {
		//涉及到指针的解引用
		expiredAt = time.Now().Add(time.Hour*time.Duration(*req.Duration))
	}

	//存入数据库
    url,err:= s.querier.CreateURL(ctx, repo.CreateURLParams{
		OriginalUrl: req.OriginalURL,
		ShortCode: shortCode,
		IsCustom: isCustom,
		ExpiredAt: expiredAt,
		})

	if err!=nil {
		return nil, err
	}

	//存入缓存 
    if err:=s.cache.SetURL(ctx, url); err!=nil {
		return nil, err
	}
	//返回结果
	return &model.CreateURLResponse{
		ShortURL: s.baseURL + "/" + url.shortCode,
		ExpiredAt: url.ExpiredAt,
		}, nil
}



//重试机制，不存在shortcode，生成一次，插入数据库，再查询，直到成功
func(s *URLService) getShortCode(ctx context.Context,n int) (string, error){
    if n>5{
		return "", errors.New("重试次数过多")
	}
	shortCode:=s.shortCodeGenerator.GenerateShortCode()

	isAvailable,err:=s.querier.IsShortCodeAvailable(ctx, shortCode)
	if err!=nil {
		return "", err
	}

	if isAvailable {
		return shortCode, nil
	}

	return s.getShortCode(ctx,n+1)
}