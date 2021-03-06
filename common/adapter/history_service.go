// Copyright (c) 2019 Temporal Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Code generated by generate-adapter. DO NOT EDIT.

package adapter

import (
	"github.com/temporalio/temporal/.gen/go/history"
	"github.com/temporalio/temporal/.gen/go/replicator"
	"github.com/temporalio/temporal/.gen/go/shared"
	"github.com/temporalio/temporal/.gen/proto/historyservice"
)

// ToProtoStartWorkflowExecutionResponse ...
func ToProtoHistoryStartWorkflowExecutionResponse(in *shared.StartWorkflowExecutionResponse) *historyservice.StartWorkflowExecutionResponse {
	if in == nil {
		return nil
	}
	return &historyservice.StartWorkflowExecutionResponse{
		RunId: in.GetRunId(),
	}
}

// ToProtoGetMutableStateResponse ...
func ToProtoGetMutableStateResponse(in *history.GetMutableStateResponse) *historyservice.GetMutableStateResponse {
	if in == nil {
		return nil
	}
	return &historyservice.GetMutableStateResponse{
		Execution:                            ToProtoWorkflowExecution(in.GetExecution()),
		WorkflowType:                         ToProtoWorkflowType(in.GetWorkflowType()),
		NextEventId:                          in.GetNextEventId(),
		PreviousStartedEventId:               in.GetPreviousStartedEventId(),
		LastFirstEventId:                     in.GetLastFirstEventId(),
		TaskList:                             ToProtoTaskList(in.GetTaskList()),
		StickyTaskList:                       ToProtoTaskList(in.GetStickyTaskList()),
		ClientLibraryVersion:                 in.GetClientLibraryVersion(),
		ClientFeatureVersion:                 in.GetClientFeatureVersion(),
		ClientImpl:                           in.GetClientImpl(),
		IsWorkflowRunning:                    in.GetIsWorkflowRunning(),
		StickyTaskListScheduleToStartTimeout: in.GetStickyTaskListScheduleToStartTimeout(),
		EventStoreVersion:                    in.GetEventStoreVersion(),
		CurrentBranchToken:                   in.GetCurrentBranchToken(),
		ReplicationInfo:                      ToProtoReplicationInfos(in.GetReplicationInfo()),
		WorkflowState:                        in.GetWorkflowState(),
		WorkflowCloseState:                   in.GetWorkflowCloseState(),
		VersionHistories:                     ToProtoVersionHistories(in.GetVersionHistories()),
		IsStickyTaskListEnabled:              in.GetIsStickyTaskListEnabled(),
	}
}

// ToProtoPollMutableStateResponse ...
func ToProtoPollMutableStateResponse(in *history.PollMutableStateResponse) *historyservice.PollMutableStateResponse {
	if in == nil {
		return nil
	}
	return &historyservice.PollMutableStateResponse{
		Execution:                            ToProtoWorkflowExecution(in.GetExecution()),
		WorkflowType:                         ToProtoWorkflowType(in.GetWorkflowType()),
		NextEventId:                          in.GetNextEventId(),
		PreviousStartedEventId:               in.GetPreviousStartedEventId(),
		LastFirstEventId:                     in.GetLastFirstEventId(),
		TaskList:                             ToProtoTaskList(in.GetTaskList()),
		StickyTaskList:                       ToProtoTaskList(in.GetStickyTaskList()),
		ClientLibraryVersion:                 in.GetClientLibraryVersion(),
		ClientFeatureVersion:                 in.GetClientFeatureVersion(),
		ClientImpl:                           in.GetClientImpl(),
		StickyTaskListScheduleToStartTimeout: in.GetStickyTaskListScheduleToStartTimeout(),
		CurrentBranchToken:                   in.GetCurrentBranchToken(),
		ReplicationInfo:                      ToProtoReplicationInfos(in.GetReplicationInfo()),
		VersionHistories:                     ToProtoVersionHistories(in.GetVersionHistories()),
		WorkflowState:                        in.GetWorkflowState(),
		WorkflowCloseState:                   in.GetWorkflowCloseState(),
	}
}

