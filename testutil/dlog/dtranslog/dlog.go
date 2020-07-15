package dtranslog

import (
	"github.com/filecoin-project/go-data-transfer/testutil/util"
	"go.uber.org/zap"
)

var L *zap.Logger

func init() {
	L = util.GetXDebugLog("data-trans")
}
