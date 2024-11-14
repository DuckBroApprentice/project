package routers

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/go-programming-tour-book/blog-service/internal/routers/api/v1"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	/*apiv1 型別 RouterGroup
	type RouterGroup struct {
		Handlers HandlersChain  //type HandlersChain []HandlerFunc  //type HandlerFunc func(*Context)
		basePath string
		engine   *Engine //type Engine struct{...}
		root     bool
	}
	*/
	article := v1.NewArticle()
	tag := v1.NewTag()
	apiv1 := r.Group("/api/v1")
	{
		//apiv1.方法("路徑")
		apiv1.POST("/tags", tag.Create)
		/*
			func (group *RouterGroup) POST(relativePath string, handlers ...HandlerFunc) IRoutes {
			return group.handle(http.MethodPost, relativePath, handlers)
			}
		*/
		apiv1.DELETE("/tags/:id", tag.Delete)
		/*
			func (group *RouterGroup) DELETE(relativePath string, handlers ...HandlerFunc) IRoutes {
			return group.handle(http.MethodDelete, relativePath, handlers)
			}
		*/
		apiv1.PUT("/tags/:id", tag.Update)
		/*
			func (group *RouterGroup) PUT(relativePath string, handlers ...HandlerFunc) IRoutes {
			return group.handle(http.MethodPut, relativePath, handlers)
			}
		*/
		apiv1.PATCH("/tags/:id/state", tag.List)
		/*
			func (group *RouterGroup) PATCH(relativePath string, handlers ...HandlerFunc) IRoutes {
			return group.handle(http.MethodPatch, relativePath, handlers)
			}
		*/
		apiv1.GET("/tags")
		/*
			func (group *RouterGroup) GET(relativePath string, handlers ...HandlerFunc) IRoutes {
			return group.handle(http.MethodGet, relativePath, handlers)
			}
		*/

		apiv1.POST("/articles", article.Create)
		apiv1.DELETE("/articles/:id", article.Delete)
		apiv1.PUT("/articles/:id", article.Update)
		apiv1.PATCH("/articles/:id/state", article.Update)
		apiv1.GET("/articles/:id", article.Get)
		apiv1.GET("/articles", article.List)
	}

	return r
}
