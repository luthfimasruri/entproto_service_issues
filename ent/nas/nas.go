// Code generated by entc, DO NOT EDIT.

package nas

const (
	// Label holds the string label denoting the nas type in the database.
	Label = "nas"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldNasname holds the string denoting the nasname field in the database.
	FieldNasname = "nasname"
	// Table holds the table name of the nas in the database.
	Table = "nas"
)

// Columns holds all SQL columns for nas fields.
var Columns = []string{
	FieldID,
	FieldNasname,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}