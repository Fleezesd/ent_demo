package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Car struct {
	ent.Schema
}

// Fields of the Car.
func (Car) Fields() []ent.Field {
	return []ent.Field{
		field.String("model"),
		field.Time("registered_at"),
	}
}

// Edges of the Car.
func (Car) Edges() []ent.Edge {
	return []ent.Edge{
		// Create an inverse-edge called "owner" of type `User`
		// and reference it to the "cars" edge (in User schema)
		// explicitly using the `Ref` method.

		// edge (BackRef)	edge.From
		edge.From("owner", User.Type).
			// Ref sets the referenced-edge of this inverse edge. ref设置逆向边的参考边
			Ref("cars").
			// setting the edge to unique, ensure
			// that a car can have only one owner.
			Unique(),
	}
}
