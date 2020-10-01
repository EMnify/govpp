// Code generated by GoVPP's binapi-generator. DO NOT EDIT.

package vxlan_gbp

import (
	"context"
	"fmt"
	api "git.fd.io/govpp.git/api"
	vpe "git.fd.io/govpp.git/internal/testbinapi/binapi2001/vpe"
	"io"
)

// RPCService defines RPC service  vxlan_gbp.
type RPCService interface {
	SwInterfaceSetVxlanGbpBypass(ctx context.Context, in *SwInterfaceSetVxlanGbpBypass) (*SwInterfaceSetVxlanGbpBypassReply, error)
	VxlanGbpTunnelAddDel(ctx context.Context, in *VxlanGbpTunnelAddDel) (*VxlanGbpTunnelAddDelReply, error)
	VxlanGbpTunnelDump(ctx context.Context, in *VxlanGbpTunnelDump) (RPCService_VxlanGbpTunnelDumpClient, error)
}

type serviceClient struct {
	conn api.Connection
}

func NewServiceClient(conn api.Connection) RPCService {
	return &serviceClient{conn}
}

func (c *serviceClient) SwInterfaceSetVxlanGbpBypass(ctx context.Context, in *SwInterfaceSetVxlanGbpBypass) (*SwInterfaceSetVxlanGbpBypassReply, error) {
	out := new(SwInterfaceSetVxlanGbpBypassReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) VxlanGbpTunnelAddDel(ctx context.Context, in *VxlanGbpTunnelAddDel) (*VxlanGbpTunnelAddDelReply, error) {
	out := new(VxlanGbpTunnelAddDelReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) VxlanGbpTunnelDump(ctx context.Context, in *VxlanGbpTunnelDump) (RPCService_VxlanGbpTunnelDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_VxlanGbpTunnelDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&vpe.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_VxlanGbpTunnelDumpClient interface {
	Recv() (*VxlanGbpTunnelDetails, error)
	api.Stream
}

type serviceClient_VxlanGbpTunnelDumpClient struct {
	api.Stream
}

func (c *serviceClient_VxlanGbpTunnelDumpClient) Recv() (*VxlanGbpTunnelDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *VxlanGbpTunnelDetails:
		return m, nil
	case *vpe.ControlPingReply:
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}
