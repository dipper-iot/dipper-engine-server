// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// RuleChanColumns holds the columns for the "rule_chan" table.
	RuleChanColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "name", Type: field.TypeString, Size: 256},
		{Name: "description", Type: field.TypeString, Size: 1000},
		{Name: "root_node", Type: field.TypeString, Size: 256},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"activated", "deactivated"}},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// RuleChanTable holds the schema information for the "rule_chan" table.
	RuleChanTable = &schema.Table{
		Name:       "rule_chan",
		Columns:    RuleChanColumns,
		PrimaryKey: []*schema.Column{RuleChanColumns[0]},
	}
	// RuleNodesColumns holds the columns for the "rule_nodes" table.
	RuleNodesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "node_id", Type: field.TypeString, Size: 20},
		{Name: "rule_id", Type: field.TypeString, Size: 100},
		{Name: "option", Type: field.TypeJSON},
		{Name: "debug", Type: field.TypeBool, Default: false},
		{Name: "end", Type: field.TypeBool, Default: false},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "chain_id", Type: field.TypeUint64, Nullable: true},
	}
	// RuleNodesTable holds the schema information for the "rule_nodes" table.
	RuleNodesTable = &schema.Table{
		Name:       "rule_nodes",
		Columns:    RuleNodesColumns,
		PrimaryKey: []*schema.Column{RuleNodesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "rule_nodes_rule_chan_rules",
				Columns:    []*schema.Column{RuleNodesColumns[8]},
				RefColumns: []*schema.Column{RuleChanColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// SessionsColumns holds the columns for the "sessions" table.
	SessionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "is_test", Type: field.TypeBool, Default: false},
		{Name: "infinite", Type: field.TypeBool, Default: false},
		{Name: "data", Type: field.TypeJSON},
		{Name: "result", Type: field.TypeJSON},
		{Name: "end_count", Type: field.TypeInt, Default: 0},
		{Name: "timeout", Type: field.TypeInt, Default: 30},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "chain_id", Type: field.TypeUint64, Nullable: true},
	}
	// SessionsTable holds the schema information for the "sessions" table.
	SessionsTable = &schema.Table{
		Name:       "sessions",
		Columns:    SessionsColumns,
		PrimaryKey: []*schema.Column{SessionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "sessions_rule_chan_sessions",
				Columns:    []*schema.Column{SessionsColumns[9]},
				RefColumns: []*schema.Column{RuleChanColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		RuleChanTable,
		RuleNodesTable,
		SessionsTable,
	}
)

func init() {
	RuleChanTable.Annotation = &entsql.Annotation{
		Table: "rule_chan",
	}
	RuleNodesTable.ForeignKeys[0].RefTable = RuleChanTable
	RuleNodesTable.Annotation = &entsql.Annotation{
		Table: "rule_nodes",
	}
	SessionsTable.ForeignKeys[0].RefTable = RuleChanTable
	SessionsTable.Annotation = &entsql.Annotation{
		Table: "sessions",
	}
}
