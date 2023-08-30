package schema

import (
	"database/sql"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"github.com/shopspring/decimal"
)

// Amount is a custom Go type that's convertible to the basic float64 type.
type Amount float64

// Card schema.
type Card struct {
	ent.Schema
}

// Fields of the Card.
func (Card) Fields() []ent.Field {
	return []ent.Field{
		field.Float("amount").
			// 自定义类型
			GoType(Amount(0)),
		field.String("name").
			Optional().
			// A ValueScanner type.   scan value均需实现
			GoType(&sql.NullString{}),
		field.Float("decimal").
			// A ValueScanner type mixed with SchemaType.
			GoType(decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL:    "decimal(6,2)",
				dialect.Postgres: "numeric",
			}), // 针对不同数据库不同类型
	}
}
