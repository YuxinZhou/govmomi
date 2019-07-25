package types

import (
	"github.com/vmware/govmomi/vim25/types"
	"time"
)

type VsanPerfDiagnoseResponse struct {
	Returnval []VsanPerfDiagnosticResult `xml:"returnval"`
}

type VsanPerfDiagnose VsanPerfDiagnoseRequestType

type VsanPerfDiagnoseRequestType struct {
	This              types.ManagedObjectReference  `xml:"_this"`
	PerfDiagnoseQuery VsanPerfDiagnoseQuerySpec     `xml:"perfDiagnoseQuery"`
	Cluster           *types.ManagedObjectReference `xml:"cluster,omitempty"`
}

type VsanPerfQueryClusterHealthResponse struct {
	Returnval []types.DynamicData `xml:"returnval"`
}

type VsanPerfQueryClusterHealth VsanPerfQueryClusterHealthRequestType

type VsanPerfQueryClusterHealthRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Cluster types.ManagedObjectReference `xml:"cluster"`
}

type VsanPerfStopStatsCollectorResponse struct {
	Returnval bool `xml:"returnval"`
}

type VsanPerfStopStatsCollector VsanPerfStopStatsCollectorRequestType

type VsanPerfStopStatsCollectorRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
}

type GetVsanPerfDiagnosisResultResponse struct {
	Returnval []VsanPerfDiagnosticResult `xml:"returnval"`
}

type GetVsanPerfDiagnosisResult GetVsanPerfDiagnosisResultRequestType

type GetVsanPerfDiagnosisResultRequestType struct {
	This    types.ManagedObjectReference  `xml:"_this"`
	Task    types.ManagedObjectReference  `xml:"task"`
	Cluster *types.ManagedObjectReference `xml:"cluster,omitempty"`
}

type VsanPerfToggleVerboseModeResponse struct {
}

type VsanPerfToggleVerboseMode VsanPerfToggleVerboseModeRequestType

type VsanPerfToggleVerboseModeRequestType struct {
	This        types.ManagedObjectReference  `xml:"_this"`
	Cluster     *types.ManagedObjectReference `xml:"cluster,omitempty"`
	VerboseMode bool                          `xml:"verboseMode"`
}

type VsanPerfDiagnoseTaskResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type VsanPerfDiagnoseTask VsanPerfDiagnoseTaskRequestType

type VsanPerfDiagnoseTaskRequestType struct {
	This              types.ManagedObjectReference  `xml:"_this"`
	PerfDiagnoseQuery VsanPerfDiagnoseQuerySpec     `xml:"perfDiagnoseQuery"`
	Cluster           *types.ManagedObjectReference `xml:"cluster,omitempty"`
}

type QueryVsanPerfTopEntitiesResponse struct {
	Returnval []VsanPerfTopEntities `xml:"returnval"`
}

type QueryVsanPerfTopEntities QueryVsanPerfTopEntitiesRequestType

type QueryVsanPerfTopEntitiesRequestType struct {
	This          types.ManagedObjectReference  `xml:"_this"`
	Cluster       *types.ManagedObjectReference `xml:"cluster,omitempty"`
	SortAscending *bool                         `xml:"sortAscending"`
	PointInTime   string                        `xml:"pointInTime,omitempty"`
	NumEntities   int32                         `xml:"numEntities,omitempty"`
}

type VsanPerfSetClusterMembersResponse struct {
	Returnval bool `xml:"returnval"`
}

type VsanPerfSetClusterMembers VsanPerfSetClusterMembersRequestType

type VsanPerfSetClusterMembersRequestType struct {
	This           types.ManagedObjectReference `xml:"_this"`
	MemberInfo     []VsanPerfMemberInfo         `xml:"memberInfo"`
	ConfigGen      *VsanConfigGeneration        `xml:"configGen,omitempty"`
	SubClusterUuid string                       `xml:"subClusterUuid,omitempty"`
}

type QueryHostNetworkConfigInfoVnicResponse struct {
	Returnval []types.HostVirtualNic `xml:"returnval"`
}

type QueryHostNetworkConfigInfoVnic QueryHostNetworkConfigInfoVnicRequestType

type QueryHostNetworkConfigInfoVnicRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
}

type SubmitFeedbackForDiagnosisResultResponse struct {
	Returnval bool `xml:"returnval"`
}

type SubmitFeedbackForDiagnosisResult SubmitFeedbackForDiagnosisResultRequestType

type SubmitFeedbackForDiagnosisResultRequestType struct {
	This          types.ManagedObjectReference `xml:"_this"`
	FeedbackSpec  VsanPerfDiagnoseFeedbackSpec `xml:"feedbackSpec"`
	FeedbackValue bool                         `xml:"feedbackValue"`
	Comment       string                       `xml:"comment,omitempty"`
	Cluster       types.ManagedObjectReference `xml:"cluster"`
}

type VsanPerfLoginResponse struct {
	Returnval bool `xml:"returnval"`
}

type VsanPerfLogin VsanPerfLoginRequestType

type VsanPerfLoginRequestType struct {
	This  types.ManagedObjectReference `xml:"_this"`
	Token string                       `xml:"token"`
}

type VsanPerfQueryTimeRangesResponse struct {
	Returnval []VsanPerfTimeRange `xml:"returnval"`
}

type VsanPerfQueryTimeRanges VsanPerfQueryTimeRangesRequestType

type VsanPerfQueryTimeRangesRequestType struct {
	This      types.ManagedObjectReference  `xml:"_this"`
	Cluster   *types.ManagedObjectReference `xml:"cluster,omitempty"`
	QuerySpec VsanPerfTimeRangeQuerySpec    `xml:"querySpec"`
}

type VsanPerfSetStatsObjectPolicyResponse struct {
	Returnval bool `xml:"returnval"`
}

type VsanPerfSetStatsObjectPolicy VsanPerfSetStatsObjectPolicyRequestType

type VsanPerfSetStatsObjectPolicyRequestType struct {
	This    types.ManagedObjectReference     `xml:"_this"`
	Cluster *types.ManagedObjectReference    `xml:"cluster,omitempty"`
	Profile *types.VirtualMachineProfileSpec `xml:"profile,omitempty"`
}

type VsanPerfStartStatsCollectorResponse struct {
	Returnval bool `xml:"returnval"`
}

type VsanPerfStartStatsCollector VsanPerfStartStatsCollectorRequestType

type VsanPerfStartStatsCollectorRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
}

type VsanPerfInjectFakeHistoryResponse struct {
	Returnval bool `xml:"returnval"`
}

type VsanPerfInjectFakeHistory VsanPerfInjectFakeHistoryRequestType

type VsanPerfInjectFakeHistoryRequestType struct {
	This           types.ManagedObjectReference `xml:"_this"`
	Cluster        types.ManagedObjectReference `xml:"cluster"`
	StartTimestamp int64                        `xml:"startTimestamp"`
	EndTimestamp   int64                        `xml:"endTimestamp"`
}

type VsanPerfCreateStatsObjectResponse struct {
	Returnval string `xml:"returnval"`
}

type VsanPerfCreateStatsObject VsanPerfCreateStatsObjectRequestType

type VsanPerfCreateStatsObjectRequestType struct {
	This    types.ManagedObjectReference     `xml:"_this"`
	Cluster *types.ManagedObjectReference    `xml:"cluster,omitempty"`
	Profile *types.VirtualMachineProfileSpec `xml:"profile,omitempty"`
}

type VsanPerfQueryPerfResponse struct {
	Returnval []VsanPerfEntityMetricCSV `xml:"returnval"`
}

type VsanPerfQueryPerf VsanPerfQueryPerfRequestType

type VsanPerfQueryPerfRequestType struct {
	This       types.ManagedObjectReference  `xml:"_this"`
	QuerySpecs []VsanPerfQuerySpec           `xml:"querySpecs"`
	Cluster    *types.ManagedObjectReference `xml:"cluster,omitempty"`
}

type VsanPerfDiagnoseFromFileResponse struct {
	Returnval []VsanPerfDiagnosticResult `xml:"returnval"`
}

type VsanPerfDiagnoseFromFile VsanPerfDiagnoseFromFileRequestType

type VsanPerfDiagnoseFromFileRequestType struct {
	This              types.ManagedObjectReference  `xml:"_this"`
	PerfSvcResultFile string                        `xml:"perfSvcResultFile"`
	PerfDiagnoseQuery VsanPerfDiagnoseQuerySpec     `xml:"perfDiagnoseQuery"`
	Cluster           *types.ManagedObjectReference `xml:"cluster,omitempty"`
}

type VsanPerfCreateStatsObjectTaskResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type VsanPerfCreateStatsObjectTask VsanPerfCreateStatsObjectTaskRequestType

type VsanPerfCreateStatsObjectTaskRequestType struct {
	This    types.ManagedObjectReference     `xml:"_this"`
	Cluster *types.ManagedObjectReference    `xml:"cluster,omitempty"`
	Profile *types.VirtualMachineProfileSpec `xml:"profile,omitempty"`
}

type VsanPerfDeleteStatsObjectTaskResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type VsanPerfDeleteStatsObjectTask VsanPerfDeleteStatsObjectTaskRequestType

type VsanPerfDeleteStatsObjectTaskRequestType struct {
	This    types.ManagedObjectReference  `xml:"_this"`
	Cluster *types.ManagedObjectReference `xml:"cluster,omitempty"`
}

type VsanPerfGetAggregatedEntityTypesResponse struct {
	Returnval []VsanPerfEntityType `xml:"returnval"`
}

type VsanPerfGetAggregatedEntityTypes VsanPerfGetAggregatedEntityTypesRequestType

type VsanPerfGetAggregatedEntityTypesRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
}

type VsanPerfDeleteTimeRangeResponse struct {
}

type VsanPerfDeleteTimeRange VsanPerfDeleteTimeRangeRequestType

type VsanPerfDeleteTimeRangeRequestType struct {
	This    types.ManagedObjectReference  `xml:"_this"`
	Cluster *types.ManagedObjectReference `xml:"cluster,omitempty"`
	Name    string                        `xml:"name"`
}

type QueryVsanHostConfigInfoResponse struct {
	Returnval types.VsanHostConfigInfo `xml:"returnval"`
}

type QueryVsanHostConfigInfo QueryVsanHostConfigInfoRequestType

type QueryVsanHostConfigInfoRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
}

type VsanPerfQueryStatsObjectInformationResponse struct {
	Returnval VsanObjectInformation `xml:"returnval"`
}

type VsanPerfQueryStatsObjectInformation VsanPerfQueryStatsObjectInformationRequestType

type VsanPerfQueryStatsObjectInformationRequestType struct {
	This    types.ManagedObjectReference  `xml:"_this"`
	Cluster *types.ManagedObjectReference `xml:"cluster,omitempty"`
}

type VsanPerfDeleteStatsObjectResponse struct {
	Returnval bool `xml:"returnval"`
}

type VsanPerfDeleteStatsObject VsanPerfDeleteStatsObjectRequestType

type VsanPerfDeleteStatsObjectRequestType struct {
	This    types.ManagedObjectReference  `xml:"_this"`
	Cluster *types.ManagedObjectReference `xml:"cluster,omitempty"`
}

type VsanPerfSaveTimeRangesResponse struct {
}

type VsanPerfSaveTimeRanges VsanPerfSaveTimeRangesRequestType

type VsanPerfSaveTimeRangesRequestType struct {
	This       types.ManagedObjectReference  `xml:"_this"`
	Cluster    *types.ManagedObjectReference `xml:"cluster,omitempty"`
	TimeRanges []VsanPerfTimeRange           `xml:"timeRanges"`
}

type VsanPerfQueryNodeInformationResponse struct {
	Returnval []VsanPerfNodeInformation `xml:"returnval"`
}

type VsanPerfQueryNodeInformation VsanPerfQueryNodeInformationRequestType

type VsanPerfQueryNodeInformationRequestType struct {
	This    types.ManagedObjectReference  `xml:"_this"`
	Cluster *types.ManagedObjectReference `xml:"cluster,omitempty"`
}

type VsanPerfGetSupportedDiagnosticExceptionsResponse struct {
	Returnval []VsanPerfDiagnosticException `xml:"returnval"`
}

type VsanPerfGetSupportedDiagnosticExceptions VsanPerfGetSupportedDiagnosticExceptionsRequestType

type VsanPerfGetSupportedDiagnosticExceptionsRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
}

type VsanPerfGetVcMoRefFromPerfEntityRefIdResponse struct {
	Returnval []VsanPerfEntityInfo `xml:"returnval"`
}

type VsanPerfGetVcMoRefFromPerfEntityRefId VsanPerfGetVcMoRefFromPerfEntityRefIdRequestType

type VsanPerfGetVcMoRefFromPerfEntityRefIdRequestType struct {
	This         types.ManagedObjectReference `xml:"_this"`
	Cluster      types.ManagedObjectReference `xml:"cluster"`
	EntityRefIds []string                     `xml:"entityRefIds"`
}

type VsanPerfGetSupportedEntityTypesResponse struct {
	Returnval []VsanPerfEntityType `xml:"returnval"`
}

type VsanPerfGetSupportedEntityTypes VsanPerfGetSupportedEntityTypesRequestType

type VsanPerfGetSupportedEntityTypesRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
}

type VsanPerformanceManager struct {
	types.ManagedObjectReference
}

type InitializeDiskMappingsResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type InitializeDiskMappings InitializeDiskMappingsRequestType

type InitializeDiskMappingsRequestType struct {
	This types.ManagedObjectReference       `xml:"_this"`
	Spec VimVsanHostDiskMappingCreationSpec `xml:"spec"`
}

type QueryDiskMappingsResponse struct {
	Returnval []VimVsanHostDiskMapInfoEx `xml:"returnval"`
}

type QueryDiskMappings QueryDiskMappingsRequestType

type QueryDiskMappingsRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
	Host types.ManagedObjectReference `xml:"host"`
}

type QueryClusterDataEfficiencyCapacityStateResponse struct {
	Returnval VimVsanDataEfficiencyCapacityState `xml:"returnval"`
}

type QueryClusterDataEfficiencyCapacityState QueryClusterDataEfficiencyCapacityStateRequestType

type QueryClusterDataEfficiencyCapacityStateRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Cluster types.ManagedObjectReference `xml:"cluster"`
}

type RetrieveAllFlashCapabilityResponse struct {
	Returnval VimVsanHostVsanHostCapability `xml:"returnval"`
}

type RetrieveAllFlashCapability RetrieveAllFlashCapabilityRequestType

type RetrieveAllFlashCapabilityRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
	Host types.ManagedObjectReference `xml:"host"`
}

type RetrieveAllFlashCapabilitiesResponse struct {
	Returnval []VimVsanHostVsanHostCapability `xml:"returnval"`
}

type RetrieveAllFlashCapabilities RetrieveAllFlashCapabilitiesRequestType

type RetrieveAllFlashCapabilitiesRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Cluster types.ManagedObjectReference `xml:"cluster"`
}

type RebuildDiskMappingResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type RebuildDiskMapping RebuildDiskMappingRequestType

type RebuildDiskMappingRequestType struct {
	This            types.ManagedObjectReference `xml:"_this"`
	Host            types.ManagedObjectReference `xml:"host"`
	Mapping         types.VsanHostDiskMapping    `xml:"mapping"`
	MaintenanceSpec types.HostMaintenanceSpec    `xml:"maintenanceSpec"`
}

type VimClusterVsanVcDiskManagementSystem struct {
	types.ManagedObjectReference
}

type VsanClusterMgmtPersistStressOptionsResponse struct {
	Returnval bool `xml:"returnval"`
}

type VsanClusterMgmtPersistStressOptions VsanClusterMgmtPersistStressOptionsRequestType

type VsanClusterMgmtPersistStressOptionsRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
}

type VsanClusterMgmtClearStressOptionsResponse struct {
	Returnval bool `xml:"returnval"`
}

type VsanClusterMgmtClearStressOptions VsanClusterMgmtClearStressOptionsRequestType

type VsanClusterMgmtClearStressOptionsRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
}

type VsanClusterMgmtServiceRestartResponse struct {
	Returnval bool `xml:"returnval"`
}

type VsanClusterMgmtServiceRestart VsanClusterMgmtServiceRestartRequestType

type VsanClusterMgmtServiceRestartRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
}

type VsanGetPersistedClusterStateResponse struct {
	Returnval VsanClusterPersistedState `xml:"returnval"`
}

type VsanGetPersistedClusterState VsanGetPersistedClusterStateRequestType

type VsanGetPersistedClusterStateRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Cluster types.ManagedObjectReference `xml:"cluster"`
}

type VsanClusterMgmtGetStatsResponse struct {
	Returnval string `xml:"returnval"`
}

type VsanClusterMgmtGetStats VsanClusterMgmtGetStatsRequestType

type VsanClusterMgmtGetStatsRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
}

type VsanClusterMgmtSetStressOptionResponse struct {
	Returnval bool `xml:"returnval"`
}

type VsanClusterMgmtSetStressOption VsanClusterMgmtSetStressOptionRequestType

type VsanClusterMgmtSetStressOptionRequestType struct {
	This  types.ManagedObjectReference `xml:"_this"`
	Key   string                       `xml:"key"`
	Value int64                        `xml:"value"`
}

type VsanClusterMgmtPurgeAllPersistedStateResponse struct {
	Returnval bool `xml:"returnval"`
}

type VsanClusterMgmtPurgeAllPersistedState VsanClusterMgmtPurgeAllPersistedStateRequestType

type VsanClusterMgmtPurgeAllPersistedStateRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
}

type VsanClusterMgmtDebugSystem struct {
	types.ManagedObjectReference
}

type VsanCompleteMigrateVmsToVdsResponse struct {
}

type VsanCompleteMigrateVmsToVds VsanCompleteMigrateVmsToVdsRequestType

type VsanCompleteMigrateVmsToVdsRequestType struct {
	This     types.ManagedObjectReference `xml:"_this"`
	JobId    string                       `xml:"jobId"`
	NewState string                       `xml:"newState"`
}

type VsanMigrateVmsToVdsResponse struct {
	Returnval string `xml:"returnval"`
}

type VsanMigrateVmsToVds VsanMigrateVmsToVdsRequestType

type VsanMigrateVmsToVdsRequestType struct {
	This          types.ManagedObjectReference `xml:"_this"`
	VmConfigSpecs []VsanVmVdsMigrationSpec     `xml:"vmConfigSpecs"`
	VdsUuid       string                       `xml:"vdsUuid"`
	TimeoutSec    int64                        `xml:"timeoutSec"`
	Revert        *bool                        `xml:"revert"`
}

type VsanHostVdsSystem struct {
	types.ManagedObjectReference
}

type VSANHostGetStretchedClusterInfoFromCmmdsResponse struct {
	Returnval []VimHostVSANStretchedClusterHostInfo `xml:"returnval"`
}

type VSANHostGetStretchedClusterInfoFromCmmds VSANHostGetStretchedClusterInfoFromCmmdsRequestType

type VSANHostGetStretchedClusterInfoFromCmmdsRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
}

type VSANWitnessJoinVsanClusterResponse struct {
}

type VSANWitnessJoinVsanCluster VSANWitnessJoinVsanClusterRequestType

type VSANWitnessJoinVsanClusterRequestType struct {
	This               types.ManagedObjectReference `xml:"_this"`
	ClusterUuid        string                       `xml:"clusterUuid"`
	PreferredFd        string                       `xml:"preferredFd"`
	DisableVsanAllowed *bool                        `xml:"disableVsanAllowed"`
}

type VSANWitnessSetPreferredFaultDomainResponse struct {
}

type VSANWitnessSetPreferredFaultDomain VSANWitnessSetPreferredFaultDomainRequestType

type VSANWitnessSetPreferredFaultDomainRequestType struct {
	This        types.ManagedObjectReference `xml:"_this"`
	PreferredFd string                       `xml:"preferredFd"`
}

type VSANHostAddUnicastAgentResponse struct {
}

type VSANHostAddUnicastAgent VSANHostAddUnicastAgentRequestType

type VSANHostAddUnicastAgentRequestType struct {
	This           types.ManagedObjectReference `xml:"_this"`
	WitnessAddress string                       `xml:"witnessAddress"`
	WitnessPort    int32                        `xml:"witnessPort,omitempty"`
	Overwrite      *bool                        `xml:"overwrite"`
	WitnessVersion string                       `xml:"witnessVersion,omitempty"`
	UcastAgentUuid string                       `xml:"UcastAgentUuid,omitempty"`
	NodeType       string                       `xml:"nodeType,omitempty"`
}

type VSANClusterGetPreferredFaultDomainResponse struct {
	Returnval VimHostVSANCmmdsPreferredFaultDomainInfo `xml:"returnval"`
}

type VSANClusterGetPreferredFaultDomain VSANClusterGetPreferredFaultDomainRequestType

type VSANClusterGetPreferredFaultDomainRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
}

type VSANWitnessLeaveVsanClusterResponse struct {
}

type VSANWitnessLeaveVsanCluster VSANWitnessLeaveVsanClusterRequestType

type VSANWitnessLeaveVsanClusterRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
}

type VSANHostGetStretchedClusterCapabilityResponse struct {
	Returnval VimHostVSANStretchedClusterHostCapability `xml:"returnval"`
}

type VSANHostGetStretchedClusterCapability VSANHostGetStretchedClusterCapabilityRequestType

type VSANHostGetStretchedClusterCapabilityRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
}

type VSANHostRemoveUnicastAgentResponse struct {
}

type VSANHostRemoveUnicastAgent VSANHostRemoveUnicastAgentRequestType

type VSANHostRemoveUnicastAgentRequestType struct {
	This            types.ManagedObjectReference `xml:"_this"`
	WitnessAddress  string                       `xml:"witnessAddress"`
	IgnoreExistence *bool                        `xml:"ignoreExistence"`
}

type VSANHostListUnicastAgentResponse struct {
	Returnval string `xml:"returnval"`
}

type VSANHostListUnicastAgent VSANHostListUnicastAgentRequestType

type VSANHostListUnicastAgentRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
}

type VimHostVsanStretchedClusterSystem struct {
	types.ManagedObjectReference
}

type VsanVitGetIscsiTargetServiceStatusResponse struct {
	Returnval bool `xml:"returnval"`
}

type VsanVitGetIscsiTargetServiceStatus VsanVitGetIscsiTargetServiceStatusRequestType

type VsanVitGetIscsiTargetServiceStatusRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Cluster types.ManagedObjectReference `xml:"cluster"`
}

type VsanVitRemoveIscsiInitiatorsFromTargetResponse struct {
}

type VsanVitRemoveIscsiInitiatorsFromTarget VsanVitRemoveIscsiInitiatorsFromTargetRequestType

type VsanVitRemoveIscsiInitiatorsFromTargetRequestType struct {
	This           types.ManagedObjectReference `xml:"_this"`
	Cluster        types.ManagedObjectReference `xml:"cluster"`
	TargetAlias    string                       `xml:"targetAlias"`
	InitiatorNames []string                     `xml:"initiatorNames"`
}

type VsanVitRemoveIscsiInitiatorsFromGroupResponse struct {
}

type VsanVitRemoveIscsiInitiatorsFromGroup VsanVitRemoveIscsiInitiatorsFromGroupRequestType

type VsanVitRemoveIscsiInitiatorsFromGroupRequestType struct {
	This               types.ManagedObjectReference `xml:"_this"`
	Cluster            types.ManagedObjectReference `xml:"cluster"`
	InitiatorGroupName string                       `xml:"initiatorGroupName"`
	InitiatorNames     []string                     `xml:"initiatorNames"`
}

type VsanVitEditIscsiLUNResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type VsanVitEditIscsiLUN VsanVitEditIscsiLUNRequestType

type VsanVitEditIscsiLUNRequestType struct {
	This        types.ManagedObjectReference `xml:"_this"`
	Cluster     types.ManagedObjectReference `xml:"cluster"`
	TargetAlias string                       `xml:"targetAlias"`
	LunSpec     VsanIscsiLUNSpec             `xml:"lunSpec"`
}

type VsanVitGetIscsiTargetServiceProcessesStatusResponse struct {
	Returnval []types.KeyValue `xml:"returnval"`
}

type VsanVitGetIscsiTargetServiceProcessesStatus VsanVitGetIscsiTargetServiceProcessesStatusRequestType

type VsanVitGetIscsiTargetServiceProcessesStatusRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Cluster types.ManagedObjectReference `xml:"cluster"`
}

type VsanVitGetIscsiLUNResponse struct {
	Returnval VsanIscsiLUN `xml:"returnval"`
}

type VsanVitGetIscsiLUN VsanVitGetIscsiLUNRequestType

type VsanVitGetIscsiLUNRequestType struct {
	This        types.ManagedObjectReference `xml:"_this"`
	Cluster     types.ManagedObjectReference `xml:"cluster"`
	TargetAlias string                       `xml:"targetAlias"`
	LunId       int32                        `xml:"lunId"`
}

type VsanVitEditIscsiTargetResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type VsanVitEditIscsiTarget VsanVitEditIscsiTargetRequestType

type VsanVitEditIscsiTargetRequestType struct {
	This       types.ManagedObjectReference `xml:"_this"`
	Cluster    types.ManagedObjectReference `xml:"cluster"`
	TargetSpec VsanIscsiTargetSpec          `xml:"targetSpec"`
}

type VsanVitAddIscsiInitiatorsToGroupResponse struct {
}

type VsanVitAddIscsiInitiatorsToGroup VsanVitAddIscsiInitiatorsToGroupRequestType

type VsanVitAddIscsiInitiatorsToGroupRequestType struct {
	This               types.ManagedObjectReference `xml:"_this"`
	Cluster            types.ManagedObjectReference `xml:"cluster"`
	InitiatorGroupName string                       `xml:"initiatorGroupName"`
	InitiatorNames     []string                     `xml:"initiatorNames"`
}

type VsanVitAddIscsiInitiatorsToTargetResponse struct {
}

type VsanVitAddIscsiInitiatorsToTarget VsanVitAddIscsiInitiatorsToTargetRequestType

type VsanVitAddIscsiInitiatorsToTargetRequestType struct {
	This           types.ManagedObjectReference `xml:"_this"`
	Cluster        types.ManagedObjectReference `xml:"cluster"`
	TargetAlias    string                       `xml:"targetAlias"`
	InitiatorNames []string                     `xml:"initiatorNames"`
}

type VsanVitQueryIscsiTargetServiceVersionResponse struct {
	Returnval string `xml:"returnval"`
}

type VsanVitQueryIscsiTargetServiceVersion VsanVitQueryIscsiTargetServiceVersionRequestType

type VsanVitQueryIscsiTargetServiceVersionRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
}

type VsanVitAddIscsiTargetToGroupResponse struct {
}

type VsanVitAddIscsiTargetToGroup VsanVitAddIscsiTargetToGroupRequestType

type VsanVitAddIscsiTargetToGroupRequestType struct {
	This               types.ManagedObjectReference `xml:"_this"`
	Cluster            types.ManagedObjectReference `xml:"cluster"`
	InitiatorGroupName string                       `xml:"initiatorGroupName"`
	TargetAlias        string                       `xml:"targetAlias"`
}

type VsanVitRemoveIscsiTargetFromGroupResponse struct {
}

type VsanVitRemoveIscsiTargetFromGroup VsanVitRemoveIscsiTargetFromGroupRequestType

type VsanVitRemoveIscsiTargetFromGroupRequestType struct {
	This               types.ManagedObjectReference `xml:"_this"`
	Cluster            types.ManagedObjectReference `xml:"cluster"`
	InitiatorGroupName string                       `xml:"initiatorGroupName"`
	TargetAlias        string                       `xml:"targetAlias"`
}

type VsanVitGetIscsiLUNsResponse struct {
	Returnval []VsanIscsiLUN `xml:"returnval"`
}

type VsanVitGetIscsiLUNs VsanVitGetIscsiLUNsRequestType

type VsanVitGetIscsiLUNsRequestType struct {
	This          types.ManagedObjectReference `xml:"_this"`
	Cluster       types.ManagedObjectReference `xml:"cluster"`
	TargetAliases []string                     `xml:"targetAliases,omitempty"`
}

type VsanVitRemoveIscsiLUNResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type VsanVitRemoveIscsiLUN VsanVitRemoveIscsiLUNRequestType

type VsanVitRemoveIscsiLUNRequestType struct {
	This        types.ManagedObjectReference `xml:"_this"`
	Cluster     types.ManagedObjectReference `xml:"cluster"`
	TargetAlias string                       `xml:"targetAlias"`
	LunId       int32                        `xml:"lunId"`
}

type VsanVitGetIscsiInitiatorGroupResponse struct {
	Returnval VsanIscsiInitiatorGroup `xml:"returnval"`
}

type VsanVitGetIscsiInitiatorGroup VsanVitGetIscsiInitiatorGroupRequestType

type VsanVitGetIscsiInitiatorGroupRequestType struct {
	This               types.ManagedObjectReference `xml:"_this"`
	Cluster            types.ManagedObjectReference `xml:"cluster"`
	InitiatorGroupName string                       `xml:"initiatorGroupName"`
}

type VsanVitEditIscsiTargetServiceResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type VsanVitEditIscsiTargetService VsanVitEditIscsiTargetServiceRequestType

type VsanVitEditIscsiTargetServiceRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Cluster types.ManagedObjectReference `xml:"cluster"`
	Spec    *VsanIscsiTargetServiceSpec  `xml:"spec,omitempty"`
}

type VsanVitRemoveIscsiInitiatorGroupResponse struct {
}

type VsanVitRemoveIscsiInitiatorGroup VsanVitRemoveIscsiInitiatorGroupRequestType

type VsanVitRemoveIscsiInitiatorGroupRequestType struct {
	This               types.ManagedObjectReference `xml:"_this"`
	Cluster            types.ManagedObjectReference `xml:"cluster"`
	InitiatorGroupName string                       `xml:"initiatorGroupName"`
}

type VsanVitCreateHomeObjectResponse struct {
}

type VsanVitCreateHomeObject VsanVitCreateHomeObjectRequestType

type VsanVitCreateHomeObjectRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Cluster types.ManagedObjectReference `xml:"cluster"`
	Spec    *VsanIscsiHomeObjectSpec     `xml:"spec,omitempty"`
}

type VsanVitGetHomeObjectResponse struct {
	Returnval VsanObjectInformation `xml:"returnval"`
}

type VsanVitGetHomeObject VsanVitGetHomeObjectRequestType

type VsanVitGetHomeObjectRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Cluster types.ManagedObjectReference `xml:"cluster"`
}

type VsanVitGetIscsiTargetResponse struct {
	Returnval VsanIscsiTarget `xml:"returnval"`
}

type VsanVitGetIscsiTarget VsanVitGetIscsiTargetRequestType

type VsanVitGetIscsiTargetRequestType struct {
	This        types.ManagedObjectReference `xml:"_this"`
	Cluster     types.ManagedObjectReference `xml:"cluster"`
	TargetAlias string                       `xml:"targetAlias"`
}

type VsanVitRemoveIscsiTargetResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type VsanVitRemoveIscsiTarget VsanVitRemoveIscsiTargetRequestType

type VsanVitRemoveIscsiTargetRequestType struct {
	This        types.ManagedObjectReference `xml:"_this"`
	Cluster     types.ManagedObjectReference `xml:"cluster"`
	TargetAlias string                       `xml:"targetAlias"`
}

type VsanVitAddIscsiLUNResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type VsanVitAddIscsiLUN VsanVitAddIscsiLUNRequestType

type VsanVitAddIscsiLUNRequestType struct {
	This        types.ManagedObjectReference `xml:"_this"`
	Cluster     types.ManagedObjectReference `xml:"cluster"`
	TargetAlias string                       `xml:"targetAlias"`
	LunSpec     VsanIscsiLUNSpec             `xml:"lunSpec"`
}

type VsanVitGetIscsiInitiatorGroupsResponse struct {
	Returnval []VsanIscsiInitiatorGroup `xml:"returnval"`
}

type VsanVitGetIscsiInitiatorGroups VsanVitGetIscsiInitiatorGroupsRequestType

type VsanVitGetIscsiInitiatorGroupsRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Cluster types.ManagedObjectReference `xml:"cluster"`
}

type VsanVitEnableIscsiTargetServiceResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type VsanVitEnableIscsiTargetService VsanVitEnableIscsiTargetServiceRequestType

type VsanVitEnableIscsiTargetServiceRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Cluster types.ManagedObjectReference `xml:"cluster"`
	Spec    VsanIscsiTargetServiceSpec   `xml:"spec"`
}

type VsanVitGetIscsiTargetsResponse struct {
	Returnval []VsanIscsiTarget `xml:"returnval"`
}

type VsanVitGetIscsiTargets VsanVitGetIscsiTargetsRequestType

type VsanVitGetIscsiTargetsRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Cluster types.ManagedObjectReference `xml:"cluster"`
}

type VsanVitAddIscsiTargetResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type VsanVitAddIscsiTarget VsanVitAddIscsiTargetRequestType

type VsanVitAddIscsiTargetRequestType struct {
	This       types.ManagedObjectReference `xml:"_this"`
	Cluster    types.ManagedObjectReference `xml:"cluster"`
	TargetSpec VsanIscsiTargetSpec          `xml:"targetSpec"`
}

type VsanVitAddIscsiInitiatorGroupResponse struct {
}

type VsanVitAddIscsiInitiatorGroup VsanVitAddIscsiInitiatorGroupRequestType

type VsanVitAddIscsiInitiatorGroupRequestType struct {
	This               types.ManagedObjectReference `xml:"_this"`
	Cluster            types.ManagedObjectReference `xml:"cluster"`
	InitiatorGroupName string                       `xml:"initiatorGroupName"`
}

type VsanVitGetIscsiTargetServiceConfigResponse struct {
	Returnval VsanIscsiTargetServiceConfig `xml:"returnval"`
}

type VsanVitGetIscsiTargetServiceConfig VsanVitGetIscsiTargetServiceConfigRequestType

type VsanVitGetIscsiTargetServiceConfigRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Cluster types.ManagedObjectReference `xml:"cluster"`
}

type VsanIscsiTargetSystem struct {
	types.ManagedObjectReference
}

type PerformVsanUpgradeExResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type PerformVsanUpgradeEx PerformVsanUpgradeExRequestType

type PerformVsanUpgradeExRequestType struct {
	This                   types.ManagedObjectReference   `xml:"_this"`
	Cluster                types.ManagedObjectReference   `xml:"cluster"`
	PerformObjectUpgrade   *bool                          `xml:"performObjectUpgrade"`
	DowngradeFormat        *bool                          `xml:"downgradeFormat"`
	AllowReducedRedundancy *bool                          `xml:"allowReducedRedundancy"`
	ExcludeHosts           []types.ManagedObjectReference `xml:"excludeHosts,omitempty"`
	Spec                   *VsanDiskFormatConversionSpec  `xml:"spec,omitempty"`
}

type VsanQueryUpgradeStatusExResponse struct {
	Returnval VsanUpgradeStatusEx `xml:"returnval"`
}

type VsanQueryUpgradeStatusEx VsanQueryUpgradeStatusExRequestType

type VsanQueryUpgradeStatusExRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Cluster types.ManagedObjectReference `xml:"cluster"`
}

type RetrieveSupportedVsanFormatVersionResponse struct {
	Returnval int32 `xml:"returnval"`
}

type RetrieveSupportedVsanFormatVersion RetrieveSupportedVsanFormatVersionRequestType

type RetrieveSupportedVsanFormatVersionRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Cluster types.ManagedObjectReference `xml:"cluster"`
}

type PerformVsanUpgradePreflightCheckExResponse struct {
	Returnval VsanDiskFormatConversionCheckResult `xml:"returnval"`
}

type PerformVsanUpgradePreflightCheckEx PerformVsanUpgradePreflightCheckExRequestType

type PerformVsanUpgradePreflightCheckExRequestType struct {
	This            types.ManagedObjectReference  `xml:"_this"`
	Cluster         types.ManagedObjectReference  `xml:"cluster"`
	DowngradeFormat *bool                         `xml:"downgradeFormat"`
	Spec            *VsanDiskFormatConversionSpec `xml:"spec,omitempty"`
}

type PerformVsanUpgradePreflightAsyncCheck_TaskResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type PerformVsanUpgradePreflightAsyncCheck_Task PerformVsanUpgradePreflightAsyncCheck_TaskRequestType

type PerformVsanUpgradePreflightAsyncCheck_TaskRequestType struct {
	This            types.ManagedObjectReference  `xml:"_this"`
	Cluster         types.ManagedObjectReference  `xml:"cluster"`
	DowngradeFormat *bool                         `xml:"downgradeFormat"`
	Spec            *VsanDiskFormatConversionSpec `xml:"spec,omitempty"`
}

type VsanUpgradeSystemEx struct {
	types.ManagedObjectReference
}

type VsanQuerySpaceUsageResponse struct {
	Returnval VsanSpaceUsage `xml:"returnval"`
}

type VsanQuerySpaceUsage VsanQuerySpaceUsageRequestType

type VsanQuerySpaceUsageRequestType struct {
	This               types.ManagedObjectReference      `xml:"_this"`
	Cluster            types.ManagedObjectReference      `xml:"cluster"`
	StoragePolicies    []types.VirtualMachineProfileSpec `xml:"storagePolicies,omitempty"`
	WhatifCapacityOnly *bool                             `xml:"whatifCapacityOnly"`
}

type VsanQueryObjectContainerSpaceUsageResponse struct {
	Returnval []VsanContainerSpaceUsage `xml:"returnval"`
}

type VsanQueryObjectContainerSpaceUsage VsanQueryObjectContainerSpaceUsageRequestType

type VsanQueryObjectContainerSpaceUsageRequestType struct {
	This      types.ManagedObjectReference `xml:"_this"`
	Cluster   types.ManagedObjectReference `xml:"cluster"`
	QuerySpec VsanSpaceQuerySpec           `xml:"querySpec"`
}

type VsanQueryRemoteDataProtectionSpaceUsageResponse struct {
	Returnval VsanDataProtectionRemoteSpaceUsage `xml:"returnval"`
}

type VsanQueryRemoteDataProtectionSpaceUsage VsanQueryRemoteDataProtectionSpaceUsageRequestType