// ToProtoResetStickyTaskListResponse ...
func ToProtoResetStickyTaskListResponse(in *history.ResetStickyTaskListResponse) *historyservice.ResetStickyTaskListResponse {
	if in == nil {
		return nil
	}
	return &historyservice.ResetStickyTaskListResponse{}
}

// ToProtoRecordDecisionTaskStartedResponse ...
func ToProtoRecordDecisionTaskStartedResponse(in *history.RecordDecisionTaskStartedResponse) *historyservice.RecordDecisionTaskStartedResponse {
	if in == nil {
		return nil
	}
	return &historyservice.RecordDecisionTaskStartedResponse{
		WorkflowType:              ToProtoWorkflowType(in.GetWorkflowType()),
		PreviousStartedEventId:    in.GetPreviousStartedEventId(),
		ScheduledEventId:          in.GetScheduledEventId(),
		StartedEventId:            in.GetStartedEventId(),
		NextEventId:               in.GetNextEventId(),
		Attempt:                   in.GetAttempt(),
		StickyExecutionEnabled:    in.GetStickyExecutionEnabled(),
		DecisionInfo:              ToProtoTransientDecisionInfo(in.GetDecisionInfo()),
		WorkflowExecutionTaskList: ToProtoTaskList(in.GetWorkflowExecutionTaskList()),
		EventStoreVersion:         in.GetEventStoreVersion(),
		BranchToken:               in.GetBranchToken(),
		ScheduledTimestamp:        in.GetScheduledTimestamp(),
		StartedTimestamp:          in.GetStartedTimestamp(),
		Queries:                   ToProtoWorkflowQueries(in.GetQueries()),
	}
}

// ToProtoRecordActivityTaskStartedResponse ...
func ToProtoRecordActivityTaskStartedResponse(in *history.RecordActivityTaskStartedResponse) *historyservice.RecordActivityTaskStartedResponse {
	if in == nil {
		return nil
	}
	return &historyservice.RecordActivityTaskStartedResponse{
		ScheduledEvent:                  ToProtoHistoryEvent(in.GetScheduledEvent()),
		StartedTimestamp:                in.GetStartedTimestamp(),
		Attempt:                         in.GetAttempt(),
		ScheduledTimestampOfThisAttempt: in.GetScheduledTimestampOfThisAttempt(),
		HeartbeatDetails:                in.GetHeartbeatDetails(),
		WorkflowType:                    ToProtoWorkflowType(in.GetWorkflowType()),
		WorkflowDomain:                  in.GetWorkflowDomain(),
	}
}

// ToProtoRespondDecisionTaskCompletedResponse ...
func ToProtoHistoryRespondDecisionTaskCompletedResponse(in *history.RespondDecisionTaskCompletedResponse) *historyservice.RespondDecisionTaskCompletedResponse {
	if in == nil {
		return nil
	}
	return &historyservice.RespondDecisionTaskCompletedResponse{
		StartedResponse: ToProtoRecordDecisionTaskStartedResponse(in.GetStartedResponse()),
	}
}

// ToProtoRecordActivityTaskHeartbeatResponse ...
func ToProtoHistoryRecordActivityTaskHeartbeatResponse(in *shared.RecordActivityTaskHeartbeatResponse) *historyservice.RecordActivityTaskHeartbeatResponse {
	if in == nil {
		return nil
	}
	return &historyservice.RecordActivityTaskHeartbeatResponse{
		CancelRequested: in.GetCancelRequested(),
	}
}

// ToProtoResetWorkflowExecutionResponse ...
func ToProtoHistorySignalWithStartWorkflowExecutionResponse(in *shared.StartWorkflowExecutionResponse) *historyservice.SignalWithStartWorkflowExecutionResponse {
	if in == nil {
		return nil
	}
	return &historyservice.SignalWithStartWorkflowExecutionResponse{
		RunId: in.GetRunId(),
	}
}

