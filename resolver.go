package socks5

import (
	"net"

	"golang.org/x/net/context"
)

// NameResolver is used to implement custom name resolution
type NameResolver interface {
	Resolve(ctx context.Context, name string) (context.Context, net.IP, error)
}

// DNSResolver uses the system DNS to resolve host names
type DNSResolver struct{}

func (d DNSResolver) Resolve(ctx context.Context, name string) (context.Context, net.IP, error) {
     r := &net.Resolver{
        PreferGo: true,
        Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
            d := net.Dialer{}
            return d.DialContext(ctx, "udp", "8.8.8.8:53")
        },
    }

    // Chỉ lấy IPv4
    ips, err := r.LookupIP(ctx, "ip4", name)
    if err != nil {
        return ctx, nil, err
    }
    if len(ips) > 0 {
        return ctx, ips[0], nil
    }

    return ctx, nil, net.ErrClosed
}
