// Code generated by GoVPP's binapi-generator. DO NOT EDIT.

package ipsec

import (
	"context"
	"fmt"
	"io"

	api "git.fd.io/govpp.git/api"
	vpe "git.fd.io/govpp.git/binapi/vpe"
)

// RPCService defines RPC service ipsec.
type RPCService interface {
	IpsecBackendDump(ctx context.Context, in *IpsecBackendDump) (RPCService_IpsecBackendDumpClient, error)
	IpsecInterfaceAddDelSpd(ctx context.Context, in *IpsecInterfaceAddDelSpd) (*IpsecInterfaceAddDelSpdReply, error)
	IpsecSaDump(ctx context.Context, in *IpsecSaDump) (RPCService_IpsecSaDumpClient, error)
	IpsecSadEntryAddDel(ctx context.Context, in *IpsecSadEntryAddDel) (*IpsecSadEntryAddDelReply, error)
	IpsecSelectBackend(ctx context.Context, in *IpsecSelectBackend) (*IpsecSelectBackendReply, error)
	IpsecSpdAddDel(ctx context.Context, in *IpsecSpdAddDel) (*IpsecSpdAddDelReply, error)
	IpsecSpdDump(ctx context.Context, in *IpsecSpdDump) (RPCService_IpsecSpdDumpClient, error)
	IpsecSpdEntryAddDel(ctx context.Context, in *IpsecSpdEntryAddDel) (*IpsecSpdEntryAddDelReply, error)
	IpsecSpdInterfaceDump(ctx context.Context, in *IpsecSpdInterfaceDump) (RPCService_IpsecSpdInterfaceDumpClient, error)
	IpsecSpdsDump(ctx context.Context, in *IpsecSpdsDump) (RPCService_IpsecSpdsDumpClient, error)
	IpsecTunnelIfAddDel(ctx context.Context, in *IpsecTunnelIfAddDel) (*IpsecTunnelIfAddDelReply, error)
	IpsecTunnelIfSetSa(ctx context.Context, in *IpsecTunnelIfSetSa) (*IpsecTunnelIfSetSaReply, error)
	IpsecTunnelProtectDel(ctx context.Context, in *IpsecTunnelProtectDel) (*IpsecTunnelProtectDelReply, error)
	IpsecTunnelProtectDump(ctx context.Context, in *IpsecTunnelProtectDump) (RPCService_IpsecTunnelProtectDumpClient, error)
	IpsecTunnelProtectUpdate(ctx context.Context, in *IpsecTunnelProtectUpdate) (*IpsecTunnelProtectUpdateReply, error)
}

type serviceClient struct {
	conn api.Connection
}

func NewServiceClient(conn api.Connection) RPCService {
	return &serviceClient{conn}
}

