package server

import (
	"context"
	"strings"

	"github.com/anywhy/lvmd/pkg/commands"
	pb "github.com/anywhy/lvmd/pkg/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

// Server lvm grpc server
type Server struct {
}

// NewServer new server
func NewServer() pb.LVMServer {
	return Server{}
}

// ListLV list lv
func (s Server) ListLV(ctx context.Context, in *pb.ListLVRequest) (*pb.ListLVReply, error) {
	lvs, err := commands.ListLV(ctx, in.VolumeGroup)
	if err != nil {
		return nil, grpc.Errorf(codes.Internal, "failed to list LVs: %v\nCommandOutput: %v", err, lvs)
	}

	pblvs := make([]*pb.LogicalVolume, len(lvs))
	for i, v := range lvs {
		pblvs[i] = v.ToProto()
	}
	return &pb.ListLVReply{Volumes: pblvs}, nil
}

// CreateLV create lv
func (s Server) CreateLV(ctx context.Context, in *pb.CreateLVRequest) (*pb.CreateLVReply, error) {
	log, err := commands.CreateLV(ctx, in.VolumeGroup, in.Name, in.Size, in.Mirrors, in.Tags)
	if err != nil {
		return nil, grpc.Errorf(codes.Internal, "failed to create lv: %v\nCommandOutput: %v", err, streamline(log))
	}
	return &pb.CreateLVReply{CommandOutput: log}, nil
}

// RemoveLV remove lv
func (s Server) RemoveLV(ctx context.Context, in *pb.RemoveLVRequest) (*pb.RemoveLVReply, error) {
	log, err := commands.RemoveLV(ctx, in.VolumeGroup, in.Name)
	if err != nil {
		return nil, grpc.Errorf(codes.Internal, "failed to remove lv: %v\nCommandOutput: %v", err, streamline(log))
	}
	return &pb.RemoveLVReply{CommandOutput: log}, nil
}

// CloneLV clone lv
func (s Server) CloneLV(ctx context.Context, in *pb.CloneLVRequest) (*pb.CloneLVReply, error) {
	log, err := commands.CloneLV(ctx, in.SourceName, in.DestName)
	if err != nil {
		return nil, grpc.Errorf(codes.Internal, "failed to clone lv: %v\nCommandOutput: %v", err, streamline(log))
	}
	return &pb.CloneLVReply{CommandOutput: log}, nil
}

// ExtendLV extend lv
func (s Server) ExtendLV(ctx context.Context, in *pb.ExtendLVRequest) (*pb.ExtendLVReply, error) {
	log, err := commands.ExtendLV(ctx, in.GetSize(), in.GetName())
	if err != nil {
		return nil, grpc.Errorf(codes.Internal, "failed to extend lv: %v\nCommandOutput: %v", err, streamline(log))
	}
	return &pb.ExtendLVReply{CommandOutput: log}, nil
}

// ListVG list vg
func (s Server) ListVG(ctx context.Context, in *pb.ListVGRequest) (*pb.ListVGReply, error) {
	vgs, err := commands.ListVG(ctx)
	if err != nil {
		return nil, grpc.Errorf(codes.Internal, "failed to list LVs: %v\nCommandOutput: %v", err, vgs)
	}

	pbvgs := make([]*pb.VolumeGroup, len(vgs))
	for i, v := range vgs {
		pbvgs[i] = v.ToProto()
	}
	return &pb.ListVGReply{VolumeGroups: pbvgs}, nil
}

// CreateVG create vg
func (s Server) CreateVG(ctx context.Context, in *pb.CreateVGRequest) (*pb.CreateVGReply, error) {
	log, err := commands.CreateVG(ctx, in.Name, in.PhysicalVolume, in.Tags)
	if err != nil {
		return nil, grpc.Errorf(codes.Internal, "failed to create vg: %v\nCommandOutput: %v", err, streamline(log))
	}
	return &pb.CreateVGReply{CommandOutput: log}, nil
}

// ExtendVG extend vg
func (s Server) ExtendVG(ctx context.Context, in *pb.ExtendVGRequest) (*pb.ExtendVGReply, error) {
	log, err := commands.ExtendVG(ctx, in.Name, in.PhysicalVolume)
	if err != nil {
		return nil, grpc.Errorf(codes.Internal, "failed to extend vg: %v\nCommandOutput: %v", err, streamline(log))
	}
	return &pb.ExtendVGReply{CommandOutput: log}, nil
}

// ReduceVG reduce vg
func (s Server) ReduceVG(ctx context.Context, in *pb.ExtendVGRequest) (*pb.ExtendVGReply, error) {
	log, err := commands.ReduceVG(ctx, in.Name, in.PhysicalVolume)
	if err != nil {
		return nil, grpc.Errorf(codes.Internal, "failed to reduce vg: %v\nCommandOutput: %v", err, streamline(log))
	}
	return &pb.ExtendVGReply{CommandOutput: log}, nil
}

// RemoveVG remove vg
func (s Server) RemoveVG(ctx context.Context, in *pb.CreateVGRequest) (*pb.RemoveVGReply, error) {
	log, err := commands.RemoveVG(ctx, in.Name)
	if err != nil {
		return nil, grpc.Errorf(codes.Internal, "failed to remove vg: %v\nCommandOutput: %v", err, streamline(log))
	}
	return &pb.RemoveVGReply{CommandOutput: log}, nil
}

// AddTagLV add lv tag
func (s Server) AddTagLV(ctx context.Context, in *pb.AddTagLVRequest) (*pb.AddTagLVReply, error) {
	log, err := commands.AddTagLV(ctx, in.VolumeGroup, in.Name, in.Tags)
	if err != nil {
		return nil, grpc.Errorf(codes.Internal, "failed to add tags to lv: %v\nCommandOutput: %v", err, streamline(log))
	}
	return &pb.AddTagLVReply{CommandOutput: log}, nil
}

// RemoveTagLV remove lv tag
func (s Server) RemoveTagLV(ctx context.Context, in *pb.RemoveTagLVRequest) (*pb.RemoveTagLVReply, error) {
	log, err := commands.RemoveTagLV(ctx, in.VolumeGroup, in.Name, in.Tags)
	if err != nil {
		return nil, grpc.Errorf(codes.Internal, "failed to remove tags from lv: %v\nCommandOutput: %v", err, streamline(log))
	}
	return &pb.RemoveTagLVReply{CommandOutput: log}, nil
}

// CreatePV create pv
func (s Server) CreatePV(ctx context.Context, in *pb.CreatePVRequest) (*pb.CreatePVReply, error) {
	log, err := commands.CreatePV(ctx, in.Block)
	if err != nil {
		return nil, grpc.Errorf(codes.Internal, "failed to create pv: %v\nCommandOutput: %v", err, streamline(log))
	}
	return &pb.CreatePVReply{CommandOutput: log}, nil
}

// RemovePV remove pv
func (s Server) RemovePV(ctx context.Context, in *pb.RemovePVRequest) (*pb.RemovePVReply, error) {
	log, err := commands.RemovePV(ctx, in.Block)
	if err != nil {
		return nil, grpc.Errorf(codes.Internal, "failed to remove pv: %v\nCommandOutput: %v", err, streamline(log))
	}
	return &pb.RemovePVReply{CommandOutput: log}, nil
}

// ListPV list pv
func (s Server) ListPV(ctx context.Context, in *pb.ListPVRequest) (*pb.ListPVReply, error) {
	pvs, err := commands.ListPV(ctx)
	if err != nil {
		return nil, grpc.Errorf(codes.Internal, "failed to list pv: %v\nCommandOutput: %v", err, pvs)
	}
	pbpvs := make([]*pb.PVInfo, len(pvs))
	for i, v := range pvs {
		pbpvs[i] = v.ToProto()
	}
	return &pb.ListPVReply{Pvinfos: pbpvs}, nil
}

// Validate validate
func (s Server) Validate(ctx context.Context, in *pb.ValidateRequest) (*pb.ValidateReply, error) {
	v, err := commands.Validate(ctx, in.Block)
	if err != nil {
		return nil, grpc.Errorf(codes.Internal, "failed to validate block: %v\nCommandOutput: %v", err, v)
	}
	return &pb.ValidateReply{Validate: v}, nil
}

// Destory destory block
func (s Server) Destory(ctx context.Context, in *pb.DestoryRequest) (*pb.DestoryReply, error) {
	log, err := commands.Destory(ctx, in.Block)
	if err != nil {
		return nil, grpc.Errorf(codes.Internal, "failed to destory block: %v\nCommandOutput: %v", err, streamline(log))
	}
	return &pb.DestoryReply{CommandOutput: log}, nil
}

// Match match block
func (s Server) Match(ctx context.Context, in *pb.MatchRequest) (*pb.MatchReply, error) {
	log := commands.Match(ctx, in.Block)
	return &pb.MatchReply{CommandOutput: log}, nil
}

// GetPVNum get pv number
func (s Server) GetPVNum(ctx context.Context, in *pb.CreateVGRequest) (*pb.GetPVNumReply, error) {
	log, err := commands.GetPVNum(ctx, in.Name)
	if err != nil {
		return nil, grpc.Errorf(codes.Internal, "failed to get vg's pv num: %v\nCommandOutput: %v", err, streamline(log))
	}
	return &pb.GetPVNumReply{CommandOutput: log}, nil
}

func streamline(out string) string {
	var res string
	for _, l := range strings.Split(out, "\n") {
		if len(l) == 0 {
			continue
		}
		if !strings.Contains(l, "/etc/lvm/cache/.cache") {
			res = res + l + "\n"
		}
	}
	return strings.TrimSpace(res)
}
