package middlewares

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/dollarkillerx/graphql_template/internal/pkg/errs"
	"github.com/dollarkillerx/graphql_template/internal/utils"
	"github.com/pkg/errors"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// MiddleError 全局錯誤處理
func MiddleError(ctx context.Context, err error) *gqlerror.Error {
	utils.Logger.Errorf("[ReqID]:%s\n[Message]:%+v", utils.GetRequestIdByContext(ctx), err)

	var gqlErr *errs.Error
	if !errors.As(err, &gqlErr) {
		return gqlerror.WrapPath(graphql.GetPath(ctx), errs.SystemError(err))
	}

	return gqlerror.WrapPath(graphql.GetPath(ctx), err)
}
