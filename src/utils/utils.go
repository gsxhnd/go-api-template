package utils

import (
	"github.com/google/wire"
)

var UtilsSet = wire.NewSet(
	NewLogger, NewTracer, NewConfig,
)
var version = ""
