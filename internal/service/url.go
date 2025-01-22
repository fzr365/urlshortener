package service


//shortcode生成的接口
type ShortCodeGenerator interface {
	GenerateShortCode() string
}

type URLService struct {
    querier repo.Querier
	shortCodeGenerator ShortCodeGenerator
}

func(s *URLService) CreateURL(ctx context.Context, req model.CreateURLRequest) (*model.CreateURLResponse, error) {
	var shortCode string
	var is_custom bool
	var expiredAt time.Time

	//判断是否是自定义的shortcode
	if req.CustomCode != "" {
		//查询数据库
	    isAvailable,err:= s.querier.IsShortCodeAvailable(ctx, req.CustomCode)
		if err!=nil {
			return nil, err
		}
		//没有错误
		//已经存在于数据库中
		if !isAvailable {
			return nil, fmt.Error("别名已存在(custom_code already exists)")
		}
		//更新别名和自定义标签
		shortCode = req.CustomCode
		is_custom = true
	} else {
		shortCode = s.shortCodeGenerator.GenerateShortCode()
		is_custom = false
	}
	//存入数据库
    s.querier.CreateURL(ctx, repo.CreateURLParams{})

	//存入缓存

	//返回结果
	return &model.CreateURLResponse{}, nil
}