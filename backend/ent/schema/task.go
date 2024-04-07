package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Task struct {
	ent.Schema
}

func (Task) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Immutable(),
		field.String("title"),
		field.String("description"),
		field.Time("due_date"),
		field.Enum("status").Values("TODO", "IN_PROGRESS", "DONE"),
		field.Time("created_at"),
		field.Time("updated_at"),
	}
}

func (Task) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("project", Project.Type).Ref("tasks").Unique(),
		edge.From("assigned_to", User.Type).Ref("assigned_tasks"),
	}
}

func (Task) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
