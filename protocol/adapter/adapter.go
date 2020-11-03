package adapter

import (
	"errors"
	"git.zgwit.com/zgwit/iot-admin/base"
	"sync"
)

type CodeMap map[string]uint8

type Adapter interface {
	Name() string
	Version() string

	Attach(link base.Link)

	Read(slave uint8, code uint8, offset uint16, size uint16) ([]byte, error)
	Write(slave uint8, code uint8, offset uint16, buf []byte) error
}

type Factory func(opts string) (Adapter, error)

type adapter struct {
	Name    string           `json:"name"`
	Codes   CodeMap `json:"codes"`
	factory Factory
}

//可以改为普通map
var protocols sync.Map //adapter

func GetProtocols() []adapter {
	ps := make([]adapter, 0)
	protocols.Range(func(key, value interface{}) bool {
		ps = append(ps, value.(adapter))
		return true
	})
	return ps
}

func CreateAdapter(name string, opts string) (Adapter, error) {
	if v, ok := protocols.Load(name); ok && v != nil {
		p := v.(*adapter)
		return p.factory(opts)
	}
	return nil, errors.New("找不到协议")
}

func RegisterAdapter(name string, codes CodeMap, factory Factory) {
	protocols.Store(name, &adapter{
		Name:    name,
		Codes:   codes,
		factory: factory,
	})
}