type VsanQueryRemoteDataProtectionSpaceUsageRequestType struct {
	This            types.ManagedObjectReference             `xml:"_this"`
	Cluster         types.ManagedObjectReference             `xml:"cluster"`
	TargetSiteInfos []VsanDataProtectionRemoteTargetSiteInfo `xml:"targetSiteInfos"`
}

type VsanSpaceReportSystem struct {
	types.ManagedObjectReference
}

type VsanRetrieveProfilingTraceResponse struct {
	Returnval string `xml:"returnval"`
}

type VsanRetrieveProfilingTrace VsanRetrieveProfilingTraceRequestType

type VsanRetrieveProfilingTraceRequestType struct {
	This    types.ManagedObjectReference  `xml:"_this"`
	Cluster *types.ManagedObjectReference `xml:"cluster,omitempty"`
}

type VsanRemediateVsanClusterResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type VsanRemediateVsanCluster VsanRemediateVsanClusterRequestType

type VsanRemediateVsanClusterRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Cluster types.ManagedObjectReference `xml:"cluster"`
}

type VsanRemediateDataProtectionConfigResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type VsanRemediateDataProtectionConfig VsanRemediateDataProtectionConfigRequestType

type VsanRemediateDataProtectionConfigRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Cluster types.ManagedObjectReference `xml:"cluster"`
}

type VsanReconfigProfilingResponse struct {
}

type VsanReconfigProfiling VsanReconfigProfilingRequestType

type VsanReconfigProfilingRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Enabled bool                         `xml:"enabled"`
}

type VsanRemediateVcResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type VsanRemediateVc VsanRemediateVcRequestType

type VsanRemediateVcRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
}

type VsanRemediateVsanHostResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type VsanRemediateVsanHost VsanRemediateVsanHostRequestType

type VsanRemediateVsanHostRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
	Host types.ManagedObjectReference `xml:"host"`
}

type VsanClusterMgmtInternalSystem struct {
	types.ManagedObjectReference
}

type VsanGetCapabilitiesResponse struct {
	Returnval []VsanCapability `xml:"returnval"`
}

type VsanGetCapabilities VsanGetCapabilitiesRequestType

type VsanGetCapabilitiesRequestType struct {
	This    types.ManagedObjectReference   `xml:"_this"`
	Targets []types.ManagedObjectReference `xml:"targets,omitempty"`
}

type VsanCapabilitySystem struct {
	types.ManagedObjectReference
}

type SetVsanVumConfigResponse struct {
}

type SetVsanVumConfig SetVsanVumConfigRequestType

type SetVsanVumConfigRequestType struct {
	This       types.ManagedObjectReference `xml:"_this"`
	ConfigSpec VsanVumSystemConfigSpec      `xml:"configSpec"`
}

type VsanHostUpdateFirmwareResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type VsanHostUpdateFirmware VsanHostUpdateFirmwareRequestType

type VsanHostUpdateFirmwareRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
	Host types.ManagedObjectReference `xml:"host"`
}

type FetchIsoDepotCookieResponse struct {
}

type FetchIsoDepotCookie FetchIsoDepotCookieRequestType

type FetchIsoDepotCookieRequestType struct {
	This     types.ManagedObjectReference `xml:"_this"`
	Username string                       `xml:"username"`
	Password string                       `xml:"password"`
}

type GetVsanVumConfigResponse struct {
	Returnval VsanVumSystemConfig `xml:"returnval"`
}

type GetVsanVumConfig GetVsanVumConfigRequestType

type GetVsanVumConfigRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
}

type VsanVcUploadReleaseDbResponse struct {
}

type VsanVcUploadReleaseDb VsanVcUploadReleaseDbRequestType

type VsanVcUploadReleaseDbRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
	Db   string                       `xml:"db"`
}

type VsanVumSystem struct {
	types.ManagedObjectReference
}

type QueryIoInsightResponse struct {
	Returnval []VsanHostIoInsightInfo `xml:"returnval"`
}

type QueryIoInsight QueryIoInsightRequestType

type QueryIoInsightRequestType struct {
	This                types.ManagedObjectReference  `xml:"_this"`
	Cluster             *types.ManagedObjectReference `xml:"cluster,omitempty"`
	HostsIoInsightInfos []VsanHostIoInsightInfo       `xml:"hostsIoInsightInfos,omitempty"`
}

type StartIoInsightResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type StartIoInsight StartIoInsightRequestType

type StartIoInsightRequestType struct {
	This        types.ManagedObjectReference   `xml:"_this"`
	Cluster     *types.ManagedObjectReference  `xml:"cluster,omitempty"`
	RunName     string                         `xml:"runName,omitempty"`
	DurationSec int64                          `xml:"durationSec,omitempty"`
	TargetHosts []types.ManagedObjectReference `xml:"targetHosts,omitempty"`
	TargetVMs   []types.ManagedObjectReference `xml:"targetVMs,omitempty"`
}

type StopIoInsightResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type StopIoInsight StopIoInsightRequestType

type StopIoInsightRequestType struct {
	This                types.ManagedObjectReference  `xml:"_this"`
	Cluster             *types.ManagedObjectReference `xml:"cluster,omitempty"`
	HostsIoInsightInfos []VsanHostIoInsightInfo       `xml:"hostsIoInsightInfos,omitempty"`
}

type VsanIoInsightManager struct {
	types.ManagedObjectReference
}

type VsanGetResourceCheckStatusResponse struct {
	Returnval VsanResourceCheckStatus `xml:"returnval"`
}

type VsanGetResourceCheckStatus VsanGetResourceCheckStatusRequestType

type VsanGetResourceCheckStatusRequestType struct {
	This              types.ManagedObjectReference  `xml:"_this"`
	ResourceCheckSpec *VsanResourceCheckSpec        `xml:"resourceCheckSpec,omitempty"`
	Cluster           *types.ManagedObjectReference `xml:"cluster,omitempty"`
}

type VsanPerformResourceCheckResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type VsanPerformResourceCheck VsanPerformResourceCheckRequestType

type VsanPerformResourceCheckRequestType struct {
	This              types.ManagedObjectReference  `xml:"_this"`
	ResourceCheckSpec VsanResourceCheckSpec         `xml:"resourceCheckSpec"`
	Cluster           *types.ManagedObjectReference `xml:"cluster,omitempty"`
}

type VsanHostCancelResourceCheckResponse struct {
	Returnval bool `xml:"returnval"`
}

type VsanHostCancelResourceCheck VsanHostCancelResourceCheckRequestType

type VsanHostCancelResourceCheckRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
}

type VsanHostPerformResourceCheckResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type VsanHostPerformResourceCheck VsanHostPerformResourceCheckRequestType

type VsanHostPerformResourceCheckRequestType struct {
	This              types.ManagedObjectReference `xml:"_this"`
	ResourceCheckSpec VsanResourceCheckSpec        `xml:"resourceCheckSpec"`
}

type VsanResourceCheckSystem struct {
	types.ManagedObjectReference
}

type UpgradeFormatVersionForDiskMappingResponse struct {
}

type UpgradeFormatVersionForDiskMapping UpgradeFormatVersionForDiskMappingRequestType

type UpgradeFormatVersionForDiskMappingRequestType struct {
	This                   types.ManagedObjectReference `xml:"_this"`
	CacheTierCanonicalName string                       `xml:"cacheTierCanonicalName"`
}

type VsanRealignVsanSparseOfflineResponse struct {
	Returnval VimHostVsanRealignResult `xml:"returnval"`
}

type VsanRealignVsanSparseOffline VsanRealignVsanSparseOfflineRequestType

type VsanRealignVsanSparseOfflineRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Path    string                       `xml:"path"`
	StartAt int64                        `xml:"startAt"`
}

type RetrieveVsanDiskManagementSystemCapabilityResponse struct {
	Returnval VimVsanHostVsanDiskManagementSystemCapability `xml:"returnval"`
}

type RetrieveVsanDiskManagementSystemCapability RetrieveVsanDiskManagementSystemCapabilityRequestType

type RetrieveVsanDiskManagementSystemCapabilityRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
}

type VsanRemoveConcatResponse struct {
	Returnval bool `xml:"returnval"`
}

type VsanRemoveConcat VsanRemoveConcatRequestType

type VsanRemoveConcatRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
}

type QueryHostDiskMappingsResponse struct {
	Returnval []VimVsanHostDiskMapInfoEx `xml:"returnval"`
}

type QueryHostDiskMappings QueryHostDiskMappingsRequestType

type QueryHostDiskMappingsRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
}

type VsanQueryUnalignedStatusResponse struct {
	Returnval VimHostVsanComponentAlignment `xml:"returnval"`
}

type VsanQueryUnalignedStatus VsanQueryUnalignedStatusRequestType

type VsanQueryUnalignedStatusRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
}

type QueryDataEfficiencyCapacityStateResponse struct {
	Returnval VimVsanDataEfficiencyCapacityState `xml:"returnval"`
}

type QueryDataEfficiencyCapacityState QueryDataEfficiencyCapacityStateRequestType

type QueryDataEfficiencyCapacityStateRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
}

type FixVsanObjectsResponse struct {
}

type FixVsanObjects FixVsanObjectsRequestType

type FixVsanObjectsRequestType struct {
	This  types.ManagedObjectReference `xml:"_this"`
	Uuids []string                     `xml:"uuids,omitempty"`
}

type VsanRealignVsanSparseOnlineResponse struct {
	Returnval VimHostVsanRealignResult `xml:"returnval"`
}

type VsanRealignVsanSparseOnline VsanRealignVsanSparseOnlineRequestType

type VsanRealignVsanSparseOnlineRequestType struct {
	This      types.ManagedObjectReference `xml:"_this"`
	UuidChain []string                     `xml:"uuidChain"`
	StartAt   int64                        `xml:"startAt"`
}

type UpdateCapacityFlashStatusForDisksResponse struct {
	Returnval bool `xml:"returnval"`
}

type UpdateCapacityFlashStatusForDisks UpdateCapacityFlashStatusForDisksRequestType

type UpdateCapacityFlashStatusForDisksRequestType struct {
	This           types.ManagedObjectReference `xml:"_this"`
	CanonicalNames []string                     `xml:"canonicalNames"`
	Enabled        bool                         `xml:"enabled"`
}

type VsanAlignObjectSizeResponse struct {
	Returnval bool `xml:"returnval"`
}

type VsanAlignObjectSize VsanAlignObjectSizeRequestType

type VsanAlignObjectSizeRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
	Uuid string                       `xml:"uuid"`
}

type VsanQueryNamespaceUuidsResponse struct {
	Returnval []string `xml:"returnval"`
}

type VsanQueryNamespaceUuids VsanQueryNamespaceUuidsRequestType

type VsanQueryNamespaceUuidsRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
}

type VsanScanObjects_TaskResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type VsanScanObjects_Task VsanScanObjects_TaskRequestType

type VsanScanObjects_TaskRequestType struct {
	This       types.ManagedObjectReference `xml:"_this"`
	Version    int64                        `xml:"version,omitempty"`
	Fixobjects *bool                        `xml:"fixobjects"`
}

type VsanGetOpenV2VsanSparseChainResponse struct {
	Returnval []string `xml:"returnval"`
}

type VsanGetOpenV2VsanSparseChain VsanGetOpenV2VsanSparseChainRequestType

type VsanGetOpenV2VsanSparseChainRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
}

type VsanScanNamespaceResponse struct {
	Returnval []string `xml:"returnval"`
}

type VsanScanNamespace VsanScanNamespaceRequestType

type VsanScanNamespaceRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
	Uuid string                       `xml:"uuid"`
}

type VimHostVsanDiskManagementSystem struct {
	types.ManagedObjectReference
}

type VsanVibInstall_TaskResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type VsanVibInstall_Task VsanVibInstall_TaskRequestType

type VsanVibInstall_TaskRequestType struct {
	This            types.ManagedObjectReference  `xml:"_this"`
	Cluster         *types.ManagedObjectReference `xml:"cluster,omitempty"`
	VibSpecs        []VsanVibSpec                 `xml:"vibSpecs,omitempty"`
	ScanResults     []VsanVibScanResult           `xml:"scanResults,omitempty"`
	FirmwareSpecs   []VsanHclFirmwareUpdateSpec   `xml:"firmwareSpecs,omitempty"`
	MaintenanceSpec *types.HostMaintenanceSpec    `xml:"maintenanceSpec,omitempty"`
	Rolling         *bool                         `xml:"rolling"`
	NoSigCheck      *bool                         `xml:"noSigCheck"`
}

type VsanVibInstallPreflightCheckResponse struct {
	Returnval VsanVibInstallPreflightStatus `xml:"returnval"`
}

type VsanVibInstallPreflightCheck VsanVibInstallPreflightCheckRequestType

type VsanVibInstallPreflightCheckRequestType struct {
	This    types.ManagedObjectReference  `xml:"_this"`
	Cluster *types.ManagedObjectReference `xml:"cluster,omitempty"`
}

type VsanVibScanResponse struct {
	Returnval []VsanVibScanResult `xml:"returnval"`
}

type VsanVibScan VsanVibScanRequestType

type VsanVibScanRequestType struct {
	This     types.ManagedObjectReference  `xml:"_this"`
	Cluster  *types.ManagedObjectReference `xml:"cluster,omitempty"`
	VibSpecs []VsanVibSpec                 `xml:"vibSpecs"`
}

type VsanUpdateManager struct {
	types.ManagedObjectReference
}

type VsanGetTaskStatusResponse struct {
	Returnval []types.ObjectContent `xml:"returnval"`
}

type VsanGetTaskStatus VsanGetTaskStatusRequestType

type VsanGetTaskStatusRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
}

type VsanTaskTracker struct {
	types.ManagedObjectReference
}

type RetrieveLookupTableResponse struct {
	Returnval string `xml:"returnval"`
}

type RetrieveLookupTable RetrieveLookupTableRequestType

type RetrieveLookupTableRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
}

type VsanRetrievePropertiesResponse struct {
	Returnval []types.ObjectContent `xml:"returnval"`
}

type VsanRetrieveProperties VsanRetrievePropertiesRequestType

type VsanRetrievePropertiesRequestType struct {
	This               types.ManagedObjectReference `xml:"_this"`
	MassCollectorSpecs []VsanMassCollectorSpec      `xml:"massCollectorSpecs"`
}

type VsanRetrievePropertiesJsonResponse struct {
	Returnval string `xml:"returnval"`
}

type VsanRetrievePropertiesJson VsanRetrievePropertiesJsonRequestType

type VsanRetrievePropertiesJsonRequestType struct {
	This               types.ManagedObjectReference `xml:"_this"`
	MassCollectorSpecs []VsanMassCollectorSpec      `xml:"massCollectorSpecs"`
}

type VsanMassCollector struct {
	types.ManagedObjectReference
}

type VsanUnmountDiskMappingExResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type VsanUnmountDiskMappingEx VsanUnmountDiskMappingExRequestType

type VsanUnmountDiskMappingExRequestType struct {
	This            types.ManagedObjectReference `xml:"_this"`
	Mappings        []types.VsanHostDiskMapping  `xml:"mappings"`
	MaintenanceSpec *types.HostMaintenanceSpec   `xml:"maintenanceSpec,omitempty"`
	Timeout         int32                        `xml:"timeout,omitempty"`
}

type QueryVsanAwsCredentialsResponse struct {
	Returnval VsanHostVsanAwsCredentials `xml:"returnval"`
}

type QueryVsanAwsCredentials QueryVsanAwsCredentialsRequestType

type QueryVsanAwsCredentialsRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
}

type VsanQuerySyncingVsanObjectsResponse struct {
	Returnval VsanHostVsanObjectSyncQueryResult `xml:"returnval"`
}

type VsanQuerySyncingVsanObjects VsanQuerySyncingVsanObjectsRequestType

type VsanQuerySyncingVsanObjectsRequestType struct {
	This           types.ManagedObjectReference `xml:"_this"`
	Uuids          []string                     `xml:"uuids,omitempty"`
	Start          int32                        `xml:"start,omitempty"`
	Limit          int32                        `xml:"limit,omitempty"`
	IncludeSummary *bool                        `xml:"includeSummary"`
}

type RefreshVsanSystemResponse struct {
}

type RefreshVsanSystem RefreshVsanSystemRequestType

type RefreshVsanSystemRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
}

type UpdateVmMetadataResponse struct {
	Returnval []types.VirtualMachineMetadataManagerVmMetadataResult `xml:"returnval"`
}

type UpdateVmMetadata UpdateVmMetadataRequestType

type UpdateVmMetadataRequestType struct {
	This              types.ManagedObjectReference                         `xml:"_this"`
	VmMetadataInput   []types.VirtualMachineMetadataManagerVmMetadataInput `xml:"vmMetadataInput"`
	UpdateTimeoutMsec int32                                                `xml:"updateTimeoutMsec"`
}

type VsanQueryHostDrsStatsResponse struct {
	Returnval VsanHostDrsStats `xml:"returnval"`
}

type VsanQueryHostDrsStats VsanQueryHostDrsStatsRequestType

type VsanQueryHostDrsStatsRequestType struct {
	This      types.ManagedObjectReference `xml:"_this"`
	HostUuids []string                     `xml:"hostUuids"`
	Vms       []string                     `xml:"vms,omitempty"`
}

type VmMonitorUnwatchObjectsResponse struct {
}

type VmMonitorUnwatchObjects VmMonitorUnwatchObjectsRequestType

type VmMonitorUnwatchObjectsRequestType struct {
	This     types.ManagedObjectReference `xml:"_this"`
	ObjUuids []string                     `xml:"objUuids"`
}

type VsanExitMaintenanceMode_TaskResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type VsanExitMaintenanceMode_Task VsanExitMaintenanceMode_TaskRequestType

type VsanExitMaintenanceMode_TaskRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Timeout int32                        `xml:"timeout,omitempty"`
}

type VsanGetLocalAffinityVmListResponse struct {
	Returnval []string `xml:"returnval"`
}

type VsanGetLocalAffinityVmList VsanGetLocalAffinityVmListRequestType

type VsanGetLocalAffinityVmListRequestType struct {
	This                 types.ManagedObjectReference `xml:"_this"`
	IncludePoweredOffVms *bool                        `xml:"includePoweredOffVms"`
}

type VsanGetAssociatedSpbmProfileResponse struct {
	Returnval []VsanObjectProfileInfo `xml:"returnval"`
}

type VsanGetAssociatedSpbmProfile VsanGetAssociatedSpbmProfileRequestType

type VsanGetAssociatedSpbmProfileRequestType struct {
	This            types.ManagedObjectReference `xml:"_this"`
	VsanObjectUuids []string                     `xml:"vsanObjectUuids"`
}

type VsanQueryWhatIfEvacuationResultResponse struct {
	Returnval VsanWhatIfEvacResult `xml:"returnval"`
}

type VsanQueryWhatIfEvacuationResult VsanQueryWhatIfEvacuationResultRequestType

type VsanQueryWhatIfEvacuationResultRequestType struct {
	This           types.ManagedObjectReference `xml:"_this"`
	EvacEntityUuid string                       `xml:"evacEntityUuid"`
}

type ClearVmMetadataResponse struct {
	Returnval []types.VirtualMachineMetadataManagerVmMetadataResult `xml:"returnval"`
}

type ClearVmMetadata ClearVmMetadataRequestType

type ClearVmMetadataRequestType struct {
	This             types.ManagedObjectReference `xml:"_this"`
	ClearTimeoutMsec int32                        `xml:"clearTimeoutMsec"`
}

type GetVsanNicsResponse struct {
	Returnval []types.KeyAnyValue `xml:"returnval"`
}

type GetVsanNics GetVsanNicsRequestType

type GetVsanNicsRequestType struct {
	This        types.ManagedObjectReference `xml:"_this"`
	TrafficType string                       `xml:"trafficType,omitempty"`
}

type VsanHostGetRuntimeStatsResponse struct {
	Returnval VsanHostRuntimeStats `xml:"returnval"`
}

type VsanHostGetRuntimeStats VsanHostGetRuntimeStatsRequestType

type VsanHostGetRuntimeStatsRequestType struct {
	This  types.ManagedObjectReference `xml:"_this"`
	Stats []string                     `xml:"stats,omitempty"`
}

type VsanGetAssociatedObjectsResponse struct {
	Returnval VsanHostAssociatedObjectsResult `xml:"returnval"`
}

type VsanGetAssociatedObjects VsanGetAssociatedObjectsRequestType

type VsanGetAssociatedObjectsRequestType struct {
	This                 types.ManagedObjectReference `xml:"_this"`
	SpbmProfileId        string                       `xml:"spbmProfileId"`
	SpbmGenerationNumber int32                        `xml:"spbmGenerationNumber,omitempty"`
	Start                int32                        `xml:"start,omitempty"`
	Limit                int32                        `xml:"limit,omitempty"`
}

type VmMonitorWatchObjectsResponse struct {
}

type VmMonitorWatchObjects VmMonitorWatchObjectsRequestType

type VmMonitorWatchObjectsRequestType struct {
	This     types.ManagedObjectReference `xml:"_this"`
	ObjUuids []string                     `xml:"objUuids"`
}

type QueryVmsAccessibilityResponse struct {
	Returnval []types.DynamicData `xml:"returnval"`
}

type QueryVmsAccessibility QueryVmsAccessibilityRequestType

type QueryVmsAccessibilityRequestType struct {
	This  types.ManagedObjectReference `xml:"_this"`
	VmIds []string                     `xml:"vmIds"`
}

type GetDomObjectsResponse struct {
	Returnval []string `xml:"returnval"`
}

type GetDomObjects GetDomObjectsRequestType

type GetDomObjectsRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
}

type VsanEnterMaintenanceMode_TaskResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type VsanEnterMaintenanceMode_Task VsanEnterMaintenanceMode_TaskRequestType

type VsanEnterMaintenanceMode_TaskRequestType struct {
	This            types.ManagedObjectReference `xml:"_this"`
	MaintenanceSpec *types.HostMaintenanceSpec   `xml:"maintenanceSpec,omitempty"`
	Timeout         int32                        `xml:"timeout,omitempty"`
}

type GetTopLevelDirectoriesResponse struct {
	Returnval []types.KeyValue `xml:"returnval"`
}

type GetTopLevelDirectories GetTopLevelDirectoriesRequestType

type GetTopLevelDirectoriesRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
	Cid  string                       `xml:"cid,omitempty"`
}

type GetVsanRuntimeInfoResponse struct {
	Returnval types.VsanHostRuntimeInfo `xml:"returnval"`
}

type GetVsanRuntimeInfo GetVsanRuntimeInfoRequestType

type GetVsanRuntimeInfoRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
}

type GetDomClientRoleStatsResponse struct {
	Returnval []types.KeyValue `xml:"returnval"`
}

type GetDomClientRoleStats GetDomClientRoleStatsRequestType

type GetDomClientRoleStatsRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
}

type VsanGetAboutInfoExResponse struct {
	Returnval VsanHostAboutInfoEx `xml:"returnval"`
}

type VsanGetAboutInfoEx VsanGetAboutInfoExRequestType

type VsanGetAboutInfoExRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
}

type QueryVmMetadataResponse struct {
	Returnval []types.VirtualMachineMetadataManagerVmMetadataResult `xml:"returnval"`
}

type QueryVmMetadata QueryVmMetadataRequestType

type QueryVmMetadataRequestType struct {
	This  types.ManagedObjectReference `xml:"_this"`
	VmIds []string                     `xml:"vmIds,omitempty"`
	DsUrl string                       `xml:"dsUrl,omitempty"`
}

type RefreshVsanVmknicsResponse struct {
}

type RefreshVsanVmknics RefreshVsanVmknicsRequestType

type RefreshVsanVmknicsRequestType struct {
	This   types.ManagedObjectReference `xml:"_this"`
	Vmknic string                       `xml:"vmknic,omitempty"`
}

type GetDomObjectStatsResponse struct {
	Returnval []types.KeyValue `xml:"returnval"`
}

type GetDomObjectStats GetDomObjectStatsRequestType

type GetDomObjectStatsRequestType struct {
	This   types.ManagedObjectReference `xml:"_this"`
	DomObj string                       `xml:"domObj"`
}

type QueryObjectsAccessibilityResponse struct {
	Returnval []types.DynamicData `xml:"returnval"`
}

type QueryObjectsAccessibility QueryObjectsAccessibilityRequestType

type QueryObjectsAccessibilityRequestType struct {
	This        types.ManagedObjectReference `xml:"_this"`
	ObjectUUIDs []string                     `xml:"objectUUIDs"`
}

type VsanSystemEx struct {
	types.ManagedObjectReference
}

type VsanClusterReconfigResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type VsanClusterReconfig VsanClusterReconfigRequestType

type VsanClusterReconfigRequestType struct {
	This             types.ManagedObjectReference `xml:"_this"`
	Cluster          types.ManagedObjectReference `xml:"cluster"`
	VsanReconfigSpec VimVsanReconfigSpec          `xml:"vsanReconfigSpec"`
}

type VsanGetDataProtectionLoadBalancersResponse struct {
	Returnval DataProtectionLoadBalancersInfo `xml:"returnval"`
}

type VsanGetDataProtectionLoadBalancers VsanGetDataProtectionLoadBalancersRequestType

type VsanGetDataProtectionLoadBalancersRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Cluster types.ManagedObjectReference `xml:"cluster"`
}

type VsanClusterGetRuntimeStatsResponse struct {
	Returnval []VsanRuntimeStatsHostMap `xml:"returnval"`
}

type VsanClusterGetRuntimeStats VsanClusterGetRuntimeStatsRequestType

type VsanClusterGetRuntimeStatsRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Cluster types.ManagedObjectReference `xml:"cluster"`
	Stats   []string                     `xml:"stats,omitempty"`
}

type VsanQueryClusterDrsStatsResponse struct {
	Returnval []VsanHostDrsStats `xml:"returnval"`
}

type VsanQueryClusterDrsStats VsanQueryClusterDrsStatsRequestType

type VsanQueryClusterDrsStatsRequestType struct {
	This    types.ManagedObjectReference   `xml:"_this"`
	Cluster types.ManagedObjectReference   `xml:"cluster"`
	Vms     []types.ManagedObjectReference `xml:"vms,omitempty"`
}

type VsanFindClusterByDatastoreUrlResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type VsanFindClusterByDatastoreUrl VsanFindClusterByDatastoreUrlRequestType

type VsanFindClusterByDatastoreUrlRequestType struct {
	This         types.ManagedObjectReference `xml:"_this"`
	DatastoreUrl string                       `xml:"datastoreUrl"`
}

type VsanClusterGetConfigResponse struct {
	Returnval VsanConfigInfoEx `xml:"returnval"`
}

type VsanClusterGetConfig VsanClusterGetConfigRequestType

type VsanClusterGetConfigRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Cluster types.ManagedObjectReference `xml:"cluster"`
}

type VsanEncryptedClusterRekey_TaskResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type VsanEncryptedClusterRekey_Task VsanEncryptedClusterRekey_TaskRequestType

type VsanEncryptedClusterRekey_TaskRequestType struct {
	This                   types.ManagedObjectReference `xml:"_this"`
	EncryptedCluster       types.ManagedObjectReference `xml:"encryptedCluster"`
	DeepRekey              *bool                        `xml:"deepRekey"`
	AllowReducedRedundancy *bool                        `xml:"allowReducedRedundancy"`
}

type VsanVcClusterConfigSystem struct {
	types.ManagedObjectReference
}

type VsanFlrSessionQueryVolumesResponse struct {
	Returnval VsanFlrQueryVolumesResult `xml:"returnval"`
}

type VsanFlrSessionQueryVolumes VsanFlrSessionQueryVolumesRequestType

type VsanFlrSessionQueryVolumesRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
}

type VsanFlrSessionRetrieveFileResponse struct {
	Returnval VsanFlrRetrieveFileResult `xml:"returnval"`
}

type VsanFlrSessionRetrieveFile VsanFlrSessionRetrieveFileRequestType

type VsanFlrSessionRetrieveFileRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
	Spec VsanFlrRetrieveFileSpec      `xml:"spec"`
}

type VsanFlrSessionQueryFileResponse struct {
	Returnval VsanFlrQueryFileResult `xml:"returnval"`
}

type VsanFlrSessionQueryFile VsanFlrSessionQueryFileRequestType

type VsanFlrSessionQueryFileRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
	Spec VsanFlrQueryFileSpec         `xml:"spec"`
}

type VsanFlrSession struct {
	types.ManagedObjectReference
}

type VsanDownloadFileServiceOvfResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type VsanDownloadFileServiceOvf VsanDownloadFileServiceOvfRequestType

type VsanDownloadFileServiceOvfRequestType struct {
	This        types.ManagedObjectReference `xml:"_this"`
	DownloadUrl string                       `xml:"downloadUrl"`
}

type VsanClusterCreateFsDomainResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type VsanClusterCreateFsDomain VsanClusterCreateFsDomainRequestType

type VsanClusterCreateFsDomainRequestType struct {
	This         types.ManagedObjectReference  `xml:"_this"`
	DomainConfig VsanFileServiceDomainConfig   `xml:"domainConfig"`
	Cluster      *types.ManagedObjectReference `xml:"cluster,omitempty"`
}

type VsanPerformFileServiceEnablePreflightCheckResponse struct {
	Returnval VsanFileServicePreflightCheckResult `xml:"returnval"`
}

type VsanPerformFileServiceEnablePreflightCheck VsanPerformFileServiceEnablePreflightCheckRequestType

type VsanPerformFileServiceEnablePreflightCheckRequestType struct {
	This         types.ManagedObjectReference `xml:"_this"`
	Cluster      types.ManagedObjectReference `xml:"cluster"`
	DomainConfig *VsanFileServiceDomainConfig `xml:"domainConfig,omitempty"`
}

type VsanRemediateFileServiceResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type VsanRemediateFileService VsanRemediateFileServiceRequestType

type VsanRemediateFileServiceRequestType struct {
	This         types.ManagedObjectReference `xml:"_this"`
	Config       VsanFileServiceConfig        `xml:"config"`
	CreateRootFs *bool                        `xml:"createRootFs"`
}

type VsanClusterQueryFsDomainsResponse struct {
	Returnval []VsanFileServiceDomain `xml:"returnval"`
}

type VsanClusterQueryFsDomains VsanClusterQueryFsDomainsRequestType

type VsanClusterQueryFsDomainsRequestType struct {
	This      types.ManagedObjectReference    `xml:"_this"`
	QuerySpec *VsanFileServiceDomainQuerySpec `xml:"querySpec,omitempty"`
	Cluster   *types.ManagedObjectReference   `xml:"cluster,omitempty"`
}

type VsanQueryFileServiceOvfsResponse struct {
	Returnval []VsanFileServiceOvfSpec `xml:"returnval"`
}

type VsanQueryFileServiceOvfs VsanQueryFileServiceOvfsRequestType

type VsanQueryFileServiceOvfsRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
}

type VsanReconfigureFileShareResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type VsanReconfigureFileShare VsanReconfigureFileShareRequestType

type VsanReconfigureFileShareRequestType struct {
	This      types.ManagedObjectReference  `xml:"_this"`
	ShareUuid string                        `xml:"shareUuid"`
	Config    VsanFileShareConfig           `xml:"config"`
	Cluster   *types.ManagedObjectReference `xml:"cluster,omitempty"`
}

type VsanClusterReconfigureFsDomainResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type VsanClusterReconfigureFsDomain VsanClusterReconfigureFsDomainRequestType

type VsanClusterReconfigureFsDomainRequestType struct {
	This         types.ManagedObjectReference  `xml:"_this"`
	DomainUuid   string                        `xml:"domainUuid"`
	DomainConfig VsanFileServiceDomainConfig   `xml:"domainConfig"`
	Cluster      *types.ManagedObjectReference `xml:"cluster,omitempty"`
}

type VsanFindOvfDownloadUrlResponse struct {
	Returnval string `xml:"returnval"`
}

type VsanFindOvfDownloadUrl VsanFindOvfDownloadUrlRequestType

type VsanFindOvfDownloadUrlRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Cluster types.ManagedObjectReference `xml:"cluster"`
}

type VsanClusterQueryFileSharesResponse struct {
	Returnval []VsanFileShare `xml:"returnval"`
}

type VsanClusterQueryFileShares VsanClusterQueryFileSharesRequestType

type VsanClusterQueryFileSharesRequestType struct {
	This      types.ManagedObjectReference  `xml:"_this"`
	QuerySpec VsanFileShareQuerySpec        `xml:"querySpec"`
	Cluster   *types.ManagedObjectReference `xml:"cluster,omitempty"`
}

type VsanClusterRemoveShareResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type VsanClusterRemoveShare VsanClusterRemoveShareRequestType

type VsanClusterRemoveShareRequestType struct {
	This      types.ManagedObjectReference  `xml:"_this"`
	ShareUuid string                        `xml:"shareUuid"`
	Cluster   *types.ManagedObjectReference `xml:"cluster,omitempty"`
}

type VsanClusterRemoveFsDomainResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type VsanClusterRemoveFsDomain VsanClusterRemoveFsDomainRequestType

type VsanClusterRemoveFsDomainRequestType struct {
	This       types.ManagedObjectReference  `xml:"_this"`
	DomainUuid string                        `xml:"domainUuid"`
	Cluster    *types.ManagedObjectReference `xml:"cluster,omitempty"`
}

type VsanCreateFileShareResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type VsanCreateFileShare VsanCreateFileShareRequestType

type VsanCreateFileShareRequestType struct {
	This    types.ManagedObjectReference  `xml:"_this"`
	Config  VsanFileShareConfig           `xml:"config"`
	Cluster *types.ManagedObjectReference `xml:"cluster,omitempty"`
}

type VsanFileServiceSystem struct {
	types.ManagedObjectReference
}

type VsanHostQueryAdvCfgResponse struct {
	Returnval []types.OptionValue `xml:"returnval"`
}

type VsanHostQueryAdvCfg VsanHostQueryAdvCfgRequestType

type VsanHostQueryAdvCfgRequestType struct {
	This                 types.ManagedObjectReference `xml:"_this"`
	Options              []string                     `xml:"options"`
	IncludeAllAdvOptions *bool                        `xml:"includeAllAdvOptions"`
	NonDefaultOnly       *bool                        `xml:"nonDefaultOnly"`
}

type VsanHostQueryRunIperfClientResponse struct {
	Returnval VsanNetworkLoadTestResult `xml:"returnval"`
}

type VsanHostQueryRunIperfClient VsanHostQueryRunIperfClientRequestType

type VsanHostQueryRunIperfClientRequestType struct {
	This        types.ManagedObjectReference `xml:"_this"`
	Multicast   bool                         `xml:"multicast"`
	ServerIp    string                       `xml:"serverIp"`
	DurationSec int32                        `xml:"durationSec,omitempty"`
}

type VsanQueryVsanConfigsByFilterResponse struct {
	Returnval VsanHostConfigInfoEx `xml:"returnval"`
}

type VsanQueryVsanConfigsByFilter VsanQueryVsanConfigsByFilterRequestType

type VsanQueryVsanConfigsByFilterRequestType struct {
	This            types.ManagedObjectReference `xml:"_this"`
	RequiredConfigs []string                     `xml:"requiredConfigs"`
}

type VsanHostQueryObjectHealthSummaryResponse struct {
	Returnval VsanObjectOverallHealth `xml:"returnval"`
}

type VsanHostQueryObjectHealthSummary VsanHostQueryObjectHealthSummaryRequestType

type VsanHostQueryObjectHealthSummaryRequestType struct {
	This                          types.ManagedObjectReference `xml:"_this"`
	ObjUuids                      []string                     `xml:"objUuids,omitempty"`
	IncludeObjUuids               *bool                        `xml:"includeObjUuids"`
	LocalHostOnly                 *bool                        `xml:"localHostOnly"`
	IncludeDataProtectionHealth   *bool                        `xml:"includeDataProtectionHealth"`
	IncludeNonComplianceObjDetail *bool                        `xml:"includeNonComplianceObjDetail"`
}

type RefreshVsanHealthGenerationIdResponse struct {
	Returnval int32 `xml:"returnval"`
}

type RefreshVsanHealthGenerationId RefreshVsanHealthGenerationIdRequestType

type RefreshVsanHealthGenerationIdRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
}

type VsanStopProactiveRebalanceResponse struct {
	Returnval bool `xml:"returnval"`
}

type VsanStopProactiveRebalance VsanStopProactiveRebalanceRequestType

type VsanStopProactiveRebalanceRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
}

type VsanLocalPropertyCollectorResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type VsanLocalPropertyCollector VsanLocalPropertyCollectorRequestType

type VsanLocalPropertyCollectorRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
}

type VsanHostQueryFileServiceHealthSummaryResponse struct {
	Returnval VsanFileServiceHealthSummary `xml:"returnval"`
}

type VsanHostQueryFileServiceHealthSummary VsanHostQueryFileServiceHealthSummaryRequestType

type VsanHostQueryFileServiceHealthSummaryRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
}

type VsanHostQueryDataProtectionCfgResponse struct {
	Returnval []types.OptionValue `xml:"returnval"`
}

type VsanHostQueryDataProtectionCfg VsanHostQueryDataProtectionCfgRequestType

type VsanHostQueryDataProtectionCfgRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
}

type VsanLocalHostSystemResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type VsanLocalHostSystem VsanLocalHostSystemRequestType

type VsanLocalHostSystemRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
}

type VsanHostClomdLivenessResponse struct {
	Returnval bool `xml:"returnval"`
}

type VsanHostClomdLiveness VsanHostClomdLivenessRequestType

type VsanHostClomdLivenessRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
}

type VsanHostRepairImmediateObjectsResponse struct {
	Returnval VsanRepairObjectsResult `xml:"returnval"`
}

type VsanHostRepairImmediateObjects VsanHostRepairImmediateObjectsRequestType

type VsanHostRepairImmediateObjectsRequestType struct {
	This       types.ManagedObjectReference `xml:"_this"`
	Uuids      []string                     `xml:"uuids,omitempty"`
	RepairType string                       `xml:"repairType,omitempty"`
}

type VsanHostQueryVerifyNetworkSettingsResponse struct {
	Returnval VsanNetworkHealthResult `xml:"returnval"`
}

