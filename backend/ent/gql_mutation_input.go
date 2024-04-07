// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/D-Wtnbe/todo-app-graphql-golang/ent/task"
)

// CreateProjectInput represents a mutation input for creating projects.
type CreateProjectInput struct {
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	UserID      *int
	TaskIDs     []int
}

// Mutate applies the CreateProjectInput on the ProjectMutation builder.
func (i *CreateProjectInput) Mutate(m *ProjectMutation) {
	m.SetName(i.Name)
	m.SetDescription(i.Description)
	m.SetCreatedAt(i.CreatedAt)
	m.SetUpdatedAt(i.UpdatedAt)
	if v := i.UserID; v != nil {
		m.SetUserID(*v)
	}
	if v := i.TaskIDs; len(v) > 0 {
		m.AddTaskIDs(v...)
	}
}

// SetInput applies the change-set in the CreateProjectInput on the ProjectCreate builder.
func (c *ProjectCreate) SetInput(i CreateProjectInput) *ProjectCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateProjectInput represents a mutation input for updating projects.
type UpdateProjectInput struct {
	Name          *string
	Description   *string
	CreatedAt     *time.Time
	UpdatedAt     *time.Time
	ClearUser     bool
	UserID        *int
	ClearTasks    bool
	AddTaskIDs    []int
	RemoveTaskIDs []int
}

// Mutate applies the UpdateProjectInput on the ProjectMutation builder.
func (i *UpdateProjectInput) Mutate(m *ProjectMutation) {
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
	if v := i.Description; v != nil {
		m.SetDescription(*v)
	}
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if i.ClearUser {
		m.ClearUser()
	}
	if v := i.UserID; v != nil {
		m.SetUserID(*v)
	}
	if i.ClearTasks {
		m.ClearTasks()
	}
	if v := i.AddTaskIDs; len(v) > 0 {
		m.AddTaskIDs(v...)
	}
	if v := i.RemoveTaskIDs; len(v) > 0 {
		m.RemoveTaskIDs(v...)
	}
}

// SetInput applies the change-set in the UpdateProjectInput on the ProjectUpdate builder.
func (c *ProjectUpdate) SetInput(i UpdateProjectInput) *ProjectUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateProjectInput on the ProjectUpdateOne builder.
func (c *ProjectUpdateOne) SetInput(i UpdateProjectInput) *ProjectUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateTaskInput represents a mutation input for creating tasks.
type CreateTaskInput struct {
	Title         string
	Description   string
	DueDate       time.Time
	Status        task.Status
	CreatedAt     time.Time
	UpdatedAt     time.Time
	ProjectID     *int
	AssignedToIDs []int
}

// Mutate applies the CreateTaskInput on the TaskMutation builder.
func (i *CreateTaskInput) Mutate(m *TaskMutation) {
	m.SetTitle(i.Title)
	m.SetDescription(i.Description)
	m.SetDueDate(i.DueDate)
	m.SetStatus(i.Status)
	m.SetCreatedAt(i.CreatedAt)
	m.SetUpdatedAt(i.UpdatedAt)
	if v := i.ProjectID; v != nil {
		m.SetProjectID(*v)
	}
	if v := i.AssignedToIDs; len(v) > 0 {
		m.AddAssignedToIDs(v...)
	}
}

// SetInput applies the change-set in the CreateTaskInput on the TaskCreate builder.
func (c *TaskCreate) SetInput(i CreateTaskInput) *TaskCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateTaskInput represents a mutation input for updating tasks.
type UpdateTaskInput struct {
	Title               *string
	Description         *string
	DueDate             *time.Time
	Status              *task.Status
	CreatedAt           *time.Time
	UpdatedAt           *time.Time
	ClearProject        bool
	ProjectID           *int
	ClearAssignedTo     bool
	AddAssignedToIDs    []int
	RemoveAssignedToIDs []int
}

// Mutate applies the UpdateTaskInput on the TaskMutation builder.
func (i *UpdateTaskInput) Mutate(m *TaskMutation) {
	if v := i.Title; v != nil {
		m.SetTitle(*v)
	}
	if v := i.Description; v != nil {
		m.SetDescription(*v)
	}
	if v := i.DueDate; v != nil {
		m.SetDueDate(*v)
	}
	if v := i.Status; v != nil {
		m.SetStatus(*v)
	}
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if i.ClearProject {
		m.ClearProject()
	}
	if v := i.ProjectID; v != nil {
		m.SetProjectID(*v)
	}
	if i.ClearAssignedTo {
		m.ClearAssignedTo()
	}
	if v := i.AddAssignedToIDs; len(v) > 0 {
		m.AddAssignedToIDs(v...)
	}
	if v := i.RemoveAssignedToIDs; len(v) > 0 {
		m.RemoveAssignedToIDs(v...)
	}
}

