package methods

import (
	"context"

	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vsan/types"
)

type VsanPerfDiagnoseBody struct {
	Req    *types.VsanPerfDiagnose         `xml:"urn:vsan VsanPerfDiagnose,omitempty"`
	Res    *types.VsanPerfDiagnoseResponse `xml:"urn:vsan VsanPerfDiagnoseResponse,omitempty"`
	Fault_ *soap.Fault                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanPerfDiagnoseBody) Fault() *soap.Fault { return b.Fault_ }

func VsanPerfDiagnose(ctx context.Context, r soap.RoundTripper, req *types.VsanPerfDiagnose) (*types.VsanPerfDiagnoseResponse, error) {
	var reqBody, resBody VsanPerfDiagnoseBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanPerfQueryClusterHealthBody struct {
	Req    *types.VsanPerfQueryClusterHealth         `xml:"urn:vsan VsanPerfQueryClusterHealth,omitempty"`
	Res    *types.VsanPerfQueryClusterHealthResponse `xml:"urn:vsan VsanPerfQueryClusterHealthResponse,omitempty"`
	Fault_ *soap.Fault                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanPerfQueryClusterHealthBody) Fault() *soap.Fault { return b.Fault_ }

func VsanPerfQueryClusterHealth(ctx context.Context, r soap.RoundTripper, req *types.VsanPerfQueryClusterHealth) (*types.VsanPerfQueryClusterHealthResponse, error) {
	var reqBody, resBody VsanPerfQueryClusterHealthBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanPerfStopStatsCollectorBody struct {
	Req    *types.VsanPerfStopStatsCollector         `xml:"urn:vsan VsanPerfStopStatsCollector,omitempty"`
	Res    *types.VsanPerfStopStatsCollectorResponse `xml:"urn:vsan VsanPerfStopStatsCollectorResponse,omitempty"`
	Fault_ *soap.Fault                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanPerfStopStatsCollectorBody) Fault() *soap.Fault { return b.Fault_ }

func VsanPerfStopStatsCollector(ctx context.Context, r soap.RoundTripper, req *types.VsanPerfStopStatsCollector) (*types.VsanPerfStopStatsCollectorResponse, error) {
	var reqBody, resBody VsanPerfStopStatsCollectorBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type GetVsanPerfDiagnosisResultBody struct {
	Req    *types.GetVsanPerfDiagnosisResult         `xml:"urn:vsan GetVsanPerfDiagnosisResult,omitempty"`
	Res    *types.GetVsanPerfDiagnosisResultResponse `xml:"urn:vsan GetVsanPerfDiagnosisResultResponse,omitempty"`
	Fault_ *soap.Fault                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *GetVsanPerfDiagnosisResultBody) Fault() *soap.Fault { return b.Fault_ }

func GetVsanPerfDiagnosisResult(ctx context.Context, r soap.RoundTripper, req *types.GetVsanPerfDiagnosisResult) (*types.GetVsanPerfDiagnosisResultResponse, error) {
	var reqBody, resBody GetVsanPerfDiagnosisResultBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanPerfToggleVerboseModeBody struct {
	Req    *types.VsanPerfToggleVerboseMode         `xml:"urn:vsan VsanPerfToggleVerboseMode,omitempty"`
	Res    *types.VsanPerfToggleVerboseModeResponse `xml:"urn:vsan VsanPerfToggleVerboseModeResponse,omitempty"`
	Fault_ *soap.Fault                              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanPerfToggleVerboseModeBody) Fault() *soap.Fault { return b.Fault_ }

func VsanPerfToggleVerboseMode(ctx context.Context, r soap.RoundTripper, req *types.VsanPerfToggleVerboseMode) (*types.VsanPerfToggleVerboseModeResponse, error) {
	var reqBody, resBody VsanPerfToggleVerboseModeBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanPerfDiagnoseTaskBody struct {
	Req    *types.VsanPerfDiagnoseTask         `xml:"urn:vsan VsanPerfDiagnoseTask,omitempty"`
	Res    *types.VsanPerfDiagnoseTaskResponse `xml:"urn:vsan VsanPerfDiagnoseTaskResponse,omitempty"`
	Fault_ *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanPerfDiagnoseTaskBody) Fault() *soap.Fault { return b.Fault_ }

func VsanPerfDiagnoseTask(ctx context.Context, r soap.RoundTripper, req *types.VsanPerfDiagnoseTask) (*types.VsanPerfDiagnoseTaskResponse, error) {
	var reqBody, resBody VsanPerfDiagnoseTaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryVsanPerfTopEntitiesBody struct {
	Req    *types.QueryVsanPerfTopEntities         `xml:"urn:vsan QueryVsanPerfTopEntities,omitempty"`
	Res    *types.QueryVsanPerfTopEntitiesResponse `xml:"urn:vsan QueryVsanPerfTopEntitiesResponse,omitempty"`
	Fault_ *soap.Fault                             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryVsanPerfTopEntitiesBody) Fault() *soap.Fault { return b.Fault_ }

func QueryVsanPerfTopEntities(ctx context.Context, r soap.RoundTripper, req *types.QueryVsanPerfTopEntities) (*types.QueryVsanPerfTopEntitiesResponse, error) {
	var reqBody, resBody QueryVsanPerfTopEntitiesBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanPerfSetClusterMembersBody struct {
	Req    *types.VsanPerfSetClusterMembers         `xml:"urn:vsan VsanPerfSetClusterMembers,omitempty"`
	Res    *types.VsanPerfSetClusterMembersResponse `xml:"urn:vsan VsanPerfSetClusterMembersResponse,omitempty"`
	Fault_ *soap.Fault                              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanPerfSetClusterMembersBody) Fault() *soap.Fault { return b.Fault_ }

func VsanPerfSetClusterMembers(ctx context.Context, r soap.RoundTripper, req *types.VsanPerfSetClusterMembers) (*types.VsanPerfSetClusterMembersResponse, error) {
	var reqBody, resBody VsanPerfSetClusterMembersBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryHostNetworkConfigInfoVnicBody struct {
	Req    *types.QueryHostNetworkConfigInfoVnic         `xml:"urn:vsan QueryHostNetworkConfigInfoVnic,omitempty"`
	Res    *types.QueryHostNetworkConfigInfoVnicResponse `xml:"urn:vsan QueryHostNetworkConfigInfoVnicResponse,omitempty"`
	Fault_ *soap.Fault                                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryHostNetworkConfigInfoVnicBody) Fault() *soap.Fault { return b.Fault_ }

func QueryHostNetworkConfigInfoVnic(ctx context.Context, r soap.RoundTripper, req *types.QueryHostNetworkConfigInfoVnic) (*types.QueryHostNetworkConfigInfoVnicResponse, error) {
	var reqBody, resBody QueryHostNetworkConfigInfoVnicBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type SubmitFeedbackForDiagnosisResultBody struct {
	Req    *types.SubmitFeedbackForDiagnosisResult         `xml:"urn:vsan SubmitFeedbackForDiagnosisResult,omitempty"`
	Res    *types.SubmitFeedbackForDiagnosisResultResponse `xml:"urn:vsan SubmitFeedbackForDiagnosisResultResponse,omitempty"`
	Fault_ *soap.Fault                                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *SubmitFeedbackForDiagnosisResultBody) Fault() *soap.Fault { return b.Fault_ }

func SubmitFeedbackForDiagnosisResult(ctx context.Context, r soap.RoundTripper, req *types.SubmitFeedbackForDiagnosisResult) (*types.SubmitFeedbackForDiagnosisResultResponse, error) {
	var reqBody, resBody SubmitFeedbackForDiagnosisResultBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanPerfLoginBody struct {
	Req    *types.VsanPerfLogin         `xml:"urn:vsan VsanPerfLogin,omitempty"`
	Res    *types.VsanPerfLoginResponse `xml:"urn:vsan VsanPerfLoginResponse,omitempty"`
	Fault_ *soap.Fault                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanPerfLoginBody) Fault() *soap.Fault { return b.Fault_ }

func VsanPerfLogin(ctx context.Context, r soap.RoundTripper, req *types.VsanPerfLogin) (*types.VsanPerfLoginResponse, error) {
	var reqBody, resBody VsanPerfLoginBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanPerfQueryTimeRangesBody struct {
	Req    *types.VsanPerfQueryTimeRanges         `xml:"urn:vsan VsanPerfQueryTimeRanges,omitempty"`
	Res    *types.VsanPerfQueryTimeRangesResponse `xml:"urn:vsan VsanPerfQueryTimeRangesResponse,omitempty"`
	Fault_ *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanPerfQueryTimeRangesBody) Fault() *soap.Fault { return b.Fault_ }

func VsanPerfQueryTimeRanges(ctx context.Context, r soap.RoundTripper, req *types.VsanPerfQueryTimeRanges) (*types.VsanPerfQueryTimeRangesResponse, error) {
	var reqBody, resBody VsanPerfQueryTimeRangesBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanPerfSetStatsObjectPolicyBody struct {
	Req    *types.VsanPerfSetStatsObjectPolicy         `xml:"urn:vsan VsanPerfSetStatsObjectPolicy,omitempty"`
	Res    *types.VsanPerfSetStatsObjectPolicyResponse `xml:"urn:vsan VsanPerfSetStatsObjectPolicyResponse,omitempty"`
	Fault_ *soap.Fault                                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanPerfSetStatsObjectPolicyBody) Fault() *soap.Fault { return b.Fault_ }

func VsanPerfSetStatsObjectPolicy(ctx context.Context, r soap.RoundTripper, req *types.VsanPerfSetStatsObjectPolicy) (*types.VsanPerfSetStatsObjectPolicyResponse, error) {
	var reqBody, resBody VsanPerfSetStatsObjectPolicyBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanPerfStartStatsCollectorBody struct {
	Req    *types.VsanPerfStartStatsCollector         `xml:"urn:vsan VsanPerfStartStatsCollector,omitempty"`
	Res    *types.VsanPerfStartStatsCollectorResponse `xml:"urn:vsan VsanPerfStartStatsCollectorResponse,omitempty"`
	Fault_ *soap.Fault                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanPerfStartStatsCollectorBody) Fault() *soap.Fault { return b.Fault_ }

func VsanPerfStartStatsCollector(ctx context.Context, r soap.RoundTripper, req *types.VsanPerfStartStatsCollector) (*types.VsanPerfStartStatsCollectorResponse, error) {
	var reqBody, resBody VsanPerfStartStatsCollectorBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanPerfInjectFakeHistoryBody struct {
	Req    *types.VsanPerfInjectFakeHistory         `xml:"urn:vsan VsanPerfInjectFakeHistory,omitempty"`
	Res    *types.VsanPerfInjectFakeHistoryResponse `xml:"urn:vsan VsanPerfInjectFakeHistoryResponse,omitempty"`
	Fault_ *soap.Fault                              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanPerfInjectFakeHistoryBody) Fault() *soap.Fault { return b.Fault_ }

func VsanPerfInjectFakeHistory(ctx context.Context, r soap.RoundTripper, req *types.VsanPerfInjectFakeHistory) (*types.VsanPerfInjectFakeHistoryResponse, error) {
	var reqBody, resBody VsanPerfInjectFakeHistoryBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanPerfCreateStatsObjectBody struct {
	Req    *types.VsanPerfCreateStatsObject         `xml:"urn:vsan VsanPerfCreateStatsObject,omitempty"`
	Res    *types.VsanPerfCreateStatsObjectResponse `xml:"urn:vsan VsanPerfCreateStatsObjectResponse,omitempty"`
	Fault_ *soap.Fault                              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanPerfCreateStatsObjectBody) Fault() *soap.Fault { return b.Fault_ }

func VsanPerfCreateStatsObject(ctx context.Context, r soap.RoundTripper, req *types.VsanPerfCreateStatsObject) (*types.VsanPerfCreateStatsObjectResponse, error) {
	var reqBody, resBody VsanPerfCreateStatsObjectBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanPerfQueryPerfBody struct {
	Req    *types.VsanPerfQueryPerf         `xml:"urn:vsan VsanPerfQueryPerf,omitempty"`
	Res    *types.VsanPerfQueryPerfResponse `xml:"urn:vsan VsanPerfQueryPerfResponse,omitempty"`
	Fault_ *soap.Fault                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanPerfQueryPerfBody) Fault() *soap.Fault { return b.Fault_ }

func VsanPerfQueryPerf(ctx context.Context, r soap.RoundTripper, req *types.VsanPerfQueryPerf) (*types.VsanPerfQueryPerfResponse, error) {
	var reqBody, resBody VsanPerfQueryPerfBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanPerfDiagnoseFromFileBody struct {
	Req    *types.VsanPerfDiagnoseFromFile         `xml:"urn:vsan VsanPerfDiagnoseFromFile,omitempty"`
	Res    *types.VsanPerfDiagnoseFromFileResponse `xml:"urn:vsan VsanPerfDiagnoseFromFileResponse,omitempty"`
	Fault_ *soap.Fault                             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanPerfDiagnoseFromFileBody) Fault() *soap.Fault { return b.Fault_ }

func VsanPerfDiagnoseFromFile(ctx context.Context, r soap.RoundTripper, req *types.VsanPerfDiagnoseFromFile) (*types.VsanPerfDiagnoseFromFileResponse, error) {
	var reqBody, resBody VsanPerfDiagnoseFromFileBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanPerfCreateStatsObjectTaskBody struct {
	Req    *types.VsanPerfCreateStatsObjectTask         `xml:"urn:vsan VsanPerfCreateStatsObjectTask,omitempty"`
	Res    *types.VsanPerfCreateStatsObjectTaskResponse `xml:"urn:vsan VsanPerfCreateStatsObjectTaskResponse,omitempty"`
	Fault_ *soap.Fault                                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanPerfCreateStatsObjectTaskBody) Fault() *soap.Fault { return b.Fault_ }

func VsanPerfCreateStatsObjectTask(ctx context.Context, r soap.RoundTripper, req *types.VsanPerfCreateStatsObjectTask) (*types.VsanPerfCreateStatsObjectTaskResponse, error) {
	var reqBody, resBody VsanPerfCreateStatsObjectTaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanPerfDeleteStatsObjectTaskBody struct {
	Req    *types.VsanPerfDeleteStatsObjectTask         `xml:"urn:vsan VsanPerfDeleteStatsObjectTask,omitempty"`
	Res    *types.VsanPerfDeleteStatsObjectTaskResponse `xml:"urn:vsan VsanPerfDeleteStatsObjectTaskResponse,omitempty"`
	Fault_ *soap.Fault                                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanPerfDeleteStatsObjectTaskBody) Fault() *soap.Fault { return b.Fault_ }

func VsanPerfDeleteStatsObjectTask(ctx context.Context, r soap.RoundTripper, req *types.VsanPerfDeleteStatsObjectTask) (*types.VsanPerfDeleteStatsObjectTaskResponse, error) {
	var reqBody, resBody VsanPerfDeleteStatsObjectTaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanPerfGetAggregatedEntityTypesBody struct {
	Req    *types.VsanPerfGetAggregatedEntityTypes         `xml:"urn:vsan VsanPerfGetAggregatedEntityTypes,omitempty"`
	Res    *types.VsanPerfGetAggregatedEntityTypesResponse `xml:"urn:vsan VsanPerfGetAggregatedEntityTypesResponse,omitempty"`
	Fault_ *soap.Fault                                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanPerfGetAggregatedEntityTypesBody) Fault() *soap.Fault { return b.Fault_ }

func VsanPerfGetAggregatedEntityTypes(ctx context.Context, r soap.RoundTripper, req *types.VsanPerfGetAggregatedEntityTypes) (*types.VsanPerfGetAggregatedEntityTypesResponse, error) {
	var reqBody, resBody VsanPerfGetAggregatedEntityTypesBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanPerfDeleteTimeRangeBody struct {
	Req    *types.VsanPerfDeleteTimeRange         `xml:"urn:vsan VsanPerfDeleteTimeRange,omitempty"`
	Res    *types.VsanPerfDeleteTimeRangeResponse `xml:"urn:vsan VsanPerfDeleteTimeRangeResponse,omitempty"`
	Fault_ *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanPerfDeleteTimeRangeBody) Fault() *soap.Fault { return b.Fault_ }

func VsanPerfDeleteTimeRange(ctx context.Context, r soap.RoundTripper, req *types.VsanPerfDeleteTimeRange) (*types.VsanPerfDeleteTimeRangeResponse, error) {
	var reqBody, resBody VsanPerfDeleteTimeRangeBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryVsanHostConfigInfoBody struct {
	Req    *types.QueryVsanHostConfigInfo         `xml:"urn:vsan QueryVsanHostConfigInfo,omitempty"`
	Res    *types.QueryVsanHostConfigInfoResponse `xml:"urn:vsan QueryVsanHostConfigInfoResponse,omitempty"`
	Fault_ *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryVsanHostConfigInfoBody) Fault() *soap.Fault { return b.Fault_ }

func QueryVsanHostConfigInfo(ctx context.Context, r soap.RoundTripper, req *types.QueryVsanHostConfigInfo) (*types.QueryVsanHostConfigInfoResponse, error) {
	var reqBody, resBody QueryVsanHostConfigInfoBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanPerfQueryStatsObjectInformationBody struct {
	Req    *types.VsanPerfQueryStatsObjectInformation         `xml:"urn:vsan VsanPerfQueryStatsObjectInformation,omitempty"`
	Res    *types.VsanPerfQueryStatsObjectInformationResponse `xml:"urn:vsan VsanPerfQueryStatsObjectInformationResponse,omitempty"`
	Fault_ *soap.Fault                                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanPerfQueryStatsObjectInformationBody) Fault() *soap.Fault { return b.Fault_ }

func VsanPerfQueryStatsObjectInformation(ctx context.Context, r soap.RoundTripper, req *types.VsanPerfQueryStatsObjectInformation) (*types.VsanPerfQueryStatsObjectInformationResponse, error) {
	var reqBody, resBody VsanPerfQueryStatsObjectInformationBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanPerfDeleteStatsObjectBody struct {
	Req    *types.VsanPerfDeleteStatsObject         `xml:"urn:vsan VsanPerfDeleteStatsObject,omitempty"`
	Res    *types.VsanPerfDeleteStatsObjectResponse `xml:"urn:vsan VsanPerfDeleteStatsObjectResponse,omitempty"`
	Fault_ *soap.Fault                              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanPerfDeleteStatsObjectBody) Fault() *soap.Fault { return b.Fault_ }

func VsanPerfDeleteStatsObject(ctx context.Context, r soap.RoundTripper, req *types.VsanPerfDeleteStatsObject) (*types.VsanPerfDeleteStatsObjectResponse, error) {
	var reqBody, resBody VsanPerfDeleteStatsObjectBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanPerfSaveTimeRangesBody struct {
	Req    *types.VsanPerfSaveTimeRanges         `xml:"urn:vsan VsanPerfSaveTimeRanges,omitempty"`
	Res    *types.VsanPerfSaveTimeRangesResponse `xml:"urn:vsan VsanPerfSaveTimeRangesResponse,omitempty"`
	Fault_ *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanPerfSaveTimeRangesBody) Fault() *soap.Fault { return b.Fault_ }

func VsanPerfSaveTimeRanges(ctx context.Context, r soap.RoundTripper, req *types.VsanPerfSaveTimeRanges) (*types.VsanPerfSaveTimeRangesResponse, error) {
	var reqBody, resBody VsanPerfSaveTimeRangesBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanPerfQueryNodeInformationBody struct {
	Req    *types.VsanPerfQueryNodeInformation         `xml:"urn:vsan VsanPerfQueryNodeInformation,omitempty"`
	Res    *types.VsanPerfQueryNodeInformationResponse `xml:"urn:vsan VsanPerfQueryNodeInformationResponse,omitempty"`
	Fault_ *soap.Fault                                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanPerfQueryNodeInformationBody) Fault() *soap.Fault { return b.Fault_ }

func VsanPerfQueryNodeInformation(ctx context.Context, r soap.RoundTripper, req *types.VsanPerfQueryNodeInformation) (*types.VsanPerfQueryNodeInformationResponse, error) {
	var reqBody, resBody VsanPerfQueryNodeInformationBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanPerfGetSupportedDiagnosticExceptionsBody struct {
	Req    *types.VsanPerfGetSupportedDiagnosticExceptions         `xml:"urn:vsan VsanPerfGetSupportedDiagnosticExceptions,omitempty"`
	Res    *types.VsanPerfGetSupportedDiagnosticExceptionsResponse `xml:"urn:vsan VsanPerfGetSupportedDiagnosticExceptionsResponse,omitempty"`
	Fault_ *soap.Fault                                             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanPerfGetSupportedDiagnosticExceptionsBody) Fault() *soap.Fault { return b.Fault_ }

func VsanPerfGetSupportedDiagnosticExceptions(ctx context.Context, r soap.RoundTripper, req *types.VsanPerfGetSupportedDiagnosticExceptions) (*types.VsanPerfGetSupportedDiagnosticExceptionsResponse, error) {
	var reqBody, resBody VsanPerfGetSupportedDiagnosticExceptionsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanPerfGetVcMoRefFromPerfEntityRefIdBody struct {
	Req    *types.VsanPerfGetVcMoRefFromPerfEntityRefId         `xml:"urn:vsan VsanPerfGetVcMoRefFromPerfEntityRefId,omitempty"`
	Res    *types.VsanPerfGetVcMoRefFromPerfEntityRefIdResponse `xml:"urn:vsan VsanPerfGetVcMoRefFromPerfEntityRefIdResponse,omitempty"`
	Fault_ *soap.Fault                                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanPerfGetVcMoRefFromPerfEntityRefIdBody) Fault() *soap.Fault { return b.Fault_ }

func VsanPerfGetVcMoRefFromPerfEntityRefId(ctx context.Context, r soap.RoundTripper, req *types.VsanPerfGetVcMoRefFromPerfEntityRefId) (*types.VsanPerfGetVcMoRefFromPerfEntityRefIdResponse, error) {
	var reqBody, resBody VsanPerfGetVcMoRefFromPerfEntityRefIdBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanPerfGetSupportedEntityTypesBody struct {
	Req    *types.VsanPerfGetSupportedEntityTypes         `xml:"urn:vsan VsanPerfGetSupportedEntityTypes,omitempty"`
	Res    *types.VsanPerfGetSupportedEntityTypesResponse `xml:"urn:vsan VsanPerfGetSupportedEntityTypesResponse,omitempty"`
	Fault_ *soap.Fault                                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanPerfGetSupportedEntityTypesBody) Fault() *soap.Fault { return b.Fault_ }

func VsanPerfGetSupportedEntityTypes(ctx context.Context, r soap.RoundTripper, req *types.VsanPerfGetSupportedEntityTypes) (*types.VsanPerfGetSupportedEntityTypesResponse, error) {
	var reqBody, resBody VsanPerfGetSupportedEntityTypesBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type InitializeDiskMappingsBody struct {
	Req    *types.InitializeDiskMappings         `xml:"urn:vsan InitializeDiskMappings,omitempty"`
	Res    *types.InitializeDiskMappingsResponse `xml:"urn:vsan InitializeDiskMappingsResponse,omitempty"`
	Fault_ *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *InitializeDiskMappingsBody) Fault() *soap.Fault { return b.Fault_ }

func InitializeDiskMappings(ctx context.Context, r soap.RoundTripper, req *types.InitializeDiskMappings) (*types.InitializeDiskMappingsResponse, error) {
	var reqBody, resBody InitializeDiskMappingsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryDiskMappingsBody struct {
	Req    *types.QueryDiskMappings         `xml:"urn:vsan QueryDiskMappings,omitempty"`
	Res    *types.QueryDiskMappingsResponse `xml:"urn:vsan QueryDiskMappingsResponse,omitempty"`
	Fault_ *soap.Fault                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryDiskMappingsBody) Fault() *soap.Fault { return b.Fault_ }

func QueryDiskMappings(ctx context.Context, r soap.RoundTripper, req *types.QueryDiskMappings) (*types.QueryDiskMappingsResponse, error) {
	var reqBody, resBody QueryDiskMappingsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryClusterDataEfficiencyCapacityStateBody struct {
	Req    *types.QueryClusterDataEfficiencyCapacityState         `xml:"urn:vsan QueryClusterDataEfficiencyCapacityState,omitempty"`
	Res    *types.QueryClusterDataEfficiencyCapacityStateResponse `xml:"urn:vsan QueryClusterDataEfficiencyCapacityStateResponse,omitempty"`
	Fault_ *soap.Fault                                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryClusterDataEfficiencyCapacityStateBody) Fault() *soap.Fault { return b.Fault_ }

func QueryClusterDataEfficiencyCapacityState(ctx context.Context, r soap.RoundTripper, req *types.QueryClusterDataEfficiencyCapacityState) (*types.QueryClusterDataEfficiencyCapacityStateResponse, error) {
	var reqBody, resBody QueryClusterDataEfficiencyCapacityStateBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RetrieveAllFlashCapabilityBody struct {
	Req    *types.RetrieveAllFlashCapability         `xml:"urn:vsan RetrieveAllFlashCapability,omitempty"`
	Res    *types.RetrieveAllFlashCapabilityResponse `xml:"urn:vsan RetrieveAllFlashCapabilityResponse,omitempty"`
	Fault_ *soap.Fault                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RetrieveAllFlashCapabilityBody) Fault() *soap.Fault { return b.Fault_ }

func RetrieveAllFlashCapability(ctx context.Context, r soap.RoundTripper, req *types.RetrieveAllFlashCapability) (*types.RetrieveAllFlashCapabilityResponse, error) {
	var reqBody, resBody RetrieveAllFlashCapabilityBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RetrieveAllFlashCapabilitiesBody struct {
	Req    *types.RetrieveAllFlashCapabilities         `xml:"urn:vsan RetrieveAllFlashCapabilities,omitempty"`
	Res    *types.RetrieveAllFlashCapabilitiesResponse `xml:"urn:vsan RetrieveAllFlashCapabilitiesResponse,omitempty"`
	Fault_ *soap.Fault                                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RetrieveAllFlashCapabilitiesBody) Fault() *soap.Fault { return b.Fault_ }

func RetrieveAllFlashCapabilities(ctx context.Context, r soap.RoundTripper, req *types.RetrieveAllFlashCapabilities) (*types.RetrieveAllFlashCapabilitiesResponse, error) {
	var reqBody, resBody RetrieveAllFlashCapabilitiesBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RebuildDiskMappingBody struct {
	Req    *types.RebuildDiskMapping         `xml:"urn:vsan RebuildDiskMapping,omitempty"`
	Res    *types.RebuildDiskMappingResponse `xml:"urn:vsan RebuildDiskMappingResponse,omitempty"`
	Fault_ *soap.Fault                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RebuildDiskMappingBody) Fault() *soap.Fault { return b.Fault_ }

func RebuildDiskMapping(ctx context.Context, r soap.RoundTripper, req *types.RebuildDiskMapping) (*types.RebuildDiskMappingResponse, error) {
	var reqBody, resBody RebuildDiskMappingBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanClusterMgmtPersistStressOptionsBody struct {
	Req    *types.VsanClusterMgmtPersistStressOptions         `xml:"urn:vsan VsanClusterMgmtPersistStressOptions,omitempty"`
	Res    *types.VsanClusterMgmtPersistStressOptionsResponse `xml:"urn:vsan VsanClusterMgmtPersistStressOptionsResponse,omitempty"`
	Fault_ *soap.Fault                                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanClusterMgmtPersistStressOptionsBody) Fault() *soap.Fault { return b.Fault_ }

func VsanClusterMgmtPersistStressOptions(ctx context.Context, r soap.RoundTripper, req *types.VsanClusterMgmtPersistStressOptions) (*types.VsanClusterMgmtPersistStressOptionsResponse, error) {
	var reqBody, resBody VsanClusterMgmtPersistStressOptionsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanClusterMgmtClearStressOptionsBody struct {
	Req    *types.VsanClusterMgmtClearStressOptions         `xml:"urn:vsan VsanClusterMgmtClearStressOptions,omitempty"`
	Res    *types.VsanClusterMgmtClearStressOptionsResponse `xml:"urn:vsan VsanClusterMgmtClearStressOptionsResponse,omitempty"`
	Fault_ *soap.Fault                                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanClusterMgmtClearStressOptionsBody) Fault() *soap.Fault { return b.Fault_ }

func VsanClusterMgmtClearStressOptions(ctx context.Context, r soap.RoundTripper, req *types.VsanClusterMgmtClearStressOptions) (*types.VsanClusterMgmtClearStressOptionsResponse, error) {
	var reqBody, resBody VsanClusterMgmtClearStressOptionsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanClusterMgmtServiceRestartBody struct {
	Req    *types.VsanClusterMgmtServiceRestart         `xml:"urn:vsan VsanClusterMgmtServiceRestart,omitempty"`
	Res    *types.VsanClusterMgmtServiceRestartResponse `xml:"urn:vsan VsanClusterMgmtServiceRestartResponse,omitempty"`
	Fault_ *soap.Fault                                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanClusterMgmtServiceRestartBody) Fault() *soap.Fault { return b.Fault_ }

func VsanClusterMgmtServiceRestart(ctx context.Context, r soap.RoundTripper, req *types.VsanClusterMgmtServiceRestart) (*types.VsanClusterMgmtServiceRestartResponse, error) {
	var reqBody, resBody VsanClusterMgmtServiceRestartBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanGetPersistedClusterStateBody struct {
	Req    *types.VsanGetPersistedClusterState         `xml:"urn:vsan VsanGetPersistedClusterState,omitempty"`
	Res    *types.VsanGetPersistedClusterStateResponse `xml:"urn:vsan VsanGetPersistedClusterStateResponse,omitempty"`
	Fault_ *soap.Fault                                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanGetPersistedClusterStateBody) Fault() *soap.Fault { return b.Fault_ }

func VsanGetPersistedClusterState(ctx context.Context, r soap.RoundTripper, req *types.VsanGetPersistedClusterState) (*types.VsanGetPersistedClusterStateResponse, error) {
	var reqBody, resBody VsanGetPersistedClusterStateBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanClusterMgmtGetStatsBody struct {
	Req    *types.VsanClusterMgmtGetStats         `xml:"urn:vsan VsanClusterMgmtGetStats,omitempty"`
	Res    *types.VsanClusterMgmtGetStatsResponse `xml:"urn:vsan VsanClusterMgmtGetStatsResponse,omitempty"`
	Fault_ *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanClusterMgmtGetStatsBody) Fault() *soap.Fault { return b.Fault_ }

func VsanClusterMgmtGetStats(ctx context.Context, r soap.RoundTripper, req *types.VsanClusterMgmtGetStats) (*types.VsanClusterMgmtGetStatsResponse, error) {
	var reqBody, resBody VsanClusterMgmtGetStatsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanClusterMgmtSetStressOptionBody struct {
	Req    *types.VsanClusterMgmtSetStressOption         `xml:"urn:vsan VsanClusterMgmtSetStressOption,omitempty"`
	Res    *types.VsanClusterMgmtSetStressOptionResponse `xml:"urn:vsan VsanClusterMgmtSetStressOptionResponse,omitempty"`
	Fault_ *soap.Fault                                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanClusterMgmtSetStressOptionBody) Fault() *soap.Fault { return b.Fault_ }

func VsanClusterMgmtSetStressOption(ctx context.Context, r soap.RoundTripper, req *types.VsanClusterMgmtSetStressOption) (*types.VsanClusterMgmtSetStressOptionResponse, error) {
	var reqBody, resBody VsanClusterMgmtSetStressOptionBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanClusterMgmtPurgeAllPersistedStateBody struct {
	Req    *types.VsanClusterMgmtPurgeAllPersistedState         `xml:"urn:vsan VsanClusterMgmtPurgeAllPersistedState,omitempty"`
	Res    *types.VsanClusterMgmtPurgeAllPersistedStateResponse `xml:"urn:vsan VsanClusterMgmtPurgeAllPersistedStateResponse,omitempty"`
	Fault_ *soap.Fault                                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanClusterMgmtPurgeAllPersistedStateBody) Fault() *soap.Fault { return b.Fault_ }

func VsanClusterMgmtPurgeAllPersistedState(ctx context.Context, r soap.RoundTripper, req *types.VsanClusterMgmtPurgeAllPersistedState) (*types.VsanClusterMgmtPurgeAllPersistedStateResponse, error) {
	var reqBody, resBody VsanClusterMgmtPurgeAllPersistedStateBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanCompleteMigrateVmsToVdsBody struct {
	Req    *types.VsanCompleteMigrateVmsToVds         `xml:"urn:vsan VsanCompleteMigrateVmsToVds,omitempty"`
	Res    *types.VsanCompleteMigrateVmsToVdsResponse `xml:"urn:vsan VsanCompleteMigrateVmsToVdsResponse,omitempty"`
	Fault_ *soap.Fault                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanCompleteMigrateVmsToVdsBody) Fault() *soap.Fault { return b.Fault_ }

func VsanCompleteMigrateVmsToVds(ctx context.Context, r soap.RoundTripper, req *types.VsanCompleteMigrateVmsToVds) (*types.VsanCompleteMigrateVmsToVdsResponse, error) {
	var reqBody, resBody VsanCompleteMigrateVmsToVdsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanMigrateVmsToVdsBody struct {
	Req    *types.VsanMigrateVmsToVds         `xml:"urn:vsan VsanMigrateVmsToVds,omitempty"`
	Res    *types.VsanMigrateVmsToVdsResponse `xml:"urn:vsan VsanMigrateVmsToVdsResponse,omitempty"`
	Fault_ *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanMigrateVmsToVdsBody) Fault() *soap.Fault { return b.Fault_ }

func VsanMigrateVmsToVds(ctx context.Context, r soap.RoundTripper, req *types.VsanMigrateVmsToVds) (*types.VsanMigrateVmsToVdsResponse, error) {
	var reqBody, resBody VsanMigrateVmsToVdsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VSANHostGetStretchedClusterInfoFromCmmdsBody struct {
	Req    *types.VSANHostGetStretchedClusterInfoFromCmmds         `xml:"urn:vsan VSANHostGetStretchedClusterInfoFromCmmds,omitempty"`
	Res    *types.VSANHostGetStretchedClusterInfoFromCmmdsResponse `xml:"urn:vsan VSANHostGetStretchedClusterInfoFromCmmdsResponse,omitempty"`
	Fault_ *soap.Fault                                             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VSANHostGetStretchedClusterInfoFromCmmdsBody) Fault() *soap.Fault { return b.Fault_ }

func VSANHostGetStretchedClusterInfoFromCmmds(ctx context.Context, r soap.RoundTripper, req *types.VSANHostGetStretchedClusterInfoFromCmmds) (*types.VSANHostGetStretchedClusterInfoFromCmmdsResponse, error) {
	var reqBody, resBody VSANHostGetStretchedClusterInfoFromCmmdsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VSANWitnessJoinVsanClusterBody struct {
	Req    *types.VSANWitnessJoinVsanCluster         `xml:"urn:vsan VSANWitnessJoinVsanCluster,omitempty"`
	Res    *types.VSANWitnessJoinVsanClusterResponse `xml:"urn:vsan VSANWitnessJoinVsanClusterResponse,omitempty"`
	Fault_ *soap.Fault                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VSANWitnessJoinVsanClusterBody) Fault() *soap.Fault { return b.Fault_ }

func VSANWitnessJoinVsanCluster(ctx context.Context, r soap.RoundTripper, req *types.VSANWitnessJoinVsanCluster) (*types.VSANWitnessJoinVsanClusterResponse, error) {
	var reqBody, resBody VSANWitnessJoinVsanClusterBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VSANWitnessSetPreferredFaultDomainBody struct {
	Req    *types.VSANWitnessSetPreferredFaultDomain         `xml:"urn:vsan VSANWitnessSetPreferredFaultDomain,omitempty"`
	Res    *types.VSANWitnessSetPreferredFaultDomainResponse `xml:"urn:vsan VSANWitnessSetPreferredFaultDomainResponse,omitempty"`
	Fault_ *soap.Fault                                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VSANWitnessSetPreferredFaultDomainBody) Fault() *soap.Fault { return b.Fault_ }

func VSANWitnessSetPreferredFaultDomain(ctx context.Context, r soap.RoundTripper, req *types.VSANWitnessSetPreferredFaultDomain) (*types.VSANWitnessSetPreferredFaultDomainResponse, error) {
	var reqBody, resBody VSANWitnessSetPreferredFaultDomainBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VSANHostAddUnicastAgentBody struct {
	Req    *types.VSANHostAddUnicastAgent         `xml:"urn:vsan VSANHostAddUnicastAgent,omitempty"`
	Res    *types.VSANHostAddUnicastAgentResponse `xml:"urn:vsan VSANHostAddUnicastAgentResponse,omitempty"`
	Fault_ *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VSANHostAddUnicastAgentBody) Fault() *soap.Fault { return b.Fault_ }

func VSANHostAddUnicastAgent(ctx context.Context, r soap.RoundTripper, req *types.VSANHostAddUnicastAgent) (*types.VSANHostAddUnicastAgentResponse, error) {
	var reqBody, resBody VSANHostAddUnicastAgentBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VSANClusterGetPreferredFaultDomainBody struct {
	Req    *types.VSANClusterGetPreferredFaultDomain         `xml:"urn:vsan VSANClusterGetPreferredFaultDomain,omitempty"`
	Res    *types.VSANClusterGetPreferredFaultDomainResponse `xml:"urn:vsan VSANClusterGetPreferredFaultDomainResponse,omitempty"`
	Fault_ *soap.Fault                                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VSANClusterGetPreferredFaultDomainBody) Fault() *soap.Fault { return b.Fault_ }

func VSANClusterGetPreferredFaultDomain(ctx context.Context, r soap.RoundTripper, req *types.VSANClusterGetPreferredFaultDomain) (*types.VSANClusterGetPreferredFaultDomainResponse, error) {
	var reqBody, resBody VSANClusterGetPreferredFaultDomainBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VSANWitnessLeaveVsanClusterBody struct {
	Req    *types.VSANWitnessLeaveVsanCluster         `xml:"urn:vsan VSANWitnessLeaveVsanCluster,omitempty"`
	Res    *types.VSANWitnessLeaveVsanClusterResponse `xml:"urn:vsan VSANWitnessLeaveVsanClusterResponse,omitempty"`
	Fault_ *soap.Fault                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VSANWitnessLeaveVsanClusterBody) Fault() *soap.Fault { return b.Fault_ }

func VSANWitnessLeaveVsanCluster(ctx context.Context, r soap.RoundTripper, req *types.VSANWitnessLeaveVsanCluster) (*types.VSANWitnessLeaveVsanClusterResponse, error) {
	var reqBody, resBody VSANWitnessLeaveVsanClusterBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VSANHostGetStretchedClusterCapabilityBody struct {
	Req    *types.VSANHostGetStretchedClusterCapability         `xml:"urn:vsan VSANHostGetStretchedClusterCapability,omitempty"`
	Res    *types.VSANHostGetStretchedClusterCapabilityResponse `xml:"urn:vsan VSANHostGetStretchedClusterCapabilityResponse,omitempty"`
	Fault_ *soap.Fault                                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VSANHostGetStretchedClusterCapabilityBody) Fault() *soap.Fault { return b.Fault_ }

func VSANHostGetStretchedClusterCapability(ctx context.Context, r soap.RoundTripper, req *types.VSANHostGetStretchedClusterCapability) (*types.VSANHostGetStretchedClusterCapabilityResponse, error) {
	var reqBody, resBody VSANHostGetStretchedClusterCapabilityBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VSANHostRemoveUnicastAgentBody struct {
	Req    *types.VSANHostRemoveUnicastAgent         `xml:"urn:vsan VSANHostRemoveUnicastAgent,omitempty"`
	Res    *types.VSANHostRemoveUnicastAgentResponse `xml:"urn:vsan VSANHostRemoveUnicastAgentResponse,omitempty"`
	Fault_ *soap.Fault                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VSANHostRemoveUnicastAgentBody) Fault() *soap.Fault { return b.Fault_ }

func VSANHostRemoveUnicastAgent(ctx context.Context, r soap.RoundTripper, req *types.VSANHostRemoveUnicastAgent) (*types.VSANHostRemoveUnicastAgentResponse, error) {
	var reqBody, resBody VSANHostRemoveUnicastAgentBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VSANHostListUnicastAgentBody struct {
	Req    *types.VSANHostListUnicastAgent         `xml:"urn:vsan VSANHostListUnicastAgent,omitempty"`
	Res    *types.VSANHostListUnicastAgentResponse `xml:"urn:vsan VSANHostListUnicastAgentResponse,omitempty"`
	Fault_ *soap.Fault                             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VSANHostListUnicastAgentBody) Fault() *soap.Fault { return b.Fault_ }

func VSANHostListUnicastAgent(ctx context.Context, r soap.RoundTripper, req *types.VSANHostListUnicastAgent) (*types.VSANHostListUnicastAgentResponse, error) {
	var reqBody, resBody VSANHostListUnicastAgentBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanVitGetIscsiTargetServiceStatusBody struct {
	Req    *types.VsanVitGetIscsiTargetServiceStatus         `xml:"urn:vsan VsanVitGetIscsiTargetServiceStatus,omitempty"`
	Res    *types.VsanVitGetIscsiTargetServiceStatusResponse `xml:"urn:vsan VsanVitGetIscsiTargetServiceStatusResponse,omitempty"`
	Fault_ *soap.Fault                                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanVitGetIscsiTargetServiceStatusBody) Fault() *soap.Fault { return b.Fault_ }

func VsanVitGetIscsiTargetServiceStatus(ctx context.Context, r soap.RoundTripper, req *types.VsanVitGetIscsiTargetServiceStatus) (*types.VsanVitGetIscsiTargetServiceStatusResponse, error) {
	var reqBody, resBody VsanVitGetIscsiTargetServiceStatusBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanVitRemoveIscsiInitiatorsFromTargetBody struct {
	Req    *types.VsanVitRemoveIscsiInitiatorsFromTarget         `xml:"urn:vsan VsanVitRemoveIscsiInitiatorsFromTarget,omitempty"`
	Res    *types.VsanVitRemoveIscsiInitiatorsFromTargetResponse `xml:"urn:vsan VsanVitRemoveIscsiInitiatorsFromTargetResponse,omitempty"`
	Fault_ *soap.Fault                                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanVitRemoveIscsiInitiatorsFromTargetBody) Fault() *soap.Fault { return b.Fault_ }

func VsanVitRemoveIscsiInitiatorsFromTarget(ctx context.Context, r soap.RoundTripper, req *types.VsanVitRemoveIscsiInitiatorsFromTarget) (*types.VsanVitRemoveIscsiInitiatorsFromTargetResponse, error) {
	var reqBody, resBody VsanVitRemoveIscsiInitiatorsFromTargetBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanVitRemoveIscsiInitiatorsFromGroupBody struct {
	Req    *types.VsanVitRemoveIscsiInitiatorsFromGroup         `xml:"urn:vsan VsanVitRemoveIscsiInitiatorsFromGroup,omitempty"`
	Res    *types.VsanVitRemoveIscsiInitiatorsFromGroupResponse `xml:"urn:vsan VsanVitRemoveIscsiInitiatorsFromGroupResponse,omitempty"`
	Fault_ *soap.Fault                                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanVitRemoveIscsiInitiatorsFromGroupBody) Fault() *soap.Fault { return b.Fault_ }

func VsanVitRemoveIscsiInitiatorsFromGroup(ctx context.Context, r soap.RoundTripper, req *types.VsanVitRemoveIscsiInitiatorsFromGroup) (*types.VsanVitRemoveIscsiInitiatorsFromGroupResponse, error) {
	var reqBody, resBody VsanVitRemoveIscsiInitiatorsFromGroupBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanVitEditIscsiLUNBody struct {
	Req    *types.VsanVitEditIscsiLUN         `xml:"urn:vsan VsanVitEditIscsiLUN,omitempty"`
	Res    *types.VsanVitEditIscsiLUNResponse `xml:"urn:vsan VsanVitEditIscsiLUNResponse,omitempty"`
	Fault_ *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanVitEditIscsiLUNBody) Fault() *soap.Fault { return b.Fault_ }

func VsanVitEditIscsiLUN(ctx context.Context, r soap.RoundTripper, req *types.VsanVitEditIscsiLUN) (*types.VsanVitEditIscsiLUNResponse, error) {
	var reqBody, resBody VsanVitEditIscsiLUNBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanVitGetIscsiTargetServiceProcessesStatusBody struct {
	Req    *types.VsanVitGetIscsiTargetServiceProcessesStatus         `xml:"urn:vsan VsanVitGetIscsiTargetServiceProcessesStatus,omitempty"`
	Res    *types.VsanVitGetIscsiTargetServiceProcessesStatusResponse `xml:"urn:vsan VsanVitGetIscsiTargetServiceProcessesStatusResponse,omitempty"`
	Fault_ *soap.Fault                                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanVitGetIscsiTargetServiceProcessesStatusBody) Fault() *soap.Fault { return b.Fault_ }

func VsanVitGetIscsiTargetServiceProcessesStatus(ctx context.Context, r soap.RoundTripper, req *types.VsanVitGetIscsiTargetServiceProcessesStatus) (*types.VsanVitGetIscsiTargetServiceProcessesStatusResponse, error) {
	var reqBody, resBody VsanVitGetIscsiTargetServiceProcessesStatusBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanVitGetIscsiLUNBody struct {
	Req    *types.VsanVitGetIscsiLUN         `xml:"urn:vsan VsanVitGetIscsiLUN,omitempty"`
	Res    *types.VsanVitGetIscsiLUNResponse `xml:"urn:vsan VsanVitGetIscsiLUNResponse,omitempty"`
	Fault_ *soap.Fault                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanVitGetIscsiLUNBody) Fault() *soap.Fault { return b.Fault_ }

func VsanVitGetIscsiLUN(ctx context.Context, r soap.RoundTripper, req *types.VsanVitGetIscsiLUN) (*types.VsanVitGetIscsiLUNResponse, error) {
	var reqBody, resBody VsanVitGetIscsiLUNBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanVitEditIscsiTargetBody struct {
	Req    *types.VsanVitEditIscsiTarget         `xml:"urn:vsan VsanVitEditIscsiTarget,omitempty"`
	Res    *types.VsanVitEditIscsiTargetResponse `xml:"urn:vsan VsanVitEditIscsiTargetResponse,omitempty"`
	Fault_ *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanVitEditIscsiTargetBody) Fault() *soap.Fault { return b.Fault_ }

func VsanVitEditIscsiTarget(ctx context.Context, r soap.RoundTripper, req *types.VsanVitEditIscsiTarget) (*types.VsanVitEditIscsiTargetResponse, error) {
	var reqBody, resBody VsanVitEditIscsiTargetBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanVitAddIscsiInitiatorsToGroupBody struct {
	Req    *types.VsanVitAddIscsiInitiatorsToGroup         `xml:"urn:vsan VsanVitAddIscsiInitiatorsToGroup,omitempty"`
	Res    *types.VsanVitAddIscsiInitiatorsToGroupResponse `xml:"urn:vsan VsanVitAddIscsiInitiatorsToGroupResponse,omitempty"`
	Fault_ *soap.Fault                                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanVitAddIscsiInitiatorsToGroupBody) Fault() *soap.Fault { return b.Fault_ }

func VsanVitAddIscsiInitiatorsToGroup(ctx context.Context, r soap.RoundTripper, req *types.VsanVitAddIscsiInitiatorsToGroup) (*types.VsanVitAddIscsiInitiatorsToGroupResponse, error) {
	var reqBody, resBody VsanVitAddIscsiInitiatorsToGroupBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanVitAddIscsiInitiatorsToTargetBody struct {
	Req    *types.VsanVitAddIscsiInitiatorsToTarget         `xml:"urn:vsan VsanVitAddIscsiInitiatorsToTarget,omitempty"`
	Res    *types.VsanVitAddIscsiInitiatorsToTargetResponse `xml:"urn:vsan VsanVitAddIscsiInitiatorsToTargetResponse,omitempty"`
	Fault_ *soap.Fault                                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanVitAddIscsiInitiatorsToTargetBody) Fault() *soap.Fault { return b.Fault_ }

func VsanVitAddIscsiInitiatorsToTarget(ctx context.Context, r soap.RoundTripper, req *types.VsanVitAddIscsiInitiatorsToTarget) (*types.VsanVitAddIscsiInitiatorsToTargetResponse, error) {
	var reqBody, resBody VsanVitAddIscsiInitiatorsToTargetBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanVitQueryIscsiTargetServiceVersionBody struct {
	Req    *types.VsanVitQueryIscsiTargetServiceVersion         `xml:"urn:vsan VsanVitQueryIscsiTargetServiceVersion,omitempty"`
	Res    *types.VsanVitQueryIscsiTargetServiceVersionResponse `xml:"urn:vsan VsanVitQueryIscsiTargetServiceVersionResponse,omitempty"`
	Fault_ *soap.Fault                                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanVitQueryIscsiTargetServiceVersionBody) Fault() *soap.Fault { return b.Fault_ }

func VsanVitQueryIscsiTargetServiceVersion(ctx context.Context, r soap.RoundTripper, req *types.VsanVitQueryIscsiTargetServiceVersion) (*types.VsanVitQueryIscsiTargetServiceVersionResponse, error) {
	var reqBody, resBody VsanVitQueryIscsiTargetServiceVersionBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanVitAddIscsiTargetToGroupBody struct {
	Req    *types.VsanVitAddIscsiTargetToGroup         `xml:"urn:vsan VsanVitAddIscsiTargetToGroup,omitempty"`
	Res    *types.VsanVitAddIscsiTargetToGroupResponse `xml:"urn:vsan VsanVitAddIscsiTargetToGroupResponse,omitempty"`
	Fault_ *soap.Fault                                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanVitAddIscsiTargetToGroupBody) Fault() *soap.Fault { return b.Fault_ }

func VsanVitAddIscsiTargetToGroup(ctx context.Context, r soap.RoundTripper, req *types.VsanVitAddIscsiTargetToGroup) (*types.VsanVitAddIscsiTargetToGroupResponse, error) {
	var reqBody, resBody VsanVitAddIscsiTargetToGroupBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanVitRemoveIscsiTargetFromGroupBody struct {
	Req    *types.VsanVitRemoveIscsiTargetFromGroup         `xml:"urn:vsan VsanVitRemoveIscsiTargetFromGroup,omitempty"`
	Res    *types.VsanVitRemoveIscsiTargetFromGroupResponse `xml:"urn:vsan VsanVitRemoveIscsiTargetFromGroupResponse,omitempty"`
	Fault_ *soap.Fault                                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanVitRemoveIscsiTargetFromGroupBody) Fault() *soap.Fault { return b.Fault_ }

func VsanVitRemoveIscsiTargetFromGroup(ctx context.Context, r soap.RoundTripper, req *types.VsanVitRemoveIscsiTargetFromGroup) (*types.VsanVitRemoveIscsiTargetFromGroupResponse, error) {
	var reqBody, resBody VsanVitRemoveIscsiTargetFromGroupBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanVitGetIscsiLUNsBody struct {
	Req    *types.VsanVitGetIscsiLUNs         `xml:"urn:vsan VsanVitGetIscsiLUNs,omitempty"`
	Res    *types.VsanVitGetIscsiLUNsResponse `xml:"urn:vsan VsanVitGetIscsiLUNsResponse,omitempty"`
	Fault_ *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanVitGetIscsiLUNsBody) Fault() *soap.Fault { return b.Fault_ }

func VsanVitGetIscsiLUNs(ctx context.Context, r soap.RoundTripper, req *types.VsanVitGetIscsiLUNs) (*types.VsanVitGetIscsiLUNsResponse, error) {
	var reqBody, resBody VsanVitGetIscsiLUNsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanVitRemoveIscsiLUNBody struct {
	Req    *types.VsanVitRemoveIscsiLUN         `xml:"urn:vsan VsanVitRemoveIscsiLUN,omitempty"`
	Res    *types.VsanVitRemoveIscsiLUNResponse `xml:"urn:vsan VsanVitRemoveIscsiLUNResponse,omitempty"`
	Fault_ *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanVitRemoveIscsiLUNBody) Fault() *soap.Fault { return b.Fault_ }

func VsanVitRemoveIscsiLUN(ctx context.Context, r soap.RoundTripper, req *types.VsanVitRemoveIscsiLUN) (*types.VsanVitRemoveIscsiLUNResponse, error) {
	var reqBody, resBody VsanVitRemoveIscsiLUNBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanVitGetIscsiInitiatorGroupBody struct {
	Req    *types.VsanVitGetIscsiInitiatorGroup         `xml:"urn:vsan VsanVitGetIscsiInitiatorGroup,omitempty"`
	Res    *types.VsanVitGetIscsiInitiatorGroupResponse `xml:"urn:vsan VsanVitGetIscsiInitiatorGroupResponse,omitempty"`
	Fault_ *soap.Fault                                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanVitGetIscsiInitiatorGroupBody) Fault() *soap.Fault { return b.Fault_ }

func VsanVitGetIscsiInitiatorGroup(ctx context.Context, r soap.RoundTripper, req *types.VsanVitGetIscsiInitiatorGroup) (*types.VsanVitGetIscsiInitiatorGroupResponse, error) {
	var reqBody, resBody VsanVitGetIscsiInitiatorGroupBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanVitEditIscsiTargetServiceBody struct {
	Req    *types.VsanVitEditIscsiTargetService         `xml:"urn:vsan VsanVitEditIscsiTargetService,omitempty"`
	Res    *types.VsanVitEditIscsiTargetServiceResponse `xml:"urn:vsan VsanVitEditIscsiTargetServiceResponse,omitempty"`
	Fault_ *soap.Fault                                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanVitEditIscsiTargetServiceBody) Fault() *soap.Fault { return b.Fault_ }

func VsanVitEditIscsiTargetService(ctx context.Context, r soap.RoundTripper, req *types.VsanVitEditIscsiTargetService) (*types.VsanVitEditIscsiTargetServiceResponse, error) {
	var reqBody, resBody VsanVitEditIscsiTargetServiceBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanVitRemoveIscsiInitiatorGroupBody struct {
	Req    *types.VsanVitRemoveIscsiInitiatorGroup         `xml:"urn:vsan VsanVitRemoveIscsiInitiatorGroup,omitempty"`
	Res    *types.VsanVitRemoveIscsiInitiatorGroupResponse `xml:"urn:vsan VsanVitRemoveIscsiInitiatorGroupResponse,omitempty"`
	Fault_ *soap.Fault                                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanVitRemoveIscsiInitiatorGroupBody) Fault() *soap.Fault { return b.Fault_ }

func VsanVitRemoveIscsiInitiatorGroup(ctx context.Context, r soap.RoundTripper, req *types.VsanVitRemoveIscsiInitiatorGroup) (*types.VsanVitRemoveIscsiInitiatorGroupResponse, error) {
	var reqBody, resBody VsanVitRemoveIscsiInitiatorGroupBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanVitCreateHomeObjectBody struct {
	Req    *types.VsanVitCreateHomeObject         `xml:"urn:vsan VsanVitCreateHomeObject,omitempty"`
	Res    *types.VsanVitCreateHomeObjectResponse `xml:"urn:vsan VsanVitCreateHomeObjectResponse,omitempty"`
	Fault_ *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanVitCreateHomeObjectBody) Fault() *soap.Fault { return b.Fault_ }

func VsanVitCreateHomeObject(ctx context.Context, r soap.RoundTripper, req *types.VsanVitCreateHomeObject) (*types.VsanVitCreateHomeObjectResponse, error) {
	var reqBody, resBody VsanVitCreateHomeObjectBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanVitGetHomeObjectBody struct {
	Req    *types.VsanVitGetHomeObject         `xml:"urn:vsan VsanVitGetHomeObject,omitempty"`
	Res    *types.VsanVitGetHomeObjectResponse `xml:"urn:vsan VsanVitGetHomeObjectResponse,omitempty"`
	Fault_ *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanVitGetHomeObjectBody) Fault() *soap.Fault { return b.Fault_ }

func VsanVitGetHomeObject(ctx context.Context, r soap.RoundTripper, req *types.VsanVitGetHomeObject) (*types.VsanVitGetHomeObjectResponse, error) {
	var reqBody, resBody VsanVitGetHomeObjectBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanVitGetIscsiTargetBody struct {
	Req    *types.VsanVitGetIscsiTarget         `xml:"urn:vsan VsanVitGetIscsiTarget,omitempty"`
	Res    *types.VsanVitGetIscsiTargetResponse `xml:"urn:vsan VsanVitGetIscsiTargetResponse,omitempty"`
	Fault_ *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanVitGetIscsiTargetBody) Fault() *soap.Fault { return b.Fault_ }

func VsanVitGetIscsiTarget(ctx context.Context, r soap.RoundTripper, req *types.VsanVitGetIscsiTarget) (*types.VsanVitGetIscsiTargetResponse, error) {
	var reqBody, resBody VsanVitGetIscsiTargetBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanVitRemoveIscsiTargetBody struct {
	Req    *types.VsanVitRemoveIscsiTarget         `xml:"urn:vsan VsanVitRemoveIscsiTarget,omitempty"`
	Res    *types.VsanVitRemoveIscsiTargetResponse `xml:"urn:vsan VsanVitRemoveIscsiTargetResponse,omitempty"`
	Fault_ *soap.Fault                             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanVitRemoveIscsiTargetBody) Fault() *soap.Fault { return b.Fault_ }

func VsanVitRemoveIscsiTarget(ctx context.Context, r soap.RoundTripper, req *types.VsanVitRemoveIscsiTarget) (*types.VsanVitRemoveIscsiTargetResponse, error) {
	var reqBody, resBody VsanVitRemoveIscsiTargetBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanVitAddIscsiLUNBody struct {
	Req    *types.VsanVitAddIscsiLUN         `xml:"urn:vsan VsanVitAddIscsiLUN,omitempty"`
	Res    *types.VsanVitAddIscsiLUNResponse `xml:"urn:vsan VsanVitAddIscsiLUNResponse,omitempty"`
	Fault_ *soap.Fault                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanVitAddIscsiLUNBody) Fault() *soap.Fault { return b.Fault_ }

func VsanVitAddIscsiLUN(ctx context.Context, r soap.RoundTripper, req *types.VsanVitAddIscsiLUN) (*types.VsanVitAddIscsiLUNResponse, error) {
	var reqBody, resBody VsanVitAddIscsiLUNBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanVitGetIscsiInitiatorGroupsBody struct {
	Req    *types.VsanVitGetIscsiInitiatorGroups         `xml:"urn:vsan VsanVitGetIscsiInitiatorGroups,omitempty"`
	Res    *types.VsanVitGetIscsiInitiatorGroupsResponse `xml:"urn:vsan VsanVitGetIscsiInitiatorGroupsResponse,omitempty"`
	Fault_ *soap.Fault                                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanVitGetIscsiInitiatorGroupsBody) Fault() *soap.Fault { return b.Fault_ }

func VsanVitGetIscsiInitiatorGroups(ctx context.Context, r soap.RoundTripper, req *types.VsanVitGetIscsiInitiatorGroups) (*types.VsanVitGetIscsiInitiatorGroupsResponse, error) {
	var reqBody, resBody VsanVitGetIscsiInitiatorGroupsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanVitEnableIscsiTargetServiceBody struct {
	Req    *types.VsanVitEnableIscsiTargetService         `xml:"urn:vsan VsanVitEnableIscsiTargetService,omitempty"`
	Res    *types.VsanVitEnableIscsiTargetServiceResponse `xml:"urn:vsan VsanVitEnableIscsiTargetServiceResponse,omitempty"`
	Fault_ *soap.Fault                                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanVitEnableIscsiTargetServiceBody) Fault() *soap.Fault { return b.Fault_ }

func VsanVitEnableIscsiTargetService(ctx context.Context, r soap.RoundTripper, req *types.VsanVitEnableIscsiTargetService) (*types.VsanVitEnableIscsiTargetServiceResponse, error) {
	var reqBody, resBody VsanVitEnableIscsiTargetServiceBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanVitGetIscsiTargetsBody struct {
	Req    *types.VsanVitGetIscsiTargets         `xml:"urn:vsan VsanVitGetIscsiTargets,omitempty"`
	Res    *types.VsanVitGetIscsiTargetsResponse `xml:"urn:vsan VsanVitGetIscsiTargetsResponse,omitempty"`
	Fault_ *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanVitGetIscsiTargetsBody) Fault() *soap.Fault { return b.Fault_ }

func VsanVitGetIscsiTargets(ctx context.Context, r soap.RoundTripper, req *types.VsanVitGetIscsiTargets) (*types.VsanVitGetIscsiTargetsResponse, error) {
	var reqBody, resBody VsanVitGetIscsiTargetsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanVitAddIscsiTargetBody struct {
	Req    *types.VsanVitAddIscsiTarget         `xml:"urn:vsan VsanVitAddIscsiTarget,omitempty"`
	Res    *types.VsanVitAddIscsiTargetResponse `xml:"urn:vsan VsanVitAddIscsiTargetResponse,omitempty"`
	Fault_ *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanVitAddIscsiTargetBody) Fault() *soap.Fault { return b.Fault_ }

func VsanVitAddIscsiTarget(ctx context.Context, r soap.RoundTripper, req *types.VsanVitAddIscsiTarget) (*types.VsanVitAddIscsiTargetResponse, error) {
	var reqBody, resBody VsanVitAddIscsiTargetBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanVitAddIscsiInitiatorGroupBody struct {
	Req    *types.VsanVitAddIscsiInitiatorGroup         `xml:"urn:vsan VsanVitAddIscsiInitiatorGroup,omitempty"`
	Res    *types.VsanVitAddIscsiInitiatorGroupResponse `xml:"urn:vsan VsanVitAddIscsiInitiatorGroupResponse,omitempty"`
	Fault_ *soap.Fault                                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanVitAddIscsiInitiatorGroupBody) Fault() *soap.Fault { return b.Fault_ }

func VsanVitAddIscsiInitiatorGroup(ctx context.Context, r soap.RoundTripper, req *types.VsanVitAddIscsiInitiatorGroup) (*types.VsanVitAddIscsiInitiatorGroupResponse, error) {
	var reqBody, resBody VsanVitAddIscsiInitiatorGroupBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanVitGetIscsiTargetServiceConfigBody struct {
	Req    *types.VsanVitGetIscsiTargetServiceConfig         `xml:"urn:vsan VsanVitGetIscsiTargetServiceConfig,omitempty"`
	Res    *types.VsanVitGetIscsiTargetServiceConfigResponse `xml:"urn:vsan VsanVitGetIscsiTargetServiceConfigResponse,omitempty"`
	Fault_ *soap.Fault                                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanVitGetIscsiTargetServiceConfigBody) Fault() *soap.Fault { return b.Fault_ }

func VsanVitGetIscsiTargetServiceConfig(ctx context.Context, r soap.RoundTripper, req *types.VsanVitGetIscsiTargetServiceConfig) (*types.VsanVitGetIscsiTargetServiceConfigResponse, error) {
	var reqBody, resBody VsanVitGetIscsiTargetServiceConfigBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type PerformVsanUpgradeExBody struct {
	Req    *types.PerformVsanUpgradeEx         `xml:"urn:vsan PerformVsanUpgradeEx,omitempty"`
	Res    *types.PerformVsanUpgradeExResponse `xml:"urn:vsan PerformVsanUpgradeExResponse,omitempty"`
	Fault_ *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *PerformVsanUpgradeExBody) Fault() *soap.Fault { return b.Fault_ }

func PerformVsanUpgradeEx(ctx context.Context, r soap.RoundTripper, req *types.PerformVsanUpgradeEx) (*types.PerformVsanUpgradeExResponse, error) {
	var reqBody, resBody PerformVsanUpgradeExBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanQueryUpgradeStatusExBody struct {
	Req    *types.VsanQueryUpgradeStatusEx         `xml:"urn:vsan VsanQueryUpgradeStatusEx,omitempty"`
	Res    *types.VsanQueryUpgradeStatusExResponse `xml:"urn:vsan VsanQueryUpgradeStatusExResponse,omitempty"`
	Fault_ *soap.Fault                             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanQueryUpgradeStatusExBody) Fault() *soap.Fault { return b.Fault_ }

func VsanQueryUpgradeStatusEx(ctx context.Context, r soap.RoundTripper, req *types.VsanQueryUpgradeStatusEx) (*types.VsanQueryUpgradeStatusExResponse, error) {
	var reqBody, resBody VsanQueryUpgradeStatusExBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RetrieveSupportedVsanFormatVersionBody struct {
	Req    *types.RetrieveSupportedVsanFormatVersion         `xml:"urn:vsan RetrieveSupportedVsanFormatVersion,omitempty"`
	Res    *types.RetrieveSupportedVsanFormatVersionResponse `xml:"urn:vsan RetrieveSupportedVsanFormatVersionResponse,omitempty"`
	Fault_ *soap.Fault                                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RetrieveSupportedVsanFormatVersionBody) Fault() *soap.Fault { return b.Fault_ }

func RetrieveSupportedVsanFormatVersion(ctx context.Context, r soap.RoundTripper, req *types.RetrieveSupportedVsanFormatVersion) (*types.RetrieveSupportedVsanFormatVersionResponse, error) {
	var reqBody, resBody RetrieveSupportedVsanFormatVersionBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type PerformVsanUpgradePreflightCheckExBody struct {
	Req    *types.PerformVsanUpgradePreflightCheckEx         `xml:"urn:vsan PerformVsanUpgradePreflightCheckEx,omitempty"`
	Res    *types.PerformVsanUpgradePreflightCheckExResponse `xml:"urn:vsan PerformVsanUpgradePreflightCheckExResponse,omitempty"`
	Fault_ *soap.Fault                                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *PerformVsanUpgradePreflightCheckExBody) Fault() *soap.Fault { return b.Fault_ }

func PerformVsanUpgradePreflightCheckEx(ctx context.Context, r soap.RoundTripper, req *types.PerformVsanUpgradePreflightCheckEx) (*types.PerformVsanUpgradePreflightCheckExResponse, error) {
	var reqBody, resBody PerformVsanUpgradePreflightCheckExBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type PerformVsanUpgradePreflightAsyncCheck_TaskBody struct {
	Req    *types.PerformVsanUpgradePreflightAsyncCheck_Task         `xml:"urn:vsan PerformVsanUpgradePreflightAsyncCheck_Task,omitempty"`
	Res    *types.PerformVsanUpgradePreflightAsyncCheck_TaskResponse `xml:"urn:vsan PerformVsanUpgradePreflightAsyncCheck_TaskResponse,omitempty"`
	Fault_ *soap.Fault                                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *PerformVsanUpgradePreflightAsyncCheck_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func PerformVsanUpgradePreflightAsyncCheck_Task(ctx context.Context, r soap.RoundTripper, req *types.PerformVsanUpgradePreflightAsyncCheck_Task) (*types.PerformVsanUpgradePreflightAsyncCheck_TaskResponse, error) {
	var reqBody, resBody PerformVsanUpgradePreflightAsyncCheck_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanQuerySpaceUsageBody struct {
	Req    *types.VsanQuerySpaceUsage         `xml:"urn:vsan VsanQuerySpaceUsage,omitempty"`
	Res    *types.VsanQuerySpaceUsageResponse `xml:"urn:vsan VsanQuerySpaceUsageResponse,omitempty"`
	Fault_ *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanQuerySpaceUsageBody) Fault() *soap.Fault { return b.Fault_ }

func VsanQuerySpaceUsage(ctx context.Context, r soap.RoundTripper, req *types.VsanQuerySpaceUsage) (*types.VsanQuerySpaceUsageResponse, error) {
	var reqBody, resBody VsanQuerySpaceUsageBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanQueryObjectContainerSpaceUsageBody struct {
	Req    *types.VsanQueryObjectContainerSpaceUsage         `xml:"urn:vsan VsanQueryObjectContainerSpaceUsage,omitempty"`
	Res    *types.VsanQueryObjectContainerSpaceUsageResponse `xml:"urn:vsan VsanQueryObjectContainerSpaceUsageResponse,omitempty"`
	Fault_ *soap.Fault                                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanQueryObjectContainerSpaceUsageBody) Fault() *soap.Fault { return b.Fault_ }

func VsanQueryObjectContainerSpaceUsage(ctx context.Context, r soap.RoundTripper, req *types.VsanQueryObjectContainerSpaceUsage) (*types.VsanQueryObjectContainerSpaceUsageResponse, error) {
	var reqBody, resBody VsanQueryObjectContainerSpaceUsageBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanQueryRemoteDataProtectionSpaceUsageBody struct {
	Req    *types.VsanQueryRemoteDataProtectionSpaceUsage         `xml:"urn:vsan VsanQueryRemoteDataProtectionSpaceUsage,omitempty"`
	Res    *types.VsanQueryRemoteDataProtectionSpaceUsageResponse `xml:"urn:vsan VsanQueryRemoteDataProtectionSpaceUsageResponse,omitempty"`
	Fault_ *soap.Fault                                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanQueryRemoteDataProtectionSpaceUsageBody) Fault() *soap.Fault { return b.Fault_ }

func VsanQueryRemoteDataProtectionSpaceUsage(ctx context.Context, r soap.RoundTripper, req *types.VsanQueryRemoteDataProtectionSpaceUsage) (*types.VsanQueryRemoteDataProtectionSpaceUsageResponse, error) {
	var reqBody, resBody VsanQueryRemoteDataProtectionSpaceUsageBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanRetrieveProfilingTraceBody struct {
	Req    *types.VsanRetrieveProfilingTrace         `xml:"urn:vsan VsanRetrieveProfilingTrace,omitempty"`
	Res    *types.VsanRetrieveProfilingTraceResponse `xml:"urn:vsan VsanRetrieveProfilingTraceResponse,omitempty"`
	Fault_ *soap.Fault                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanRetrieveProfilingTraceBody) Fault() *soap.Fault { return b.Fault_ }

func VsanRetrieveProfilingTrace(ctx context.Context, r soap.RoundTripper, req *types.VsanRetrieveProfilingTrace) (*types.VsanRetrieveProfilingTraceResponse, error) {
	var reqBody, resBody VsanRetrieveProfilingTraceBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanRemediateVsanClusterBody struct {
	Req    *types.VsanRemediateVsanCluster         `xml:"urn:vsan VsanRemediateVsanCluster,omitempty"`
	Res    *types.VsanRemediateVsanClusterResponse `xml:"urn:vsan VsanRemediateVsanClusterResponse,omitempty"`
	Fault_ *soap.Fault                             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanRemediateVsanClusterBody) Fault() *soap.Fault { return b.Fault_ }

func VsanRemediateVsanCluster(ctx context.Context, r soap.RoundTripper, req *types.VsanRemediateVsanCluster) (*types.VsanRemediateVsanClusterResponse, error) {
	var reqBody, resBody VsanRemediateVsanClusterBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanRemediateDataProtectionConfigBody struct {
	Req    *types.VsanRemediateDataProtectionConfig         `xml:"urn:vsan VsanRemediateDataProtectionConfig,omitempty"`
	Res    *types.VsanRemediateDataProtectionConfigResponse `xml:"urn:vsan VsanRemediateDataProtectionConfigResponse,omitempty"`
	Fault_ *soap.Fault                                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanRemediateDataProtectionConfigBody) Fault() *soap.Fault { return b.Fault_ }

func VsanRemediateDataProtectionConfig(ctx context.Context, r soap.RoundTripper, req *types.VsanRemediateDataProtectionConfig) (*types.VsanRemediateDataProtectionConfigResponse, error) {
	var reqBody, resBody VsanRemediateDataProtectionConfigBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanReconfigProfilingBody struct {
	Req    *types.VsanReconfigProfiling         `xml:"urn:vsan VsanReconfigProfiling,omitempty"`
	Res    *types.VsanReconfigProfilingResponse `xml:"urn:vsan VsanReconfigProfilingResponse,omitempty"`
	Fault_ *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanReconfigProfilingBody) Fault() *soap.Fault { return b.Fault_ }

func VsanReconfigProfiling(ctx context.Context, r soap.RoundTripper, req *types.VsanReconfigProfiling) (*types.VsanReconfigProfilingResponse, error) {
	var reqBody, resBody VsanReconfigProfilingBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanRemediateVcBody struct {
	Req    *types.VsanRemediateVc         `xml:"urn:vsan VsanRemediateVc,omitempty"`
	Res    *types.VsanRemediateVcResponse `xml:"urn:vsan VsanRemediateVcResponse,omitempty"`
	Fault_ *soap.Fault                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanRemediateVcBody) Fault() *soap.Fault { return b.Fault_ }

func VsanRemediateVc(ctx context.Context, r soap.RoundTripper, req *types.VsanRemediateVc) (*types.VsanRemediateVcResponse, error) {
	var reqBody, resBody VsanRemediateVcBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanRemediateVsanHostBody struct {
	Req    *types.VsanRemediateVsanHost         `xml:"urn:vsan VsanRemediateVsanHost,omitempty"`
	Res    *types.VsanRemediateVsanHostResponse `xml:"urn:vsan VsanRemediateVsanHostResponse,omitempty"`
	Fault_ *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanRemediateVsanHostBody) Fault() *soap.Fault { return b.Fault_ }

func VsanRemediateVsanHost(ctx context.Context, r soap.RoundTripper, req *types.VsanRemediateVsanHost) (*types.VsanRemediateVsanHostResponse, error) {
	var reqBody, resBody VsanRemediateVsanHostBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanGetCapabilitiesBody struct {
	Req    *types.VsanGetCapabilities         `xml:"urn:vsan VsanGetCapabilities,omitempty"`
	Res    *types.VsanGetCapabilitiesResponse `xml:"urn:vsan VsanGetCapabilitiesResponse,omitempty"`
	Fault_ *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanGetCapabilitiesBody) Fault() *soap.Fault { return b.Fault_ }

func VsanGetCapabilities(ctx context.Context, r soap.RoundTripper, req *types.VsanGetCapabilities) (*types.VsanGetCapabilitiesResponse, error) {
	var reqBody, resBody VsanGetCapabilitiesBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type SetVsanVumConfigBody struct {
	Req    *types.SetVsanVumConfig         `xml:"urn:vsan SetVsanVumConfig,omitempty"`
	Res    *types.SetVsanVumConfigResponse `xml:"urn:vsan SetVsanVumConfigResponse,omitempty"`
	Fault_ *soap.Fault                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *SetVsanVumConfigBody) Fault() *soap.Fault { return b.Fault_ }

func SetVsanVumConfig(ctx context.Context, r soap.RoundTripper, req *types.SetVsanVumConfig) (*types.SetVsanVumConfigResponse, error) {
	var reqBody, resBody SetVsanVumConfigBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanHostUpdateFirmwareBody struct {
	Req    *types.VsanHostUpdateFirmware         `xml:"urn:vsan VsanHostUpdateFirmware,omitempty"`
	Res    *types.VsanHostUpdateFirmwareResponse `xml:"urn:vsan VsanHostUpdateFirmwareResponse,omitempty"`
	Fault_ *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanHostUpdateFirmwareBody) Fault() *soap.Fault { return b.Fault_ }

func VsanHostUpdateFirmware(ctx context.Context, r soap.RoundTripper, req *types.VsanHostUpdateFirmware) (*types.VsanHostUpdateFirmwareResponse, error) {
	var reqBody, resBody VsanHostUpdateFirmwareBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type FetchIsoDepotCookieBody struct {
	Req    *types.FetchIsoDepotCookie         `xml:"urn:vsan FetchIsoDepotCookie,omitempty"`
	Res    *types.FetchIsoDepotCookieResponse `xml:"urn:vsan FetchIsoDepotCookieResponse,omitempty"`
	Fault_ *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *FetchIsoDepotCookieBody) Fault() *soap.Fault { return b.Fault_ }

func FetchIsoDepotCookie(ctx context.Context, r soap.RoundTripper, req *types.FetchIsoDepotCookie) (*types.FetchIsoDepotCookieResponse, error) {
	var reqBody, resBody FetchIsoDepotCookieBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type GetVsanVumConfigBody struct {
	Req    *types.GetVsanVumConfig         `xml:"urn:vsan GetVsanVumConfig,omitempty"`
	Res    *types.GetVsanVumConfigResponse `xml:"urn:vsan GetVsanVumConfigResponse,omitempty"`
	Fault_ *soap.Fault                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *GetVsanVumConfigBody) Fault() *soap.Fault { return b.Fault_ }

func GetVsanVumConfig(ctx context.Context, r soap.RoundTripper, req *types.GetVsanVumConfig) (*types.GetVsanVumConfigResponse, error) {
	var reqBody, resBody GetVsanVumConfigBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanVcUploadReleaseDbBody struct {
	Req    *types.VsanVcUploadReleaseDb         `xml:"urn:vsan VsanVcUploadReleaseDb,omitempty"`
	Res    *types.VsanVcUploadReleaseDbResponse `xml:"urn:vsan VsanVcUploadReleaseDbResponse,omitempty"`
	Fault_ *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanVcUploadReleaseDbBody) Fault() *soap.Fault { return b.Fault_ }

func VsanVcUploadReleaseDb(ctx context.Context, r soap.RoundTripper, req *types.VsanVcUploadReleaseDb) (*types.VsanVcUploadReleaseDbResponse, error) {
	var reqBody, resBody VsanVcUploadReleaseDbBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryIoInsightBody struct {
	Req    *types.QueryIoInsight         `xml:"urn:vsan QueryIoInsight,omitempty"`
	Res    *types.QueryIoInsightResponse `xml:"urn:vsan QueryIoInsightResponse,omitempty"`
	Fault_ *soap.Fault                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryIoInsightBody) Fault() *soap.Fault { return b.Fault_ }

func QueryIoInsight(ctx context.Context, r soap.RoundTripper, req *types.QueryIoInsight) (*types.QueryIoInsightResponse, error) {
	var reqBody, resBody QueryIoInsightBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type StartIoInsightBody struct {
	Req    *types.StartIoInsight         `xml:"urn:vsan StartIoInsight,omitempty"`
	Res    *types.StartIoInsightResponse `xml:"urn:vsan StartIoInsightResponse,omitempty"`
	Fault_ *soap.Fault                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *StartIoInsightBody) Fault() *soap.Fault { return b.Fault_ }

func StartIoInsight(ctx context.Context, r soap.RoundTripper, req *types.StartIoInsight) (*types.StartIoInsightResponse, error) {
	var reqBody, resBody StartIoInsightBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type StopIoInsightBody struct {
	Req    *types.StopIoInsight         `xml:"urn:vsan StopIoInsight,omitempty"`
	Res    *types.StopIoInsightResponse `xml:"urn:vsan StopIoInsightResponse,omitempty"`
	Fault_ *soap.Fault                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *StopIoInsightBody) Fault() *soap.Fault { return b.Fault_ }

func StopIoInsight(ctx context.Context, r soap.RoundTripper, req *types.StopIoInsight) (*types.StopIoInsightResponse, error) {
	var reqBody, resBody StopIoInsightBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanGetResourceCheckStatusBody struct {
	Req    *types.VsanGetResourceCheckStatus         `xml:"urn:vsan VsanGetResourceCheckStatus,omitempty"`
	Res    *types.VsanGetResourceCheckStatusResponse `xml:"urn:vsan VsanGetResourceCheckStatusResponse,omitempty"`
	Fault_ *soap.Fault                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanGetResourceCheckStatusBody) Fault() *soap.Fault { return b.Fault_ }

func VsanGetResourceCheckStatus(ctx context.Context, r soap.RoundTripper, req *types.VsanGetResourceCheckStatus) (*types.VsanGetResourceCheckStatusResponse, error) {
	var reqBody, resBody VsanGetResourceCheckStatusBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanPerformResourceCheckBody struct {
	Req    *types.VsanPerformResourceCheck         `xml:"urn:vsan VsanPerformResourceCheck,omitempty"`
	Res    *types.VsanPerformResourceCheckResponse `xml:"urn:vsan VsanPerformResourceCheckResponse,omitempty"`
	Fault_ *soap.Fault                             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanPerformResourceCheckBody) Fault() *soap.Fault { return b.Fault_ }

func VsanPerformResourceCheck(ctx context.Context, r soap.RoundTripper, req *types.VsanPerformResourceCheck) (*types.VsanPerformResourceCheckResponse, error) {
	var reqBody, resBody VsanPerformResourceCheckBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanHostCancelResourceCheckBody struct {
	Req    *types.VsanHostCancelResourceCheck         `xml:"urn:vsan VsanHostCancelResourceCheck,omitempty"`
	Res    *types.VsanHostCancelResourceCheckResponse `xml:"urn:vsan VsanHostCancelResourceCheckResponse,omitempty"`
	Fault_ *soap.Fault                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanHostCancelResourceCheckBody) Fault() *soap.Fault { return b.Fault_ }

func VsanHostCancelResourceCheck(ctx context.Context, r soap.RoundTripper, req *types.VsanHostCancelResourceCheck) (*types.VsanHostCancelResourceCheckResponse, error) {
	var reqBody, resBody VsanHostCancelResourceCheckBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanHostPerformResourceCheckBody struct {
	Req    *types.VsanHostPerformResourceCheck         `xml:"urn:vsan VsanHostPerformResourceCheck,omitempty"`
	Res    *types.VsanHostPerformResourceCheckResponse `xml:"urn:vsan VsanHostPerformResourceCheckResponse,omitempty"`
	Fault_ *soap.Fault                                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanHostPerformResourceCheckBody) Fault() *soap.Fault { return b.Fault_ }

func VsanHostPerformResourceCheck(ctx context.Context, r soap.RoundTripper, req *types.VsanHostPerformResourceCheck) (*types.VsanHostPerformResourceCheckResponse, error) {
	var reqBody, resBody VsanHostPerformResourceCheckBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UpgradeFormatVersionForDiskMappingBody struct {
	Req    *types.UpgradeFormatVersionForDiskMapping         `xml:"urn:vsan UpgradeFormatVersionForDiskMapping,omitempty"`
	Res    *types.UpgradeFormatVersionForDiskMappingResponse `xml:"urn:vsan UpgradeFormatVersionForDiskMappingResponse,omitempty"`
	Fault_ *soap.Fault                                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpgradeFormatVersionForDiskMappingBody) Fault() *soap.Fault { return b.Fault_ }

func UpgradeFormatVersionForDiskMapping(ctx context.Context, r soap.RoundTripper, req *types.UpgradeFormatVersionForDiskMapping) (*types.UpgradeFormatVersionForDiskMappingResponse, error) {
	var reqBody, resBody UpgradeFormatVersionForDiskMappingBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanRealignVsanSparseOfflineBody struct {
	Req    *types.VsanRealignVsanSparseOffline         `xml:"urn:vsan VsanRealignVsanSparseOffline,omitempty"`
	Res    *types.VsanRealignVsanSparseOfflineResponse `xml:"urn:vsan VsanRealignVsanSparseOfflineResponse,omitempty"`
	Fault_ *soap.Fault                                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanRealignVsanSparseOfflineBody) Fault() *soap.Fault { return b.Fault_ }

func VsanRealignVsanSparseOffline(ctx context.Context, r soap.RoundTripper, req *types.VsanRealignVsanSparseOffline) (*types.VsanRealignVsanSparseOfflineResponse, error) {
	var reqBody, resBody VsanRealignVsanSparseOfflineBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RetrieveVsanDiskManagementSystemCapabilityBody struct {
	Req    *types.RetrieveVsanDiskManagementSystemCapability         `xml:"urn:vsan RetrieveVsanDiskManagementSystemCapability,omitempty"`
	Res    *types.RetrieveVsanDiskManagementSystemCapabilityResponse `xml:"urn:vsan RetrieveVsanDiskManagementSystemCapabilityResponse,omitempty"`
	Fault_ *soap.Fault                                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RetrieveVsanDiskManagementSystemCapabilityBody) Fault() *soap.Fault { return b.Fault_ }

func RetrieveVsanDiskManagementSystemCapability(ctx context.Context, r soap.RoundTripper, req *types.RetrieveVsanDiskManagementSystemCapability) (*types.RetrieveVsanDiskManagementSystemCapabilityResponse, error) {
	var reqBody, resBody RetrieveVsanDiskManagementSystemCapabilityBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanRemoveConcatBody struct {
	Req    *types.VsanRemoveConcat         `xml:"urn:vsan VsanRemoveConcat,omitempty"`
	Res    *types.VsanRemoveConcatResponse `xml:"urn:vsan VsanRemoveConcatResponse,omitempty"`
	Fault_ *soap.Fault                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanRemoveConcatBody) Fault() *soap.Fault { return b.Fault_ }

func VsanRemoveConcat(ctx context.Context, r soap.RoundTripper, req *types.VsanRemoveConcat) (*types.VsanRemoveConcatResponse, error) {
	var reqBody, resBody VsanRemoveConcatBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryHostDiskMappingsBody struct {
	Req    *types.QueryHostDiskMappings         `xml:"urn:vsan QueryHostDiskMappings,omitempty"`
	Res    *types.QueryHostDiskMappingsResponse `xml:"urn:vsan QueryHostDiskMappingsResponse,omitempty"`
	Fault_ *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryHostDiskMappingsBody) Fault() *soap.Fault { return b.Fault_ }

func QueryHostDiskMappings(ctx context.Context, r soap.RoundTripper, req *types.QueryHostDiskMappings) (*types.QueryHostDiskMappingsResponse, error) {
	var reqBody, resBody QueryHostDiskMappingsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanQueryUnalignedStatusBody struct {
	Req    *types.VsanQueryUnalignedStatus         `xml:"urn:vsan VsanQueryUnalignedStatus,omitempty"`
	Res    *types.VsanQueryUnalignedStatusResponse `xml:"urn:vsan VsanQueryUnalignedStatusResponse,omitempty"`
	Fault_ *soap.Fault                             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanQueryUnalignedStatusBody) Fault() *soap.Fault { return b.Fault_ }

func VsanQueryUnalignedStatus(ctx context.Context, r soap.RoundTripper, req *types.VsanQueryUnalignedStatus) (*types.VsanQueryUnalignedStatusResponse, error) {
	var reqBody, resBody VsanQueryUnalignedStatusBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryDataEfficiencyCapacityStateBody struct {
	Req    *types.QueryDataEfficiencyCapacityState         `xml:"urn:vsan QueryDataEfficiencyCapacityState,omitempty"`
	Res    *types.QueryDataEfficiencyCapacityStateResponse `xml:"urn:vsan QueryDataEfficiencyCapacityStateResponse,omitempty"`
	Fault_ *soap.Fault                                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryDataEfficiencyCapacityStateBody) Fault() *soap.Fault { return b.Fault_ }

func QueryDataEfficiencyCapacityState(ctx context.Context, r soap.RoundTripper, req *types.QueryDataEfficiencyCapacityState) (*types.QueryDataEfficiencyCapacityStateResponse, error) {
	var reqBody, resBody QueryDataEfficiencyCapacityStateBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type FixVsanObjectsBody struct {
	Req    *types.FixVsanObjects         `xml:"urn:vsan FixVsanObjects,omitempty"`
	Res    *types.FixVsanObjectsResponse `xml:"urn:vsan FixVsanObjectsResponse,omitempty"`
	Fault_ *soap.Fault                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *FixVsanObjectsBody) Fault() *soap.Fault { return b.Fault_ }

func FixVsanObjects(ctx context.Context, r soap.RoundTripper, req *types.FixVsanObjects) (*types.FixVsanObjectsResponse, error) {
	var reqBody, resBody FixVsanObjectsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanRealignVsanSparseOnlineBody struct {
	Req    *types.VsanRealignVsanSparseOnline         `xml:"urn:vsan VsanRealignVsanSparseOnline,omitempty"`
	Res    *types.VsanRealignVsanSparseOnlineResponse `xml:"urn:vsan VsanRealignVsanSparseOnlineResponse,omitempty"`
	Fault_ *soap.Fault                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanRealignVsanSparseOnlineBody) Fault() *soap.Fault { return b.Fault_ }

func VsanRealignVsanSparseOnline(ctx context.Context, r soap.RoundTripper, req *types.VsanRealignVsanSparseOnline) (*types.VsanRealignVsanSparseOnlineResponse, error) {
	var reqBody, resBody VsanRealignVsanSparseOnlineBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UpdateCapacityFlashStatusForDisksBody struct {
	Req    *types.UpdateCapacityFlashStatusForDisks         `xml:"urn:vsan UpdateCapacityFlashStatusForDisks,omitempty"`
	Res    *types.UpdateCapacityFlashStatusForDisksResponse `xml:"urn:vsan UpdateCapacityFlashStatusForDisksResponse,omitempty"`
	Fault_ *soap.Fault                                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateCapacityFlashStatusForDisksBody) Fault() *soap.Fault { return b.Fault_ }

func UpdateCapacityFlashStatusForDisks(ctx context.Context, r soap.RoundTripper, req *types.UpdateCapacityFlashStatusForDisks) (*types.UpdateCapacityFlashStatusForDisksResponse, error) {
	var reqBody, resBody UpdateCapacityFlashStatusForDisksBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanAlignObjectSizeBody struct {
	Req    *types.VsanAlignObjectSize         `xml:"urn:vsan VsanAlignObjectSize,omitempty"`
	Res    *types.VsanAlignObjectSizeResponse `xml:"urn:vsan VsanAlignObjectSizeResponse,omitempty"`
	Fault_ *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanAlignObjectSizeBody) Fault() *soap.Fault { return b.Fault_ }

func VsanAlignObjectSize(ctx context.Context, r soap.RoundTripper, req *types.VsanAlignObjectSize) (*types.VsanAlignObjectSizeResponse, error) {
	var reqBody, resBody VsanAlignObjectSizeBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanQueryNamespaceUuidsBody struct {
	Req    *types.VsanQueryNamespaceUuids         `xml:"urn:vsan VsanQueryNamespaceUuids,omitempty"`
	Res    *types.VsanQueryNamespaceUuidsResponse `xml:"urn:vsan VsanQueryNamespaceUuidsResponse,omitempty"`
	Fault_ *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanQueryNamespaceUuidsBody) Fault() *soap.Fault { return b.Fault_ }

func VsanQueryNamespaceUuids(ctx context.Context, r soap.RoundTripper, req *types.VsanQueryNamespaceUuids) (*types.VsanQueryNamespaceUuidsResponse, error) {
	var reqBody, resBody VsanQueryNamespaceUuidsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanScanObjects_TaskBody struct {
	Req    *types.VsanScanObjects_Task         `xml:"urn:vsan VsanScanObjects_Task,omitempty"`
	Res    *types.VsanScanObjects_TaskResponse `xml:"urn:vsan VsanScanObjects_TaskResponse,omitempty"`
	Fault_ *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanScanObjects_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func VsanScanObjects_Task(ctx context.Context, r soap.RoundTripper, req *types.VsanScanObjects_Task) (*types.VsanScanObjects_TaskResponse, error) {
	var reqBody, resBody VsanScanObjects_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanGetOpenV2VsanSparseChainBody struct {
	Req    *types.VsanGetOpenV2VsanSparseChain         `xml:"urn:vsan VsanGetOpenV2VsanSparseChain,omitempty"`
	Res    *types.VsanGetOpenV2VsanSparseChainResponse `xml:"urn:vsan VsanGetOpenV2VsanSparseChainResponse,omitempty"`
	Fault_ *soap.Fault                                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanGetOpenV2VsanSparseChainBody) Fault() *soap.Fault { return b.Fault_ }

func VsanGetOpenV2VsanSparseChain(ctx context.Context, r soap.RoundTripper, req *types.VsanGetOpenV2VsanSparseChain) (*types.VsanGetOpenV2VsanSparseChainResponse, error) {
	var reqBody, resBody VsanGetOpenV2VsanSparseChainBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanScanNamespaceBody struct {
	Req    *types.VsanScanNamespace         `xml:"urn:vsan VsanScanNamespace,omitempty"`
	Res    *types.VsanScanNamespaceResponse `xml:"urn:vsan VsanScanNamespaceResponse,omitempty"`
	Fault_ *soap.Fault                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanScanNamespaceBody) Fault() *soap.Fault { return b.Fault_ }

func VsanScanNamespace(ctx context.Context, r soap.RoundTripper, req *types.VsanScanNamespace) (*types.VsanScanNamespaceResponse, error) {
	var reqBody, resBody VsanScanNamespaceBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanVibInstall_TaskBody struct {
	Req    *types.VsanVibInstall_Task         `xml:"urn:vsan VsanVibInstall_Task,omitempty"`
	Res    *types.VsanVibInstall_TaskResponse `xml:"urn:vsan VsanVibInstall_TaskResponse,omitempty"`
	Fault_ *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanVibInstall_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func VsanVibInstall_Task(ctx context.Context, r soap.RoundTripper, req *types.VsanVibInstall_Task) (*types.VsanVibInstall_TaskResponse, error) {
	var reqBody, resBody VsanVibInstall_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanVibInstallPreflightCheckBody struct {
	Req    *types.VsanVibInstallPreflightCheck         `xml:"urn:vsan VsanVibInstallPreflightCheck,omitempty"`
	Res    *types.VsanVibInstallPreflightCheckResponse `xml:"urn:vsan VsanVibInstallPreflightCheckResponse,omitempty"`
	Fault_ *soap.Fault                                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanVibInstallPreflightCheckBody) Fault() *soap.Fault { return b.Fault_ }

func VsanVibInstallPreflightCheck(ctx context.Context, r soap.RoundTripper, req *types.VsanVibInstallPreflightCheck) (*types.VsanVibInstallPreflightCheckResponse, error) {
	var reqBody, resBody VsanVibInstallPreflightCheckBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanVibScanBody struct {
	Req    *types.VsanVibScan         `xml:"urn:vsan VsanVibScan,omitempty"`
	Res    *types.VsanVibScanResponse `xml:"urn:vsan VsanVibScanResponse,omitempty"`
	Fault_ *soap.Fault                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanVibScanBody) Fault() *soap.Fault { return b.Fault_ }

func VsanVibScan(ctx context.Context, r soap.RoundTripper, req *types.VsanVibScan) (*types.VsanVibScanResponse, error) {
	var reqBody, resBody VsanVibScanBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanGetTaskStatusBody struct {
	Req    *types.VsanGetTaskStatus         `xml:"urn:vsan VsanGetTaskStatus,omitempty"`
	Res    *types.VsanGetTaskStatusResponse `xml:"urn:vsan VsanGetTaskStatusResponse,omitempty"`
	Fault_ *soap.Fault                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanGetTaskStatusBody) Fault() *soap.Fault { return b.Fault_ }

func VsanGetTaskStatus(ctx context.Context, r soap.RoundTripper, req *types.VsanGetTaskStatus) (*types.VsanGetTaskStatusResponse, error) {
	var reqBody, resBody VsanGetTaskStatusBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RetrieveLookupTableBody struct {
	Req    *types.RetrieveLookupTable         `xml:"urn:vsan RetrieveLookupTable,omitempty"`
	Res    *types.RetrieveLookupTableResponse `xml:"urn:vsan RetrieveLookupTableResponse,omitempty"`
	Fault_ *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RetrieveLookupTableBody) Fault() *soap.Fault { return b.Fault_ }

func RetrieveLookupTable(ctx context.Context, r soap.RoundTripper, req *types.RetrieveLookupTable) (*types.RetrieveLookupTableResponse, error) {
	var reqBody, resBody RetrieveLookupTableBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanRetrievePropertiesBody struct {
	Req    *types.VsanRetrieveProperties         `xml:"urn:vsan VsanRetrieveProperties,omitempty"`
	Res    *types.VsanRetrievePropertiesResponse `xml:"urn:vsan VsanRetrievePropertiesResponse,omitempty"`
	Fault_ *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanRetrievePropertiesBody) Fault() *soap.Fault { return b.Fault_ }

func VsanRetrieveProperties(ctx context.Context, r soap.RoundTripper, req *types.VsanRetrieveProperties) (*types.VsanRetrievePropertiesResponse, error) {
	var reqBody, resBody VsanRetrievePropertiesBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanRetrievePropertiesJsonBody struct {
	Req    *types.VsanRetrievePropertiesJson         `xml:"urn:vsan VsanRetrievePropertiesJson,omitempty"`
	Res    *types.VsanRetrievePropertiesJsonResponse `xml:"urn:vsan VsanRetrievePropertiesJsonResponse,omitempty"`
	Fault_ *soap.Fault                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanRetrievePropertiesJsonBody) Fault() *soap.Fault { return b.Fault_ }

func VsanRetrievePropertiesJson(ctx context.Context, r soap.RoundTripper, req *types.VsanRetrievePropertiesJson) (*types.VsanRetrievePropertiesJsonResponse, error) {
	var reqBody, resBody VsanRetrievePropertiesJsonBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanUnmountDiskMappingExBody struct {
	Req    *types.VsanUnmountDiskMappingEx         `xml:"urn:vsan VsanUnmountDiskMappingEx,omitempty"`
	Res    *types.VsanUnmountDiskMappingExResponse `xml:"urn:vsan VsanUnmountDiskMappingExResponse,omitempty"`
	Fault_ *soap.Fault                             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanUnmountDiskMappingExBody) Fault() *soap.Fault { return b.Fault_ }

func VsanUnmountDiskMappingEx(ctx context.Context, r soap.RoundTripper, req *types.VsanUnmountDiskMappingEx) (*types.VsanUnmountDiskMappingExResponse, error) {
	var reqBody, resBody VsanUnmountDiskMappingExBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryVsanAwsCredentialsBody struct {
	Req    *types.QueryVsanAwsCredentials         `xml:"urn:vsan QueryVsanAwsCredentials,omitempty"`
	Res    *types.QueryVsanAwsCredentialsResponse `xml:"urn:vsan QueryVsanAwsCredentialsResponse,omitempty"`
	Fault_ *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryVsanAwsCredentialsBody) Fault() *soap.Fault { return b.Fault_ }

func QueryVsanAwsCredentials(ctx context.Context, r soap.RoundTripper, req *types.QueryVsanAwsCredentials) (*types.QueryVsanAwsCredentialsResponse, error) {
	var reqBody, resBody QueryVsanAwsCredentialsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanQuerySyncingVsanObjectsBody struct {
	Req    *types.VsanQuerySyncingVsanObjects         `xml:"urn:vsan VsanQuerySyncingVsanObjects,omitempty"`
	Res    *types.VsanQuerySyncingVsanObjectsResponse `xml:"urn:vsan VsanQuerySyncingVsanObjectsResponse,omitempty"`
	Fault_ *soap.Fault                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanQuerySyncingVsanObjectsBody) Fault() *soap.Fault { return b.Fault_ }

func VsanQuerySyncingVsanObjects(ctx context.Context, r soap.RoundTripper, req *types.VsanQuerySyncingVsanObjects) (*types.VsanQuerySyncingVsanObjectsResponse, error) {
	var reqBody, resBody VsanQuerySyncingVsanObjectsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RefreshVsanSystemBody struct {
	Req    *types.RefreshVsanSystem         `xml:"urn:vsan RefreshVsanSystem,omitempty"`
	Res    *types.RefreshVsanSystemResponse `xml:"urn:vsan RefreshVsanSystemResponse,omitempty"`
	Fault_ *soap.Fault                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RefreshVsanSystemBody) Fault() *soap.Fault { return b.Fault_ }

func RefreshVsanSystem(ctx context.Context, r soap.RoundTripper, req *types.RefreshVsanSystem) (*types.RefreshVsanSystemResponse, error) {
	var reqBody, resBody RefreshVsanSystemBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UpdateVmMetadataBody struct {
	Req    *types.UpdateVmMetadata         `xml:"urn:vsan UpdateVmMetadata,omitempty"`
	Res    *types.UpdateVmMetadataResponse `xml:"urn:vsan UpdateVmMetadataResponse,omitempty"`
	Fault_ *soap.Fault                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateVmMetadataBody) Fault() *soap.Fault { return b.Fault_ }

func UpdateVmMetadata(ctx context.Context, r soap.RoundTripper, req *types.UpdateVmMetadata) (*types.UpdateVmMetadataResponse, error) {
	var reqBody, resBody UpdateVmMetadataBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanQueryHostDrsStatsBody struct {
	Req    *types.VsanQueryHostDrsStats         `xml:"urn:vsan VsanQueryHostDrsStats,omitempty"`
	Res    *types.VsanQueryHostDrsStatsResponse `xml:"urn:vsan VsanQueryHostDrsStatsResponse,omitempty"`
	Fault_ *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanQueryHostDrsStatsBody) Fault() *soap.Fault { return b.Fault_ }

func VsanQueryHostDrsStats(ctx context.Context, r soap.RoundTripper, req *types.VsanQueryHostDrsStats) (*types.VsanQueryHostDrsStatsResponse, error) {
	var reqBody, resBody VsanQueryHostDrsStatsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VmMonitorUnwatchObjectsBody struct {
	Req    *types.VmMonitorUnwatchObjects         `xml:"urn:vsan VmMonitorUnwatchObjects,omitempty"`
	Res    *types.VmMonitorUnwatchObjectsResponse `xml:"urn:vsan VmMonitorUnwatchObjectsResponse,omitempty"`
	Fault_ *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VmMonitorUnwatchObjectsBody) Fault() *soap.Fault { return b.Fault_ }

func VmMonitorUnwatchObjects(ctx context.Context, r soap.RoundTripper, req *types.VmMonitorUnwatchObjects) (*types.VmMonitorUnwatchObjectsResponse, error) {
	var reqBody, resBody VmMonitorUnwatchObjectsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanExitMaintenanceMode_TaskBody struct {
	Req    *types.VsanExitMaintenanceMode_Task         `xml:"urn:vsan VsanExitMaintenanceMode_Task,omitempty"`
	Res    *types.VsanExitMaintenanceMode_TaskResponse `xml:"urn:vsan VsanExitMaintenanceMode_TaskResponse,omitempty"`
	Fault_ *soap.Fault                                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanExitMaintenanceMode_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func VsanExitMaintenanceMode_Task(ctx context.Context, r soap.RoundTripper, req *types.VsanExitMaintenanceMode_Task) (*types.VsanExitMaintenanceMode_TaskResponse, error) {
	var reqBody, resBody VsanExitMaintenanceMode_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanGetLocalAffinityVmListBody struct {
	Req    *types.VsanGetLocalAffinityVmList         `xml:"urn:vsan VsanGetLocalAffinityVmList,omitempty"`
	Res    *types.VsanGetLocalAffinityVmListResponse `xml:"urn:vsan VsanGetLocalAffinityVmListResponse,omitempty"`
	Fault_ *soap.Fault                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanGetLocalAffinityVmListBody) Fault() *soap.Fault { return b.Fault_ }

func VsanGetLocalAffinityVmList(ctx context.Context, r soap.RoundTripper, req *types.VsanGetLocalAffinityVmList) (*types.VsanGetLocalAffinityVmListResponse, error) {
	var reqBody, resBody VsanGetLocalAffinityVmListBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanGetAssociatedSpbmProfileBody struct {
	Req    *types.VsanGetAssociatedSpbmProfile         `xml:"urn:vsan VsanGetAssociatedSpbmProfile,omitempty"`
	Res    *types.VsanGetAssociatedSpbmProfileResponse `xml:"urn:vsan VsanGetAssociatedSpbmProfileResponse,omitempty"`
	Fault_ *soap.Fault                                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanGetAssociatedSpbmProfileBody) Fault() *soap.Fault { return b.Fault_ }

func VsanGetAssociatedSpbmProfile(ctx context.Context, r soap.RoundTripper, req *types.VsanGetAssociatedSpbmProfile) (*types.VsanGetAssociatedSpbmProfileResponse, error) {
	var reqBody, resBody VsanGetAssociatedSpbmProfileBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanQueryWhatIfEvacuationResultBody struct {
	Req    *types.VsanQueryWhatIfEvacuationResult         `xml:"urn:vsan VsanQueryWhatIfEvacuationResult,omitempty"`
	Res    *types.VsanQueryWhatIfEvacuationResultResponse `xml:"urn:vsan VsanQueryWhatIfEvacuationResultResponse,omitempty"`
	Fault_ *soap.Fault                                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanQueryWhatIfEvacuationResultBody) Fault() *soap.Fault { return b.Fault_ }

func VsanQueryWhatIfEvacuationResult(ctx context.Context, r soap.RoundTripper, req *types.VsanQueryWhatIfEvacuationResult) (*types.VsanQueryWhatIfEvacuationResultResponse, error) {
	var reqBody, resBody VsanQueryWhatIfEvacuationResultBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ClearVmMetadataBody struct {
	Req    *types.ClearVmMetadata         `xml:"urn:vsan ClearVmMetadata,omitempty"`
	Res    *types.ClearVmMetadataResponse `xml:"urn:vsan ClearVmMetadataResponse,omitempty"`
	Fault_ *soap.Fault                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ClearVmMetadataBody) Fault() *soap.Fault { return b.Fault_ }

func ClearVmMetadata(ctx context.Context, r soap.RoundTripper, req *types.ClearVmMetadata) (*types.ClearVmMetadataResponse, error) {
	var reqBody, resBody ClearVmMetadataBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type GetVsanNicsBody struct {
	Req    *types.GetVsanNics         `xml:"urn:vsan GetVsanNics,omitempty"`
	Res    *types.GetVsanNicsResponse `xml:"urn:vsan GetVsanNicsResponse,omitempty"`
	Fault_ *soap.Fault                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *GetVsanNicsBody) Fault() *soap.Fault { return b.Fault_ }

func GetVsanNics(ctx context.Context, r soap.RoundTripper, req *types.GetVsanNics) (*types.GetVsanNicsResponse, error) {
	var reqBody, resBody GetVsanNicsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanHostGetRuntimeStatsBody struct {
	Req    *types.VsanHostGetRuntimeStats         `xml:"urn:vsan VsanHostGetRuntimeStats,omitempty"`
	Res    *types.VsanHostGetRuntimeStatsResponse `xml:"urn:vsan VsanHostGetRuntimeStatsResponse,omitempty"`
	Fault_ *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanHostGetRuntimeStatsBody) Fault() *soap.Fault { return b.Fault_ }

func VsanHostGetRuntimeStats(ctx context.Context, r soap.RoundTripper, req *types.VsanHostGetRuntimeStats) (*types.VsanHostGetRuntimeStatsResponse, error) {
	var reqBody, resBody VsanHostGetRuntimeStatsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanGetAssociatedObjectsBody struct {
	Req    *types.VsanGetAssociatedObjects         `xml:"urn:vsan VsanGetAssociatedObjects,omitempty"`
	Res    *types.VsanGetAssociatedObjectsResponse `xml:"urn:vsan VsanGetAssociatedObjectsResponse,omitempty"`
	Fault_ *soap.Fault                             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanGetAssociatedObjectsBody) Fault() *soap.Fault { return b.Fault_ }

func VsanGetAssociatedObjects(ctx context.Context, r soap.RoundTripper, req *types.VsanGetAssociatedObjects) (*types.VsanGetAssociatedObjectsResponse, error) {
	var reqBody, resBody VsanGetAssociatedObjectsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VmMonitorWatchObjectsBody struct {
	Req    *types.VmMonitorWatchObjects         `xml:"urn:vsan VmMonitorWatchObjects,omitempty"`
	Res    *types.VmMonitorWatchObjectsResponse `xml:"urn:vsan VmMonitorWatchObjectsResponse,omitempty"`
	Fault_ *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VmMonitorWatchObjectsBody) Fault() *soap.Fault { return b.Fault_ }

func VmMonitorWatchObjects(ctx context.Context, r soap.RoundTripper, req *types.VmMonitorWatchObjects) (*types.VmMonitorWatchObjectsResponse, error) {
	var reqBody, resBody VmMonitorWatchObjectsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryVmsAccessibilityBody struct {
	Req    *types.QueryVmsAccessibility         `xml:"urn:vsan QueryVmsAccessibility,omitempty"`
	Res    *types.QueryVmsAccessibilityResponse `xml:"urn:vsan QueryVmsAccessibilityResponse,omitempty"`
	Fault_ *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryVmsAccessibilityBody) Fault() *soap.Fault { return b.Fault_ }

func QueryVmsAccessibility(ctx context.Context, r soap.RoundTripper, req *types.QueryVmsAccessibility) (*types.QueryVmsAccessibilityResponse, error) {
	var reqBody, resBody QueryVmsAccessibilityBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type GetDomObjectsBody struct {
	Req    *types.GetDomObjects         `xml:"urn:vsan GetDomObjects,omitempty"`
	Res    *types.GetDomObjectsResponse `xml:"urn:vsan GetDomObjectsResponse,omitempty"`
	Fault_ *soap.Fault                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *GetDomObjectsBody) Fault() *soap.Fault { return b.Fault_ }

func GetDomObjects(ctx context.Context, r soap.RoundTripper, req *types.GetDomObjects) (*types.GetDomObjectsResponse, error) {
	var reqBody, resBody GetDomObjectsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanEnterMaintenanceMode_TaskBody struct {
	Req    *types.VsanEnterMaintenanceMode_Task         `xml:"urn:vsan VsanEnterMaintenanceMode_Task,omitempty"`
	Res    *types.VsanEnterMaintenanceMode_TaskResponse `xml:"urn:vsan VsanEnterMaintenanceMode_TaskResponse,omitempty"`
	Fault_ *soap.Fault                                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanEnterMaintenanceMode_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func VsanEnterMaintenanceMode_Task(ctx context.Context, r soap.RoundTripper, req *types.VsanEnterMaintenanceMode_Task) (*types.VsanEnterMaintenanceMode_TaskResponse, error) {
	var reqBody, resBody VsanEnterMaintenanceMode_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type GetTopLevelDirectoriesBody struct {
	Req    *types.GetTopLevelDirectories         `xml:"urn:vsan GetTopLevelDirectories,omitempty"`
	Res    *types.GetTopLevelDirectoriesResponse `xml:"urn:vsan GetTopLevelDirectoriesResponse,omitempty"`
	Fault_ *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *GetTopLevelDirectoriesBody) Fault() *soap.Fault { return b.Fault_ }

func GetTopLevelDirectories(ctx context.Context, r soap.RoundTripper, req *types.GetTopLevelDirectories) (*types.GetTopLevelDirectoriesResponse, error) {
	var reqBody, resBody GetTopLevelDirectoriesBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type GetVsanRuntimeInfoBody struct {
	Req    *types.GetVsanRuntimeInfo         `xml:"urn:vsan GetVsanRuntimeInfo,omitempty"`
	Res    *types.GetVsanRuntimeInfoResponse `xml:"urn:vsan GetVsanRuntimeInfoResponse,omitempty"`
	Fault_ *soap.Fault                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *GetVsanRuntimeInfoBody) Fault() *soap.Fault { return b.Fault_ }

func GetVsanRuntimeInfo(ctx context.Context, r soap.RoundTripper, req *types.GetVsanRuntimeInfo) (*types.GetVsanRuntimeInfoResponse, error) {
	var reqBody, resBody GetVsanRuntimeInfoBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type GetDomClientRoleStatsBody struct {
	Req    *types.GetDomClientRoleStats         `xml:"urn:vsan GetDomClientRoleStats,omitempty"`
	Res    *types.GetDomClientRoleStatsResponse `xml:"urn:vsan GetDomClientRoleStatsResponse,omitempty"`
	Fault_ *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *GetDomClientRoleStatsBody) Fault() *soap.Fault { return b.Fault_ }

func GetDomClientRoleStats(ctx context.Context, r soap.RoundTripper, req *types.GetDomClientRoleStats) (*types.GetDomClientRoleStatsResponse, error) {
	var reqBody, resBody GetDomClientRoleStatsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanGetAboutInfoExBody struct {
	Req    *types.VsanGetAboutInfoEx         `xml:"urn:vsan VsanGetAboutInfoEx,omitempty"`
	Res    *types.VsanGetAboutInfoExResponse `xml:"urn:vsan VsanGetAboutInfoExResponse,omitempty"`
	Fault_ *soap.Fault                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanGetAboutInfoExBody) Fault() *soap.Fault { return b.Fault_ }

func VsanGetAboutInfoEx(ctx context.Context, r soap.RoundTripper, req *types.VsanGetAboutInfoEx) (*types.VsanGetAboutInfoExResponse, error) {
	var reqBody, resBody VsanGetAboutInfoExBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryVmMetadataBody struct {
	Req    *types.QueryVmMetadata         `xml:"urn:vsan QueryVmMetadata,omitempty"`
	Res    *types.QueryVmMetadataResponse `xml:"urn:vsan QueryVmMetadataResponse,omitempty"`
	Fault_ *soap.Fault                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryVmMetadataBody) Fault() *soap.Fault { return b.Fault_ }

func QueryVmMetadata(ctx context.Context, r soap.RoundTripper, req *types.QueryVmMetadata) (*types.QueryVmMetadataResponse, error) {
	var reqBody, resBody QueryVmMetadataBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RefreshVsanVmknicsBody struct {
	Req    *types.RefreshVsanVmknics         `xml:"urn:vsan RefreshVsanVmknics,omitempty"`
	Res    *types.RefreshVsanVmknicsResponse `xml:"urn:vsan RefreshVsanVmknicsResponse,omitempty"`
	Fault_ *soap.Fault                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RefreshVsanVmknicsBody) Fault() *soap.Fault { return b.Fault_ }

func RefreshVsanVmknics(ctx context.Context, r soap.RoundTripper, req *types.RefreshVsanVmknics) (*types.RefreshVsanVmknicsResponse, error) {
	var reqBody, resBody RefreshVsanVmknicsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type GetDomObjectStatsBody struct {
	Req    *types.GetDomObjectStats         `xml:"urn:vsan GetDomObjectStats,omitempty"`
	Res    *types.GetDomObjectStatsResponse `xml:"urn:vsan GetDomObjectStatsResponse,omitempty"`
	Fault_ *soap.Fault                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *GetDomObjectStatsBody) Fault() *soap.Fault { return b.Fault_ }

func GetDomObjectStats(ctx context.Context, r soap.RoundTripper, req *types.GetDomObjectStats) (*types.GetDomObjectStatsResponse, error) {
	var reqBody, resBody GetDomObjectStatsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryObjectsAccessibilityBody struct {
	Req    *types.QueryObjectsAccessibility         `xml:"urn:vsan QueryObjectsAccessibility,omitempty"`
	Res    *types.QueryObjectsAccessibilityResponse `xml:"urn:vsan QueryObjectsAccessibilityResponse,omitempty"`
	Fault_ *soap.Fault                              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryObjectsAccessibilityBody) Fault() *soap.Fault { return b.Fault_ }

func QueryObjectsAccessibility(ctx context.Context, r soap.RoundTripper, req *types.QueryObjectsAccessibility) (*types.QueryObjectsAccessibilityResponse, error) {
	var reqBody, resBody QueryObjectsAccessibilityBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanClusterReconfigBody struct {
	Req    *types.VsanClusterReconfig         `xml:"urn:vsan VsanClusterReconfig,omitempty"`
	Res    *types.VsanClusterReconfigResponse `xml:"urn:vsan VsanClusterReconfigResponse,omitempty"`
	Fault_ *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanClusterReconfigBody) Fault() *soap.Fault { return b.Fault_ }

func VsanClusterReconfig(ctx context.Context, r soap.RoundTripper, req *types.VsanClusterReconfig) (*types.VsanClusterReconfigResponse, error) {
	var reqBody, resBody VsanClusterReconfigBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanGetDataProtectionLoadBalancersBody struct {
	Req    *types.VsanGetDataProtectionLoadBalancers         `xml:"urn:vsan VsanGetDataProtectionLoadBalancers,omitempty"`
	Res    *types.VsanGetDataProtectionLoadBalancersResponse `xml:"urn:vsan VsanGetDataProtectionLoadBalancersResponse,omitempty"`
	Fault_ *soap.Fault                                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanGetDataProtectionLoadBalancersBody) Fault() *soap.Fault { return b.Fault_ }

func VsanGetDataProtectionLoadBalancers(ctx context.Context, r soap.RoundTripper, req *types.VsanGetDataProtectionLoadBalancers) (*types.VsanGetDataProtectionLoadBalancersResponse, error) {
	var reqBody, resBody VsanGetDataProtectionLoadBalancersBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanClusterGetRuntimeStatsBody struct {
	Req    *types.VsanClusterGetRuntimeStats         `xml:"urn:vsan VsanClusterGetRuntimeStats,omitempty"`
	Res    *types.VsanClusterGetRuntimeStatsResponse `xml:"urn:vsan VsanClusterGetRuntimeStatsResponse,omitempty"`
	Fault_ *soap.Fault                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanClusterGetRuntimeStatsBody) Fault() *soap.Fault { return b.Fault_ }

func VsanClusterGetRuntimeStats(ctx context.Context, r soap.RoundTripper, req *types.VsanClusterGetRuntimeStats) (*types.VsanClusterGetRuntimeStatsResponse, error) {
	var reqBody, resBody VsanClusterGetRuntimeStatsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanQueryClusterDrsStatsBody struct {
	Req    *types.VsanQueryClusterDrsStats         `xml:"urn:vsan VsanQueryClusterDrsStats,omitempty"`
	Res    *types.VsanQueryClusterDrsStatsResponse `xml:"urn:vsan VsanQueryClusterDrsStatsResponse,omitempty"`
	Fault_ *soap.Fault                             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanQueryClusterDrsStatsBody) Fault() *soap.Fault { return b.Fault_ }

func VsanQueryClusterDrsStats(ctx context.Context, r soap.RoundTripper, req *types.VsanQueryClusterDrsStats) (*types.VsanQueryClusterDrsStatsResponse, error) {
	var reqBody, resBody VsanQueryClusterDrsStatsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanFindClusterByDatastoreUrlBody struct {
	Req    *types.VsanFindClusterByDatastoreUrl         `xml:"urn:vsan VsanFindClusterByDatastoreUrl,omitempty"`
	Res    *types.VsanFindClusterByDatastoreUrlResponse `xml:"urn:vsan VsanFindClusterByDatastoreUrlResponse,omitempty"`
	Fault_ *soap.Fault                                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanFindClusterByDatastoreUrlBody) Fault() *soap.Fault { return b.Fault_ }

func VsanFindClusterByDatastoreUrl(ctx context.Context, r soap.RoundTripper, req *types.VsanFindClusterByDatastoreUrl) (*types.VsanFindClusterByDatastoreUrlResponse, error) {
	var reqBody, resBody VsanFindClusterByDatastoreUrlBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanClusterGetConfigBody struct {
	Req    *types.VsanClusterGetConfig         `xml:"urn:vsan VsanClusterGetConfig,omitempty"`
	Res    *types.VsanClusterGetConfigResponse `xml:"urn:vsan VsanClusterGetConfigResponse,omitempty"`
	Fault_ *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanClusterGetConfigBody) Fault() *soap.Fault { return b.Fault_ }

func VsanClusterGetConfig(ctx context.Context, r soap.RoundTripper, req *types.VsanClusterGetConfig) (*types.VsanClusterGetConfigResponse, error) {
	var reqBody, resBody VsanClusterGetConfigBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanEncryptedClusterRekey_TaskBody struct {
	Req    *types.VsanEncryptedClusterRekey_Task         `xml:"urn:vsan VsanEncryptedClusterRekey_Task,omitempty"`
	Res    *types.VsanEncryptedClusterRekey_TaskResponse `xml:"urn:vsan VsanEncryptedClusterRekey_TaskResponse,omitempty"`
	Fault_ *soap.Fault                                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanEncryptedClusterRekey_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func VsanEncryptedClusterRekey_Task(ctx context.Context, r soap.RoundTripper, req *types.VsanEncryptedClusterRekey_Task) (*types.VsanEncryptedClusterRekey_TaskResponse, error) {
	var reqBody, resBody VsanEncryptedClusterRekey_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanFlrSessionQueryVolumesBody struct {
	Req    *types.VsanFlrSessionQueryVolumes         `xml:"urn:vsan VsanFlrSessionQueryVolumes,omitempty"`
	Res    *types.VsanFlrSessionQueryVolumesResponse `xml:"urn:vsan VsanFlrSessionQueryVolumesResponse,omitempty"`
	Fault_ *soap.Fault                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanFlrSessionQueryVolumesBody) Fault() *soap.Fault { return b.Fault_ }

func VsanFlrSessionQueryVolumes(ctx context.Context, r soap.RoundTripper, req *types.VsanFlrSessionQueryVolumes) (*types.VsanFlrSessionQueryVolumesResponse, error) {
	var reqBody, resBody VsanFlrSessionQueryVolumesBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanFlrSessionRetrieveFileBody struct {
	Req    *types.VsanFlrSessionRetrieveFile         `xml:"urn:vsan VsanFlrSessionRetrieveFile,omitempty"`
	Res    *types.VsanFlrSessionRetrieveFileResponse `xml:"urn:vsan VsanFlrSessionRetrieveFileResponse,omitempty"`
	Fault_ *soap.Fault                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanFlrSessionRetrieveFileBody) Fault() *soap.Fault { return b.Fault_ }

func VsanFlrSessionRetrieveFile(ctx context.Context, r soap.RoundTripper, req *types.VsanFlrSessionRetrieveFile) (*types.VsanFlrSessionRetrieveFileResponse, error) {
	var reqBody, resBody VsanFlrSessionRetrieveFileBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanFlrSessionQueryFileBody struct {
	Req    *types.VsanFlrSessionQueryFile         `xml:"urn:vsan VsanFlrSessionQueryFile,omitempty"`
	Res    *types.VsanFlrSessionQueryFileResponse `xml:"urn:vsan VsanFlrSessionQueryFileResponse,omitempty"`
	Fault_ *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanFlrSessionQueryFileBody) Fault() *soap.Fault { return b.Fault_ }

func VsanFlrSessionQueryFile(ctx context.Context, r soap.RoundTripper, req *types.VsanFlrSessionQueryFile) (*types.VsanFlrSessionQueryFileResponse, error) {
	var reqBody, resBody VsanFlrSessionQueryFileBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanDownloadFileServiceOvfBody struct {
	Req    *types.VsanDownloadFileServiceOvf         `xml:"urn:vsan VsanDownloadFileServiceOvf,omitempty"`
	Res    *types.VsanDownloadFileServiceOvfResponse `xml:"urn:vsan VsanDownloadFileServiceOvfResponse,omitempty"`
	Fault_ *soap.Fault                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanDownloadFileServiceOvfBody) Fault() *soap.Fault { return b.Fault_ }

func VsanDownloadFileServiceOvf(ctx context.Context, r soap.RoundTripper, req *types.VsanDownloadFileServiceOvf) (*types.VsanDownloadFileServiceOvfResponse, error) {
	var reqBody, resBody VsanDownloadFileServiceOvfBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanClusterCreateFsDomainBody struct {
	Req    *types.VsanClusterCreateFsDomain         `xml:"urn:vsan VsanClusterCreateFsDomain,omitempty"`
	Res    *types.VsanClusterCreateFsDomainResponse `xml:"urn:vsan VsanClusterCreateFsDomainResponse,omitempty"`
	Fault_ *soap.Fault                              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanClusterCreateFsDomainBody) Fault() *soap.Fault { return b.Fault_ }

func VsanClusterCreateFsDomain(ctx context.Context, r soap.RoundTripper, req *types.VsanClusterCreateFsDomain) (*types.VsanClusterCreateFsDomainResponse, error) {
	var reqBody, resBody VsanClusterCreateFsDomainBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanPerformFileServiceEnablePreflightCheckBody struct {
	Req    *types.VsanPerformFileServiceEnablePreflightCheck         `xml:"urn:vsan VsanPerformFileServiceEnablePreflightCheck,omitempty"`
	Res    *types.VsanPerformFileServiceEnablePreflightCheckResponse `xml:"urn:vsan VsanPerformFileServiceEnablePreflightCheckResponse,omitempty"`
	Fault_ *soap.Fault                                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanPerformFileServiceEnablePreflightCheckBody) Fault() *soap.Fault { return b.Fault_ }

func VsanPerformFileServiceEnablePreflightCheck(ctx context.Context, r soap.RoundTripper, req *types.VsanPerformFileServiceEnablePreflightCheck) (*types.VsanPerformFileServiceEnablePreflightCheckResponse, error) {
	var reqBody, resBody VsanPerformFileServiceEnablePreflightCheckBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanRemediateFileServiceBody struct {
	Req    *types.VsanRemediateFileService         `xml:"urn:vsan VsanRemediateFileService,omitempty"`
	Res    *types.VsanRemediateFileServiceResponse `xml:"urn:vsan VsanRemediateFileServiceResponse,omitempty"`
	Fault_ *soap.Fault                             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanRemediateFileServiceBody) Fault() *soap.Fault { return b.Fault_ }

func VsanRemediateFileService(ctx context.Context, r soap.RoundTripper, req *types.VsanRemediateFileService) (*types.VsanRemediateFileServiceResponse, error) {
	var reqBody, resBody VsanRemediateFileServiceBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanClusterQueryFsDomainsBody struct {
	Req    *types.VsanClusterQueryFsDomains         `xml:"urn:vsan VsanClusterQueryFsDomains,omitempty"`
	Res    *types.VsanClusterQueryFsDomainsResponse `xml:"urn:vsan VsanClusterQueryFsDomainsResponse,omitempty"`
	Fault_ *soap.Fault                              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanClusterQueryFsDomainsBody) Fault() *soap.Fault { return b.Fault_ }

func VsanClusterQueryFsDomains(ctx context.Context, r soap.RoundTripper, req *types.VsanClusterQueryFsDomains) (*types.VsanClusterQueryFsDomainsResponse, error) {
	var reqBody, resBody VsanClusterQueryFsDomainsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanQueryFileServiceOvfsBody struct {
	Req    *types.VsanQueryFileServiceOvfs         `xml:"urn:vsan VsanQueryFileServiceOvfs,omitempty"`
	Res    *types.VsanQueryFileServiceOvfsResponse `xml:"urn:vsan VsanQueryFileServiceOvfsResponse,omitempty"`
	Fault_ *soap.Fault                             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanQueryFileServiceOvfsBody) Fault() *soap.Fault { return b.Fault_ }

func VsanQueryFileServiceOvfs(ctx context.Context, r soap.RoundTripper, req *types.VsanQueryFileServiceOvfs) (*types.VsanQueryFileServiceOvfsResponse, error) {
	var reqBody, resBody VsanQueryFileServiceOvfsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanReconfigureFileShareBody struct {
	Req    *types.VsanReconfigureFileShare         `xml:"urn:vsan VsanReconfigureFileShare,omitempty"`
	Res    *types.VsanReconfigureFileShareResponse `xml:"urn:vsan VsanReconfigureFileShareResponse,omitempty"`
	Fault_ *soap.Fault                             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanReconfigureFileShareBody) Fault() *soap.Fault { return b.Fault_ }

func VsanReconfigureFileShare(ctx context.Context, r soap.RoundTripper, req *types.VsanReconfigureFileShare) (*types.VsanReconfigureFileShareResponse, error) {
	var reqBody, resBody VsanReconfigureFileShareBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanClusterReconfigureFsDomainBody struct {
	Req    *types.VsanClusterReconfigureFsDomain         `xml:"urn:vsan VsanClusterReconfigureFsDomain,omitempty"`
	Res    *types.VsanClusterReconfigureFsDomainResponse `xml:"urn:vsan VsanClusterReconfigureFsDomainResponse,omitempty"`
	Fault_ *soap.Fault                                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanClusterReconfigureFsDomainBody) Fault() *soap.Fault { return b.Fault_ }

func VsanClusterReconfigureFsDomain(ctx context.Context, r soap.RoundTripper, req *types.VsanClusterReconfigureFsDomain) (*types.VsanClusterReconfigureFsDomainResponse, error) {
	var reqBody, resBody VsanClusterReconfigureFsDomainBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanFindOvfDownloadUrlBody struct {
	Req    *types.VsanFindOvfDownloadUrl         `xml:"urn:vsan VsanFindOvfDownloadUrl,omitempty"`
	Res    *types.VsanFindOvfDownloadUrlResponse `xml:"urn:vsan VsanFindOvfDownloadUrlResponse,omitempty"`
	Fault_ *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanFindOvfDownloadUrlBody) Fault() *soap.Fault { return b.Fault_ }

func VsanFindOvfDownloadUrl(ctx context.Context, r soap.RoundTripper, req *types.VsanFindOvfDownloadUrl) (*types.VsanFindOvfDownloadUrlResponse, error) {
	var reqBody, resBody VsanFindOvfDownloadUrlBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanClusterQueryFileSharesBody struct {
	Req    *types.VsanClusterQueryFileShares         `xml:"urn:vsan VsanClusterQueryFileShares,omitempty"`
	Res    *types.VsanClusterQueryFileSharesResponse `xml:"urn:vsan VsanClusterQueryFileSharesResponse,omitempty"`
	Fault_ *soap.Fault                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanClusterQueryFileSharesBody) Fault() *soap.Fault { return b.Fault_ }

func VsanClusterQueryFileShares(ctx context.Context, r soap.RoundTripper, req *types.VsanClusterQueryFileShares) (*types.VsanClusterQueryFileSharesResponse, error) {
	var reqBody, resBody VsanClusterQueryFileSharesBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanClusterRemoveShareBody struct {
	Req    *types.VsanClusterRemoveShare         `xml:"urn:vsan VsanClusterRemoveShare,omitempty"`
	Res    *types.VsanClusterRemoveShareResponse `xml:"urn:vsan VsanClusterRemoveShareResponse,omitempty"`
	Fault_ *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanClusterRemoveShareBody) Fault() *soap.Fault { return b.Fault_ }

func VsanClusterRemoveShare(ctx context.Context, r soap.RoundTripper, req *types.VsanClusterRemoveShare) (*types.VsanClusterRemoveShareResponse, error) {
	var reqBody, resBody VsanClusterRemoveShareBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanClusterRemoveFsDomainBody struct {
	Req    *types.VsanClusterRemoveFsDomain         `xml:"urn:vsan VsanClusterRemoveFsDomain,omitempty"`
	Res    *types.VsanClusterRemoveFsDomainResponse `xml:"urn:vsan VsanClusterRemoveFsDomainResponse,omitempty"`
	Fault_ *soap.Fault                              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanClusterRemoveFsDomainBody) Fault() *soap.Fault { return b.Fault_ }

func VsanClusterRemoveFsDomain(ctx context.Context, r soap.RoundTripper, req *types.VsanClusterRemoveFsDomain) (*types.VsanClusterRemoveFsDomainResponse, error) {
	var reqBody, resBody VsanClusterRemoveFsDomainBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanCreateFileShareBody struct {
	Req    *types.VsanCreateFileShare         `xml:"urn:vsan VsanCreateFileShare,omitempty"`
	Res    *types.VsanCreateFileShareResponse `xml:"urn:vsan VsanCreateFileShareResponse,omitempty"`
	Fault_ *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanCreateFileShareBody) Fault() *soap.Fault { return b.Fault_ }

func VsanCreateFileShare(ctx context.Context, r soap.RoundTripper, req *types.VsanCreateFileShare) (*types.VsanCreateFileShareResponse, error) {
	var reqBody, resBody VsanCreateFileShareBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanHostQueryAdvCfgBody struct {
	Req    *types.VsanHostQueryAdvCfg         `xml:"urn:vsan VsanHostQueryAdvCfg,omitempty"`
	Res    *types.VsanHostQueryAdvCfgResponse `xml:"urn:vsan VsanHostQueryAdvCfgResponse,omitempty"`
	Fault_ *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanHostQueryAdvCfgBody) Fault() *soap.Fault { return b.Fault_ }

func VsanHostQueryAdvCfg(ctx context.Context, r soap.RoundTripper, req *types.VsanHostQueryAdvCfg) (*types.VsanHostQueryAdvCfgResponse, error) {
	var reqBody, resBody VsanHostQueryAdvCfgBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanHostQueryRunIperfClientBody struct {
	Req    *types.VsanHostQueryRunIperfClient         `xml:"urn:vsan VsanHostQueryRunIperfClient,omitempty"`
	Res    *types.VsanHostQueryRunIperfClientResponse `xml:"urn:vsan VsanHostQueryRunIperfClientResponse,omitempty"`
	Fault_ *soap.Fault                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanHostQueryRunIperfClientBody) Fault() *soap.Fault { return b.Fault_ }

func VsanHostQueryRunIperfClient(ctx context.Context, r soap.RoundTripper, req *types.VsanHostQueryRunIperfClient) (*types.VsanHostQueryRunIperfClientResponse, error) {
	var reqBody, resBody VsanHostQueryRunIperfClientBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanQueryVsanConfigsByFilterBody struct {
	Req    *types.VsanQueryVsanConfigsByFilter         `xml:"urn:vsan VsanQueryVsanConfigsByFilter,omitempty"`
	Res    *types.VsanQueryVsanConfigsByFilterResponse `xml:"urn:vsan VsanQueryVsanConfigsByFilterResponse,omitempty"`
	Fault_ *soap.Fault                                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanQueryVsanConfigsByFilterBody) Fault() *soap.Fault { return b.Fault_ }

func VsanQueryVsanConfigsByFilter(ctx context.Context, r soap.RoundTripper, req *types.VsanQueryVsanConfigsByFilter) (*types.VsanQueryVsanConfigsByFilterResponse, error) {
	var reqBody, resBody VsanQueryVsanConfigsByFilterBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanHostQueryObjectHealthSummaryBody struct {
	Req    *types.VsanHostQueryObjectHealthSummary         `xml:"urn:vsan VsanHostQueryObjectHealthSummary,omitempty"`
	Res    *types.VsanHostQueryObjectHealthSummaryResponse `xml:"urn:vsan VsanHostQueryObjectHealthSummaryResponse,omitempty"`
	Fault_ *soap.Fault                                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanHostQueryObjectHealthSummaryBody) Fault() *soap.Fault { return b.Fault_ }

func VsanHostQueryObjectHealthSummary(ctx context.Context, r soap.RoundTripper, req *types.VsanHostQueryObjectHealthSummary) (*types.VsanHostQueryObjectHealthSummaryResponse, error) {
	var reqBody, resBody VsanHostQueryObjectHealthSummaryBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RefreshVsanHealthGenerationIdBody struct {
	Req    *types.RefreshVsanHealthGenerationId         `xml:"urn:vsan RefreshVsanHealthGenerationId,omitempty"`
	Res    *types.RefreshVsanHealthGenerationIdResponse `xml:"urn:vsan RefreshVsanHealthGenerationIdResponse,omitempty"`
	Fault_ *soap.Fault                                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RefreshVsanHealthGenerationIdBody) Fault() *soap.Fault { return b.Fault_ }

func RefreshVsanHealthGenerationId(ctx context.Context, r soap.RoundTripper, req *types.RefreshVsanHealthGenerationId) (*types.RefreshVsanHealthGenerationIdResponse, error) {
	var reqBody, resBody RefreshVsanHealthGenerationIdBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanStopProactiveRebalanceBody struct {
	Req    *types.VsanStopProactiveRebalance         `xml:"urn:vsan VsanStopProactiveRebalance,omitempty"`
	Res    *types.VsanStopProactiveRebalanceResponse `xml:"urn:vsan VsanStopProactiveRebalanceResponse,omitempty"`
	Fault_ *soap.Fault                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanStopProactiveRebalanceBody) Fault() *soap.Fault { return b.Fault_ }

func VsanStopProactiveRebalance(ctx context.Context, r soap.RoundTripper, req *types.VsanStopProactiveRebalance) (*types.VsanStopProactiveRebalanceResponse, error) {
	var reqBody, resBody VsanStopProactiveRebalanceBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanLocalPropertyCollectorBody struct {
	Req    *types.VsanLocalPropertyCollector         `xml:"urn:vsan VsanLocalPropertyCollector,omitempty"`
	Res    *types.VsanLocalPropertyCollectorResponse `xml:"urn:vsan VsanLocalPropertyCollectorResponse,omitempty"`
	Fault_ *soap.Fault                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanLocalPropertyCollectorBody) Fault() *soap.Fault { return b.Fault_ }

func VsanLocalPropertyCollector(ctx context.Context, r soap.RoundTripper, req *types.VsanLocalPropertyCollector) (*types.VsanLocalPropertyCollectorResponse, error) {
	var reqBody, resBody VsanLocalPropertyCollectorBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanHostQueryFileServiceHealthSummaryBody struct {
	Req    *types.VsanHostQueryFileServiceHealthSummary         `xml:"urn:vsan VsanHostQueryFileServiceHealthSummary,omitempty"`
	Res    *types.VsanHostQueryFileServiceHealthSummaryResponse `xml:"urn:vsan VsanHostQueryFileServiceHealthSummaryResponse,omitempty"`
	Fault_ *soap.Fault                                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanHostQueryFileServiceHealthSummaryBody) Fault() *soap.Fault { return b.Fault_ }

func VsanHostQueryFileServiceHealthSummary(ctx context.Context, r soap.RoundTripper, req *types.VsanHostQueryFileServiceHealthSummary) (*types.VsanHostQueryFileServiceHealthSummaryResponse, error) {
	var reqBody, resBody VsanHostQueryFileServiceHealthSummaryBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanHostQueryDataProtectionCfgBody struct {
	Req    *types.VsanHostQueryDataProtectionCfg         `xml:"urn:vsan VsanHostQueryDataProtectionCfg,omitempty"`
	Res    *types.VsanHostQueryDataProtectionCfgResponse `xml:"urn:vsan VsanHostQueryDataProtectionCfgResponse,omitempty"`
	Fault_ *soap.Fault                                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanHostQueryDataProtectionCfgBody) Fault() *soap.Fault { return b.Fault_ }

func VsanHostQueryDataProtectionCfg(ctx context.Context, r soap.RoundTripper, req *types.VsanHostQueryDataProtectionCfg) (*types.VsanHostQueryDataProtectionCfgResponse, error) {
	var reqBody, resBody VsanHostQueryDataProtectionCfgBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanLocalHostSystemBody struct {
	Req    *types.VsanLocalHostSystem         `xml:"urn:vsan VsanLocalHostSystem,omitempty"`
	Res    *types.VsanLocalHostSystemResponse `xml:"urn:vsan VsanLocalHostSystemResponse,omitempty"`
	Fault_ *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanLocalHostSystemBody) Fault() *soap.Fault { return b.Fault_ }

func VsanLocalHostSystem(ctx context.Context, r soap.RoundTripper, req *types.VsanLocalHostSystem) (*types.VsanLocalHostSystemResponse, error) {
	var reqBody, resBody VsanLocalHostSystemBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanHostClomdLivenessBody struct {
	Req    *types.VsanHostClomdLiveness         `xml:"urn:vsan VsanHostClomdLiveness,omitempty"`
	Res    *types.VsanHostClomdLivenessResponse `xml:"urn:vsan VsanHostClomdLivenessResponse,omitempty"`
	Fault_ *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanHostClomdLivenessBody) Fault() *soap.Fault { return b.Fault_ }

func VsanHostClomdLiveness(ctx context.Context, r soap.RoundTripper, req *types.VsanHostClomdLiveness) (*types.VsanHostClomdLivenessResponse, error) {
	var reqBody, resBody VsanHostClomdLivenessBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanHostRepairImmediateObjectsBody struct {
	Req    *types.VsanHostRepairImmediateObjects         `xml:"urn:vsan VsanHostRepairImmediateObjects,omitempty"`
	Res    *types.VsanHostRepairImmediateObjectsResponse `xml:"urn:vsan VsanHostRepairImmediateObjectsResponse,omitempty"`
	Fault_ *soap.Fault                                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanHostRepairImmediateObjectsBody) Fault() *soap.Fault { return b.Fault_ }

func VsanHostRepairImmediateObjects(ctx context.Context, r soap.RoundTripper, req *types.VsanHostRepairImmediateObjects) (*types.VsanHostRepairImmediateObjectsResponse, error) {
	var reqBody, resBody VsanHostRepairImmediateObjectsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanHostQueryVerifyNetworkSettingsBody struct {
	Req    *types.VsanHostQueryVerifyNetworkSettings         `xml:"urn:vsan VsanHostQueryVerifyNetworkSettings,omitempty"`
	Res    *types.VsanHostQueryVerifyNetworkSettingsResponse `xml:"urn:vsan VsanHostQueryVerifyNetworkSettingsResponse,omitempty"`
	Fault_ *soap.Fault                                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanHostQueryVerifyNetworkSettingsBody) Fault() *soap.Fault { return b.Fault_ }

func VsanHostQueryVerifyNetworkSettings(ctx context.Context, r soap.RoundTripper, req *types.VsanHostQueryVerifyNetworkSettings) (*types.VsanHostQueryVerifyNetworkSettingsResponse, error) {
	var reqBody, resBody VsanHostQueryVerifyNetworkSettingsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanHostCleanupVmdkLoadTestBody struct {
	Req    *types.VsanHostCleanupVmdkLoadTest         `xml:"urn:vsan VsanHostCleanupVmdkLoadTest,omitempty"`
	Res    *types.VsanHostCleanupVmdkLoadTestResponse `xml:"urn:vsan VsanHostCleanupVmdkLoadTestResponse,omitempty"`
	Fault_ *soap.Fault                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanHostCleanupVmdkLoadTestBody) Fault() *soap.Fault { return b.Fault_ }

func VsanHostCleanupVmdkLoadTest(ctx context.Context, r soap.RoundTripper, req *types.VsanHostCleanupVmdkLoadTest) (*types.VsanHostCleanupVmdkLoadTestResponse, error) {
	var reqBody, resBody VsanHostCleanupVmdkLoadTestBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanStartProactiveRebalanceBody struct {
	Req    *types.VsanStartProactiveRebalance         `xml:"urn:vsan VsanStartProactiveRebalance,omitempty"`
	Res    *types.VsanStartProactiveRebalanceResponse `xml:"urn:vsan VsanStartProactiveRebalanceResponse,omitempty"`
	Fault_ *soap.Fault                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanStartProactiveRebalanceBody) Fault() *soap.Fault { return b.Fault_ }

func VsanStartProactiveRebalance(ctx context.Context, r soap.RoundTripper, req *types.VsanStartProactiveRebalance) (*types.VsanStartProactiveRebalanceResponse, error) {
	var reqBody, resBody VsanStartProactiveRebalanceBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanHostQueryEncryptionHealthSummaryBody struct {
	Req    *types.VsanHostQueryEncryptionHealthSummary         `xml:"urn:vsan VsanHostQueryEncryptionHealthSummary,omitempty"`
	Res    *types.VsanHostQueryEncryptionHealthSummaryResponse `xml:"urn:vsan VsanHostQueryEncryptionHealthSummaryResponse,omitempty"`
	Fault_ *soap.Fault                                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanHostQueryEncryptionHealthSummaryBody) Fault() *soap.Fault { return b.Fault_ }

func VsanHostQueryEncryptionHealthSummary(ctx context.Context, r soap.RoundTripper, req *types.VsanHostQueryEncryptionHealthSummary) (*types.VsanHostQueryEncryptionHealthSummaryResponse, error) {
	var reqBody, resBody VsanHostQueryEncryptionHealthSummaryBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanFlashScsiControllerFirmware_TaskBody struct {
	Req    *types.VsanFlashScsiControllerFirmware_Task         `xml:"urn:vsan VsanFlashScsiControllerFirmware_Task,omitempty"`
	Res    *types.VsanFlashScsiControllerFirmware_TaskResponse `xml:"urn:vsan VsanFlashScsiControllerFirmware_TaskResponse,omitempty"`
	Fault_ *soap.Fault                                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanFlashScsiControllerFirmware_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func VsanFlashScsiControllerFirmware_Task(ctx context.Context, r soap.RoundTripper, req *types.VsanFlashScsiControllerFirmware_Task) (*types.VsanFlashScsiControllerFirmware_TaskResponse, error) {
	var reqBody, resBody VsanFlashScsiControllerFirmware_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type PrepareHostForClusterRebootWithNAMMBody struct {
	Req    *types.PrepareHostForClusterRebootWithNAMM         `xml:"urn:vsan PrepareHostForClusterRebootWithNAMM,omitempty"`
	Res    *types.PrepareHostForClusterRebootWithNAMMResponse `xml:"urn:vsan PrepareHostForClusterRebootWithNAMMResponse,omitempty"`
	Fault_ *soap.Fault                                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *PrepareHostForClusterRebootWithNAMMBody) Fault() *soap.Fault { return b.Fault_ }

func PrepareHostForClusterRebootWithNAMM(ctx context.Context, r soap.RoundTripper, req *types.PrepareHostForClusterRebootWithNAMM) (*types.PrepareHostForClusterRebootWithNAMMResponse, error) {
	var reqBody, resBody PrepareHostForClusterRebootWithNAMMBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanQueryHostEMMStateBody struct {
	Req    *types.VsanQueryHostEMMState         `xml:"urn:vsan VsanQueryHostEMMState,omitempty"`
	Res    *types.VsanQueryHostEMMStateResponse `xml:"urn:vsan VsanQueryHostEMMStateResponse,omitempty"`
	Fault_ *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanQueryHostEMMStateBody) Fault() *soap.Fault { return b.Fault_ }

func VsanQueryHostEMMState(ctx context.Context, r soap.RoundTripper, req *types.VsanQueryHostEMMState) (*types.VsanQueryHostEMMStateResponse, error) {
	var reqBody, resBody VsanQueryHostEMMStateBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanRemediateScsiControllerIssuesBody struct {
	Req    *types.VsanRemediateScsiControllerIssues         `xml:"urn:vsan VsanRemediateScsiControllerIssues,omitempty"`
	Res    *types.VsanRemediateScsiControllerIssuesResponse `xml:"urn:vsan VsanRemediateScsiControllerIssuesResponse,omitempty"`
	Fault_ *soap.Fault                                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanRemediateScsiControllerIssuesBody) Fault() *soap.Fault { return b.Fault_ }

func VsanRemediateScsiControllerIssues(ctx context.Context, r soap.RoundTripper, req *types.VsanRemediateScsiControllerIssues) (*types.VsanRemediateScsiControllerIssuesResponse, error) {
	var reqBody, resBody VsanRemediateScsiControllerIssuesBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanHostQueryVmProcessListBody struct {
	Req    *types.VsanHostQueryVmProcessList         `xml:"urn:vsan VsanHostQueryVmProcessList,omitempty"`
	Res    *types.VsanHostQueryVmProcessListResponse `xml:"urn:vsan VsanHostQueryVmProcessListResponse,omitempty"`
	Fault_ *soap.Fault                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanHostQueryVmProcessListBody) Fault() *soap.Fault { return b.Fault_ }

func VsanHostQueryVmProcessList(ctx context.Context, r soap.RoundTripper, req *types.VsanHostQueryVmProcessList) (*types.VsanHostQueryVmProcessListResponse, error) {
	var reqBody, resBody VsanHostQueryVmProcessListBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanWaitForVsanHealthGenerationIdChangeBody struct {
	Req    *types.VsanWaitForVsanHealthGenerationIdChange         `xml:"urn:vsan VsanWaitForVsanHealthGenerationIdChange,omitempty"`
	Res    *types.VsanWaitForVsanHealthGenerationIdChangeResponse `xml:"urn:vsan VsanWaitForVsanHealthGenerationIdChangeResponse,omitempty"`
	Fault_ *soap.Fault                                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanWaitForVsanHealthGenerationIdChangeBody) Fault() *soap.Fault { return b.Fault_ }

func VsanWaitForVsanHealthGenerationIdChange(ctx context.Context, r soap.RoundTripper, req *types.VsanWaitForVsanHealthGenerationIdChange) (*types.VsanWaitForVsanHealthGenerationIdChangeResponse, error) {
	var reqBody, resBody VsanWaitForVsanHealthGenerationIdChangeBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanHostQueryHealthSystemVersionBody struct {
	Req    *types.VsanHostQueryHealthSystemVersion         `xml:"urn:vsan VsanHostQueryHealthSystemVersion,omitempty"`
	Res    *types.VsanHostQueryHealthSystemVersionResponse `xml:"urn:vsan VsanHostQueryHealthSystemVersionResponse,omitempty"`
	Fault_ *soap.Fault                                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanHostQueryHealthSystemVersionBody) Fault() *soap.Fault { return b.Fault_ }

func VsanHostQueryHealthSystemVersion(ctx context.Context, r soap.RoundTripper, req *types.VsanHostQueryHealthSystemVersion) (*types.VsanHostQueryHealthSystemVersionResponse, error) {
	var reqBody, resBody VsanHostQueryHealthSystemVersionBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanGetHclInfoBody struct {
	Req    *types.VsanGetHclInfo         `xml:"urn:vsan VsanGetHclInfo,omitempty"`
	Res    *types.VsanGetHclInfoResponse `xml:"urn:vsan VsanGetHclInfoResponse,omitempty"`
	Fault_ *soap.Fault                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanGetHclInfoBody) Fault() *soap.Fault { return b.Fault_ }

func VsanGetHclInfo(ctx context.Context, r soap.RoundTripper, req *types.VsanGetHclInfo) (*types.VsanGetHclInfoResponse, error) {
	var reqBody, resBody VsanGetHclInfoBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanHostRunVmdkLoadTestBody struct {
	Req    *types.VsanHostRunVmdkLoadTest         `xml:"urn:vsan VsanHostRunVmdkLoadTest,omitempty"`
	Res    *types.VsanHostRunVmdkLoadTestResponse `xml:"urn:vsan VsanHostRunVmdkLoadTestResponse,omitempty"`
	Fault_ *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanHostRunVmdkLoadTestBody) Fault() *soap.Fault { return b.Fault_ }

func VsanHostRunVmdkLoadTest(ctx context.Context, r soap.RoundTripper, req *types.VsanHostRunVmdkLoadTest) (*types.VsanHostRunVmdkLoadTestResponse, error) {
	var reqBody, resBody VsanHostRunVmdkLoadTestBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RestoreHostForClusterRebootWithNAMMBody struct {
	Req    *types.RestoreHostForClusterRebootWithNAMM         `xml:"urn:vsan RestoreHostForClusterRebootWithNAMM,omitempty"`
	Res    *types.RestoreHostForClusterRebootWithNAMMResponse `xml:"urn:vsan RestoreHostForClusterRebootWithNAMMResponse,omitempty"`
	Fault_ *soap.Fault                                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RestoreHostForClusterRebootWithNAMMBody) Fault() *soap.Fault { return b.Fault_ }

func RestoreHostForClusterRebootWithNAMM(ctx context.Context, r soap.RoundTripper, req *types.RestoreHostForClusterRebootWithNAMM) (*types.RestoreHostForClusterRebootWithNAMMResponse, error) {
	var reqBody, resBody RestoreHostForClusterRebootWithNAMMBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanHostQuerySmartStatsBody struct {
	Req    *types.VsanHostQuerySmartStats         `xml:"urn:vsan VsanHostQuerySmartStats,omitempty"`
	Res    *types.VsanHostQuerySmartStatsResponse `xml:"urn:vsan VsanHostQuerySmartStatsResponse,omitempty"`
	Fault_ *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanHostQuerySmartStatsBody) Fault() *soap.Fault { return b.Fault_ }

func VsanHostQuerySmartStats(ctx context.Context, r soap.RoundTripper, req *types.VsanHostQuerySmartStats) (*types.VsanHostQuerySmartStatsResponse, error) {
	var reqBody, resBody VsanHostQuerySmartStatsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanHostPrepareVmdkLoadTestBody struct {
	Req    *types.VsanHostPrepareVmdkLoadTest         `xml:"urn:vsan VsanHostPrepareVmdkLoadTest,omitempty"`
	Res    *types.VsanHostPrepareVmdkLoadTestResponse `xml:"urn:vsan VsanHostPrepareVmdkLoadTestResponse,omitempty"`
	Fault_ *soap.Fault                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanHostPrepareVmdkLoadTestBody) Fault() *soap.Fault { return b.Fault_ }

func VsanHostPrepareVmdkLoadTest(ctx context.Context, r soap.RoundTripper, req *types.VsanHostPrepareVmdkLoadTest) (*types.VsanHostPrepareVmdkLoadTestResponse, error) {
	var reqBody, resBody VsanHostPrepareVmdkLoadTestBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanHostQueryRunIperfServerBody struct {
	Req    *types.VsanHostQueryRunIperfServer         `xml:"urn:vsan VsanHostQueryRunIperfServer,omitempty"`
	Res    *types.VsanHostQueryRunIperfServerResponse `xml:"urn:vsan VsanHostQueryRunIperfServerResponse,omitempty"`
	Fault_ *soap.Fault                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanHostQueryRunIperfServerBody) Fault() *soap.Fault { return b.Fault_ }

func VsanHostQueryRunIperfServer(ctx context.Context, r soap.RoundTripper, req *types.VsanHostQueryRunIperfServer) (*types.VsanHostQueryRunIperfServerResponse, error) {
	var reqBody, resBody VsanHostQueryRunIperfServerBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanGetProactiveRebalanceInfoBody struct {
	Req    *types.VsanGetProactiveRebalanceInfo         `xml:"urn:vsan VsanGetProactiveRebalanceInfo,omitempty"`
	Res    *types.VsanGetProactiveRebalanceInfoResponse `xml:"urn:vsan VsanGetProactiveRebalanceInfoResponse,omitempty"`
	Fault_ *soap.Fault                                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanGetProactiveRebalanceInfoBody) Fault() *soap.Fault { return b.Fault_ }

func VsanGetProactiveRebalanceInfo(ctx context.Context, r soap.RoundTripper, req *types.VsanGetProactiveRebalanceInfo) (*types.VsanGetProactiveRebalanceInfoResponse, error) {
	var reqBody, resBody VsanGetProactiveRebalanceInfoBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanHostQueryPhysicalDiskHealthSummaryBody struct {
	Req    *types.VsanHostQueryPhysicalDiskHealthSummary         `xml:"urn:vsan VsanHostQueryPhysicalDiskHealthSummary,omitempty"`
	Res    *types.VsanHostQueryPhysicalDiskHealthSummaryResponse `xml:"urn:vsan VsanHostQueryPhysicalDiskHealthSummaryResponse,omitempty"`
	Fault_ *soap.Fault                                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanHostQueryPhysicalDiskHealthSummaryBody) Fault() *soap.Fault { return b.Fault_ }

func VsanHostQueryPhysicalDiskHealthSummary(ctx context.Context, r soap.RoundTripper, req *types.VsanHostQueryPhysicalDiskHealthSummary) (*types.VsanHostQueryPhysicalDiskHealthSummaryResponse, error) {
	var reqBody, resBody VsanHostQueryPhysicalDiskHealthSummaryBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanHostQueryVsanPcapBody struct {
	Req    *types.VsanHostQueryVsanPcap         `xml:"urn:vsan VsanHostQueryVsanPcap,omitempty"`
	Res    *types.VsanHostQueryVsanPcapResponse `xml:"urn:vsan VsanHostQueryVsanPcapResponse,omitempty"`
	Fault_ *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanHostQueryVsanPcapBody) Fault() *soap.Fault { return b.Fault_ }

func VsanHostQueryVsanPcap(ctx context.Context, r soap.RoundTripper, req *types.VsanHostQueryVsanPcap) (*types.VsanHostQueryVsanPcapResponse, error) {
	var reqBody, resBody VsanHostQueryVsanPcapBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanHostQueryHostInfoByUuidsBody struct {
	Req    *types.VsanHostQueryHostInfoByUuids         `xml:"urn:vsan VsanHostQueryHostInfoByUuids,omitempty"`
	Res    *types.VsanHostQueryHostInfoByUuidsResponse `xml:"urn:vsan VsanHostQueryHostInfoByUuidsResponse,omitempty"`
	Fault_ *soap.Fault                                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanHostQueryHostInfoByUuidsBody) Fault() *soap.Fault { return b.Fault_ }

func VsanHostQueryHostInfoByUuids(ctx context.Context, r soap.RoundTripper, req *types.VsanHostQueryHostInfoByUuids) (*types.VsanHostQueryHostInfoByUuidsResponse, error) {
	var reqBody, resBody VsanHostQueryHostInfoByUuidsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanCheckArchivalAccessibilityBody struct {
	Req    *types.VsanCheckArchivalAccessibility         `xml:"urn:vsan VsanCheckArchivalAccessibility,omitempty"`
	Res    *types.VsanCheckArchivalAccessibilityResponse `xml:"urn:vsan VsanCheckArchivalAccessibilityResponse,omitempty"`
	Fault_ *soap.Fault                                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanCheckArchivalAccessibilityBody) Fault() *soap.Fault { return b.Fault_ }

func VsanCheckArchivalAccessibility(ctx context.Context, r soap.RoundTripper, req *types.VsanCheckArchivalAccessibility) (*types.VsanCheckArchivalAccessibilityResponse, error) {
	var reqBody, resBody VsanCheckArchivalAccessibilityBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanQueryDiskRebalanceStatusBody struct {
	Req    *types.VsanQueryDiskRebalanceStatus         `xml:"urn:vsan VsanQueryDiskRebalanceStatus,omitempty"`
	Res    *types.VsanQueryDiskRebalanceStatusResponse `xml:"urn:vsan VsanQueryDiskRebalanceStatusResponse,omitempty"`
	Fault_ *soap.Fault                                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanQueryDiskRebalanceStatusBody) Fault() *soap.Fault { return b.Fault_ }

func VsanQueryDiskRebalanceStatus(ctx context.Context, r soap.RoundTripper, req *types.VsanQueryDiskRebalanceStatus) (*types.VsanQueryDiskRebalanceStatusResponse, error) {
	var reqBody, resBody VsanQueryDiskRebalanceStatusBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanHostCreateVmHealthTestBody struct {
	Req    *types.VsanHostCreateVmHealthTest         `xml:"urn:vsan VsanHostCreateVmHealthTest,omitempty"`
	Res    *types.VsanHostCreateVmHealthTestResponse `xml:"urn:vsan VsanHostCreateVmHealthTestResponse,omitempty"`
	Fault_ *soap.Fault                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanHostCreateVmHealthTestBody) Fault() *soap.Fault { return b.Fault_ }

func VsanHostCreateVmHealthTest(ctx context.Context, r soap.RoundTripper, req *types.VsanHostCreateVmHealthTest) (*types.VsanHostCreateVmHealthTestResponse, error) {
	var reqBody, resBody VsanHostCreateVmHealthTestBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanHostDpdLivenessBody struct {
	Req    *types.VsanHostDpdLiveness         `xml:"urn:vsan VsanHostDpdLiveness,omitempty"`
	Res    *types.VsanHostDpdLivenessResponse `xml:"urn:vsan VsanHostDpdLivenessResponse,omitempty"`
	Fault_ *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanHostDpdLivenessBody) Fault() *soap.Fault { return b.Fault_ }

func VsanHostDpdLiveness(ctx context.Context, r soap.RoundTripper, req *types.VsanHostDpdLiveness) (*types.VsanHostDpdLivenessResponse, error) {
	var reqBody, resBody VsanHostDpdLivenessBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanHostQueryCheckLimitsBody struct {
	Req    *types.VsanHostQueryCheckLimits         `xml:"urn:vsan VsanHostQueryCheckLimits,omitempty"`
	Res    *types.VsanHostQueryCheckLimitsResponse `xml:"urn:vsan VsanHostQueryCheckLimitsResponse,omitempty"`
	Fault_ *soap.Fault                             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanHostQueryCheckLimitsBody) Fault() *soap.Fault { return b.Fault_ }

func VsanHostQueryCheckLimits(ctx context.Context, r soap.RoundTripper, req *types.VsanHostQueryCheckLimits) (*types.VsanHostQueryCheckLimitsResponse, error) {
	var reqBody, resBody VsanHostQueryCheckLimitsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanCheckDatastoreUsageBody struct {
	Req    *types.VsanCheckDatastoreUsage         `xml:"urn:vsan VsanCheckDatastoreUsage,omitempty"`
	Res    *types.VsanCheckDatastoreUsageResponse `xml:"urn:vsan VsanCheckDatastoreUsageResponse,omitempty"`
	Fault_ *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanCheckDatastoreUsageBody) Fault() *soap.Fault { return b.Fault_ }

func VsanCheckDatastoreUsage(ctx context.Context, r soap.RoundTripper, req *types.VsanCheckDatastoreUsage) (*types.VsanCheckDatastoreUsageResponse, error) {
	var reqBody, resBody VsanCheckDatastoreUsageBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanPurgeHclFilesBody struct {
	Req    *types.VsanPurgeHclFiles         `xml:"urn:vsan VsanPurgeHclFiles,omitempty"`
	Res    *types.VsanPurgeHclFilesResponse `xml:"urn:vsan VsanPurgeHclFilesResponse,omitempty"`
	Fault_ *soap.Fault                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanPurgeHclFilesBody) Fault() *soap.Fault { return b.Fault_ }

func VsanPurgeHclFiles(ctx context.Context, r soap.RoundTripper, req *types.VsanPurgeHclFiles) (*types.VsanPurgeHclFilesResponse, error) {
	var reqBody, resBody VsanPurgeHclFilesBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanQueryVcClusterCreateVmHealthHistoryTestBody struct {
	Req    *types.VsanQueryVcClusterCreateVmHealthHistoryTest         `xml:"urn:vsan VsanQueryVcClusterCreateVmHealthHistoryTest,omitempty"`
	Res    *types.VsanQueryVcClusterCreateVmHealthHistoryTestResponse `xml:"urn:vsan VsanQueryVcClusterCreateVmHealthHistoryTestResponse,omitempty"`
	Fault_ *soap.Fault                                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanQueryVcClusterCreateVmHealthHistoryTestBody) Fault() *soap.Fault { return b.Fault_ }

func VsanQueryVcClusterCreateVmHealthHistoryTest(ctx context.Context, r soap.RoundTripper, req *types.VsanQueryVcClusterCreateVmHealthHistoryTest) (*types.VsanQueryVcClusterCreateVmHealthHistoryTestResponse, error) {
	var reqBody, resBody VsanQueryVcClusterCreateVmHealthHistoryTestBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanQueryVcClusterAdvCfgSyncBody struct {
	Req    *types.VsanQueryVcClusterAdvCfgSync         `xml:"urn:vsan VsanQueryVcClusterAdvCfgSync,omitempty"`
	Res    *types.VsanQueryVcClusterAdvCfgSyncResponse `xml:"urn:vsan VsanQueryVcClusterAdvCfgSyncResponse,omitempty"`
	Fault_ *soap.Fault                                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanQueryVcClusterAdvCfgSyncBody) Fault() *soap.Fault { return b.Fault_ }

func VsanQueryVcClusterAdvCfgSync(ctx context.Context, r soap.RoundTripper, req *types.VsanQueryVcClusterAdvCfgSync) (*types.VsanQueryVcClusterAdvCfgSyncResponse, error) {
	var reqBody, resBody VsanQueryVcClusterAdvCfgSyncBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanQueryVcClusterObjExtAttrsBody struct {
	Req    *types.VsanQueryVcClusterObjExtAttrs         `xml:"urn:vsan VsanQueryVcClusterObjExtAttrs,omitempty"`
	Res    *types.VsanQueryVcClusterObjExtAttrsResponse `xml:"urn:vsan VsanQueryVcClusterObjExtAttrsResponse,omitempty"`
	Fault_ *soap.Fault                                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanQueryVcClusterObjExtAttrsBody) Fault() *soap.Fault { return b.Fault_ }

func VsanQueryVcClusterObjExtAttrs(ctx context.Context, r soap.RoundTripper, req *types.VsanQueryVcClusterObjExtAttrs) (*types.VsanQueryVcClusterObjExtAttrsResponse, error) {
	var reqBody, resBody VsanQueryVcClusterObjExtAttrsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanVcClusterQueryObjectHealthSummaryBody struct {
	Req    *types.VsanVcClusterQueryObjectHealthSummary         `xml:"urn:vsan VsanVcClusterQueryObjectHealthSummary,omitempty"`
	Res    *types.VsanVcClusterQueryObjectHealthSummaryResponse `xml:"urn:vsan VsanVcClusterQueryObjectHealthSummaryResponse,omitempty"`
	Fault_ *soap.Fault                                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanVcClusterQueryObjectHealthSummaryBody) Fault() *soap.Fault { return b.Fault_ }

func VsanVcClusterQueryObjectHealthSummary(ctx context.Context, r soap.RoundTripper, req *types.VsanVcClusterQueryObjectHealthSummary) (*types.VsanVcClusterQueryObjectHealthSummaryResponse, error) {
	var reqBody, resBody VsanVcClusterQueryObjectHealthSummaryBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanHealthSetLogLevelBody struct {
	Req    *types.VsanHealthSetLogLevel         `xml:"urn:vsan VsanHealthSetLogLevel,omitempty"`
	Res    *types.VsanHealthSetLogLevelResponse `xml:"urn:vsan VsanHealthSetLogLevelResponse,omitempty"`
	Fault_ *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanHealthSetLogLevelBody) Fault() *soap.Fault { return b.Fault_ }

func VsanHealthSetLogLevel(ctx context.Context, r soap.RoundTripper, req *types.VsanHealthSetLogLevel) (*types.VsanHealthSetLogLevelResponse, error) {
	var reqBody, resBody VsanHealthSetLogLevelBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanHealthTestVsanClusterTelemetryProxyBody struct {
	Req    *types.VsanHealthTestVsanClusterTelemetryProxy         `xml:"urn:vsan VsanHealthTestVsanClusterTelemetryProxy,omitempty"`
	Res    *types.VsanHealthTestVsanClusterTelemetryProxyResponse `xml:"urn:vsan VsanHealthTestVsanClusterTelemetryProxyResponse,omitempty"`
	Fault_ *soap.Fault                                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanHealthTestVsanClusterTelemetryProxyBody) Fault() *soap.Fault { return b.Fault_ }

func VsanHealthTestVsanClusterTelemetryProxy(ctx context.Context, r soap.RoundTripper, req *types.VsanHealthTestVsanClusterTelemetryProxy) (*types.VsanHealthTestVsanClusterTelemetryProxyResponse, error) {
	var reqBody, resBody VsanHealthTestVsanClusterTelemetryProxyBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanQueryVerifyVcClusterNetworkSettingsBody struct {
	Req    *types.VsanQueryVerifyVcClusterNetworkSettings         `xml:"urn:vsan VsanQueryVerifyVcClusterNetworkSettings,omitempty"`
	Res    *types.VsanQueryVerifyVcClusterNetworkSettingsResponse `xml:"urn:vsan VsanQueryVerifyVcClusterNetworkSettingsResponse,omitempty"`
	Fault_ *soap.Fault                                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanQueryVerifyVcClusterNetworkSettingsBody) Fault() *soap.Fault { return b.Fault_ }

func VsanQueryVerifyVcClusterNetworkSettings(ctx context.Context, r soap.RoundTripper, req *types.VsanQueryVerifyVcClusterNetworkSettings) (*types.VsanQueryVerifyVcClusterNetworkSettingsResponse, error) {
	var reqBody, resBody VsanQueryVerifyVcClusterNetworkSettingsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanVcUploadHclDbBody struct {
	Req    *types.VsanVcUploadHclDb         `xml:"urn:vsan VsanVcUploadHclDb,omitempty"`
	Res    *types.VsanVcUploadHclDbResponse `xml:"urn:vsan VsanVcUploadHclDbResponse,omitempty"`
	Fault_ *soap.Fault                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanVcUploadHclDbBody) Fault() *soap.Fault { return b.Fault_ }

func VsanVcUploadHclDb(ctx context.Context, r soap.RoundTripper, req *types.VsanVcUploadHclDb) (*types.VsanVcUploadHclDbResponse, error) {
	var reqBody, resBody VsanVcUploadHclDbBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanQueryVcClusterSmartStatsSummaryBody struct {
	Req    *types.VsanQueryVcClusterSmartStatsSummary         `xml:"urn:vsan VsanQueryVcClusterSmartStatsSummary,omitempty"`
	Res    *types.VsanQueryVcClusterSmartStatsSummaryResponse `xml:"urn:vsan VsanQueryVcClusterSmartStatsSummaryResponse,omitempty"`
	Fault_ *soap.Fault                                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanQueryVcClusterSmartStatsSummaryBody) Fault() *soap.Fault { return b.Fault_ }

func VsanQueryVcClusterSmartStatsSummary(ctx context.Context, r soap.RoundTripper, req *types.VsanQueryVcClusterSmartStatsSummary) (*types.VsanQueryVcClusterSmartStatsSummaryResponse, error) {
	var reqBody, resBody VsanQueryVcClusterSmartStatsSummaryBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanVcUpdateHclDbFromWebBody struct {
	Req    *types.VsanVcUpdateHclDbFromWeb         `xml:"urn:vsan VsanVcUpdateHclDbFromWeb,omitempty"`
	Res    *types.VsanVcUpdateHclDbFromWebResponse `xml:"urn:vsan VsanVcUpdateHclDbFromWebResponse,omitempty"`
	Fault_ *soap.Fault                             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanVcUpdateHclDbFromWebBody) Fault() *soap.Fault { return b.Fault_ }

func VsanVcUpdateHclDbFromWeb(ctx context.Context, r soap.RoundTripper, req *types.VsanVcUpdateHclDbFromWeb) (*types.VsanVcUpdateHclDbFromWebResponse, error) {
	var reqBody, resBody VsanVcUpdateHclDbFromWebBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanHealthGetVsanClusterSilentChecksBody struct {
	Req    *types.VsanHealthGetVsanClusterSilentChecks         `xml:"urn:vsan VsanHealthGetVsanClusterSilentChecks,omitempty"`
	Res    *types.VsanHealthGetVsanClusterSilentChecksResponse `xml:"urn:vsan VsanHealthGetVsanClusterSilentChecksResponse,omitempty"`
	Fault_ *soap.Fault                                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanHealthGetVsanClusterSilentChecksBody) Fault() *soap.Fault { return b.Fault_ }

func VsanHealthGetVsanClusterSilentChecks(ctx context.Context, r soap.RoundTripper, req *types.VsanHealthGetVsanClusterSilentChecks) (*types.VsanHealthGetVsanClusterSilentChecksResponse, error) {
	var reqBody, resBody VsanHealthGetVsanClusterSilentChecksBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanClusterQueryFileServiceHealthSummaryBody struct {
	Req    *types.VsanClusterQueryFileServiceHealthSummary         `xml:"urn:vsan VsanClusterQueryFileServiceHealthSummary,omitempty"`
	Res    *types.VsanClusterQueryFileServiceHealthSummaryResponse `xml:"urn:vsan VsanClusterQueryFileServiceHealthSummaryResponse,omitempty"`
	Fault_ *soap.Fault                                             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanClusterQueryFileServiceHealthSummaryBody) Fault() *soap.Fault { return b.Fault_ }

func VsanClusterQueryFileServiceHealthSummary(ctx context.Context, r soap.RoundTripper, req *types.VsanClusterQueryFileServiceHealthSummary) (*types.VsanClusterQueryFileServiceHealthSummaryResponse, error) {
	var reqBody, resBody VsanClusterQueryFileServiceHealthSummaryBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type PrepareClusterRebootWithNAMMBody struct {
	Req    *types.PrepareClusterRebootWithNAMM         `xml:"urn:vsan PrepareClusterRebootWithNAMM,omitempty"`
	Res    *types.PrepareClusterRebootWithNAMMResponse `xml:"urn:vsan PrepareClusterRebootWithNAMMResponse,omitempty"`
	Fault_ *soap.Fault                                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *PrepareClusterRebootWithNAMMBody) Fault() *soap.Fault { return b.Fault_ }

func PrepareClusterRebootWithNAMM(ctx context.Context, r soap.RoundTripper, req *types.PrepareClusterRebootWithNAMM) (*types.PrepareClusterRebootWithNAMMResponse, error) {
	var reqBody, resBody PrepareClusterRebootWithNAMMBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanVcClusterClomdLivenessBody struct {
	Req    *types.VsanVcClusterClomdLiveness         `xml:"urn:vsan VsanVcClusterClomdLiveness,omitempty"`
	Res    *types.VsanVcClusterClomdLivenessResponse `xml:"urn:vsan VsanVcClusterClomdLivenessResponse,omitempty"`
	Fault_ *soap.Fault                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanVcClusterClomdLivenessBody) Fault() *soap.Fault { return b.Fault_ }

func VsanVcClusterClomdLiveness(ctx context.Context, r soap.RoundTripper, req *types.VsanVcClusterClomdLiveness) (*types.VsanVcClusterClomdLivenessResponse, error) {
	var reqBody, resBody VsanVcClusterClomdLivenessBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanClusterHealthExtensionManagmentPreCheckBody struct {
	Req    *types.VsanClusterHealthExtensionManagmentPreCheck         `xml:"urn:vsan VsanClusterHealthExtensionManagmentPreCheck,omitempty"`
	Res    *types.VsanClusterHealthExtensionManagmentPreCheckResponse `xml:"urn:vsan VsanClusterHealthExtensionManagmentPreCheckResponse,omitempty"`
	Fault_ *soap.Fault                                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanClusterHealthExtensionManagmentPreCheckBody) Fault() *soap.Fault { return b.Fault_ }

func VsanClusterHealthExtensionManagmentPreCheck(ctx context.Context, r soap.RoundTripper, req *types.VsanClusterHealthExtensionManagmentPreCheck) (*types.VsanClusterHealthExtensionManagmentPreCheckResponse, error) {
	var reqBody, resBody VsanClusterHealthExtensionManagmentPreCheckBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanHealthRepairClusterObjectsImmediateBody struct {
	Req    *types.VsanHealthRepairClusterObjectsImmediate         `xml:"urn:vsan VsanHealthRepairClusterObjectsImmediate,omitempty"`
	Res    *types.VsanHealthRepairClusterObjectsImmediateResponse `xml:"urn:vsan VsanHealthRepairClusterObjectsImmediateResponse,omitempty"`
	Fault_ *soap.Fault                                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanHealthRepairClusterObjectsImmediateBody) Fault() *soap.Fault { return b.Fault_ }

func VsanHealthRepairClusterObjectsImmediate(ctx context.Context, r soap.RoundTripper, req *types.VsanHealthRepairClusterObjectsImmediate) (*types.VsanHealthRepairClusterObjectsImmediateResponse, error) {
	var reqBody, resBody VsanHealthRepairClusterObjectsImmediateBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanQueryVcClusterNetworkPerfTestBody struct {
	Req    *types.VsanQueryVcClusterNetworkPerfTest         `xml:"urn:vsan VsanQueryVcClusterNetworkPerfTest,omitempty"`
	Res    *types.VsanQueryVcClusterNetworkPerfTestResponse `xml:"urn:vsan VsanQueryVcClusterNetworkPerfTestResponse,omitempty"`
	Fault_ *soap.Fault                                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanQueryVcClusterNetworkPerfTestBody) Fault() *soap.Fault { return b.Fault_ }

func VsanQueryVcClusterNetworkPerfTest(ctx context.Context, r soap.RoundTripper, req *types.VsanQueryVcClusterNetworkPerfTest) (*types.VsanQueryVcClusterNetworkPerfTestResponse, error) {
	var reqBody, resBody VsanQueryVcClusterNetworkPerfTestBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanQueryVcClusterVmdkLoadHistoryTestBody struct {
	Req    *types.VsanQueryVcClusterVmdkLoadHistoryTest         `xml:"urn:vsan VsanQueryVcClusterVmdkLoadHistoryTest,omitempty"`
	Res    *types.VsanQueryVcClusterVmdkLoadHistoryTestResponse `xml:"urn:vsan VsanQueryVcClusterVmdkLoadHistoryTestResponse,omitempty"`
	Fault_ *soap.Fault                                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanQueryVcClusterVmdkLoadHistoryTestBody) Fault() *soap.Fault { return b.Fault_ }

func VsanQueryVcClusterVmdkLoadHistoryTest(ctx context.Context, r soap.RoundTripper, req *types.VsanQueryVcClusterVmdkLoadHistoryTest) (*types.VsanQueryVcClusterVmdkLoadHistoryTestResponse, error) {
	var reqBody, resBody VsanQueryVcClusterVmdkLoadHistoryTestBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanHealthIsRebalanceRunningBody struct {
	Req    *types.VsanHealthIsRebalanceRunning         `xml:"urn:vsan VsanHealthIsRebalanceRunning,omitempty"`
	Res    *types.VsanHealthIsRebalanceRunningResponse `xml:"urn:vsan VsanHealthIsRebalanceRunningResponse,omitempty"`
	Fault_ *soap.Fault                                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanHealthIsRebalanceRunningBody) Fault() *soap.Fault { return b.Fault_ }

func VsanHealthIsRebalanceRunning(ctx context.Context, r soap.RoundTripper, req *types.VsanHealthIsRebalanceRunning) (*types.VsanHealthIsRebalanceRunningResponse, error) {
	var reqBody, resBody VsanHealthIsRebalanceRunningBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanQueryVcClusterCreateVmHealthTestBody struct {
	Req    *types.VsanQueryVcClusterCreateVmHealthTest         `xml:"urn:vsan VsanQueryVcClusterCreateVmHealthTest,omitempty"`
	Res    *types.VsanQueryVcClusterCreateVmHealthTestResponse `xml:"urn:vsan VsanQueryVcClusterCreateVmHealthTestResponse,omitempty"`
	Fault_ *soap.Fault                                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanQueryVcClusterCreateVmHealthTestBody) Fault() *soap.Fault { return b.Fault_ }

func VsanQueryVcClusterCreateVmHealthTest(ctx context.Context, r soap.RoundTripper, req *types.VsanQueryVcClusterCreateVmHealthTest) (*types.VsanQueryVcClusterCreateVmHealthTestResponse, error) {
	var reqBody, resBody VsanQueryVcClusterCreateVmHealthTestBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanHealthQueryVsanProxyConfigBody struct {
	Req    *types.VsanHealthQueryVsanProxyConfig         `xml:"urn:vsan VsanHealthQueryVsanProxyConfig,omitempty"`
	Res    *types.VsanHealthQueryVsanProxyConfigResponse `xml:"urn:vsan VsanHealthQueryVsanProxyConfigResponse,omitempty"`
	Fault_ *soap.Fault                                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanHealthQueryVsanProxyConfigBody) Fault() *soap.Fault { return b.Fault_ }

func VsanHealthQueryVsanProxyConfig(ctx context.Context, r soap.RoundTripper, req *types.VsanHealthQueryVsanProxyConfig) (*types.VsanHealthQueryVsanProxyConfigResponse, error) {
	var reqBody, resBody VsanHealthQueryVsanProxyConfigBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanHealthQueryVsanClusterHealthCheckIntervalBody struct {
	Req    *types.VsanHealthQueryVsanClusterHealthCheckInterval         `xml:"urn:vsan VsanHealthQueryVsanClusterHealthCheckInterval,omitempty"`
	Res    *types.VsanHealthQueryVsanClusterHealthCheckIntervalResponse `xml:"urn:vsan VsanHealthQueryVsanClusterHealthCheckIntervalResponse,omitempty"`
	Fault_ *soap.Fault                                                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanHealthQueryVsanClusterHealthCheckIntervalBody) Fault() *soap.Fault { return b.Fault_ }

func VsanHealthQueryVsanClusterHealthCheckInterval(ctx context.Context, r soap.RoundTripper, req *types.VsanHealthQueryVsanClusterHealthCheckInterval) (*types.VsanHealthQueryVsanClusterHealthCheckIntervalResponse, error) {
	var reqBody, resBody VsanHealthQueryVsanClusterHealthCheckIntervalBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanQueryVcClusterEncryptionHealthSummaryBody struct {
	Req    *types.VsanQueryVcClusterEncryptionHealthSummary         `xml:"urn:vsan VsanQueryVcClusterEncryptionHealthSummary,omitempty"`
	Res    *types.VsanQueryVcClusterEncryptionHealthSummaryResponse `xml:"urn:vsan VsanQueryVcClusterEncryptionHealthSummaryResponse,omitempty"`
	Fault_ *soap.Fault                                              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanQueryVcClusterEncryptionHealthSummaryBody) Fault() *soap.Fault { return b.Fault_ }

func VsanQueryVcClusterEncryptionHealthSummary(ctx context.Context, r soap.RoundTripper, req *types.VsanQueryVcClusterEncryptionHealthSummary) (*types.VsanQueryVcClusterEncryptionHealthSummaryResponse, error) {
	var reqBody, resBody VsanQueryVcClusterEncryptionHealthSummaryBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanQueryAllSupportedHealthChecksBody struct {
	Req    *types.VsanQueryAllSupportedHealthChecks         `xml:"urn:vsan VsanQueryAllSupportedHealthChecks,omitempty"`
	Res    *types.VsanQueryAllSupportedHealthChecksResponse `xml:"urn:vsan VsanQueryAllSupportedHealthChecksResponse,omitempty"`
	Fault_ *soap.Fault                                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanQueryAllSupportedHealthChecksBody) Fault() *soap.Fault { return b.Fault_ }

func VsanQueryAllSupportedHealthChecks(ctx context.Context, r soap.RoundTripper, req *types.VsanQueryAllSupportedHealthChecks) (*types.VsanQueryAllSupportedHealthChecksResponse, error) {
	var reqBody, resBody VsanQueryAllSupportedHealthChecksBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanVcClusterGetHclInfoBody struct {
	Req    *types.VsanVcClusterGetHclInfo         `xml:"urn:vsan VsanVcClusterGetHclInfo,omitempty"`
	Res    *types.VsanVcClusterGetHclInfoResponse `xml:"urn:vsan VsanVcClusterGetHclInfoResponse,omitempty"`
	Fault_ *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanVcClusterGetHclInfoBody) Fault() *soap.Fault { return b.Fault_ }

func VsanVcClusterGetHclInfo(ctx context.Context, r soap.RoundTripper, req *types.VsanVcClusterGetHclInfo) (*types.VsanVcClusterGetHclInfoResponse, error) {
	var reqBody, resBody VsanVcClusterGetHclInfoBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanHealthPrepareClusterBody struct {
	Req    *types.VsanHealthPrepareCluster         `xml:"urn:vsan VsanHealthPrepareCluster,omitempty"`
	Res    *types.VsanHealthPrepareClusterResponse `xml:"urn:vsan VsanHealthPrepareClusterResponse,omitempty"`
	Fault_ *soap.Fault                             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanHealthPrepareClusterBody) Fault() *soap.Fault { return b.Fault_ }

func VsanHealthPrepareCluster(ctx context.Context, r soap.RoundTripper, req *types.VsanHealthPrepareCluster) (*types.VsanHealthPrepareClusterResponse, error) {
	var reqBody, resBody VsanHealthPrepareClusterBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanQueryAttachToSrHistoryBody struct {
	Req    *types.VsanQueryAttachToSrHistory         `xml:"urn:vsan VsanQueryAttachToSrHistory,omitempty"`
	Res    *types.VsanQueryAttachToSrHistoryResponse `xml:"urn:vsan VsanQueryAttachToSrHistoryResponse,omitempty"`
	Fault_ *soap.Fault                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanQueryAttachToSrHistoryBody) Fault() *soap.Fault { return b.Fault_ }

func VsanQueryAttachToSrHistory(ctx context.Context, r soap.RoundTripper, req *types.VsanQueryAttachToSrHistory) (*types.VsanQueryAttachToSrHistoryResponse, error) {
	var reqBody, resBody VsanQueryAttachToSrHistoryBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanVcClusterArchivalAccessibilityBody struct {
	Req    *types.VsanVcClusterArchivalAccessibility         `xml:"urn:vsan VsanVcClusterArchivalAccessibility,omitempty"`
	Res    *types.VsanVcClusterArchivalAccessibilityResponse `xml:"urn:vsan VsanVcClusterArchivalAccessibilityResponse,omitempty"`
	Fault_ *soap.Fault                                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanVcClusterArchivalAccessibilityBody) Fault() *soap.Fault { return b.Fault_ }

func VsanVcClusterArchivalAccessibility(ctx context.Context, r soap.RoundTripper, req *types.VsanVcClusterArchivalAccessibility) (*types.VsanVcClusterArchivalAccessibilityResponse, error) {
	var reqBody, resBody VsanVcClusterArchivalAccessibilityBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanGetHclConstraintsBody struct {
	Req    *types.VsanGetHclConstraints         `xml:"urn:vsan VsanGetHclConstraints,omitempty"`
	Res    *types.VsanGetHclConstraintsResponse `xml:"urn:vsan VsanGetHclConstraintsResponse,omitempty"`
	Fault_ *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanGetHclConstraintsBody) Fault() *soap.Fault { return b.Fault_ }

func VsanGetHclConstraints(ctx context.Context, r soap.RoundTripper, req *types.VsanGetHclConstraints) (*types.VsanGetHclConstraintsResponse, error) {
	var reqBody, resBody VsanGetHclConstraintsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanRebalanceClusterBody struct {
	Req    *types.VsanRebalanceCluster         `xml:"urn:vsan VsanRebalanceCluster,omitempty"`
	Res    *types.VsanRebalanceClusterResponse `xml:"urn:vsan VsanRebalanceClusterResponse,omitempty"`
	Fault_ *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanRebalanceClusterBody) Fault() *soap.Fault { return b.Fault_ }

func VsanRebalanceCluster(ctx context.Context, r soap.RoundTripper, req *types.VsanRebalanceCluster) (*types.VsanRebalanceClusterResponse, error) {
	var reqBody, resBody VsanRebalanceClusterBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanVcClusterRunVmdkLoadTestBody struct {
	Req    *types.VsanVcClusterRunVmdkLoadTest         `xml:"urn:vsan VsanVcClusterRunVmdkLoadTest,omitempty"`
	Res    *types.VsanVcClusterRunVmdkLoadTestResponse `xml:"urn:vsan VsanVcClusterRunVmdkLoadTestResponse,omitempty"`
	Fault_ *soap.Fault                                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanVcClusterRunVmdkLoadTestBody) Fault() *soap.Fault { return b.Fault_ }

func VsanVcClusterRunVmdkLoadTest(ctx context.Context, r soap.RoundTripper, req *types.VsanVcClusterRunVmdkLoadTest) (*types.VsanVcClusterRunVmdkLoadTestResponse, error) {
	var reqBody, resBody VsanVcClusterRunVmdkLoadTestBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanHealthSendVsanTelemetryBody struct {
	Req    *types.VsanHealthSendVsanTelemetry         `xml:"urn:vsan VsanHealthSendVsanTelemetry,omitempty"`
	Res    *types.VsanHealthSendVsanTelemetryResponse `xml:"urn:vsan VsanHealthSendVsanTelemetryResponse,omitempty"`
	Fault_ *soap.Fault                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanHealthSendVsanTelemetryBody) Fault() *soap.Fault { return b.Fault_ }

func VsanHealthSendVsanTelemetry(ctx context.Context, r soap.RoundTripper, req *types.VsanHealthSendVsanTelemetry) (*types.VsanHealthSendVsanTelemetryResponse, error) {
	var reqBody, resBody VsanHealthSendVsanTelemetryBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanRemediateDataProtectionConfigInClusterBody struct {
	Req    *types.VsanRemediateDataProtectionConfigInCluster         `xml:"urn:vsan VsanRemediateDataProtectionConfigInCluster,omitempty"`
	Res    *types.VsanRemediateDataProtectionConfigInClusterResponse `xml:"urn:vsan VsanRemediateDataProtectionConfigInClusterResponse,omitempty"`
	Fault_ *soap.Fault                                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanRemediateDataProtectionConfigInClusterBody) Fault() *soap.Fault { return b.Fault_ }

func VsanRemediateDataProtectionConfigInCluster(ctx context.Context, r soap.RoundTripper, req *types.VsanRemediateDataProtectionConfigInCluster) (*types.VsanRemediateDataProtectionConfigInClusterResponse, error) {
	var reqBody, resBody VsanRemediateDataProtectionConfigInClusterBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanQueryVcClusterNetworkPerfHistoryTestBody struct {
	Req    *types.VsanQueryVcClusterNetworkPerfHistoryTest         `xml:"urn:vsan VsanQueryVcClusterNetworkPerfHistoryTest,omitempty"`
	Res    *types.VsanQueryVcClusterNetworkPerfHistoryTestResponse `xml:"urn:vsan VsanQueryVcClusterNetworkPerfHistoryTestResponse,omitempty"`
	Fault_ *soap.Fault                                             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanQueryVcClusterNetworkPerfHistoryTestBody) Fault() *soap.Fault { return b.Fault_ }

func VsanQueryVcClusterNetworkPerfHistoryTest(ctx context.Context, r soap.RoundTripper, req *types.VsanQueryVcClusterNetworkPerfHistoryTest) (*types.VsanQueryVcClusterNetworkPerfHistoryTestResponse, error) {
	var reqBody, resBody VsanQueryVcClusterNetworkPerfHistoryTestBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanQueryVcClusterHistoryHealthSummaryBody struct {
	Req    *types.VsanQueryVcClusterHistoryHealthSummary         `xml:"urn:vsan VsanQueryVcClusterHistoryHealthSummary,omitempty"`
	Res    *types.VsanQueryVcClusterHistoryHealthSummaryResponse `xml:"urn:vsan VsanQueryVcClusterHistoryHealthSummaryResponse,omitempty"`
	Fault_ *soap.Fault                                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanQueryVcClusterHistoryHealthSummaryBody) Fault() *soap.Fault { return b.Fault_ }

func VsanQueryVcClusterHistoryHealthSummary(ctx context.Context, r soap.RoundTripper, req *types.VsanQueryVcClusterHistoryHealthSummary) (*types.VsanQueryVcClusterHistoryHealthSummaryResponse, error) {
	var reqBody, resBody VsanQueryVcClusterHistoryHealthSummaryBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanQueryVcClusterHealthSummaryBody struct {
	Req    *types.VsanQueryVcClusterHealthSummary         `xml:"urn:vsan VsanQueryVcClusterHealthSummary,omitempty"`
	Res    *types.VsanQueryVcClusterHealthSummaryResponse `xml:"urn:vsan VsanQueryVcClusterHealthSummaryResponse,omitempty"`
	Fault_ *soap.Fault                                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanQueryVcClusterHealthSummaryBody) Fault() *soap.Fault { return b.Fault_ }

func VsanQueryVcClusterHealthSummary(ctx context.Context, r soap.RoundTripper, req *types.VsanQueryVcClusterHealthSummary) (*types.VsanQueryVcClusterHealthSummaryResponse, error) {
	var reqBody, resBody VsanQueryVcClusterHealthSummaryBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanQueryVcClusterHealthSummaryTaskBody struct {
	Req    *types.VsanQueryVcClusterHealthSummaryTask         `xml:"urn:vsan VsanQueryVcClusterHealthSummaryTask,omitempty"`
	Res    *types.VsanQueryVcClusterHealthSummaryTaskResponse `xml:"urn:vsan VsanQueryVcClusterHealthSummaryTaskResponse,omitempty"`
	Fault_ *soap.Fault                                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanQueryVcClusterHealthSummaryTaskBody) Fault() *soap.Fault { return b.Fault_ }

func VsanQueryVcClusterHealthSummaryTask(ctx context.Context, r soap.RoundTripper, req *types.VsanQueryVcClusterHealthSummaryTask) (*types.VsanQueryVcClusterHealthSummaryTaskResponse, error) {
	var reqBody, resBody VsanQueryVcClusterHealthSummaryTaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanHealthGetClusterStatusBody struct {
	Req    *types.VsanHealthGetClusterStatus         `xml:"urn:vsan VsanHealthGetClusterStatus,omitempty"`
	Res    *types.VsanHealthGetClusterStatusResponse `xml:"urn:vsan VsanHealthGetClusterStatusResponse,omitempty"`
	Fault_ *soap.Fault                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanHealthGetClusterStatusBody) Fault() *soap.Fault { return b.Fault_ }

func VsanHealthGetClusterStatus(ctx context.Context, r soap.RoundTripper, req *types.VsanHealthGetClusterStatus) (*types.VsanHealthGetClusterStatusResponse, error) {
	var reqBody, resBody VsanHealthGetClusterStatusBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanStopRebalanceClusterBody struct {
	Req    *types.VsanStopRebalanceCluster         `xml:"urn:vsan VsanStopRebalanceCluster,omitempty"`
	Res    *types.VsanStopRebalanceClusterResponse `xml:"urn:vsan VsanStopRebalanceClusterResponse,omitempty"`
	Fault_ *soap.Fault                             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanStopRebalanceClusterBody) Fault() *soap.Fault { return b.Fault_ }

func VsanStopRebalanceCluster(ctx context.Context, r soap.RoundTripper, req *types.VsanStopRebalanceCluster) (*types.VsanStopRebalanceClusterResponse, error) {
	var reqBody, resBody VsanStopRebalanceClusterBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanHealthQueryVsanClusterHealthConfigBody struct {
	Req    *types.VsanHealthQueryVsanClusterHealthConfig         `xml:"urn:vsan VsanHealthQueryVsanClusterHealthConfig,omitempty"`
	Res    *types.VsanHealthQueryVsanClusterHealthConfigResponse `xml:"urn:vsan VsanHealthQueryVsanClusterHealthConfigResponse,omitempty"`
	Fault_ *soap.Fault                                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanHealthQueryVsanClusterHealthConfigBody) Fault() *soap.Fault { return b.Fault_ }

func VsanHealthQueryVsanClusterHealthConfig(ctx context.Context, r soap.RoundTripper, req *types.VsanHealthQueryVsanClusterHealthConfig) (*types.VsanHealthQueryVsanClusterHealthConfigResponse, error) {
	var reqBody, resBody VsanHealthQueryVsanClusterHealthConfigBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanAttachVsanSupportBundleToSrBody struct {
	Req    *types.VsanAttachVsanSupportBundleToSr         `xml:"urn:vsan VsanAttachVsanSupportBundleToSr,omitempty"`
	Res    *types.VsanAttachVsanSupportBundleToSrResponse `xml:"urn:vsan VsanAttachVsanSupportBundleToSrResponse,omitempty"`
	Fault_ *soap.Fault                                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanAttachVsanSupportBundleToSrBody) Fault() *soap.Fault { return b.Fault_ }

func VsanAttachVsanSupportBundleToSr(ctx context.Context, r soap.RoundTripper, req *types.VsanAttachVsanSupportBundleToSr) (*types.VsanAttachVsanSupportBundleToSrResponse, error) {
	var reqBody, resBody VsanAttachVsanSupportBundleToSrBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanDownloadHclFile_TaskBody struct {
	Req    *types.VsanDownloadHclFile_Task         `xml:"urn:vsan VsanDownloadHclFile_Task,omitempty"`
	Res    *types.VsanDownloadHclFile_TaskResponse `xml:"urn:vsan VsanDownloadHclFile_TaskResponse,omitempty"`
	Fault_ *soap.Fault                             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanDownloadHclFile_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func VsanDownloadHclFile_Task(ctx context.Context, r soap.RoundTripper, req *types.VsanDownloadHclFile_Task) (*types.VsanDownloadHclFile_TaskResponse, error) {
	var reqBody, resBody VsanDownloadHclFile_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanVcQueryClusterCaptureVsanPcapBody struct {
	Req    *types.VsanVcQueryClusterCaptureVsanPcap         `xml:"urn:vsan VsanVcQueryClusterCaptureVsanPcap,omitempty"`
	Res    *types.VsanVcQueryClusterCaptureVsanPcapResponse `xml:"urn:vsan VsanVcQueryClusterCaptureVsanPcapResponse,omitempty"`
	Fault_ *soap.Fault                                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanVcQueryClusterCaptureVsanPcapBody) Fault() *soap.Fault { return b.Fault_ }

func VsanVcQueryClusterCaptureVsanPcap(ctx context.Context, r soap.RoundTripper, req *types.VsanVcQueryClusterCaptureVsanPcap) (*types.VsanVcQueryClusterCaptureVsanPcapResponse, error) {
	var reqBody, resBody VsanVcQueryClusterCaptureVsanPcapBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanQueryVcClusterVmdkWorkloadTypesBody struct {
	Req    *types.VsanQueryVcClusterVmdkWorkloadTypes         `xml:"urn:vsan VsanQueryVcClusterVmdkWorkloadTypes,omitempty"`
	Res    *types.VsanQueryVcClusterVmdkWorkloadTypesResponse `xml:"urn:vsan VsanQueryVcClusterVmdkWorkloadTypesResponse,omitempty"`
	Fault_ *soap.Fault                                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanQueryVcClusterVmdkWorkloadTypesBody) Fault() *soap.Fault { return b.Fault_ }

func VsanQueryVcClusterVmdkWorkloadTypes(ctx context.Context, r soap.RoundTripper, req *types.VsanQueryVcClusterVmdkWorkloadTypes) (*types.VsanQueryVcClusterVmdkWorkloadTypesResponse, error) {
	var reqBody, resBody VsanQueryVcClusterVmdkWorkloadTypesBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanQueryVcClusterPhysicalDiskHealthSummaryBody struct {
	Req    *types.VsanQueryVcClusterPhysicalDiskHealthSummary         `xml:"urn:vsan VsanQueryVcClusterPhysicalDiskHealthSummary,omitempty"`
	Res    *types.VsanQueryVcClusterPhysicalDiskHealthSummaryResponse `xml:"urn:vsan VsanQueryVcClusterPhysicalDiskHealthSummaryResponse,omitempty"`
	Fault_ *soap.Fault                                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanQueryVcClusterPhysicalDiskHealthSummaryBody) Fault() *soap.Fault { return b.Fault_ }

func VsanQueryVcClusterPhysicalDiskHealthSummary(ctx context.Context, r soap.RoundTripper, req *types.VsanQueryVcClusterPhysicalDiskHealthSummary) (*types.VsanQueryVcClusterPhysicalDiskHealthSummaryResponse, error) {
	var reqBody, resBody VsanQueryVcClusterPhysicalDiskHealthSummaryBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanHealthSetVsanClusterSilentChecksBody struct {
	Req    *types.VsanHealthSetVsanClusterSilentChecks         `xml:"urn:vsan VsanHealthSetVsanClusterSilentChecks,omitempty"`
	Res    *types.VsanHealthSetVsanClusterSilentChecksResponse `xml:"urn:vsan VsanHealthSetVsanClusterSilentChecksResponse,omitempty"`
	Fault_ *soap.Fault                                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanHealthSetVsanClusterSilentChecksBody) Fault() *soap.Fault { return b.Fault_ }

func VsanHealthSetVsanClusterSilentChecks(ctx context.Context, r soap.RoundTripper, req *types.VsanHealthSetVsanClusterSilentChecks) (*types.VsanHealthSetVsanClusterSilentChecksResponse, error) {
	var reqBody, resBody VsanHealthSetVsanClusterSilentChecksBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanQueryVcClusterCheckLimitsBody struct {
	Req    *types.VsanQueryVcClusterCheckLimits         `xml:"urn:vsan VsanQueryVcClusterCheckLimits,omitempty"`
	Res    *types.VsanQueryVcClusterCheckLimitsResponse `xml:"urn:vsan VsanQueryVcClusterCheckLimitsResponse,omitempty"`
	Fault_ *soap.Fault                                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanQueryVcClusterCheckLimitsBody) Fault() *soap.Fault { return b.Fault_ }

func VsanQueryVcClusterCheckLimits(ctx context.Context, r soap.RoundTripper, req *types.VsanQueryVcClusterCheckLimits) (*types.VsanQueryVcClusterCheckLimitsResponse, error) {
	var reqBody, resBody VsanQueryVcClusterCheckLimitsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanVcClusterQueryVerifyHealthSystemVersionsBody struct {
	Req    *types.VsanVcClusterQueryVerifyHealthSystemVersions         `xml:"urn:vsan VsanVcClusterQueryVerifyHealthSystemVersions,omitempty"`
	Res    *types.VsanVcClusterQueryVerifyHealthSystemVersionsResponse `xml:"urn:vsan VsanVcClusterQueryVerifyHealthSystemVersionsResponse,omitempty"`
	Fault_ *soap.Fault                                                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanVcClusterQueryVerifyHealthSystemVersionsBody) Fault() *soap.Fault { return b.Fault_ }

func VsanVcClusterQueryVerifyHealthSystemVersions(ctx context.Context, r soap.RoundTripper, req *types.VsanVcClusterQueryVerifyHealthSystemVersions) (*types.VsanVcClusterQueryVerifyHealthSystemVersionsResponse, error) {
	var reqBody, resBody VsanVcClusterQueryVerifyHealthSystemVersionsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanHealthSetVsanClusterTelemetryConfigBody struct {
	Req    *types.VsanHealthSetVsanClusterTelemetryConfig         `xml:"urn:vsan VsanHealthSetVsanClusterTelemetryConfig,omitempty"`
	Res    *types.VsanHealthSetVsanClusterTelemetryConfigResponse `xml:"urn:vsan VsanHealthSetVsanClusterTelemetryConfigResponse,omitempty"`
	Fault_ *soap.Fault                                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanHealthSetVsanClusterTelemetryConfigBody) Fault() *soap.Fault { return b.Fault_ }

func VsanHealthSetVsanClusterTelemetryConfig(ctx context.Context, r soap.RoundTripper, req *types.VsanHealthSetVsanClusterTelemetryConfig) (*types.VsanHealthSetVsanClusterTelemetryConfigResponse, error) {
	var reqBody, resBody VsanHealthSetVsanClusterTelemetryConfigBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RestoreClusterRebootWithNAMMBody struct {
	Req    *types.RestoreClusterRebootWithNAMM         `xml:"urn:vsan RestoreClusterRebootWithNAMM,omitempty"`
	Res    *types.RestoreClusterRebootWithNAMMResponse `xml:"urn:vsan RestoreClusterRebootWithNAMMResponse,omitempty"`
	Fault_ *soap.Fault                                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RestoreClusterRebootWithNAMMBody) Fault() *soap.Fault { return b.Fault_ }

func RestoreClusterRebootWithNAMM(ctx context.Context, r soap.RoundTripper, req *types.RestoreClusterRebootWithNAMM) (*types.RestoreClusterRebootWithNAMMResponse, error) {
	var reqBody, resBody RestoreClusterRebootWithNAMMBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanQueryVcClusterDataProtectionCfgSyncBody struct {
	Req    *types.VsanQueryVcClusterDataProtectionCfgSync         `xml:"urn:vsan VsanQueryVcClusterDataProtectionCfgSync,omitempty"`
	Res    *types.VsanQueryVcClusterDataProtectionCfgSyncResponse `xml:"urn:vsan VsanQueryVcClusterDataProtectionCfgSyncResponse,omitempty"`
	Fault_ *soap.Fault                                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanQueryVcClusterDataProtectionCfgSyncBody) Fault() *soap.Fault { return b.Fault_ }

func VsanQueryVcClusterDataProtectionCfgSync(ctx context.Context, r soap.RoundTripper, req *types.VsanQueryVcClusterDataProtectionCfgSync) (*types.VsanQueryVcClusterDataProtectionCfgSyncResponse, error) {
	var reqBody, resBody VsanQueryVcClusterDataProtectionCfgSyncBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanDownloadAndInstallVendorTool_TaskBody struct {
	Req    *types.VsanDownloadAndInstallVendorTool_Task         `xml:"urn:vsan VsanDownloadAndInstallVendorTool_Task,omitempty"`
	Res    *types.VsanDownloadAndInstallVendorTool_TaskResponse `xml:"urn:vsan VsanDownloadAndInstallVendorTool_TaskResponse,omitempty"`
	Fault_ *soap.Fault                                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanDownloadAndInstallVendorTool_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func VsanDownloadAndInstallVendorTool_Task(ctx context.Context, r soap.RoundTripper, req *types.VsanDownloadAndInstallVendorTool_Task) (*types.VsanDownloadAndInstallVendorTool_TaskResponse, error) {
	var reqBody, resBody VsanDownloadAndInstallVendorTool_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanHealthSetVsanClusterHealthCheckIntervalBody struct {
	Req    *types.VsanHealthSetVsanClusterHealthCheckInterval         `xml:"urn:vsan VsanHealthSetVsanClusterHealthCheckInterval,omitempty"`
	Res    *types.VsanHealthSetVsanClusterHealthCheckIntervalResponse `xml:"urn:vsan VsanHealthSetVsanClusterHealthCheckIntervalResponse,omitempty"`
	Fault_ *soap.Fault                                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanHealthSetVsanClusterHealthCheckIntervalBody) Fault() *soap.Fault { return b.Fault_ }

func VsanHealthSetVsanClusterHealthCheckInterval(ctx context.Context, r soap.RoundTripper, req *types.VsanHealthSetVsanClusterHealthCheckInterval) (*types.VsanHealthSetVsanClusterHealthCheckIntervalResponse, error) {
	var reqBody, resBody VsanHealthSetVsanClusterHealthCheckIntervalBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanVcClusterDpdLivenessBody struct {
	Req    *types.VsanVcClusterDpdLiveness         `xml:"urn:vsan VsanVcClusterDpdLiveness,omitempty"`
	Res    *types.VsanVcClusterDpdLivenessResponse `xml:"urn:vsan VsanVcClusterDpdLivenessResponse,omitempty"`
	Fault_ *soap.Fault                             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanVcClusterDpdLivenessBody) Fault() *soap.Fault { return b.Fault_ }

func VsanVcClusterDpdLiveness(ctx context.Context, r soap.RoundTripper, req *types.VsanVcClusterDpdLiveness) (*types.VsanVcClusterDpdLivenessResponse, error) {
	var reqBody, resBody VsanVcClusterDpdLivenessBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanHealthUninstallClusterBody struct {
	Req    *types.VsanHealthUninstallCluster         `xml:"urn:vsan VsanHealthUninstallCluster,omitempty"`
	Res    *types.VsanHealthUninstallClusterResponse `xml:"urn:vsan VsanHealthUninstallClusterResponse,omitempty"`
	Fault_ *soap.Fault                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanHealthUninstallClusterBody) Fault() *soap.Fault { return b.Fault_ }

func VsanHealthUninstallCluster(ctx context.Context, r soap.RoundTripper, req *types.VsanHealthUninstallCluster) (*types.VsanHealthUninstallClusterResponse, error) {
	var reqBody, resBody VsanHealthUninstallClusterBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanVcClusterDatastoreUsageBody struct {
	Req    *types.VsanVcClusterDatastoreUsage         `xml:"urn:vsan VsanVcClusterDatastoreUsage,omitempty"`
	Res    *types.VsanVcClusterDatastoreUsageResponse `xml:"urn:vsan VsanVcClusterDatastoreUsageResponse,omitempty"`
	Fault_ *soap.Fault                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanVcClusterDatastoreUsageBody) Fault() *soap.Fault { return b.Fault_ }

func VsanVcClusterDatastoreUsage(ctx context.Context, r soap.RoundTripper, req *types.VsanVcClusterDatastoreUsage) (*types.VsanVcClusterDatastoreUsageResponse, error) {
	var reqBody, resBody VsanVcClusterDatastoreUsageBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CaptureInternalStatsBody struct {
	Req    *types.CaptureInternalStats         `xml:"urn:vsan CaptureInternalStats,omitempty"`
	Res    *types.CaptureInternalStatsResponse `xml:"urn:vsan CaptureInternalStatsResponse,omitempty"`
	Fault_ *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CaptureInternalStatsBody) Fault() *soap.Fault { return b.Fault_ }

func CaptureInternalStats(ctx context.Context, r soap.RoundTripper, req *types.CaptureInternalStats) (*types.CaptureInternalStatsResponse, error) {
	var reqBody, resBody CaptureInternalStatsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanConvertDeploymentSpecFromJsonBody struct {
	Req    *types.VsanConvertDeploymentSpecFromJson         `xml:"urn:vsan VsanConvertDeploymentSpecFromJson,omitempty"`
	Res    *types.VsanConvertDeploymentSpecFromJsonResponse `xml:"urn:vsan VsanConvertDeploymentSpecFromJsonResponse,omitempty"`
	Fault_ *soap.Fault                                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanConvertDeploymentSpecFromJsonBody) Fault() *soap.Fault { return b.Fault_ }

func VsanConvertDeploymentSpecFromJson(ctx context.Context, r soap.RoundTripper, req *types.VsanConvertDeploymentSpecFromJson) (*types.VsanConvertDeploymentSpecFromJsonResponse, error) {
	var reqBody, resBody VsanConvertDeploymentSpecFromJsonBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanVcsaGetBootstrapTasksBody struct {
	Req    *types.VsanVcsaGetBootstrapTasks         `xml:"urn:vsan VsanVcsaGetBootstrapTasks,omitempty"`
	Res    *types.VsanVcsaGetBootstrapTasksResponse `xml:"urn:vsan VsanVcsaGetBootstrapTasksResponse,omitempty"`
	Fault_ *soap.Fault                              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanVcsaGetBootstrapTasksBody) Fault() *soap.Fault { return b.Fault_ }

func VsanVcsaGetBootstrapTasks(ctx context.Context, r soap.RoundTripper, req *types.VsanVcsaGetBootstrapTasks) (*types.VsanVcsaGetBootstrapTasksResponse, error) {
	var reqBody, resBody VsanVcsaGetBootstrapTasksBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanValidateDeploymentConfigBody struct {
	Req    *types.VsanValidateDeploymentConfig         `xml:"urn:vsan VsanValidateDeploymentConfig,omitempty"`
	Res    *types.VsanValidateDeploymentConfigResponse `xml:"urn:vsan VsanValidateDeploymentConfigResponse,omitempty"`
	Fault_ *soap.Fault                                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanValidateDeploymentConfigBody) Fault() *soap.Fault { return b.Fault_ }

func VsanValidateDeploymentConfig(ctx context.Context, r soap.RoundTripper, req *types.VsanValidateDeploymentConfig) (*types.VsanValidateDeploymentConfigResponse, error) {
	var reqBody, resBody VsanValidateDeploymentConfigBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanPostConfigForVcsaBody struct {
	Req    *types.VsanPostConfigForVcsa         `xml:"urn:vsan VsanPostConfigForVcsa,omitempty"`
	Res    *types.VsanPostConfigForVcsaResponse `xml:"urn:vsan VsanPostConfigForVcsaResponse,omitempty"`
	Fault_ *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanPostConfigForVcsaBody) Fault() *soap.Fault { return b.Fault_ }

func VsanPostConfigForVcsa(ctx context.Context, r soap.RoundTripper, req *types.VsanPostConfigForVcsa) (*types.VsanPostConfigForVcsaResponse, error) {
	var reqBody, resBody VsanPostConfigForVcsaBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanVcsaGetBootstrapProgressBody struct {
	Req    *types.VsanVcsaGetBootstrapProgress         `xml:"urn:vsan VsanVcsaGetBootstrapProgress,omitempty"`
	Res    *types.VsanVcsaGetBootstrapProgressResponse `xml:"urn:vsan VsanVcsaGetBootstrapProgressResponse,omitempty"`
	Fault_ *soap.Fault                                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanVcsaGetBootstrapProgressBody) Fault() *soap.Fault { return b.Fault_ }

func VsanVcsaGetBootstrapProgress(ctx context.Context, r soap.RoundTripper, req *types.VsanVcsaGetBootstrapProgress) (*types.VsanVcsaGetBootstrapProgressResponse, error) {
	var reqBody, resBody VsanVcsaGetBootstrapProgressBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanPrepareVsanForVcsaBody struct {
	Req    *types.VsanPrepareVsanForVcsa         `xml:"urn:vsan VsanPrepareVsanForVcsa,omitempty"`
	Res    *types.VsanPrepareVsanForVcsaResponse `xml:"urn:vsan VsanPrepareVsanForVcsaResponse,omitempty"`
	Fault_ *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanPrepareVsanForVcsaBody) Fault() *soap.Fault { return b.Fault_ }

func VsanPrepareVsanForVcsa(ctx context.Context, r soap.RoundTripper, req *types.VsanPrepareVsanForVcsa) (*types.VsanPrepareVsanForVcsaResponse, error) {
	var reqBody, resBody VsanPrepareVsanForVcsaBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanVcsaCancelBootstrapTaskBody struct {
	Req    *types.VsanVcsaCancelBootstrapTask         `xml:"urn:vsan VsanVcsaCancelBootstrapTask,omitempty"`
	Res    *types.VsanVcsaCancelBootstrapTaskResponse `xml:"urn:vsan VsanVcsaCancelBootstrapTaskResponse,omitempty"`
	Fault_ *soap.Fault                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanVcsaCancelBootstrapTaskBody) Fault() *soap.Fault { return b.Fault_ }

func VsanVcsaCancelBootstrapTask(ctx context.Context, r soap.RoundTripper, req *types.VsanVcsaCancelBootstrapTask) (*types.VsanVcsaCancelBootstrapTaskResponse, error) {
	var reqBody, resBody VsanVcsaCancelBootstrapTaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ComplianceResourceCheck_TaskBody struct {
	Req    *types.ComplianceResourceCheck_Task         `xml:"urn:vsan ComplianceResourceCheck_Task,omitempty"`
	Res    *types.ComplianceResourceCheck_TaskResponse `xml:"urn:vsan ComplianceResourceCheck_TaskResponse,omitempty"`
	Fault_ *soap.Fault                                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ComplianceResourceCheck_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func ComplianceResourceCheck_Task(ctx context.Context, r soap.RoundTripper, req *types.ComplianceResourceCheck_Task) (*types.ComplianceResourceCheck_TaskResponse, error) {
	var reqBody, resBody ComplianceResourceCheck_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type WhatIfDecomCheck_TaskBody struct {
	Req    *types.WhatIfDecomCheck_Task         `xml:"urn:vsan WhatIfDecomCheck_Task,omitempty"`
	Res    *types.WhatIfDecomCheck_TaskResponse `xml:"urn:vsan WhatIfDecomCheck_TaskResponse,omitempty"`
	Fault_ *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *WhatIfDecomCheck_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func WhatIfDecomCheck_Task(ctx context.Context, r soap.RoundTripper, req *types.WhatIfDecomCheck_Task) (*types.WhatIfDecomCheck_TaskResponse, error) {
	var reqBody, resBody WhatIfDecomCheck_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CnsSearchLabelsBody struct {
	Req    *types.CnsSearchLabels         `xml:"urn:vsan CnsSearchLabels,omitempty"`
	Res    *types.CnsSearchLabelsResponse `xml:"urn:vsan CnsSearchLabelsResponse,omitempty"`
	Fault_ *soap.Fault                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CnsSearchLabelsBody) Fault() *soap.Fault { return b.Fault_ }

func CnsSearchLabels(ctx context.Context, r soap.RoundTripper, req *types.CnsSearchLabels) (*types.CnsSearchLabelsResponse, error) {
	var reqBody, resBody CnsSearchLabelsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CnsExtendVolumeBody struct {
	Req    *types.CnsExtendVolume         `xml:"urn:vsan CnsExtendVolume,omitempty"`
	Res    *types.CnsExtendVolumeResponse `xml:"urn:vsan CnsExtendVolumeResponse,omitempty"`
	Fault_ *soap.Fault                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CnsExtendVolumeBody) Fault() *soap.Fault { return b.Fault_ }

func CnsExtendVolume(ctx context.Context, r soap.RoundTripper, req *types.CnsExtendVolume) (*types.CnsExtendVolumeResponse, error) {
	var reqBody, resBody CnsExtendVolumeBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CnsCreateVolumeBody struct {
	Req    *types.CnsCreateVolume         `xml:"urn:vsan CnsCreateVolume,omitempty"`
	Res    *types.CnsCreateVolumeResponse `xml:"urn:vsan CnsCreateVolumeResponse,omitempty"`
	Fault_ *soap.Fault                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CnsCreateVolumeBody) Fault() *soap.Fault { return b.Fault_ }

func CnsCreateVolume(ctx context.Context, r soap.RoundTripper, req *types.CnsCreateVolume) (*types.CnsCreateVolumeResponse, error) {
	var reqBody, resBody CnsCreateVolumeBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CnsGetLastFullSyncTimeBody struct {
	Req    *types.CnsGetLastFullSyncTime         `xml:"urn:vsan CnsGetLastFullSyncTime,omitempty"`
	Res    *types.CnsGetLastFullSyncTimeResponse `xml:"urn:vsan CnsGetLastFullSyncTimeResponse,omitempty"`
	Fault_ *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CnsGetLastFullSyncTimeBody) Fault() *soap.Fault { return b.Fault_ }

func CnsGetLastFullSyncTime(ctx context.Context, r soap.RoundTripper, req *types.CnsGetLastFullSyncTime) (*types.CnsGetLastFullSyncTimeResponse, error) {
	var reqBody, resBody CnsGetLastFullSyncTimeBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CnsAttachVolumeBody struct {
	Req    *types.CnsAttachVolume         `xml:"urn:vsan CnsAttachVolume,omitempty"`
	Res    *types.CnsAttachVolumeResponse `xml:"urn:vsan CnsAttachVolumeResponse,omitempty"`
	Fault_ *soap.Fault                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CnsAttachVolumeBody) Fault() *soap.Fault { return b.Fault_ }

func CnsAttachVolume(ctx context.Context, r soap.RoundTripper, req *types.CnsAttachVolume) (*types.CnsAttachVolumeResponse, error) {
	var reqBody, resBody CnsAttachVolumeBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CnsTriggerFullSyncBody struct {
	Req    *types.CnsTriggerFullSync         `xml:"urn:vsan CnsTriggerFullSync,omitempty"`
	Res    *types.CnsTriggerFullSyncResponse `xml:"urn:vsan CnsTriggerFullSyncResponse,omitempty"`
	Fault_ *soap.Fault                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CnsTriggerFullSyncBody) Fault() *soap.Fault { return b.Fault_ }

func CnsTriggerFullSync(ctx context.Context, r soap.RoundTripper, req *types.CnsTriggerFullSync) (*types.CnsTriggerFullSyncResponse, error) {
	var reqBody, resBody CnsTriggerFullSyncBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CnsUpdateVolumeMetadataBody struct {
	Req    *types.CnsUpdateVolumeMetadata         `xml:"urn:vsan CnsUpdateVolumeMetadata,omitempty"`
	Res    *types.CnsUpdateVolumeMetadataResponse `xml:"urn:vsan CnsUpdateVolumeMetadataResponse,omitempty"`
	Fault_ *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CnsUpdateVolumeMetadataBody) Fault() *soap.Fault { return b.Fault_ }

func CnsUpdateVolumeMetadata(ctx context.Context, r soap.RoundTripper, req *types.CnsUpdateVolumeMetadata) (*types.CnsUpdateVolumeMetadataResponse, error) {
	var reqBody, resBody CnsUpdateVolumeMetadataBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CnsQueryVolumeBody struct {
	Req    *types.CnsQueryVolume         `xml:"urn:vsan CnsQueryVolume,omitempty"`
	Res    *types.CnsQueryVolumeResponse `xml:"urn:vsan CnsQueryVolumeResponse,omitempty"`
	Fault_ *soap.Fault                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CnsQueryVolumeBody) Fault() *soap.Fault { return b.Fault_ }

func CnsQueryVolume(ctx context.Context, r soap.RoundTripper, req *types.CnsQueryVolume) (*types.CnsQueryVolumeResponse, error) {
	var reqBody, resBody CnsQueryVolumeBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CnsDetachVolumeBody struct {
	Req    *types.CnsDetachVolume         `xml:"urn:vsan CnsDetachVolume,omitempty"`
	Res    *types.CnsDetachVolumeResponse `xml:"urn:vsan CnsDetachVolumeResponse,omitempty"`
	Fault_ *soap.Fault                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CnsDetachVolumeBody) Fault() *soap.Fault { return b.Fault_ }

func CnsDetachVolume(ctx context.Context, r soap.RoundTripper, req *types.CnsDetachVolume) (*types.CnsDetachVolumeResponse, error) {
	var reqBody, resBody CnsDetachVolumeBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CnsQueryAllVolumeBody struct {
	Req    *types.CnsQueryAllVolume         `xml:"urn:vsan CnsQueryAllVolume,omitempty"`
	Res    *types.CnsQueryAllVolumeResponse `xml:"urn:vsan CnsQueryAllVolumeResponse,omitempty"`
	Fault_ *soap.Fault                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CnsQueryAllVolumeBody) Fault() *soap.Fault { return b.Fault_ }

func CnsQueryAllVolume(ctx context.Context, r soap.RoundTripper, req *types.CnsQueryAllVolume) (*types.CnsQueryAllVolumeResponse, error) {
	var reqBody, resBody CnsQueryAllVolumeBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CnsDeleteVolumeBody struct {
	Req    *types.CnsDeleteVolume         `xml:"urn:vsan CnsDeleteVolume,omitempty"`
	Res    *types.CnsDeleteVolumeResponse `xml:"urn:vsan CnsDeleteVolumeResponse,omitempty"`
	Fault_ *soap.Fault                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CnsDeleteVolumeBody) Fault() *soap.Fault { return b.Fault_ }

func CnsDeleteVolume(ctx context.Context, r soap.RoundTripper, req *types.CnsDeleteVolume) (*types.CnsDeleteVolumeResponse, error) {
	var reqBody, resBody CnsDeleteVolumeBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanQueryDecomResourceCheckAsync_TaskBody struct {
	Req    *types.VsanQueryDecomResourceCheckAsync_Task         `xml:"urn:vsan VsanQueryDecomResourceCheckAsync_Task,omitempty"`
	Res    *types.VsanQueryDecomResourceCheckAsync_TaskResponse `xml:"urn:vsan VsanQueryDecomResourceCheckAsync_TaskResponse,omitempty"`
	Fault_ *soap.Fault                                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanQueryDecomResourceCheckAsync_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func VsanQueryDecomResourceCheckAsync_Task(ctx context.Context, r soap.RoundTripper, req *types.VsanQueryDecomResourceCheckAsync_Task) (*types.VsanQueryDecomResourceCheckAsync_TaskResponse, error) {
	var reqBody, resBody VsanQueryDecomResourceCheckAsync_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanQueryComplianceResourceCheckAsync_TaskBody struct {
	Req    *types.VsanQueryComplianceResourceCheckAsync_Task         `xml:"urn:vsan VsanQueryComplianceResourceCheckAsync_Task,omitempty"`
	Res    *types.VsanQueryComplianceResourceCheckAsync_TaskResponse `xml:"urn:vsan VsanQueryComplianceResourceCheckAsync_TaskResponse,omitempty"`
	Fault_ *soap.Fault                                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanQueryComplianceResourceCheckAsync_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func VsanQueryComplianceResourceCheckAsync_Task(ctx context.Context, r soap.RoundTripper, req *types.VsanQueryComplianceResourceCheckAsync_Task) (*types.VsanQueryComplianceResourceCheckAsync_TaskResponse, error) {
	var reqBody, resBody VsanQueryComplianceResourceCheckAsync_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanGetComplianceResourceCheckStatusBody struct {
	Req    *types.VsanGetComplianceResourceCheckStatus         `xml:"urn:vsan VsanGetComplianceResourceCheckStatus,omitempty"`
	Res    *types.VsanGetComplianceResourceCheckStatusResponse `xml:"urn:vsan VsanGetComplianceResourceCheckStatusResponse,omitempty"`
	Fault_ *soap.Fault                                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanGetComplianceResourceCheckStatusBody) Fault() *soap.Fault { return b.Fault_ }

func VsanGetComplianceResourceCheckStatus(ctx context.Context, r soap.RoundTripper, req *types.VsanGetComplianceResourceCheckStatus) (*types.VsanGetComplianceResourceCheckStatusResponse, error) {
	var reqBody, resBody VsanGetComplianceResourceCheckStatusBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanGetResourceCheckStatus2Body struct {
	Req    *types.VsanGetResourceCheckStatus2         `xml:"urn:vsan VsanGetResourceCheckStatus2,omitempty"`
	Res    *types.VsanGetResourceCheckStatus2Response `xml:"urn:vsan VsanGetResourceCheckStatus2Response,omitempty"`
	Fault_ *soap.Fault                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanGetResourceCheckStatus2Body) Fault() *soap.Fault { return b.Fault_ }

func VsanGetResourceCheckStatus2(ctx context.Context, r soap.RoundTripper, req *types.VsanGetResourceCheckStatus2) (*types.VsanGetResourceCheckStatus2Response, error) {
	var reqBody, resBody VsanGetResourceCheckStatus2Body

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanPrepareDefragBody struct {
	Req    *types.VsanPrepareDefrag         `xml:"urn:vsan VsanPrepareDefrag,omitempty"`
	Res    *types.VsanPrepareDefragResponse `xml:"urn:vsan VsanPrepareDefragResponse,omitempty"`
	Fault_ *soap.Fault                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanPrepareDefragBody) Fault() *soap.Fault { return b.Fault_ }

func VsanPrepareDefrag(ctx context.Context, r soap.RoundTripper, req *types.VsanPrepareDefrag) (*types.VsanPrepareDefragResponse, error) {
	var reqBody, resBody VsanPrepareDefragBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanCleanupDefragBody struct {
	Req    *types.VsanCleanupDefrag         `xml:"urn:vsan VsanCleanupDefrag,omitempty"`
	Res    *types.VsanCleanupDefragResponse `xml:"urn:vsan VsanCleanupDefragResponse,omitempty"`
	Fault_ *soap.Fault                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanCleanupDefragBody) Fault() *soap.Fault { return b.Fault_ }

func VsanCleanupDefrag(ctx context.Context, r soap.RoundTripper, req *types.VsanCleanupDefrag) (*types.VsanCleanupDefragResponse, error) {
	var reqBody, resBody VsanCleanupDefragBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanVdsMigrateVssBody struct {
	Req    *types.VsanVdsMigrateVss         `xml:"urn:vsan VsanVdsMigrateVss,omitempty"`
	Res    *types.VsanVdsMigrateVssResponse `xml:"urn:vsan VsanVdsMigrateVssResponse,omitempty"`
	Fault_ *soap.Fault                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanVdsMigrateVssBody) Fault() *soap.Fault { return b.Fault_ }

func VsanVdsMigrateVss(ctx context.Context, r soap.RoundTripper, req *types.VsanVdsMigrateVss) (*types.VsanVdsMigrateVssResponse, error) {
	var reqBody, resBody VsanVdsMigrateVssBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanVdsGetMigrationPlanBody struct {
	Req    *types.VsanVdsGetMigrationPlan         `xml:"urn:vsan VsanVdsGetMigrationPlan,omitempty"`
	Res    *types.VsanVdsGetMigrationPlanResponse `xml:"urn:vsan VsanVdsGetMigrationPlanResponse,omitempty"`
	Fault_ *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanVdsGetMigrationPlanBody) Fault() *soap.Fault { return b.Fault_ }

func VsanVdsGetMigrationPlan(ctx context.Context, r soap.RoundTripper, req *types.VsanVdsGetMigrationPlan) (*types.VsanVdsGetMigrationPlanResponse, error) {
	var reqBody, resBody VsanVdsGetMigrationPlanBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanVssMigrateVdsBody struct {
	Req    *types.VsanVssMigrateVds         `xml:"urn:vsan VsanVssMigrateVds,omitempty"`
	Res    *types.VsanVssMigrateVdsResponse `xml:"urn:vsan VsanVssMigrateVdsResponse,omitempty"`
	Fault_ *soap.Fault                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanVssMigrateVdsBody) Fault() *soap.Fault { return b.Fault_ }

func VsanVssMigrateVds(ctx context.Context, r soap.RoundTripper, req *types.VsanVssMigrateVds) (*types.VsanVssMigrateVdsResponse, error) {
	var reqBody, resBody VsanVssMigrateVdsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanRollbackVdsToVssBody struct {
	Req    *types.VsanRollbackVdsToVss         `xml:"urn:vsan VsanRollbackVdsToVss,omitempty"`
	Res    *types.VsanRollbackVdsToVssResponse `xml:"urn:vsan VsanRollbackVdsToVssResponse,omitempty"`
	Fault_ *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanRollbackVdsToVssBody) Fault() *soap.Fault { return b.Fault_ }

func VsanRollbackVdsToVss(ctx context.Context, r soap.RoundTripper, req *types.VsanRollbackVdsToVss) (*types.VsanRollbackVdsToVssResponse, error) {
	var reqBody, resBody VsanRollbackVdsToVssBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VosSetVsanObjectPolicyBody struct {
	Req    *types.VosSetVsanObjectPolicy         `xml:"urn:vsan VosSetVsanObjectPolicy,omitempty"`
	Res    *types.VosSetVsanObjectPolicyResponse `xml:"urn:vsan VosSetVsanObjectPolicyResponse,omitempty"`
	Fault_ *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VosSetVsanObjectPolicyBody) Fault() *soap.Fault { return b.Fault_ }

func VosSetVsanObjectPolicy(ctx context.Context, r soap.RoundTripper, req *types.VosSetVsanObjectPolicy) (*types.VosSetVsanObjectPolicyResponse, error) {
	var reqBody, resBody VosSetVsanObjectPolicyBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanDeleteObjects_TaskBody struct {
	Req    *types.VsanDeleteObjects_Task         `xml:"urn:vsan VsanDeleteObjects_Task,omitempty"`
	Res    *types.VsanDeleteObjects_TaskResponse `xml:"urn:vsan VsanDeleteObjects_TaskResponse,omitempty"`
	Fault_ *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanDeleteObjects_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func VsanDeleteObjects_Task(ctx context.Context, r soap.RoundTripper, req *types.VsanDeleteObjects_Task) (*types.VsanDeleteObjects_TaskResponse, error) {
	var reqBody, resBody VsanDeleteObjects_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VosQueryVsanObjectInformationBody struct {
	Req    *types.VosQueryVsanObjectInformation         `xml:"urn:vsan VosQueryVsanObjectInformation,omitempty"`
	Res    *types.VosQueryVsanObjectInformationResponse `xml:"urn:vsan VosQueryVsanObjectInformationResponse,omitempty"`
	Fault_ *soap.Fault                                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VosQueryVsanObjectInformationBody) Fault() *soap.Fault { return b.Fault_ }

func VosQueryVsanObjectInformation(ctx context.Context, r soap.RoundTripper, req *types.VosQueryVsanObjectInformation) (*types.VosQueryVsanObjectInformationResponse, error) {
	var reqBody, resBody VosQueryVsanObjectInformationBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RelayoutObjectsBody struct {
	Req    *types.RelayoutObjects         `xml:"urn:vsan RelayoutObjects,omitempty"`
	Res    *types.RelayoutObjectsResponse `xml:"urn:vsan RelayoutObjectsResponse,omitempty"`
	Fault_ *soap.Fault                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RelayoutObjectsBody) Fault() *soap.Fault { return b.Fault_ }

func RelayoutObjects(ctx context.Context, r soap.RoundTripper, req *types.RelayoutObjects) (*types.RelayoutObjectsResponse, error) {
	var reqBody, resBody RelayoutObjectsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanQueryInaccessibleVmSwapObjectsBody struct {
	Req    *types.VsanQueryInaccessibleVmSwapObjects         `xml:"urn:vsan VsanQueryInaccessibleVmSwapObjects,omitempty"`
	Res    *types.VsanQueryInaccessibleVmSwapObjectsResponse `xml:"urn:vsan VsanQueryInaccessibleVmSwapObjectsResponse,omitempty"`
	Fault_ *soap.Fault                                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanQueryInaccessibleVmSwapObjectsBody) Fault() *soap.Fault { return b.Fault_ }

func VsanQueryInaccessibleVmSwapObjects(ctx context.Context, r soap.RoundTripper, req *types.VsanQueryInaccessibleVmSwapObjects) (*types.VsanQueryInaccessibleVmSwapObjectsResponse, error) {
	var reqBody, resBody VsanQueryInaccessibleVmSwapObjectsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QuerySyncingVsanObjectsSummaryBody struct {
	Req    *types.QuerySyncingVsanObjectsSummary         `xml:"urn:vsan QuerySyncingVsanObjectsSummary,omitempty"`
	Res    *types.QuerySyncingVsanObjectsSummaryResponse `xml:"urn:vsan QuerySyncingVsanObjectsSummaryResponse,omitempty"`
	Fault_ *soap.Fault                                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QuerySyncingVsanObjectsSummaryBody) Fault() *soap.Fault { return b.Fault_ }

func QuerySyncingVsanObjectsSummary(ctx context.Context, r soap.RoundTripper, req *types.QuerySyncingVsanObjectsSummary) (*types.QuerySyncingVsanObjectsSummaryResponse, error) {
	var reqBody, resBody QuerySyncingVsanObjectsSummaryBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanQueryInactiveDataProtectionSpaceUsageBody struct {
	Req    *types.VsanQueryInactiveDataProtectionSpaceUsage         `xml:"urn:vsan VsanQueryInactiveDataProtectionSpaceUsage,omitempty"`
	Res    *types.VsanQueryInactiveDataProtectionSpaceUsageResponse `xml:"urn:vsan VsanQueryInactiveDataProtectionSpaceUsageResponse,omitempty"`
	Fault_ *soap.Fault                                              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanQueryInactiveDataProtectionSpaceUsageBody) Fault() *soap.Fault { return b.Fault_ }

func VsanQueryInactiveDataProtectionSpaceUsage(ctx context.Context, r soap.RoundTripper, req *types.VsanQueryInactiveDataProtectionSpaceUsage) (*types.VsanQueryInactiveDataProtectionSpaceUsageResponse, error) {
	var reqBody, resBody VsanQueryInactiveDataProtectionSpaceUsageBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanQueryObjectIdentitiesBody struct {
	Req    *types.VsanQueryObjectIdentities         `xml:"urn:vsan VsanQueryObjectIdentities,omitempty"`
	Res    *types.VsanQueryObjectIdentitiesResponse `xml:"urn:vsan VsanQueryObjectIdentitiesResponse,omitempty"`
	Fault_ *soap.Fault                              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanQueryObjectIdentitiesBody) Fault() *soap.Fault { return b.Fault_ }

func VsanQueryObjectIdentities(ctx context.Context, r soap.RoundTripper, req *types.VsanQueryObjectIdentities) (*types.VsanQueryObjectIdentitiesResponse, error) {
	var reqBody, resBody VsanQueryObjectIdentitiesBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type HostSpbmRetrieveAllPolicyInfoBody struct {
	Req    *types.HostSpbmRetrieveAllPolicyInfo         `xml:"urn:vsan HostSpbmRetrieveAllPolicyInfo,omitempty"`
	Res    *types.HostSpbmRetrieveAllPolicyInfoResponse `xml:"urn:vsan HostSpbmRetrieveAllPolicyInfoResponse,omitempty"`
	Fault_ *soap.Fault                                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *HostSpbmRetrieveAllPolicyInfoBody) Fault() *soap.Fault { return b.Fault_ }

func HostSpbmRetrieveAllPolicyInfo(ctx context.Context, r soap.RoundTripper, req *types.HostSpbmRetrieveAllPolicyInfo) (*types.HostSpbmRetrieveAllPolicyInfoResponse, error) {
	var reqBody, resBody HostSpbmRetrieveAllPolicyInfoBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type HostSpbmComputeHashInfoBody struct {
	Req    *types.HostSpbmComputeHashInfo         `xml:"urn:vsan HostSpbmComputeHashInfo,omitempty"`
	Res    *types.HostSpbmComputeHashInfoResponse `xml:"urn:vsan HostSpbmComputeHashInfoResponse,omitempty"`
	Fault_ *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *HostSpbmComputeHashInfoBody) Fault() *soap.Fault { return b.Fault_ }

func HostSpbmComputeHashInfo(ctx context.Context, r soap.RoundTripper, req *types.HostSpbmComputeHashInfo) (*types.HostSpbmComputeHashInfoResponse, error) {
	var reqBody, resBody HostSpbmComputeHashInfoBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type HostSpbmRetrieveAllDatastoreInfoBody struct {
	Req    *types.HostSpbmRetrieveAllDatastoreInfo         `xml:"urn:vsan HostSpbmRetrieveAllDatastoreInfo,omitempty"`
	Res    *types.HostSpbmRetrieveAllDatastoreInfoResponse `xml:"urn:vsan HostSpbmRetrieveAllDatastoreInfoResponse,omitempty"`
	Fault_ *soap.Fault                                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *HostSpbmRetrieveAllDatastoreInfoBody) Fault() *soap.Fault { return b.Fault_ }

func HostSpbmRetrieveAllDatastoreInfo(ctx context.Context, r soap.RoundTripper, req *types.HostSpbmRetrieveAllDatastoreInfo) (*types.HostSpbmRetrieveAllDatastoreInfoResponse, error) {
	var reqBody, resBody HostSpbmRetrieveAllDatastoreInfoBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type HostSpbmFetchDefaultPolicyBlobBody struct {
	Req    *types.HostSpbmFetchDefaultPolicyBlob         `xml:"urn:vsan HostSpbmFetchDefaultPolicyBlob,omitempty"`
	Res    *types.HostSpbmFetchDefaultPolicyBlobResponse `xml:"urn:vsan HostSpbmFetchDefaultPolicyBlobResponse,omitempty"`
	Fault_ *soap.Fault                                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *HostSpbmFetchDefaultPolicyBlobBody) Fault() *soap.Fault { return b.Fault_ }

func HostSpbmFetchDefaultPolicyBlob(ctx context.Context, r soap.RoundTripper, req *types.HostSpbmFetchDefaultPolicyBlob) (*types.HostSpbmFetchDefaultPolicyBlobResponse, error) {
	var reqBody, resBody HostSpbmFetchDefaultPolicyBlobBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type HostSpbmPushPolicyInfoBody struct {
	Req    *types.HostSpbmPushPolicyInfo         `xml:"urn:vsan HostSpbmPushPolicyInfo,omitempty"`
	Res    *types.HostSpbmPushPolicyInfoResponse `xml:"urn:vsan HostSpbmPushPolicyInfoResponse,omitempty"`
	Fault_ *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *HostSpbmPushPolicyInfoBody) Fault() *soap.Fault { return b.Fault_ }

func HostSpbmPushPolicyInfo(ctx context.Context, r soap.RoundTripper, req *types.HostSpbmPushPolicyInfo) (*types.HostSpbmPushPolicyInfoResponse, error) {
	var reqBody, resBody HostSpbmPushPolicyInfoBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type HostSpbmPushDatastoreInfoBody struct {
	Req    *types.HostSpbmPushDatastoreInfo         `xml:"urn:vsan HostSpbmPushDatastoreInfo,omitempty"`
	Res    *types.HostSpbmPushDatastoreInfoResponse `xml:"urn:vsan HostSpbmPushDatastoreInfoResponse,omitempty"`
	Fault_ *soap.Fault                              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *HostSpbmPushDatastoreInfoBody) Fault() *soap.Fault { return b.Fault_ }

func HostSpbmPushDatastoreInfo(ctx context.Context, r soap.RoundTripper, req *types.HostSpbmPushDatastoreInfo) (*types.HostSpbmPushDatastoreInfoResponse, error) {
	var reqBody, resBody HostSpbmPushDatastoreInfoBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type HostSpbmFetchApplicablePolicyBlobBody struct {
	Req    *types.HostSpbmFetchApplicablePolicyBlob         `xml:"urn:vsan HostSpbmFetchApplicablePolicyBlob,omitempty"`
	Res    *types.HostSpbmFetchApplicablePolicyBlobResponse `xml:"urn:vsan HostSpbmFetchApplicablePolicyBlobResponse,omitempty"`
	Fault_ *soap.Fault                                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *HostSpbmFetchApplicablePolicyBlobBody) Fault() *soap.Fault { return b.Fault_ }

func HostSpbmFetchApplicablePolicyBlob(ctx context.Context, r soap.RoundTripper, req *types.HostSpbmFetchApplicablePolicyBlob) (*types.HostSpbmFetchApplicablePolicyBlobResponse, error) {
	var reqBody, resBody HostSpbmFetchApplicablePolicyBlobBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanRecoveryRecoverComponentBody struct {
	Req    *types.VsanRecoveryRecoverComponent         `xml:"urn:vsan VsanRecoveryRecoverComponent,omitempty"`
	Res    *types.VsanRecoveryRecoverComponentResponse `xml:"urn:vsan VsanRecoveryRecoverComponentResponse,omitempty"`
	Fault_ *soap.Fault                                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanRecoveryRecoverComponentBody) Fault() *soap.Fault { return b.Fault_ }

func VsanRecoveryRecoverComponent(ctx context.Context, r soap.RoundTripper, req *types.VsanRecoveryRecoverComponent) (*types.VsanRecoveryRecoverComponentResponse, error) {
	var reqBody, resBody VsanRecoveryRecoverComponentBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanRecoveryServiceGetDummyCmmdsModeBody struct {
	Req    *types.VsanRecoveryServiceGetDummyCmmdsMode         `xml:"urn:vsan VsanRecoveryServiceGetDummyCmmdsMode,omitempty"`
	Res    *types.VsanRecoveryServiceGetDummyCmmdsModeResponse `xml:"urn:vsan VsanRecoveryServiceGetDummyCmmdsModeResponse,omitempty"`
	Fault_ *soap.Fault                                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanRecoveryServiceGetDummyCmmdsModeBody) Fault() *soap.Fault { return b.Fault_ }

func VsanRecoveryServiceGetDummyCmmdsMode(ctx context.Context, r soap.RoundTripper, req *types.VsanRecoveryServiceGetDummyCmmdsMode) (*types.VsanRecoveryServiceGetDummyCmmdsModeResponse, error) {
	var reqBody, resBody VsanRecoveryServiceGetDummyCmmdsModeBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanRecoveryGetRecoverableComponentsBody struct {
	Req    *types.VsanRecoveryGetRecoverableComponents         `xml:"urn:vsan VsanRecoveryGetRecoverableComponents,omitempty"`
	Res    *types.VsanRecoveryGetRecoverableComponentsResponse `xml:"urn:vsan VsanRecoveryGetRecoverableComponentsResponse,omitempty"`
	Fault_ *soap.Fault                                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanRecoveryGetRecoverableComponentsBody) Fault() *soap.Fault { return b.Fault_ }

func VsanRecoveryGetRecoverableComponents(ctx context.Context, r soap.RoundTripper, req *types.VsanRecoveryGetRecoverableComponents) (*types.VsanRecoveryGetRecoverableComponentsResponse, error) {
	var reqBody, resBody VsanRecoveryGetRecoverableComponentsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanRecoveryGetAllObjectsStatusBody struct {
	Req    *types.VsanRecoveryGetAllObjectsStatus         `xml:"urn:vsan VsanRecoveryGetAllObjectsStatus,omitempty"`
	Res    *types.VsanRecoveryGetAllObjectsStatusResponse `xml:"urn:vsan VsanRecoveryGetAllObjectsStatusResponse,omitempty"`
	Fault_ *soap.Fault                                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanRecoveryGetAllObjectsStatusBody) Fault() *soap.Fault { return b.Fault_ }

func VsanRecoveryGetAllObjectsStatus(ctx context.Context, r soap.RoundTripper, req *types.VsanRecoveryGetAllObjectsStatus) (*types.VsanRecoveryGetAllObjectsStatusResponse, error) {
	var reqBody, resBody VsanRecoveryGetAllObjectsStatusBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanRecoveryGetObjectRecoveryInfoBody struct {
	Req    *types.VsanRecoveryGetObjectRecoveryInfo         `xml:"urn:vsan VsanRecoveryGetObjectRecoveryInfo,omitempty"`
	Res    *types.VsanRecoveryGetObjectRecoveryInfoResponse `xml:"urn:vsan VsanRecoveryGetObjectRecoveryInfoResponse,omitempty"`
	Fault_ *soap.Fault                                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanRecoveryGetObjectRecoveryInfoBody) Fault() *soap.Fault { return b.Fault_ }

func VsanRecoveryGetObjectRecoveryInfo(ctx context.Context, r soap.RoundTripper, req *types.VsanRecoveryGetObjectRecoveryInfo) (*types.VsanRecoveryGetObjectRecoveryInfoResponse, error) {
	var reqBody, resBody VsanRecoveryGetObjectRecoveryInfoBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanRecoveryDeleteLsomCompLocalBody struct {
	Req    *types.VsanRecoveryDeleteLsomCompLocal         `xml:"urn:vsan VsanRecoveryDeleteLsomCompLocal,omitempty"`
	Res    *types.VsanRecoveryDeleteLsomCompLocalResponse `xml:"urn:vsan VsanRecoveryDeleteLsomCompLocalResponse,omitempty"`
	Fault_ *soap.Fault                                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanRecoveryDeleteLsomCompLocalBody) Fault() *soap.Fault { return b.Fault_ }

func VsanRecoveryDeleteLsomCompLocal(ctx context.Context, r soap.RoundTripper, req *types.VsanRecoveryDeleteLsomCompLocal) (*types.VsanRecoveryDeleteLsomCompLocalResponse, error) {
	var reqBody, resBody VsanRecoveryDeleteLsomCompLocalBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanRecoveryGetDiskUuidForLocalCompBody struct {
	Req    *types.VsanRecoveryGetDiskUuidForLocalComp         `xml:"urn:vsan VsanRecoveryGetDiskUuidForLocalComp,omitempty"`
	Res    *types.VsanRecoveryGetDiskUuidForLocalCompResponse `xml:"urn:vsan VsanRecoveryGetDiskUuidForLocalCompResponse,omitempty"`
	Fault_ *soap.Fault                                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanRecoveryGetDiskUuidForLocalCompBody) Fault() *soap.Fault { return b.Fault_ }

func VsanRecoveryGetDiskUuidForLocalComp(ctx context.Context, r soap.RoundTripper, req *types.VsanRecoveryGetDiskUuidForLocalComp) (*types.VsanRecoveryGetDiskUuidForLocalCompResponse, error) {
	var reqBody, resBody VsanRecoveryGetDiskUuidForLocalCompBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanRecoveryCreateObjectBody struct {
	Req    *types.VsanRecoveryCreateObject         `xml:"urn:vsan VsanRecoveryCreateObject,omitempty"`
	Res    *types.VsanRecoveryCreateObjectResponse `xml:"urn:vsan VsanRecoveryCreateObjectResponse,omitempty"`
	Fault_ *soap.Fault                             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanRecoveryCreateObjectBody) Fault() *soap.Fault { return b.Fault_ }

func VsanRecoveryCreateObject(ctx context.Context, r soap.RoundTripper, req *types.VsanRecoveryCreateObject) (*types.VsanRecoveryCreateObjectResponse, error) {
	var reqBody, resBody VsanRecoveryCreateObjectBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanRecoveryGetDiskAndHostForCompBody struct {
	Req    *types.VsanRecoveryGetDiskAndHostForComp         `xml:"urn:vsan VsanRecoveryGetDiskAndHostForComp,omitempty"`
	Res    *types.VsanRecoveryGetDiskAndHostForCompResponse `xml:"urn:vsan VsanRecoveryGetDiskAndHostForCompResponse,omitempty"`
	Fault_ *soap.Fault                                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanRecoveryGetDiskAndHostForCompBody) Fault() *soap.Fault { return b.Fault_ }

func VsanRecoveryGetDiskAndHostForComp(ctx context.Context, r soap.RoundTripper, req *types.VsanRecoveryGetDiskAndHostForComp) (*types.VsanRecoveryGetDiskAndHostForCompResponse, error) {
	var reqBody, resBody VsanRecoveryGetDiskAndHostForCompBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanRecoveryServiceVersionBody struct {
	Req    *types.VsanRecoveryServiceVersion         `xml:"urn:vsan VsanRecoveryServiceVersion,omitempty"`
	Res    *types.VsanRecoveryServiceVersionResponse `xml:"urn:vsan VsanRecoveryServiceVersionResponse,omitempty"`
	Fault_ *soap.Fault                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanRecoveryServiceVersionBody) Fault() *soap.Fault { return b.Fault_ }

func VsanRecoveryServiceVersion(ctx context.Context, r soap.RoundTripper, req *types.VsanRecoveryServiceVersion) (*types.VsanRecoveryServiceVersionResponse, error) {
	var reqBody, resBody VsanRecoveryServiceVersionBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanRecoveryGetTaskStatusBody struct {
	Req    *types.VsanRecoveryGetTaskStatus         `xml:"urn:vsan VsanRecoveryGetTaskStatus,omitempty"`
	Res    *types.VsanRecoveryGetTaskStatusResponse `xml:"urn:vsan VsanRecoveryGetTaskStatusResponse,omitempty"`
	Fault_ *soap.Fault                              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanRecoveryGetTaskStatusBody) Fault() *soap.Fault { return b.Fault_ }

func VsanRecoveryGetTaskStatus(ctx context.Context, r soap.RoundTripper, req *types.VsanRecoveryGetTaskStatus) (*types.VsanRecoveryGetTaskStatusResponse, error) {
	var reqBody, resBody VsanRecoveryGetTaskStatusBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanRecoveryServiceSetDummyCmmdsModeBody struct {
	Req    *types.VsanRecoveryServiceSetDummyCmmdsMode         `xml:"urn:vsan VsanRecoveryServiceSetDummyCmmdsMode,omitempty"`
	Res    *types.VsanRecoveryServiceSetDummyCmmdsModeResponse `xml:"urn:vsan VsanRecoveryServiceSetDummyCmmdsModeResponse,omitempty"`
	Fault_ *soap.Fault                                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanRecoveryServiceSetDummyCmmdsModeBody) Fault() *soap.Fault { return b.Fault_ }

func VsanRecoveryServiceSetDummyCmmdsMode(ctx context.Context, r soap.RoundTripper, req *types.VsanRecoveryServiceSetDummyCmmdsMode) (*types.VsanRecoveryServiceSetDummyCmmdsModeResponse, error) {
	var reqBody, resBody VsanRecoveryServiceSetDummyCmmdsModeBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanRecoveryGetObjectPropsBody struct {
	Req    *types.VsanRecoveryGetObjectProps         `xml:"urn:vsan VsanRecoveryGetObjectProps,omitempty"`
	Res    *types.VsanRecoveryGetObjectPropsResponse `xml:"urn:vsan VsanRecoveryGetObjectPropsResponse,omitempty"`
	Fault_ *soap.Fault                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanRecoveryGetObjectPropsBody) Fault() *soap.Fault { return b.Fault_ }

func VsanRecoveryGetObjectProps(ctx context.Context, r soap.RoundTripper, req *types.VsanRecoveryGetObjectProps) (*types.VsanRecoveryGetObjectPropsResponse, error) {
	var reqBody, resBody VsanRecoveryGetObjectPropsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanRecoveryRecoverObjectBody struct {
	Req    *types.VsanRecoveryRecoverObject         `xml:"urn:vsan VsanRecoveryRecoverObject,omitempty"`
	Res    *types.VsanRecoveryRecoverObjectResponse `xml:"urn:vsan VsanRecoveryRecoverObjectResponse,omitempty"`
	Fault_ *soap.Fault                              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanRecoveryRecoverObjectBody) Fault() *soap.Fault { return b.Fault_ }

func VsanRecoveryRecoverObject(ctx context.Context, r soap.RoundTripper, req *types.VsanRecoveryRecoverObject) (*types.VsanRecoveryRecoverObjectResponse, error) {
	var reqBody, resBody VsanRecoveryRecoverObjectBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VSANVcIsWitnessHostBody struct {
	Req    *types.VSANVcIsWitnessHost         `xml:"urn:vsan VSANVcIsWitnessHost,omitempty"`
	Res    *types.VSANVcIsWitnessHostResponse `xml:"urn:vsan VSANVcIsWitnessHostResponse,omitempty"`
	Fault_ *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VSANVcIsWitnessHostBody) Fault() *soap.Fault { return b.Fault_ }

func VSANVcIsWitnessHost(ctx context.Context, r soap.RoundTripper, req *types.VSANVcIsWitnessHost) (*types.VSANVcIsWitnessHostResponse, error) {
	var reqBody, resBody VSANVcIsWitnessHostBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VSANVcSetPreferredFaultDomainBody struct {
	Req    *types.VSANVcSetPreferredFaultDomain         `xml:"urn:vsan VSANVcSetPreferredFaultDomain,omitempty"`
	Res    *types.VSANVcSetPreferredFaultDomainResponse `xml:"urn:vsan VSANVcSetPreferredFaultDomainResponse,omitempty"`
	Fault_ *soap.Fault                                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VSANVcSetPreferredFaultDomainBody) Fault() *soap.Fault { return b.Fault_ }

func VSANVcSetPreferredFaultDomain(ctx context.Context, r soap.RoundTripper, req *types.VSANVcSetPreferredFaultDomain) (*types.VSANVcSetPreferredFaultDomainResponse, error) {
	var reqBody, resBody VSANVcSetPreferredFaultDomainBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VSANVcReconcileUnicastAgentsBody struct {
	Req    *types.VSANVcReconcileUnicastAgents         `xml:"urn:vsan VSANVcReconcileUnicastAgents,omitempty"`
	Res    *types.VSANVcReconcileUnicastAgentsResponse `xml:"urn:vsan VSANVcReconcileUnicastAgentsResponse,omitempty"`
	Fault_ *soap.Fault                                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VSANVcReconcileUnicastAgentsBody) Fault() *soap.Fault { return b.Fault_ }

func VSANVcReconcileUnicastAgents(ctx context.Context, r soap.RoundTripper, req *types.VSANVcReconcileUnicastAgents) (*types.VSANVcReconcileUnicastAgentsResponse, error) {
	var reqBody, resBody VSANVcReconcileUnicastAgentsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VSANVcGetPreferredFaultDomainBody struct {
	Req    *types.VSANVcGetPreferredFaultDomain         `xml:"urn:vsan VSANVcGetPreferredFaultDomain,omitempty"`
	Res    *types.VSANVcGetPreferredFaultDomainResponse `xml:"urn:vsan VSANVcGetPreferredFaultDomainResponse,omitempty"`
	Fault_ *soap.Fault                                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VSANVcGetPreferredFaultDomainBody) Fault() *soap.Fault { return b.Fault_ }

func VSANVcGetPreferredFaultDomain(ctx context.Context, r soap.RoundTripper, req *types.VSANVcGetPreferredFaultDomain) (*types.VSANVcGetPreferredFaultDomainResponse, error) {
	var reqBody, resBody VSANVcGetPreferredFaultDomainBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VSANIsWitnessVirtualApplianceBody struct {
	Req    *types.VSANIsWitnessVirtualAppliance         `xml:"urn:vsan VSANIsWitnessVirtualAppliance,omitempty"`
	Res    *types.VSANIsWitnessVirtualApplianceResponse `xml:"urn:vsan VSANIsWitnessVirtualApplianceResponse,omitempty"`
	Fault_ *soap.Fault                                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VSANIsWitnessVirtualApplianceBody) Fault() *soap.Fault { return b.Fault_ }

func VSANIsWitnessVirtualAppliance(ctx context.Context, r soap.RoundTripper, req *types.VSANIsWitnessVirtualAppliance) (*types.VSANIsWitnessVirtualApplianceResponse, error) {
	var reqBody, resBody VSANIsWitnessVirtualApplianceBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VSANVcAddWitnessHostBody struct {
	Req    *types.VSANVcAddWitnessHost         `xml:"urn:vsan VSANVcAddWitnessHost,omitempty"`
	Res    *types.VSANVcAddWitnessHostResponse `xml:"urn:vsan VSANVcAddWitnessHostResponse,omitempty"`
	Fault_ *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VSANVcAddWitnessHostBody) Fault() *soap.Fault { return b.Fault_ }

func VSANVcAddWitnessHost(ctx context.Context, r soap.RoundTripper, req *types.VSANVcAddWitnessHost) (*types.VSANVcAddWitnessHostResponse, error) {
	var reqBody, resBody VSANVcAddWitnessHostBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VSANVcRetrieveStretchedClusterHostCapabilityBody struct {
	Req    *types.VSANVcRetrieveStretchedClusterHostCapability         `xml:"urn:vsan VSANVcRetrieveStretchedClusterHostCapability,omitempty"`
	Res    *types.VSANVcRetrieveStretchedClusterHostCapabilityResponse `xml:"urn:vsan VSANVcRetrieveStretchedClusterHostCapabilityResponse,omitempty"`
	Fault_ *soap.Fault                                                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VSANVcRetrieveStretchedClusterHostCapabilityBody) Fault() *soap.Fault { return b.Fault_ }

func VSANVcRetrieveStretchedClusterHostCapability(ctx context.Context, r soap.RoundTripper, req *types.VSANVcRetrieveStretchedClusterHostCapability) (*types.VSANVcRetrieveStretchedClusterHostCapabilityResponse, error) {
	var reqBody, resBody VSANVcRetrieveStretchedClusterHostCapabilityBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VSANVcGetNumOfWitnessHostsBody struct {
	Req    *types.VSANVcGetNumOfWitnessHosts         `xml:"urn:vsan VSANVcGetNumOfWitnessHosts,omitempty"`
	Res    *types.VSANVcGetNumOfWitnessHostsResponse `xml:"urn:vsan VSANVcGetNumOfWitnessHostsResponse,omitempty"`
	Fault_ *soap.Fault                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VSANVcGetNumOfWitnessHostsBody) Fault() *soap.Fault { return b.Fault_ }

func VSANVcGetNumOfWitnessHosts(ctx context.Context, r soap.RoundTripper, req *types.VSANVcGetNumOfWitnessHosts) (*types.VSANVcGetNumOfWitnessHostsResponse, error) {
	var reqBody, resBody VSANVcGetNumOfWitnessHostsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VSANVcGetStretchedClusterConfigIssuesBody struct {
	Req    *types.VSANVcGetStretchedClusterConfigIssues         `xml:"urn:vsan VSANVcGetStretchedClusterConfigIssues,omitempty"`
	Res    *types.VSANVcGetStretchedClusterConfigIssuesResponse `xml:"urn:vsan VSANVcGetStretchedClusterConfigIssuesResponse,omitempty"`
	Fault_ *soap.Fault                                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VSANVcGetStretchedClusterConfigIssuesBody) Fault() *soap.Fault { return b.Fault_ }

func VSANVcGetStretchedClusterConfigIssues(ctx context.Context, r soap.RoundTripper, req *types.VSANVcGetStretchedClusterConfigIssues) (*types.VSANVcGetStretchedClusterConfigIssuesResponse, error) {
	var reqBody, resBody VSANVcGetStretchedClusterConfigIssuesBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VSANVcGetWitnessHostsBody struct {
	Req    *types.VSANVcGetWitnessHosts         `xml:"urn:vsan VSANVcGetWitnessHosts,omitempty"`
	Res    *types.VSANVcGetWitnessHostsResponse `xml:"urn:vsan VSANVcGetWitnessHostsResponse,omitempty"`
	Fault_ *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VSANVcGetWitnessHostsBody) Fault() *soap.Fault { return b.Fault_ }

func VSANVcGetWitnessHosts(ctx context.Context, r soap.RoundTripper, req *types.VSANVcGetWitnessHosts) (*types.VSANVcGetWitnessHostsResponse, error) {
	var reqBody, resBody VSANVcGetWitnessHostsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VSANVcRetrieveStretchedClusterVcCapabilityBody struct {
	Req    *types.VSANVcRetrieveStretchedClusterVcCapability         `xml:"urn:vsan VSANVcRetrieveStretchedClusterVcCapability,omitempty"`
	Res    *types.VSANVcRetrieveStretchedClusterVcCapabilityResponse `xml:"urn:vsan VSANVcRetrieveStretchedClusterVcCapabilityResponse,omitempty"`
	Fault_ *soap.Fault                                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VSANVcRetrieveStretchedClusterVcCapabilityBody) Fault() *soap.Fault { return b.Fault_ }

func VSANVcRetrieveStretchedClusterVcCapability(ctx context.Context, r soap.RoundTripper, req *types.VSANVcRetrieveStretchedClusterVcCapability) (*types.VSANVcRetrieveStretchedClusterVcCapabilityResponse, error) {
	var reqBody, resBody VSANVcRetrieveStretchedClusterVcCapabilityBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VSANVcConvertToStretchedClusterBody struct {
	Req    *types.VSANVcConvertToStretchedCluster         `xml:"urn:vsan VSANVcConvertToStretchedCluster,omitempty"`
	Res    *types.VSANVcConvertToStretchedClusterResponse `xml:"urn:vsan VSANVcConvertToStretchedClusterResponse,omitempty"`
	Fault_ *soap.Fault                                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VSANVcConvertToStretchedClusterBody) Fault() *soap.Fault { return b.Fault_ }

func VSANVcConvertToStretchedCluster(ctx context.Context, r soap.RoundTripper, req *types.VSANVcConvertToStretchedCluster) (*types.VSANVcConvertToStretchedClusterResponse, error) {
	var reqBody, resBody VSANVcConvertToStretchedClusterBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VSANVcRemoveWitnessHostBody struct {
	Req    *types.VSANVcRemoveWitnessHost         `xml:"urn:vsan VSANVcRemoveWitnessHost,omitempty"`
	Res    *types.VSANVcRemoveWitnessHostResponse `xml:"urn:vsan VSANVcRemoveWitnessHostResponse,omitempty"`
	Fault_ *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VSANVcRemoveWitnessHostBody) Fault() *soap.Fault { return b.Fault_ }

func VSANVcRemoveWitnessHost(ctx context.Context, r soap.RoundTripper, req *types.VSANVcRemoveWitnessHost) (*types.VSANVcRemoveWitnessHostResponse, error) {
	var reqBody, resBody VSANVcRemoveWitnessHostBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanQueryClusterPhysicalDiskHealthSummaryBody struct {
	Req    *types.VsanQueryClusterPhysicalDiskHealthSummary         `xml:"urn:vsan VsanQueryClusterPhysicalDiskHealthSummary,omitempty"`
	Res    *types.VsanQueryClusterPhysicalDiskHealthSummaryResponse `xml:"urn:vsan VsanQueryClusterPhysicalDiskHealthSummaryResponse,omitempty"`
	Fault_ *soap.Fault                                              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanQueryClusterPhysicalDiskHealthSummaryBody) Fault() *soap.Fault { return b.Fault_ }

func VsanQueryClusterPhysicalDiskHealthSummary(ctx context.Context, r soap.RoundTripper, req *types.VsanQueryClusterPhysicalDiskHealthSummary) (*types.VsanQueryClusterPhysicalDiskHealthSummaryResponse, error) {
	var reqBody, resBody VsanQueryClusterPhysicalDiskHealthSummaryBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanQueryClusterNetworkPerfTestBody struct {
	Req    *types.VsanQueryClusterNetworkPerfTest         `xml:"urn:vsan VsanQueryClusterNetworkPerfTest,omitempty"`
	Res    *types.VsanQueryClusterNetworkPerfTestResponse `xml:"urn:vsan VsanQueryClusterNetworkPerfTestResponse,omitempty"`
	Fault_ *soap.Fault                                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanQueryClusterNetworkPerfTestBody) Fault() *soap.Fault { return b.Fault_ }

func VsanQueryClusterNetworkPerfTest(ctx context.Context, r soap.RoundTripper, req *types.VsanQueryClusterNetworkPerfTest) (*types.VsanQueryClusterNetworkPerfTestResponse, error) {
	var reqBody, resBody VsanQueryClusterNetworkPerfTestBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanQueryClusterAdvCfgSyncBody struct {
	Req    *types.VsanQueryClusterAdvCfgSync         `xml:"urn:vsan VsanQueryClusterAdvCfgSync,omitempty"`
	Res    *types.VsanQueryClusterAdvCfgSyncResponse `xml:"urn:vsan VsanQueryClusterAdvCfgSyncResponse,omitempty"`
	Fault_ *soap.Fault                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanQueryClusterAdvCfgSyncBody) Fault() *soap.Fault { return b.Fault_ }

func VsanQueryClusterAdvCfgSync(ctx context.Context, r soap.RoundTripper, req *types.VsanQueryClusterAdvCfgSync) (*types.VsanQueryClusterAdvCfgSyncResponse, error) {
	var reqBody, resBody VsanQueryClusterAdvCfgSyncBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanCheckClusterArchivalAccessibilityBody struct {
	Req    *types.VsanCheckClusterArchivalAccessibility         `xml:"urn:vsan VsanCheckClusterArchivalAccessibility,omitempty"`
	Res    *types.VsanCheckClusterArchivalAccessibilityResponse `xml:"urn:vsan VsanCheckClusterArchivalAccessibilityResponse,omitempty"`
	Fault_ *soap.Fault                                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanCheckClusterArchivalAccessibilityBody) Fault() *soap.Fault { return b.Fault_ }

func VsanCheckClusterArchivalAccessibility(ctx context.Context, r soap.RoundTripper, req *types.VsanCheckClusterArchivalAccessibility) (*types.VsanCheckClusterArchivalAccessibilityResponse, error) {
	var reqBody, resBody VsanCheckClusterArchivalAccessibilityBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanCheckClusterDpdLivenessBody struct {
	Req    *types.VsanCheckClusterDpdLiveness         `xml:"urn:vsan VsanCheckClusterDpdLiveness,omitempty"`
	Res    *types.VsanCheckClusterDpdLivenessResponse `xml:"urn:vsan VsanCheckClusterDpdLivenessResponse,omitempty"`
	Fault_ *soap.Fault                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanCheckClusterDpdLivenessBody) Fault() *soap.Fault { return b.Fault_ }

func VsanCheckClusterDpdLiveness(ctx context.Context, r soap.RoundTripper, req *types.VsanCheckClusterDpdLiveness) (*types.VsanCheckClusterDpdLivenessResponse, error) {
	var reqBody, resBody VsanCheckClusterDpdLivenessBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanRepairClusterImmediateObjectsBody struct {
	Req    *types.VsanRepairClusterImmediateObjects         `xml:"urn:vsan VsanRepairClusterImmediateObjects,omitempty"`
	Res    *types.VsanRepairClusterImmediateObjectsResponse `xml:"urn:vsan VsanRepairClusterImmediateObjectsResponse,omitempty"`
	Fault_ *soap.Fault                                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanRepairClusterImmediateObjectsBody) Fault() *soap.Fault { return b.Fault_ }

func VsanRepairClusterImmediateObjects(ctx context.Context, r soap.RoundTripper, req *types.VsanRepairClusterImmediateObjects) (*types.VsanRepairClusterImmediateObjectsResponse, error) {
	var reqBody, resBody VsanRepairClusterImmediateObjectsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanQueryVerifyClusterNetworkSettingsBody struct {
	Req    *types.VsanQueryVerifyClusterNetworkSettings         `xml:"urn:vsan VsanQueryVerifyClusterNetworkSettings,omitempty"`
	Res    *types.VsanQueryVerifyClusterNetworkSettingsResponse `xml:"urn:vsan VsanQueryVerifyClusterNetworkSettingsResponse,omitempty"`
	Fault_ *soap.Fault                                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanQueryVerifyClusterNetworkSettingsBody) Fault() *soap.Fault { return b.Fault_ }

func VsanQueryVerifyClusterNetworkSettings(ctx context.Context, r soap.RoundTripper, req *types.VsanQueryVerifyClusterNetworkSettings) (*types.VsanQueryVerifyClusterNetworkSettingsResponse, error) {
	var reqBody, resBody VsanQueryVerifyClusterNetworkSettingsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanQueryClusterCreateVmHealthTestBody struct {
	Req    *types.VsanQueryClusterCreateVmHealthTest         `xml:"urn:vsan VsanQueryClusterCreateVmHealthTest,omitempty"`
	Res    *types.VsanQueryClusterCreateVmHealthTestResponse `xml:"urn:vsan VsanQueryClusterCreateVmHealthTestResponse,omitempty"`
	Fault_ *soap.Fault                                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanQueryClusterCreateVmHealthTestBody) Fault() *soap.Fault { return b.Fault_ }

func VsanQueryClusterCreateVmHealthTest(ctx context.Context, r soap.RoundTripper, req *types.VsanQueryClusterCreateVmHealthTest) (*types.VsanQueryClusterCreateVmHealthTestResponse, error) {
	var reqBody, resBody VsanQueryClusterCreateVmHealthTestBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanQueryClusterHealthSystemVersionsBody struct {
	Req    *types.VsanQueryClusterHealthSystemVersions         `xml:"urn:vsan VsanQueryClusterHealthSystemVersions,omitempty"`
	Res    *types.VsanQueryClusterHealthSystemVersionsResponse `xml:"urn:vsan VsanQueryClusterHealthSystemVersionsResponse,omitempty"`
	Fault_ *soap.Fault                                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanQueryClusterHealthSystemVersionsBody) Fault() *soap.Fault { return b.Fault_ }

func VsanQueryClusterHealthSystemVersions(ctx context.Context, r soap.RoundTripper, req *types.VsanQueryClusterHealthSystemVersions) (*types.VsanQueryClusterHealthSystemVersionsResponse, error) {
	var reqBody, resBody VsanQueryClusterHealthSystemVersionsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanClusterGetHclInfoBody struct {
	Req    *types.VsanClusterGetHclInfo         `xml:"urn:vsan VsanClusterGetHclInfo,omitempty"`
	Res    *types.VsanClusterGetHclInfoResponse `xml:"urn:vsan VsanClusterGetHclInfoResponse,omitempty"`
	Fault_ *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanClusterGetHclInfoBody) Fault() *soap.Fault { return b.Fault_ }

func VsanClusterGetHclInfo(ctx context.Context, r soap.RoundTripper, req *types.VsanClusterGetHclInfo) (*types.VsanClusterGetHclInfoResponse, error) {
	var reqBody, resBody VsanClusterGetHclInfoBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanQueryClusterCheckLimitsBody struct {
	Req    *types.VsanQueryClusterCheckLimits         `xml:"urn:vsan VsanQueryClusterCheckLimits,omitempty"`
	Res    *types.VsanQueryClusterCheckLimitsResponse `xml:"urn:vsan VsanQueryClusterCheckLimitsResponse,omitempty"`
	Fault_ *soap.Fault                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanQueryClusterCheckLimitsBody) Fault() *soap.Fault { return b.Fault_ }

func VsanQueryClusterCheckLimits(ctx context.Context, r soap.RoundTripper, req *types.VsanQueryClusterCheckLimits) (*types.VsanQueryClusterCheckLimitsResponse, error) {
	var reqBody, resBody VsanQueryClusterCheckLimitsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanQueryClusterCaptureVsanPcapBody struct {
	Req    *types.VsanQueryClusterCaptureVsanPcap         `xml:"urn:vsan VsanQueryClusterCaptureVsanPcap,omitempty"`
	Res    *types.VsanQueryClusterCaptureVsanPcapResponse `xml:"urn:vsan VsanQueryClusterCaptureVsanPcapResponse,omitempty"`
	Fault_ *soap.Fault                                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanQueryClusterCaptureVsanPcapBody) Fault() *soap.Fault { return b.Fault_ }

func VsanQueryClusterCaptureVsanPcap(ctx context.Context, r soap.RoundTripper, req *types.VsanQueryClusterCaptureVsanPcap) (*types.VsanQueryClusterCaptureVsanPcapResponse, error) {
	var reqBody, resBody VsanQueryClusterCaptureVsanPcapBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanCheckClusterClomdLivenessBody struct {
	Req    *types.VsanCheckClusterClomdLiveness         `xml:"urn:vsan VsanCheckClusterClomdLiveness,omitempty"`
	Res    *types.VsanCheckClusterClomdLivenessResponse `xml:"urn:vsan VsanCheckClusterClomdLivenessResponse,omitempty"`
	Fault_ *soap.Fault                                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanCheckClusterClomdLivenessBody) Fault() *soap.Fault { return b.Fault_ }

func VsanCheckClusterClomdLiveness(ctx context.Context, r soap.RoundTripper, req *types.VsanCheckClusterClomdLiveness) (*types.VsanCheckClusterClomdLivenessResponse, error) {
	var reqBody, resBody VsanCheckClusterClomdLivenessBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanQueryMobStatusBody struct {
	Req    *types.VsanQueryMobStatus         `xml:"urn:vsan VsanQueryMobStatus,omitempty"`
	Res    *types.VsanQueryMobStatusResponse `xml:"urn:vsan VsanQueryMobStatusResponse,omitempty"`
	Fault_ *soap.Fault                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanQueryMobStatusBody) Fault() *soap.Fault { return b.Fault_ }

func VsanQueryMobStatus(ctx context.Context, r soap.RoundTripper, req *types.VsanQueryMobStatus) (*types.VsanQueryMobStatusResponse, error) {
	var reqBody, resBody VsanQueryMobStatusBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanQueryMemoryStatsBody struct {
	Req    *types.VsanQueryMemoryStats         `xml:"urn:vsan VsanQueryMemoryStats,omitempty"`
	Res    *types.VsanQueryMemoryStatsResponse `xml:"urn:vsan VsanQueryMemoryStatsResponse,omitempty"`
	Fault_ *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanQueryMemoryStatsBody) Fault() *soap.Fault { return b.Fault_ }

func VsanQueryMemoryStats(ctx context.Context, r soap.RoundTripper, req *types.VsanQueryMemoryStats) (*types.VsanQueryMemoryStatsResponse, error) {
	var reqBody, resBody VsanQueryMemoryStatsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanStopMobServiceBody struct {
	Req    *types.VsanStopMobService         `xml:"urn:vsan VsanStopMobService,omitempty"`
	Res    *types.VsanStopMobServiceResponse `xml:"urn:vsan VsanStopMobServiceResponse,omitempty"`
	Fault_ *soap.Fault                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanStopMobServiceBody) Fault() *soap.Fault { return b.Fault_ }

func VsanStopMobService(ctx context.Context, r soap.RoundTripper, req *types.VsanStopMobService) (*types.VsanStopMobServiceResponse, error) {
	var reqBody, resBody VsanStopMobServiceBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanStartMobServiceBody struct {
	Req    *types.VsanStartMobService         `xml:"urn:vsan VsanStartMobService,omitempty"`
	Res    *types.VsanStartMobServiceResponse `xml:"urn:vsan VsanStartMobServiceResponse,omitempty"`
	Fault_ *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanStartMobServiceBody) Fault() *soap.Fault { return b.Fault_ }

func VsanStartMobService(ctx context.Context, r soap.RoundTripper, req *types.VsanStartMobService) (*types.VsanStartMobServiceResponse, error) {
	var reqBody, resBody VsanStartMobServiceBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanFlrServiceGetStateBody struct {
	Req    *types.VsanFlrServiceGetState         `xml:"urn:vsan VsanFlrServiceGetState,omitempty"`
	Res    *types.VsanFlrServiceGetStateResponse `xml:"urn:vsan VsanFlrServiceGetStateResponse,omitempty"`
	Fault_ *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanFlrServiceGetStateBody) Fault() *soap.Fault { return b.Fault_ }

func VsanFlrServiceGetState(ctx context.Context, r soap.RoundTripper, req *types.VsanFlrServiceGetState) (*types.VsanFlrServiceGetStateResponse, error) {
	var reqBody, resBody VsanFlrServiceGetStateBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanFlrServiceDestroySessionBody struct {
	Req    *types.VsanFlrServiceDestroySession         `xml:"urn:vsan VsanFlrServiceDestroySession,omitempty"`
	Res    *types.VsanFlrServiceDestroySessionResponse `xml:"urn:vsan VsanFlrServiceDestroySessionResponse,omitempty"`
	Fault_ *soap.Fault                                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanFlrServiceDestroySessionBody) Fault() *soap.Fault { return b.Fault_ }

func VsanFlrServiceDestroySession(ctx context.Context, r soap.RoundTripper, req *types.VsanFlrServiceDestroySession) (*types.VsanFlrServiceDestroySessionResponse, error) {
	var reqBody, resBody VsanFlrServiceDestroySessionBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanFlrServiceCreateSessionBody struct {
	Req    *types.VsanFlrServiceCreateSession         `xml:"urn:vsan VsanFlrServiceCreateSession,omitempty"`
	Res    *types.VsanFlrServiceCreateSessionResponse `xml:"urn:vsan VsanFlrServiceCreateSessionResponse,omitempty"`
	Fault_ *soap.Fault                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanFlrServiceCreateSessionBody) Fault() *soap.Fault { return b.Fault_ }

func VsanFlrServiceCreateSession(ctx context.Context, r soap.RoundTripper, req *types.VsanFlrServiceCreateSession) (*types.VsanFlrServiceCreateSessionResponse, error) {
	var reqBody, resBody VsanFlrServiceCreateSessionBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanGetClusterPhoneHomeDataBody struct {
	Req    *types.VsanGetClusterPhoneHomeData         `xml:"urn:vsan VsanGetClusterPhoneHomeData,omitempty"`
	Res    *types.VsanGetClusterPhoneHomeDataResponse `xml:"urn:vsan VsanGetClusterPhoneHomeDataResponse,omitempty"`
	Fault_ *soap.Fault                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanGetClusterPhoneHomeDataBody) Fault() *soap.Fault { return b.Fault_ }

func VsanGetClusterPhoneHomeData(ctx context.Context, r soap.RoundTripper, req *types.VsanGetClusterPhoneHomeData) (*types.VsanGetClusterPhoneHomeDataResponse, error) {
	var reqBody, resBody VsanGetClusterPhoneHomeDataBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanQueryClusterHostLogBundleBody struct {
	Req    *types.VsanQueryClusterHostLogBundle         `xml:"urn:vsan VsanQueryClusterHostLogBundle,omitempty"`
	Res    *types.VsanQueryClusterHostLogBundleResponse `xml:"urn:vsan VsanQueryClusterHostLogBundleResponse,omitempty"`
	Fault_ *soap.Fault                                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanQueryClusterHostLogBundleBody) Fault() *soap.Fault { return b.Fault_ }

func VsanQueryClusterHostLogBundle(ctx context.Context, r soap.RoundTripper, req *types.VsanQueryClusterHostLogBundle) (*types.VsanQueryClusterHostLogBundleResponse, error) {
	var reqBody, resBody VsanQueryClusterHostLogBundleBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanPhoneHomeStatusBody struct {
	Req    *types.VsanPhoneHomeStatus         `xml:"urn:vsan VsanPhoneHomeStatus,omitempty"`
	Res    *types.VsanPhoneHomeStatusResponse `xml:"urn:vsan VsanPhoneHomeStatusResponse,omitempty"`
	Fault_ *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanPhoneHomeStatusBody) Fault() *soap.Fault { return b.Fault_ }

func VsanPhoneHomeStatus(ctx context.Context, r soap.RoundTripper, req *types.VsanPhoneHomeStatus) (*types.VsanPhoneHomeStatusResponse, error) {
	var reqBody, resBody VsanPhoneHomeStatusBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanVcQueryClusterHostLogBundleBody struct {
	Req    *types.VsanVcQueryClusterHostLogBundle         `xml:"urn:vsan VsanVcQueryClusterHostLogBundle,omitempty"`
	Res    *types.VsanVcQueryClusterHostLogBundleResponse `xml:"urn:vsan VsanVcQueryClusterHostLogBundleResponse,omitempty"`
	Fault_ *soap.Fault                                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanVcQueryClusterHostLogBundleBody) Fault() *soap.Fault { return b.Fault_ }

func VsanVcQueryClusterHostLogBundle(ctx context.Context, r soap.RoundTripper, req *types.VsanVcQueryClusterHostLogBundle) (*types.VsanVcQueryClusterHostLogBundleResponse, error) {
	var reqBody, resBody VsanVcQueryClusterHostLogBundleBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanVcSetClusterLogsCatalogBody struct {
	Req    *types.VsanVcSetClusterLogsCatalog         `xml:"urn:vsan VsanVcSetClusterLogsCatalog,omitempty"`
	Res    *types.VsanVcSetClusterLogsCatalogResponse `xml:"urn:vsan VsanVcSetClusterLogsCatalogResponse,omitempty"`
	Fault_ *soap.Fault                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanVcSetClusterLogsCatalogBody) Fault() *soap.Fault { return b.Fault_ }

func VsanVcSetClusterLogsCatalog(ctx context.Context, r soap.RoundTripper, req *types.VsanVcSetClusterLogsCatalog) (*types.VsanVcSetClusterLogsCatalogResponse, error) {
	var reqBody, resBody VsanVcSetClusterLogsCatalogBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanTriggerDiagnosticCollectionBody struct {
	Req    *types.VsanTriggerDiagnosticCollection         `xml:"urn:vsan VsanTriggerDiagnosticCollection,omitempty"`
	Res    *types.VsanTriggerDiagnosticCollectionResponse `xml:"urn:vsan VsanTriggerDiagnosticCollectionResponse,omitempty"`
	Fault_ *soap.Fault                                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanTriggerDiagnosticCollectionBody) Fault() *soap.Fault { return b.Fault_ }

func VsanTriggerDiagnosticCollection(ctx context.Context, r soap.RoundTripper, req *types.VsanTriggerDiagnosticCollection) (*types.VsanTriggerDiagnosticCollectionResponse, error) {
	var reqBody, resBody VsanTriggerDiagnosticCollectionBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryVsanCloudHealthDiagnosticDataBody struct {
	Req    *types.QueryVsanCloudHealthDiagnosticData         `xml:"urn:vsan QueryVsanCloudHealthDiagnosticData,omitempty"`
	Res    *types.QueryVsanCloudHealthDiagnosticDataResponse `xml:"urn:vsan QueryVsanCloudHealthDiagnosticDataResponse,omitempty"`
	Fault_ *soap.Fault                                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryVsanCloudHealthDiagnosticDataBody) Fault() *soap.Fault { return b.Fault_ }

func QueryVsanCloudHealthDiagnosticData(ctx context.Context, r soap.RoundTripper, req *types.QueryVsanCloudHealthDiagnosticData) (*types.QueryVsanCloudHealthDiagnosticDataResponse, error) {
	var reqBody, resBody QueryVsanCloudHealthDiagnosticDataBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanGetCoreDumpPartitionInfoBody struct {
	Req    *types.VsanGetCoreDumpPartitionInfo         `xml:"urn:vsan VsanGetCoreDumpPartitionInfo,omitempty"`
	Res    *types.VsanGetCoreDumpPartitionInfoResponse `xml:"urn:vsan VsanGetCoreDumpPartitionInfoResponse,omitempty"`
	Fault_ *soap.Fault                                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanGetCoreDumpPartitionInfoBody) Fault() *soap.Fault { return b.Fault_ }

func VsanGetCoreDumpPartitionInfo(ctx context.Context, r soap.RoundTripper, req *types.VsanGetCoreDumpPartitionInfo) (*types.VsanGetCoreDumpPartitionInfoResponse, error) {
	var reqBody, resBody VsanGetCoreDumpPartitionInfoBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryVsanCloudHealthStatusBody struct {
	Req    *types.QueryVsanCloudHealthStatus         `xml:"urn:vsan QueryVsanCloudHealthStatus,omitempty"`
	Res    *types.QueryVsanCloudHealthStatusResponse `xml:"urn:vsan QueryVsanCloudHealthStatusResponse,omitempty"`
	Fault_ *soap.Fault                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryVsanCloudHealthStatusBody) Fault() *soap.Fault { return b.Fault_ }

func QueryVsanCloudHealthStatus(ctx context.Context, r soap.RoundTripper, req *types.QueryVsanCloudHealthStatus) (*types.QueryVsanCloudHealthStatusResponse, error) {
	var reqBody, resBody QueryVsanCloudHealthStatusBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanQueryCMMDSPhoneHomeDataBody struct {
	Req    *types.VsanQueryCMMDSPhoneHomeData         `xml:"urn:vsan VsanQueryCMMDSPhoneHomeData,omitempty"`
	Res    *types.VsanQueryCMMDSPhoneHomeDataResponse `xml:"urn:vsan VsanQueryCMMDSPhoneHomeDataResponse,omitempty"`
	Fault_ *soap.Fault                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanQueryCMMDSPhoneHomeDataBody) Fault() *soap.Fault { return b.Fault_ }

func VsanQueryCMMDSPhoneHomeData(ctx context.Context, r soap.RoundTripper, req *types.VsanQueryCMMDSPhoneHomeData) (*types.VsanQueryCMMDSPhoneHomeDataResponse, error) {
	var reqBody, resBody VsanQueryCMMDSPhoneHomeDataBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanQueryVcLogBundleBody struct {
	Req    *types.VsanQueryVcLogBundle         `xml:"urn:vsan VsanQueryVcLogBundle,omitempty"`
	Res    *types.VsanQueryVcLogBundleResponse `xml:"urn:vsan VsanQueryVcLogBundleResponse,omitempty"`
	Fault_ *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanQueryVcLogBundleBody) Fault() *soap.Fault { return b.Fault_ }

func VsanQueryVcLogBundle(ctx context.Context, r soap.RoundTripper, req *types.VsanQueryVcLogBundle) (*types.VsanQueryVcLogBundleResponse, error) {
	var reqBody, resBody VsanQueryVcLogBundleBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanSetLastHostBundleCollectionTimeBody struct {
	Req    *types.VsanSetLastHostBundleCollectionTime         `xml:"urn:vsan VsanSetLastHostBundleCollectionTime,omitempty"`
	Res    *types.VsanSetLastHostBundleCollectionTimeResponse `xml:"urn:vsan VsanSetLastHostBundleCollectionTimeResponse,omitempty"`
	Fault_ *soap.Fault                                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanSetLastHostBundleCollectionTimeBody) Fault() *soap.Fault { return b.Fault_ }

func VsanSetLastHostBundleCollectionTime(ctx context.Context, r soap.RoundTripper, req *types.VsanSetLastHostBundleCollectionTime) (*types.VsanSetLastHostBundleCollectionTimeResponse, error) {
	var reqBody, resBody VsanSetLastHostBundleCollectionTimeBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanGetPhoneHomeObfuscationMapBody struct {
	Req    *types.VsanGetPhoneHomeObfuscationMap         `xml:"urn:vsan VsanGetPhoneHomeObfuscationMap,omitempty"`
	Res    *types.VsanGetPhoneHomeObfuscationMapResponse `xml:"urn:vsan VsanGetPhoneHomeObfuscationMapResponse,omitempty"`
	Fault_ *soap.Fault                                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanGetPhoneHomeObfuscationMapBody) Fault() *soap.Fault { return b.Fault_ }

func VsanGetPhoneHomeObfuscationMap(ctx context.Context, r soap.RoundTripper, req *types.VsanGetPhoneHomeObfuscationMap) (*types.VsanGetPhoneHomeObfuscationMapResponse, error) {
	var reqBody, resBody VsanGetPhoneHomeObfuscationMapBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanEnablePerfNetworkDiagnosticResultsBody struct {
	Req    *types.VsanEnablePerfNetworkDiagnosticResults         `xml:"urn:vsan VsanEnablePerfNetworkDiagnosticResults,omitempty"`
	Res    *types.VsanEnablePerfNetworkDiagnosticResultsResponse `xml:"urn:vsan VsanEnablePerfNetworkDiagnosticResultsResponse,omitempty"`
	Fault_ *soap.Fault                                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanEnablePerfNetworkDiagnosticResultsBody) Fault() *soap.Fault { return b.Fault_ }

func VsanEnablePerfNetworkDiagnosticResults(ctx context.Context, r soap.RoundTripper, req *types.VsanEnablePerfNetworkDiagnosticResults) (*types.VsanEnablePerfNetworkDiagnosticResultsResponse, error) {
	var reqBody, resBody VsanEnablePerfNetworkDiagnosticResultsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanPhoneHomeQueryPhysicalVsanDisksBody struct {
	Req    *types.VsanPhoneHomeQueryPhysicalVsanDisks         `xml:"urn:vsan VsanPhoneHomeQueryPhysicalVsanDisks,omitempty"`
	Res    *types.VsanPhoneHomeQueryPhysicalVsanDisksResponse `xml:"urn:vsan VsanPhoneHomeQueryPhysicalVsanDisksResponse,omitempty"`
	Fault_ *soap.Fault                                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanPhoneHomeQueryPhysicalVsanDisksBody) Fault() *soap.Fault { return b.Fault_ }

func VsanPhoneHomeQueryPhysicalVsanDisks(ctx context.Context, r soap.RoundTripper, req *types.VsanPhoneHomeQueryPhysicalVsanDisks) (*types.VsanPhoneHomeQueryPhysicalVsanDisksResponse, error) {
	var reqBody, resBody VsanPhoneHomeQueryPhysicalVsanDisksBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanCountLogTraceLinesBody struct {
	Req    *types.VsanCountLogTraceLines         `xml:"urn:vsan VsanCountLogTraceLines,omitempty"`
	Res    *types.VsanCountLogTraceLinesResponse `xml:"urn:vsan VsanCountLogTraceLinesResponse,omitempty"`
	Fault_ *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanCountLogTraceLinesBody) Fault() *soap.Fault { return b.Fault_ }

func VsanCountLogTraceLines(ctx context.Context, r soap.RoundTripper, req *types.VsanCountLogTraceLines) (*types.VsanCountLogTraceLinesResponse, error) {
	var reqBody, resBody VsanCountLogTraceLinesBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanPerformOnlineHealthCheckBody struct {
	Req    *types.VsanPerformOnlineHealthCheck         `xml:"urn:vsan VsanPerformOnlineHealthCheck,omitempty"`
	Res    *types.VsanPerformOnlineHealthCheckResponse `xml:"urn:vsan VsanPerformOnlineHealthCheckResponse,omitempty"`
	Fault_ *soap.Fault                                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanPerformOnlineHealthCheckBody) Fault() *soap.Fault { return b.Fault_ }

func VsanPerformOnlineHealthCheck(ctx context.Context, r soap.RoundTripper, req *types.VsanPerformOnlineHealthCheck) (*types.VsanPerformOnlineHealthCheckResponse, error) {
	var reqBody, resBody VsanPerformOnlineHealthCheckBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanPhoneHomeGetEsxCliDataBody struct {
	Req    *types.VsanPhoneHomeGetEsxCliData         `xml:"urn:vsan VsanPhoneHomeGetEsxCliData,omitempty"`
	Res    *types.VsanPhoneHomeGetEsxCliDataResponse `xml:"urn:vsan VsanPhoneHomeGetEsxCliDataResponse,omitempty"`
	Fault_ *soap.Fault                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanPhoneHomeGetEsxCliDataBody) Fault() *soap.Fault { return b.Fault_ }

func VsanPhoneHomeGetEsxCliData(ctx context.Context, r soap.RoundTripper, req *types.VsanPhoneHomeGetEsxCliData) (*types.VsanPhoneHomeGetEsxCliDataResponse, error) {
	var reqBody, resBody VsanPhoneHomeGetEsxCliDataBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryVsanCloudHealthClusterHostsProfilingDataBody struct {
	Req    *types.QueryVsanCloudHealthClusterHostsProfilingData         `xml:"urn:vsan QueryVsanCloudHealthClusterHostsProfilingData,omitempty"`
	Res    *types.QueryVsanCloudHealthClusterHostsProfilingDataResponse `xml:"urn:vsan QueryVsanCloudHealthClusterHostsProfilingDataResponse,omitempty"`
	Fault_ *soap.Fault                                                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryVsanCloudHealthClusterHostsProfilingDataBody) Fault() *soap.Fault { return b.Fault_ }

func QueryVsanCloudHealthClusterHostsProfilingData(ctx context.Context, r soap.RoundTripper, req *types.QueryVsanCloudHealthClusterHostsProfilingData) (*types.QueryVsanCloudHealthClusterHostsProfilingDataResponse, error) {
	var reqBody, resBody QueryVsanCloudHealthClusterHostsProfilingDataBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanDisablePerfNetworkDiagnosticResultsBody struct {
	Req    *types.VsanDisablePerfNetworkDiagnosticResults         `xml:"urn:vsan VsanDisablePerfNetworkDiagnosticResults,omitempty"`
	Res    *types.VsanDisablePerfNetworkDiagnosticResultsResponse `xml:"urn:vsan VsanDisablePerfNetworkDiagnosticResultsResponse,omitempty"`
	Fault_ *soap.Fault                                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanDisablePerfNetworkDiagnosticResultsBody) Fault() *soap.Fault { return b.Fault_ }

func VsanDisablePerfNetworkDiagnosticResults(ctx context.Context, r soap.RoundTripper, req *types.VsanDisablePerfNetworkDiagnosticResults) (*types.VsanDisablePerfNetworkDiagnosticResultsResponse, error) {
	var reqBody, resBody VsanDisablePerfNetworkDiagnosticResultsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanSetPhoneHomeResourceIdBody struct {
	Req    *types.VsanSetPhoneHomeResourceId         `xml:"urn:vsan VsanSetPhoneHomeResourceId,omitempty"`
	Res    *types.VsanSetPhoneHomeResourceIdResponse `xml:"urn:vsan VsanSetPhoneHomeResourceIdResponse,omitempty"`
	Fault_ *soap.Fault                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanSetPhoneHomeResourceIdBody) Fault() *soap.Fault { return b.Fault_ }

func VsanSetPhoneHomeResourceId(ctx context.Context, r soap.RoundTripper, req *types.VsanSetPhoneHomeResourceId) (*types.VsanSetPhoneHomeResourceIdResponse, error) {
	var reqBody, resBody VsanSetPhoneHomeResourceIdBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanResetLastHostBundleCollectionTimeBody struct {
	Req    *types.VsanResetLastHostBundleCollectionTime         `xml:"urn:vsan VsanResetLastHostBundleCollectionTime,omitempty"`
	Res    *types.VsanResetLastHostBundleCollectionTimeResponse `xml:"urn:vsan VsanResetLastHostBundleCollectionTimeResponse,omitempty"`
	Fault_ *soap.Fault                                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanResetLastHostBundleCollectionTimeBody) Fault() *soap.Fault { return b.Fault_ }

func VsanResetLastHostBundleCollectionTime(ctx context.Context, r soap.RoundTripper, req *types.VsanResetLastHostBundleCollectionTime) (*types.VsanResetLastHostBundleCollectionTimeResponse, error) {
	var reqBody, resBody VsanResetLastHostBundleCollectionTimeBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanVcQueryClusterLogsCatalogVersionBody struct {
	Req    *types.VsanVcQueryClusterLogsCatalogVersion         `xml:"urn:vsan VsanVcQueryClusterLogsCatalogVersion,omitempty"`
	Res    *types.VsanVcQueryClusterLogsCatalogVersionResponse `xml:"urn:vsan VsanVcQueryClusterLogsCatalogVersionResponse,omitempty"`
	Fault_ *soap.Fault                                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanVcQueryClusterLogsCatalogVersionBody) Fault() *soap.Fault { return b.Fault_ }

func VsanVcQueryClusterLogsCatalogVersion(ctx context.Context, r soap.RoundTripper, req *types.VsanVcQueryClusterLogsCatalogVersion) (*types.VsanVcQueryClusterLogsCatalogVersionResponse, error) {
	var reqBody, resBody VsanVcQueryClusterLogsCatalogVersionBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanQueryHostPatchBody struct {
	Req    *types.VsanQueryHostPatch         `xml:"urn:vsan VsanQueryHostPatch,omitempty"`
	Res    *types.VsanQueryHostPatchResponse `xml:"urn:vsan VsanQueryHostPatchResponse,omitempty"`
	Fault_ *soap.Fault                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanQueryHostPatchBody) Fault() *soap.Fault { return b.Fault_ }

func VsanQueryHostPatch(ctx context.Context, r soap.RoundTripper, req *types.VsanQueryHostPatch) (*types.VsanQueryHostPatchResponse, error) {
	var reqBody, resBody VsanQueryHostPatchBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryVsanCloudHealthProfilingDataBody struct {
	Req    *types.QueryVsanCloudHealthProfilingData         `xml:"urn:vsan QueryVsanCloudHealthProfilingData,omitempty"`
	Res    *types.QueryVsanCloudHealthProfilingDataResponse `xml:"urn:vsan QueryVsanCloudHealthProfilingDataResponse,omitempty"`
	Fault_ *soap.Fault                                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryVsanCloudHealthProfilingDataBody) Fault() *soap.Fault { return b.Fault_ }

func QueryVsanCloudHealthProfilingData(ctx context.Context, r soap.RoundTripper, req *types.QueryVsanCloudHealthProfilingData) (*types.QueryVsanCloudHealthProfilingDataResponse, error) {
	var reqBody, resBody QueryVsanCloudHealthProfilingDataBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanCloudHealthRunDaemonBody struct {
	Req    *types.VsanCloudHealthRunDaemon         `xml:"urn:vsan VsanCloudHealthRunDaemon,omitempty"`
	Res    *types.VsanCloudHealthRunDaemonResponse `xml:"urn:vsan VsanCloudHealthRunDaemonResponse,omitempty"`
	Fault_ *soap.Fault                             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanCloudHealthRunDaemonBody) Fault() *soap.Fault { return b.Fault_ }

func VsanCloudHealthRunDaemon(ctx context.Context, r soap.RoundTripper, req *types.VsanCloudHealthRunDaemon) (*types.VsanCloudHealthRunDaemonResponse, error) {
	var reqBody, resBody VsanCloudHealthRunDaemonBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}
