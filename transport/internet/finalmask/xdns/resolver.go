package xdns

import (
	"errors"
	"net"

	"github.com/xtls/xray-core/common/serial"
	"github.com/xtls/xray-core/transport/internet/finalmask"
)

type Resolver interface {
	Addr() *net.UDPAddr
	EDNS0() *int32
	Read(p []byte) (int, error)
	Send(p []byte)
	Close()
}

func NewResolver(proto *serial.TypedMessage, dialer *finalmask.Dialer) (Resolver, error) {
	config, err := proto.GetInstance()
	if err != nil {
		return nil, err
	}
	switch v := config.(type) {
	case *TCPResolverProto:
		return NewTCPResolver(v, dialer)
	case *UDPResolverProto:
		return NewUDPResolver(v, dialer)
	default:
		return nil, errors.New("unknown proto")
	}
}

func validateResolverEDNS0(value *int32) error {
	if value != nil && *value != 0 && (*value < 512 || *value > 4096) {
		return errors.New("resolver edns0 must be 0 or between 512 and 4096")
	}
	return nil
}
