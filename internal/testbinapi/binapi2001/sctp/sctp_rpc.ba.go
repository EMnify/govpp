// Code generated by GoVPP's binapi-generator. DO NOT EDIT.

package sctp

import (
	"context"
	api "git.fd.io/govpp.git/api"
)

// RPCService defines RPC service  sctp.
type RPCService interface {
	SctpAddSrcDstConnection(ctx context.Context, in *SctpAddSrcDstConnection) (*SctpAddSrcDstConnectionReply, error)
	SctpConfig(ctx context.Context, in *SctpConfig) (*SctpConfigReply, error)
	SctpDelSrcDstConnection(ctx context.Context, in *SctpDelSrcDstConnection) (*SctpDelSrcDstConnectionReply, error)
}

type serviceClient struct {
	conn api.Connection
}

func NewServiceClient(conn api.Connection) RPCService {
	return &serviceClient{conn}
}

func (c *serviceClient) SctpAddSrcDstConnection(ctx context.Context, in *SctpAddSrcDstConnection) (*SctpAddSrcDstConnectionReply, error) {
	out := new(SctpAddSrcDstConnectionReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) SctpConfig(ctx context.Context, in *SctpConfig) (*SctpConfigReply, error) {
	out := new(SctpConfigReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) SctpDelSrcDstConnection(ctx context.Context, in *SctpDelSrcDstConnection) (*SctpDelSrcDstConnectionReply, error) {
	out := new(SctpDelSrcDstConnectionReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
