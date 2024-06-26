package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.45

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/D-Wtnbe/todo-app-graphql-golang/ent"
	"github.com/D-Wtnbe/todo-app-graphql-golang/resolver/model"
)

// Node is the resolver for the node field.
func (r *queryResolver) Node(ctx context.Context, id string) (ent.Noder, error) {
	panic(fmt.Errorf("not implemented: Node - node"))
}

// Nodes is the resolver for the nodes field.
func (r *queryResolver) Nodes(ctx context.Context, ids []string) ([]ent.Noder, error) {
	panic(fmt.Errorf("not implemented: Nodes - nodes"))
}

// Projects is the resolver for the projects field.
func (r *queryResolver) Projects(ctx context.Context) ([]*model.Project, error) {
	entProjects, err := r.client.Project.Query().All(ctx)
	if err != nil {
		log.Fatalf("failed querying projects: %v", err)
	}
	modelProjects := make([]*model.Project, len(entProjects))
	for i, entProject := range entProjects {
		modelProjects[i] = &model.Project{
			ID:          strconv.Itoa(entProject.ID),
			Name:        entProject.Name,
			Description: entProject.Description,
			CreatedAt:   entProject.CreatedAt,
			UpdatedAt:   entProject.UpdatedAt,
			User: &model.User{
				ID: strconv.Itoa(entProject.Edges.User.ID),
			},
			Tasks: []*model.Task{
				{
					ID: strconv.Itoa(entProject.Edges.Tasks[0].ID),
				},
			},
		}
	}
	return modelProjects, nil
}

// Tasks is the resolver for the tasks field.
func (r *queryResolver) Tasks(ctx context.Context) ([]*model.Task, error) {
	entTasks, err := r.client.Task.Query().All(ctx)
	if err != nil {
		log.Fatalf("failed querying tasks: %v", err)
	}
	modelTasks := make([]*model.Task, len(entTasks))
	for i, entTask := range entTasks {
		modelTasks[i] = &model.Task{
			ID:          strconv.Itoa(entTask.ID),
			Title:       entTask.Title,
			Description: entTask.Description,
			DueDate:     entTask.DueDate,
			Status:      entTask.Status,
			CreatedAt:   entTask.CreatedAt,
			UpdatedAt:   entTask.UpdatedAt,
			Project: &model.Project{
				ID: strconv.Itoa(entTask.Edges.Project.ID),
			},
			AssignedTo: []*model.User{
				{
					ID: strconv.Itoa(entTask.Edges.AssignedTo[i].ID),
				},
			},
		}
	}
	return modelTasks, nil
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	entTodos, err := r.client.Todo.Query().All(ctx)
	if err != nil {
		log.Fatalf("failed querying todos: %v", err)
	}
	modelTodos := make([]*model.Todo, len(entTodos))
	for i, entTodo := range entTodos {
		modelTodos[i] = &model.Todo{
			ID:    strconv.Itoa(entTodo.ID),
			Name:  entTodo.Name,
			Email: entTodo.Email,
			Done:  entTodo.Done,
		}
	}
	return modelTodos, nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	entUsers, err := r.client.User.Query().All(ctx)
	if err != nil {
		log.Fatalf("failed querying users: %v", err)
	}
	modelUsers := make([]*model.User, len(entUsers))
	for i, entUser := range entUsers {
		modelUsers[i] = &model.User{
			ID:           strconv.Itoa(entUser.ID),
			Name:         entUser.Name,
			Email:        entUser.Email,
			PasswordHash: entUser.PasswordHash,
			Projects: []*model.Project{
				{
					ID: strconv.Itoa(entUser.Edges.Projects[i].ID),
				},
			},
			AssignedTasks: []*model.Task{
				{
					ID: strconv.Itoa(entUser.Edges.AssignedTasks[i].ID),
				},
			},
		}
	}
	return modelUsers, nil
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
