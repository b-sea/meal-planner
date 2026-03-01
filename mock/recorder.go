package mock

import (
	"time"

	"github.com/b-sea/go-server/server"
	"github.com/b-sea/meal-planner/mariadb"
)

var _ mariadb.Recorder = (*Recorder)(nil)

type Recorder struct {
	server.NoOpRecorder
}

func (r *Recorder) ObserveMariadbTxDuration(string, time.Duration) {}
