// Code generated by MockGen. DO NOT EDIT.
// Source: ./interface.go

// Package proposals is a generated GoMock package.
package proposals

import (
	context "context"
	reflect "reflect"
	time "time"

	types "github.com/spacemeshos/go-spacemesh/common/types"
	tortoise "github.com/spacemeshos/go-spacemesh/tortoise"
	gomock "go.uber.org/mock/gomock"
)

// MockmeshProvider is a mock of meshProvider interface.
type MockmeshProvider struct {
	ctrl     *gomock.Controller
	recorder *MockmeshProviderMockRecorder
}

// MockmeshProviderMockRecorder is the mock recorder for MockmeshProvider.
type MockmeshProviderMockRecorder struct {
	mock *MockmeshProvider
}

// NewMockmeshProvider creates a new mock instance.
func NewMockmeshProvider(ctrl *gomock.Controller) *MockmeshProvider {
	mock := &MockmeshProvider{ctrl: ctrl}
	mock.recorder = &MockmeshProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockmeshProvider) EXPECT() *MockmeshProviderMockRecorder {
	return m.recorder
}

// AddBallot mocks base method.
func (m *MockmeshProvider) AddBallot(arg0 context.Context, arg1 *types.Ballot) (*types.MalfeasanceProof, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddBallot", arg0, arg1)
	ret0, _ := ret[0].(*types.MalfeasanceProof)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddBallot indicates an expected call of AddBallot.
func (mr *MockmeshProviderMockRecorder) AddBallot(arg0, arg1 interface{}) *meshProviderAddBallotCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddBallot", reflect.TypeOf((*MockmeshProvider)(nil).AddBallot), arg0, arg1)
	return &meshProviderAddBallotCall{Call: call}
}

// meshProviderAddBallotCall wrap *gomock.Call
type meshProviderAddBallotCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *meshProviderAddBallotCall) Return(arg0 *types.MalfeasanceProof, arg1 error) *meshProviderAddBallotCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *meshProviderAddBallotCall) Do(f func(context.Context, *types.Ballot) (*types.MalfeasanceProof, error)) *meshProviderAddBallotCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *meshProviderAddBallotCall) DoAndReturn(f func(context.Context, *types.Ballot) (*types.MalfeasanceProof, error)) *meshProviderAddBallotCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// AddTXsFromProposal mocks base method.
func (m *MockmeshProvider) AddTXsFromProposal(arg0 context.Context, arg1 types.LayerID, arg2 types.ProposalID, arg3 []types.TransactionID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddTXsFromProposal", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddTXsFromProposal indicates an expected call of AddTXsFromProposal.
func (mr *MockmeshProviderMockRecorder) AddTXsFromProposal(arg0, arg1, arg2, arg3 interface{}) *meshProviderAddTXsFromProposalCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTXsFromProposal", reflect.TypeOf((*MockmeshProvider)(nil).AddTXsFromProposal), arg0, arg1, arg2, arg3)
	return &meshProviderAddTXsFromProposalCall{Call: call}
}

// meshProviderAddTXsFromProposalCall wrap *gomock.Call
type meshProviderAddTXsFromProposalCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *meshProviderAddTXsFromProposalCall) Return(arg0 error) *meshProviderAddTXsFromProposalCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *meshProviderAddTXsFromProposalCall) Do(f func(context.Context, types.LayerID, types.ProposalID, []types.TransactionID) error) *meshProviderAddTXsFromProposalCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *meshProviderAddTXsFromProposalCall) DoAndReturn(f func(context.Context, types.LayerID, types.ProposalID, []types.TransactionID) error) *meshProviderAddTXsFromProposalCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ProcessedLayer mocks base method.
func (m *MockmeshProvider) ProcessedLayer() types.LayerID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessedLayer")
	ret0, _ := ret[0].(types.LayerID)
	return ret0
}

// ProcessedLayer indicates an expected call of ProcessedLayer.
func (mr *MockmeshProviderMockRecorder) ProcessedLayer() *meshProviderProcessedLayerCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessedLayer", reflect.TypeOf((*MockmeshProvider)(nil).ProcessedLayer))
	return &meshProviderProcessedLayerCall{Call: call}
}

// meshProviderProcessedLayerCall wrap *gomock.Call
type meshProviderProcessedLayerCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *meshProviderProcessedLayerCall) Return(arg0 types.LayerID) *meshProviderProcessedLayerCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *meshProviderProcessedLayerCall) Do(f func() types.LayerID) *meshProviderProcessedLayerCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *meshProviderProcessedLayerCall) DoAndReturn(f func() types.LayerID) *meshProviderProcessedLayerCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockeligibilityValidator is a mock of eligibilityValidator interface.
type MockeligibilityValidator struct {
	ctrl     *gomock.Controller
	recorder *MockeligibilityValidatorMockRecorder
}

