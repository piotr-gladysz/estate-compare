package crawler

import (
	"errors"
	"github.com/piotr-gladysz/estate-compare/pkg/worker/admin"
)

var _ admin.ProcessorControl = (*SitesProcessor)(nil)

var AlreadyRunningError = errors.New("processor already running")
var NotRunningError = errors.New("processor not running")
var NotStartedError = errors.New("processor not started")

func (s *SitesProcessor) RunProcessing() error {
	if s.procJob == nil {
		return NotStartedError
	}
	s.procMux.Lock()
	defer s.procMux.Unlock()

	if s.procCtx != nil {
		return AlreadyRunningError
	}

	return s.procJob.RunNow()
}

func (s *SitesProcessor) StopProcessing() error {
	if s.procJob == nil {
		return NotStartedError
	}

	s.procMux.Lock()
	defer s.procMux.Unlock()

	if s.procCtx == nil {
		return NotRunningError
	}

	s.procCancel()
	return nil
}

func (s *SitesProcessor) GetProcessingState() (admin.ProcessorState, error) {
	ret := admin.ProcessorState{}
	if s.procJob == nil {
		return ret, NotStartedError
	}

	lastRun, err := s.procJob.LastRun()
	if err != nil {
		return ret, err
	}

	nextRun, err := s.procJob.NextRun()
	if err != nil {
		return ret, err
	}

	s.procMux.Lock()
	defer s.procMux.Unlock()

	return admin.ProcessorState{
		LastRun:   lastRun,
		NextRun:   nextRun,
		IsRunning: s.procCtx != nil,
	}, nil

}
