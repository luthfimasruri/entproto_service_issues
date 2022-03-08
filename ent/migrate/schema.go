// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// NasColumns holds the columns for the "nas" table.
	NasColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "nasname", Type: field.TypeString},
	}
	// NasTable holds the schema information for the "nas" table.
	NasTable = &schema.Table{
		Name:       "nas",
		Columns:    NasColumns,
		PrimaryKey: []*schema.Column{NasColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		NasTable,
	}
)

func init() {
}