type VsanHostQueryVerifyNetworkSettings VsanHostQueryVerifyNetworkSettingsRequestType

type VsanHostQueryVerifyNetworkSettingsRequestType struct {
	This                          types.ManagedObjectReference `xml:"_this"`
	Peers                         []string                     `xml:"peers,omitempty"`
	ROBOStretchedClusterWitnesses []string                     `xml:"ROBOStretchedClusterWitnesses,omitempty"`
	VMotionPeers                  []string                     `xml:"vMotionPeers,omitempty"`
}

type VsanHostCleanupVmdkLoadTestResponse struct {
	Returnval string `xml:"returnval"`
}

type VsanHostCleanupVmdkLoadTest VsanHostCleanupVmdkLoadTestRequestType

type VsanHostCleanupVmdkLoadTestRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Runname string                       `xml:"runname"`
	Specs   []VsanVmdkLoadTestSpec       `xml:"specs,omitempty"`
}

type VsanStartProactiveRebalanceResponse struct {
	Returnval bool `xml:"returnval"`
}

type VsanStartProactiveRebalance VsanStartProactiveRebalanceRequestType

type VsanStartProactiveRebalanceRequestType struct {
	This              types.ManagedObjectReference `xml:"_this"`
	TimeSpan          int32                        `xml:"timeSpan,omitempty"`
	VarianceThreshold float32                      `xml:"varianceThreshold,omitempty"`
	TimeThreshold     int32                        `xml:"timeThreshold,omitempty"`
	RateThreshold     int32                        `xml:"rateThreshold,omitempty"`
}

type VsanHostQueryEncryptionHealthSummaryResponse struct {
	Returnval VsanEncryptionHealthSummary `xml:"returnval"`
}

type VsanHostQueryEncryptionHealthSummary VsanHostQueryEncryptionHealthSummaryRequestType

type VsanHostQueryEncryptionHealthSummaryRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
}

type VsanFlashScsiControllerFirmware_TaskResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type VsanFlashScsiControllerFirmware_Task VsanFlashScsiControllerFirmware_TaskRequestType

type VsanFlashScsiControllerFirmware_TaskRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
	Spec VsanHclFirmwareUpdateSpec    `xml:"spec"`
}

type PrepareHostForClusterRebootWithNAMMResponse struct {
	Returnval bool `xml:"returnval"`
}

type PrepareHostForClusterRebootWithNAMM PrepareHostForClusterRebootWithNAMMRequestType

type PrepareHostForClusterRebootWithNAMMRequestType struct {
	This         types.ManagedObjectReference `xml:"_this"`
	VerifyOnly   bool                         `xml:"verifyOnly"`
	ScheduleTime int32                        `xml:"scheduleTime,omitempty"`
}

type VsanQueryHostEMMStateResponse struct {
	Returnval VsanHostEMMSummary `xml:"returnval"`
}

type VsanQueryHostEMMState VsanQueryHostEMMStateRequestType

type VsanQueryHostEMMStateRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
}

type VsanRemediateScsiControllerIssuesResponse struct {
}

type VsanRemediateScsiControllerIssues VsanRemediateScsiControllerIssuesRequestType

type VsanRemediateScsiControllerIssuesRequestType struct {
	This            types.ManagedObjectReference `xml:"_this"`
	HbaDevice       string                       `xml:"hbaDevice"`
	RemediateIssues []string                     `xml:"remediateIssues,omitempty"`
}

type VsanHostQueryVmProcessListResponse struct {
	Returnval []string `xml:"returnval"`
}

type VsanHostQueryVmProcessList VsanHostQueryVmProcessListRequestType

type VsanHostQueryVmProcessListRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
}

type VsanWaitForVsanHealthGenerationIdChangeResponse struct {
	Returnval bool `xml:"returnval"`
}

type VsanWaitForVsanHealthGenerationIdChange VsanWaitForVsanHealthGenerationIdChangeRequestType

type VsanWaitForVsanHealthGenerationIdChangeRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Timeout int32                        `xml:"timeout"`
}

type VsanHostQueryHealthSystemVersionResponse struct {
	Returnval string `xml:"returnval"`
}

type VsanHostQueryHealthSystemVersion VsanHostQueryHealthSystemVersionRequestType

type VsanHostQueryHealthSystemVersionRequestType struct {
	This           types.ManagedObjectReference `xml:"_this"`
	DisplayVersion *bool                        `xml:"displayVersion"`
}

type VsanGetHclInfoResponse struct {
	Returnval VsanHostHclInfo `xml:"returnval"`
}

type VsanGetHclInfo VsanGetHclInfoRequestType

type VsanGetHclInfoRequestType struct {
	This              types.ManagedObjectReference `xml:"_this"`
	IncludeVendorInfo *bool                        `xml:"includeVendorInfo"`
}

type VsanHostRunVmdkLoadTestResponse struct {
	Returnval []VsanVmdkLoadTestResult `xml:"returnval"`
}

type VsanHostRunVmdkLoadTest VsanHostRunVmdkLoadTestRequestType

type VsanHostRunVmdkLoadTestRequestType struct {
	This        types.ManagedObjectReference `xml:"_this"`
	Runname     string                       `xml:"runname"`
	DurationSec int32                        `xml:"durationSec"`
	Specs       []VsanVmdkLoadTestSpec       `xml:"specs"`
}

type RestoreHostForClusterRebootWithNAMMResponse struct {
	Returnval bool `xml:"returnval"`
}

type RestoreHostForClusterRebootWithNAMM RestoreHostForClusterRebootWithNAMMRequestType

type RestoreHostForClusterRebootWithNAMMRequestType struct {
	This         types.ManagedObjectReference `xml:"_this"`
	VerifyOnly   bool                         `xml:"verifyOnly"`
	ScheduleTime int32                        `xml:"scheduleTime,omitempty"`
}

type VsanHostQuerySmartStatsResponse struct {
	Returnval VsanSmartStatsHostSummary `xml:"returnval"`
}

type VsanHostQuerySmartStats VsanHostQuerySmartStatsRequestType

type VsanHostQuerySmartStatsRequestType struct {
	This            types.ManagedObjectReference `xml:"_this"`
	Disks           []string                     `xml:"disks,omitempty"`
	IncludeAllDisks *bool                        `xml:"includeAllDisks"`
}

type VsanHostPrepareVmdkLoadTestResponse struct {
	Returnval string `xml:"returnval"`
}

type VsanHostPrepareVmdkLoadTest VsanHostPrepareVmdkLoadTestRequestType

type VsanHostPrepareVmdkLoadTestRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Runname string                       `xml:"runname"`
	Specs   []VsanVmdkLoadTestSpec       `xml:"specs"`
}

type VsanHostQueryRunIperfServerResponse struct {
	Returnval VsanNetworkLoadTestResult `xml:"returnval"`
}

type VsanHostQueryRunIperfServer VsanHostQueryRunIperfServerRequestType

type VsanHostQueryRunIperfServerRequestType struct {
	This        types.ManagedObjectReference `xml:"_this"`
	Multicast   bool                         `xml:"multicast"`
	ServerIp    string                       `xml:"serverIp,omitempty"`
	DurationSec int32                        `xml:"durationSec,omitempty"`
}

type VsanGetProactiveRebalanceInfoResponse struct {
	Returnval VsanProactiveRebalanceInfoEx `xml:"returnval"`
}

type VsanGetProactiveRebalanceInfo VsanGetProactiveRebalanceInfoRequestType

type VsanGetProactiveRebalanceInfoRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
}

type VsanHostQueryPhysicalDiskHealthSummaryResponse struct {
	Returnval VsanPhysicalDiskHealthSummary `xml:"returnval"`
}

type VsanHostQueryPhysicalDiskHealthSummary VsanHostQueryPhysicalDiskHealthSummaryRequestType

type VsanHostQueryPhysicalDiskHealthSummaryRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
}

type VsanHostQueryVsanPcapResponse struct {
	Returnval VsanVsanPcapResult `xml:"returnval"`
}

type VsanHostQueryVsanPcap VsanHostQueryVsanPcapRequestType

type VsanHostQueryVsanPcapRequestType struct {
	This               types.ManagedObjectReference `xml:"_this"`
	Vmknic             string                       `xml:"vmknic"`
	Duration           int32                        `xml:"duration"`
	IncludeRawPcap     *bool                        `xml:"includeRawPcap"`
	IncludeIgmp        *bool                        `xml:"includeIgmp"`
	CmmdsMsgTypeFilter []string                     `xml:"cmmdsMsgTypeFilter,omitempty"`
	CmmdsPorts         []int32                      `xml:"cmmdsPorts,omitempty"`
	ClusterUuid        string                       `xml:"clusterUuid,omitempty"`
}

type VsanHostQueryHostInfoByUuidsResponse struct {
	Returnval []VsanQueryResultHostInfo `xml:"returnval"`
}

type VsanHostQueryHostInfoByUuids VsanHostQueryHostInfoByUuidsRequestType

type VsanHostQueryHostInfoByUuidsRequestType struct {
	This  types.ManagedObjectReference `xml:"_this"`
	Uuids []string                     `xml:"uuids"`
}

type VsanCheckArchivalAccessibilityResponse struct {
	Returnval VsanArchivalAccessibilityResult `xml:"returnval"`
}

type VsanCheckArchivalAccessibility VsanCheckArchivalAccessibilityRequestType

type VsanCheckArchivalAccessibilityRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
}

type VsanQueryDiskRebalanceStatusResponse struct {
	Returnval []VsanDiskRebalanceResult `xml:"returnval"`
}

type VsanQueryDiskRebalanceStatus VsanQueryDiskRebalanceStatusRequestType

type VsanQueryDiskRebalanceStatusRequestType struct {
	This         types.ManagedObjectReference `xml:"_this"`
	DiskUuidList []string                     `xml:"diskUuidList,omitempty"`
}

type VsanHostCreateVmHealthTestResponse struct {
	Returnval VsanHostCreateVmHealthTestResult `xml:"returnval"`
}

type VsanHostCreateVmHealthTest VsanHostCreateVmHealthTestRequestType

type VsanHostCreateVmHealthTestRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Timeout int32                        `xml:"timeout"`
}

type VsanHostDpdLivenessResponse struct {
	Returnval bool `xml:"returnval"`
}

type VsanHostDpdLiveness VsanHostDpdLivenessRequestType

type VsanHostDpdLivenessRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
}

type VsanHostQueryCheckLimitsResponse struct {
	Returnval VsanLimitHealthResult `xml:"returnval"`
}

type VsanHostQueryCheckLimits VsanHostQueryCheckLimitsRequestType

type VsanHostQueryCheckLimitsRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
}

type VsanCheckDatastoreUsageResponse struct {
	Returnval VsanDatastoreUsageResult `xml:"returnval"`
}

type VsanCheckDatastoreUsage VsanCheckDatastoreUsageRequestType

type VsanCheckDatastoreUsageRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
}

type HostVsanHealthSystem struct {
	types.ManagedObjectReference
}

type VsanPurgeHclFilesResponse struct {
}

type VsanPurgeHclFiles VsanPurgeHclFilesRequestType

type VsanPurgeHclFilesRequestType struct {
	This     types.ManagedObjectReference `xml:"_this"`
	Sha1sums []string                     `xml:"sha1sums"`
}

type VsanQueryVcClusterCreateVmHealthHistoryTestResponse struct {
	Returnval []VsanClusterCreateVmHealthTestResult `xml:"returnval"`
}

type VsanQueryVcClusterCreateVmHealthHistoryTest VsanQueryVcClusterCreateVmHealthHistoryTestRequestType

type VsanQueryVcClusterCreateVmHealthHistoryTestRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Cluster types.ManagedObjectReference `xml:"cluster"`
	Count   int32                        `xml:"count,omitempty"`
}

type VsanQueryVcClusterAdvCfgSyncResponse struct {
	Returnval []VsanClusterAdvCfgSyncResult `xml:"returnval"`
}

type VsanQueryVcClusterAdvCfgSync VsanQueryVcClusterAdvCfgSyncRequestType

type VsanQueryVcClusterAdvCfgSyncRequestType struct {
	This                 types.ManagedObjectReference `xml:"_this"`
	Cluster              types.ManagedObjectReference `xml:"cluster"`
	Options              []string                     `xml:"options,omitempty"`
	IncludeAllAdvOptions *bool                        `xml:"includeAllAdvOptions"`
	NonDefaultOnly       *bool                        `xml:"nonDefaultOnly"`
}

type VsanQueryVcClusterObjExtAttrsResponse struct {
	Returnval []VsanClusterObjectExtAttrs `xml:"returnval"`
}

type VsanQueryVcClusterObjExtAttrs VsanQueryVcClusterObjExtAttrsRequestType

type VsanQueryVcClusterObjExtAttrsRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Cluster types.ManagedObjectReference `xml:"cluster"`
	Uuids   []string                     `xml:"uuids"`
}

type VsanVcClusterQueryObjectHealthSummaryResponse struct {
	Returnval VsanObjectOverallHealth `xml:"returnval"`
}

type VsanVcClusterQueryObjectHealthSummary VsanVcClusterQueryObjectHealthSummaryRequestType

type VsanVcClusterQueryObjectHealthSummaryRequestType struct {
	This                        types.ManagedObjectReference `xml:"_this"`
	Cluster                     types.ManagedObjectReference `xml:"cluster"`
	ObjUuids                    []string                     `xml:"objUuids,omitempty"`
	IncludeObjUuids             *bool                        `xml:"includeObjUuids"`
	IncludeDataProtectionHealth *bool                        `xml:"includeDataProtectionHealth"`
}

type VsanHealthSetLogLevelResponse struct {
}

type VsanHealthSetLogLevel VsanHealthSetLogLevelRequestType

type VsanHealthSetLogLevelRequestType struct {
	This  types.ManagedObjectReference `xml:"_this"`
	Level string                       `xml:"level,omitempty"`
}

type VsanHealthTestVsanClusterTelemetryProxyResponse struct {
	Returnval bool `xml:"returnval"`
}

type VsanHealthTestVsanClusterTelemetryProxy VsanHealthTestVsanClusterTelemetryProxyRequestType

type VsanHealthTestVsanClusterTelemetryProxyRequestType struct {
	This        types.ManagedObjectReference    `xml:"_this"`
	ProxyConfig VsanClusterTelemetryProxyConfig `xml:"proxyConfig"`
}

type VsanQueryVerifyVcClusterNetworkSettingsResponse struct {
	Returnval VsanClusterNetworkHealthResult `xml:"returnval"`
}

type VsanQueryVerifyVcClusterNetworkSettings VsanQueryVerifyVcClusterNetworkSettingsRequestType

type VsanQueryVerifyVcClusterNetworkSettingsRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Cluster types.ManagedObjectReference `xml:"cluster"`
}

type VsanVcUploadHclDbResponse struct {
	Returnval bool `xml:"returnval"`
}

type VsanVcUploadHclDb VsanVcUploadHclDbRequestType

type VsanVcUploadHclDbRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
	Db   string                       `xml:"db"`
}

type VsanQueryVcClusterSmartStatsSummaryResponse struct {
	Returnval []VsanSmartStatsHostSummary `xml:"returnval"`
}

type VsanQueryVcClusterSmartStatsSummary VsanQueryVcClusterSmartStatsSummaryRequestType

type VsanQueryVcClusterSmartStatsSummaryRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Cluster types.ManagedObjectReference `xml:"cluster"`
}

type VsanVcUpdateHclDbFromWebResponse struct {
	Returnval bool `xml:"returnval"`
}

type VsanVcUpdateHclDbFromWeb VsanVcUpdateHclDbFromWebRequestType

type VsanVcUpdateHclDbFromWebRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
	Url  string                       `xml:"url,omitempty"`
}

type VsanHealthGetVsanClusterSilentChecksResponse struct {
	Returnval []string `xml:"returnval"`
}

type VsanHealthGetVsanClusterSilentChecks VsanHealthGetVsanClusterSilentChecksRequestType

type VsanHealthGetVsanClusterSilentChecksRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Cluster types.ManagedObjectReference `xml:"cluster"`
}

type VsanClusterQueryFileServiceHealthSummaryResponse struct {
	Returnval VsanClusterFileServiceHealthSummary `xml:"returnval"`
}

type VsanClusterQueryFileServiceHealthSummary VsanClusterQueryFileServiceHealthSummaryRequestType

type VsanClusterQueryFileServiceHealthSummaryRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Cluster types.ManagedObjectReference `xml:"cluster"`
}

type PrepareClusterRebootWithNAMMResponse struct {
	Returnval bool `xml:"returnval"`
}

type PrepareClusterRebootWithNAMM PrepareClusterRebootWithNAMMRequestType

type PrepareClusterRebootWithNAMMRequestType struct {
	This         types.ManagedObjectReference `xml:"_this"`
	Cluster      types.ManagedObjectReference `xml:"cluster"`
	ScheduleTime int32                        `xml:"scheduleTime,omitempty"`
}

type VsanVcClusterClomdLivenessResponse struct {
	Returnval VsanClusterClomdLivenessResult `xml:"returnval"`
}

type VsanVcClusterClomdLiveness VsanVcClusterClomdLivenessRequestType

type VsanVcClusterClomdLivenessRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Cluster types.ManagedObjectReference `xml:"cluster"`
}

type VsanClusterHealthExtensionManagmentPreCheckResponse struct {
	Returnval VsanHealthExtMgmtPreCheckResult `xml:"returnval"`
}

type VsanClusterHealthExtensionManagmentPreCheck VsanClusterHealthExtensionManagmentPreCheckRequestType

type VsanClusterHealthExtensionManagmentPreCheckRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Cluster types.ManagedObjectReference `xml:"cluster"`
	Install bool                         `xml:"install"`
	Feature string                       `xml:"feature,omitempty"`
}

type VsanHealthRepairClusterObjectsImmediateResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type VsanHealthRepairClusterObjectsImmediate VsanHealthRepairClusterObjectsImmediateRequestType

type VsanHealthRepairClusterObjectsImmediateRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Cluster types.ManagedObjectReference `xml:"cluster"`
	Uuids   []string                     `xml:"uuids,omitempty"`
}

type VsanQueryVcClusterNetworkPerfTestResponse struct {
	Returnval VsanClusterNetworkLoadTestResult `xml:"returnval"`
}

type VsanQueryVcClusterNetworkPerfTest VsanQueryVcClusterNetworkPerfTestRequestType

type VsanQueryVcClusterNetworkPerfTestRequestType struct {
	This        types.ManagedObjectReference `xml:"_this"`
	Cluster     types.ManagedObjectReference `xml:"cluster"`
	Multicast   bool                         `xml:"multicast"`
	DurationSec int32                        `xml:"durationSec,omitempty"`
}

type VsanQueryVcClusterVmdkLoadHistoryTestResponse struct {
	Returnval []VsanClusterVmdkLoadTestResult `xml:"returnval"`
}

type VsanQueryVcClusterVmdkLoadHistoryTest VsanQueryVcClusterVmdkLoadHistoryTestRequestType

type VsanQueryVcClusterVmdkLoadHistoryTestRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Cluster types.ManagedObjectReference `xml:"cluster"`
	Count   int32                        `xml:"count,omitempty"`
	TaskId  string                       `xml:"taskId,omitempty"`
}

type VsanHealthIsRebalanceRunningResponse struct {
	Returnval bool `xml:"returnval"`
}

type VsanHealthIsRebalanceRunning VsanHealthIsRebalanceRunningRequestType

type VsanHealthIsRebalanceRunningRequestType struct {
	This        types.ManagedObjectReference   `xml:"_this"`
	Cluster     types.ManagedObjectReference   `xml:"cluster"`
	TargetHosts []types.ManagedObjectReference `xml:"targetHosts,omitempty"`
}

type VsanQueryVcClusterCreateVmHealthTestResponse struct {
	Returnval VsanClusterCreateVmHealthTestResult `xml:"returnval"`
}

type VsanQueryVcClusterCreateVmHealthTest VsanQueryVcClusterCreateVmHealthTestRequestType

type VsanQueryVcClusterCreateVmHealthTestRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Cluster types.ManagedObjectReference `xml:"cluster"`
	Timeout int32                        `xml:"timeout"`
}

type VsanHealthQueryVsanProxyConfigResponse struct {
	Returnval VsanClusterTelemetryProxyConfig `xml:"returnval"`
}

type VsanHealthQueryVsanProxyConfig VsanHealthQueryVsanProxyConfigRequestType

type VsanHealthQueryVsanProxyConfigRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
}

type VsanHealthQueryVsanClusterHealthCheckIntervalResponse struct {
	Returnval int32 `xml:"returnval"`
}

type VsanHealthQueryVsanClusterHealthCheckInterval VsanHealthQueryVsanClusterHealthCheckIntervalRequestType

type VsanHealthQueryVsanClusterHealthCheckIntervalRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Cluster types.ManagedObjectReference `xml:"cluster"`
}

type VsanQueryVcClusterEncryptionHealthSummaryResponse struct {
	Returnval VsanClusterEncryptionHealthSummary `xml:"returnval"`
}

type VsanQueryVcClusterEncryptionHealthSummary VsanQueryVcClusterEncryptionHealthSummaryRequestType

type VsanQueryVcClusterEncryptionHealthSummaryRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Cluster types.ManagedObjectReference `xml:"cluster"`
}

type VsanQueryAllSupportedHealthChecksResponse struct {
	Returnval []VsanClusterHealthCheckInfo `xml:"returnval"`
}

type VsanQueryAllSupportedHealthChecks VsanQueryAllSupportedHealthChecksRequestType

type VsanQueryAllSupportedHealthChecksRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
}

type VsanVcClusterGetHclInfoResponse struct {
	Returnval VsanClusterHclInfo `xml:"returnval"`
}

type VsanVcClusterGetHclInfo VsanVcClusterGetHclInfoRequestType

type VsanVcClusterGetHclInfoRequestType struct {
	This               types.ManagedObjectReference  `xml:"_this"`
	Cluster            *types.ManagedObjectReference `xml:"cluster,omitempty"`
	IncludeHostsResult *bool                         `xml:"includeHostsResult"`
	IncludeVendorInfo  *bool                         `xml:"includeVendorInfo"`
	EsxRelease         string                        `xml:"esxRelease,omitempty"`
}

type VsanHealthPrepareClusterResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type VsanHealthPrepareCluster VsanHealthPrepareClusterRequestType

type VsanHealthPrepareClusterRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Cluster types.ManagedObjectReference `xml:"cluster"`
	Feature string                       `xml:"feature,omitempty"`
}

type VsanQueryAttachToSrHistoryResponse struct {
	Returnval []VsanAttachToSrOperation `xml:"returnval"`
}

type VsanQueryAttachToSrHistory VsanQueryAttachToSrHistoryRequestType

type VsanQueryAttachToSrHistoryRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Cluster types.ManagedObjectReference `xml:"cluster"`
	Count   int32                        `xml:"count,omitempty"`
	TaskId  string                       `xml:"taskId,omitempty"`
}

type VsanVcClusterArchivalAccessibilityResponse struct {
	Returnval []VsanArchivalAccessibilityResult `xml:"returnval"`
}

type VsanVcClusterArchivalAccessibility VsanVcClusterArchivalAccessibilityRequestType

type VsanVcClusterArchivalAccessibilityRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Cluster types.ManagedObjectReference `xml:"cluster"`
}

type VsanGetHclConstraintsResponse struct {
	Returnval VsanHclReleaseConstraint `xml:"returnval"`
}

type VsanGetHclConstraints VsanGetHclConstraintsRequestType

type VsanGetHclConstraintsRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Cluster types.ManagedObjectReference `xml:"cluster"`
	Release string                       `xml:"release"`
}

type VsanRebalanceClusterResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type VsanRebalanceCluster VsanRebalanceClusterRequestType

type VsanRebalanceClusterRequestType struct {
	This        types.ManagedObjectReference   `xml:"_this"`
	Cluster     types.ManagedObjectReference   `xml:"cluster"`
	TargetHosts []types.ManagedObjectReference `xml:"targetHosts,omitempty"`
}

type VsanVcClusterRunVmdkLoadTestResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type VsanVcClusterRunVmdkLoadTest VsanVcClusterRunVmdkLoadTestRequestType

type VsanVcClusterRunVmdkLoadTestRequestType struct {
	This        types.ManagedObjectReference `xml:"_this"`
	Cluster     types.ManagedObjectReference `xml:"cluster"`
	Runname     string                       `xml:"runname"`
	DurationSec int32                        `xml:"durationSec,omitempty"`
	Specs       []VsanVmdkLoadTestSpec       `xml:"specs,omitempty"`
	Action      string                       `xml:"action,omitempty"`
}

type VsanHealthSendVsanTelemetryResponse struct {
}

type VsanHealthSendVsanTelemetry VsanHealthSendVsanTelemetryRequestType

type VsanHealthSendVsanTelemetryRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Cluster types.ManagedObjectReference `xml:"cluster"`
}

type VsanRemediateDataProtectionConfigInClusterResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type VsanRemediateDataProtectionConfigInCluster VsanRemediateDataProtectionConfigInClusterRequestType

type VsanRemediateDataProtectionConfigInClusterRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Cluster types.ManagedObjectReference `xml:"cluster"`
}

type VsanQueryVcClusterNetworkPerfHistoryTestResponse struct {
	Returnval []VsanClusterNetworkLoadTestResult `xml:"returnval"`
}

type VsanQueryVcClusterNetworkPerfHistoryTest VsanQueryVcClusterNetworkPerfHistoryTestRequestType

type VsanQueryVcClusterNetworkPerfHistoryTestRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Cluster types.ManagedObjectReference `xml:"cluster"`
	Count   int32                        `xml:"count,omitempty"`
}

type VsanQueryVcClusterHistoryHealthSummaryResponse struct {
	Returnval []VsanClusterHealthSummary `xml:"returnval"`
}

type VsanQueryVcClusterHistoryHealthSummary VsanQueryVcClusterHistoryHealthSummaryRequestType

type VsanQueryVcClusterHistoryHealthSummaryRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
	Spec VsanHistoryItemQuerySpec     `xml:"spec"`
}

type VsanQueryVcClusterHealthSummaryResponse struct {
	Returnval VsanClusterHealthSummary `xml:"returnval"`
}

type VsanQueryVcClusterHealthSummary VsanQueryVcClusterHealthSummaryRequestType

type VsanQueryVcClusterHealthSummaryRequestType struct {
	This                        types.ManagedObjectReference   `xml:"_this"`
	Cluster                     *types.ManagedObjectReference  `xml:"cluster,omitempty"`
	VmCreateTimeout             int32                          `xml:"vmCreateTimeout,omitempty"`
	ObjUuids                    []string                       `xml:"objUuids,omitempty"`
	IncludeObjUuids             *bool                          `xml:"includeObjUuids"`
	Fields                      []string                       `xml:"fields,omitempty"`
	FetchFromCache              *bool                          `xml:"fetchFromCache"`
	Perspective                 string                         `xml:"perspective,omitempty"`
	Hosts                       []types.ManagedObjectReference `xml:"hosts,omitempty"`
	IncludeDataProtectionHealth *bool                          `xml:"includeDataProtectionHealth"`
}

type VsanQueryVcClusterHealthSummaryTaskResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type VsanQueryVcClusterHealthSummaryTask VsanQueryVcClusterHealthSummaryTaskRequestType

type VsanQueryVcClusterHealthSummaryTaskRequestType struct {
	This                        types.ManagedObjectReference   `xml:"_this"`
	Cluster                     types.ManagedObjectReference   `xml:"cluster"`
	Hosts                       []types.ManagedObjectReference `xml:"hosts,omitempty"`
	IncludeDataProtectionHealth *bool                          `xml:"includeDataProtectionHealth"`
}

type VsanHealthGetClusterStatusResponse struct {
	Returnval string `xml:"returnval"`
}

type VsanHealthGetClusterStatus VsanHealthGetClusterStatusRequestType

type VsanHealthGetClusterStatusRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Cluster types.ManagedObjectReference `xml:"cluster"`
	Feature string                       `xml:"feature,omitempty"`
}

type VsanStopRebalanceClusterResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type VsanStopRebalanceCluster VsanStopRebalanceClusterRequestType

type VsanStopRebalanceClusterRequestType struct {
	This        types.ManagedObjectReference   `xml:"_this"`
	Cluster     types.ManagedObjectReference   `xml:"cluster"`
	TargetHosts []types.ManagedObjectReference `xml:"targetHosts,omitempty"`
}

type VsanHealthQueryVsanClusterHealthConfigResponse struct {
	Returnval VsanClusterHealthConfigs `xml:"returnval"`
}

type VsanHealthQueryVsanClusterHealthConfig VsanHealthQueryVsanClusterHealthConfigRequestType

type VsanHealthQueryVsanClusterHealthConfigRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Cluster types.ManagedObjectReference `xml:"cluster"`
}

type VsanAttachVsanSupportBundleToSrResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type VsanAttachVsanSupportBundleToSr VsanAttachVsanSupportBundleToSrRequestType

type VsanAttachVsanSupportBundleToSrRequestType struct {
	This     types.ManagedObjectReference `xml:"_this"`
	Cluster  types.ManagedObjectReference `xml:"cluster"`
	SrNumber string                       `xml:"srNumber"`
}

type VsanDownloadHclFile_TaskResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type VsanDownloadHclFile_Task VsanDownloadHclFile_TaskRequestType

type VsanDownloadHclFile_TaskRequestType struct {
	This     types.ManagedObjectReference `xml:"_this"`
	Sha1sums []string                     `xml:"sha1sums"`
}

type VsanVcQueryClusterCaptureVsanPcapResponse struct {
	Returnval VsanVsanClusterPcapResult `xml:"returnval"`
}

type VsanVcQueryClusterCaptureVsanPcap VsanVcQueryClusterCaptureVsanPcapRequestType

type VsanVcQueryClusterCaptureVsanPcapRequestType struct {
	This               types.ManagedObjectReference   `xml:"_this"`
	Cluster            types.ManagedObjectReference   `xml:"cluster"`
	Duration           int32                          `xml:"duration"`
	Vmknic             []VsanClusterHostVmknicMapping `xml:"vmknic,omitempty"`
	IncludeRawPcap     *bool                          `xml:"includeRawPcap"`
	IncludeIgmp        *bool                          `xml:"includeIgmp"`
	CmmdsMsgTypeFilter []string                       `xml:"cmmdsMsgTypeFilter,omitempty"`
	CmmdsPorts         []int32                        `xml:"cmmdsPorts,omitempty"`
	ClusterUuid        string                         `xml:"clusterUuid,omitempty"`
}

type VsanQueryVcClusterVmdkWorkloadTypesResponse struct {
	Returnval []VsanStorageWorkloadType `xml:"returnval"`
}

type VsanQueryVcClusterVmdkWorkloadTypes VsanQueryVcClusterVmdkWorkloadTypesRequestType

type VsanQueryVcClusterVmdkWorkloadTypesRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
}

type VsanQueryVcClusterPhysicalDiskHealthSummaryResponse struct {
	Returnval []VsanPhysicalDiskHealthSummary `xml:"returnval"`
}

type VsanQueryVcClusterPhysicalDiskHealthSummary VsanQueryVcClusterPhysicalDiskHealthSummaryRequestType

type VsanQueryVcClusterPhysicalDiskHealthSummaryRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Cluster types.ManagedObjectReference `xml:"cluster"`
}

type VsanHealthSetVsanClusterSilentChecksResponse struct {
	Returnval bool `xml:"returnval"`
}

type VsanHealthSetVsanClusterSilentChecks VsanHealthSetVsanClusterSilentChecksRequestType

type VsanHealthSetVsanClusterSilentChecksRequestType struct {
	This               types.ManagedObjectReference `xml:"_this"`
	Cluster            types.ManagedObjectReference `xml:"cluster"`
	AddSilentChecks    []string                     `xml:"addSilentChecks,omitempty"`
	RemoveSilentChecks []string                     `xml:"removeSilentChecks,omitempty"`
}

type VsanQueryVcClusterCheckLimitsResponse struct {
	Returnval VsanClusterLimitHealthResult `xml:"returnval"`
}

type VsanQueryVcClusterCheckLimits VsanQueryVcClusterCheckLimitsRequestType

type VsanQueryVcClusterCheckLimitsRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Cluster types.ManagedObjectReference `xml:"cluster"`
}

type VsanVcClusterQueryVerifyHealthSystemVersionsResponse struct {
	Returnval VsanClusterHealthSystemVersionResult `xml:"returnval"`
}

type VsanVcClusterQueryVerifyHealthSystemVersions VsanVcClusterQueryVerifyHealthSystemVersionsRequestType

type VsanVcClusterQueryVerifyHealthSystemVersionsRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Cluster types.ManagedObjectReference `xml:"cluster"`
}

type VsanHealthSetVsanClusterTelemetryConfigResponse struct {
}

type VsanHealthSetVsanClusterTelemetryConfig VsanHealthSetVsanClusterTelemetryConfigRequestType

type VsanHealthSetVsanClusterTelemetryConfigRequestType struct {
	This                    types.ManagedObjectReference `xml:"_this"`
	Cluster                 types.ManagedObjectReference `xml:"cluster"`
	VsanClusterHealthConfig VsanClusterHealthConfigs     `xml:"vsanClusterHealthConfig"`
}

type RestoreClusterRebootWithNAMMResponse struct {
	Returnval bool `xml:"returnval"`
}

type RestoreClusterRebootWithNAMM RestoreClusterRebootWithNAMMRequestType

type RestoreClusterRebootWithNAMMRequestType struct {
	This         types.ManagedObjectReference `xml:"_this"`
	Cluster      types.ManagedObjectReference `xml:"cluster"`
	ScheduleTime int32                        `xml:"scheduleTime,omitempty"`
}

type VsanQueryVcClusterDataProtectionCfgSyncResponse struct {
	Returnval []VsanClusterDataProtectionCfgSyncResult `xml:"returnval"`
}

type VsanQueryVcClusterDataProtectionCfgSync VsanQueryVcClusterDataProtectionCfgSyncRequestType

type VsanQueryVcClusterDataProtectionCfgSyncRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Cluster types.ManagedObjectReference `xml:"cluster"`
}

type VsanDownloadAndInstallVendorTool_TaskResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type VsanDownloadAndInstallVendorTool_Task VsanDownloadAndInstallVendorTool_TaskRequestType

type VsanDownloadAndInstallVendorTool_TaskRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Cluster types.ManagedObjectReference `xml:"cluster"`
}

type VsanHealthSetVsanClusterHealthCheckIntervalResponse struct {
}

type VsanHealthSetVsanClusterHealthCheckInterval VsanHealthSetVsanClusterHealthCheckIntervalRequestType

type VsanHealthSetVsanClusterHealthCheckIntervalRequestType struct {
	This                           types.ManagedObjectReference `xml:"_this"`
	Cluster                        types.ManagedObjectReference `xml:"cluster"`
	VsanClusterHealthCheckInterval int32                        `xml:"vsanClusterHealthCheckInterval"`
}

type VsanVcClusterDpdLivenessResponse struct {
	Returnval VsanClusterDpdLivenessResult `xml:"returnval"`
}

type VsanVcClusterDpdLiveness VsanVcClusterDpdLivenessRequestType

type VsanVcClusterDpdLivenessRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Cluster types.ManagedObjectReference `xml:"cluster"`
}

type VsanHealthUninstallClusterResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type VsanHealthUninstallCluster VsanHealthUninstallClusterRequestType

type VsanHealthUninstallClusterRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Cluster types.ManagedObjectReference `xml:"cluster"`
	Feature string                       `xml:"feature,omitempty"`
}

type VsanVcClusterDatastoreUsageResponse struct {
	Returnval VsanClusterDatastoreUsageResult `xml:"returnval"`
}

type VsanVcClusterDatastoreUsage VsanVcClusterDatastoreUsageRequestType

type VsanVcClusterDatastoreUsageRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Cluster types.ManagedObjectReference `xml:"cluster"`
}

type VsanVcClusterHealthSystem struct {
	types.ManagedObjectReference
}

type CaptureInternalStatsResponse struct {
	Returnval string `xml:"returnval"`
}

type CaptureInternalStats CaptureInternalStatsRequestType

type CaptureInternalStatsRequestType struct {
	This         types.ManagedObjectReference `xml:"_this"`
	CallerNodeId string                       `xml:"callerNodeId,omitempty"`
	Interval     int32                        `xml:"interval,omitempty"`
	VerboseMode  *bool                        `xml:"verboseMode"`
}

type VsanInternalStatsProvider struct {
	types.ManagedObjectReference
}

type VsanConvertDeploymentSpecFromJsonResponse struct {
	Returnval VsanVcsaDeploymentSpecValidationResult `xml:"returnval"`
}

type VsanConvertDeploymentSpecFromJson VsanConvertDeploymentSpecFromJsonRequestType

type VsanConvertDeploymentSpecFromJsonRequestType struct {
	This         types.ManagedObjectReference `xml:"_this"`
	VcConfigJson string                       `xml:"vcConfigJson"`
}

type VsanVcsaGetBootstrapTasksResponse struct {
	Returnval []VsanVcsaBootstrapOntoVsanSpec `xml:"returnval"`
}

type VsanVcsaGetBootstrapTasks VsanVcsaGetBootstrapTasksRequestType

type VsanVcsaGetBootstrapTasksRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
}

type VsanValidateDeploymentConfigResponse struct {
	Returnval VsanVcsaDeploymentSpecValidationResult `xml:"returnval"`
}

type VsanValidateDeploymentConfig VsanValidateDeploymentConfigRequestType

type VsanValidateDeploymentConfigRequestType struct {
	This     types.ManagedObjectReference `xml:"_this"`
	VcConfig VsanVcsaDeploymentSpec       `xml:"vcConfig"`
}

type VsanPostConfigForVcsaResponse struct {
	Returnval string `xml:"returnval"`
}

type VsanPostConfigForVcsa VsanPostConfigForVcsaRequestType

type VsanPostConfigForVcsaRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
	Spec VsanVcPostDeployConfigSpec   `xml:"spec"`
}

type VsanVcsaGetBootstrapProgressResponse struct {
	Returnval []VsanVcsaDeploymentProgress `xml:"returnval"`
}

type VsanVcsaGetBootstrapProgress VsanVcsaGetBootstrapProgressRequestType

