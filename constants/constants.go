package constants

import (
	"time"

	"github.com/marianina8/timetravel/utils"
)

var (
	Now           = utils.FormatDestination(time.Now())
	LastDeparture = utils.FormatDestination(time.Now().Add(-4 * time.Minute))
)
