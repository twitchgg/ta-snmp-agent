package server

import "ntsc.ac.cn/ta-snmp-agent/internal/common"

type GeneralServer struct {
	tasClient *common.TASClient
}

func NewGeneralServer(tc *common.TASClient) (*GeneralServer, error) {
	return &GeneralServer{
		tasClient: tc,
	}, nil
}

func (s *GeneralServer) Start() chan error {
	errChan := make(chan error, 1)
	if err := s.tasClient.Regist(); err != nil {
		errChan <- err
		return errChan
	}
	s.genps()
	s.genlxd()
	return errChan
}