// MockeligibilityValidatorMockRecorder is the mock recorder for MockeligibilityValidator.
type MockeligibilityValidatorMockRecorder struct {
	mock *MockeligibilityValidator
}

// NewMockeligibilityValidator creates a new mock instance.
func NewMockeligibilityValidator(ctrl *gomock.Controller) *MockeligibilityValidator {
	mock := &MockeligibilityValidator{ctrl: ctrl}
	mock.recorder = &MockeligibilityValidatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockeligibilityValidator) EXPECT() *MockeligibilityValidatorMockRecorder {
	return m.recorder
}

// CheckEligibility mocks base method.
func (m *MockeligibilityValidator) CheckEligibility(arg0 context.Context, arg1 *types.Ballot) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckEligibility", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckEligibility indicates an expected call of CheckEligibility.
func (mr *MockeligibilityValidatorMockRecorder) CheckEligibility(arg0, arg1 interface{}) *eligibilityValidatorCheckEligibilityCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckEligibility", reflect.TypeOf((*MockeligibilityValidator)(nil).CheckEligibility), arg0, arg1)
	return &eligibilityValidatorCheckEligibilityCall{Call: call}
}

// eligibilityValidatorCheckEligibilityCall wrap *gomock.Call
type eligibilityValidatorCheckEligibilityCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *eligibilityValidatorCheckEligibilityCall) Return(arg0 bool, arg1 error) *eligibilityValidatorCheckEligibilityCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *eligibilityValidatorCheckEligibilityCall) Do(f func(context.Context, *types.Ballot) (bool, error)) *eligibilityValidatorCheckEligibilityCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *eligibilityValidatorCheckEligibilityCall) DoAndReturn(f func(context.Context, *types.Ballot) (bool, error)) *eligibilityValidatorCheckEligibilityCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MocktortoiseProvider is a mock of tortoiseProvider interface.
type MocktortoiseProvider struct {
	ctrl     *gomock.Controller
	recorder *MocktortoiseProviderMockRecorder
}

// MocktortoiseProviderMockRecorder is the mock recorder for MocktortoiseProvider.
type MocktortoiseProviderMockRecorder struct {
	mock *MocktortoiseProvider
}

// NewMocktortoiseProvider creates a new mock instance.
func NewMocktortoiseProvider(ctrl *gomock.Controller) *MocktortoiseProvider {
	mock := &MocktortoiseProvider{ctrl: ctrl}
	mock.recorder = &MocktortoiseProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MocktortoiseProvider) EXPECT() *MocktortoiseProviderMockRecorder {
	return m.recorder
}

// DecodeBallot mocks base method.
func (m *MocktortoiseProvider) DecodeBallot(arg0 *types.BallotTortoiseData) (*tortoise.DecodedBallot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DecodeBallot", arg0)
	ret0, _ := ret[0].(*tortoise.DecodedBallot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DecodeBallot indicates an expected call of DecodeBallot.
func (mr *MocktortoiseProviderMockRecorder) DecodeBallot(arg0 interface{}) *tortoiseProviderDecodeBallotCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DecodeBallot", reflect.TypeOf((*MocktortoiseProvider)(nil).DecodeBallot), arg0)
	return &tortoiseProviderDecodeBallotCall{Call: call}
}

// tortoiseProviderDecodeBallotCall wrap *gomock.Call
type tortoiseProviderDecodeBallotCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *tortoiseProviderDecodeBallotCall) Return(arg0 *tortoise.DecodedBallot, arg1 error) *tortoiseProviderDecodeBallotCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *tortoiseProviderDecodeBallotCall) Do(f func(*types.BallotTortoiseData) (*tortoise.DecodedBallot, error)) *tortoiseProviderDecodeBallotCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *tortoiseProviderDecodeBallotCall) DoAndReturn(f func(*types.BallotTortoiseData) (*tortoise.DecodedBallot, error)) *tortoiseProviderDecodeBallotCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetBallot mocks base method.
func (m *MocktortoiseProvider) GetBallot(arg0 types.BallotID) *tortoise.BallotData {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBallot", arg0)
	ret0, _ := ret[0].(*tortoise.BallotData)
	return ret0
}

// GetBallot indicates an expected call of GetBallot.
func (mr *MocktortoiseProviderMockRecorder) GetBallot(arg0 interface{}) *tortoiseProviderGetBallotCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBallot", reflect.TypeOf((*MocktortoiseProvider)(nil).GetBallot), arg0)
	return &tortoiseProviderGetBallotCall{Call: call}
}