type VsanVcsaGetBootstrapProgressRequestType struct {
	This   types.ManagedObjectReference `xml:"_this"`
	TaskId []string                     `xml:"taskId"`
}

type VsanPrepareVsanForVcsaResponse struct {
	Returnval string `xml:"returnval"`
}

type VsanPrepareVsanForVcsa VsanPrepareVsanForVcsaRequestType

type VsanPrepareVsanForVcsaRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
	Spec VsanPrepareVsanForVcsaSpec   `xml:"spec"`
}

type VsanVcsaCancelBootstrapTaskResponse struct {
}

type VsanVcsaCancelBootstrapTask VsanVcsaCancelBootstrapTaskRequestType

type VsanVcsaCancelBootstrapTaskRequestType struct {
	This     types.ManagedObjectReference `xml:"_this"`
	TaskId   string                       `xml:"taskId"`
	DeleteVm bool                         `xml:"deleteVm"`
}

type VsanVcsaDeployerSystem struct {
	types.ManagedObjectReference
}

type ComplianceResourceCheck_TaskResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type ComplianceResourceCheck_Task ComplianceResourceCheck_TaskRequestType

type ComplianceResourceCheck_TaskRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
}

type WhatIfDecomCheck_TaskResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type WhatIfDecomCheck_Task WhatIfDecomCheck_TaskRequestType

type WhatIfDecomCheck_TaskRequestType struct {
	This            types.ManagedObjectReference `xml:"_this"`
	Entities        []string                     `xml:"entities"`
	MaintenanceSpec types.HostMaintenanceSpec    `xml:"maintenanceSpec"`
}

type VimHostVsanPrecheckerSystem struct {
	types.ManagedObjectReference
}

type CnsSearchLabelsResponse struct {
	Returnval CnsSearchLabelResult `xml:"returnval"`
}

type CnsSearchLabels CnsSearchLabelsRequestType

type CnsSearchLabelsRequestType struct {
	This            types.ManagedObjectReference `xml:"_this"`
	SearchLabelSpec CnsSearchLabelSpec           `xml:"searchLabelSpec"`
}

type CnsExtendVolumeResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type CnsExtendVolume CnsExtendVolumeRequestType

type CnsExtendVolumeRequestType struct {
	This        types.ManagedObjectReference `xml:"_this"`
	ExtendSpecs []CnsVolumeExtendSpec        `xml:"extendSpecs"`
}

type CnsCreateVolumeResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type CnsCreateVolume CnsCreateVolumeRequestType

type CnsCreateVolumeRequestType struct {
	This        types.ManagedObjectReference `xml:"_this"`
	CreateSpecs []CnsVolumeCreateSpec        `xml:"createSpecs"`
}

type CnsGetLastFullSyncTimeResponse struct {
	Returnval time.Time `xml:"returnval"`
}

type CnsGetLastFullSyncTime CnsGetLastFullSyncTimeRequestType

type CnsGetLastFullSyncTimeRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
}

type CnsAttachVolumeResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type CnsAttachVolume CnsAttachVolumeRequestType

type CnsAttachVolumeRequestType struct {
	This        types.ManagedObjectReference `xml:"_this"`
	AttachSpecs []CnsVolumeAttachDetachSpec  `xml:"attachSpecs"`
}

type CnsTriggerFullSyncResponse struct {
}

type CnsTriggerFullSync CnsTriggerFullSyncRequestType

type CnsTriggerFullSyncRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
}

type CnsUpdateVolumeMetadataResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type CnsUpdateVolumeMetadata CnsUpdateVolumeMetadataRequestType

type CnsUpdateVolumeMetadataRequestType struct {
	This        types.ManagedObjectReference  `xml:"_this"`
	UpdateSpecs []CnsVolumeMetadataUpdateSpec `xml:"updateSpecs"`
}

type CnsQueryVolumeResponse struct {
	Returnval CnsQueryResult `xml:"returnval"`
}

type CnsQueryVolume CnsQueryVolumeRequestType

type CnsQueryVolumeRequestType struct {
	This   types.ManagedObjectReference `xml:"_this"`
	Filter CnsQueryFilter               `xml:"filter"`
}

type CnsDetachVolumeResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type CnsDetachVolume CnsDetachVolumeRequestType

type CnsDetachVolumeRequestType struct {
	This        types.ManagedObjectReference `xml:"_this"`
	DetachSpecs []CnsVolumeAttachDetachSpec  `xml:"detachSpecs"`
}

type CnsQueryAllVolumeResponse struct {
	Returnval CnsQueryResult `xml:"returnval"`
}

type CnsQueryAllVolume CnsQueryAllVolumeRequestType

type CnsQueryAllVolumeRequestType struct {
	This      types.ManagedObjectReference `xml:"_this"`
	Filter    CnsQueryFilter               `xml:"filter"`
	Selection CnsQuerySelection            `xml:"selection"`
}

type CnsDeleteVolumeResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type CnsDeleteVolume CnsDeleteVolumeRequestType

type CnsDeleteVolumeRequestType struct {
	This       types.ManagedObjectReference `xml:"_this"`
	VolumeIds  []CnsVolumeId                `xml:"volumeIds"`
	DeleteDisk bool                         `xml:"deleteDisk"`
}

type CnsVolumeManager struct {
	types.ManagedObjectReference
}

type VsanQueryDecomResourceCheckAsync_TaskResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type VsanQueryDecomResourceCheckAsync_Task VsanQueryDecomResourceCheckAsync_TaskRequestType

type VsanQueryDecomResourceCheckAsync_TaskRequestType struct {
	This            types.ManagedObjectReference `xml:"_this"`
	Cluster         types.ManagedObjectReference `xml:"cluster"`
	Entities        []string                     `xml:"entities"`
	MaintenanceSpec types.HostMaintenanceSpec    `xml:"maintenanceSpec"`
}

type VsanQueryComplianceResourceCheckAsync_TaskResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type VsanQueryComplianceResourceCheckAsync_Task VsanQueryComplianceResourceCheckAsync_TaskRequestType

type VsanQueryComplianceResourceCheckAsync_TaskRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Cluster types.ManagedObjectReference `xml:"cluster"`
}

type VsanGetComplianceResourceCheckStatusResponse struct {
	Returnval VsanComplianceResourceCheckStatus `xml:"returnval"`
}

type VsanGetComplianceResourceCheckStatus VsanGetComplianceResourceCheckStatusRequestType

type VsanGetComplianceResourceCheckStatusRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Cluster types.ManagedObjectReference `xml:"cluster"`
}

type VsanGetResourceCheckStatus2Response struct {
	Returnval VsanComplianceResourceCheckStatus `xml:"returnval"`
}

type VsanGetResourceCheckStatus2 VsanGetResourceCheckStatus2RequestType

type VsanGetResourceCheckStatus2RequestType struct {
	This            types.ManagedObjectReference `xml:"_this"`
	Cluster         types.ManagedObjectReference `xml:"cluster"`
	QueryType       string                       `xml:"queryType"`
	Entities        []string                     `xml:"entities,omitempty"`
	MaintenanceSpec *types.HostMaintenanceSpec   `xml:"maintenanceSpec,omitempty"`
}

type VsanVcPrecheckerSystem struct {
	types.ManagedObjectReference
}

type VsanPrepareDefragResponse struct {
	Returnval bool `xml:"returnval"`
}

type VsanPrepareDefrag VsanPrepareDefragRequestType

type VsanPrepareDefragRequestType struct {
	This      types.ManagedObjectReference `xml:"_this"`
	Cluster   types.ManagedObjectReference `xml:"cluster"`
	VmxPath   string                       `xml:"vmxPath"`
	DiskPaths []string                     `xml:"diskPaths"`
}

type VsanCleanupDefragResponse struct {
	Returnval bool `xml:"returnval"`
}

type VsanCleanupDefrag VsanCleanupDefragRequestType

type VsanCleanupDefragRequestType struct {
	This      types.ManagedObjectReference `xml:"_this"`
	Cluster   types.ManagedObjectReference `xml:"cluster"`
	VmxPath   string                       `xml:"vmxPath"`
	DiskPaths []string                     `xml:"diskPaths"`
}

type VsanDataProtectionSpaceSystem struct {
	types.ManagedObjectReference
}

type VsanVdsMigrateVssResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type VsanVdsMigrateVss VsanVdsMigrateVssRequestType

type VsanVdsMigrateVssRequestType struct {
	This          types.ManagedObjectReference   `xml:"_this"`
	Cluster       types.ManagedObjectReference   `xml:"cluster"`
	MigrationPlan *VsanVdsMigrationPlan          `xml:"migrationPlan,omitempty"`
	VswitchName   string                         `xml:"vswitchName,omitempty"`
	VdsName       string                         `xml:"vdsName,omitempty"`
	VmnicDevices  []string                       `xml:"vmnicDevices,omitempty"`
	InfraVm       []types.ManagedObjectReference `xml:"infraVm,omitempty"`
	Vds           *types.ManagedObjectReference  `xml:"vds,omitempty"`
	Hosts         []types.ManagedObjectReference `xml:"hosts,omitempty"`
}

type VsanVdsGetMigrationPlanResponse struct {
	Returnval VsanVdsMigrationPlan `xml:"returnval"`
}

type VsanVdsGetMigrationPlan VsanVdsGetMigrationPlanRequestType

type VsanVdsGetMigrationPlanRequestType struct {
	This         types.ManagedObjectReference   `xml:"_this"`
	Cluster      types.ManagedObjectReference   `xml:"cluster"`
	VswitchName  string                         `xml:"vswitchName,omitempty"`
	VdsName      string                         `xml:"vdsName,omitempty"`
	VmnicDevices []string                       `xml:"vmnicDevices,omitempty"`
	InfraVm      []types.ManagedObjectReference `xml:"infraVm,omitempty"`
	Vds          *types.ManagedObjectReference  `xml:"vds,omitempty"`
	Hosts        []types.ManagedObjectReference `xml:"hosts,omitempty"`
}

type VsanVssMigrateVdsResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type VsanVssMigrateVds VsanVssMigrateVdsRequestType

type VsanVssMigrateVdsRequestType struct {
	This         types.ManagedObjectReference   `xml:"_this"`
	Cluster      *types.ManagedObjectReference  `xml:"cluster,omitempty"`
	Hosts        []types.ManagedObjectReference `xml:"hosts,omitempty"`
	Vds          types.ManagedObjectReference   `xml:"vds"`
	VswitchName  string                         `xml:"vswitchName,omitempty"`
	VmnicDevices []string                       `xml:"vmnicDevices,omitempty"`
	InfraVm      []types.ManagedObjectReference `xml:"infraVm,omitempty"`
}

type VsanRollbackVdsToVssResponse struct {
	Returnval bool `xml:"returnval"`
}

type VsanRollbackVdsToVss VsanRollbackVdsToVssRequestType

type VsanRollbackVdsToVssRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
	Task types.ManagedObjectReference `xml:"task"`
}

type VsanVdsSystem struct {
	types.ManagedObjectReference
}

type VosSetVsanObjectPolicyResponse struct {
	Returnval bool `xml:"returnval"`
}

type VosSetVsanObjectPolicy VosSetVsanObjectPolicyRequestType

type VosSetVsanObjectPolicyRequestType struct {
	This           types.ManagedObjectReference     `xml:"_this"`
	Cluster        *types.ManagedObjectReference    `xml:"cluster,omitempty"`
	VsanObjectUuid string                           `xml:"vsanObjectUuid"`
	Profile        *types.VirtualMachineProfileSpec `xml:"profile,omitempty"`
}

type VsanDeleteObjects_TaskResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type VsanDeleteObjects_Task VsanDeleteObjects_TaskRequestType

type VsanDeleteObjects_TaskRequestType struct {
	This     types.ManagedObjectReference  `xml:"_this"`
	Cluster  *types.ManagedObjectReference `xml:"cluster,omitempty"`
	ObjUuids []string                      `xml:"objUuids"`
	Force    *bool                         `xml:"force"`
}

type VosQueryVsanObjectInformationResponse struct {
	Returnval []VsanObjectInformation `xml:"returnval"`
}

type VosQueryVsanObjectInformation VosQueryVsanObjectInformationRequestType

type VosQueryVsanObjectInformationRequestType struct {
	This                 types.ManagedObjectReference  `xml:"_this"`
	Cluster              *types.ManagedObjectReference `xml:"cluster,omitempty"`
	VsanObjectQuerySpecs []VsanObjectQuerySpec         `xml:"vsanObjectQuerySpecs"`
}

type RelayoutObjectsResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type RelayoutObjects RelayoutObjectsRequestType

type RelayoutObjectsRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
}

type VsanQueryInaccessibleVmSwapObjectsResponse struct {
	Returnval []string `xml:"returnval"`
}

type VsanQueryInaccessibleVmSwapObjects VsanQueryInaccessibleVmSwapObjectsRequestType

type VsanQueryInaccessibleVmSwapObjectsRequestType struct {
	This    types.ManagedObjectReference  `xml:"_this"`
	Cluster *types.ManagedObjectReference `xml:"cluster,omitempty"`
}

type QuerySyncingVsanObjectsSummaryResponse struct {
	Returnval VsanHostVsanObjectSyncQueryResult `xml:"returnval"`
}

type QuerySyncingVsanObjectsSummary QuerySyncingVsanObjectsSummaryRequestType

type QuerySyncingVsanObjectsSummaryRequestType struct {
	This                types.ManagedObjectReference `xml:"_this"`
	Cluster             types.ManagedObjectReference `xml:"cluster"`
	SyncingObjectFilter *VsanSyncingObjectFilter     `xml:"syncingObjectFilter,omitempty"`
}

type VsanQueryInactiveDataProtectionSpaceUsageResponse struct {
	Returnval []VsanDataProtectionSpaceUsage `xml:"returnval"`
}

type VsanQueryInactiveDataProtectionSpaceUsage VsanQueryInactiveDataProtectionSpaceUsageRequestType

type VsanQueryInactiveDataProtectionSpaceUsageRequestType struct {
	This                      types.ManagedObjectReference             `xml:"_this"`
	Cluster                   types.ManagedObjectReference             `xml:"cluster"`
	DataProtectionTargetInfos []VsanDataProtectionRemoteTargetSiteInfo `xml:"dataProtectionTargetInfos"`
}

type VsanQueryObjectIdentitiesResponse struct {
	Returnval VsanObjectIdentityAndHealth `xml:"returnval"`
}

type VsanQueryObjectIdentities VsanQueryObjectIdentitiesRequestType

type VsanQueryObjectIdentitiesRequestType struct {
	This                        types.ManagedObjectReference  `xml:"_this"`
	Cluster                     *types.ManagedObjectReference `xml:"cluster,omitempty"`
	ObjUuids                    []string                      `xml:"objUuids,omitempty"`
	ObjTypes                    []string                      `xml:"objTypes,omitempty"`
	IncludeHealth               *bool                         `xml:"includeHealth"`
	IncludeObjIdentity          *bool                         `xml:"includeObjIdentity"`
	IncludeSpaceSummary         *bool                         `xml:"includeSpaceSummary"`
	IncludeDataProtectionHealth *bool                         `xml:"includeDataProtectionHealth"`
}

type VsanObjectSystem struct {
	types.ManagedObjectReference
}

type HostSpbmRetrieveAllPolicyInfoResponse struct {
	Returnval []HostSpbmPolicyInfo `xml:"returnval"`
}

type HostSpbmRetrieveAllPolicyInfo HostSpbmRetrieveAllPolicyInfoRequestType

type HostSpbmRetrieveAllPolicyInfoRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
}

type HostSpbmComputeHashInfoResponse struct {
	Returnval HostSpbmHashInfo `xml:"returnval"`
}

type HostSpbmComputeHashInfo HostSpbmComputeHashInfoRequestType

type HostSpbmComputeHashInfoRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
}

type HostSpbmRetrieveAllDatastoreInfoResponse struct {
	Returnval []HostSpbmDatastoreInfo `xml:"returnval"`
}

type HostSpbmRetrieveAllDatastoreInfo HostSpbmRetrieveAllDatastoreInfoRequestType

type HostSpbmRetrieveAllDatastoreInfoRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
}

type HostSpbmFetchDefaultPolicyBlobResponse struct {
	Returnval string `xml:"returnval"`
}

type HostSpbmFetchDefaultPolicyBlob HostSpbmFetchDefaultPolicyBlobRequestType

type HostSpbmFetchDefaultPolicyBlobRequestType struct {
	This         types.ManagedObjectReference `xml:"_this"`
	DatastoreUrl string                       `xml:"datastoreUrl"`
}

type HostSpbmPushPolicyInfoResponse struct {
}

type HostSpbmPushPolicyInfo HostSpbmPushPolicyInfoRequestType

type HostSpbmPushPolicyInfoRequestType struct {
	This           types.ManagedObjectReference `xml:"_this"`
	SpbmPolicyInfo []HostSpbmPolicyInfo         `xml:"spbmPolicyInfo"`
}

type HostSpbmPushDatastoreInfoResponse struct {
}

type HostSpbmPushDatastoreInfo HostSpbmPushDatastoreInfoRequestType

type HostSpbmPushDatastoreInfoRequestType struct {
	This              types.ManagedObjectReference `xml:"_this"`
	SpbmDatastoreInfo []HostSpbmDatastoreInfo      `xml:"spbmDatastoreInfo"`
}

type HostSpbmFetchApplicablePolicyBlobResponse struct {
	Returnval string `xml:"returnval"`
}

type HostSpbmFetchApplicablePolicyBlob HostSpbmFetchApplicablePolicyBlobRequestType

type HostSpbmFetchApplicablePolicyBlobRequestType struct {
	This         types.ManagedObjectReference `xml:"_this"`
	ProfileId    string                       `xml:"profileId"`
	DatastoreUrl string                       `xml:"datastoreUrl"`
}

type HostSpbm struct {
	types.ManagedObjectReference
}

type VsanRecoveryRecoverComponentResponse struct {
	Returnval string `xml:"returnval"`
}

type VsanRecoveryRecoverComponent VsanRecoveryRecoverComponentRequestType

type VsanRecoveryRecoverComponentRequestType struct {
	This     types.ManagedObjectReference     `xml:"_this"`
	CompInfo VimHostVsanRecoveryComponentInfo `xml:"compInfo"`
	DestUuid string                           `xml:"destUuid"`
	Parent   string                           `xml:"parent,omitempty"`
}

type VsanRecoveryServiceGetDummyCmmdsModeResponse struct {
	Returnval bool `xml:"returnval"`
}

type VsanRecoveryServiceGetDummyCmmdsMode VsanRecoveryServiceGetDummyCmmdsModeRequestType

type VsanRecoveryServiceGetDummyCmmdsModeRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
}

type VsanRecoveryGetRecoverableComponentsResponse struct {
	Returnval []VimHostVsanRecoveryComponentInfo `xml:"returnval"`
}

type VsanRecoveryGetRecoverableComponents VsanRecoveryGetRecoverableComponentsRequestType

type VsanRecoveryGetRecoverableComponentsRequestType struct {
	This       types.ManagedObjectReference `xml:"_this"`
	ObjectUuid string                       `xml:"objectUuid"`
	Flags      *VimHostVsanRecoveryFlags    `xml:"flags,omitempty"`
}

type VsanRecoveryGetAllObjectsStatusResponse struct {
	Returnval []VimHostVsanRecoveryObjectInfo `xml:"returnval"`
}

type VsanRecoveryGetAllObjectsStatus VsanRecoveryGetAllObjectsStatusRequestType

type VsanRecoveryGetAllObjectsStatusRequestType struct {
	This  types.ManagedObjectReference `xml:"_this"`
	Flags *VimHostVsanRecoveryFlags    `xml:"flags,omitempty"`
}

type VsanRecoveryGetObjectRecoveryInfoResponse struct {
	Returnval VimHostVsanRecoveryObjectInfo `xml:"returnval"`
}

type VsanRecoveryGetObjectRecoveryInfo VsanRecoveryGetObjectRecoveryInfoRequestType

type VsanRecoveryGetObjectRecoveryInfoRequestType struct {
	This       types.ManagedObjectReference `xml:"_this"`
	ObjectUuid string                       `xml:"objectUuid"`
	Flags      *VimHostVsanRecoveryFlags    `xml:"flags,omitempty"`
}

type VsanRecoveryDeleteLsomCompLocalResponse struct {
	Returnval bool `xml:"returnval"`
}

type VsanRecoveryDeleteLsomCompLocal VsanRecoveryDeleteLsomCompLocalRequestType

type VsanRecoveryDeleteLsomCompLocalRequestType struct {
	This     types.ManagedObjectReference `xml:"_this"`
	CompUuid string                       `xml:"compUuid"`
	DiskUuid string                       `xml:"diskUuid"`
	Force    *bool                        `xml:"force"`
}

type VsanRecoveryGetDiskUuidForLocalCompResponse struct {
	Returnval string `xml:"returnval"`
}

type VsanRecoveryGetDiskUuidForLocalComp VsanRecoveryGetDiskUuidForLocalCompRequestType

type VsanRecoveryGetDiskUuidForLocalCompRequestType struct {
	This     types.ManagedObjectReference `xml:"_this"`
	CompUuid string                       `xml:"compUuid"`
}

type VsanRecoveryCreateObjectResponse struct {
	Returnval string `xml:"returnval"`
}

type VsanRecoveryCreateObject VsanRecoveryCreateObjectRequestType

type VsanRecoveryCreateObjectRequestType struct {
	This        types.ManagedObjectReference   `xml:"_this"`
	ObjectProps VimHostVsanRecoveryObjectProps `xml:"objectProps"`
	Policy      string                         `xml:"policy,omitempty"`
}

type VsanRecoveryGetDiskAndHostForCompResponse struct {
	Returnval VimHostVsanRecoveryHostDiskInfo `xml:"returnval"`
}

type VsanRecoveryGetDiskAndHostForComp VsanRecoveryGetDiskAndHostForCompRequestType

type VsanRecoveryGetDiskAndHostForCompRequestType struct {
	This     types.ManagedObjectReference `xml:"_this"`
	CompUuid string                       `xml:"compUuid"`
}

type VsanRecoveryServiceVersionResponse struct {
	Returnval string `xml:"returnval"`
}

type VsanRecoveryServiceVersion VsanRecoveryServiceVersionRequestType

type VsanRecoveryServiceVersionRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
}

type VsanRecoveryGetTaskStatusResponse struct {
	Returnval VimHostVsanRecoveryStatus `xml:"returnval"`
}

type VsanRecoveryGetTaskStatus VsanRecoveryGetTaskStatusRequestType

type VsanRecoveryGetTaskStatusRequestType struct {
	This   types.ManagedObjectReference `xml:"_this"`
	TaskId string                       `xml:"taskId"`
}

type VsanRecoveryServiceSetDummyCmmdsModeResponse struct {
}

type VsanRecoveryServiceSetDummyCmmdsMode VsanRecoveryServiceSetDummyCmmdsModeRequestType

type VsanRecoveryServiceSetDummyCmmdsModeRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
	Db   string                       `xml:"db,omitempty"`
}

type VsanRecoveryGetObjectPropsResponse struct {
	Returnval VimHostVsanRecoveryObjectProps `xml:"returnval"`
}

type VsanRecoveryGetObjectProps VsanRecoveryGetObjectPropsRequestType

type VsanRecoveryGetObjectPropsRequestType struct {
	This       types.ManagedObjectReference `xml:"_this"`
	ObjectUuid string                       `xml:"objectUuid"`
}

type VsanRecoveryRecoverObjectResponse struct {
	Returnval string `xml:"returnval"`
}

type VsanRecoveryRecoverObject VsanRecoveryRecoverObjectRequestType

type VsanRecoveryRecoverObjectRequestType struct {
	This        types.ManagedObjectReference  `xml:"_this"`
	ObjectInfo  VimHostVsanRecoveryObjectInfo `xml:"objectInfo"`
	DestObjUuid string                        `xml:"destObjUuid"`
}

type VimHostVsanRecoveryService struct {
	types.ManagedObjectReference
}

type VSANVcIsWitnessHostResponse struct {
	Returnval bool `xml:"returnval"`
}

type VSANVcIsWitnessHost VSANVcIsWitnessHostRequestType

type VSANVcIsWitnessHostRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
	Host types.ManagedObjectReference `xml:"host"`
}

type VSANVcSetPreferredFaultDomainResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type VSANVcSetPreferredFaultDomain VSANVcSetPreferredFaultDomainRequestType

type VSANVcSetPreferredFaultDomainRequestType struct {
	This        types.ManagedObjectReference  `xml:"_this"`
	Cluster     types.ManagedObjectReference  `xml:"cluster"`
	PreferredFd string                        `xml:"preferredFd"`
	WitnessHost *types.ManagedObjectReference `xml:"witnessHost,omitempty"`
}

type VSANVcReconcileUnicastAgentsResponse struct {
	Returnval bool `xml:"returnval"`
}

type VSANVcReconcileUnicastAgents VSANVcReconcileUnicastAgentsRequestType

type VSANVcReconcileUnicastAgentsRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Cluster types.ManagedObjectReference `xml:"cluster"`
}

type VSANVcGetPreferredFaultDomainResponse struct {
	Returnval VimClusterVSANPreferredFaultDomainInfo `xml:"returnval"`
}

type VSANVcGetPreferredFaultDomain VSANVcGetPreferredFaultDomainRequestType

type VSANVcGetPreferredFaultDomainRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Cluster types.ManagedObjectReference `xml:"cluster"`
}

type VSANIsWitnessVirtualApplianceResponse struct {
	Returnval []VsanHostVirtualApplianceInfo `xml:"returnval"`
}

type VSANIsWitnessVirtualAppliance VSANIsWitnessVirtualApplianceRequestType

type VSANIsWitnessVirtualApplianceRequestType struct {
	This  types.ManagedObjectReference   `xml:"_this"`
	Hosts []types.ManagedObjectReference `xml:"hosts"`
}

type VSANVcAddWitnessHostResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type VSANVcAddWitnessHost VSANVcAddWitnessHostRequestType

type VSANVcAddWitnessHostRequestType struct {
	This         types.ManagedObjectReference `xml:"_this"`
	Cluster      types.ManagedObjectReference `xml:"cluster"`
	WitnessHost  types.ManagedObjectReference `xml:"witnessHost"`
	PreferredFd  string                       `xml:"preferredFd"`
	DiskMapping  *types.VsanHostDiskMapping   `xml:"diskMapping,omitempty"`
	MetadataMode *bool                        `xml:"metadataMode"`
}

type VSANVcRetrieveStretchedClusterHostCapabilityResponse struct {
	Returnval VimClusterVSANStretchedClusterCapability `xml:"returnval"`
}

type VSANVcRetrieveStretchedClusterHostCapability VSANVcRetrieveStretchedClusterHostCapabilityRequestType

type VSANVcRetrieveStretchedClusterHostCapabilityRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
	Host types.ManagedObjectReference `xml:"host"`
}

type VSANVcGetNumOfWitnessHostsResponse struct {
	Returnval int32 `xml:"returnval"`
}

type VSANVcGetNumOfWitnessHosts VSANVcGetNumOfWitnessHostsRequestType

type VSANVcGetNumOfWitnessHostsRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Cluster types.ManagedObjectReference `xml:"cluster"`
}

type VSANVcGetStretchedClusterConfigIssuesResponse struct {
	Returnval []VimClusterVSANStretchedClusterConfigIssueInfo `xml:"returnval"`
}

type VSANVcGetStretchedClusterConfigIssues VSANVcGetStretchedClusterConfigIssuesRequestType

type VSANVcGetStretchedClusterConfigIssuesRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Cluster types.ManagedObjectReference `xml:"cluster"`
}

type VSANVcGetWitnessHostsResponse struct {
	Returnval []VimClusterVSANWitnessHostInfo `xml:"returnval"`
}

type VSANVcGetWitnessHosts VSANVcGetWitnessHostsRequestType

type VSANVcGetWitnessHostsRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Cluster types.ManagedObjectReference `xml:"cluster"`
}

type VSANVcRetrieveStretchedClusterVcCapabilityResponse struct {
	Returnval []VimClusterVSANStretchedClusterCapability `xml:"returnval"`
}

type VSANVcRetrieveStretchedClusterVcCapability VSANVcRetrieveStretchedClusterVcCapabilityRequestType

type VSANVcRetrieveStretchedClusterVcCapabilityRequestType struct {
	This               types.ManagedObjectReference `xml:"_this"`
	Cluster            types.ManagedObjectReference `xml:"cluster"`
	VerifyAllConnected *bool                        `xml:"verifyAllConnected"`
}

type VSANVcConvertToStretchedClusterResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type VSANVcConvertToStretchedCluster VSANVcConvertToStretchedClusterRequestType

type VSANVcConvertToStretchedClusterRequestType struct {
	This              types.ManagedObjectReference                    `xml:"_this"`
	Cluster           types.ManagedObjectReference                    `xml:"cluster"`
	FaultDomainConfig VimClusterVSANStretchedClusterFaultDomainConfig `xml:"faultDomainConfig"`
	WitnessHost       types.ManagedObjectReference                    `xml:"witnessHost"`
	PreferredFd       string                                          `xml:"preferredFd"`
	DiskMapping       *types.VsanHostDiskMapping                      `xml:"diskMapping,omitempty"`
}

type VSANVcRemoveWitnessHostResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type VSANVcRemoveWitnessHost VSANVcRemoveWitnessHostRequestType

type VSANVcRemoveWitnessHostRequestType struct {
	This           types.ManagedObjectReference  `xml:"_this"`
	Cluster        types.ManagedObjectReference  `xml:"cluster"`
	WitnessHost    *types.ManagedObjectReference `xml:"witnessHost,omitempty"`
	WitnessAddress string                        `xml:"witnessAddress,omitempty"`
}

type VimClusterVsanVcStretchedClusterSystem struct {
	types.ManagedObjectReference
}

type VsanQueryClusterPhysicalDiskHealthSummaryResponse struct {
	Returnval []VsanPhysicalDiskHealthSummary `xml:"returnval"`
}

type VsanQueryClusterPhysicalDiskHealthSummary VsanQueryClusterPhysicalDiskHealthSummaryRequestType

type VsanQueryClusterPhysicalDiskHealthSummaryRequestType struct {
	This            types.ManagedObjectReference `xml:"_this"`
	Hosts           []string                     `xml:"hosts"`
	EsxRootPassword string                       `xml:"esxRootPassword"`
}

type VsanQueryClusterNetworkPerfTestResponse struct {
	Returnval VsanClusterNetworkLoadTestResult `xml:"returnval"`
}

type VsanQueryClusterNetworkPerfTest VsanQueryClusterNetworkPerfTestRequestType

type VsanQueryClusterNetworkPerfTestRequestType struct {
	This            types.ManagedObjectReference `xml:"_this"`
	Hosts           []string                     `xml:"hosts"`
	EsxRootPassword string                       `xml:"esxRootPassword"`
	Multicast       bool                         `xml:"multicast"`
	DurationSec     int32                        `xml:"durationSec,omitempty"`
}

type VsanQueryClusterAdvCfgSyncResponse struct {
	Returnval []VsanClusterAdvCfgSyncResult `xml:"returnval"`
}

type VsanQueryClusterAdvCfgSync VsanQueryClusterAdvCfgSyncRequestType

type VsanQueryClusterAdvCfgSyncRequestType struct {
	This            types.ManagedObjectReference `xml:"_this"`
	Hosts           []string                     `xml:"hosts"`
	EsxRootPassword string                       `xml:"esxRootPassword"`
	Options         []string                     `xml:"options,omitempty"`
}

type VsanCheckClusterArchivalAccessibilityResponse struct {
	Returnval []VsanArchivalAccessibilityResult `xml:"returnval"`
}

type VsanCheckClusterArchivalAccessibility VsanCheckClusterArchivalAccessibilityRequestType

type VsanCheckClusterArchivalAccessibilityRequestType struct {
	This            types.ManagedObjectReference `xml:"_this"`
	Hosts           []string                     `xml:"hosts"`
	EsxRootPassword string                       `xml:"esxRootPassword"`
}

type VsanCheckClusterDpdLivenessResponse struct {
	Returnval VsanClusterDpdLivenessResult `xml:"returnval"`
}

type VsanCheckClusterDpdLiveness VsanCheckClusterDpdLivenessRequestType

type VsanCheckClusterDpdLivenessRequestType struct {
	This            types.ManagedObjectReference `xml:"_this"`
	Hosts           []string                     `xml:"hosts"`
	EsxRootPassword string                       `xml:"esxRootPassword"`
}

type VsanRepairClusterImmediateObjectsResponse struct {
	Returnval VsanClusterHealthSystemObjectsRepairResult `xml:"returnval"`
}

type VsanRepairClusterImmediateObjects VsanRepairClusterImmediateObjectsRequestType

type VsanRepairClusterImmediateObjectsRequestType struct {
	This            types.ManagedObjectReference `xml:"_this"`
	Hosts           []string                     `xml:"hosts"`
	EsxRootPassword string                       `xml:"esxRootPassword"`
	Uuids           []string                     `xml:"uuids,omitempty"`
}

type VsanQueryVerifyClusterNetworkSettingsResponse struct {
	Returnval VsanClusterNetworkHealthResult `xml:"returnval"`
}

type VsanQueryVerifyClusterNetworkSettings VsanQueryVerifyClusterNetworkSettingsRequestType

type VsanQueryVerifyClusterNetworkSettingsRequestType struct {
	This            types.ManagedObjectReference `xml:"_this"`
	Hosts           []string                     `xml:"hosts"`
	EsxRootPassword string                       `xml:"esxRootPassword"`
}

type VsanQueryClusterCreateVmHealthTestResponse struct {
	Returnval VsanClusterCreateVmHealthTestResult `xml:"returnval"`
}

type VsanQueryClusterCreateVmHealthTest VsanQueryClusterCreateVmHealthTestRequestType

type VsanQueryClusterCreateVmHealthTestRequestType struct {
	This            types.ManagedObjectReference `xml:"_this"`
	Hosts           []string                     `xml:"hosts"`
	EsxRootPassword string                       `xml:"esxRootPassword"`
	Timeout         int32                        `xml:"timeout"`
}

type VsanQueryClusterHealthSystemVersionsResponse struct {
	Returnval VsanClusterHealthSystemVersionResult `xml:"returnval"`
}

type VsanQueryClusterHealthSystemVersions VsanQueryClusterHealthSystemVersionsRequestType

type VsanQueryClusterHealthSystemVersionsRequestType struct {
	This            types.ManagedObjectReference `xml:"_this"`
	Hosts           []string                     `xml:"hosts"`
	EsxRootPassword string                       `xml:"esxRootPassword"`
}

type VsanClusterGetHclInfoResponse struct {
	Returnval VsanClusterHclInfo `xml:"returnval"`
}

type VsanClusterGetHclInfo VsanClusterGetHclInfoRequestType

type VsanClusterGetHclInfoRequestType struct {
	This            types.ManagedObjectReference `xml:"_this"`
	Hosts           []string                     `xml:"hosts"`
	EsxRootPassword string                       `xml:"esxRootPassword"`
}

type VsanQueryClusterCheckLimitsResponse struct {
	Returnval VsanClusterLimitHealthResult `xml:"returnval"`
}

type VsanQueryClusterCheckLimits VsanQueryClusterCheckLimitsRequestType

type VsanQueryClusterCheckLimitsRequestType struct {
	This            types.ManagedObjectReference `xml:"_this"`
	Hosts           []string                     `xml:"hosts"`
	EsxRootPassword string                       `xml:"esxRootPassword"`
}

type VsanQueryClusterCaptureVsanPcapResponse struct {
	Returnval VsanVsanClusterPcapResult `xml:"returnval"`
}

type VsanQueryClusterCaptureVsanPcap VsanQueryClusterCaptureVsanPcapRequestType

type VsanQueryClusterCaptureVsanPcapRequestType struct {
	This               types.ManagedObjectReference   `xml:"_this"`
	Hosts              []string                       `xml:"hosts"`
	EsxRootPassword    string                         `xml:"esxRootPassword"`
	Duration           int32                          `xml:"duration"`
	Vmknic             []VsanClusterHostVmknicMapping `xml:"vmknic,omitempty"`
	IncludeRawPcap     *bool                          `xml:"includeRawPcap"`
	IncludeIgmp        *bool                          `xml:"includeIgmp"`
	CmmdsMsgTypeFilter []string                       `xml:"cmmdsMsgTypeFilter,omitempty"`
	CmmdsPorts         []int32                        `xml:"cmmdsPorts,omitempty"`
	ClusterUuid        string                         `xml:"clusterUuid,omitempty"`
}

type VsanCheckClusterClomdLivenessResponse struct {
	Returnval VsanClusterClomdLivenessResult `xml:"returnval"`
}

type VsanCheckClusterClomdLiveness VsanCheckClusterClomdLivenessRequestType

type VsanCheckClusterClomdLivenessRequestType struct {
	This            types.ManagedObjectReference `xml:"_this"`
	Hosts           []string                     `xml:"hosts"`
	EsxRootPassword string                       `xml:"esxRootPassword"`
}

type VsanClusterHealthSystem struct {
	types.ManagedObjectReference
}

type VsanQueryMobStatusResponse struct {
	Returnval string `xml:"returnval"`
}

type VsanQueryMobStatus VsanQueryMobStatusRequestType

type VsanQueryMobStatusRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
}

type VsanQueryMemoryStatsResponse struct {
	Returnval string `xml:"returnval"`
}

type VsanQueryMemoryStats VsanQueryMemoryStatsRequestType

type VsanQueryMemoryStatsRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
}

type VsanStopMobServiceResponse struct {
}

type VsanStopMobService VsanStopMobServiceRequestType

type VsanStopMobServiceRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
}

type VsanStartMobServiceResponse struct {
}

type VsanStartMobService VsanStartMobServiceRequestType

type VsanStartMobServiceRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
}

type VsanDebugSystem struct {
	types.ManagedObjectReference

	Entity []types.ManagedObjectReference `xml:"entity"`
}

type VsanFlrServiceGetStateResponse struct {
	Returnval VsanFlrHealthState `xml:"returnval"`
}

type VsanFlrServiceGetState VsanFlrServiceGetStateRequestType

type VsanFlrServiceGetStateRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Session *VsanFlrSession              `xml:"session,omitempty"`
}

type VsanFlrServiceDestroySessionResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type VsanFlrServiceDestroySession VsanFlrServiceDestroySessionRequestType

type VsanFlrServiceDestroySessionRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Session string                       `xml:"session"`
}

type VsanFlrServiceCreateSessionResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type VsanFlrServiceCreateSession VsanFlrServiceCreateSessionRequestType

type VsanFlrServiceCreateSessionRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Cluster types.ManagedObjectReference `xml:"cluster"`
	Spec    VsanFlrSessionSpec           `xml:"spec"`
}

type VsanFlrService struct {
	types.ManagedObjectReference
}

type VsanGetClusterPhoneHomeDataResponse struct {
	Returnval string `xml:"returnval"`
}

type VsanGetClusterPhoneHomeData VsanGetClusterPhoneHomeDataRequestType

type VsanGetClusterPhoneHomeDataRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Cluster types.ManagedObjectReference `xml:"cluster"`
}

type VsanQueryClusterHostLogBundleResponse struct {
	Returnval string `xml:"returnval"`
}

type VsanQueryClusterHostLogBundle VsanQueryClusterHostLogBundleRequestType

type VsanQueryClusterHostLogBundleRequestType struct {
	This      types.ManagedObjectReference `xml:"_this"`
	Host      types.ManagedObjectReference `xml:"host"`
	Manifests []string                     `xml:"manifests"`
}

type VsanPhoneHomeStatusResponse struct {
	Returnval string `xml:"returnval"`
}

type VsanPhoneHomeStatus VsanPhoneHomeStatusRequestType

type VsanPhoneHomeStatusRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
}

type VsanVcQueryClusterHostLogBundleResponse struct {
	Returnval string `xml:"returnval"`
}

type VsanVcQueryClusterHostLogBundle VsanVcQueryClusterHostLogBundleRequestType

type VsanVcQueryClusterHostLogBundleRequestType struct {
	This      types.ManagedObjectReference `xml:"_this"`
	Cluster   types.ManagedObjectReference `xml:"cluster"`
	Manifests []string                     `xml:"manifests"`
}

type VsanVcSetClusterLogsCatalogResponse struct {
	Returnval string `xml:"returnval"`
}

type VsanVcSetClusterLogsCatalog VsanVcSetClusterLogsCatalogRequestType

type VsanVcSetClusterLogsCatalogRequestType struct {
	This               types.ManagedObjectReference `xml:"_this"`
	Cluster            types.ManagedObjectReference `xml:"cluster"`
	LogsCatalog        string                       `xml:"logsCatalog"`
	LogsCatalogVersion string                       `xml:"logsCatalogVersion,omitempty"`
}

type VsanTriggerDiagnosticCollectionResponse struct {
}

type VsanTriggerDiagnosticCollection VsanTriggerDiagnosticCollectionRequestType

type VsanTriggerDiagnosticCollectionRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Cluster types.ManagedObjectReference `xml:"cluster"`
}

type QueryVsanCloudHealthDiagnosticDataResponse struct {
	Returnval string `xml:"returnval"`
}

type QueryVsanCloudHealthDiagnosticData QueryVsanCloudHealthDiagnosticDataRequestType

type QueryVsanCloudHealthDiagnosticDataRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
}

type VsanGetCoreDumpPartitionInfoResponse struct {
	Returnval types.HostDiskPartitionInfo `xml:"returnval"`
}

type VsanGetCoreDumpPartitionInfo VsanGetCoreDumpPartitionInfoRequestType

type VsanGetCoreDumpPartitionInfoRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
	Host types.ManagedObjectReference `xml:"host"`
}

type QueryVsanCloudHealthStatusResponse struct {
	Returnval VsanCloudHealthStatus `xml:"returnval"`
}

type QueryVsanCloudHealthStatus QueryVsanCloudHealthStatusRequestType

type QueryVsanCloudHealthStatusRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
}

type VsanQueryCMMDSPhoneHomeDataResponse struct {
	Returnval string `xml:"returnval"`
}

type VsanQueryCMMDSPhoneHomeData VsanQueryCMMDSPhoneHomeDataRequestType

type VsanQueryCMMDSPhoneHomeDataRequestType struct {
	This       types.ManagedObjectReference `xml:"_this"`
	Cluster    types.ManagedObjectReference `xml:"cluster"`
	QueryTypes []string                     `xml:"queryTypes,omitempty"`
}

type VsanQueryVcLogBundleResponse struct {
	Returnval string `xml:"returnval"`
}

type VsanQueryVcLogBundle VsanQueryVcLogBundleRequestType

type VsanQueryVcLogBundleRequestType struct {
	This      types.ManagedObjectReference `xml:"_this"`
	Manifests []string                     `xml:"manifests"`
}

type VsanSetLastHostBundleCollectionTimeResponse struct {
	Returnval string `xml:"returnval"`
}

type VsanSetLastHostBundleCollectionTime VsanSetLastHostBundleCollectionTimeRequestType

type VsanSetLastHostBundleCollectionTimeRequestType struct {
	This               types.ManagedObjectReference `xml:"_this"`
	Host               types.ManagedObjectReference `xml:"host"`
	LastCollectionTime string                       `xml:"lastCollectionTime"`
}

type VsanGetPhoneHomeObfuscationMapResponse struct {
	Returnval string `xml:"returnval"`
}

type VsanGetPhoneHomeObfuscationMap VsanGetPhoneHomeObfuscationMapRequestType

type VsanGetPhoneHomeObfuscationMapRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Cluster types.ManagedObjectReference `xml:"cluster"`
}

type VsanEnablePerfNetworkDiagnosticResultsResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type VsanEnablePerfNetworkDiagnosticResults VsanEnablePerfNetworkDiagnosticResultsRequestType

type VsanEnablePerfNetworkDiagnosticResultsRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Cluster types.ManagedObjectReference `xml:"cluster"`
}

type VsanPhoneHomeQueryPhysicalVsanDisksResponse struct {
	Returnval string `xml:"returnval"`
}

type VsanPhoneHomeQueryPhysicalVsanDisks VsanPhoneHomeQueryPhysicalVsanDisksRequestType

type VsanPhoneHomeQueryPhysicalVsanDisksRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Cluster types.ManagedObjectReference `xml:"cluster"`
	Props   []string                     `xml:"props,omitempty"`
}

type VsanCountLogTraceLinesResponse struct {
	Returnval string `xml:"returnval"`
}

type VsanCountLogTraceLines VsanCountLogTraceLinesRequestType

type VsanCountLogTraceLinesRequestType struct {
	This            types.ManagedObjectReference `xml:"_this"`
	Host            types.ManagedObjectReference `xml:"host"`
	CountSpec       string                       `xml:"countSpec"`
	ResetCollection *bool                        `xml:"resetCollection"`
}

type VsanPerformOnlineHealthCheckResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type VsanPerformOnlineHealthCheck VsanPerformOnlineHealthCheckRequestType

type VsanPerformOnlineHealthCheckRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Cluster types.ManagedObjectReference `xml:"cluster"`
}

type VsanPhoneHomeGetEsxCliDataResponse struct {
	Returnval string `xml:"returnval"`
}

type VsanPhoneHomeGetEsxCliData VsanPhoneHomeGetEsxCliDataRequestType

type VsanPhoneHomeGetEsxCliDataRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Cluster types.ManagedObjectReference `xml:"cluster"`
	Cmds    []string                     `xml:"cmds"`
}

type QueryVsanCloudHealthClusterHostsProfilingDataResponse struct {
	Returnval string `xml:"returnval"`
}

type QueryVsanCloudHealthClusterHostsProfilingData QueryVsanCloudHealthClusterHostsProfilingDataRequestType

type QueryVsanCloudHealthClusterHostsProfilingDataRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Cluster types.ManagedObjectReference `xml:"cluster"`
}

type VsanDisablePerfNetworkDiagnosticResultsResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type VsanDisablePerfNetworkDiagnosticResults VsanDisablePerfNetworkDiagnosticResultsRequestType

type VsanDisablePerfNetworkDiagnosticResultsRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Cluster types.ManagedObjectReference `xml:"cluster"`
}

type VsanSetPhoneHomeResourceIdResponse struct {
}

type VsanSetPhoneHomeResourceId VsanSetPhoneHomeResourceIdRequestType

type VsanSetPhoneHomeResourceIdRequestType struct {
	This       types.ManagedObjectReference `xml:"_this"`
	ResourceId string                       `xml:"resourceId"`
}

type VsanResetLastHostBundleCollectionTimeResponse struct {
}

type VsanResetLastHostBundleCollectionTime VsanResetLastHostBundleCollectionTimeRequestType

type VsanResetLastHostBundleCollectionTimeRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Cluster types.ManagedObjectReference `xml:"cluster"`
}

type VsanVcQueryClusterLogsCatalogVersionResponse struct {
	Returnval string `xml:"returnval"`
}

type VsanVcQueryClusterLogsCatalogVersion VsanVcQueryClusterLogsCatalogVersionRequestType

type VsanVcQueryClusterLogsCatalogVersionRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Cluster types.ManagedObjectReference `xml:"cluster"`
}

type VsanQueryHostPatchResponse struct {
	Returnval types.HostPatchManagerResult `xml:"returnval"`
}

type VsanQueryHostPatch VsanQueryHostPatchRequestType

type VsanQueryHostPatchRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
	Host types.ManagedObjectReference `xml:"host"`
}

type QueryVsanCloudHealthProfilingDataResponse struct {
	Returnval string `xml:"returnval"`
}

type QueryVsanCloudHealthProfilingData QueryVsanCloudHealthProfilingDataRequestType

type QueryVsanCloudHealthProfilingDataRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
}

type VsanCloudHealthRunDaemonResponse struct {
	Returnval string `xml:"returnval"`
}

type VsanCloudHealthRunDaemon VsanCloudHealthRunDaemonRequestType

type VsanCloudHealthRunDaemonRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
}

type VsanPhoneHomeSystem struct {
	types.ManagedObjectReference
}

type VimVsanDataEfficiencyCapacityState struct {
	types.DynamicData

	LogicalCapacity      int64 `xml:"logicalCapacity,omitempty"`
	LogicalCapacityUsed  int64 `xml:"logicalCapacityUsed,omitempty"`
	PhysicalCapacity     int64 `xml:"physicalCapacity,omitempty"`
	PhysicalCapacityUsed int64 `xml:"physicalCapacityUsed,omitempty"`
	DedupMetadataSize    int64 `xml:"dedupMetadataSize,omitempty"`
}

type VsanEncryptionHealthSummary struct {
	types.DynamicData

	Hostname         string                     `xml:"hostname,omitempty"`
	EncryptionInfo   *VsanHostEncryptionInfo    `xml:"encryptionInfo,omitempty"`
	OverallKmsHealth string                     `xml:"overallKmsHealth"`
	KmsHealth        []VsanKmsHealth            `xml:"kmsHealth,omitempty"`
	EncryptionIssues []string                   `xml:"encryptionIssues,omitempty"`
	DiskResults      []VsanDiskEncryptionHealth `xml:"diskResults,omitempty"`
	Error            *types.MethodFault         `xml:"error,omitempty"`
	AesniEnabled     *bool                      `xml:"aesniEnabled"`
}

type VsanBurnInTest struct {
	types.DynamicData

	Testname string `xml:"testname"`
	Workload string `xml:"workload,omitempty"`
	Duration int64  `xml:"duration"`
	Result   string `xml:"result"`
}

type VsanWhatIfEvacDetail struct {
	types.DynamicData

	Success                        *bool    `xml:"success"`
	BytesToSync                    int64    `xml:"bytesToSync,omitempty"`
	InaccessibleObjects            []string `xml:"inaccessibleObjects,omitempty"`
	IncompliantObjects             []string `xml:"incompliantObjects,omitempty"`
	ExtraSpaceNeeded               int64    `xml:"extraSpaceNeeded,omitempty"`
	FailedDueToInaccessibleObjects *bool    `xml:"failedDueToInaccessibleObjects"`
}

type VsanPerfTimeRange struct {
	types.DynamicData

	Name      string    `xml:"name"`
	StartTime time.Time `xml:"startTime"`
	EndTime   time.Time `xml:"endTime"`
}

type VsanIscsiLUNCommonInfo struct {
	types.DynamicData

	LunId   int32  `xml:"lunId,omitempty"`
	Alias   string `xml:"alias,omitempty"`
	LunSize int64  `xml:"lunSize"`
	Status  string `xml:"status,omitempty"`
}

type VimClusterVsanFaultDomainSpec struct {
	types.DynamicData

	Hosts []types.ManagedObjectReference `xml:"hosts"`
	Name  string                         `xml:"name"`
}

type VsanUnsupportedHighDiskVersionIssue struct {
	types.VsanUpgradeSystemPreflightCheckIssue

	Hosts []types.ManagedObjectReference `xml:"hosts"`
}

type VsanClusterHealthGroup struct {
	types.DynamicData

	GroupId      string                        `xml:"groupId"`
	GroupName    string                        `xml:"groupName"`
	GroupHealth  string                        `xml:"groupHealth"`
	GroupTests   []VsanClusterHealthTest       `xml:"groupTests,omitempty"`
	GroupDetails []VsanClusterHealthResultBase `xml:"groupDetails,omitempty"`
}

type VsanHostFwComponent struct {
	types.DynamicData

	Name             string   `xml:"name"`
	Url              string   `xml:"url,omitempty"`
	Sha1sum          string   `xml:"sha1sum,omitempty"`
	CurrentVersion   string   `xml:"currentVersion,omitempty"`
	SuggestedVersion string   `xml:"suggestedVersion,omitempty"`
	ComponentID      []string `xml:"componentID,omitempty"`
}

type VsanSpaceUsageDetailResult struct {
	types.DynamicData

	SpaceUsageByObjectType []VsanObjectSpaceSummary `xml:"spaceUsageByObjectType,omitempty"`
}

type VsanAttachToSrOperation struct {
	types.DynamicData

	Task      *types.ManagedObjectReference `xml:"task,omitempty"`
	Success   *bool                         `xml:"success"`
	Timestamp *time.Time                    `xml:"timestamp"`
	SrNumber  string                        `xml:"srNumber"`
}

type VsanPerfsvcConfig struct {
	types.DynamicData

	Enabled        bool                             `xml:"enabled"`
	Profile        *types.VirtualMachineProfileSpec `xml:"profile,omitempty"`
	DiagnosticMode *bool                            `xml:"diagnosticMode"`
	VerboseMode    *bool                            `xml:"verboseMode"`
}

type VsanObjectSpaceSummary struct {
	types.DynamicData

	ObjType            string                             `xml:"objType,omitempty"`
	OverheadB          int64                              `xml:"overheadB,omitempty"`
	TemporaryOverheadB int64                              `xml:"temporaryOverheadB,omitempty"`
	PrimaryCapacityB   int64                              `xml:"primaryCapacityB,omitempty"`
	ProvisionCapacityB int64                              `xml:"provisionCapacityB,omitempty"`
	ReservedCapacityB  int64                              `xml:"reservedCapacityB,omitempty"`
	OverReservedB      int64                              `xml:"overReservedB,omitempty"`
	PhysicalUsedB      int64                              `xml:"physicalUsedB,omitempty"`
	UsedB              int64                              `xml:"usedB,omitempty"`
	DpSpaceUsageInfo   *VsanDataProtectionSpaceUsageStats `xml:"dpSpaceUsageInfo,omitempty"`
}

type VsanClusterHealthResultValues struct {
	VsanClusterHealthResultBase

	Values []string `xml:"values,omitempty"`
}

type VsanPerfGraph struct {
	types.DynamicData

	Id          string             `xml:"id"`
	Metrics     []VsanPerfMetricId `xml:"metrics"`
	Unit        string             `xml:"unit"`
	Threshold   *VsanPerfThreshold `xml:"threshold,omitempty"`
	Name        string             `xml:"name,omitempty"`
	Description string             `xml:"description,omitempty"`
}

type VimClusterVSANStretchedClusterConfigIssueInfo struct {
	types.DynamicData

	Type    string                        `xml:"type"`
	Target  *types.ManagedObjectReference `xml:"target,omitempty"`
	Summary string                        `xml:"summary,omitempty"`
}

type VsanHclFirmwareUpdateSpec struct {
	types.DynamicData

	Host              types.ManagedObjectReference `xml:"host"`
	HbaDevice         string                       `xml:"hbaDevice"`
	FwFiles           []VsanHclFirmwareFile        `xml:"fwFiles"`
	AllowDowngrade    *bool                        `xml:"allowDowngrade"`
	FirmwareComponent []VsanHostFwComponent        `xml:"firmwareComponent,omitempty"`
}

type VsanVcsaBootstrapOntoVsanSpec struct {
	types.DynamicData

	VsanDiskMappingCreationSpec *VimVsanHostDiskMappingCreationSpec `xml:"vsanDiskMappingCreationSpec,omitempty"`
	VcConfig                    VsanVcsaDeploymentSpec              `xml:"vcConfig"`
	Portgroup                   types.ManagedObjectReference        `xml:"portgroup"`
	VmName                      string                              `xml:"vmName"`
	VsanBootstrapOnly           *bool                               `xml:"vsanBootstrapOnly"`
	VcsaOvaPath                 string                              `xml:"vcsaOvaPath,omitempty"`
	VcPostDeployConfig          *VsanVcPostDeployConfigSpec         `xml:"vcPostDeployConfig,omitempty"`
	TaskId                      string                              `xml:"taskId,omitempty"`
}

type VsanPerfEntityInfo struct {
	types.DynamicData

	EntityRefId        string `xml:"entityRefId"`
	EntityRefType      string `xml:"entityRefType,omitempty"`
	EntityName         string `xml:"entityName,omitempty"`
	EntityRelatedMoRef string `xml:"entityRelatedMoRef,omitempty"`
}

type CnsVolumeMetadata struct {
	types.DynamicData

	ContainerCluster CnsContainerCluster `xml:"containerCluster"`
	EntityMetadata   []CnsEntityMetadata `xml:"entityMetadata,omitempty"`
}

type VsanFlrMountSpec struct {
	types.DynamicData
}

type VsanClusterBalancePerDiskInfo struct {
	types.DynamicData

	Uuid                   string `xml:"uuid,omitempty"`
	Fullness               int64  `xml:"fullness"`
	Variance               int64  `xml:"variance"`
	FullnessAboveThreshold int64  `xml:"fullnessAboveThreshold"`
	DataToMoveB            int64  `xml:"dataToMoveB"`
}

type VsanUnicastAddressInfo struct {
	types.DynamicData

	Address string `xml:"address"`
	Port    int32  `xml:"port,omitempty"`
}

type VsanVmSpaceUsage struct {
	VsanContainerSpaceUsage

	Vm types.ManagedObjectReference `xml:"vm"`
}

type VsanFileServiceAdServiceHealthSummary struct {
	types.DynamicData

	DomainName   string `xml:"domainName,omitempty"`
	FileServerIp string `xml:"fileServerIp,omitempty"`
	Health       string `xml:"health,omitempty"`
	Description  string `xml:"description,omitempty"`
}

type CnsQuerySelection struct {
	types.DynamicData

	Names []string `xml:"names,omitempty"`
}

type VsanNetworkHealthResult struct {
	types.DynamicData

	Host              *types.ManagedObjectReference `xml:"host,omitempty"`
	Hostname          string                        `xml:"hostname,omitempty"`
	VsanVmknicPresent *bool                         `xml:"vsanVmknicPresent"`
	IpSubnets         []string                      `xml:"ipSubnets,omitempty"`
	IssueFound        *bool                         `xml:"issueFound"`
	PeerHealth        []VsanNetworkPeerHealthResult `xml:"peerHealth,omitempty"`
	VMotionHealth     []VsanNetworkPeerHealthResult `xml:"vMotionHealth,omitempty"`
	MulticastConfig   string                        `xml:"multicastConfig,omitempty"`
	UnicastConfig     string                        `xml:"unicastConfig,omitempty"`
	InUnicast         *bool                         `xml:"inUnicast"`
}

type VsanIscsiTargetServiceDefaultConfigSpec struct {
	types.DynamicData

	NetworkInterface    string                   `xml:"networkInterface,omitempty"`
	Port                int32                    `xml:"port,omitempty"`
	IscsiTargetAuthSpec *VsanIscsiTargetAuthSpec `xml:"iscsiTargetAuthSpec,omitempty"`
}

type VsanStoragePolicyStatus struct {
	types.DynamicData

	Id            string `xml:"id,omitempty"`
	ExpectedValue string `xml:"expectedValue,omitempty"`
	CurrentValue  string `xml:"currentValue,omitempty"`
}

type VsanPerfMemberInfo struct {
	types.DynamicData

	Thumbprint          string                   `xml:"thumbprint"`
	MemberUuid          string                   `xml:"memberUuid,omitempty"`
	IsSupportUnicast    *bool                    `xml:"isSupportUnicast"`
	UnicastAddressInfos []VsanUnicastAddressInfo `xml:"unicastAddressInfos,omitempty"`
	Hostname            string                   `xml:"hostname,omitempty"`
}

type VsanPerfMetricId struct {
	types.DynamicData

	Label                  string `xml:"label"`
	Group                  string `xml:"group,omitempty"`
	RollupType             string `xml:"rollupType,omitempty"`
	StatsType              string `xml:"statsType,omitempty"`
	Name                   string `xml:"name,omitempty"`
	Description            string `xml:"description,omitempty"`
	MetricsCollectInterval int32  `xml:"metricsCollectInterval,omitempty"`
}

type VsanFileServiceIpConfig struct {
	types.HostIpConfig

	IsPrimary *bool  `xml:"isPrimary"`
	Gateway   string `xml:"gateway"`
}

type HostSpbmDatastoreInfo struct {
	types.DynamicData

	DatastoreUrl     string `xml:"datastoreUrl"`
	Namespace        string `xml:"namespace"`
	DefaultProfileId string `xml:"defaultProfileId"`
}

type VsanBrokenDiskChainIssue struct {
	types.VsanUpgradeSystemPreflightCheckIssue

	Uuids []string `xml:"uuids"`
}

type VsanMetricProfile struct {
	types.DynamicData

	AuthToken string `xml:"authToken"`
}

type VsanFlrRetrieveFileSpec struct {
	VsanFlrQueryFileSpec
}

type VsanDiskGroupResourceCheckResult struct {
	EntityResourceCheckDetails

	CacheTierDisk     *VsanDiskResourceCheckResult  `xml:"cacheTierDisk,omitempty"`
	CapacityTierDisks []VsanDiskResourceCheckResult `xml:"capacityTierDisks,omitempty"`
}

type VsanDiskRebalanceResult struct {
	types.DynamicData

	Status               string  `xml:"status"`
	BytesMoving          int64   `xml:"bytesMoving,omitempty"`
	RemainingBytesToMove int64   `xml:"remainingBytesToMove,omitempty"`
	DiskUsage            float32 `xml:"diskUsage,omitempty"`
	MaxDiskUsage         float32 `xml:"maxDiskUsage,omitempty"`
	MinDiskUsage         float32 `xml:"minDiskUsage,omitempty"`
	AvgDiskUsage         float32 `xml:"avgDiskUsage,omitempty"`
}

type VsanContainerSpaceUsage struct {
	types.DynamicData

	ContainerId            string                   `xml:"containerId,omitempty"`
	SpaceUsageByObjectType []VsanObjectSpaceSummary `xml:"spaceUsageByObjectType,omitempty"`
}

type VsanFileServiceDomainConfig struct {
	types.DynamicData

	Name               string                    `xml:"name,omitempty"`
	SecurityFlavor     string                    `xml:"securityFlavor,omitempty"`
	FileServerIpConfig []VsanFileServiceIpConfig `xml:"fileServerIpConfig,omitempty"`
}

type VsanHostPropertyRetrieveIssue struct {
	types.VsanUpgradeSystemPreflightCheckIssue

	Hosts []types.ManagedObjectReference `xml:"hosts"`
}

type VsanVsanClusterPcapResult struct {
	types.DynamicData

	Pkts        []string                   `xml:"pkts,omitempty"`
	Groups      []VsanVsanClusterPcapGroup `xml:"groups,omitempty"`
	Issues      []string                   `xml:"issues,omitempty"`
	HostResults []VsanVsanPcapResult       `xml:"hostResults,omitempty"`
}

type VsanPerfDiagnoseQuerySpec struct {
	types.DynamicData

	StartTime time.Time `xml:"startTime"`
	EndTime   time.Time `xml:"endTime"`
	QueryType string    `xml:"queryType"`
	Context   string    `xml:"context,omitempty"`
}

type VsanDiskFormatConversionCheckResult struct {
	types.VsanUpgradeSystemPreflightCheckResult

	IsSupported            bool  `xml:"isSupported"`
	TargetVersion          int32 `xml:"targetVersion,omitempty"`
	IsDataMovementRequired *bool `xml:"isDataMovementRequired"`
}

type VimHostVsanRecoveryStatus struct {
	types.DynamicData

	Speed          int64  `xml:"speed"`
	Progress       int32  `xml:"progress"`
	TotalBytes     int64  `xml:"totalBytes"`
	CompletedBytes int64  `xml:"completedBytes"`
	Message        string `xml:"message"`
	HadError       bool   `xml:"hadError"`
}

type DataProtectionLoadBalancersInfo struct {
	types.DynamicData

	BasicLoadBalancerInfo    []DataProtectionLoadBalancerBasicInfo    `xml:"basicLoadBalancerInfo,omitempty"`
	DetailedLoadBalancerInfo []DataProtectionLoadBalancerDetailedInfo `xml:"detailedLoadBalancerInfo,omitempty"`
}

type VsanArchivalAccessibilityStatus struct {
	types.DynamicData

	ArchivalUrl string             `xml:"archivalUrl"`
	HealthState string             `xml:"healthState"`
	Error       *types.MethodFault `xml:"error,omitempty"`
}

type VsanHostVsanObjectSyncState struct {
	types.DynamicData

	Uuid       string                       `xml:"uuid"`
	Components []VsanHostComponentSyncState `xml:"components"`
}

type VsanVibSpec struct {
	types.DynamicData

	Host        types.ManagedObjectReference `xml:"host"`
	MetaUrl     string                       `xml:"metaUrl,omitempty"`
	MetaSha1Sum string                       `xml:"metaSha1Sum,omitempty"`
	VibUrl      string                       `xml:"vibUrl"`
	VibSha1Sum  string                       `xml:"vibSha1Sum"`
}

type VsanResourceCheckSpec struct {
	types.DynamicData

	Operation       string                        `xml:"operation"`
	Entities        []string                      `xml:"entities,omitempty"`
	MaintenanceSpec *types.HostMaintenanceSpec    `xml:"maintenanceSpec,omitempty"`
	Parent          *types.ManagedObjectReference `xml:"parent,omitempty"`
}

type VsanClusterHealthResultKeyValuePair struct {
	types.DynamicData

	Key   string `xml:"key,omitempty"`
	Value string `xml:"value,omitempty"`
}

type VsanVnicVdsMigrationSpec struct {
	types.DynamicData

	Key        int32                              `xml:"key"`
	VdsBacking types.BaseVirtualDeviceBackingInfo `xml:"vdsBacking"`
}

type VsanSpaceQuerySpec struct {
	types.DynamicData

	ContainerType string   `xml:"containerType"`
	ContainerIds  []string `xml:"containerIds,omitempty"`
}

type VsanFaultDomainResourceCheckResult struct {
	EntityResourceCheckDetails

	Hosts []VsanHostResourceCheckResult `xml:"hosts,omitempty"`
}

type VsanFlrQueryVolumesResult struct {
	types.DynamicData

	Volumes []VsanFlrVolume `xml:"volumes,omitempty"`
}

type VsanVcKmipServersHealth struct {
	types.DynamicData

	Health               string             `xml:"health,omitempty"`
	Error                *types.MethodFault `xml:"error,omitempty"`
	KmsProviderId        string             `xml:"kmsProviderId,omitempty"`
	KmsHealth            []VsanKmsHealth    `xml:"kmsHealth,omitempty"`
	ClientCertHealth     string             `xml:"clientCertHealth,omitempty"`
	ClientCertExpireDate *time.Time         `xml:"clientCertExpireDate"`
	IsAwsKms             *bool              `xml:"isAwsKms"`
	CmkHealth            string             `xml:"cmkHealth,omitempty"`
}

type VimHostVsanRecoveryObjectInfo struct {
	types.DynamicData

	Uuid        string                             `xml:"uuid"`
	Recoverable bool                               `xml:"recoverable"`
	Flags       VimHostVsanRecoveryFlags           `xml:"flags"`
	Errs        []string                           `xml:"errs,omitempty"`
	Msgs        []string                           `xml:"msgs,omitempty"`
	Props       VimHostVsanRecoveryObjectProps     `xml:"props"`
	Comps       []VimHostVsanRecoveryComponentInfo `xml:"comps,omitempty"`
}

type VsanClusterDatastoreUsageResult struct {
	types.DynamicData

	DatastoreUsageResult            []VsanHostDatastoreUsageResult `xml:"datastoreUsageResult,omitempty"`
	ClusterConfiguredUsageThreshold int32                          `xml:"clusterConfiguredUsageThreshold,omitempty"`
	IssueFound                      bool                           `xml:"issueFound"`
}

type VsanFlrHealthState struct {
	types.DynamicData

	Status string `xml:"status,omitempty"`
}

type VsanConfigCheckResult struct {
	types.DynamicData

	VsanEnabled bool                  `xml:"vsanEnabled"`
	Issues      []VsanConfigBaseIssue `xml:"issues,omitempty"`
}

type VsanFileServicePreflightCheckResult struct {
	types.DynamicData

	OvfInstalled          string `xml:"ovfInstalled,omitempty"`
	HostVersion           string `xml:"hostVersion,omitempty"`
	MixedModeIssue        string `xml:"mixedModeIssue,omitempty"`
	NetworkPartitionIssue string `xml:"networkPartitionIssue,omitempty"`
	VsanDatastoreIssue    string `xml:"vsanDatastoreIssue,omitempty"`
	DomainConfigIssue     string `xml:"domainConfigIssue,omitempty"`
}

type VsanNetworkConfigPortgroupWithNoRedundancyIssue struct {
	VsanNetworkConfigBaseIssue

	Host          types.ManagedObjectReference  `xml:"host"`
	PortgroupName string                        `xml:"portgroupName,omitempty"`
	Vds           *types.ManagedObjectReference `xml:"vds,omitempty"`
	Pg            *types.ManagedObjectReference `xml:"pg,omitempty"`
	NumPnics      int64                         `xml:"numPnics"`
}

type VsanNetworkVMotionVmknicNotFountIssue struct {
	VsanNetworkConfigBaseIssue

	HostWithoutVmotionVmknic types.ManagedObjectReference `xml:"hostWithoutVmotionVmknic"`
}

type DataProtectionLoadBalancerBasicInfo struct {
	types.DynamicData

	HostUuid string `xml:"hostUuid"`
}

type ResyncIopsInfo struct {
	types.DynamicData

	ResyncIops int32 `xml:"resyncIops"`
}

type VimVsanReconfigSpec struct {
	types.SDDCBase

	VsanClusterConfig      *types.VsanClusterConfigInfo          `xml:"vsanClusterConfig,omitempty"`
	DataEfficiencyConfig   *VsanDataEfficiencyConfig             `xml:"dataEfficiencyConfig,omitempty"`
	DiskMappingSpec        *VimClusterVsanDiskMappingsConfigSpec `xml:"diskMappingSpec,omitempty"`
	FaultDomainsSpec       *VimClusterVsanFaultDomainsConfigSpec `xml:"faultDomainsSpec,omitempty"`
	Modify                 bool                                  `xml:"modify"`
	AllowReducedRedundancy *bool                                 `xml:"allowReducedRedundancy"`
	ResyncIopsLimitConfig  *ResyncIopsInfo                       `xml:"resyncIopsLimitConfig,omitempty"`
	IscsiSpec              *VsanIscsiTargetServiceSpec           `xml:"iscsiSpec,omitempty"`
	DataEncryptionConfig   *VsanDataEncryptionConfig             `xml:"dataEncryptionConfig,omitempty"`
	ExtendedConfig         *VsanExtendedConfig                   `xml:"extendedConfig,omitempty"`
	DataProtectionConfig   *VsanDataProtectionInfo               `xml:"dataProtectionConfig,omitempty"`
	DatastoreConfig        *VsanDatastoreConfig                  `xml:"datastoreConfig,omitempty"`
	PerfsvcConfig          *VsanPerfsvcConfig                    `xml:"perfsvcConfig,omitempty"`
	BmcConfig              []VimClusterVsanBmcSpec               `xml:"bmcConfig,omitempty"`
	FileServiceConfig      *VsanFileServiceConfig                `xml:"fileServiceConfig,omitempty"`
	UnmapConfig            *VsanUnmapConfig                      `xml:"unmapConfig,omitempty"`
	RdmaConfig             *VsanRdmaConfig                       `xml:"rdmaConfig,omitempty"`
	VumConfig              *VsanVumConfig                        `xml:"vumConfig,omitempty"`
	MetricsConfig          *VsanMetricsConfig                    `xml:"metricsConfig,omitempty"`
}

type VsanClusterDataProtectionCfgSyncResult struct {
	types.DynamicData

	InSync       bool                                         `xml:"inSync"`
	Name         string                                       `xml:"name"`
	ClusterValue string                                       `xml:"clusterValue"`
	HostValues   []VsanClusterDataProtectionCfgSyncHostResult `xml:"hostValues,omitempty"`
}

type VsanHistoryItemQuerySpec struct {
	types.DynamicData

	Clusters []types.ManagedObjectReference `xml:"clusters,omitempty"`
	CleanAll *bool                          `xml:"cleanAll"`
	Start    *time.Time                     `xml:"start"`
	End      *time.Time                     `xml:"end"`
}

type VimVsanHostDiskMappingCreationSpec struct {
	types.DynamicData

	Host          types.ManagedObjectReference `xml:"host"`
	CacheDisks    []types.HostScsiDisk         `xml:"cacheDisks,omitempty"`
	CapacityDisks []types.HostScsiDisk         `xml:"capacityDisks"`
	CreationType  string                       `xml:"creationType"`
}

type VsanObjectTypeRule struct {
	types.DynamicData

	ObjectType string   `xml:"objectType,omitempty"`
	Attributes []string `xml:"attributes,omitempty"`
}

type VsanResourceCheckResult struct {
	EntityResourceCheckDetails

	Timestamp           time.Time                            `xml:"timestamp"`
	Status              string                               `xml:"status"`
	Messages            []types.LocalizableMessage           `xml:"messages,omitempty"`
	FaultDomains        []VsanFaultDomainResourceCheckResult `xml:"faultDomains,omitempty"`
	DataToMove          int64                                `xml:"dataToMove,omitempty"`
	NonCompliantObjects []string                             `xml:"nonCompliantObjects,omitempty"`
	InaccessibleObjects []string                             `xml:"inaccessibleObjects,omitempty"`
	CapacityThreshold   *VsanHealthThreshold                 `xml:"capacityThreshold,omitempty"`
	Health              *VsanClusterHealthSummary            `xml:"health,omitempty"`
}

type VsanDataProtectionSpaceUsage struct {
	types.DynamicData

	Error                *types.MethodFault                 `xml:"error,omitempty"`
	Warning              *types.MethodFault                 `xml:"warning,omitempty"`
	ProtectionTargetInfo VsanDataProtectionTargetSiteInfo   `xml:"protectionTargetInfo"`
	DpUsageStats         *VsanDataProtectionSpaceUsageStats `xml:"dpUsageStats,omitempty"`
}

type CnsStoreFault struct {
	CnsFault
}

type VsanHostVsanObjectSyncQueryResult struct {
	types.DynamicData

	TotalObjectsToSync           int64                             `xml:"totalObjectsToSync,omitempty"`
	TotalBytesToSync             int64                             `xml:"totalBytesToSync,omitempty"`
	TotalRecoveryETA             int64                             `xml:"totalRecoveryETA,omitempty"`
	Objects                      []VsanHostVsanObjectSyncState     `xml:"objects,omitempty"`
	SyncingObjectRecoveryDetails *VsanSyncingObjectRecoveryDetails `xml:"syncingObjectRecoveryDetails,omitempty"`
}

type VsanLimitHealthResult struct {
	types.DynamicData

	Hostname                string `xml:"hostname,omitempty"`
	IssueFound              bool   `xml:"issueFound"`
	MaxComponents           int32  `xml:"maxComponents"`
	FreeComponents          int32  `xml:"freeComponents"`
	ComponentLimitHealth    string `xml:"componentLimitHealth"`
	LowestFreeDiskSpacePct  int32  `xml:"lowestFreeDiskSpacePct"`
	UsedDiskSpaceB          int64  `xml:"usedDiskSpaceB"`
	TotalDiskSpaceB         int64  `xml:"totalDiskSpaceB"`
	DiskFreeSpaceHealth     string `xml:"diskFreeSpaceHealth"`
	ReservedRcSizeB         int64  `xml:"reservedRcSizeB"`
	TotalRcSizeB            int64  `xml:"totalRcSizeB"`
	RcFreeReservationHealth string `xml:"rcFreeReservationHealth"`
	CdReservedSizeB         int64  `xml:"cdReservedSizeB,omitempty"`
}

type VsanIscsiTargetServiceConfig struct {
	types.DynamicData

	DefaultConfig *VsanIscsiTargetServiceDefaultConfigSpec `xml:"defaultConfig,omitempty"`
	Enabled       *bool                                    `xml:"enabled"`
}

type CnsQueryResult struct {
	types.DynamicData

	Volumes []CnsVolume `xml:"volumes,omitempty"`
	Cursor  CnsCursor   `xml:"cursor"`
}

type CnsVSANFileCreateSpec struct {
	CnsFileCreateSpec

	SoftQuotaInMb int64                        `xml:"softQuotaInMb,omitempty"`
	Permission    []VsanFileShareNetPermission `xml:"permission,omitempty"`
}

