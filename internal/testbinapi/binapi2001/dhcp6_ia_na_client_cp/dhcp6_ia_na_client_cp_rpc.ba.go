// Code generated by GoVPP's binapi-generator. DO NOT EDIT.

package dhcp6_ia_na_client_cp

import (
	"context"
	api "git.fd.io/govpp.git/api"
)

// RPCService defines RPC service  dhcp6_ia_na_client_cp.
type RPCService interface {
	DHCP6ClientEnableDisable(ctx context.Context, in *DHCP6ClientEnableDisable) (*DHCP6ClientEnableDisableReply, error)
}

type serviceClient struct {
	conn api.Connection
}

func NewServiceClient(conn api.Connection) RPCService {
	return &serviceClient{conn}
}

func (c *serviceClient) DHCP6ClientEnableDisable(ctx context.Context, in *DHCP6ClientEnableDisable) (*DHCP6ClientEnableDisableReply, error) {
	out := new(DHCP6ClientEnableDisableReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
