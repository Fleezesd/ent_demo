package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("age").
			Positive(),
		field.String("name").
			StructTag(`gqlgen:"gql_name"`),
		field.String("username").
			Unique().
			Immutable().
			StorageKey("user-name"), // 设置列名
		field.String("password").
			Sensitive(), // Sensitive 敏感字段
		field.Enum("size").
			Values("big", "small"),
		field.Time("created_at").
			Default(time.Now).
			Annotations(
				entsql.Default("CURRENT_TIMESTAMP"),
			),
		// 该字段设置的有些问题 后续修改
		// field.String("field").
		// 	Optional().
		// 	Annotations(
		// 		entsql.DefaultExpr("lower(username)"),
		// 	),
		field.String("default_exprs").
			Optional().
			Annotations(
				entsql.DefaultExprs(map[string]string{
					dialect.MySQL:    "TO_BASE64('ent')", // base64位
					dialect.SQLite:   "hex('ent')",       // hex 16进制编码
					dialect.Postgres: "md5('ent)",        // md5 加密
				},
				),
			),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	// edge (Relation)  edge.To
	return []ent.Edge{
		edge.To("cars", Car.Type),
		// Create an inverse-edge called "groups" of type `Group`
		// and reference it to the "users" edge (in Group schema)
		// explicitly using the `Ref` method.
		edge.From("groups", Group.Type).
			Ref("users"),
	}
}
