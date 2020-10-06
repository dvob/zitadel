package sql

import (
	"testing"

	"github.com/caos/zitadel/internal/eventstore/v2/repository"
	_ "github.com/lib/pq"
)

func TestCRDB_placeholder(t *testing.T) {
	type args struct {
		query string
	}
	type res struct {
		query string
	}
	tests := []struct {
		name string
		args args
		res  res
	}{
		{
			name: "no placeholders",
			args: args{
				query: "SELECT * FROM eventstore.events",
			},
			res: res{
				query: "SELECT * FROM eventstore.events",
			},
		},
		{
			name: "one placeholder",
			args: args{
				query: "SELECT * FROM eventstore.events WHERE aggregate_type = ?",
			},
			res: res{
				query: "SELECT * FROM eventstore.events WHERE aggregate_type = $1",
			},
		},
		{
			name: "multiple placeholders",
			args: args{
				query: "SELECT * FROM eventstore.events WHERE aggregate_type = ? AND aggregate_id = ? LIMIT ?",
			},
			res: res{
				query: "SELECT * FROM eventstore.events WHERE aggregate_type = $1 AND aggregate_id = $2 LIMIT $3",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &CRDB{}
			if query := db.placeholder(tt.args.query); query != tt.res.query {
				t.Errorf("CRDB.placeholder() = %v, want %v", query, tt.res.query)
			}
		})
	}
}

func TestCRDB_operation(t *testing.T) {
	type res struct {
		op string
	}
	type args struct {
		operation repository.Operation
	}
	tests := []struct {
		name string
		args args
		res  res
	}{
		{
			name: "no op",
			args: args{
				operation: repository.Operation(-1),
			},
			res: res{
				op: "",
			},
		},
		{
			name: "greater",
			args: args{
				operation: repository.Operation_Greater,
			},
			res: res{
				op: ">",
			},
		},
		{
			name: "less",
			args: args{
				operation: repository.Operation_Less,
			},
			res: res{
				op: "<",
			},
		},
		{
			name: "equals",
			args: args{
				operation: repository.Operation_Equals,
			},
			res: res{
				op: "=",
			},
		},
		{
			name: "in",
			args: args{
				operation: repository.Operation_In,
			},
			res: res{
				op: "=",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &CRDB{}
			if got := db.operation(tt.args.operation); got != tt.res.op {
				t.Errorf("CRDB.operation() = %v, want %v", got, tt.res.op)
			}
		})
	}
}

func TestCRDB_conditionFormat(t *testing.T) {
	type res struct {
		format string
	}
	type args struct {
		operation repository.Operation
	}
	tests := []struct {
		name string
		args args
		res  res
	}{
		{
			name: "default",
			args: args{
				operation: repository.Operation_Equals,
			},
			res: res{
				format: "%s %s ?",
			},
		},
		{
			name: "in",
			args: args{
				operation: repository.Operation_In,
			},
			res: res{
				format: "%s %s ANY(?)",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &CRDB{}
			if got := db.conditionFormat(tt.args.operation); got != tt.res.format {
				t.Errorf("CRDB.conditionFormat() = %v, want %v", got, tt.res.format)
			}
		})
	}
}

func TestCRDB_columnName(t *testing.T) {
	type res struct {
		name string
	}
	type args struct {
		field repository.Field
	}
	tests := []struct {
		name string
		args args
		res  res
	}{
		{
			name: "invalid field",
			args: args{
				field: repository.Field(-1),
			},
			res: res{
				name: "",
			},
		},
		{
			name: "aggregate id",
			args: args{
				field: repository.Field_AggregateID,
			},
			res: res{
				name: "aggregate_id",
			},
		},
		{
			name: "aggregate type",
			args: args{
				field: repository.Field_AggregateType,
			},
			res: res{
				name: "aggregate_type",
			},
		},
		{
			name: "editor service",
			args: args{
				field: repository.Field_EditorService,
			},
			res: res{
				name: "editor_service",
			},
		},
		{
			name: "editor user",
			args: args{
				field: repository.Field_EditorUser,
			},
			res: res{
				name: "editor_user",
			},
		},
		{
			name: "event type",
			args: args{
				field: repository.Field_EventType,
			},
			res: res{
				name: "event_type",
			},
		},
		{
			name: "latest sequence",
			args: args{
				field: repository.Field_LatestSequence,
			},
			res: res{
				name: "event_sequence",
			},
		},
		{
			name: "resource owner",
			args: args{
				field: repository.Field_ResourceOwner,
			},
			res: res{
				name: "resource_owner",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &CRDB{}
			if got := db.columnName(tt.args.field); got != tt.res.name {
				t.Errorf("CRDB.operation() = %v, want %v", got, tt.res.name)
			}
		})
	}
}