package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"strings"

	pb "github.com/anywhy/lvmd/pkg/proto"
	"github.com/anywhy/lvmd/pkg/server"
	"github.com/golang/glog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	endpoint = flag.String("endpoint", "unix://tmp/lvmd.sock", "lvmd endpoint")
)

func main() {
	flag.Parse()

	proto, addr, err := parseEndpoint(*endpoint)
	if err != nil {
		glog.Fatal(err.Error())
	}
	if proto == "unix" {
		addr = "/" + addr
		if err := os.Remove(addr); err != nil && !os.IsNotExist(err) {
			glog.Fatalf("Failed to remove %s, error: %s", addr, err.Error())
		}
	}

	listener, err := net.Listen(proto, addr)
	if err != nil {
		glog.Fatalf("Failed to listen: %v", err)
	}

	svr := server.NewServer()
	server := grpc.NewServer()
	reflection.Register(server)
	pb.RegisterLVMServer(server, svr)
	if err = server.Serve(listener); err != nil {
		glog.Fatalf("run grpc server failed:%s", err.Error())
	}
}

func parseEndpoint(ep string) (string, string, error) {
	if strings.HasPrefix(strings.ToLower(ep), "unix://") || strings.HasPrefix(strings.ToLower(ep), "tcp://") {
		s := strings.SplitN(ep, "://", 2)
		if s[1] != "" {
			return s[0], s[1], nil
		}
	}
	return "", "", fmt.Errorf("Invalid endpoint: %v", ep)
}