type VsanObjectOverallHealth struct {
	types.DynamicData

	ObjectHealthDetail      []VsanObjectHealth            `xml:"objectHealthDetail,omitempty"`
	ObjectsComplianceDetail []VsanStorageComplianceResult `xml:"objectsComplianceDetail,omitempty"`
	ObjectVersionCompliance *bool                         `xml:"objectVersionCompliance"`
}

type VsanDataProtectionPairingInfo struct {
	types.DynamicData

	PairingId          string                                   `xml:"pairingId,omitempty"`
	PeerSite           *DataProtectionPeerSiteInfo              `xml:"peerSite,omitempty"`
	PeerClusterUuid    string                                   `xml:"peerClusterUuid"`
	PeerDatastoreUrl   string                                   `xml:"peerDatastoreUrl"`
	LocalDatastoreUrl  string                                   `xml:"localDatastoreUrl"`
	PeerClusterName    string                                   `xml:"peerClusterName,omitempty"`
	PeerDatastoreName  string                                   `xml:"peerDatastoreName,omitempty"`
	LocalLoadBalancers []DataProtectionLoadBalancerBasicInfo    `xml:"localLoadBalancers,omitempty"`
	PeerLoadBalancers  []DataProtectionLoadBalancerDetailedInfo `xml:"peerLoadBalancers,omitempty"`
	DeletePairing      *bool                                    `xml:"deletePairing"`
}

type VsanClusterHealthResultColumnInfo struct {
	types.DynamicData

	Label string `xml:"label"`
	Type  string `xml:"type"`
}

type VsanVdsPgMigrationSpec struct {
	types.DynamicData

	VssPgName       string                       `xml:"vssPgName"`
	DvPgName        string                       `xml:"dvPgName"`
	VdsPgSetting    types.VMwareDVSPortSetting   `xml:"vdsPgSetting"`
	VdsPgType       string                       `xml:"vdsPgType"`
	Hosts           []VsanVdsPgMigrationHostInfo `xml:"hosts,omitempty"`
	CollisionRename bool                         `xml:"collisionRename"`
}

type VsanClusterNetworkHealthResult struct {
	types.DynamicData

	HostResults                []VsanNetworkHealthResult         `xml:"hostResults,omitempty"`
	IssueFound                 *bool                             `xml:"issueFound"`
	VsanVmknicPresent          *bool                             `xml:"vsanVmknicPresent"`
	MatchingMulticastConfig    *bool                             `xml:"matchingMulticastConfig"`
	MatchingIpSubnets          *bool                             `xml:"matchingIpSubnets"`
	PingTestSuccess            *bool                             `xml:"pingTestSuccess"`
	LargePingTestSuccess       *bool                             `xml:"largePingTestSuccess"`
	HostLatencyCheckSuccess    *bool                             `xml:"hostLatencyCheckSuccess"`
	PotentialMulticastIssue    *bool                             `xml:"potentialMulticastIssue"`
	OtherHostsInVsanCluster    []string                          `xml:"otherHostsInVsanCluster,omitempty"`
	Partitions                 []VsanClusterNetworkPartitionInfo `xml:"partitions,omitempty"`
	HostsWithVsanDisabled      []string                          `xml:"hostsWithVsanDisabled,omitempty"`
	HostsDisconnected          []string                          `xml:"hostsDisconnected,omitempty"`
	HostsCommFailure           []string                          `xml:"hostsCommFailure,omitempty"`
	HostsInEsxMaintenanceMode  []string                          `xml:"hostsInEsxMaintenanceMode,omitempty"`
	HostsInVsanMaintenanceMode []string                          `xml:"hostsInVsanMaintenanceMode,omitempty"`
	InfoAboutUnexpectedHosts   []VsanQueryResultHostInfo         `xml:"infoAboutUnexpectedHosts,omitempty"`
	ClusterInUnicastMode       *bool                             `xml:"clusterInUnicastMode"`
}

type HostSpbmPolicyInfo struct {
	types.DynamicData

	ProfileId      string                   `xml:"profileId"`
	Name           string                   `xml:"name"`
	Description    string                   `xml:"description,omitempty"`
	GenerationId   int64                    `xml:"generationId"`
	PolicyBlobInfo []HostSpbmPolicyBlobInfo `xml:"policyBlobInfo"`
}

type VsanIscsiInitiatorGroup struct {
	types.DynamicData

	Name       string                     `xml:"name"`
	Initiators []string                   `xml:"initiators,omitempty"`
	Targets    []VsanIscsiTargetBasicInfo `xml:"targets,omitempty"`
}

type VsanDiskUnhealthIssue struct {
	types.VsanUpgradeSystemPreflightCheckIssue

	Uuids []string `xml:"uuids"`
}

type VimHostVsanComponentAlignment struct {
	types.DynamicData

	UnalignedObjects               []string `xml:"unalignedObjects,omitempty"`
	ObjectsWithUnalignedComponents []string `xml:"objectsWithUnalignedComponents,omitempty"`
	TotalBytesToSync               int64    `xml:"totalBytesToSync,omitempty"`
}

type VsanResourceCheckStatus struct {
	types.DynamicData

	Status     string                        `xml:"status"`
	Result     *VsanResourceCheckResult      `xml:"result,omitempty"`
	Task       *VsanResourceCheckTaskDetails `xml:"task,omitempty"`
	ParentTask *VsanResourceCheckTaskDetails `xml:"parentTask,omitempty"`
}

type VsanVumConfig struct {
	types.DynamicData

	BaselinePreferenceType string `xml:"baselinePreferenceType,typeattr"`
}

type VsanDiskFormatConversionSpec struct {
	types.DynamicData

	DataEfficiencyConfig *VsanDataEfficiencyConfig `xml:"dataEfficiencyConfig,omitempty"`
	DataEncryptionConfig *VsanDataEncryptionConfig `xml:"dataEncryptionConfig,omitempty"`
	SkipHostRemediation  *bool                     `xml:"skipHostRemediation"`
	AllowDataMovement    *bool                     `xml:"allowDataMovement"`
}

type VsanObjectQuerySpec struct {
	types.DynamicData

	Uuid                    string `xml:"uuid"`
	SpbmProfileGenerationId string `xml:"spbmProfileGenerationId,omitempty"`
}

type EntityResourceCheckDetails struct {
	types.DynamicData

	Name                       string `xml:"name,omitempty"`
	Uuid                       string `xml:"uuid,omitempty"`
	IsNew                      *bool  `xml:"isNew"`
	Capacity                   int64  `xml:"capacity,omitempty"`
	PostOperationCapacity      int64  `xml:"postOperationCapacity,omitempty"`
	UsedCapacity               int64  `xml:"usedCapacity,omitempty"`
	PostOperationUsedCapacity  int64  `xml:"postOperationUsedCapacity,omitempty"`
	AdditionalRequiredCapacity int64  `xml:"additionalRequiredCapacity,omitempty"`
	MaxComponents              int64  `xml:"maxComponents,omitempty"`
	Components                 int64  `xml:"components,omitempty"`
}

type VsanFileServiceDomain struct {
	types.DynamicData

	Uuid   string                       `xml:"uuid"`
	Config *VsanFileServiceDomainConfig `xml:"config,omitempty"`
}

type CnsFault struct {
	types.VimFault

	Reason string `xml:"reason"`
}

type VsanResourceHealth struct {
	types.DynamicData

	Resource    string `xml:"resource"`
	Health      string `xml:"health"`
	Description string `xml:"description,omitempty"`
}

type VsanDatastoreConfig struct {
	types.DynamicData

	Datastores       []VsanDatastoreSpec       `xml:"datastores,omitempty"`
	RemoteDatastores []VsanRemoteDatastoreSpec `xml:"remoteDatastores,omitempty"`
}

type VimClusterVsanHostPersistedState struct {
	types.DynamicData

	ClusterUuid      string                `xml:"clusterUuid"`
	PasswdLastUpdate *time.Time            `xml:"passwdLastUpdate"`
	ConnectSpec      types.HostConnectSpec `xml:"connectSpec"`
}

type VsanClusterDataProtectionCfgSyncHostResult struct {
	types.DynamicData

	Hostname string `xml:"hostname"`
	Value    string `xml:"value"`
}

type VsanHostCimProviderInfo struct {
	types.DynamicData

	CimProviderSupported  *bool              `xml:"cimProviderSupported"`
	InstalledCIMProvider  string             `xml:"installedCIMProvider,omitempty"`
	CimProviderOnHcl      []string           `xml:"cimProviderOnHcl,omitempty"`
	CimProviderLinksOnHcl []VsanDownloadItem `xml:"cimProviderLinksOnHcl,omitempty"`
}

type VsanNestJsonComparator struct {
	VsanComparator

	NestedComparators []VsanJsonComparator `xml:"nestedComparators,omitempty"`
	Conjoiner         string               `xml:"conjoiner,omitempty"`
}

type VsanIscsiTargetSpec struct {
	VsanIscsiTargetCommonInfo

	StoragePolicy *types.VirtualMachineProfileSpec `xml:"storagePolicy,omitempty"`
	NewAlias      string                           `xml:"newAlias,omitempty"`
}

type VsanDataEncryptionConfig struct {
	types.DynamicData

	EncryptionEnabled   bool                 `xml:"encryptionEnabled"`
	KmsProviderId       *types.KeyProviderId `xml:"kmsProviderId,omitempty"`
	KekId               string               `xml:"kekId,omitempty"`
	HostKeyId           string               `xml:"hostKeyId,omitempty"`
	DekGenerationId     int64                `xml:"dekGenerationId,omitempty"`
	Changing            *bool                `xml:"changing"`
	EraseDisksBeforeUse *bool                `xml:"eraseDisksBeforeUse"`
}

type VsanKmsHealth struct {
	types.DynamicData

	ServerName     string             `xml:"serverName"`
	Health         string             `xml:"health"`
	Error          *types.MethodFault `xml:"error,omitempty"`
	TrustHealth    string             `xml:"trustHealth,omitempty"`
	CertHealth     string             `xml:"certHealth,omitempty"`
	CertExpireDate *time.Time         `xml:"certExpireDate"`
}

type VsanFileServiceShareHealthSummary struct {
	types.DynamicData

	OverallHealth string                   `xml:"overallHealth,omitempty"`
	DomainName    string                   `xml:"domainName,omitempty"`
	ShareUuid     string                   `xml:"shareUuid,omitempty"`
	ShareName     string                   `xml:"shareName,omitempty"`
	ObjectHealth  *VsanObjectOverallHealth `xml:"objectHealth,omitempty"`
	Description   string                   `xml:"description,omitempty"`
}

type VsanResourceCheckTaskDetails struct {
	types.DynamicData

	Task            types.ManagedObjectReference  `xml:"task"`
	Host            *types.ManagedObjectReference `xml:"host,omitempty"`
	HostUuid        string                        `xml:"hostUuid,omitempty"`
	MaintenanceSpec *types.HostMaintenanceSpec    `xml:"maintenanceSpec,omitempty"`
}

type VsanFlrQueryFileSpec struct {
	types.DynamicData

	VolumeId string `xml:"volumeId,omitempty"`
	FileId   string `xml:"fileId,omitempty"`
}

type VsanHealthThreshold struct {
	types.DynamicData

	YellowValue int64 `xml:"yellowValue"`
	RedValue    int64 `xml:"redValue"`
}

type VimClusterVSANStretchedClusterFaultDomainConfig struct {
	types.DynamicData

	FirstFdName   string                         `xml:"firstFdName"`
	FirstFdHosts  []types.ManagedObjectReference `xml:"firstFdHosts"`
	SecondFdName  string                         `xml:"secondFdName"`
	SecondFdHosts []types.ManagedObjectReference `xml:"secondFdHosts"`
}

type VsanUnmapConfig struct {
	types.DynamicData

	Enable bool `xml:"enable"`
}

type VsanDiskEncryptionHealth struct {
	types.DynamicData

	DiskHealth       *VsanPhysicalDiskHealth `xml:"diskHealth,omitempty"`
	EncryptionIssues []string                `xml:"encryptionIssues,omitempty"`
}

type VsanRuntimeStatsHostMap struct {
	types.DynamicData

	Host  types.ManagedObjectReference `xml:"host"`
	Stats *VsanHostRuntimeStats        `xml:"stats,omitempty"`
}

type VsanFileServiceRootFsHealth struct {
	types.DynamicData

	Created     *bool  `xml:"created"`
	Health      string `xml:"health,omitempty"`
	Description string `xml:"description,omitempty"`
}

type VsanResourceConstraint struct {
	types.DynamicData

	TargetType string `xml:"targetType,omitempty"`
}

type VimHostVSANCmmdsFaultDomainInfo struct {
	types.DynamicData

	FaultDomainId   string `xml:"faultDomainId"`
	FaultDomainName string `xml:"faultDomainName"`
}

type VsanClusterHealthConfigs struct {
	types.DynamicData

	EnableVsanTelemetry   *bool                                 `xml:"enableVsanTelemetry"`
	VsanTelemetryInterval int32                                 `xml:"vsanTelemetryInterval,omitempty"`
	VsanTelemetryProxy    *VsanClusterTelemetryProxyConfig      `xml:"vsanTelemetryProxy,omitempty"`
	Configs               []VsanClusterHealthResultKeyValuePair `xml:"configs,omitempty"`
}

type VsanObjectIdentityAndHealth struct {
	types.DynamicData

	Identities   []VsanObjectIdentity     `xml:"identities,omitempty"`
	Health       *VsanObjectOverallHealth `xml:"health,omitempty"`
	SpaceSummary []VsanObjectSpaceSummary `xml:"spaceSummary,omitempty"`
	RawData      string                   `xml:"rawData,omitempty"`
}

type VsanClusterHealthResultTable struct {
	VsanClusterHealthResultBase

	Columns []VsanClusterHealthResultColumnInfo `xml:"columns,omitempty"`
	Rows    []VsanClusterHealthResultRow        `xml:"rows,omitempty"`
}

type VsanStorageOperationalStatus struct {
	types.DynamicData

	Healthy           *bool      `xml:"healthy"`
	OperationETA      *time.Time `xml:"operationETA"`
	OperationProgress int64      `xml:"operationProgress,omitempty"`
	Transitional      *bool      `xml:"transitional"`
}

type VsanVcPostDeployConfigSpec struct {
	types.DynamicData

	DcName                   string                    `xml:"dcName,omitempty"`
	ClusterName              string                    `xml:"clusterName,omitempty"`
	FirstHost                *types.HostConnectSpec    `xml:"firstHost,omitempty"`
	HostsToAdd               []types.HostConnectSpec   `xml:"hostsToAdd,omitempty"`
	VsanDataEfficiencyConfig *VsanDataEfficiencyConfig `xml:"vsanDataEfficiencyConfig,omitempty"`
	VsanLicenseKey           string                    `xml:"vsanLicenseKey,omitempty"`
	HostLicenseKey           string                    `xml:"hostLicenseKey,omitempty"`
	TaskId                   string                    `xml:"taskId,omitempty"`
	VsanDataEncryptionConfig *VsanHostEncryptionInfo   `xml:"vsanDataEncryptionConfig,omitempty"`
}

type VsanSpaceUsage struct {
	types.DynamicData

	TotalCapacityB    int64                               `xml:"totalCapacityB"`
	FreeCapacityB     int64                               `xml:"freeCapacityB,omitempty"`
	SpaceOverview     *VsanObjectSpaceSummary             `xml:"spaceOverview,omitempty"`
	SpaceDetail       *VsanSpaceUsageDetailResult         `xml:"spaceDetail,omitempty"`
	EfficientCapacity *VimVsanDataEfficiencyCapacityState `xml:"efficientCapacity,omitempty"`
	WhatifCapacities  []VsanWhatifCapacity                `xml:"whatifCapacities,omitempty"`
	UncommittedB      int64                               `xml:"uncommittedB,omitempty"`
}

type VsanHclDiskInfo struct {
	types.DynamicData

	DeviceName       string              `xml:"deviceName"`
	Model            string              `xml:"model,omitempty"`
	IsSsd            *bool               `xml:"isSsd"`
	VsanDisk         bool                `xml:"vsanDisk"`
	Issues           []types.MethodFault `xml:"issues,omitempty"`
	RemediableIssues []string            `xml:"remediableIssues,omitempty"`
}

type VsanFlrRetrieveFileResult struct {
	types.DynamicData

	Url        string `xml:"url"`
	Thumbprint string `xml:"thumbprint,omitempty"`
}

type VsanFileServiceHealthSummary struct {
	types.DynamicData

	Hostname         string                                  `xml:"hostname,omitempty"`
	OverallHealth    string                                  `xml:"overallHealth,omitempty"`
	Enabled          *bool                                   `xml:"enabled"`
	VdfsdStatus      *VsanResourceHealth                     `xml:"vdfsdStatus,omitempty"`
	FsvmStatus       *VsanResourceHealth                     `xml:"fsvmStatus,omitempty"`
	RootFsStatus     *VsanFileServiceRootFsHealth            `xml:"rootFsStatus,omitempty"`
	AdServiceHealth  []VsanFileServiceAdServiceHealthSummary `xml:"adServiceHealth,omitempty"`
	FileServerHealth []VsanFileServerHealthSummary           `xml:"fileServerHealth,omitempty"`
	FileShareHealth  []VsanFileServiceShareHealthSummary     `xml:"fileShareHealth,omitempty"`
}

type VsanVmdkLoadTestSpec struct {
	types.DynamicData

	VmdkCreateSpec     *types.FileBackedVirtualDiskSpec `xml:"vmdkCreateSpec,omitempty"`
	VmdkIOSpec         *VsanVmdkIOLoadSpec              `xml:"vmdkIOSpec,omitempty"`
	VmdkIOSpecSequence []VsanVmdkIOLoadSpec             `xml:"vmdkIOSpecSequence,omitempty"`
	StepDurationSec    int64                            `xml:"stepDurationSec,omitempty"`
}

type VsanArchivalAccessibilityResult struct {
	types.DynamicData

	Hostname             string                            `xml:"hostname,omitempty"`
	ArchivalHealthStatus []VsanArchivalAccessibilityStatus `xml:"archivalHealthStatus,omitempty"`
}

type CnsVolumeExtendSpec struct {
	types.DynamicData

	VolumeId     CnsVolumeId `xml:"volumeId"`
	CapacityInMb int64       `xml:"capacityInMb"`
	VolumeType   string      `xml:"volumeType,omitempty"`
}

type VsanClusterProactiveTestResult struct {
	types.DynamicData

	OverallStatus            string                 `xml:"overallStatus"`
	OverallStatusDescription string                 `xml:"overallStatusDescription"`
	Timestamp                time.Time              `xml:"timestamp"`
	HealthTest               *VsanClusterHealthTest `xml:"healthTest,omitempty"`
}

type VimVsanHostVsanHostCapability struct {
	types.DynamicData

	Host        types.ManagedObjectReference `xml:"host"`
	IsSupported bool                         `xml:"isSupported"`
	IsLicensed  bool                         `xml:"isLicensed"`
}

type VimHostVsanRecoveryObjectProps struct {
	types.DynamicData

	Hosts               []string `xml:"hosts,omitempty"`
	ExtendedAttrib      string   `xml:"extendedAttrib"`
	Size                int64    `xml:"size"`
	HasQuorum           bool     `xml:"hasQuorum"`
	HasDataAvailability bool     `xml:"hasDataAvailability"`
	RecoverySupported   bool     `xml:"recoverySupported"`
	ConfigRootType      string   `xml:"configRootType"`
}

type VsanPerfThreshold struct {
	types.DynamicData

	Direction string `xml:"direction"`
	Yellow    string `xml:"yellow,omitempty"`
	Red       string `xml:"red,omitempty"`
}

type VsanHostIpConfigEx struct {
	types.VsanHostIpConfig

	UpstreamIpV6Address   string `xml:"upstreamIpV6Address,omitempty"`
	DownstreamIpV6Address string `xml:"downstreamIpV6Address,omitempty"`
}

type VsanSyncingObjectFilter struct {
	types.DynamicData

	ResyncType      string `xml:"resyncType,omitempty"`
	ResyncStatus    string `xml:"resyncStatus,omitempty"`
	NumberOfObjects int64  `xml:"numberOfObjects,omitempty"`
	Offset          int64  `xml:"offset,omitempty"`
}

type VsanSmartDiskStats struct {
	types.DynamicData

	Disk  string               `xml:"disk"`
	Stats []VsanSmartParameter `xml:"stats,omitempty"`
	Error *types.MethodFault   `xml:"error,omitempty"`
}

type CnsFileBackingDetails struct {
	CnsBackingObjectDetails

	BackingFileId string `xml:"backingFileId,omitempty"`
}

type VsanClusterVMsHealthOverallResult struct {
	types.DynamicData

	HealthStateList    []VsanClusterVMsHealthSummaryResult `xml:"healthStateList,omitempty"`
	OverallHealthState string                              `xml:"overallHealthState,omitempty"`
}

type CnsEntityMetadata struct {
	types.DynamicData

	EntityName       string               `xml:"entityName"`
	Labels           []types.KeyValue     `xml:"labels,omitempty"`
	Delete           *bool                `xml:"delete"`
	ContainerCluster *CnsContainerCluster `xml:"containerCluster,omitempty"`
}

type VsanHostHealthSystemStatusResult struct {
	types.DynamicData

	Hostname string   `xml:"hostname"`
	Status   string   `xml:"status"`
	Issues   []string `xml:"issues,omitempty"`
}

type VsanClusterAdvCfgSyncResult struct {
	types.DynamicData

	InSync     bool                              `xml:"inSync"`
	Name       string                            `xml:"name"`
	HostValues []VsanClusterAdvCfgSyncHostResult `xml:"hostValues,omitempty"`
}

type VsanQueryResultHostInfo struct {
	types.DynamicData

	Uuid              string   `xml:"uuid,omitempty"`
	HostnameInCmmds   string   `xml:"hostnameInCmmds,omitempty"`
	VsanIpv4Addresses []string `xml:"vsanIpv4Addresses,omitempty"`
}

type VimVsanVsanScanObjectsResult struct {
	types.DynamicData

	Issues  []VimVsanVsanScanObjectsIssue `xml:"issues,omitempty"`
	Fixed   []VimVsanVsanScanObjectsIssue `xml:"fixed,omitempty"`
	Fixable []VimVsanVsanScanObjectsIssue `xml:"fixable,omitempty"`
}

type VimVsanHostDiskMapInfoEx struct {
	types.DynamicData

	Mapping             types.VsanHostDiskMapping `xml:"mapping"`
	IsMounted           bool                      `xml:"isMounted"`
	UnlockedEncrypted   *bool                     `xml:"unlockedEncrypted"`
	IsAllFlash          bool                      `xml:"isAllFlash"`
	IsDataEfficiency    *bool                     `xml:"isDataEfficiency"`
	EncryptionInfo      *VsanDataEncryptionConfig `xml:"encryptionInfo,omitempty"`
	DiskgroupCapability []string                  `xml:"diskgroupCapability,omitempty"`
}

type VsanClusterHealthSummary struct {
	types.DynamicData

	ClusterStatus               *VsanClusterHealthSystemStatusResult     `xml:"clusterStatus,omitempty"`
	Timestamp                   *time.Time                               `xml:"timestamp"`
	ClusterVersions             *VsanClusterHealthSystemVersionResult    `xml:"clusterVersions,omitempty"`
	ObjectHealth                *VsanObjectOverallHealth                 `xml:"objectHealth,omitempty"`
	VmHealth                    *VsanClusterVMsHealthOverallResult       `xml:"vmHealth,omitempty"`
	NetworkHealth               *VsanClusterNetworkHealthResult          `xml:"networkHealth,omitempty"`
	LimitHealth                 *VsanClusterLimitHealthResult            `xml:"limitHealth,omitempty"`
	AdvCfgSync                  []VsanClusterAdvCfgSyncResult            `xml:"advCfgSync,omitempty"`
	CreateVmHealth              []VsanHostCreateVmHealthTestResult       `xml:"createVmHealth,omitempty"`
	PhysicalDisksHealth         []VsanPhysicalDiskHealthSummary          `xml:"physicalDisksHealth,omitempty"`
	EncryptionHealth            *VsanClusterEncryptionHealthSummary      `xml:"encryptionHealth,omitempty"`
	FileServiceHealth           *VsanClusterFileServiceHealthSummary     `xml:"fileServiceHealth,omitempty"`
	HclInfo                     *VsanClusterHclInfo                      `xml:"hclInfo,omitempty"`
	Groups                      []VsanClusterHealthGroup                 `xml:"groups,omitempty"`
	OverallHealth               string                                   `xml:"overallHealth"`
	OverallHealthDescription    string                                   `xml:"overallHealthDescription"`
	ClomdLiveness               *VsanClusterClomdLivenessResult          `xml:"clomdLiveness,omitempty"`
	DiskBalance                 *VsanClusterBalanceSummary               `xml:"diskBalance,omitempty"`
	GenericCluster              *VsanGenericClusterBestPracticeHealth    `xml:"genericCluster,omitempty"`
	NetworkConfig               *VsanNetworkConfigBestPracticeHealth     `xml:"networkConfig,omitempty"`
	VsanConfig                  *VsanConfigCheckResult                   `xml:"vsanConfig,omitempty"`
	BurnInTest                  *VsanBurnInTestCheckResult               `xml:"burnInTest,omitempty"`
	PerfsvcHealth               *VsanPerfsvcHealthResult                 `xml:"perfsvcHealth,omitempty"`
	Cluster                     *types.ManagedObjectReference            `xml:"cluster,omitempty"`
	DpdLiveness                 *VsanClusterDpdLivenessResult            `xml:"dpdLiveness,omitempty"`
	DatastoreUsage              *VsanClusterDatastoreUsageResult         `xml:"datastoreUsage,omitempty"`
	DataProtectionCfgSync       []VsanClusterDataProtectionCfgSyncResult `xml:"dataProtectionCfgSync,omitempty"`
	ArchivalAccessibilityStatus []VsanArchivalAccessibilityResult        `xml:"archivalAccessibilityStatus,omitempty"`
}

type VsanPerfEntityType struct {
	types.DynamicData

	Name        string          `xml:"name"`
	Id          string          `xml:"id"`
	Graphs      []VsanPerfGraph `xml:"graphs"`
	Description string          `xml:"description,omitempty"`
}

type VsanNetworkLoadTestResult struct {
	types.DynamicData

	Hostname      string  `xml:"hostname"`
	Status        string  `xml:"status,omitempty"`
	Client        bool    `xml:"client"`
	BandwidthBps  int64   `xml:"bandwidthBps"`
	TotalBytes    int64   `xml:"totalBytes"`
	LostDatagrams int64   `xml:"lostDatagrams,omitempty"`
	LossPct       int64   `xml:"lossPct,omitempty"`
	SentDatagrams int64   `xml:"sentDatagrams,omitempty"`
	JitterMs      float32 `xml:"jitterMs,omitempty"`
}

type VsanPhysicalDiskHealthSummary struct {
	types.DynamicData

	OverallHealth        string                   `xml:"overallHealth"`
	HeapsWithIssues      []VsanResourceHealth     `xml:"heapsWithIssues,omitempty"`
	SlabsWithIssues      []VsanResourceHealth     `xml:"slabsWithIssues,omitempty"`
	Disks                []VsanPhysicalDiskHealth `xml:"disks,omitempty"`
	ComponentsWithIssues []VsanResourceHealth     `xml:"componentsWithIssues,omitempty"`
	Hostname             string                   `xml:"hostname,omitempty"`
	HostDedupScope       int32                    `xml:"hostDedupScope,omitempty"`
	Error                *types.MethodFault       `xml:"error,omitempty"`
}

type VsanNetworkConfigVsanNotOnVdsIssue struct {
	VsanNetworkConfigBaseIssue

	Host   types.ManagedObjectReference `xml:"host"`
	Vmknic string                       `xml:"vmknic"`
}

type VsanNetworkConfigBaseIssue struct {
	types.DynamicData
}

type VsanDatastoreSpec struct {
	types.DynamicData

	Uuid string `xml:"uuid"`
	Name string `xml:"name"`
}

type VsanJsonFilterRule struct {
	types.DynamicData

	FilterComparator *VsanComparator `xml:"filterComparator,omitempty"`
	ComparablePath   []string        `xml:"comparablePath,omitempty"`
	KeysWithStrVal   []string        `xml:"keysWithStrVal,omitempty"`
	PropertyName     string          `xml:"propertyName,omitempty"`
}

type VsanWhatIfEvacResult struct {
	types.DynamicData

	NoAction     VsanWhatIfEvacDetail `xml:"noAction"`
	EnsureAccess VsanWhatIfEvacDetail `xml:"ensureAccess"`
	EvacAllData  VsanWhatIfEvacDetail `xml:"evacAllData"`
}

type VsanVcsaDeploymentSpecValidationResult struct {
	types.DynamicData

	Valid             bool                    `xml:"valid"`
	Errors            []types.MethodFault     `xml:"errors,omitempty"`
	InvalidProperties []string                `xml:"invalidProperties,omitempty"`
	DeploymentSpec    *VsanVcsaDeploymentSpec `xml:"deploymentSpec,omitempty"`
}

type VsanDataProtectionTargetSiteInfo struct {
	types.DynamicData

	TargetType string `xml:"targetType"`
}

type CnsNfsFileShareBackingDetails struct {
	CnsFileBackingDetails

	Address string `xml:"address,omitempty"`
}

type VimHostVsanRecoveryFlags struct {
	types.DynamicData

	IncludeAbsent   bool `xml:"includeAbsent"`
	IncludeDegraded bool `xml:"includeDegraded"`
	IncludeActive   bool `xml:"includeActive"`
	Force           bool `xml:"force"`
	RecoverAll      bool `xml:"recoverAll"`
}

type VsanConfigInfoEx struct {
	types.VsanClusterConfigInfo

	DataEfficiencyConfig  *VsanDataEfficiencyConfig     `xml:"dataEfficiencyConfig,omitempty"`
	ResyncIopsLimitConfig *ResyncIopsInfo               `xml:"resyncIopsLimitConfig,omitempty"`
	IscsiConfig           *VsanIscsiTargetServiceConfig `xml:"iscsiConfig,omitempty"`
	DataEncryptionConfig  *VsanDataEncryptionConfig     `xml:"dataEncryptionConfig,omitempty"`
	BmcConfig             []VimClusterVsanBmcSpec       `xml:"bmcConfig,omitempty"`
	ExtendedConfig        *VsanExtendedConfig           `xml:"extendedConfig,omitempty"`
	DataProtectionConfig  *VsanDataProtectionInfo       `xml:"dataProtectionConfig,omitempty"`
	DatastoreConfig       *VsanDatastoreConfig          `xml:"datastoreConfig,omitempty"`
	PerfsvcConfig         *VsanPerfsvcConfig            `xml:"perfsvcConfig,omitempty"`
	FileServiceConfig     *VsanFileServiceConfig        `xml:"fileServiceConfig,omitempty"`
	UnmapConfig           *VsanUnmapConfig              `xml:"unmapConfig,omitempty"`
	RdmaConfig            *VsanRdmaConfig               `xml:"rdmaConfig,omitempty"`
	VumConfig             *VsanVumConfig                `xml:"vumConfig,omitempty"`
	MetricsConfig         *VsanMetricsConfig            `xml:"metricsConfig,omitempty"`
}

type VsanClusterMetroConfig struct {
	types.DynamicData

	Enabled                bool                          `xml:"enabled"`
	WitnessNode            *types.ManagedObjectReference `xml:"witnessNode,omitempty"`
	WitnessNodeUuid        string                        `xml:"witnessNodeUuid,omitempty"`
	WitnessNodeIPAddresses string                        `xml:"witnessNodeIPAddresses,omitempty"`
	PreferredFdName        string                        `xml:"preferredFdName,omitempty"`
	PreferredFdUuid        string                        `xml:"preferredFdUuid,omitempty"`
	MetadataMode           *bool                         `xml:"metadataMode"`
	WitnessHosts           []VsanWitnessConfig           `xml:"witnessHosts,omitempty"`
}

type VsanHostResourceCheckResult struct {
	EntityResourceCheckDetails

	Host       *types.ManagedObjectReference      `xml:"host,omitempty"`
	DiskGroups []VsanDiskGroupResourceCheckResult `xml:"diskGroups,omitempty"`
}

type VsanClusterAdvCfgSyncHostResult struct {
	types.DynamicData

	Hostname  string `xml:"hostname"`
	Value     string `xml:"value"`
	IsDefault *bool  `xml:"isDefault"`
}

type VsanClusterDpdLivenessResult struct {
	types.DynamicData

	DpdLivenessResult []VsanHostDpdLivenessResult `xml:"dpdLivenessResult,omitempty"`
	IssueFound        bool                        `xml:"issueFound"`
}

type RepairTimerInfo struct {
	types.DynamicData

	MaxTimeToRepair            int32 `xml:"maxTimeToRepair"`
	MinTimeToRepair            int32 `xml:"minTimeToRepair"`
	ObjectCount                int32 `xml:"objectCount"`
	ObjectCountWithRepairTimer int32 `xml:"objectCountWithRepairTimer,omitempty"`
}

type VsanHostIoInsightInfo struct {
	types.DynamicData

	Host              types.ManagedObjectReference `xml:"host"`
	IoinsightWorldId  int64                        `xml:"ioinsightWorldId,omitempty"`
	FaultMessage      string                       `xml:"faultMessage,omitempty"`
	IoinsightInternal *VsanIoInsightInternal       `xml:"ioinsightInternal,omitempty"`
}

type VsanDataEfficiencyConfig struct {
	types.DynamicData

	DedupEnabled       bool  `xml:"dedupEnabled"`
	CompressionEnabled *bool `xml:"compressionEnabled"`
}

type VimHostVSANCmmdsPreferredFaultDomainInfo struct {
	types.DynamicData

	PreferredFaultDomainId   string `xml:"preferredFaultDomainId"`
	PreferredFaultDomainName string `xml:"preferredFaultDomainName"`
}

type VsanPerfTopEntity struct {
	types.DynamicData

	EntityRefId string `xml:"entityRefId"`
	Value       string `xml:"value"`
}

type VsanSyncingObjectRecoveryDetails struct {
	types.DynamicData

	ActivelySyncingObjectRecoveryETA int64 `xml:"activelySyncingObjectRecoveryETA,omitempty"`
	QueuedForSyncObjectRecoveryETA   int64 `xml:"queuedForSyncObjectRecoveryETA,omitempty"`
	SuspendedObjectRecoveryETA       int64 `xml:"suspendedObjectRecoveryETA,omitempty"`
	ActiveObjectsToSync              int64 `xml:"activeObjectsToSync,omitempty"`
	QueuedObjectsToSync              int64 `xml:"queuedObjectsToSync,omitempty"`
	SuspendedObjectsToSync           int64 `xml:"suspendedObjectsToSync,omitempty"`
	BytesToSyncForActiveObjects      int64 `xml:"bytesToSyncForActiveObjects,omitempty"`
	BytesToSyncForQueuedObjects      int64 `xml:"bytesToSyncForQueuedObjects,omitempty"`
	BytesToSyncForSuspendedObjects   int64 `xml:"bytesToSyncForSuspendedObjects,omitempty"`
}

type VsanHclFirmwareFile struct {
	types.DynamicData

	FileType      string `xml:"fileType"`
	FilenameOrUrl string `xml:"filenameOrUrl"`
	Sha1sum       string `xml:"sha1sum"`
}

type VsanPerfDiagnosticException struct {
	types.DynamicData

	ExceptionId      string `xml:"exceptionId"`
	ExceptionMessage string `xml:"exceptionMessage"`
	ExceptionDetails string `xml:"exceptionDetails"`
	ExceptionUrl     string `xml:"exceptionUrl"`
}

type VsanUnknownScanIssue struct {
	types.VsanUpgradeSystemPreflightCheckIssue

	Uuids []string `xml:"uuids"`
}

type VsanClusterHostVmknicMapping struct {
	types.DynamicData

	Host   string `xml:"host"`
	Vmknic string `xml:"vmknic"`
}

type VsanClusterFileServiceHealthSummary struct {
	types.DynamicData

	OverallHealth     string                                         `xml:"overallHealth,omitempty"`
	HostResults       []VsanFileServiceHealthSummary                 `xml:"hostResults,omitempty"`
	VcAdServiceHealth []VsanClusterFileServiceAdServiceHealthSummary `xml:"vcAdServiceHealth,omitempty"`
}

type VsanStorageWorkloadType struct {
	types.DynamicData

	Specs       []VsanVmdkLoadTestSpec `xml:"specs"`
	TypeId      string                 `xml:"typeId"`
	Name        string                 `xml:"name"`
	Description string                 `xml:"description"`
	Duration    int64                  `xml:"duration,omitempty"`
}

type DataProtectionLoadBalancerDetailedInfo struct {
	DataProtectionLoadBalancerBasicInfo

	Url        string `xml:"url"`
	Thumbprint string `xml:"thumbprint,omitempty"`
	PublicKey  string `xml:"publicKey"`
}

type VsanObjectPolicyIssue struct {
	types.VsanUpgradeSystemPreflightCheckIssue

	Uuids []string `xml:"uuids"`
}

type VimHostVsanRealignResult struct {
	types.DynamicData

	Status   *bool `xml:"status"`
	Progress int64 `xml:"progress,omitempty"`
	Total    int64 `xml:"total,omitempty"`
}

type VsanPerfQuerySpec struct {
	types.DynamicData

	EntityRefId string     `xml:"entityRefId"`
	StartTime   *time.Time `xml:"startTime"`
	EndTime     *time.Time `xml:"endTime"`
	Group       string     `xml:"group,omitempty"`
	Labels      []string   `xml:"labels,omitempty"`
	Interval    int32      `xml:"interval,omitempty"`
}

type VsanProactiveRebalanceInfoEx struct {
	types.DynamicData

	Running           *bool              `xml:"running"`
	StartTs           *time.Time         `xml:"startTs"`
	StopTs            *time.Time         `xml:"stopTs"`
	VarianceThreshold float32            `xml:"varianceThreshold,omitempty"`
	TimeThreshold     int32              `xml:"timeThreshold,omitempty"`
	RateThreshold     int32              `xml:"rateThreshold,omitempty"`
	Hostname          string             `xml:"hostname,omitempty"`
	Error             *types.MethodFault `xml:"error,omitempty"`
}