// ToProtoResetWorkflowExecutionResponse ...
func ToProtoHistoryResetWorkflowExecutionResponse(in *shared.ResetWorkflowExecutionResponse) *historyservice.ResetWorkflowExecutionResponse {
	if in == nil {
		return nil
	}
	return &historyservice.ResetWorkflowExecutionResponse{
		RunId: in.GetRunId(),
	}
}

// ToProtoDescribeWorkflowExecutionResponse ...
func ToProtoHistoryDescribeWorkflowExecutionResponse(in *shared.DescribeWorkflowExecutionResponse) *historyservice.DescribeWorkflowExecutionResponse {
	if in == nil {
		return nil
	}
	return &historyservice.DescribeWorkflowExecutionResponse{
		ExecutionConfiguration: ToProtoWorkflowExecutionConfiguration(in.GetExecutionConfiguration()),
		WorkflowExecutionInfo:  ToProtoWorkflowExecutionInfo(in.GetWorkflowExecutionInfo()),
		PendingActivities:      ToProtoPendingActivityInfos(in.GetPendingActivities()),
		PendingChildren:        ToProtoPendingChildExecutionInfos(in.GetPendingChildren()),
	}
}

// ToProtoDescribeMutableStateResponse ...
func ToProtoDescribeMutableStateResponse(in *history.DescribeMutableStateResponse) *historyservice.DescribeMutableStateResponse {
	if in == nil {
		return nil
	}
	return &historyservice.DescribeMutableStateResponse{
		MutableStateInCache:    in.GetMutableStateInCache(),
		MutableStateInDatabase: in.GetMutableStateInDatabase(),
	}
}

// ToProtoDescribeHistoryHostResponse ...
func ToProtoHistoryDescribeHistoryHostResponse(in *shared.DescribeHistoryHostResponse) *historyservice.DescribeHistoryHostResponse {
	if in == nil {
		return nil
	}
	return &historyservice.DescribeHistoryHostResponse{
		NumberOfShards:        in.GetNumberOfShards(),
		ShardIDs:              in.GetShardIDs(),
		DomainCache:           ToProtoDomainCacheInfo(in.GetDomainCache()),
		ShardControllerStatus: in.GetShardControllerStatus(),
		Address:               in.GetAddress(),
	}
}

// ToProtoGetReplicationMessagesResponse ...
func ToProtoHistoryGetReplicationMessagesResponse(in *replicator.GetReplicationMessagesResponse) *historyservice.GetReplicationMessagesResponse {
	if in == nil {
		return nil
	}
	return &historyservice.GetReplicationMessagesResponse{
		MessagesByShard: ToProtoReplicationMessagess(in.GetMessagesByShard()),
	}
}

// ToProtoGetDLQReplicationMessagesResponse ...
func ToProtoHistoryGetDLQReplicationMessagesResponse(in *replicator.GetDLQReplicationMessagesResponse) *historyservice.GetDLQReplicationMessagesResponse {
	if in == nil {
		return nil
	}
	return &historyservice.GetDLQReplicationMessagesResponse{
		ReplicationTasks: ToProtoReplicationTasks(in.GetReplicationTasks()),
	}
}

// ToProtoQueryWorkflowResponse ...
func ToProtoHistoryQueryWorkflowResponse(in *history.QueryWorkflowResponse) *historyservice.QueryWorkflowResponse {
	if in == nil {
		return nil
	}
	return &historyservice.QueryWorkflowResponse{
		Response: ToProtoQueryWorkflowResponse(in.GetResponse()),
	}
}

