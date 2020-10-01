// Code generated by GoVPP's binapi-generator. DO NOT EDIT.

package gbp

import (
	"context"
	"fmt"
	"io"

	api "git.fd.io/govpp.git/api"
	vpe "git.fd.io/govpp.git/binapi/vpe"
)

// RPCService defines RPC service  gbp.
type RPCService interface {
	GbpBridgeDomainAdd(ctx context.Context, in *GbpBridgeDomainAdd) (*GbpBridgeDomainAddReply, error)
	GbpBridgeDomainDel(ctx context.Context, in *GbpBridgeDomainDel) (*GbpBridgeDomainDelReply, error)
	GbpBridgeDomainDump(ctx context.Context, in *GbpBridgeDomainDump) (RPCService_GbpBridgeDomainDumpClient, error)
	GbpContractAddDel(ctx context.Context, in *GbpContractAddDel) (*GbpContractAddDelReply, error)
	GbpContractDump(ctx context.Context, in *GbpContractDump) (RPCService_GbpContractDumpClient, error)
	GbpEndpointAdd(ctx context.Context, in *GbpEndpointAdd) (*GbpEndpointAddReply, error)
	GbpEndpointDel(ctx context.Context, in *GbpEndpointDel) (*GbpEndpointDelReply, error)
	GbpEndpointDump(ctx context.Context, in *GbpEndpointDump) (RPCService_GbpEndpointDumpClient, error)
	GbpEndpointGroupAdd(ctx context.Context, in *GbpEndpointGroupAdd) (*GbpEndpointGroupAddReply, error)
	GbpEndpointGroupDel(ctx context.Context, in *GbpEndpointGroupDel) (*GbpEndpointGroupDelReply, error)
	GbpEndpointGroupDump(ctx context.Context, in *GbpEndpointGroupDump) (RPCService_GbpEndpointGroupDumpClient, error)
	GbpExtItfAddDel(ctx context.Context, in *GbpExtItfAddDel) (*GbpExtItfAddDelReply, error)
	GbpExtItfDump(ctx context.Context, in *GbpExtItfDump) (RPCService_GbpExtItfDumpClient, error)
	GbpRecircAddDel(ctx context.Context, in *GbpRecircAddDel) (*GbpRecircAddDelReply, error)
	GbpRecircDump(ctx context.Context, in *GbpRecircDump) (RPCService_GbpRecircDumpClient, error)
	GbpRouteDomainAdd(ctx context.Context, in *GbpRouteDomainAdd) (*GbpRouteDomainAddReply, error)
	GbpRouteDomainDel(ctx context.Context, in *GbpRouteDomainDel) (*GbpRouteDomainDelReply, error)
	GbpRouteDomainDump(ctx context.Context, in *GbpRouteDomainDump) (RPCService_GbpRouteDomainDumpClient, error)
	GbpSubnetAddDel(ctx context.Context, in *GbpSubnetAddDel) (*GbpSubnetAddDelReply, error)
	GbpSubnetDump(ctx context.Context, in *GbpSubnetDump) (RPCService_GbpSubnetDumpClient, error)
	GbpVxlanTunnelAdd(ctx context.Context, in *GbpVxlanTunnelAdd) (*GbpVxlanTunnelAddReply, error)
	GbpVxlanTunnelDel(ctx context.Context, in *GbpVxlanTunnelDel) (*GbpVxlanTunnelDelReply, error)
	GbpVxlanTunnelDump(ctx context.Context, in *GbpVxlanTunnelDump) (RPCService_GbpVxlanTunnelDumpClient, error)
}

type serviceClient struct {
	conn api.Connection
}

func NewServiceClient(conn api.Connection) RPCService {
	return &serviceClient{conn}
}