type VimVsanVsanDiskGroupComplianceResourceCheck struct {
	types.DynamicData

	Ssd             *VimVsanVsanDiskComplianceResourceCheck  `xml:"ssd,omitempty"`
	CapacityDevices []VimVsanVsanDiskComplianceResourceCheck `xml:"capacityDevices,omitempty"`
}

type VsanFailedRepairObjectResult struct {
	types.DynamicData

	Uuid       string `xml:"uuid"`
	ErrMessage string `xml:"errMessage,omitempty"`
}

type VsanClusterCreateVmHealthTestResult struct {
	types.DynamicData

	ClusterResult VsanClusterProactiveTestResult     `xml:"clusterResult"`
	HostResults   []VsanHostCreateVmHealthTestResult `xml:"hostResults,omitempty"`
}

type VimVsanVsanDiskComplianceResourceCheck struct {
	types.DynamicData

	DeviceName    string `xml:"deviceName,omitempty"`
	Uuid          string `xml:"uuid,omitempty"`
	Capacity      int64  `xml:"capacity,omitempty"`
	InitCapacity  int64  `xml:"initCapacity,omitempty"`
	FinalCapacity int64  `xml:"finalCapacity,omitempty"`
	IsNew         *bool  `xml:"isNew"`
}

type VsanClusterBalanceSummary struct {
	types.DynamicData

	VarianceThreshold int64                           `xml:"varianceThreshold"`
	Disks             []VsanClusterBalancePerDiskInfo `xml:"disks,omitempty"`
}

type CnsKubernetesEntityMetadata struct {
	CnsEntityMetadata

	EntityType     string   `xml:"entityType"`
	Namespace      string   `xml:"namespace,omitempty"`
	ReferredEntity []string `xml:"referredEntity,omitempty"`
}

type VsanHostAssociatedObjectsResult struct {
	types.DynamicData

	Data   []VsanHostAssociatedObjects `xml:"data"`
	Offset int32                       `xml:"offset"`
	Limit  int32                       `xml:"limit"`
}

type VsanConfigGeneration struct {
	types.DynamicData

	VcUuid  string `xml:"vcUuid"`
	GenNum  int64  `xml:"genNum"`
	GenTime int64  `xml:"genTime"`
}

type VsanVdsPgMigrationHostInfo struct {
	types.DynamicData

	Host          types.ManagedObjectReference `xml:"host"`
	Hostname      string                       `xml:"hostname"`
	VmknicDevices []string                     `xml:"vmknicDevices,omitempty"`
	VmVnics       []VsanVdsPgMigrationVmInfo   `xml:"vmVnics,omitempty"`
}

type VsanComplianceResourceCheckStatus struct {
	types.DynamicData

	Status string                                    `xml:"status,omitempty"`
	Result *VimVsanVsanComplianceResourceCheckReport `xml:"result,omitempty"`
}

type VsanVumSystemClusterConfig struct {
	types.DynamicData

	Cluster                types.ManagedObjectReference `xml:"cluster"`
	BaselinePreferenceType string                       `xml:"baselinePreferenceType,typeattr"`
}

type VsanIscsiTargetAuthSpec struct {
	types.DynamicData

	AuthType                    string `xml:"authType,omitempty"`
	UserNameAttachToTarget      string `xml:"userNameAttachToTarget,omitempty"`
	UserSecretAttachToTarget    string `xml:"userSecretAttachToTarget,omitempty"`
	UserNameAttachToInitiator   string `xml:"userNameAttachToInitiator,omitempty"`
	UserSecretAttachToInitiator string `xml:"userSecretAttachToInitiator,omitempty"`
}

type VsanClusterNetworkLoadTestResult struct {
	types.DynamicData

	ClusterResult VsanClusterProactiveTestResult `xml:"clusterResult"`
	HostResults   []VsanNetworkLoadTestResult    `xml:"hostResults,omitempty"`
}

type VsanHostReference struct {
	types.DynamicData

	Hostname string `xml:"hostname"`
}

type VsanNetworkConfigBestPracticeHealth struct {
	types.DynamicData

	VdsPresent bool                         `xml:"vdsPresent"`
	Issues     []VsanNetworkConfigBaseIssue `xml:"issues,omitempty"`
}

type VimClusterVsanBmcSpec struct {
	types.DynamicData

	Host        types.ManagedObjectReference `xml:"host"`
	BmcAddress  string                       `xml:"bmcAddress"`
	BmcUserName string                       `xml:"bmcUserName"`
	BmcPassword string                       `xml:"bmcPassword,omitempty"`
}

type VimHostVSANCmmdsNodeInfo struct {
	types.DynamicData

	NodeUuid  string `xml:"nodeUuid"`
	IsWitness bool   `xml:"isWitness"`
}

type VsanHostAboutInfoEx struct {
	types.DynamicData

	Name       string `xml:"name,omitempty"`
	Version    string `xml:"version,omitempty"`
	Build      string `xml:"build,omitempty"`
	BuildType  string `xml:"buildType,omitempty"`
	ApiVersion string `xml:"apiVersion,omitempty"`
}

type VsanClusterLimitHealthResult struct {
	types.DynamicData

	IssueFound              bool                                  `xml:"issueFound"`
	ComponentLimitHealth    string                                `xml:"componentLimitHealth"`
	DiskFreeSpaceHealth     string                                `xml:"diskFreeSpaceHealth"`
	RcFreeReservationHealth string                                `xml:"rcFreeReservationHealth"`
	HostResults             []VsanLimitHealthResult               `xml:"hostResults,omitempty"`
	WhatifHostFailures      []VsanClusterWhatifHostFailuresResult `xml:"whatifHostFailures,omitempty"`
	HostsCommFailure        []string                              `xml:"hostsCommFailure,omitempty"`
}

type VsanStorageComplianceResult struct {
	types.DynamicData

	CheckTime             *time.Time                    `xml:"checkTime"`
	Profile               string                        `xml:"profile,omitempty"`
	ObjectUUID            string                        `xml:"objectUUID,omitempty"`
	ComplianceStatus      string                        `xml:"complianceStatus"`
	Mismatch              bool                          `xml:"mismatch"`
	ViolatedPolicies      []VsanStoragePolicyStatus     `xml:"violatedPolicies,omitempty"`
	OperationalStatus     *VsanStorageOperationalStatus `xml:"operationalStatus,omitempty"`
	ObjPolicyGenerationId string                        `xml:"objPolicyGenerationId,omitempty"`
}

type VsanFileShareNetPermission struct {
	types.DynamicData

	Ips         string `xml:"ips"`
	Permissions string `xml:"permissions,omitempty"`
	AllowRoot   *bool  `xml:"allowRoot"`
}

type CnsSearchLabelResult struct {
	types.DynamicData

	Labels         []types.KeyValue `xml:"labels,omitempty"`
	HasMoreResults bool             `xml:"hasMoreResults"`
}

type VsanIscsiTargetCommonInfo struct {
	VsanIscsiTargetBasicInfo

	AuthSpec         *VsanIscsiTargetAuthSpec `xml:"authSpec,omitempty"`
	Port             int32                    `xml:"port,omitempty"`
	NetworkInterface string                   `xml:"networkInterface,omitempty"`
}

type VsanSmartParameter struct {
	types.DynamicData

	Parameter string `xml:"parameter,omitempty"`
	Value     int32  `xml:"value,omitempty"`
	Threshold int32  `xml:"threshold,omitempty"`
	Worst     int32  `xml:"worst,omitempty"`
}

type VsanCompliantFirmware struct {
	types.DynamicData

	FirmwareVersion  string                `xml:"firmwareVersion"`
	CompliantDrivers []VsanCompliantDriver `xml:"compliantDrivers"`
}

type VsanClusterNetworkPartitionInfo struct {
	types.DynamicData

	Hosts            []string `xml:"hosts,omitempty"`
	PartitionUnknown *bool    `xml:"partitionUnknown"`
}

type VsanClusterHclInfo struct {
	types.DynamicData

	HclDbLastUpdate *time.Time        `xml:"hclDbLastUpdate"`
	HclDbAgeHealth  string            `xml:"hclDbAgeHealth,omitempty"`
	HostResults     []VsanHostHclInfo `xml:"hostResults,omitempty"`
	UpdateItems     []VsanUpdateItem  `xml:"updateItems,omitempty"`
}

type VsanWitnessConfig struct {
	types.DynamicData

	Host                 types.ManagedObjectReference `xml:"host"`
	Uuid                 string                       `xml:"uuid,omitempty"`
	PreferredFaultDomain string                       `xml:"preferredFaultDomain,omitempty"`
	PreferredFdUuid      string                       `xml:"preferredFdUuid,omitempty"`
	NetworkAddresses     []string                     `xml:"networkAddresses,omitempty"`
}

type VsanClusterHealthResultBase struct {
	types.DynamicData

	Label string `xml:"label,omitempty"`
}

type CnsCursor struct {
	types.DynamicData

	Offset       int64 `xml:"offset"`
	Limit        int64 `xml:"limit"`
	TotalRecords int64 `xml:"totalRecords,omitempty"`
}

type VsanIscsiLUN struct {
	VsanIscsiLUNCommonInfo

	TargetAlias       string                 `xml:"targetAlias"`
	Uuid              string                 `xml:"uuid"`
	ActualSize        int64                  `xml:"actualSize"`
	ObjectInformation *VsanObjectInformation `xml:"objectInformation,omitempty"`
}

type VsanPerfDiagnosticResult struct {
	types.DynamicData

	ExceptionId         string                    `xml:"exceptionId"`
	Recommendation      string                    `xml:"recommendation,omitempty"`
	AggregationFunction string                    `xml:"aggregationFunction,omitempty"`
	AggregationData     *VsanPerfEntityMetricCSV  `xml:"aggregationData,omitempty"`
	ExceptionData       []VsanPerfEntityMetricCSV `xml:"exceptionData"`
}

type VsanNetworkConfigVdsScopeIssue struct {
	VsanNetworkConfigBaseIssue

	Vds            types.ManagedObjectReference   `xml:"vds"`
	MemberHosts    []types.ManagedObjectReference `xml:"memberHosts"`
	NonMemberHosts []types.ManagedObjectReference `xml:"nonMemberHosts"`
}

type VsanHclDriverInfo struct {
	types.DynamicData

	DriverName    string             `xml:"driverName,omitempty"`
	DriverVersion string             `xml:"driverVersion,omitempty"`
	DriverLink    *VsanDownloadItem  `xml:"driverLink,omitempty"`
	FwVersion     string             `xml:"fwVersion,omitempty"`
	FwLinks       []VsanDownloadItem `xml:"fwLinks,omitempty"`
	ToolsLinks    []VsanDownloadItem `xml:"toolsLinks,omitempty"`
	Eula          string             `xml:"eula,omitempty"`
	DriverType    string             `xml:"driverType,omitempty"`
}

type CnsVolumeAttachResult struct {
	CnsVolumeOperationResult

	DiskUUID string `xml:"diskUUID,omitempty"`
}

type VimClusterVsanFaultDomainsConfigSpec struct {
	types.DynamicData

	FaultDomains []VimClusterVsanFaultDomainSpec `xml:"faultDomains"`
	Witness      *VimClusterVsanWitnessSpec      `xml:"witness,omitempty"`
}

type VsanFileServiceOvfSpec struct {
	types.DynamicData

	Version    string                        `xml:"version,omitempty"`
	UpdateTime *time.Time                    `xml:"updateTime"`
	Task       *types.ManagedObjectReference `xml:"task,omitempty"`
}

type VsanMixedEsxVersionIssue struct {
	types.VsanUpgradeSystemPreflightCheckIssue
}

type VsanMassCollectorSpec struct {
	types.DynamicData

	Objects          []types.ManagedObjectReference    `xml:"objects,omitempty"`
	ObjectCollection string                            `xml:"objectCollection,omitempty"`
	Properties       []string                          `xml:"properties"`
	PropertiesParams []VsanMassCollectorPropertyParams `xml:"propertiesParams,omitempty"`
	Constraint       *VsanResourceConstraint           `xml:"constraint,omitempty"`
}

type VsanHclNicInfo struct {
	VsanHclCommonDeviceInfo

	RdmaCapable *bool `xml:"rdmaCapable"`
}

type VsanConfigNotAllDisksClaimedIssue struct {
	VsanConfigBaseIssue

	Host  types.ManagedObjectReference `xml:"host"`
	Disks []string                     `xml:"disks"`
}

type VsanClusterHealthTest struct {
	types.DynamicData

	TestId               string                        `xml:"testId,omitempty"`
	TestName             string                        `xml:"testName,omitempty"`
	TestDescription      string                        `xml:"testDescription,omitempty"`
	TestShortDescription string                        `xml:"testShortDescription,omitempty"`
	TestHealthyEntities  int32                         `xml:"testHealthyEntities,omitempty"`
	TestAllEntities      int32                         `xml:"testAllEntities,omitempty"`
	TestHealth           string                        `xml:"testHealth,omitempty"`
	TestDetails          []VsanClusterHealthResultBase `xml:"testDetails,omitempty"`
	TestActions          []VsanClusterHealthAction     `xml:"testActions,omitempty"`
}

type VsanHostEncryptionInfo struct {
	types.DynamicData

	Enabled             *bool                  `xml:"enabled"`
	KekId               string                 `xml:"kekId,omitempty"`
	HostKeyId           string                 `xml:"hostKeyId,omitempty"`
	KmipServers         []types.KmipServerSpec `xml:"kmipServers,omitempty"`
	KmsServerCerts      []string               `xml:"kmsServerCerts,omitempty"`
	ClientKey           string                 `xml:"clientKey,omitempty"`
	ClientCert          string                 `xml:"clientCert,omitempty"`
	DekGenerationId     int64                  `xml:"dekGenerationId,omitempty"`
	Changing            *bool                  `xml:"changing"`
	EraseDisksBeforeUse *bool                  `xml:"eraseDisksBeforeUse"`
}

type VsanHealthExtMgmtPreCheckResult struct {
	types.DynamicData

	OverallResult            bool                    `xml:"overallResult"`
	EsxVersionCheckPassed    *bool                   `xml:"esxVersionCheckPassed"`
	DrsCheckPassed           *bool                   `xml:"drsCheckPassed"`
	EamConnectionCheckPassed *bool                   `xml:"eamConnectionCheckPassed"`
	InstallStateCheckPassed  *bool                   `xml:"installStateCheckPassed"`
	Results                  []VsanClusterHealthTest `xml:"results"`
	VumRegistered            *bool                   `xml:"vumRegistered"`
}

type VsanClusterClomdLivenessResult struct {
	types.DynamicData

	ClomdLivenessResult []VsanHostClomdLivenessResult `xml:"clomdLivenessResult,omitempty"`
	IssueFound          bool                          `xml:"issueFound"`
}

type VsanDataProtectionSpaceUsageStats struct {
	types.DynamicData

	OverheadB              int64 `xml:"overheadB"`
	FragmentationOverheadB int64 `xml:"fragmentationOverheadB"`
	RaidPolicyOverheadB    int64 `xml:"raidPolicyOverheadB"`
}

type VsanDataProtectionArchivalLocation struct {
	types.DynamicData

	DatastoreUrl string `xml:"datastoreUrl"`
}

type VimClusterVsanHostDiskMapping struct {
	types.DynamicData

	Host          types.ManagedObjectReference `xml:"host"`
	CacheDisks    []types.HostScsiDisk         `xml:"cacheDisks,omitempty"`
	CapacityDisks []types.HostScsiDisk         `xml:"capacityDisks"`
	Type          string                       `xml:"type"`
}

type VsanPhysicalDiskHealth struct {
	types.DynamicData

	Name                         string                   `xml:"name"`
	Uuid                         string                   `xml:"uuid"`
	InCmmds                      bool                     `xml:"inCmmds"`
	InVsi                        bool                     `xml:"inVsi"`
	DedupScope                   int64                    `xml:"dedupScope,omitempty"`
	FormatVersion                int32                    `xml:"formatVersion,omitempty"`
	IsAllFlash                   int32                    `xml:"isAllFlash,omitempty"`
	CongestionValue              int32                    `xml:"congestionValue,omitempty"`
	CongestionArea               string                   `xml:"congestionArea,omitempty"`
	CongestionHealth             string                   `xml:"congestionHealth,omitempty"`
	MetadataHealth               string                   `xml:"metadataHealth,omitempty"`
	OperationalHealthDescription string                   `xml:"operationalHealthDescription,omitempty"`
	OperationalHealth            string                   `xml:"operationalHealth,omitempty"`
	DedupUsageHealth             string                   `xml:"dedupUsageHealth,omitempty"`
	CapacityHealth               string                   `xml:"capacityHealth,omitempty"`
	SummaryHealth                string                   `xml:"summaryHealth"`
	Capacity                     int64                    `xml:"capacity,omitempty"`
	UsedCapacity                 int64                    `xml:"usedCapacity,omitempty"`
	ReservedCapacity             int64                    `xml:"reservedCapacity,omitempty"`
	TotalBytes                   int64                    `xml:"totalBytes,omitempty"`
	FreeBytes                    int64                    `xml:"freeBytes,omitempty"`
	HashedBytes                  int64                    `xml:"hashedBytes,omitempty"`
	DedupedBytes                 int64                    `xml:"dedupedBytes,omitempty"`
	ScsiDisk                     *types.HostScsiDisk      `xml:"scsiDisk,omitempty"`
	UsedComponents               int64                    `xml:"usedComponents,omitempty"`
	MaxComponents                int64                    `xml:"maxComponents,omitempty"`
	CompLimitHealth              string                   `xml:"compLimitHealth,omitempty"`
	EncryptionEnabled            *bool                    `xml:"encryptionEnabled"`
	KmsProviderId                string                   `xml:"kmsProviderId,omitempty"`
	KekId                        string                   `xml:"kekId,omitempty"`
	DekGenerationId              int64                    `xml:"dekGenerationId,omitempty"`
	EncryptedUnlocked            *bool                    `xml:"encryptedUnlocked"`
	RebalanceResult              *VsanDiskRebalanceResult `xml:"rebalanceResult,omitempty"`
}

type VsanPerfMetricSeriesCSV struct {
	types.DynamicData

	MetricId      VsanPerfMetricId   `xml:"metricId"`
	Threshold     *VsanPerfThreshold `xml:"threshold,omitempty"`
	NumExceptions string             `xml:"numExceptions,omitempty"`
	Values        string             `xml:"values,omitempty"`
}

type VsanHclReleaseConstraint struct {
	types.DynamicData

	Cluster     types.ManagedObjectReference `xml:"cluster"`
	Release     string                       `xml:"release"`
	Constraints []VsanHclDeviceConstraint    `xml:"constraints,omitempty"`
}

type VsanRepairObjectsResult struct {
	types.DynamicData

	InQueueObjects      []string                       `xml:"inQueueObjects,omitempty"`
	FailedRepairObjects []VsanFailedRepairObjectResult `xml:"failedRepairObjects,omitempty"`
	NotInQueueObjects   []string                       `xml:"notInQueueObjects,omitempty"`
}

type VsanDataProtectionRemoteSpaceUsage struct {
	types.DynamicData

	ActiveDataProtectionsUsage   []VsanDataProtectionSpaceUsage `xml:"activeDataProtectionsUsage"`
	InactiveDataProtectionsUsage []VsanDataProtectionSpaceUsage `xml:"inactiveDataProtectionsUsage"`
}

type VsanClusterVMsHealthSummaryResult struct {
	types.DynamicData

	NumVMs          int32    `xml:"numVMs"`
	State           string   `xml:"state,omitempty"`
	Health          string   `xml:"health"`
	VmInstanceUuids []string `xml:"vmInstanceUuids,omitempty"`
}

type VsanVmVdsMigrationSpec struct {
	types.DynamicData

	VmInstanceUuid string                     `xml:"vmInstanceUuid"`
	Vnics          []VsanVnicVdsMigrationSpec `xml:"vnics"`
}

type VsanHostDatastoreUsageResult struct {
	types.DynamicData

	Hostname                       string             `xml:"hostname"`
	ConfiguredUsageThresholdOnHost int32              `xml:"configuredUsageThresholdOnHost,omitempty"`
	CurrentUsage                   float32            `xml:"currentUsage,omitempty"`
	ThresholdExceeded              *bool              `xml:"thresholdExceeded"`
	Error                          *types.MethodFault `xml:"error,omitempty"`
}

type VsanPerfMasterInformation struct {
	types.DynamicData

	SecSinceLastStatsWrite     int64      `xml:"secSinceLastStatsWrite,omitempty"`
	SecSinceLastStatsCollect   int64      `xml:"secSinceLastStatsCollect,omitempty"`
	StatsIntervalSec           int64      `xml:"statsIntervalSec"`
	CollectionFailureHostUuids []string   `xml:"collectionFailureHostUuids,omitempty"`
	RenamedStatsDirectories    []string   `xml:"renamedStatsDirectories,omitempty"`
	StatsDirectoryPercentFree  int64      `xml:"statsDirectoryPercentFree,omitempty"`
	VerboseMode                *bool      `xml:"verboseMode"`
	VerboseModeLastUpdate      *time.Time `xml:"verboseModeLastUpdate"`
}

type VimVsanVsanComplianceResourceCheckReport struct {
	types.DynamicData

	IsValid      *bool                                           `xml:"isValid"`
	FaultDomains []VimVsanVsanFaultDomainComplianceResourceCheck `xml:"faultDomains,omitempty"`
}

type VsanClusterVmdkLoadTestResult struct {
	types.DynamicData

	Task          *types.ManagedObjectReference   `xml:"task,omitempty"`
	ClusterResult *VsanClusterProactiveTestResult `xml:"clusterResult,omitempty"`
	HostResults   []VsanHostVmdkLoadTestResult    `xml:"hostResults,omitempty"`
}

type VsanHostCreateVmHealthTestResult struct {
	types.DynamicData

	Hostname string             `xml:"hostname"`
	State    string             `xml:"state"`
	Fault    *types.MethodFault `xml:"fault,omitempty"`
}

type VsanDownloadItem struct {
	types.DynamicData

	Url        string `xml:"url"`
	Sha1sum    string `xml:"sha1sum"`
	FormatType string `xml:"formatType,omitempty"`
	ItemId     string `xml:"itemId,omitempty"`
}

type VsanVcsaDeploymentProgress struct {
	types.DynamicData

	Phase         string                        `xml:"phase"`
	ProgressPct   int64                         `xml:"progressPct"`
	Message       string                        `xml:"message"`
	Success       bool                          `xml:"success"`
	Error         *types.MethodFault            `xml:"error,omitempty"`
	UpdateCounter int64                         `xml:"updateCounter"`
	TaskId        string                        `xml:"taskId,omitempty"`
	Vm            *types.ManagedObjectReference `xml:"vm,omitempty"`
}

type VsanClusterFileServiceAdServiceHealthSummary struct {
	types.DynamicData

	DomainName       string   `xml:"domainName,omitempty"`
	DnsAddress       []string `xml:"dnsAddress,omitempty"`
	AdServiceAddress string   `xml:"adServiceAddress,omitempty"`
	Health           string   `xml:"health,omitempty"`
	Description      string   `xml:"description,omitempty"`
}

type VsanDataProtectionInfo struct {
	types.DynamicData

	ArchivalTarget          *VsanDataProtectionArchivalLocation `xml:"archivalTarget,omitempty"`
	UsageThreshold          int32                               `xml:"usageThreshold,omitempty"`
	PairingInfo             []VsanDataProtectionPairingInfo     `xml:"pairingInfo,omitempty"`
	IncomingReplicationPort int32                               `xml:"incomingReplicationPort,omitempty"`
}

type VsanFileServerHealthSummary struct {
	types.DynamicData

	DomainName    string `xml:"domainName,omitempty"`
	FileServerIp  string `xml:"fileServerIp,omitempty"`
	NfsdHealth    string `xml:"nfsdHealth,omitempty"`
	NetworkHealth string `xml:"networkHealth,omitempty"`
	RootfsHealth  string `xml:"rootfsHealth,omitempty"`
	Description   string `xml:"description,omitempty"`
}

type CnsVolumeCreateResult struct {
	CnsVolumeOperationResult

	Name string `xml:"name,omitempty"`
}

type CnsVolumeAttachDetachSpec struct {
	types.DynamicData

	VolumeId CnsVolumeId                  `xml:"volumeId"`
	Vm       types.ManagedObjectReference `xml:"vm"`
}

type VimClusterVsanDiskMappingsConfigSpec struct {
	types.DynamicData

	HostDiskMappings []VimClusterVsanHostDiskMapping `xml:"hostDiskMappings"`
}

type VsanJsonComparator struct {
	VsanComparator

	Comparator      string             `xml:"comparator,omitempty"`
	ComparableValue *types.KeyAnyValue `xml:"comparableValue,omitempty"`
}

type VimVsanVsanHostComplianceResourceCheck struct {
	types.DynamicData

	Host       *types.ManagedObjectReference                 `xml:"host,omitempty"`
	Uuid       string                                        `xml:"uuid,omitempty"`
	IsNew      *bool                                         `xml:"isNew"`
	DiskGroups []VimVsanVsanDiskGroupComplianceResourceCheck `xml:"diskGroups,omitempty"`
}

type VsanHostEMMSummary struct {
	types.DynamicData

	Hostname          string `xml:"hostname,omitempty"`
	InMaintenanceMode *bool  `xml:"inMaintenanceMode"`
	InDecomState      *bool  `xml:"inDecomState"`
}

type VsanHostVmdkLoadTestResult struct {
	types.DynamicData

	Hostname     string                   `xml:"hostname"`
	IssueFound   bool                     `xml:"issueFound"`
	FaultMessage string                   `xml:"faultMessage,omitempty"`
	VmdkResults  []VsanVmdkLoadTestResult `xml:"vmdkResults,omitempty"`
}

type VsanLocalProtectionTargetSiteInfo struct {
	VsanDataProtectionTargetSiteInfo
}

type VsanCompliantDriver struct {
	types.DynamicData

	DriverName     string   `xml:"driverName"`
	DriverVersions []string `xml:"driverVersions"`
}

type VsanDataProtectionArchivalTargetSiteInfo struct {
	VsanDataProtectionTargetSiteInfo

	DatastoreUrl string `xml:"datastoreUrl"`
}

type VsanClusterObjectExtAttrs struct {
	types.DynamicData

	Uuid          string `xml:"uuid"`
	ObjectType    string `xml:"objectType,omitempty"`
	ObjectPath    string `xml:"objectPath,omitempty"`
	GroupUuid     string `xml:"groupUuid,omitempty"`
	DirectoryName string `xml:"directoryName,omitempty"`
}

type VsanFileServiceConfig struct {
	types.DynamicData

	Enabled bool                          `xml:"enabled"`
	Network *types.ManagedObjectReference `xml:"network,omitempty"`
	Domains []VsanFileServiceDomainConfig `xml:"domains,omitempty"`
}

type VsanCloudHealthStatus struct {
	types.DynamicData

	CollectorRunning     *bool  `xml:"collectorRunning"`
	LastSentTimestamp    string `xml:"lastSentTimestamp,omitempty"`
	InternetConnectivity *bool  `xml:"internetConnectivity"`
}

type VsanClusterHealthSystemObjectsRepairResult struct {
	types.DynamicData

	InRepairingQueueObjects []string                       `xml:"inRepairingQueueObjects,omitempty"`
	FailedRepairObjects     []VsanFailedRepairObjectResult `xml:"failedRepairObjects,omitempty"`
	IssueFound              bool                           `xml:"issueFound"`
}

type VsanNetworkPeerHealthResult struct {
	types.DynamicData

	Peer                    string `xml:"peer,omitempty"`
	PeerHostname            string `xml:"peerHostname,omitempty"`
	PeerVmknicName          string `xml:"peerVmknicName,omitempty"`
	SmallPingTestSuccessPct int32  `xml:"smallPingTestSuccessPct,omitempty"`
	LargePingTestSuccessPct int32  `xml:"largePingTestSuccessPct,omitempty"`
	MaxLatencyUs            int64  `xml:"maxLatencyUs,omitempty"`
	OnSameIpSubnet          *bool  `xml:"onSameIpSubnet"`
	SourceVmknicName        string `xml:"sourceVmknicName,omitempty"`
}

type VsanIscsiTargetBasicInfo struct {
	types.DynamicData

	Alias string `xml:"alias"`
	Iqn   string `xml:"iqn,omitempty"`
}

type VsanNetworkConfigVswitchWithNoRedundancyIssue struct {
	VsanNetworkConfigBaseIssue

	Host        types.ManagedObjectReference  `xml:"host"`
	VswitchName string                        `xml:"vswitchName,omitempty"`
	Vds         *types.ManagedObjectReference `xml:"vds,omitempty"`
	NumPnics    int64                         `xml:"numPnics"`
}

type VimClusterVsanWitnessSpec struct {
	types.DynamicData

	Host                     types.ManagedObjectReference `xml:"host"`
	PreferredFaultDomainName string                       `xml:"preferredFaultDomainName"`
	DiskMapping              *types.VsanHostDiskMapping   `xml:"diskMapping,omitempty"`
}

type VsanObjectIdentity struct {
	types.DynamicData

	Uuid              string                         `xml:"uuid"`
	Type              string                         `xml:"type"`
	TypeExtId         string                         `xml:"typeExtId,omitempty"`
	VmInstanceUuid    string                         `xml:"vmInstanceUuid,omitempty"`
	VmNsObjectUuid    string                         `xml:"vmNsObjectUuid,omitempty"`
	Vm                *types.ManagedObjectReference  `xml:"vm,omitempty"`
	Description       string                         `xml:"description,omitempty"`
	DpSpaceUsageInfos []VsanDataProtectionSpaceUsage `xml:"dpSpaceUsageInfos,omitempty"`
	SpbmProfileUuid   string                         `xml:"spbmProfileUuid,omitempty"`
	Metadatas         []types.KeyValue               `xml:"metadatas,omitempty"`
}

type VsanVumSystemConfig struct {
	types.DynamicData

	Enabled                *bool                        `xml:"enabled"`
	AutoCheckInterval      int32                        `xml:"autoCheckInterval,omitempty"`
	MetadataUpdateInterval int32                        `xml:"metadataUpdateInterval,omitempty"`
	ClusterConfig          []VsanVumSystemClusterConfig `xml:"clusterConfig,omitempty"`
	ReleaseDbLastUpdate    *time.Time                   `xml:"releaseDbLastUpdate"`
}

type VsanConfigBaseIssue struct {
	types.DynamicData
}

type VsanObjectProfileInfo struct {
	types.DynamicData

	VsanObjectUuid           string `xml:"vsanObjectUuid"`
	SpbmProfileId            string `xml:"spbmProfileId"`
	SpbmProfileGenerationNum int32  `xml:"spbmProfileGenerationNum"`
}

type VsanClusterConfig struct {
	types.DynamicData

	Config      types.VsanClusterConfigInfo `xml:"config"`
	Name        string                      `xml:"name"`
	Hosts       []string                    `xml:"hosts,omitempty"`
	ToBeDeleted *bool                       `xml:"toBeDeleted"`
}

type VsanFlrSessionSpec struct {
	types.DynamicData

	MountSpec  []VsanFlrMountSpec     `xml:"mountSpec"`
	SearchSpec *VsanFlrSearchFileSpec `xml:"searchSpec,omitempty"`
}

type CnsVolumeOperationResult struct {
	types.DynamicData

	VolumeId *CnsVolumeId       `xml:"volumeId,omitempty"`
	Fault    *types.MethodFault `xml:"fault,omitempty"`
}

type VsanSmartStatsHostSummary struct {
	types.DynamicData

	Hostname   string               `xml:"hostname,omitempty"`
	SmartStats []VsanSmartDiskStats `xml:"smartStats,omitempty"`
}

type VimHostVSANStretchedClusterHostInfo struct {
	types.DynamicData

	NodeInfo                 VimHostVSANCmmdsNodeInfo                  `xml:"nodeInfo"`
	FaultDomainInfo          *VimHostVSANCmmdsFaultDomainInfo          `xml:"faultDomainInfo,omitempty"`
	PreferredFaultDomainInfo *VimHostVSANCmmdsPreferredFaultDomainInfo `xml:"preferredFaultDomainInfo,omitempty"`
}

type CnsSearchLabelSpec struct {
	types.DynamicData

	KeyPrefix          string `xml:"keyPrefix,omitempty"`
	ValuePrefix        string `xml:"valuePrefix,omitempty"`
	MaxNumberOfResults int64  `xml:"maxNumberOfResults"`
}

type VsanHostComponentSyncState struct {
	types.DynamicData

	Uuid        string   `xml:"uuid"`
	DiskUuid    string   `xml:"diskUuid"`
	HostUuid    string   `xml:"hostUuid"`
	BytesToSync int64    `xml:"bytesToSync"`
	RecoveryETA int64    `xml:"recoveryETA,omitempty"`
	Reasons     []string `xml:"reasons,omitempty"`
}

type VsanClusterBurnInTestResultList struct {
	types.DynamicData

	Items []VsanBurnInTest `xml:"items,omitempty"`
	Hosts []string         `xml:"hosts,omitempty"`
}

type VsanVsanClusterPcapGroup struct {
	types.DynamicData

	Master  string   `xml:"master"`
	Members []string `xml:"members,omitempty"`
}

type VsanFileShareQuerySpec struct {
	types.DynamicData

	DomainName string   `xml:"domainName,omitempty"`
	Uuids      []string `xml:"uuids,omitempty"`
	Names      []string `xml:"names,omitempty"`
}

type VsanMassCollectorPropertyParams struct {
	types.DynamicData

	PropertyName   string              `xml:"propertyName,omitempty"`
	PropertyParams []types.KeyAnyValue `xml:"propertyParams,omitempty"`
}

type VsanHostVirtualApplianceInfo struct {
	types.DynamicData

	HostKey      types.ManagedObjectReference `xml:"hostKey"`
	IsVirtualApp bool                         `xml:"isVirtualApp"`
}

type VsanPerfNodeInformation struct {
	types.DynamicData

	Version        string                     `xml:"version"`
	Hostname       string                     `xml:"hostname,omitempty"`
	Error          *types.MethodFault         `xml:"error,omitempty"`
	IsCmmdsMaster  bool                       `xml:"isCmmdsMaster"`
	IsStatsMaster  bool                       `xml:"isStatsMaster"`
	VsanMasterUuid string                     `xml:"vsanMasterUuid,omitempty"`
	VsanNodeUuid   string                     `xml:"vsanNodeUuid,omitempty"`
	MasterInfo     *VsanPerfMasterInformation `xml:"masterInfo,omitempty"`
	DiagnosticMode *bool                      `xml:"diagnosticMode"`
}

type VsanPerfEntityMetricCSV struct {
	types.DynamicData

	EntityRefId string                    `xml:"entityRefId"`
	SampleInfo  string                    `xml:"sampleInfo,omitempty"`
	Value       []VsanPerfMetricSeriesCSV `xml:"value,omitempty"`
}

type CnsFileCreateSpec struct {
	CnsBaseCreateSpec
}

type CnsVolume struct {
	types.DynamicData

	VolumeId                     CnsVolumeId              `xml:"volumeId"`
	DatastoreUrl                 string                   `xml:"datastoreUrl,omitempty"`
	Name                         string                   `xml:"name,omitempty"`
	VolumeType                   string                   `xml:"volumeType,omitempty"`
	StoragePolicyId              string                   `xml:"storagePolicyId,omitempty"`
	Metadata                     *CnsVolumeMetadata       `xml:"metadata,omitempty"`
	BackingObjectDetails         *CnsBackingObjectDetails `xml:"backingObjectDetails,omitempty"`
	ComplianceStatus             string                   `xml:"complianceStatus,omitempty"`
	DatastoreAccessibilityStatus string                   `xml:"datastoreAccessibilityStatus,omitempty"`
}

type CnsContainerCluster struct {
	types.DynamicData

	ClusterType string `xml:"clusterType"`
	ClusterId   string `xml:"clusterId"`
	VSphereUser string `xml:"vSphereUser"`
}

type VsanPerfTopEntities struct {
	types.DynamicData

	MetricId VsanPerfMetricId    `xml:"metricId"`
	Entities []VsanPerfTopEntity `xml:"entities"`
}

type VsanHostWithHybridDiskgroupIssue struct {
	types.VsanUpgradeSystemPreflightCheckIssue

	Hosts []types.ManagedObjectReference `xml:"hosts"`
}

type VimClusterVSANStretchedClusterCapability struct {
	types.DynamicData

	HostMoId       string                                     `xml:"hostMoId"`
	ConnStatus     string                                     `xml:"connStatus,omitempty"`
	IsSupported    *bool                                      `xml:"isSupported"`
	HostCapability *VimHostVSANStretchedClusterHostCapability `xml:"hostCapability,omitempty"`
}

type VsanNetworkConfigPnicSpeedInconsistencyIssue struct {
	VsanNetworkConfigBaseIssue

	Host        types.ManagedObjectReference  `xml:"host"`
	VswitchName string                        `xml:"vswitchName,omitempty"`
	Vds         *types.ManagedObjectReference `xml:"vds,omitempty"`
	SpeedsMb    []int64                       `xml:"speedsMb"`
}

type VsanInternalExtendedConfig struct {
	types.DynamicData

	VcMaxDiskVersion int32 `xml:"vcMaxDiskVersion,omitempty"`
}

type VsanIscsiLUNSpec struct {
	VsanIscsiLUNCommonInfo

	StoragePolicy *types.VirtualMachineProfileSpec `xml:"storagePolicy,omitempty"`
	NewLunId      int32                            `xml:"newLunId,omitempty"`
}

type VsanUpdateItem struct {
	types.DynamicData

	Host            types.ManagedObjectReference `xml:"host"`
	Type            string                       `xml:"type"`
	Name            string                       `xml:"name"`
	Version         string                       `xml:"version"`
	ExistingVersion string                       `xml:"existingVersion,omitempty"`
	Present         bool                         `xml:"present"`
	VibSpec         []VsanVibSpec                `xml:"vibSpec,omitempty"`
	VibType         string                       `xml:"vibType,omitempty"`
	FirmwareSpec    *VsanHclFirmwareUpdateSpec   `xml:"firmwareSpec,omitempty"`
	DownloadInfo    []VsanDownloadItem           `xml:"downloadInfo,omitempty"`
	Eula            string                       `xml:"eula,omitempty"`
	Adapter         string                       `xml:"adapter,omitempty"`
	Key             string                       `xml:"key,omitempty"`
	Impact          string                       `xml:"impact,omitempty"`
	FirmwareUnknown *bool                        `xml:"firmwareUnknown"`
}