// ToThriftStartWorkflowExecutionRequest ...
func ToThriftHistoryStartWorkflowExecutionRequest(in *historyservice.StartWorkflowExecutionRequest) *history.StartWorkflowExecutionRequest {
	if in == nil {
		return nil
	}
	return &history.StartWorkflowExecutionRequest{
		DomainUUID:                      &in.DomainUUID,
		StartRequest:                    ToThriftStartWorkflowExecutionRequest(in.StartRequest),
		ParentExecutionInfo:             ToThriftParentExecutionInfo(in.ParentExecutionInfo),
		Attempt:                         &in.Attempt,
		ExpirationTimestamp:             &in.ExpirationTimestamp,
		ContinueAsNewInitiator:          ToThriftContinueAsNewInitiator(in.ContinueAsNewInitiator),
		ContinuedFailureReason:          &in.ContinuedFailureReason,
		ContinuedFailureDetails:         in.ContinuedFailureDetails,
		LastCompletionResult:            in.LastCompletionResult,
		FirstDecisionTaskBackoffSeconds: &in.FirstDecisionTaskBackoffSeconds,
	}
}

// ToThriftGetMutableStateRequest ...
func ToThriftGetMutableStateRequest(in *historyservice.GetMutableStateRequest) *history.GetMutableStateRequest {
	if in == nil {
		return nil
	}
	return &history.GetMutableStateRequest{
		DomainUUID:          &in.DomainUUID,
		Execution:           ToThriftWorkflowExecution(in.Execution),
		ExpectedNextEventId: &in.ExpectedNextEventId,
		CurrentBranchToken:  in.CurrentBranchToken,
	}
}

// ToThriftPollMutableStateRequest ...
func ToThriftPollMutableStateRequest(in *historyservice.PollMutableStateRequest) *history.PollMutableStateRequest {
	if in == nil {
		return nil
	}
	return &history.PollMutableStateRequest{
		DomainUUID:          &in.DomainUUID,
		Execution:           ToThriftWorkflowExecution(in.Execution),
		ExpectedNextEventId: &in.ExpectedNextEventId,
		CurrentBranchToken:  in.CurrentBranchToken,
	}
}

// ToThriftResetStickyTaskListRequest ...
func ToThriftHistoryResetStickyTaskListRequest(in *historyservice.ResetStickyTaskListRequest) *history.ResetStickyTaskListRequest {
	if in == nil {
		return nil
	}
	return &history.ResetStickyTaskListRequest{
		DomainUUID: &in.DomainUUID,
		Execution:  ToThriftWorkflowExecution(in.Execution),
	}
}

// ToThriftRecordDecisionTaskStartedRequest ...
func ToThriftRecordDecisionTaskStartedRequest(in *historyservice.RecordDecisionTaskStartedRequest) *history.RecordDecisionTaskStartedRequest {
	if in == nil {
		return nil
	}
	return &history.RecordDecisionTaskStartedRequest{
		DomainUUID:        &in.DomainUUID,
		WorkflowExecution: ToThriftWorkflowExecution(in.WorkflowExecution),
		ScheduleId:        &in.ScheduleId,
		TaskId:            &in.TaskId,
		RequestId:         &in.RequestId,
		PollRequest:       ToThriftPollForDecisionTaskRequest(in.PollRequest),
	}
}

// ToThriftRecordActivityTaskStartedRequest ...
func ToThriftRecordActivityTaskStartedRequest(in *historyservice.RecordActivityTaskStartedRequest) *history.RecordActivityTaskStartedRequest {
	if in == nil {
		return nil
	}
	return &history.RecordActivityTaskStartedRequest{
		DomainUUID:        &in.DomainUUID,
		WorkflowExecution: ToThriftWorkflowExecution(in.WorkflowExecution),
		ScheduleId:        &in.ScheduleId,
		TaskId:            &in.TaskId,
		RequestId:         &in.RequestId,
		PollRequest:       ToThriftPollForActivityTaskRequest(in.PollRequest),
	}
}