func (c *serviceClient) IpsecBackendDump(ctx context.Context, in *IpsecBackendDump) (RPCService_IpsecBackendDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_IpsecBackendDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&vpe.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_IpsecBackendDumpClient interface {
	Recv() (*IpsecBackendDetails, error)
	api.Stream
}

type serviceClient_IpsecBackendDumpClient struct {
	api.Stream
}

func (c *serviceClient_IpsecBackendDumpClient) Recv() (*IpsecBackendDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *IpsecBackendDetails:
		return m, nil
	case *vpe.ControlPingReply:
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}

func (c *serviceClient) IpsecInterfaceAddDelSpd(ctx context.Context, in *IpsecInterfaceAddDelSpd) (*IpsecInterfaceAddDelSpdReply, error) {
	out := new(IpsecInterfaceAddDelSpdReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) IpsecSaDump(ctx context.Context, in *IpsecSaDump) (RPCService_IpsecSaDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_IpsecSaDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&vpe.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_IpsecSaDumpClient interface {
	Recv() (*IpsecSaDetails, error)
	api.Stream
}

type serviceClient_IpsecSaDumpClient struct {
	api.Stream
}

func (c *serviceClient_IpsecSaDumpClient) Recv() (*IpsecSaDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *IpsecSaDetails:
		return m, nil
	case *vpe.ControlPingReply:
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}

func (c *serviceClient) IpsecSadEntryAddDel(ctx context.Context, in *IpsecSadEntryAddDel) (*IpsecSadEntryAddDelReply, error) {
	out := new(IpsecSadEntryAddDelReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) IpsecSelectBackend(ctx context.Context, in *IpsecSelectBackend) (*IpsecSelectBackendReply, error) {
	out := new(IpsecSelectBackendReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) IpsecSpdAddDel(ctx context.Context, in *IpsecSpdAddDel) (*IpsecSpdAddDelReply, error) {
	out := new(IpsecSpdAddDelReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) IpsecSpdDump(ctx context.Context, in *IpsecSpdDump) (RPCService_IpsecSpdDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_IpsecSpdDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&vpe.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_IpsecSpdDumpClient interface {
	Recv() (*IpsecSpdDetails, error)
	api.Stream
}

type serviceClient_IpsecSpdDumpClient struct {
	api.Stream
}

func (c *serviceClient_IpsecSpdDumpClient) Recv() (*IpsecSpdDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *IpsecSpdDetails:
		return m, nil
	case *vpe.ControlPingReply:
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}

func (c *serviceClient) IpsecSpdEntryAddDel(ctx context.Context, in *IpsecSpdEntryAddDel) (*IpsecSpdEntryAddDelReply, error) {
	out := new(IpsecSpdEntryAddDelReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) IpsecSpdInterfaceDump(ctx context.Context, in *IpsecSpdInterfaceDump) (RPCService_IpsecSpdInterfaceDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_IpsecSpdInterfaceDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&vpe.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_IpsecSpdInterfaceDumpClient interface {
	Recv() (*IpsecSpdInterfaceDetails, error)
	api.Stream
}

type serviceClient_IpsecSpdInterfaceDumpClient struct {
	api.Stream
}

func (c *serviceClient_IpsecSpdInterfaceDumpClient) Recv() (*IpsecSpdInterfaceDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *IpsecSpdInterfaceDetails:
		return m, nil
	case *vpe.ControlPingReply:
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}

func (c *serviceClient) IpsecSpdsDump(ctx context.Context, in *IpsecSpdsDump) (RPCService_IpsecSpdsDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_IpsecSpdsDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&vpe.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_IpsecSpdsDumpClient interface {
	Recv() (*IpsecSpdsDetails, error)
	api.Stream
}

type serviceClient_IpsecSpdsDumpClient struct {
	api.Stream
}

func (c *serviceClient_IpsecSpdsDumpClient) Recv() (*IpsecSpdsDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *IpsecSpdsDetails:
		return m, nil
	case *vpe.ControlPingReply:
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}

func (c *serviceClient) IpsecTunnelIfAddDel(ctx context.Context, in *IpsecTunnelIfAddDel) (*IpsecTunnelIfAddDelReply, error) {
	out := new(IpsecTunnelIfAddDelReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) IpsecTunnelIfSetSa(ctx context.Context, in *IpsecTunnelIfSetSa) (*IpsecTunnelIfSetSaReply, error) {
	out := new(IpsecTunnelIfSetSaReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) IpsecTunnelProtectDel(ctx context.Context, in *IpsecTunnelProtectDel) (*IpsecTunnelProtectDelReply, error) {
	out := new(IpsecTunnelProtectDelReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) IpsecTunnelProtectDump(ctx context.Context, in *IpsecTunnelProtectDump) (RPCService_IpsecTunnelProtectDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_IpsecTunnelProtectDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&vpe.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_IpsecTunnelProtectDumpClient interface {
	Recv() (*IpsecTunnelProtectDetails, error)
	api.Stream
}

type serviceClient_IpsecTunnelProtectDumpClient struct {
	api.Stream
}

func (c *serviceClient_IpsecTunnelProtectDumpClient) Recv() (*IpsecTunnelProtectDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *IpsecTunnelProtectDetails:
		return m, nil
	case *vpe.ControlPingReply:
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}

func (c *serviceClient) IpsecTunnelProtectUpdate(ctx context.Context, in *IpsecTunnelProtectUpdate) (*IpsecTunnelProtectUpdateReply, error) {
	out := new(IpsecTunnelProtectUpdateReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}
