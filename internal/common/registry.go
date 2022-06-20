package common

import (
	"context"
	"fmt"
	"time"

	"github.com/denisbrodbeck/machineid"
	"google.golang.org/protobuf/types/known/timestamppb"
	"ntsc.ac.cn/ta-registry/pkg/pb"
	"ntsc.ac.cn/ta-registry/pkg/rpc"
)

// TASClient
type TASClient struct {
	conf      *TASConfig
	machineID string
	rsc       pb.RegistryServiceClient
}

// NewTASClient create tas client
func NewTASClient(conf *TASConfig) (*TASClient, error) {
	if conf == nil {
		return nil, fmt.Errorf("rpc server config is not define")
	}
	machineID, err := machineid.ID()
	if err != nil {
		return nil, fmt.Errorf("generate machine id failed: %v", err)
	}
	if err := conf.Check(); err != nil {
		return nil, fmt.Errorf("check config failed: %v", err)
	}
	tlsConf, err := conf.GetTlsConfig(machineID)
	if err != nil {
		return nil, fmt.Errorf("generate tls config failed: %v", err)
	}
	conn, err := rpc.DialRPCConn(&rpc.DialOptions{
		RemoteAddr: conf.ManagerEndpoint,
		TLSConfig:  tlsConf,
	})
	if err != nil {
		return nil, fmt.Errorf(
			"dial management grpc connection failed: %v", err)
	}
	return &TASClient{
		conf:      conf,
		machineID: machineID,
		rsc:       pb.NewRegistryServiceClient(conn),
	}, nil
}

// MachineID machine id
func (tc *TASClient) MachineID() string {
	return tc.machineID
}

func (tc *TASClient) Regist() error {
	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(time.Second*3))
	defer cancel()
	conf, err := tc.rsc.RegistServer(ctx, &pb.RegistServerRequest{
		SysTime:   timestamppb.Now(),
		MachineID: tc.machineID,
	})
	if err != nil {
		return err
	}
	fmt.Println(conf)
	return nil
}