// tortoiseProviderGetBallotCall wrap *gomock.Call
type tortoiseProviderGetBallotCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *tortoiseProviderGetBallotCall) Return(arg0 *tortoise.BallotData) *tortoiseProviderGetBallotCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *tortoiseProviderGetBallotCall) Do(f func(types.BallotID) *tortoise.BallotData) *tortoiseProviderGetBallotCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *tortoiseProviderGetBallotCall) DoAndReturn(f func(types.BallotID) *tortoise.BallotData) *tortoiseProviderGetBallotCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetMissingActiveSet mocks base method.
func (m *MocktortoiseProvider) GetMissingActiveSet(arg0 types.EpochID, arg1 []types.ATXID) []types.ATXID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMissingActiveSet", arg0, arg1)
	ret0, _ := ret[0].([]types.ATXID)
	return ret0
}

// GetMissingActiveSet indicates an expected call of GetMissingActiveSet.
func (mr *MocktortoiseProviderMockRecorder) GetMissingActiveSet(arg0, arg1 interface{}) *tortoiseProviderGetMissingActiveSetCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMissingActiveSet", reflect.TypeOf((*MocktortoiseProvider)(nil).GetMissingActiveSet), arg0, arg1)
	return &tortoiseProviderGetMissingActiveSetCall{Call: call}
}

// tortoiseProviderGetMissingActiveSetCall wrap *gomock.Call
type tortoiseProviderGetMissingActiveSetCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *tortoiseProviderGetMissingActiveSetCall) Return(arg0 []types.ATXID) *tortoiseProviderGetMissingActiveSetCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *tortoiseProviderGetMissingActiveSetCall) Do(f func(types.EpochID, []types.ATXID) []types.ATXID) *tortoiseProviderGetMissingActiveSetCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *tortoiseProviderGetMissingActiveSetCall) DoAndReturn(f func(types.EpochID, []types.ATXID) []types.ATXID) *tortoiseProviderGetMissingActiveSetCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// StoreBallot mocks base method.
func (m *MocktortoiseProvider) StoreBallot(arg0 *tortoise.DecodedBallot) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StoreBallot", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// StoreBallot indicates an expected call of StoreBallot.
func (mr *MocktortoiseProviderMockRecorder) StoreBallot(arg0 interface{}) *tortoiseProviderStoreBallotCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreBallot", reflect.TypeOf((*MocktortoiseProvider)(nil).StoreBallot), arg0)
	return &tortoiseProviderStoreBallotCall{Call: call}
}

// tortoiseProviderStoreBallotCall wrap *gomock.Call
type tortoiseProviderStoreBallotCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *tortoiseProviderStoreBallotCall) Return(arg0 error) *tortoiseProviderStoreBallotCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *tortoiseProviderStoreBallotCall) Do(f func(*tortoise.DecodedBallot) error) *tortoiseProviderStoreBallotCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *tortoiseProviderStoreBallotCall) DoAndReturn(f func(*tortoise.DecodedBallot) error) *tortoiseProviderStoreBallotCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockvrfVerifier is a mock of vrfVerifier interface.
type MockvrfVerifier struct {
	ctrl     *gomock.Controller
	recorder *MockvrfVerifierMockRecorder
}

// MockvrfVerifierMockRecorder is the mock recorder for MockvrfVerifier.
type MockvrfVerifierMockRecorder struct {
	mock *MockvrfVerifier
}

// NewMockvrfVerifier creates a new mock instance.
func NewMockvrfVerifier(ctrl *gomock.Controller) *MockvrfVerifier {
	mock := &MockvrfVerifier{ctrl: ctrl}
	mock.recorder = &MockvrfVerifierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockvrfVerifier) EXPECT() *MockvrfVerifierMockRecorder {
	return m.recorder
}

// Verify mocks base method.
func (m *MockvrfVerifier) Verify(arg0 types.NodeID, arg1 []byte, arg2 types.VrfSignature) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Verify", arg0, arg1, arg2)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Verify indicates an expected call of Verify.
func (mr *MockvrfVerifierMockRecorder) Verify(arg0, arg1, arg2 interface{}) *vrfVerifierVerifyCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Verify", reflect.TypeOf((*MockvrfVerifier)(nil).Verify), arg0, arg1, arg2)
	return &vrfVerifierVerifyCall{Call: call}
}

// vrfVerifierVerifyCall wrap *gomock.Call
type vrfVerifierVerifyCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *vrfVerifierVerifyCall) Return(arg0 bool) *vrfVerifierVerifyCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *vrfVerifierVerifyCall) Do(f func(types.NodeID, []byte, types.VrfSignature) bool) *vrfVerifierVerifyCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *vrfVerifierVerifyCall) DoAndReturn(f func(types.NodeID, []byte, types.VrfSignature) bool) *vrfVerifierVerifyCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MocknonceFetcher is a mock of nonceFetcher interface.
type MocknonceFetcher struct {
	ctrl     *gomock.Controller
	recorder *MocknonceFetcherMockRecorder
}

