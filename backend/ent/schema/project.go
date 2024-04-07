package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Project struct {
	ent.Schema
}

func (Project) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Immutable(),
		field.String("name"),
		field.String("description"),
		field.Time("created_at"),
		field.Time("updated_at"),
	}
}

func (Project) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("projects").Unique(),
		edge.To("tasks", Task.Type),
	}
}

func (Project) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
