package user

import (
	"context"
	"encoding/json"
	"github.com/caos/zitadel/internal/errors"
	"github.com/caos/zitadel/internal/eventstore/v2"
	"github.com/caos/zitadel/internal/eventstore/v2/repository"
)

const (
	externalIDPEventPrefix   = humanEventPrefix + "externalidp."
	externalLoginEventPrefix = humanEventPrefix + "externallogin."

	//TODO: Handle unique Aggregate
	HumanExternalIDPReservedType = externalIDPEventPrefix + "reserved"
	HumanExternalIDPReleasedType = externalIDPEventPrefix + "released"

	HumanExternalIDPAddedType          = externalIDPEventPrefix + "added"
	HumanExternalIDPRemovedType        = externalIDPEventPrefix + "removed"
	HumanExternalIDPCascadeRemovedType = externalIDPEventPrefix + "cascade.removed"

	HumanExternalLoginCheckSucceededType = externalLoginEventPrefix + "check.succeeded"
)

type HumanExternalIDPReservedEvent struct {
	eventstore.BaseEvent `json:"-"`
}

func (e *HumanExternalIDPReservedEvent) Data() interface{} {
	return nil
}

func NewHumanExternalIDPReservedEvent(ctx context.Context) *HumanExternalIDPReservedEvent {
	return &HumanExternalIDPReservedEvent{
		BaseEvent: *eventstore.NewBaseEventForPush(
			ctx,
			HumanExternalIDPReservedType,
		),
	}
}

type HumanExternalIDPReleasedEvent struct {
	eventstore.BaseEvent `json:"-"`
}

func (e *HumanExternalIDPReleasedEvent) Data() interface{} {
	return nil
}

func NewHumanExternalIDPReleasedEvent(ctx context.Context) *HumanExternalIDPReleasedEvent {
	return &HumanExternalIDPReleasedEvent{
		BaseEvent: *eventstore.NewBaseEventForPush(
			ctx,
			HumanExternalIDPReleasedType,
		),
	}
}

type HumanExternalIDPAddedEvent struct {
	eventstore.BaseEvent `json:"-"`

	IDPConfigID string `json:"idpConfigId,omitempty"`
	UserID      string `json:"userId,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
}

func (e *HumanExternalIDPAddedEvent) Data() interface{} {
	return e
}

func NewHumanExternalIDPAddedEvent(ctx context.Context, idpConfigID, displayName string) *HumanExternalIDPAddedEvent {
	return &HumanExternalIDPAddedEvent{
		BaseEvent: *eventstore.NewBaseEventForPush(
			ctx,
			HumanExternalIDPAddedType,
		),
		IDPConfigID: idpConfigID,
		DisplayName: displayName,
	}
}

func HumanExternalIDPAddedEventMapper(event *repository.Event) (eventstore.EventReader, error) {
	e := &HumanExternalIDPAddedEvent{
		BaseEvent: *eventstore.BaseEventFromRepo(event),
	}

	err := json.Unmarshal(event.Data, e)
	if err != nil {
		return nil, errors.ThrowInternal(err, "USER-6M9sd", "unable to unmarshal user external idp added")
	}

	return e, nil
}

type HumanExternalIDPRemovedEvent struct {
	eventstore.BaseEvent `json:"-"`

	IDPConfigID string `json:"idpConfigId"`
	UserID      string `json:"userId,omitempty"`
}

func (e *HumanExternalIDPRemovedEvent) Data() interface{} {
	return e
}

func NewHumanExternalIDPRemovedEvent(ctx context.Context, idpConfigID, externalUserID string) *HumanExternalIDPRemovedEvent {
	return &HumanExternalIDPRemovedEvent{
		BaseEvent: *eventstore.NewBaseEventForPush(
			ctx,
			HumanExternalIDPRemovedType,
		),
		IDPConfigID: idpConfigID,
		UserID:      externalUserID,
	}
}

func HumanExternalIDPRemovedEventMapper(event *repository.Event) (eventstore.EventReader, error) {
	e := &HumanExternalIDPRemovedEvent{
		BaseEvent: *eventstore.BaseEventFromRepo(event),
	}

	err := json.Unmarshal(event.Data, e)
	if err != nil {
		return nil, errors.ThrowInternal(err, "USER-5Gm9s", "unable to unmarshal user external idp removed")
	}

	return e, nil
}

type HumanExternalIDPCascadeRemovedEvent struct {
	eventstore.BaseEvent `json:"-"`

	IDPConfigID string `json:"idpConfigId"`
	UserID      string `json:"userId,omitempty"`
}

func (e *HumanExternalIDPCascadeRemovedEvent) Data() interface{} {
	return e
}

func NewHumanExternalIDPCascadeRemovedEvent(ctx context.Context, idpConfigID, externalUserID string) *HumanExternalIDPCascadeRemovedEvent {
	return &HumanExternalIDPCascadeRemovedEvent{
		BaseEvent: *eventstore.NewBaseEventForPush(
			ctx,
			HumanExternalIDPCascadeRemovedType,
		),
		IDPConfigID: idpConfigID,
		UserID:      externalUserID,
	}
}

func HumanExternalIDPCascadeRemovedEventMapper(event *repository.Event) (eventstore.EventReader, error) {
	e := &HumanExternalIDPCascadeRemovedEvent{
		BaseEvent: *eventstore.BaseEventFromRepo(event),
	}

	err := json.Unmarshal(event.Data, e)
	if err != nil {
		return nil, errors.ThrowInternal(err, "USER-2M0sd", "unable to unmarshal user external idp cascade removed")
	}

	return e, nil
}

type HumanExternalIDPCheckSucceededEvent struct {
	eventstore.BaseEvent `json:"-"`
}

func (e *HumanExternalIDPCheckSucceededEvent) Data() interface{} {
	return nil
}

func NewHumanExternalIDPCheckSucceededEvent(ctx context.Context) *HumanExternalIDPCheckSucceededEvent {
	return &HumanExternalIDPCheckSucceededEvent{
		BaseEvent: *eventstore.NewBaseEventForPush(
			ctx,
			HumanExternalLoginCheckSucceededType,
		),
	}
}

func HumanExternalIDPCheckSucceededEventMapper(event *repository.Event) (eventstore.EventReader, error) {
	return &HumanExternalIDPCheckSucceededEvent{
		BaseEvent: *eventstore.BaseEventFromRepo(event),
	}, nil
}