func (c *serviceClient) GbpBridgeDomainAdd(ctx context.Context, in *GbpBridgeDomainAdd) (*GbpBridgeDomainAddReply, error) {
	out := new(GbpBridgeDomainAddReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GbpBridgeDomainDel(ctx context.Context, in *GbpBridgeDomainDel) (*GbpBridgeDomainDelReply, error) {
	out := new(GbpBridgeDomainDelReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GbpBridgeDomainDump(ctx context.Context, in *GbpBridgeDomainDump) (RPCService_GbpBridgeDomainDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_GbpBridgeDomainDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&vpe.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_GbpBridgeDomainDumpClient interface {
	Recv() (*GbpBridgeDomainDetails, error)
	api.Stream
}

type serviceClient_GbpBridgeDomainDumpClient struct {
	api.Stream
}

func (c *serviceClient_GbpBridgeDomainDumpClient) Recv() (*GbpBridgeDomainDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *GbpBridgeDomainDetails:
		return m, nil
	case *vpe.ControlPingReply:
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}

func (c *serviceClient) GbpContractAddDel(ctx context.Context, in *GbpContractAddDel) (*GbpContractAddDelReply, error) {
	out := new(GbpContractAddDelReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GbpContractDump(ctx context.Context, in *GbpContractDump) (RPCService_GbpContractDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_GbpContractDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&vpe.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_GbpContractDumpClient interface {
	Recv() (*GbpContractDetails, error)
	api.Stream
}

type serviceClient_GbpContractDumpClient struct {
	api.Stream
}

func (c *serviceClient_GbpContractDumpClient) Recv() (*GbpContractDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *GbpContractDetails:
		return m, nil
	case *vpe.ControlPingReply:
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}

func (c *serviceClient) GbpEndpointAdd(ctx context.Context, in *GbpEndpointAdd) (*GbpEndpointAddReply, error) {
	out := new(GbpEndpointAddReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GbpEndpointDel(ctx context.Context, in *GbpEndpointDel) (*GbpEndpointDelReply, error) {
	out := new(GbpEndpointDelReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GbpEndpointDump(ctx context.Context, in *GbpEndpointDump) (RPCService_GbpEndpointDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_GbpEndpointDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&vpe.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_GbpEndpointDumpClient interface {
	Recv() (*GbpEndpointDetails, error)
	api.Stream
}

type serviceClient_GbpEndpointDumpClient struct {
	api.Stream
}

func (c *serviceClient_GbpEndpointDumpClient) Recv() (*GbpEndpointDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *GbpEndpointDetails:
		return m, nil
	case *vpe.ControlPingReply:
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}

func (c *serviceClient) GbpEndpointGroupAdd(ctx context.Context, in *GbpEndpointGroupAdd) (*GbpEndpointGroupAddReply, error) {
	out := new(GbpEndpointGroupAddReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GbpEndpointGroupDel(ctx context.Context, in *GbpEndpointGroupDel) (*GbpEndpointGroupDelReply, error) {
	out := new(GbpEndpointGroupDelReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GbpEndpointGroupDump(ctx context.Context, in *GbpEndpointGroupDump) (RPCService_GbpEndpointGroupDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_GbpEndpointGroupDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&vpe.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_GbpEndpointGroupDumpClient interface {
	Recv() (*GbpEndpointGroupDetails, error)
	api.Stream
}

type serviceClient_GbpEndpointGroupDumpClient struct {
	api.Stream
}

func (c *serviceClient_GbpEndpointGroupDumpClient) Recv() (*GbpEndpointGroupDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *GbpEndpointGroupDetails:
		return m, nil
	case *vpe.ControlPingReply:
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}

func (c *serviceClient) GbpExtItfAddDel(ctx context.Context, in *GbpExtItfAddDel) (*GbpExtItfAddDelReply, error) {
	out := new(GbpExtItfAddDelReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GbpExtItfDump(ctx context.Context, in *GbpExtItfDump) (RPCService_GbpExtItfDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_GbpExtItfDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&vpe.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_GbpExtItfDumpClient interface {
	Recv() (*GbpExtItfDetails, error)
	api.Stream
}

type serviceClient_GbpExtItfDumpClient struct {
	api.Stream
}

func (c *serviceClient_GbpExtItfDumpClient) Recv() (*GbpExtItfDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *GbpExtItfDetails:
		return m, nil
	case *vpe.ControlPingReply:
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}

func (c *serviceClient) GbpRecircAddDel(ctx context.Context, in *GbpRecircAddDel) (*GbpRecircAddDelReply, error) {
	out := new(GbpRecircAddDelReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GbpRecircDump(ctx context.Context, in *GbpRecircDump) (RPCService_GbpRecircDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_GbpRecircDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&vpe.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_GbpRecircDumpClient interface {
	Recv() (*GbpRecircDetails, error)
	api.Stream
}

type serviceClient_GbpRecircDumpClient struct {
	api.Stream
}

func (c *serviceClient_GbpRecircDumpClient) Recv() (*GbpRecircDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *GbpRecircDetails:
		return m, nil
	case *vpe.ControlPingReply:
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}

func (c *serviceClient) GbpRouteDomainAdd(ctx context.Context, in *GbpRouteDomainAdd) (*GbpRouteDomainAddReply, error) {
	out := new(GbpRouteDomainAddReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GbpRouteDomainDel(ctx context.Context, in *GbpRouteDomainDel) (*GbpRouteDomainDelReply, error) {
	out := new(GbpRouteDomainDelReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GbpRouteDomainDump(ctx context.Context, in *GbpRouteDomainDump) (RPCService_GbpRouteDomainDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_GbpRouteDomainDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&vpe.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_GbpRouteDomainDumpClient interface {
	Recv() (*GbpRouteDomainDetails, error)
	api.Stream
}

type serviceClient_GbpRouteDomainDumpClient struct {
	api.Stream
}

func (c *serviceClient_GbpRouteDomainDumpClient) Recv() (*GbpRouteDomainDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *GbpRouteDomainDetails:
		return m, nil
	case *vpe.ControlPingReply:
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}

func (c *serviceClient) GbpSubnetAddDel(ctx context.Context, in *GbpSubnetAddDel) (*GbpSubnetAddDelReply, error) {
	out := new(GbpSubnetAddDelReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GbpSubnetDump(ctx context.Context, in *GbpSubnetDump) (RPCService_GbpSubnetDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_GbpSubnetDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&vpe.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_GbpSubnetDumpClient interface {
	Recv() (*GbpSubnetDetails, error)
	api.Stream
}

type serviceClient_GbpSubnetDumpClient struct {
	api.Stream
}

func (c *serviceClient_GbpSubnetDumpClient) Recv() (*GbpSubnetDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *GbpSubnetDetails:
		return m, nil
	case *vpe.ControlPingReply:
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}

func (c *serviceClient) GbpVxlanTunnelAdd(ctx context.Context, in *GbpVxlanTunnelAdd) (*GbpVxlanTunnelAddReply, error) {
	out := new(GbpVxlanTunnelAddReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GbpVxlanTunnelDel(ctx context.Context, in *GbpVxlanTunnelDel) (*GbpVxlanTunnelDelReply, error) {
	out := new(GbpVxlanTunnelDelReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GbpVxlanTunnelDump(ctx context.Context, in *GbpVxlanTunnelDump) (RPCService_GbpVxlanTunnelDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_GbpVxlanTunnelDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&vpe.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_GbpVxlanTunnelDumpClient interface {
	Recv() (*GbpVxlanTunnelDetails, error)
	api.Stream
}

type serviceClient_GbpVxlanTunnelDumpClient struct {
	api.Stream
}

func (c *serviceClient_GbpVxlanTunnelDumpClient) Recv() (*GbpVxlanTunnelDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *GbpVxlanTunnelDetails:
		return m, nil
	case *vpe.ControlPingReply:
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}