type VsanMetricsConfig struct {
	types.DynamicData

	Profiles []VsanMetricProfile `xml:"profiles,omitempty"`
}

type VsanObjectInaccessibleIssue struct {
	types.VsanUpgradeSystemPreflightCheckIssue

	Uuids []string `xml:"uuids"`
}

type VsanDataProtectionRemoteTargetSiteInfo struct {
	types.DynamicData
}

type VsanFileShareRuntimeInfo struct {
	types.DynamicData

	UsedCapacity    int64            `xml:"usedCapacity,omitempty"`
	Hostname        string           `xml:"hostname,omitempty"`
	Address         string           `xml:"address,omitempty"`
	VsanObjectUuids []string         `xml:"vsanObjectUuids,omitempty"`
	AccessPoints    []types.KeyValue `xml:"accessPoints,omitempty"`
}

type VsanVdsPgMigrationVmInfo struct {
	types.DynamicData

	Vm        types.ManagedObjectReference `xml:"vm"`
	VnicLabel []string                     `xml:"vnicLabel"`
}

type VsanClusterHealthAction struct {
	types.DynamicData

	ActionId          string                   `xml:"actionId"`
	ActionLabel       types.LocalizableMessage `xml:"actionLabel"`
	ActionDescription types.LocalizableMessage `xml:"actionDescription"`
	Enabled           bool                     `xml:"enabled"`
}

type VsanClusterHealthSystemVersionResult struct {
	types.DynamicData

	HostResults     []VsanHostHealthSystemVersionResult `xml:"hostResults,omitempty"`
	VcVersion       string                              `xml:"vcVersion,omitempty"`
	IssueFound      bool                                `xml:"issueFound"`
	UpgradePossible *bool                               `xml:"upgradePossible"`
}

type VsanDeviceFirmwareInfo struct {
	types.DynamicData

	Host       string `xml:"host"`
	DeviceName string `xml:"deviceName"`
	FwVersion  string `xml:"fwVersion,omitempty"`
}

type CnsVolumeMetadataUpdateSpec struct {
	types.DynamicData

	VolumeId   CnsVolumeId       `xml:"volumeId"`
	Metadata   CnsVolumeMetadata `xml:"metadata"`
	VolumeType string            `xml:"volumeType,omitempty"`
}

type VsanPerfDiagnoseFeedbackSpec struct {
	types.DynamicData

	TransactionId string `xml:"transactionId"`
	ExceptionId   string `xml:"exceptionId"`
	EntityRefId   string `xml:"entityRefId,omitempty"`
}

type VsanClusterHealthResultRow struct {
	types.DynamicData

	Values     []string                     `xml:"values"`
	NestedRows []VsanClusterHealthResultRow `xml:"nestedRows,omitempty"`
}

type CnsVolumeId struct {
	types.DynamicData

	Id string `xml:"id"`
}

type HostSpbmPolicyBlobInfo struct {
	types.DynamicData

	PolicyBlob string `xml:"policyBlob"`
	Namespace  string `xml:"namespace"`
}

type VimVsanVsanFaultDomainComplianceResourceCheck struct {
	types.DynamicData

	FdName string                                   `xml:"fdName,omitempty"`
	Uuid   string                                   `xml:"uuid,omitempty"`
	IsNew  *bool                                    `xml:"isNew"`
	Hosts  []VimVsanVsanHostComplianceResourceCheck `xml:"hosts,omitempty"`
}

type VsanClusterHealthSystemStatusResult struct {
	types.DynamicData

	Status             string                             `xml:"status"`
	GoalState          string                             `xml:"goalState"`
	UntrackedHosts     []string                           `xml:"untrackedHosts,omitempty"`
	TrackedHostsStatus []VsanHostHealthSystemStatusResult `xml:"trackedHostsStatus,omitempty"`
}

type CnsBlockBackingDetails struct {
	CnsBackingObjectDetails

	BackingDiskId string `xml:"backingDiskId,omitempty"`
}

type VsanVcsaDeploymentSpec struct {
	types.DynamicData

	OsPassword        string   `xml:"osPassword"`
	OsSshEnabled      *bool    `xml:"osSshEnabled"`
	OsTimeToolsSync   *bool    `xml:"osTimeToolsSync"`
	OsNtpServers      []string `xml:"osNtpServers,omitempty"`
	SsoPassword       string   `xml:"ssoPassword"`
	SsoDomainName     string   `xml:"ssoDomainName,omitempty"`
	SsoSiteName       string   `xml:"ssoSiteName"`
	NetworkIpFamily   string   `xml:"networkIpFamily,omitempty"`
	NetworkMode       string   `xml:"networkMode,omitempty"`
	NetworkIpAddress  string   `xml:"networkIpAddress,omitempty"`
	NetworkPrefix     int64    `xml:"networkPrefix,omitempty"`
	NetworkGateway    string   `xml:"networkGateway,omitempty"`
	NetworkDnsServers []string `xml:"networkDnsServers,omitempty"`
	NetworkHostname   string   `xml:"networkHostname,omitempty"`
	NetworkPorts      string   `xml:"networkPorts,omitempty"`
}

type VsanCompositeConstraint struct {
	VsanResourceConstraint

	NestedConstraints []VsanResourceConstraint `xml:"nestedConstraints,omitempty"`
	Conjoiner         string                   `xml:"conjoiner,omitempty"`
}

type VsanClusterWhatifHostFailuresResult struct {
	types.DynamicData

	NumFailures             int64  `xml:"numFailures"`
	TotalUsedCapacityB      int64  `xml:"totalUsedCapacityB"`
	TotalCapacityB          int64  `xml:"totalCapacityB"`
	TotalRcReservationB     int64  `xml:"totalRcReservationB"`
	TotalRcSizeB            int64  `xml:"totalRcSizeB"`
	UsedComponents          int64  `xml:"usedComponents"`
	TotalComponents         int64  `xml:"totalComponents"`
	ComponentLimitHealth    string `xml:"componentLimitHealth,omitempty"`
	DiskFreeSpaceHealth     string `xml:"diskFreeSpaceHealth,omitempty"`
	RcFreeReservationHealth string `xml:"rcFreeReservationHealth,omitempty"`
}

type VsanIscsiTarget struct {
	VsanIscsiTargetCommonInfo

	LunCount          int32                  `xml:"lunCount,omitempty"`
	ObjectInformation *VsanObjectInformation `xml:"objectInformation,omitempty"`
	IoOwnerHost       string                 `xml:"ioOwnerHost,omitempty"`
	Initiators        []string               `xml:"initiators,omitempty"`
	InitiatorGroups   []string               `xml:"initiatorGroups,omitempty"`
}

type VimHostVSANStretchedClusterHostCapability struct {
	types.DynamicData

	FeatureVersion string `xml:"featureVersion"`
}

type VsanIscsiTargetServiceSpec struct {
	VsanIscsiTargetServiceConfig

	HomeObjectStoragePolicy *types.VirtualMachineProfileSpec `xml:"homeObjectStoragePolicy,omitempty"`
}

type VsanNodeNotMaster struct {
	types.VimFault

	VsanMasterUuid               string `xml:"vsanMasterUuid,omitempty"`
	CmmdsMasterButNotStatsMaster *bool  `xml:"cmmdsMasterButNotStatsMaster"`
}

type HostSpbmHashInfo struct {
	types.DynamicData

	PolicyInfoHash    string `xml:"policyInfoHash"`
	DatastoreInfoHash string `xml:"datastoreInfoHash"`
}

type CnsVolumeCreateSpec struct {
	types.DynamicData

	Name                 string                            `xml:"name"`
	VolumeType           string                            `xml:"volumeType"`
	Datastores           []types.ManagedObjectReference    `xml:"datastores,omitempty"`
	Metadata             *CnsVolumeMetadata                `xml:"metadata,omitempty"`
	BackingObjectDetails CnsBackingObjectDetails           `xml:"backingObjectDetails"`
	Profile              []types.VirtualMachineProfileSpec `xml:"profile,omitempty"`
	CreateSpec           *CnsBaseCreateSpec                `xml:"createSpec,omitempty"`
}

type VsanProactiveRebalanceInfo struct {
	types.DynamicData

	Enabled   *bool `xml:"enabled"`
	Threshold int32 `xml:"threshold,omitempty"`
}

type VsanHostHealthSystemVersionResult struct {
	types.DynamicData

	Hostname string             `xml:"hostname"`
	Version  string             `xml:"version,omitempty"`
	Error    *types.MethodFault `xml:"error,omitempty"`
}

type VsanHclControllerInfo struct {
	types.DynamicData

	DeviceName             string                   `xml:"deviceName"`
	DeviceDisplayName      string                   `xml:"deviceDisplayName,omitempty"`
	DriverName             string                   `xml:"driverName,omitempty"`
	DriverVersion          string                   `xml:"driverVersion,omitempty"`
	VendorId               int64                    `xml:"vendorId,omitempty"`
	DeviceId               int64                    `xml:"deviceId,omitempty"`
	SubVendorId            int64                    `xml:"subVendorId,omitempty"`
	SubDeviceId            int64                    `xml:"subDeviceId,omitempty"`
	ExtraInfo              []types.KeyValue         `xml:"extraInfo,omitempty"`
	DeviceOnHcl            *bool                    `xml:"deviceOnHcl"`
	ReleaseSupported       *bool                    `xml:"releaseSupported"`
	ReleasesOnHcl          []string                 `xml:"releasesOnHcl,omitempty"`
	DriverVersionsOnHcl    []string                 `xml:"driverVersionsOnHcl,omitempty"`
	DriverVersionSupported *bool                    `xml:"driverVersionSupported"`
	FwVersionSupported     *bool                    `xml:"fwVersionSupported"`
	FwVersionOnHcl         []string                 `xml:"fwVersionOnHcl,omitempty"`
	CacheConfigSupported   *bool                    `xml:"cacheConfigSupported"`
	CacheConfigOnHcl       []string                 `xml:"cacheConfigOnHcl,omitempty"`
	RaidConfigSupported    *bool                    `xml:"raidConfigSupported"`
	RaidConfigOnHcl        []string                 `xml:"raidConfigOnHcl,omitempty"`
	FwVersion              string                   `xml:"fwVersion,omitempty"`
	RaidConfig             string                   `xml:"raidConfig,omitempty"`
	CacheConfig            string                   `xml:"cacheConfig,omitempty"`
	CimProviderInfo        *VsanHostCimProviderInfo `xml:"cimProviderInfo,omitempty"`
	UsedByVsan             *bool                    `xml:"usedByVsan"`
	Disks                  []VsanHclDiskInfo        `xml:"disks,omitempty"`
	Issues                 []types.MethodFault      `xml:"issues,omitempty"`
	RemediableIssues       []string                 `xml:"remediableIssues,omitempty"`
	DriversOnHcl           []VsanHclDriverInfo      `xml:"driversOnHcl,omitempty"`
	FwAuxVersion           string                   `xml:"fwAuxVersion,omitempty"`
	QueueDepth             int64                    `xml:"queueDepth,omitempty"`
	QueueDepthOnHcl        int64                    `xml:"queueDepthOnHcl,omitempty"`
	QueueDepthSupported    *bool                    `xml:"queueDepthSupported"`
	DiskMode               string                   `xml:"diskMode,omitempty"`
	DiskModeOnHcl          []string                 `xml:"diskModeOnHcl,omitempty"`
	DiskModeSupported      *bool                    `xml:"diskModeSupported"`
	ToolName               string                   `xml:"toolName,omitempty"`
	ToolVersion            string                   `xml:"toolVersion,omitempty"`
}

type VsanHostDpdLivenessResult struct {
	types.DynamicData

	Hostname  string             `xml:"hostname"`
	DpdStatus string             `xml:"dpdStatus"`
	Error     *types.MethodFault `xml:"error,omitempty"`
}

type VSANStretchedClusterHostVirtualApplianceStatus struct {
	types.DynamicData

	VcCluster    *types.ManagedObjectReference `xml:"vcCluster,omitempty"`
	IsVirtualApp *bool                         `xml:"isVirtualApp"`
}

type VsanVibScanResult struct {
	types.DynamicData

	Host                    types.ManagedObjectReference `xml:"host"`
	VibName                 string                       `xml:"vibName"`
	VibVersion              string                       `xml:"vibVersion"`
	ExistingVersion         string                       `xml:"existingVersion,omitempty"`
	MaintenanceModeRequired bool                         `xml:"maintenanceModeRequired"`
	RebootRequired          bool                         `xml:"rebootRequired"`
	MeetsSystemReq          bool                         `xml:"meetsSystemReq"`
	PkgDepsMetByHost        bool                         `xml:"pkgDepsMetByHost"`
}

type CnsBackingObjectDetails struct {
	types.DynamicData

	CapacityInMb int64 `xml:"capacityInMb,omitempty"`
}

type VsanDisallowDataMovementIssue struct {
	types.VsanUpgradeSystemPreflightCheckIssue
}

type VsanHostDrsStats struct {
	types.DynamicData

	Host  types.ManagedObjectReference `xml:"host"`
	Stats []byte                       `xml:"stats"`
}

type CnsQueryFilter struct {
	types.DynamicData

	VolumeIds                    []CnsVolumeId                  `xml:"volumeIds,omitempty"`
	Names                        []string                       `xml:"names,omitempty"`
	ContainerClusterIds          []string                       `xml:"containerClusterIds,omitempty"`
	StoragePolicyId              string                         `xml:"storagePolicyId,omitempty"`
	Datastores                   []types.ManagedObjectReference `xml:"datastores,omitempty"`
	Labels                       []types.KeyValue               `xml:"labels,omitempty"`
	ComplianceStatus             string                         `xml:"complianceStatus,omitempty"`
	DatastoreAccessibilityStatus string                         `xml:"datastoreAccessibilityStatus,omitempty"`
	Cursor                       *CnsCursor                     `xml:"cursor,omitempty"`
}

type VsanClusterHealthCheckInfo struct {
	types.DynamicData

	TestId    string `xml:"testId"`
	TestName  string `xml:"testName,omitempty"`
	GroupId   string `xml:"groupId"`
	GroupName string `xml:"groupName,omitempty"`
}

type VsanFlrManagedMountSpec struct {
	VsanFlrMountSpec

	Snapshot types.ManagedObjectReference `xml:"snapshot"`
}

type VsanClusterPersistedState struct {
	types.DynamicData

	GenerationId           int64                          `xml:"generationId"`
	VcCluster              *types.ManagedObjectReference  `xml:"vcCluster,omitempty"`
	VsanEnabled            bool                           `xml:"vsanEnabled"`
	VsanClusterUuid        string                         `xml:"vsanClusterUuid,omitempty"`
	BasicConfig            *VsanClusterConfig             `xml:"basicConfig,omitempty"`
	Hosts                  []types.ManagedObjectReference `xml:"hosts,omitempty"`
	MetroConfig            *VsanClusterMetroConfig        `xml:"metroConfig,omitempty"`
	DataEfficiencyConfig   *VsanDataEfficiencyConfig      `xml:"dataEfficiencyConfig,omitempty"`
	PerfMemberInfos        []VsanPerfMemberInfo           `xml:"perfMemberInfos,omitempty"`
	DataEncryptionConfig   *VsanDataEncryptionConfig      `xml:"dataEncryptionConfig,omitempty"`
	ResyncIopsLimitConfig  *ResyncIopsInfo                `xml:"resyncIopsLimitConfig,omitempty"`
	IscsiEnabled           *bool                          `xml:"iscsiEnabled"`
	FileServiceConfig      *VsanFileServiceConfig         `xml:"fileServiceConfig,omitempty"`
	VsanDatastoreName      string                         `xml:"vsanDatastoreName,omitempty"`
	ExtendedConfig         *VsanExtendedConfig            `xml:"extendedConfig,omitempty"`
	DataProtectionConfig   *VsanDataProtectionInfo        `xml:"dataProtectionConfig,omitempty"`
	DatastoreConfig        *VsanDatastoreConfig           `xml:"datastoreConfig,omitempty"`
	GenNum                 int64                          `xml:"genNum,omitempty"`
	PerfsvcConfig          *VsanPerfsvcConfig             `xml:"perfsvcConfig,omitempty"`
	UnmapConfig            *VsanUnmapConfig               `xml:"unmapConfig,omitempty"`
	VsanBaselinePreference string                         `xml:"vsanBaselinePreference,omitempty"`
	RdmaConfig             *VsanRdmaConfig                `xml:"rdmaConfig,omitempty"`
	MetricsConfig          *VsanMetricsConfig             `xml:"metricsConfig,omitempty"`
}

type VsanCapability struct {
	types.DynamicData

	Target       *types.ManagedObjectReference `xml:"target,omitempty"`
	Capabilities []string                      `xml:"capabilities,omitempty"`
	Statuses     []string                      `xml:"statuses,omitempty"`
}

type VsanGenericClusterBaseIssue struct {
	types.DynamicData
}

type VsanFlrFileEntry struct {
	types.DynamicData

	VolumeId  string               `xml:"volumeId,omitempty"`
	FileId    string               `xml:"fileId"`
	ParentId  string               `xml:"parentId,omitempty"`
	Directory bool                 `xml:"directory"`
	Name      string               `xml:"name,omitempty"`
	Path      string               `xml:"path,omitempty"`
	Attrs     *types.GuestFileInfo `xml:"attrs,omitempty"`
}

type VsanFileShareConfig struct {
	types.DynamicData

	Name          string                           `xml:"name,omitempty"`
	DomainName    string                           `xml:"domainName,omitempty"`
	Quota         string                           `xml:"quota,omitempty"`
	SoftQuota     string                           `xml:"softQuota,omitempty"`
	Labels        []types.KeyValue                 `xml:"labels,omitempty"`
	StoragePolicy *types.VirtualMachineProfileSpec `xml:"storagePolicy,omitempty"`
	Permission    []VsanFileShareNetPermission     `xml:"permission,omitempty"`
}

type DataProtectionPeerSiteInfo struct {
	types.DynamicData

	Name                    string `xml:"name"`
	LookupServiceUrl        string `xml:"lookupServiceUrl,omitempty"`
	LookupServiceThumbprint string `xml:"lookupServiceThumbprint,omitempty"`
	SiteId                  string `xml:"siteId,omitempty"`
	NodeId                  string `xml:"nodeId,omitempty"`
}

type VsanGenericClusterBestPracticeHealth struct {
	types.DynamicData

	DrsEnabled bool                          `xml:"drsEnabled"`
	HaEnabled  bool                          `xml:"haEnabled"`
	Issues     []VsanGenericClusterBaseIssue `xml:"issues,omitempty"`
}

type VsanClusterEncryptionHealthSummary struct {
	types.DynamicData

	OverallHealth string                        `xml:"overallHealth,omitempty"`
	ConfigHealth  string                        `xml:"configHealth,omitempty"`
	KmsHealth     string                        `xml:"kmsHealth,omitempty"`
	VcKmsResult   *VsanVcKmipServersHealth      `xml:"vcKmsResult,omitempty"`
	HostResults   []VsanEncryptionHealthSummary `xml:"hostResults,omitempty"`
	AesniHealth   string                        `xml:"aesniHealth,omitempty"`
}

type VsanVmdkLoadTestResult struct {
	types.DynamicData

	Success                    bool                 `xml:"success"`
	FaultMessage               string               `xml:"faultMessage,omitempty"`
	Spec                       VsanVmdkLoadTestSpec `xml:"spec"`
	ActualDurationSec          int32                `xml:"actualDurationSec,omitempty"`
	TotalBytes                 int64                `xml:"totalBytes,omitempty"`
	Iops                       int64                `xml:"iops,omitempty"`
	TputBps                    int64                `xml:"tputBps,omitempty"`
	AvgLatencyUs               int64                `xml:"avgLatencyUs,omitempty"`
	MaxLatencyUs               int64                `xml:"maxLatencyUs,omitempty"`
	NumIoAboveLatencyThreshold int64                `xml:"numIoAboveLatencyThreshold,omitempty"`
}

type VsanDatastoreUsageResult struct {
	types.DynamicData

	ConfiguredUsageThreshold int32   `xml:"configuredUsageThreshold"`
	CurrentUsage             float32 `xml:"currentUsage"`
	ThresholdExceeded        bool    `xml:"thresholdExceeded"`
}

type VsanExtendedConfig struct {
	types.DynamicData

	ObjectRepairTimer          int64                       `xml:"objectRepairTimer,omitempty"`
	DisableSiteReadLocality    *bool                       `xml:"disableSiteReadLocality"`
	EnableCustomizedSwapObject *bool                       `xml:"enableCustomizedSwapObject"`
	LargeScaleClusterSupport   *bool                       `xml:"largeScaleClusterSupport"`
	ProactiveRebalanceInfo     *VsanProactiveRebalanceInfo `xml:"proactiveRebalanceInfo,omitempty"`
}

type VsanHclCommonDeviceInfo struct {
	types.DynamicData

	DeviceName             string              `xml:"deviceName"`
	DisplayName            string              `xml:"displayName,omitempty"`
	DriverName             string              `xml:"driverName,omitempty"`
	DriverVersion          string              `xml:"driverVersion,omitempty"`
	VendorId               int64               `xml:"vendorId,omitempty"`
	DeviceId               int64               `xml:"deviceId,omitempty"`
	SubVendorId            int64               `xml:"subVendorId,omitempty"`
	SubDeviceId            int64               `xml:"subDeviceId,omitempty"`
	ExtraInfo              []types.KeyValue    `xml:"extraInfo,omitempty"`
	DeviceOnHcl            *bool               `xml:"deviceOnHcl"`
	ReleaseSupported       *bool               `xml:"releaseSupported"`
	ReleasesOnHcl          []string            `xml:"releasesOnHcl,omitempty"`
	DriverVersionsOnHcl    []string            `xml:"driverVersionsOnHcl,omitempty"`
	DriverVersionSupported *bool               `xml:"driverVersionSupported"`
	FwVersionSupported     *bool               `xml:"fwVersionSupported"`
	FwVersionOnHcl         []string            `xml:"fwVersionOnHcl,omitempty"`
	FwVersion              string              `xml:"fwVersion,omitempty"`
	DriversOnHcl           []VsanHclDriverInfo `xml:"driversOnHcl,omitempty"`
}

type VsanIoInsightInternal struct {
	types.DynamicData

	State        string                         `xml:"state"`
	MonitoredVMs []types.ManagedObjectReference `xml:"monitoredVMs,omitempty"`
}

type VsanHostRuntimeStats struct {
	types.DynamicData

	ResyncIopsInfo       *ResyncIopsInfo       `xml:"resyncIopsInfo,omitempty"`
	ConfigGeneration     *VsanConfigGeneration `xml:"configGeneration,omitempty"`
	SupportedClusterSize int32                 `xml:"supportedClusterSize,omitempty"`
	RepairTimerInfo      *RepairTimerInfo      `xml:"repairTimerInfo,omitempty"`
}

type VsanPerfTimeRangeQuerySpec struct {
	types.DynamicData

	Name          string     `xml:"name,omitempty"`
	StartTimeFrom *time.Time `xml:"startTimeFrom"`
	StartTimeTo   *time.Time `xml:"startTimeTo"`
	EndTimeFrom   *time.Time `xml:"endTimeFrom"`
	EndTimeTo     *time.Time `xml:"endTimeTo"`
}

type VsanHclDeviceConstraint struct {
	types.DynamicData

	PciId               string                   `xml:"pciId"`
	DeviceFirmwareInfos []VsanDeviceFirmwareInfo `xml:"deviceFirmwareInfos,omitempty"`
	CompliantFirmwares  []VsanCompliantFirmware  `xml:"compliantFirmwares,omitempty"`
}

type VimClusterVSANWitnessHostInfo struct {
	types.DynamicData

	NodeUuid         string                        `xml:"nodeUuid"`
	FaultDomainName  string                        `xml:"faultDomainName,omitempty"`
	PreferredFdName  string                        `xml:"preferredFdName,omitempty"`
	PreferredFdUuid  string                        `xml:"preferredFdUuid,omitempty"`
	UnicastAgentAddr string                        `xml:"unicastAgentAddr,omitempty"`
	Host             *types.ManagedObjectReference `xml:"host,omitempty"`
	MetadataMode     *bool                         `xml:"metadataMode"`
}

type VsanVumSystemConfigSpec struct {
	types.DynamicData

	ClusterConfig []VsanVumSystemClusterConfig `xml:"clusterConfig,omitempty"`
}

type VsanHigherObjectsPresentDuringDowngradeIssue struct {
	types.VsanUpgradeSystemPreflightCheckIssue

	Uuids []string `xml:"uuids"`
}

type VimVsanHostVsanDiskManagementSystemCapability struct {
	types.DynamicData

	Version string `xml:"version"`
}

type VsanFlrVolume struct {
	types.DynamicData

	Id             string `xml:"id"`
	NativeId       string `xml:"nativeId"`
	Label          string `xml:"label,omitempty"`
	FileSystemType string `xml:"fileSystemType,omitempty"`
	MountPoint     string `xml:"mountPoint,omitempty"`
	Capacity       int64  `xml:"capacity,omitempty"`
	FreeSpace      int64  `xml:"freeSpace,omitempty"`
}

type VsanUpgradeStatusEx struct {
	types.VsanUpgradeSystemUpgradeStatus

	IsPrecheck     *bool                                `xml:"isPrecheck"`
	PrecheckResult *VsanDiskFormatConversionCheckResult `xml:"precheckResult,omitempty"`
}

type VsanFileShare struct {
	types.DynamicData

	Uuid    string                    `xml:"uuid"`
	Config  *VsanFileShareConfig      `xml:"config,omitempty"`
	Runtime *VsanFileShareRuntimeInfo `xml:"runtime,omitempty"`
}

type VsanComparator struct {
	types.DynamicData
}

type VsanObjectInformation struct {
	types.DynamicData

	DirectoryName            string                       `xml:"directoryName,omitempty"`
	VsanObjectUuid           string                       `xml:"vsanObjectUuid,omitempty"`
	VsanHealth               string                       `xml:"vsanHealth,omitempty"`
	PolicyAttributes         []types.KeyValue             `xml:"policyAttributes,omitempty"`
	SpbmProfileUuid          string                       `xml:"spbmProfileUuid,omitempty"`
	SpbmProfileGenerationId  string                       `xml:"spbmProfileGenerationId,omitempty"`
	SpbmComplianceResult     *VsanStorageComplianceResult `xml:"spbmComplianceResult,omitempty"`
	VsanDataProtectionHealth string                       `xml:"vsanDataProtectionHealth,omitempty"`
}

type VsanRdmaConfig struct {
	types.DynamicData

	RdmaEnabled bool `xml:"rdmaEnabled"`
}

type VsanVmdkIOLoadSpec struct {
	types.DynamicData

	ReadPct      int32 `xml:"readPct"`
	Oio          int32 `xml:"oio"`
	IosizeB      int32 `xml:"iosizeB"`
	DataSizeMb   int64 `xml:"dataSizeMb"`
	Random       bool  `xml:"random"`
	StartOffsetB int64 `xml:"startOffsetB,omitempty"`
}

type VsanFlrDpMountSpec struct {
	VsanFlrMountSpec

	CgId          string                         `xml:"cgId"`
	CgInstanceKey string                         `xml:"cgInstanceKey,omitempty"`
	Entities      []types.ManagedObjectReference `xml:"entities,omitempty"`
	Location      string                         `xml:"location,omitempty"`
}

type VsanHostClomdLivenessResult struct {
	types.DynamicData

	Hostname  string             `xml:"hostname"`
	ClomdStat string             `xml:"clomdStat"`
	Error     *types.MethodFault `xml:"error,omitempty"`
}

type VimVsanVsanScanObjectsIssue struct {
	types.DynamicData

	Uuid  string `xml:"uuid"`
	Issue string `xml:"issue"`
}

type VsanBurnInTestCheckResult struct {
	types.DynamicData

	PassedTests       []VsanBurnInTest `xml:"passedTests,omitempty"`
	NotPerformedTests []VsanBurnInTest `xml:"notPerformedTests,omitempty"`
	FailedTests       []VsanBurnInTest `xml:"failedTests,omitempty"`
}

type VsanDiskResourceCheckResult struct {
	EntityResourceCheckDetails
}

type VsanPerfsvcHealthResult struct {
	types.DynamicData

	StatsObjectInfo             *VsanObjectInformation    `xml:"statsObjectInfo,omitempty"`
	StatsObjectConsistent       *bool                     `xml:"statsObjectConsistent"`
	StatsObjectPolicyConsistent *bool                     `xml:"statsObjectPolicyConsistent"`
	DatastoreCompatible         *bool                     `xml:"datastoreCompatible"`
	EnoughFreeSpace             *bool                     `xml:"enoughFreeSpace"`
	RemediateAction             string                    `xml:"remediateAction,omitempty"`
	HostResults                 []VsanPerfNodeInformation `xml:"hostResults,omitempty"`
	VerboseModeStatus           *bool                     `xml:"verboseModeStatus"`
}

type VsanRemoteDatastoreSpec struct {
	types.DynamicData

	ClusterUuid   string                      `xml:"clusterUuid"`
	DatastoreSpec []VsanDatastoreSpec         `xml:"datastoreSpec"`
	UnicastInfo   []VsanServerHostUnicastInfo `xml:"unicastInfo,omitempty"`
}

type VsanWitnessHostConfig struct {
	types.DynamicData

	SubClusterUuid           string `xml:"subClusterUuid"`
	PreferredFaultDomainName string `xml:"preferredFaultDomainName"`
	MetadataMode             *bool  `xml:"metadataMode"`
}

type VsanWhatifCapacity struct {
	types.DynamicData

	TotalWhatifCapacityB int64                           `xml:"totalWhatifCapacityB"`
	FreeWhatifCapacityB  int64                           `xml:"freeWhatifCapacityB"`
	StoragePolicy        types.VirtualMachineProfileSpec `xml:"storagePolicy"`
	IsSatisfiable        bool                            `xml:"isSatisfiable"`
}

type VsanPrepareVsanForVcsaSpec struct {
	types.DynamicData

	VsanDiskMappingCreationSpec *VimVsanHostDiskMappingCreationSpec `xml:"vsanDiskMappingCreationSpec,omitempty"`
	VsanDataEfficiencyConfig    *VsanDataEfficiencyConfig           `xml:"vsanDataEfficiencyConfig,omitempty"`
	TaskId                      string                              `xml:"taskId,omitempty"`
	VsanDataEncryptionConfig    *VsanHostEncryptionInfo             `xml:"vsanDataEncryptionConfig,omitempty"`
}

type VsanHostHclInfo struct {
	types.DynamicData

	Hostname    string                  `xml:"hostname"`
	HclChecked  bool                    `xml:"hclChecked"`
	ReleaseName string                  `xml:"releaseName,omitempty"`
	Error       *types.MethodFault      `xml:"error,omitempty"`
	Controllers []VsanHclControllerInfo `xml:"controllers,omitempty"`
	Pnics       []VsanHclNicInfo        `xml:"pnics,omitempty"`
}

type VsanHostPortConfigEx struct {
	types.VsanHostConfigInfoNetworkInfoPortConfig

	TrafficTypes []string `xml:"trafficTypes,omitempty"`
}

type VsanFlrQueryFileResult struct {
	types.DynamicData

	Files []VsanFlrFileEntry `xml:"files,omitempty"`
}

type CnsBaseCreateSpec struct {
	types.DynamicData
}

type CnsVolumeOperationBatchResult struct {
	types.DynamicData

	VolumeResults []CnsVolumeOperationResult `xml:"volumeResults,omitempty"`
}

type VsanVdsMigrationPlan struct {
	types.DynamicData

	VdsSpec         types.DVSCreateSpec            `xml:"vdsSpec"`
	Pgs             []VsanVdsPgMigrationSpec       `xml:"pgs,omitempty"`
	InaccessibleVms []types.ManagedObjectReference `xml:"inaccessibleVms,omitempty"`
	InfraVms        []types.ManagedObjectReference `xml:"infraVms,omitempty"`
}

type VsanIscsiHomeObjectSpec struct {
	types.DynamicData

	StoragePolicy *types.VirtualMachineProfileSpec         `xml:"storagePolicy,omitempty"`
	DefaultConfig *VsanIscsiTargetServiceDefaultConfigSpec `xml:"defaultConfig,omitempty"`
}

type VsanPropertyConstraint struct {
	VsanResourceConstraint

	PropertyName    string             `xml:"propertyName,omitempty"`
	Comparator      string             `xml:"comparator,omitempty"`
	ComparableValue *types.KeyAnyValue `xml:"comparableValue,omitempty"`
}

type VimHostVsanRecoveryComponentInfo struct {
	types.DynamicData

	Offset         int64  `xml:"offset"`
	Raid0Col       int64  `xml:"raid0Col,omitempty"`
	Raid0ColCount  int64  `xml:"raid0ColCount,omitempty"`
	Uuid           string `xml:"uuid"`
	DiskUuid       string `xml:"diskUuid"`
	Host           string `xml:"host"`
	Size           int64  `xml:"size"`
	StaleCsn       int32  `xml:"staleCsn,omitempty"`
	ComponentState int32  `xml:"componentState"`
}

type VsanHostVsanAwsCredentials struct {
	types.DynamicData

	Region          string `xml:"region,omitempty"`
	AccessKeyId     string `xml:"accessKeyId"`
	SecretAccessKey string `xml:"secretAccessKey"`
	SessionToken    string `xml:"sessionToken,omitempty"`
}

type VsanDataObfuscationRule struct {
	types.DynamicData
}

type VsanObjectHealth struct {
	types.DynamicData

	NumObjects           int32    `xml:"numObjects"`
	Health               string   `xml:"health"`
	ObjUuids             []string `xml:"objUuids,omitempty"`
	DataProtectionHealth string   `xml:"dataProtectionHealth,omitempty"`
}

type VsanClusterTelemetryProxyConfig struct {
	types.DynamicData

	Host           string `xml:"host,omitempty"`
	Port           int32  `xml:"port,omitempty"`
	User           string `xml:"user,omitempty"`
	Password       string `xml:"password,omitempty"`
	AutoDiscovered *bool  `xml:"autoDiscovered"`
}

type VsanFlrSearchFileSpec struct {
	VsanFlrQueryFileSpec

	NameRegEx string `xml:"nameRegEx,omitempty"`
}

type CnsKubernetesQueryFilter struct {
	CnsQueryFilter

	Namespaces []string `xml:"namespaces,omitempty"`
	PodNames   []string `xml:"podNames,omitempty"`
	PvcNames   []string `xml:"pvcNames,omitempty"`
	PvNames    []string `xml:"pvNames,omitempty"`
}

type VsanHostConfigInfoEx struct {
	types.VsanHostConfigInfo

	EncryptionInfo         *VsanHostEncryptionInfo     `xml:"encryptionInfo,omitempty"`
	DataEfficiencyInfo     *VsanDataEfficiencyConfig   `xml:"dataEfficiencyInfo,omitempty"`
	ResyncIopsLimitInfo    *ResyncIopsInfo             `xml:"resyncIopsLimitInfo,omitempty"`
	ExtendedConfig         *VsanExtendedConfig         `xml:"extendedConfig,omitempty"`
	DataProtectionInfo     *VsanDataProtectionInfo     `xml:"dataProtectionInfo,omitempty"`
	DatastoreInfo          *VsanDatastoreConfig        `xml:"datastoreInfo,omitempty"`
	UnmapConfig            *VsanUnmapConfig            `xml:"unmapConfig,omitempty"`
	WitnessHostConfig      []VsanWitnessHostConfig     `xml:"witnessHostConfig,omitempty"`
	RdmaConfig             *VsanRdmaConfig             `xml:"rdmaConfig,omitempty"`
	InternalExtendedConfig *VsanInternalExtendedConfig `xml:"internalExtendedConfig,omitempty"`
	MetricsConfig          *VsanMetricsConfig          `xml:"metricsConfig,omitempty"`
}

type VimClusterVSANPreferredFaultDomainInfo struct {
	types.DynamicData

	PreferredFaultDomainName string `xml:"preferredFaultDomainName,omitempty"`
	PreferredFaultDomainId   string `xml:"preferredFaultDomainId,omitempty"`
}

type VsanVsanPcapResult struct {
	types.DynamicData

	Calltime      float32            `xml:"calltime"`
	Vmknic        string             `xml:"vmknic"`
	TcpdumpFilter string             `xml:"tcpdumpFilter"`
	Snaplen       int32              `xml:"snaplen"`
	Pkts          []string           `xml:"pkts,omitempty"`
	Pcap          string             `xml:"pcap,omitempty"`
	Error         *types.MethodFault `xml:"error,omitempty"`
	Hostname      string             `xml:"hostname,omitempty"`
}

type VsanServerHostUnicastInfo struct {
	types.DynamicData

	HostUuid    string                   `xml:"hostUuid"`
	UnicastSpec []VsanUnicastAddressInfo `xml:"unicastSpec,omitempty"`
}

type VimHostVsanRecoveryHostDiskInfo struct {
	types.DynamicData

	HostUuid string `xml:"hostUuid,omitempty"`
	DiskUuid string `xml:"diskUuid,omitempty"`
}

type VsanHostAssociatedObjects struct {
	types.DynamicData

	SpbmProfileId            string   `xml:"spbmProfileId"`
	SpbmProfileGenerationNum int32    `xml:"spbmProfileGenerationNum"`
	VsanObjects              []string `xml:"vsanObjects,omitempty"`
}

type VsanFileServiceDomainQuerySpec struct {
	types.DynamicData

	Uuids []string `xml:"uuids,omitempty"`
	Names []string `xml:"names,omitempty"`
}

type VsanVibInstallPreflightStatus struct {
	types.DynamicData

	ManualVmotionRequired bool `xml:"manualVmotionRequired"`
	RollingRequired       bool `xml:"rollingRequired"`
}

type VsanRegexBasedRule struct {
	types.DynamicData

	Rules []string `xml:"rules,omitempty"`
}
