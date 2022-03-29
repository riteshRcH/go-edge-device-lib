package sync

import (
	"testing"

	ds "github.com/riteshRcH/go-edge-device-lib/datastore"
	dstest "github.com/riteshRcH/go-edge-device-lib/datastore/test"
)

func TestSync(t *testing.T) {
	dstest.SubtestAll(t, MutexWrap(ds.NewMapDatastore()))
}
