// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

func (pr *Project) User(ctx context.Context) (*User, error) {
	result, err := pr.Edges.UserOrErr()
	if IsNotLoaded(err) {
		result, err = pr.QueryUser().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (pr *Project) Tasks(ctx context.Context) (result []*Task, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = pr.NamedTasks(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = pr.Edges.TasksOrErr()
	}
	if IsNotLoaded(err) {
		result, err = pr.QueryTasks().All(ctx)
	}
	return result, err
}

func (t *Task) Project(ctx context.Context) (*Project, error) {
	result, err := t.Edges.ProjectOrErr()
	if IsNotLoaded(err) {
		result, err = t.QueryProject().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (t *Task) AssignedTo(ctx context.Context) (result []*User, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = t.NamedAssignedTo(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = t.Edges.AssignedToOrErr()
	}
	if IsNotLoaded(err) {
		result, err = t.QueryAssignedTo().All(ctx)
	}
	return result, err
}

func (u *User) Projects(ctx context.Context) (result []*Project, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = u.NamedProjects(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = u.Edges.ProjectsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = u.QueryProjects().All(ctx)
	}
	return result, err
}

func (u *User) AssignedTasks(ctx context.Context) (result []*Task, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = u.NamedAssignedTasks(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = u.Edges.AssignedTasksOrErr()
	}
	if IsNotLoaded(err) {
		result, err = u.QueryAssignedTasks().All(ctx)
	}
	return result, err
}