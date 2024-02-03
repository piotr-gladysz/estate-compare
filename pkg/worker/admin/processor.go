package admin

import (
	"context"
	"github.com/piotr-gladysz/estate-compare/pkg/api"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type ProcessorState struct {
	LastRun   time.Time
	NextRun   time.Time
	IsRunning bool
}

type ProcessorControl interface {
	RunProcessing() error
	StopProcessing() error
	GetProcessingState() (ProcessorState, error)
}

type ProcessorServer struct {
	control ProcessorControl
	api.UnimplementedProcessorServiceServer
}

func NewProcessorServer(control ProcessorControl) *ProcessorServer {
	return &ProcessorServer{control: control}
}

func (p *ProcessorServer) StartProcessing(ctx context.Context, empty *emptypb.Empty) (*emptypb.Empty, error) {
	err := p.control.RunProcessing()
	return nil, err
}

func (p *ProcessorServer) StopProcessing(ctx context.Context, empty *emptypb.Empty) (*emptypb.Empty, error) {
	err := p.control.StopProcessing()
	return nil, err
}

func (p *ProcessorServer) GetProcessingStatus(ctx context.Context, empty *emptypb.Empty) (*api.ProcessingStatus, error) {
	state, err := p.control.GetProcessingState()
	if err != nil {

		return nil, err
	}

	return &api.ProcessingStatus{
		LastRun:   timestamppb.New(state.LastRun),
		NextRun:   timestamppb.New(state.NextRun),
		IsRunning: state.IsRunning,
	}, nil
}
