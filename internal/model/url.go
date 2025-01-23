package model
import "time"
//数据结构模型（长URL和短URL）

type CreateURLRequest struct {
	OriginalURL string `json:"original_url" validate:"required,url"`
	//4-10位，字母数字
	CustomCode  string `json:"custom_code,omitempty" validate:"omitempty,min=4,max=10,alphanum"`
	//1-365天
	Duration    *int    `json:"duration,omitempty",validate:"omitempty,min=1,max=365"`
}

type CreateURLResponse struct {
	ShortURL string `json:"short_url"`
	ExpiredAt time.Time `json:"expired_at"`
}