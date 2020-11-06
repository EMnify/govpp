// Code generated by GoVPP's binapi-generator. DO NOT EDIT.

package teib

import (
	"context"
	"fmt"
	"io"

	api "git.fd.io/govpp.git/api"
	vpe "git.fd.io/govpp.git/binapi/vpe"
)

// RPCService defines RPC service teib.
type RPCService interface {
	TeibDump(ctx context.Context, in *TeibDump) (RPCService_TeibDumpClient, error)
	TeibEntryAddDel(ctx context.Context, in *TeibEntryAddDel) (*TeibEntryAddDelReply, error)
}

type serviceClient struct {
	conn api.Connection
}

func NewServiceClient(conn api.Connection) RPCService {
	return &serviceClient{conn}
}

func (c *serviceClient) TeibDump(ctx context.Context, in *TeibDump) (RPCService_TeibDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_TeibDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&vpe.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_TeibDumpClient interface {
	Recv() (*TeibDetails, error)
	api.Stream
}

type serviceClient_TeibDumpClient struct {
	api.Stream
}

func (c *serviceClient_TeibDumpClient) Recv() (*TeibDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *TeibDetails:
		return m, nil
	case *vpe.ControlPingReply:
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}

func (c *serviceClient) TeibEntryAddDel(ctx context.Context, in *TeibEntryAddDel) (*TeibEntryAddDelReply, error) {
	out := new(TeibEntryAddDelReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}
