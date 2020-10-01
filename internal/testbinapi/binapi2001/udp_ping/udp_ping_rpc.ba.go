// Code generated by GoVPP's binapi-generator. DO NOT EDIT.

package udp_ping

import (
	"context"
	api "git.fd.io/govpp.git/api"
)

// RPCService defines RPC service  udp_ping.
type RPCService interface {
	UDPPingAddDel(ctx context.Context, in *UDPPingAddDel) (*UDPPingAddDelReply, error)
	UDPPingExport(ctx context.Context, in *UDPPingExport) (*UDPPingExportReply, error)
}

type serviceClient struct {
	conn api.Connection
}

func NewServiceClient(conn api.Connection) RPCService {
	return &serviceClient{conn}
}

func (c *serviceClient) UDPPingAddDel(ctx context.Context, in *UDPPingAddDel) (*UDPPingAddDelReply, error) {
	out := new(UDPPingAddDelReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) UDPPingExport(ctx context.Context, in *UDPPingExport) (*UDPPingExportReply, error) {
	out := new(UDPPingExportReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
