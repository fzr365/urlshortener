package shortcode

import "math/rand"

//实现GenerateShortCode() string 接口

type ShortCode struct{
	len int
}

func NewShortCode(length int)*ShortCode{
	return &ShortCode{
		len:length,
	}
}


const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func (s *ShortCode) GenerateShortCode() string{
	length:=len(chars)
	result:=make([]byte, s.len)

	for i:=0;i<s.len;i++{
		//索引0-n-1
		result[i]=chars[rand.Intn(length)]
	}
	return string(result)
}