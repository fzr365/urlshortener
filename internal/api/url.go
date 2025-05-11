package api

import (
	"net/http"
    "context"
	"github.com/fzr365/urlshortener/internal/model"
	"github.com/labstack/echo/v4"
)

//需求1： 长转短，POST /api/url original_url,custom_code,duration  ,--> short_url,expired_at
//需求2： 重定向，GET /api/url/:code --> original_url,expired_at

//实现url转换的接口
type URLService interface {
	CreateURL(ctx context.Context, req model.CreateURLRequest) (*model.CreateURLResponse, error)

	GetURL(ctx context.Context, shortCode string) (string, error)
}


type URLHandler struct {
	//传入转换的接口
	urlService URLService
}


//对外函数api层
func NewURLHandler(urlService URLService) *URLHandler {
	return &URLHandler{
		urlService: urlService,
	}	
}




//需求1：post方法
func (h *URLHandler) CreateURL(c echo.Context) error {
	//把数据提取
    var req model.CreateURLRequest
	if err:=c.Bind(&req);err!=nil {
		//500 internal server error
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	//验证数据的格式
	if err:=c.Validate(req);err!=nil {
		//400 bad request
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	//调用业务函数
    resp,err:=h.urlService.CreateURL(c.Request().Context(), req)
	if err!=nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	//返回201结果
	 return c.JSON(http.StatusCreated, resp)
}


//需求2：get方法
//需求2： 重定向，GET /api/url/:shortcode --> original_url,expired_at
func(h*URLHandler) RedirectURL(c echo.Context) error {
	//获取code
	shortCode:=c.Param("code")
	//shortcode-->url调用业务函数
	originalURL,err:=h.urlService.GetURL(c.Request().Context(), shortCode)
	if err!=nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	//永久重定向(浏览器缓存)
	return c.Redirect(http.StatusPermanentRedirect, originalURL)
}


