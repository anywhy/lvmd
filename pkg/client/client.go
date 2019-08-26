package client

import (
	"time"

	pb "github.com/anywhy/lvmd/pkg/proto"
	"google.golang.org/grpc"
)

// Client lvm grpc client
type Client struct {
	pb.LVMClient
	conn *grpc.ClientConn
}

// New new lvm grpc client
func New(addr string, timeout time.Duration) (*Client, error) {
	dialOptions := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithTimeout(timeout),
	}

	conn, err := grpc.Dial(addr, dialOptions...)
	if err != nil {
		return nil, err
	}

	return &Client{
		LVMClient: pb.NewLVMClient(conn),
		conn:      conn,
	}, nil
}

// Close close grpc conn
func (c *Client) Close() error {
	return c.conn.Close()
}
