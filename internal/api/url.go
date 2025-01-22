package api

import (
	"net/http"

	"github.com/fzr365/urlshortener/internal/model"
	"github.com/labstack/echo/v4"
)

//需求1： 长转短，POST /api/url original_url,custom_code,duration  ,--> short_url,expired_at
//需求2： 重定向，GET /api/url/:short_code --> original_url,expired_at
type URLHandler struct {
	OriginalUrl string `json:"original_url"`
	CustomCode  string `json:"custom_code"`
	Duration    int    `json:"duration"`
	ShortUrl    string `json:"short_url"`
	ExpiredAt   string `json:"expired_at"`
}

//需求1：
func (h *URLHandler) CreateURL(c echo.Context) error {
	//把数据提取
    var req model.CreateURLRequest
	if err:=c.Bind(&req);err!=nil {
		//400 bad request
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	//验证数据的格式
	

	//调用业务函数

	//返回结果
	return c.JSON(200, h)
}