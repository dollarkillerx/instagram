package middlewares

import (
	"context"
	"fmt"

	"github.com/dollarkillerx/graphql_template/internal/utils"
	"runtime/debug"
)

// RecoverFunc ...
func RecoverFunc(ctx context.Context, err interface{}) (userMessage error) {
	utils.Logger.Errorf("[ReqID]:%s\n[Message]:%+v\n[Panic]:%s", utils.GetRequestIdByContext(ctx), err, string(debug.Stack()))
	return fmt.Errorf("Panic: %+v", err)
}