// ToThriftRespondDecisionTaskCompletedRequest ...
func ToThriftHistoryRespondDecisionTaskCompletedRequest(in *historyservice.RespondDecisionTaskCompletedRequest) *history.RespondDecisionTaskCompletedRequest {
	if in == nil {
		return nil
	}
	return &history.RespondDecisionTaskCompletedRequest{
		DomainUUID:      &in.DomainUUID,
		CompleteRequest: ToThriftRespondDecisionTaskCompletedRequest(in.CompleteRequest),
	}
}

// ToThriftRespondDecisionTaskFailedRequest ...
func ToThriftHistoryRespondDecisionTaskFailedRequest(in *historyservice.RespondDecisionTaskFailedRequest) *history.RespondDecisionTaskFailedRequest {
	if in == nil {
		return nil
	}
	return &history.RespondDecisionTaskFailedRequest{
		DomainUUID:    &in.DomainUUID,
		FailedRequest: ToThriftRespondDecisionTaskFailedRequest(in.FailedRequest),
	}
}

// ToThriftRecordActivityTaskHeartbeatRequest ...
func ToThriftHistoryRecordActivityTaskHeartbeatRequest(in *historyservice.RecordActivityTaskHeartbeatRequest) *history.RecordActivityTaskHeartbeatRequest {
	if in == nil {
		return nil
	}
	return &history.RecordActivityTaskHeartbeatRequest{
		DomainUUID:       &in.DomainUUID,
		HeartbeatRequest: ToThriftRecordActivityTaskHeartbeatRequest(in.HeartbeatRequest),
	}
}

// ToThriftRespondActivityTaskCompletedRequest ...
func ToThriftHistoryRespondActivityTaskCompletedRequest(in *historyservice.RespondActivityTaskCompletedRequest) *history.RespondActivityTaskCompletedRequest {
	if in == nil {
		return nil
	}
	return &history.RespondActivityTaskCompletedRequest{
		DomainUUID:      &in.DomainUUID,
		CompleteRequest: ToThriftRespondActivityTaskCompletedRequest(in.CompleteRequest),
	}
}

// ToThriftRespondActivityTaskFailedRequest ...
func ToThriftHistoryRespondActivityTaskFailedRequest(in *historyservice.RespondActivityTaskFailedRequest) *history.RespondActivityTaskFailedRequest {
	if in == nil {
		return nil
	}
	return &history.RespondActivityTaskFailedRequest{
		DomainUUID:    &in.DomainUUID,
		FailedRequest: ToThriftRespondActivityTaskFailedRequest(in.FailedRequest),
	}
}

// ToThriftRespondActivityTaskCanceledRequest ...
func ToThriftHistoryRespondActivityTaskCanceledRequest(in *historyservice.RespondActivityTaskCanceledRequest) *history.RespondActivityTaskCanceledRequest {
	if in == nil {
		return nil
	}
	return &history.RespondActivityTaskCanceledRequest{
		DomainUUID:    &in.DomainUUID,
		CancelRequest: ToThriftRespondActivityTaskCanceledRequest(in.CancelRequest),
	}
}

// ToThriftSignalWorkflowExecutionRequest ...
func ToThriftHistorySignalWorkflowExecutionRequest(in *historyservice.SignalWorkflowExecutionRequest) *history.SignalWorkflowExecutionRequest {
	if in == nil {
		return nil
	}
	return &history.SignalWorkflowExecutionRequest{
		DomainUUID:                &in.DomainUUID,
		SignalRequest:             ToThriftSignalWorkflowExecutionRequest(in.SignalRequest),
		ExternalWorkflowExecution: ToThriftWorkflowExecution(in.ExternalWorkflowExecution),
		ChildWorkflowOnly:         &in.ChildWorkflowOnly,
	}
}

