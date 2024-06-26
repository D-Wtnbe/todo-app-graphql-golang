// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ProjectsColumns holds the columns for the "projects" table.
	ProjectsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "user_projects", Type: field.TypeInt, Nullable: true},
	}
	// ProjectsTable holds the schema information for the "projects" table.
	ProjectsTable = &schema.Table{
		Name:       "projects",
		Columns:    ProjectsColumns,
		PrimaryKey: []*schema.Column{ProjectsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "projects_users_projects",
				Columns:    []*schema.Column{ProjectsColumns[5]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// TasksColumns holds the columns for the "tasks" table.
	TasksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "title", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "due_date", Type: field.TypeTime},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"TODO", "IN_PROGRESS", "DONE"}},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "project_tasks", Type: field.TypeInt, Nullable: true},
	}
	// TasksTable holds the schema information for the "tasks" table.
	TasksTable = &schema.Table{
		Name:       "tasks",
		Columns:    TasksColumns,
		PrimaryKey: []*schema.Column{TasksColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "tasks_projects_tasks",
				Columns:    []*schema.Column{TasksColumns[7]},
				RefColumns: []*schema.Column{ProjectsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// TodosColumns holds the columns for the "todos" table.
	TodosColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "email", Type: field.TypeString, Size: 255},
		{Name: "done", Type: field.TypeBool, Default: false},
	}
	// TodosTable holds the schema information for the "todos" table.
	TodosTable = &schema.Table{
		Name:       "todos",
		Columns:    TodosColumns,
		PrimaryKey: []*schema.Column{TodosColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "password_hash", Type: field.TypeString},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// UserAssignedTasksColumns holds the columns for the "user_assigned_tasks" table.
	UserAssignedTasksColumns = []*schema.Column{
		{Name: "user_id", Type: field.TypeInt},
		{Name: "task_id", Type: field.TypeInt},
	}
	// UserAssignedTasksTable holds the schema information for the "user_assigned_tasks" table.
	UserAssignedTasksTable = &schema.Table{
		Name:       "user_assigned_tasks",
		Columns:    UserAssignedTasksColumns,
		PrimaryKey: []*schema.Column{UserAssignedTasksColumns[0], UserAssignedTasksColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_assigned_tasks_user_id",
				Columns:    []*schema.Column{UserAssignedTasksColumns[0]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "user_assigned_tasks_task_id",
				Columns:    []*schema.Column{UserAssignedTasksColumns[1]},
				RefColumns: []*schema.Column{TasksColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ProjectsTable,
		TasksTable,
		TodosTable,
		UsersTable,
		UserAssignedTasksTable,
	}
)

func init() {
	ProjectsTable.ForeignKeys[0].RefTable = UsersTable
	TasksTable.ForeignKeys[0].RefTable = ProjectsTable
	UserAssignedTasksTable.ForeignKeys[0].RefTable = UsersTable
	UserAssignedTasksTable.ForeignKeys[1].RefTable = TasksTable
}
