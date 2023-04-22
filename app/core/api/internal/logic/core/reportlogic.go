package core

import (
	"blogs/app/core/api/internal/svc"
	"blogs/app/core/api/internal/types"
	"blogs/common/report"
	"context"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReportLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReportLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReportLogic {
	return &ReportLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReportLogic) Report(r *http.Request) (resp *types.ReportResponse, err error) {
	report.Report(r)
	return
}
