// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	core "blogs/app/core/api/internal/handler/core"
	"blogs/app/core/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/",
				Handler: core.ApiHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/photos/albums",
				Handler: core.GetPhotoAlbumsHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/report",
				Handler: core.ReportHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/categories/all",
				Handler: core.CategoryHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/articles/all",
				Handler: core.GetAllArticlesHandler(serverCtx),
			},
		},
	)
}