// ToThriftSignalWithStartWorkflowExecutionRequest ...
func ToThriftHistorySignalWithStartWorkflowExecutionRequest(in *historyservice.SignalWithStartWorkflowExecutionRequest) *history.SignalWithStartWorkflowExecutionRequest {
	if in == nil {
		return nil
	}
	return &history.SignalWithStartWorkflowExecutionRequest{
		DomainUUID:             &in.DomainUUID,
		SignalWithStartRequest: ToThriftSignalWithStartWorkflowExecutionRequest(in.SignalWithStartRequest),
	}
}

// ToThriftRemoveSignalMutableStateRequest ...
func ToThriftRemoveSignalMutableStateRequest(in *historyservice.RemoveSignalMutableStateRequest) *history.RemoveSignalMutableStateRequest {
	if in == nil {
		return nil
	}
	return &history.RemoveSignalMutableStateRequest{
		DomainUUID:        &in.DomainUUID,
		WorkflowExecution: ToThriftWorkflowExecution(in.WorkflowExecution),
		RequestId:         &in.RequestId,
	}
}

// ToThriftTerminateWorkflowExecutionRequest ...
func ToThriftHistoryTerminateWorkflowExecutionRequest(in *historyservice.TerminateWorkflowExecutionRequest) *history.TerminateWorkflowExecutionRequest {
	if in == nil {
		return nil
	}
	return &history.TerminateWorkflowExecutionRequest{
		DomainUUID:       &in.DomainUUID,
		TerminateRequest: ToThriftTerminateWorkflowExecutionRequest(in.TerminateRequest),
	}
}

// ToThriftResetWorkflowExecutionRequest ...
func ToThriftHistoryResetWorkflowExecutionRequest(in *historyservice.ResetWorkflowExecutionRequest) *history.ResetWorkflowExecutionRequest {
	if in == nil {
		return nil
	}
	return &history.ResetWorkflowExecutionRequest{
		DomainUUID:   &in.DomainUUID,
		ResetRequest: ToThriftResetWorkflowExecutionRequest(in.ResetRequest),
	}
}

// ToThriftRequestCancelWorkflowExecutionRequest ...
func ToThriftHistoryRequestCancelWorkflowExecutionRequest(in *historyservice.RequestCancelWorkflowExecutionRequest) *history.RequestCancelWorkflowExecutionRequest {
	if in == nil {
		return nil
	}
	return &history.RequestCancelWorkflowExecutionRequest{
		DomainUUID:                &in.DomainUUID,
		CancelRequest:             ToThriftRequestCancelWorkflowExecutionRequest(in.CancelRequest),
		ExternalInitiatedEventId:  &in.ExternalInitiatedEventId,
		ExternalWorkflowExecution: ToThriftWorkflowExecution(in.ExternalWorkflowExecution),
		ChildWorkflowOnly:         &in.ChildWorkflowOnly,
	}
}

// ToThriftScheduleDecisionTaskRequest ...
func ToThriftScheduleDecisionTaskRequest(in *historyservice.ScheduleDecisionTaskRequest) *history.ScheduleDecisionTaskRequest {
	if in == nil {
		return nil
	}
	return &history.ScheduleDecisionTaskRequest{
		DomainUUID:        &in.DomainUUID,
		WorkflowExecution: ToThriftWorkflowExecution(in.WorkflowExecution),
		IsFirstDecision:   &in.IsFirstDecision,
	}
}

// ToThriftRecordChildExecutionCompletedRequest ...
func ToThriftRecordChildExecutionCompletedRequest(in *historyservice.RecordChildExecutionCompletedRequest) *history.RecordChildExecutionCompletedRequest {
	if in == nil {
		return nil
	}
	return &history.RecordChildExecutionCompletedRequest{
		DomainUUID:         &in.DomainUUID,
		WorkflowExecution:  ToThriftWorkflowExecution(in.WorkflowExecution),
		InitiatedId:        &in.InitiatedId,
		CompletedExecution: ToThriftWorkflowExecution(in.CompletedExecution),
		CompletionEvent:    ToThriftHistoryEvent(in.CompletionEvent),
	}
}

