package infrastructure

import (
	"github.com/labstack/echo/v4"
)

// Router ルーティング
type Router struct {
	e  *echo.Echo
	lc *controllers.LinebotController
}

// NewRouter コンストラクタ
func NewRouter(e *echo.Echo, lc *controllers.LinebotController) *Router {
	return &Router{e: e}
}

// Init ルーティング設定
func (r *Router) Init() {
	r.e.POST("/linebot/callback", r.lc.CatchEvents())

	// api := r.e.Group("/api")
	// {
	// 	api.GET("/search", r.ac.Search())

	// 	favorite := api.Group("/favorite")
	// 	{
	// 		favorite.POST("/get", r.ac.GetFavorites())
	// 		favorite.POST("/add", r.ac.AddFavorites())
	// 		favorite.POST("/remove", r.ac.RemoveFavorites())
	// 	}
	// }
}