// MocknonceFetcherMockRecorder is the mock recorder for MocknonceFetcher.
type MocknonceFetcherMockRecorder struct {
	mock *MocknonceFetcher
}

// NewMocknonceFetcher creates a new mock instance.
func NewMocknonceFetcher(ctrl *gomock.Controller) *MocknonceFetcher {
	mock := &MocknonceFetcher{ctrl: ctrl}
	mock.recorder = &MocknonceFetcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MocknonceFetcher) EXPECT() *MocknonceFetcherMockRecorder {
	return m.recorder
}

// VRFNonce mocks base method.
func (m *MocknonceFetcher) VRFNonce(arg0 types.NodeID, arg1 types.EpochID) (types.VRFPostIndex, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VRFNonce", arg0, arg1)
	ret0, _ := ret[0].(types.VRFPostIndex)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VRFNonce indicates an expected call of VRFNonce.
func (mr *MocknonceFetcherMockRecorder) VRFNonce(arg0, arg1 interface{}) *nonceFetcherVRFNonceCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VRFNonce", reflect.TypeOf((*MocknonceFetcher)(nil).VRFNonce), arg0, arg1)
	return &nonceFetcherVRFNonceCall{Call: call}
}

// nonceFetcherVRFNonceCall wrap *gomock.Call
type nonceFetcherVRFNonceCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *nonceFetcherVRFNonceCall) Return(arg0 types.VRFPostIndex, arg1 error) *nonceFetcherVRFNonceCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *nonceFetcherVRFNonceCall) Do(f func(types.NodeID, types.EpochID) (types.VRFPostIndex, error)) *nonceFetcherVRFNonceCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *nonceFetcherVRFNonceCall) DoAndReturn(f func(types.NodeID, types.EpochID) (types.VRFPostIndex, error)) *nonceFetcherVRFNonceCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MocklayerClock is a mock of layerClock interface.
type MocklayerClock struct {
	ctrl     *gomock.Controller
	recorder *MocklayerClockMockRecorder
}

// MocklayerClockMockRecorder is the mock recorder for MocklayerClock.
type MocklayerClockMockRecorder struct {
	mock *MocklayerClock
}

// NewMocklayerClock creates a new mock instance.
func NewMocklayerClock(ctrl *gomock.Controller) *MocklayerClock {
	mock := &MocklayerClock{ctrl: ctrl}
	mock.recorder = &MocklayerClockMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MocklayerClock) EXPECT() *MocklayerClockMockRecorder {
	return m.recorder
}

// CurrentLayer mocks base method.
func (m *MocklayerClock) CurrentLayer() types.LayerID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CurrentLayer")
	ret0, _ := ret[0].(types.LayerID)
	return ret0
}

// CurrentLayer indicates an expected call of CurrentLayer.
func (mr *MocklayerClockMockRecorder) CurrentLayer() *layerClockCurrentLayerCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CurrentLayer", reflect.TypeOf((*MocklayerClock)(nil).CurrentLayer))
	return &layerClockCurrentLayerCall{Call: call}
}

// layerClockCurrentLayerCall wrap *gomock.Call
type layerClockCurrentLayerCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *layerClockCurrentLayerCall) Return(arg0 types.LayerID) *layerClockCurrentLayerCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *layerClockCurrentLayerCall) Do(f func() types.LayerID) *layerClockCurrentLayerCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *layerClockCurrentLayerCall) DoAndReturn(f func() types.LayerID) *layerClockCurrentLayerCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// LayerToTime mocks base method.
func (m *MocklayerClock) LayerToTime(arg0 types.LayerID) time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LayerToTime", arg0)
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// LayerToTime indicates an expected call of LayerToTime.
func (mr *MocklayerClockMockRecorder) LayerToTime(arg0 interface{}) *layerClockLayerToTimeCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LayerToTime", reflect.TypeOf((*MocklayerClock)(nil).LayerToTime), arg0)
	return &layerClockLayerToTimeCall{Call: call}
}

// layerClockLayerToTimeCall wrap *gomock.Call
type layerClockLayerToTimeCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *layerClockLayerToTimeCall) Return(arg0 time.Time) *layerClockLayerToTimeCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *layerClockLayerToTimeCall) Do(f func(types.LayerID) time.Time) *layerClockLayerToTimeCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *layerClockLayerToTimeCall) DoAndReturn(f func(types.LayerID) time.Time) *layerClockLayerToTimeCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
