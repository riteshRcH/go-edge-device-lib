package pstoreds

import (
	"bytes"
	"context"
	"encoding/gob"

	pool "github.com/riteshRcH/go-edge-device-lib/buffer-pool"
	"github.com/riteshRcH/go-edge-device-lib/core/peer"
	pstore "github.com/riteshRcH/go-edge-device-lib/core/peerstore"

	"github.com/riteshRcH/go-edge-device-lib/base32"
	ds "github.com/riteshRcH/go-edge-device-lib/datastore"
	"github.com/riteshRcH/go-edge-device-lib/datastore/query"
)

// Metadata is stored under the following db key pattern:
// /peers/metadata/<b32 peer id no padding>/<key>
var pmBase = ds.NewKey("/peers/metadata")

type dsPeerMetadata struct {
	ds ds.Datastore
}

var _ pstore.PeerMetadata = (*dsPeerMetadata)(nil)

func init() {
	// Gob registers basic types by default.
	//
	// Register complex types used by the peerstore itself.
	gob.Register(make(map[string]struct{}))
}

// NewPeerMetadata creates a metadata store backed by a persistent db. It uses gob for serialisation.
//
// See `init()` to learn which types are registered by default. Modules wishing to store
// values of other types will need to `gob.Register()` them explicitly, or else callers
// will receive runtime errors.
func NewPeerMetadata(_ context.Context, store ds.Datastore, _ Options) (*dsPeerMetadata, error) {
	return &dsPeerMetadata{store}, nil
}

func (pm *dsPeerMetadata) Get(p peer.ID, key string) (interface{}, error) {
	k := pmBase.ChildString(base32.RawStdEncoding.EncodeToString([]byte(p))).ChildString(key)
	value, err := pm.ds.Get(context.TODO(), k)
	if err != nil {
		if err == ds.ErrNotFound {
			err = pstore.ErrNotFound
		}
		return nil, err
	}

	var res interface{}
	if err := gob.NewDecoder(bytes.NewReader(value)).Decode(&res); err != nil {
		return nil, err
	}
	return res, nil
}

func (pm *dsPeerMetadata) Put(p peer.ID, key string, val interface{}) error {
	k := pmBase.ChildString(base32.RawStdEncoding.EncodeToString([]byte(p))).ChildString(key)
	var buf pool.Buffer
	if err := gob.NewEncoder(&buf).Encode(&val); err != nil {
		return err
	}
	return pm.ds.Put(context.TODO(), k, buf.Bytes())
}

func (pm *dsPeerMetadata) RemovePeer(p peer.ID) {
	result, err := pm.ds.Query(context.TODO(), query.Query{
		Prefix:   pmBase.ChildString(base32.RawStdEncoding.EncodeToString([]byte(p))).String(),
		KeysOnly: true,
	})
	if err != nil {
		log.Warnw("querying datastore when removing peer failed", "peer", p, "error", err)
		return
	}
	for entry := range result.Next() {
		pm.ds.Delete(context.TODO(), ds.NewKey(entry.Key))
	}
}