// SetInput applies the change-set in the UpdateTaskInput on the TaskUpdate builder.
func (c *TaskUpdate) SetInput(i UpdateTaskInput) *TaskUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateTaskInput on the TaskUpdateOne builder.
func (c *TaskUpdateOne) SetInput(i UpdateTaskInput) *TaskUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateTodoInput represents a mutation input for creating todos.
type CreateTodoInput struct {
	Name  string
	Email string
	Done  *bool
}

// Mutate applies the CreateTodoInput on the TodoMutation builder.
func (i *CreateTodoInput) Mutate(m *TodoMutation) {
	m.SetName(i.Name)
	m.SetEmail(i.Email)
	if v := i.Done; v != nil {
		m.SetDone(*v)
	}
}

// SetInput applies the change-set in the CreateTodoInput on the TodoCreate builder.
func (c *TodoCreate) SetInput(i CreateTodoInput) *TodoCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateTodoInput represents a mutation input for updating todos.
type UpdateTodoInput struct {
	Name  *string
	Email *string
	Done  *bool
}

// Mutate applies the UpdateTodoInput on the TodoMutation builder.
func (i *UpdateTodoInput) Mutate(m *TodoMutation) {
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
	if v := i.Email; v != nil {
		m.SetEmail(*v)
	}
	if v := i.Done; v != nil {
		m.SetDone(*v)
	}
}

// SetInput applies the change-set in the UpdateTodoInput on the TodoUpdate builder.
func (c *TodoUpdate) SetInput(i UpdateTodoInput) *TodoUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateTodoInput on the TodoUpdateOne builder.
func (c *TodoUpdateOne) SetInput(i UpdateTodoInput) *TodoUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateUserInput represents a mutation input for creating users.
type CreateUserInput struct {
	Name            string
	Email           string
	PasswordHash    string
	ProjectIDs      []int
	AssignedTaskIDs []int
}

// Mutate applies the CreateUserInput on the UserMutation builder.
func (i *CreateUserInput) Mutate(m *UserMutation) {
	m.SetName(i.Name)
	m.SetEmail(i.Email)
	m.SetPasswordHash(i.PasswordHash)
	if v := i.ProjectIDs; len(v) > 0 {
		m.AddProjectIDs(v...)
	}
	if v := i.AssignedTaskIDs; len(v) > 0 {
		m.AddAssignedTaskIDs(v...)
	}
}

// SetInput applies the change-set in the CreateUserInput on the UserCreate builder.
func (c *UserCreate) SetInput(i CreateUserInput) *UserCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateUserInput represents a mutation input for updating users.
type UpdateUserInput struct {
	Name                  *string
	Email                 *string
	PasswordHash          *string
	ClearProjects         bool
	AddProjectIDs         []int
	RemoveProjectIDs      []int
	ClearAssignedTasks    bool
	AddAssignedTaskIDs    []int
	RemoveAssignedTaskIDs []int
}

// Mutate applies the UpdateUserInput on the UserMutation builder.
func (i *UpdateUserInput) Mutate(m *UserMutation) {
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
	if v := i.Email; v != nil {
		m.SetEmail(*v)
	}
	if v := i.PasswordHash; v != nil {
		m.SetPasswordHash(*v)
	}
	if i.ClearProjects {
		m.ClearProjects()
	}
	if v := i.AddProjectIDs; len(v) > 0 {
		m.AddProjectIDs(v...)
	}
	if v := i.RemoveProjectIDs; len(v) > 0 {
		m.RemoveProjectIDs(v...)
	}
	if i.ClearAssignedTasks {
		m.ClearAssignedTasks()
	}
	if v := i.AddAssignedTaskIDs; len(v) > 0 {
		m.AddAssignedTaskIDs(v...)
	}
	if v := i.RemoveAssignedTaskIDs; len(v) > 0 {
		m.RemoveAssignedTaskIDs(v...)
	}
}

// SetInput applies the change-set in the UpdateUserInput on the UserUpdate builder.
func (c *UserUpdate) SetInput(i UpdateUserInput) *UserUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateUserInput on the UserUpdateOne builder.
func (c *UserUpdateOne) SetInput(i UpdateUserInput) *UserUpdateOne {
	i.Mutate(c.Mutation())
	return c
}