// ToThriftDescribeWorkflowExecutionRequest ...
func ToThriftHistoryDescribeWorkflowExecutionRequest(in *historyservice.DescribeWorkflowExecutionRequest) *history.DescribeWorkflowExecutionRequest {
	if in == nil {
		return nil
	}
	return &history.DescribeWorkflowExecutionRequest{
		DomainUUID: &in.DomainUUID,
		Request:    ToThriftDescribeWorkflowExecutionRequest(in.Request),
	}
}

// ToThriftReplicateEventsRequest ...
func ToThriftReplicateEventsRequest(in *historyservice.ReplicateEventsRequest) *history.ReplicateEventsRequest {
	if in == nil {
		return nil
	}
	return &history.ReplicateEventsRequest{
		SourceCluster:           &in.SourceCluster,
		DomainUUID:              &in.DomainUUID,
		WorkflowExecution:       ToThriftWorkflowExecution(in.WorkflowExecution),
		FirstEventId:            &in.FirstEventId,
		NextEventId:             &in.NextEventId,
		Version:                 &in.Version,
		ReplicationInfo:         ToThriftReplicationInfos(in.ReplicationInfo),
		History:                 ToThriftHistory(in.History),
		NewRunHistory:           ToThriftHistory(in.NewRunHistory),
		ForceBufferEvents:       &in.ForceBufferEvents,
		EventStoreVersion:       &in.EventStoreVersion,
		NewRunEventStoreVersion: &in.NewRunEventStoreVersion,
		ResetWorkflow:           &in.ResetWorkflow,
		NewRunNDC:               &in.NewRunNDC,
	}
}

// ToThriftReplicateRawEventsRequest ...
func ToThriftReplicateRawEventsRequest(in *historyservice.ReplicateRawEventsRequest) *history.ReplicateRawEventsRequest {
	if in == nil {
		return nil
	}
	return &history.ReplicateRawEventsRequest{
		DomainUUID:              &in.DomainUUID,
		WorkflowExecution:       ToThriftWorkflowExecution(in.WorkflowExecution),
		ReplicationInfo:         ToThriftReplicationInfos(in.ReplicationInfo),
		History:                 ToThriftDataBlob(in.History),
		NewRunHistory:           ToThriftDataBlob(in.NewRunHistory),
		EventStoreVersion:       &in.EventStoreVersion,
		NewRunEventStoreVersion: &in.NewRunEventStoreVersion,
	}
}

// ToThriftReplicateEventsV2Request ...
func ToThriftReplicateEventsV2Request(in *historyservice.ReplicateEventsV2Request) *history.ReplicateEventsV2Request {
	if in == nil {
		return nil
	}
	return &history.ReplicateEventsV2Request{
		DomainUUID:          &in.DomainUUID,
		WorkflowExecution:   ToThriftWorkflowExecution(in.WorkflowExecution),
		VersionHistoryItems: ToThriftVersionHistoryItems(in.VersionHistoryItems),
		Events:              ToThriftDataBlob(in.Events),
		NewRunEvents:        ToThriftDataBlob(in.NewRunEvents),
	}
}

// ToThriftSyncShardStatusRequest ...
func ToThriftSyncShardStatusRequest(in *historyservice.SyncShardStatusRequest) *history.SyncShardStatusRequest {
	if in == nil {
		return nil
	}
	return &history.SyncShardStatusRequest{
		SourceCluster: &in.SourceCluster,
		ShardId:       &in.ShardId,
		Timestamp:     &in.Timestamp,
	}
}

