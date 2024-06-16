package controllermock

import (
	"github.com/stretchr/testify/mock"
	"github.com/yyyyymmmmm/Test/pkg/aria2/common"
	"github.com/yyyyymmmmm/Test/pkg/cluster"
	"github.com/yyyyymmmmm/Test/pkg/mq"
	"github.com/yyyyymmmmm/Test/pkg/serializer"
)

type SlaveControllerMock struct {
	mock.Mock
}

func (s SlaveControllerMock) HandleHeartBeat(pingReq *serializer.NodePingReq) (serializer.NodePingResp, error) {
	args := s.Called(pingReq)
	return args.Get(0).(serializer.NodePingResp), args.Error(1)
}

func (s SlaveControllerMock) GetAria2Instance(s2 string) (common.Aria2, error) {
	args := s.Called(s2)
	return args.Get(0).(common.Aria2), args.Error(1)
}

func (s SlaveControllerMock) SendNotification(s3 string, s2 string, message mq.Message) error {
	args := s.Called(s3, s2, message)
	return args.Error(0)
}

func (s SlaveControllerMock) SubmitTask(s3 string, i interface{}, s2 string, f func(interface{})) error {
	args := s.Called(s3, i, s2, f)
	return args.Error(0)
}

func (s SlaveControllerMock) GetMasterInfo(s2 string) (*cluster.MasterInfo, error) {
	args := s.Called(s2)
	return args.Get(0).(*cluster.MasterInfo), args.Error(1)
}

func (s SlaveControllerMock) GetPolicyOauthToken(s2 string, u uint) (string, error) {
	args := s.Called(s2, u)
	return args.String(0), args.Error(1)
}
