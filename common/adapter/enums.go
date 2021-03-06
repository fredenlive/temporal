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
	"go.temporal.io/temporal-proto/enums"

	"github.com/temporalio/temporal/.gen/go/replicator"
	"github.com/temporalio/temporal/.gen/go/shared"
)

// ToThriftArchivalStatus ...
func ToThriftArchivalStatus(in enums.ArchivalStatus) *shared.ArchivalStatus {
	if in == enums.ArchivalStatusDefault {
		return nil
	}

	ret := shared.ArchivalStatus(in - 1)
	return &ret
}

// ToProtoArchivalStatus ...
func ToProtoArchivalStatus(in *shared.ArchivalStatus) enums.ArchivalStatus {
	if in == nil {
		return enums.ArchivalStatusDefault
	}
	return enums.ArchivalStatus(*in + 1)
}

func ToProtoWorkflowExecutionCloseStatus(in *shared.WorkflowExecutionCloseStatus) enums.WorkflowExecutionCloseStatus {
	if in == nil {
		return enums.WorkflowExecutionCloseStatusRunning
	}
	return enums.WorkflowExecutionCloseStatus(*in + 1)
}

func ToThriftTaskListKind(in enums.TaskListKind) *shared.TaskListKind {
	ret := shared.TaskListKind(in)
	return &ret
}

func ToThriftWorkflowIDReusePolicy(in enums.WorkflowIdReusePolicy) *shared.WorkflowIdReusePolicy {
	ret := shared.WorkflowIdReusePolicy(in)
	return &ret
}
func ToThriftHistoryEventFilterType(in enums.HistoryEventFilterType) *shared.HistoryEventFilterType {
	ret := shared.HistoryEventFilterType(in)
	return &ret
}

func ToThriftWorkflowExecutionCloseStatus(in enums.WorkflowExecutionCloseStatus) *shared.WorkflowExecutionCloseStatus {
	if in == enums.WorkflowExecutionCloseStatusRunning {
		return nil
	}

	ret := shared.WorkflowExecutionCloseStatus(in - 1)
	return &ret
}

func ToThriftDecisionType(in enums.DecisionType) *shared.DecisionType {
	ret := shared.DecisionType(in)
	return &ret
}

func ToThriftContinueAsNewInitiator(in enums.ContinueAsNewInitiator) *shared.ContinueAsNewInitiator {
	if in == enums.ContinueAsNewInitiatorNotSet {
		return nil
	}
	ret := shared.ContinueAsNewInitiator(in - 1)
	return &ret
}

// ToProtoContinueAsNewInitiator ...
func ToProtoContinueAsNewInitiator(in *shared.ContinueAsNewInitiator) enums.ContinueAsNewInitiator {
	if in == nil {
		return enums.ContinueAsNewInitiatorNotSet
	}
	return enums.ContinueAsNewInitiator(*in + 1)
}

func ToThriftParentClosePolicy(in enums.ParentClosePolicy) *shared.ParentClosePolicy {
	ret := shared.ParentClosePolicy(in)
	return &ret
}

func ToThriftQueryResultType(in enums.QueryResultType) *shared.QueryResultType {
	ret := shared.QueryResultType(in)
	return &ret
}
func ToThriftDecisionTaskFailedCause(in enums.DecisionTaskFailedCause) *shared.DecisionTaskFailedCause {
	ret := shared.DecisionTaskFailedCause(in)
	return &ret
}

func ToThriftQueryTaskCompletedType(in enums.QueryTaskCompletedType) *shared.QueryTaskCompletedType {
	ret := shared.QueryTaskCompletedType(in)
	return &ret
}

func ToThriftQueryRejectCondition(in enums.QueryRejectCondition) *shared.QueryRejectCondition {
	if in == enums.QueryRejectConditionNone {
		return nil
	}

	ret := shared.QueryRejectCondition(in - 1)
	return &ret
}

func ToThriftQueryConsistencyLevel(in enums.QueryConsistencyLevel) *shared.QueryConsistencyLevel {
	ret := shared.QueryConsistencyLevel(in)
	return &ret
}

func ToThriftTaskListType(in enums.TaskListType) *shared.TaskListType {
	ret := shared.TaskListType(in)
	return &ret
}

func ToThriftEncodingType(in enums.EncodingType) *shared.EncodingType {
	switch in {
	case enums.EncodingTypeThriftRW:
		ret := shared.EncodingTypeThriftRW
		return &ret
	case enums.EncodingTypeJSON:
		ret := shared.EncodingTypeJSON
		return &ret
	case enums.EncodingTypeProto:
		panic("EncodingTypeProto is not supported")
	}

	ret := shared.EncodingTypeThriftRW
	return &ret
}
func ToProtoEncodingType(in shared.EncodingType) enums.EncodingType {
	switch in {
	case shared.EncodingTypeThriftRW:
		return enums.EncodingTypeThriftRW
	case shared.EncodingTypeJSON:
		return enums.EncodingTypeJSON
	}

	return enums.EncodingTypeThriftRW
}
func ToThriftEventType(in enums.EventType) *shared.EventType {
	ret := shared.EventType(in)
	return &ret
}

func ToThriftTimeoutType(in enums.TimeoutType) *shared.TimeoutType {
	ret := shared.TimeoutType(in)
	return &ret
}

func ToThriftCancelExternalWorkflowExecutionFailedCause(in enums.CancelExternalWorkflowExecutionFailedCause) *shared.CancelExternalWorkflowExecutionFailedCause {
	ret := shared.CancelExternalWorkflowExecutionFailedCause(in)
	return &ret
}

func ToThriftSignalExternalWorkflowExecutionFailedCause(in enums.SignalExternalWorkflowExecutionFailedCause) *shared.SignalExternalWorkflowExecutionFailedCause {
	ret := shared.SignalExternalWorkflowExecutionFailedCause(in)
	return &ret
}

func ToThriftChildWorkflowExecutionFailedCause(in enums.ChildWorkflowExecutionFailedCause) *shared.ChildWorkflowExecutionFailedCause {
	ret := shared.ChildWorkflowExecutionFailedCause(in)
	return &ret
}

func ToThriftDomainOperation(in enums.DomainOperation) *replicator.DomainOperation {
	ret := replicator.DomainOperation(in)
	return &ret
}
func ToThriftDomainStatus(in enums.DomainStatus) *shared.DomainStatus {
	ret := shared.DomainStatus(in)
	return &ret
}
func ToThriftReplicationTaskType(in enums.ReplicationTaskType) *replicator.ReplicationTaskType {
	ret := replicator.ReplicationTaskType(in)
	return &ret
}