// ToThriftSyncActivityRequest ...
func ToThriftSyncActivityRequest(in *historyservice.SyncActivityRequest) *history.SyncActivityRequest {
	if in == nil {
		return nil
	}
	return &history.SyncActivityRequest{
		DomainId:           &in.DomainId,
		WorkflowId:         &in.WorkflowId,
		RunId:              &in.RunId,
		Version:            &in.Version,
		ScheduledId:        &in.ScheduledId,
		ScheduledTime:      &in.ScheduledTime,
		StartedId:          &in.StartedId,
		StartedTime:        &in.StartedTime,
		LastHeartbeatTime:  &in.LastHeartbeatTime,
		Details:            in.Details,
		Attempt:            &in.Attempt,
		LastFailureReason:  &in.LastFailureReason,
		LastWorkerIdentity: &in.LastWorkerIdentity,
		LastFailureDetails: in.LastFailureDetails,
		VersionHistory:     ToThriftVersionHistory(in.VersionHistory),
	}
}

// ToThriftDescribeMutableStateRequest ...
func ToThriftDescribeMutableStateRequest(in *historyservice.DescribeMutableStateRequest) *history.DescribeMutableStateRequest {
	if in == nil {
		return nil
	}
	return &history.DescribeMutableStateRequest{
		DomainUUID: &in.DomainUUID,
		Execution:  ToThriftWorkflowExecution(in.Execution),
	}
}

// ToThriftDescribeHistoryHostRequest ...
func ToThriftHistoryDescribeHistoryHostRequest(in *historyservice.DescribeHistoryHostRequest) *shared.DescribeHistoryHostRequest {
	if in == nil {
		return nil
	}
	return &shared.DescribeHistoryHostRequest{
		HostAddress:      &in.HostAddress,
		ShardIdForHost:   &in.ShardIdForHost,
		ExecutionForHost: ToThriftWorkflowExecution(in.ExecutionForHost),
	}
}

// ToThriftCloseShardRequest ...
func ToThriftHistoryCloseShardRequest(in *historyservice.CloseShardRequest) *shared.CloseShardRequest {
	if in == nil {
		return nil
	}
	return &shared.CloseShardRequest{
		ShardID: &in.ShardID,
	}
}

// ToThriftRemoveTaskRequest ...
func ToThriftHistoryRemoveTaskRequest(in *historyservice.RemoveTaskRequest) *shared.RemoveTaskRequest {
	if in == nil {
		return nil
	}
	return &shared.RemoveTaskRequest{
		ShardID: &in.ShardID,
		Type:    &in.Type,
		TaskID:  &in.TaskID,
	}
}

// ToThriftGetReplicationMessagesRequest ...
func ToThriftHistoryGetReplicationMessagesRequest(in *historyservice.GetReplicationMessagesRequest) *replicator.GetReplicationMessagesRequest {
	if in == nil {
		return nil
	}
	return &replicator.GetReplicationMessagesRequest{
		Tokens:      ToThriftReplicationTokens(in.Tokens),
		ClusterName: &in.ClusterName,
	}
}

// ToThriftGetDLQReplicationMessagesRequest ...
func ToThriftHistoryGetDLQReplicationMessagesRequest(in *historyservice.GetDLQReplicationMessagesRequest) *replicator.GetDLQReplicationMessagesRequest {
	if in == nil {
		return nil
	}
	return &replicator.GetDLQReplicationMessagesRequest{
		TaskInfos: ToThriftReplicationTaskInfos(in.TaskInfos),
	}
}

// ToThriftQueryWorkflowRequest ...
func ToThriftHistoryQueryWorkflowRequest(in *historyservice.QueryWorkflowRequest) *history.QueryWorkflowRequest {
	if in == nil {
		return nil
	}
	return &history.QueryWorkflowRequest{
		DomainUUID: &in.DomainUUID,
		Request:    ToThriftQueryWorkflowRequest(in.Request),
	}
}

// ToThriftReapplyEventsRequest ...
func ToThriftHistoryReapplyEventsRequest(in *historyservice.ReapplyEventsRequest) *history.ReapplyEventsRequest {
	if in == nil {
		return nil
	}
	return &history.ReapplyEventsRequest{
		DomainUUID: &in.DomainUUID,
		Request:    ToThriftReapplyEventsRequest(in.Request),
	}
}
