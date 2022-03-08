package schema

import (
	"database/sql/driver"
	"fmt"
	"net"

	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type Nas struct {
	ent.Schema
}

func (Nas) Fields() []ent.Field {
	return []ent.Field{
		field.String("nasname").Annotations(entproto.Field(2)),
		field.String("type").Annotations(entproto.Field(4)),
		field.String("nasipaddress").
			GoType(&Inet{}).
			SchemaType(map[string]string{
				dialect.Postgres: "inet",
			}).
			Validate(func(s string) error {
				if net.ParseIP(s) == nil {
					return fmt.Errorf("invalid value for ip %q", s)
				}
				return nil
			}).Annotations(entproto.Field(25)),
	}
}

func (Nas) Edges() []ent.Edge {
	return nil
}

func (Nas) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(),
	}
}

// Inet represents a single IP address
type Inet struct {
	net.IP
}

// Scan implements the Scanner interface
func (i *Inet) Scan(value interface{}) (err error) {
	switch v := value.(type) {
	case nil:
	case []byte:
		if i.IP = net.ParseIP(string(v)); i.IP == nil {
			err = fmt.Errorf("invalid value for ip %q", v)
		}
	case string:
		if i.IP = net.ParseIP(v); i.IP == nil {
			err = fmt.Errorf("invalid value for ip %q", v)
		}
	default:
		err = fmt.Errorf("unexpected type %T", v)
	}
	return
}

// Value implements the driver Valuer interface
func (i Inet) Value() (driver.Value, error) {
	return i.IP.String(), nil
}
