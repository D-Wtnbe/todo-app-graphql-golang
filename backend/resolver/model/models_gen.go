// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/D-Wtnbe/todo-app-graphql-golang/ent/task"
)

// CreateProjectInput is used for create Project object.
// Input was generated by ent.
type CreateProjectInput struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	UserID      *string   `json:"userID,omitempty"`
	TaskIDs     []string  `json:"taskIDs,omitempty"`
}

// CreateTaskInput is used for create Task object.
// Input was generated by ent.
type CreateTaskInput struct {
	Title         string      `json:"title"`
	Description   string      `json:"description"`
	DueDate       time.Time   `json:"dueDate"`
	Status        task.Status `json:"status"`
	CreatedAt     time.Time   `json:"createdAt"`
	UpdatedAt     time.Time   `json:"updatedAt"`
	ProjectID     *string     `json:"projectID,omitempty"`
	AssignedToIDs []string    `json:"assignedToIDs,omitempty"`
}

// CreateTodoInput is used for create Todo object.
// Input was generated by ent.
type CreateTodoInput struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Done  *bool  `json:"done,omitempty"`
}

// CreateUserInput is used for create User object.
// Input was generated by ent.
type CreateUserInput struct {
	Name            string   `json:"name"`
	Email           string   `json:"email"`
	PasswordHash    string   `json:"passwordHash"`
	ProjectIDs      []string `json:"projectIDs,omitempty"`
	AssignedTaskIDs []string `json:"assignedTaskIDs,omitempty"`
}

// Information about pagination in a connection.
// https://relay.dev/graphql/connections.htm#sec-undefined.PageInfo
type PageInfo struct {
	// When paginating forwards, are there more items?
	HasNextPage bool `json:"hasNextPage"`
	// When paginating backwards, are there more items?
	HasPreviousPage bool `json:"hasPreviousPage"`
	// When paginating backwards, the cursor to continue.
	StartCursor *string `json:"startCursor,omitempty"`
	// When paginating forwards, the cursor to continue.
	EndCursor *string `json:"endCursor,omitempty"`
}

type Project struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	User        *User     `json:"user,omitempty"`
	Tasks       []*Task   `json:"tasks,omitempty"`
}

func (Project) IsNode() {}

// ProjectWhereInput is used for filtering Project objects.
// Input was generated by ent.
type ProjectWhereInput struct {
	Not *ProjectWhereInput   `json:"not,omitempty"`
	And []*ProjectWhereInput `json:"and,omitempty"`
	Or  []*ProjectWhereInput `json:"or,omitempty"`
	// id field predicates
	ID      *string  `json:"id,omitempty"`
	IDNeq   *string  `json:"idNEQ,omitempty"`
	IDIn    []string `json:"idIn,omitempty"`
	IDNotIn []string `json:"idNotIn,omitempty"`
	IDGt    *string  `json:"idGT,omitempty"`
	IDGte   *string  `json:"idGTE,omitempty"`
	IDLt    *string  `json:"idLT,omitempty"`
	IDLte   *string  `json:"idLTE,omitempty"`
	// name field predicates
	Name             *string  `json:"name,omitempty"`
	NameNeq          *string  `json:"nameNEQ,omitempty"`
	NameIn           []string `json:"nameIn,omitempty"`
	NameNotIn        []string `json:"nameNotIn,omitempty"`
	NameGt           *string  `json:"nameGT,omitempty"`
	NameGte          *string  `json:"nameGTE,omitempty"`
	NameLt           *string  `json:"nameLT,omitempty"`
	NameLte          *string  `json:"nameLTE,omitempty"`
	NameContains     *string  `json:"nameContains,omitempty"`
	NameHasPrefix    *string  `json:"nameHasPrefix,omitempty"`
	NameHasSuffix    *string  `json:"nameHasSuffix,omitempty"`
	NameEqualFold    *string  `json:"nameEqualFold,omitempty"`
	NameContainsFold *string  `json:"nameContainsFold,omitempty"`
	// description field predicates
	Description             *string  `json:"description,omitempty"`
	DescriptionNeq          *string  `json:"descriptionNEQ,omitempty"`
	DescriptionIn           []string `json:"descriptionIn,omitempty"`
	DescriptionNotIn        []string `json:"descriptionNotIn,omitempty"`
	DescriptionGt           *string  `json:"descriptionGT,omitempty"`
	DescriptionGte          *string  `json:"descriptionGTE,omitempty"`
	DescriptionLt           *string  `json:"descriptionLT,omitempty"`
	DescriptionLte          *string  `json:"descriptionLTE,omitempty"`
	DescriptionContains     *string  `json:"descriptionContains,omitempty"`
	DescriptionHasPrefix    *string  `json:"descriptionHasPrefix,omitempty"`
	DescriptionHasSuffix    *string  `json:"descriptionHasSuffix,omitempty"`
	DescriptionEqualFold    *string  `json:"descriptionEqualFold,omitempty"`
	DescriptionContainsFold *string  `json:"descriptionContainsFold,omitempty"`
	// created_at field predicates
	CreatedAt      *time.Time   `json:"createdAt,omitempty"`
	CreatedAtNeq   *time.Time   `json:"createdAtNEQ,omitempty"`
	CreatedAtIn    []*time.Time `json:"createdAtIn,omitempty"`
	CreatedAtNotIn []*time.Time `json:"createdAtNotIn,omitempty"`
	CreatedAtGt    *time.Time   `json:"createdAtGT,omitempty"`
	CreatedAtGte   *time.Time   `json:"createdAtGTE,omitempty"`
	CreatedAtLt    *time.Time   `json:"createdAtLT,omitempty"`
	CreatedAtLte   *time.Time   `json:"createdAtLTE,omitempty"`
	// updated_at field predicates
	UpdatedAt      *time.Time   `json:"updatedAt,omitempty"`
	UpdatedAtNeq   *time.Time   `json:"updatedAtNEQ,omitempty"`
	UpdatedAtIn    []*time.Time `json:"updatedAtIn,omitempty"`
	UpdatedAtNotIn []*time.Time `json:"updatedAtNotIn,omitempty"`
	UpdatedAtGt    *time.Time   `json:"updatedAtGT,omitempty"`
	UpdatedAtGte   *time.Time   `json:"updatedAtGTE,omitempty"`
	UpdatedAtLt    *time.Time   `json:"updatedAtLT,omitempty"`
	UpdatedAtLte   *time.Time   `json:"updatedAtLTE,omitempty"`
	// user edge predicates
	HasUser     *bool             `json:"hasUser,omitempty"`
	HasUserWith []*UserWhereInput `json:"hasUserWith,omitempty"`
	// tasks edge predicates
	HasTasks     *bool             `json:"hasTasks,omitempty"`
	HasTasksWith []*TaskWhereInput `json:"hasTasksWith,omitempty"`
}

type Query struct {
}

type Task struct {
	ID          string      `json:"id"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	DueDate     time.Time   `json:"dueDate"`
	Status      task.Status `json:"status"`
	CreatedAt   time.Time   `json:"createdAt"`
	UpdatedAt   time.Time   `json:"updatedAt"`
	Project     *Project    `json:"project,omitempty"`
	AssignedTo  []*User     `json:"assignedTo,omitempty"`
}

func (Task) IsNode() {}

// TaskWhereInput is used for filtering Task objects.
// Input was generated by ent.
type TaskWhereInput struct {
	Not *TaskWhereInput   `json:"not,omitempty"`
	And []*TaskWhereInput `json:"and,omitempty"`
	Or  []*TaskWhereInput `json:"or,omitempty"`
	// id field predicates
	ID      *string  `json:"id,omitempty"`
	IDNeq   *string  `json:"idNEQ,omitempty"`
	IDIn    []string `json:"idIn,omitempty"`
	IDNotIn []string `json:"idNotIn,omitempty"`
	IDGt    *string  `json:"idGT,omitempty"`
	IDGte   *string  `json:"idGTE,omitempty"`
	IDLt    *string  `json:"idLT,omitempty"`
	IDLte   *string  `json:"idLTE,omitempty"`
	// title field predicates
	Title             *string  `json:"title,omitempty"`
	TitleNeq          *string  `json:"titleNEQ,omitempty"`
	TitleIn           []string `json:"titleIn,omitempty"`
	TitleNotIn        []string `json:"titleNotIn,omitempty"`
	TitleGt           *string  `json:"titleGT,omitempty"`
	TitleGte          *string  `json:"titleGTE,omitempty"`
	TitleLt           *string  `json:"titleLT,omitempty"`
	TitleLte          *string  `json:"titleLTE,omitempty"`
	TitleContains     *string  `json:"titleContains,omitempty"`
	TitleHasPrefix    *string  `json:"titleHasPrefix,omitempty"`
	TitleHasSuffix    *string  `json:"titleHasSuffix,omitempty"`
	TitleEqualFold    *string  `json:"titleEqualFold,omitempty"`
	TitleContainsFold *string  `json:"titleContainsFold,omitempty"`
	// description field predicates
	Description             *string  `json:"description,omitempty"`
	DescriptionNeq          *string  `json:"descriptionNEQ,omitempty"`
	DescriptionIn           []string `json:"descriptionIn,omitempty"`
	DescriptionNotIn        []string `json:"descriptionNotIn,omitempty"`
	DescriptionGt           *string  `json:"descriptionGT,omitempty"`
	DescriptionGte          *string  `json:"descriptionGTE,omitempty"`
	DescriptionLt           *string  `json:"descriptionLT,omitempty"`
	DescriptionLte          *string  `json:"descriptionLTE,omitempty"`
	DescriptionContains     *string  `json:"descriptionContains,omitempty"`
	DescriptionHasPrefix    *string  `json:"descriptionHasPrefix,omitempty"`
	DescriptionHasSuffix    *string  `json:"descriptionHasSuffix,omitempty"`
	DescriptionEqualFold    *string  `json:"descriptionEqualFold,omitempty"`
	DescriptionContainsFold *string  `json:"descriptionContainsFold,omitempty"`
	// due_date field predicates
	DueDate      *time.Time   `json:"dueDate,omitempty"`
	DueDateNeq   *time.Time   `json:"dueDateNEQ,omitempty"`
	DueDateIn    []*time.Time `json:"dueDateIn,omitempty"`
	DueDateNotIn []*time.Time `json:"dueDateNotIn,omitempty"`
	DueDateGt    *time.Time   `json:"dueDateGT,omitempty"`
	DueDateGte   *time.Time   `json:"dueDateGTE,omitempty"`
	DueDateLt    *time.Time   `json:"dueDateLT,omitempty"`
	DueDateLte   *time.Time   `json:"dueDateLTE,omitempty"`
	// status field predicates
	Status      *task.Status  `json:"status,omitempty"`
	StatusNeq   *task.Status  `json:"statusNEQ,omitempty"`
	StatusIn    []task.Status `json:"statusIn,omitempty"`
	StatusNotIn []task.Status `json:"statusNotIn,omitempty"`
	// created_at field predicates
	CreatedAt      *time.Time   `json:"createdAt,omitempty"`
	CreatedAtNeq   *time.Time   `json:"createdAtNEQ,omitempty"`
	CreatedAtIn    []*time.Time `json:"createdAtIn,omitempty"`
	CreatedAtNotIn []*time.Time `json:"createdAtNotIn,omitempty"`
	CreatedAtGt    *time.Time   `json:"createdAtGT,omitempty"`
	CreatedAtGte   *time.Time   `json:"createdAtGTE,omitempty"`
	CreatedAtLt    *time.Time   `json:"createdAtLT,omitempty"`
	CreatedAtLte   *time.Time   `json:"createdAtLTE,omitempty"`
	// updated_at field predicates
	UpdatedAt      *time.Time   `json:"updatedAt,omitempty"`
	UpdatedAtNeq   *time.Time   `json:"updatedAtNEQ,omitempty"`
	UpdatedAtIn    []*time.Time `json:"updatedAtIn,omitempty"`
	UpdatedAtNotIn []*time.Time `json:"updatedAtNotIn,omitempty"`
	UpdatedAtGt    *time.Time   `json:"updatedAtGT,omitempty"`
	UpdatedAtGte   *time.Time   `json:"updatedAtGTE,omitempty"`
	UpdatedAtLt    *time.Time   `json:"updatedAtLT,omitempty"`
	UpdatedAtLte   *time.Time   `json:"updatedAtLTE,omitempty"`
	// project edge predicates
	HasProject     *bool                `json:"hasProject,omitempty"`
	HasProjectWith []*ProjectWhereInput `json:"hasProjectWith,omitempty"`
	// assigned_to edge predicates
	HasAssignedTo     *bool             `json:"hasAssignedTo,omitempty"`
	HasAssignedToWith []*UserWhereInput `json:"hasAssignedToWith,omitempty"`
}

type Todo struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Done  bool   `json:"done"`
}

func (Todo) IsNode() {}

// TodoWhereInput is used for filtering Todo objects.
// Input was generated by ent.
type TodoWhereInput struct {
	Not *TodoWhereInput   `json:"not,omitempty"`
	And []*TodoWhereInput `json:"and,omitempty"`
	Or  []*TodoWhereInput `json:"or,omitempty"`
	// id field predicates
	ID      *string  `json:"id,omitempty"`
	IDNeq   *string  `json:"idNEQ,omitempty"`
	IDIn    []string `json:"idIn,omitempty"`
	IDNotIn []string `json:"idNotIn,omitempty"`
	IDGt    *string  `json:"idGT,omitempty"`
	IDGte   *string  `json:"idGTE,omitempty"`
	IDLt    *string  `json:"idLT,omitempty"`
	IDLte   *string  `json:"idLTE,omitempty"`
	// name field predicates
	Name             *string  `json:"name,omitempty"`
	NameNeq          *string  `json:"nameNEQ,omitempty"`
	NameIn           []string `json:"nameIn,omitempty"`
	NameNotIn        []string `json:"nameNotIn,omitempty"`
	NameGt           *string  `json:"nameGT,omitempty"`
	NameGte          *string  `json:"nameGTE,omitempty"`
	NameLt           *string  `json:"nameLT,omitempty"`
	NameLte          *string  `json:"nameLTE,omitempty"`
	NameContains     *string  `json:"nameContains,omitempty"`
	NameHasPrefix    *string  `json:"nameHasPrefix,omitempty"`
	NameHasSuffix    *string  `json:"nameHasSuffix,omitempty"`
	NameEqualFold    *string  `json:"nameEqualFold,omitempty"`
	NameContainsFold *string  `json:"nameContainsFold,omitempty"`
	// email field predicates
	Email             *string  `json:"email,omitempty"`
	EmailNeq          *string  `json:"emailNEQ,omitempty"`
	EmailIn           []string `json:"emailIn,omitempty"`
	EmailNotIn        []string `json:"emailNotIn,omitempty"`
	EmailGt           *string  `json:"emailGT,omitempty"`
	EmailGte          *string  `json:"emailGTE,omitempty"`
	EmailLt           *string  `json:"emailLT,omitempty"`
	EmailLte          *string  `json:"emailLTE,omitempty"`
	EmailContains     *string  `json:"emailContains,omitempty"`
	EmailHasPrefix    *string  `json:"emailHasPrefix,omitempty"`
	EmailHasSuffix    *string  `json:"emailHasSuffix,omitempty"`
	EmailEqualFold    *string  `json:"emailEqualFold,omitempty"`
	EmailContainsFold *string  `json:"emailContainsFold,omitempty"`
	// done field predicates
	Done    *bool `json:"done,omitempty"`
	DoneNeq *bool `json:"doneNEQ,omitempty"`
}

// UpdateProjectInput is used for update Project object.
// Input was generated by ent.
type UpdateProjectInput struct {
	Name          *string    `json:"name,omitempty"`
	Description   *string    `json:"description,omitempty"`
	CreatedAt     *time.Time `json:"createdAt,omitempty"`
	UpdatedAt     *time.Time `json:"updatedAt,omitempty"`
	UserID        *string    `json:"userID,omitempty"`
	ClearUser     *bool      `json:"clearUser,omitempty"`
	AddTaskIDs    []string   `json:"addTaskIDs,omitempty"`
	RemoveTaskIDs []string   `json:"removeTaskIDs,omitempty"`
	ClearTasks    *bool      `json:"clearTasks,omitempty"`
}

// UpdateTaskInput is used for update Task object.
// Input was generated by ent.
type UpdateTaskInput struct {
	Title               *string      `json:"title,omitempty"`
	Description         *string      `json:"description,omitempty"`
	DueDate             *time.Time   `json:"dueDate,omitempty"`
	Status              *task.Status `json:"status,omitempty"`
	CreatedAt           *time.Time   `json:"createdAt,omitempty"`
	UpdatedAt           *time.Time   `json:"updatedAt,omitempty"`
	ProjectID           *string      `json:"projectID,omitempty"`
	ClearProject        *bool        `json:"clearProject,omitempty"`
	AddAssignedToIDs    []string     `json:"addAssignedToIDs,omitempty"`
	RemoveAssignedToIDs []string     `json:"removeAssignedToIDs,omitempty"`
	ClearAssignedTo     *bool        `json:"clearAssignedTo,omitempty"`
}

// UpdateTodoInput is used for update Todo object.
// Input was generated by ent.
type UpdateTodoInput struct {
	Name  *string `json:"name,omitempty"`
	Email *string `json:"email,omitempty"`
	Done  *bool   `json:"done,omitempty"`
}

// UpdateUserInput is used for update User object.
// Input was generated by ent.
type UpdateUserInput struct {
	Name                  *string  `json:"name,omitempty"`
	Email                 *string  `json:"email,omitempty"`
	PasswordHash          *string  `json:"passwordHash,omitempty"`
	AddProjectIDs         []string `json:"addProjectIDs,omitempty"`
	RemoveProjectIDs      []string `json:"removeProjectIDs,omitempty"`
	ClearProjects         *bool    `json:"clearProjects,omitempty"`
	AddAssignedTaskIDs    []string `json:"addAssignedTaskIDs,omitempty"`
	RemoveAssignedTaskIDs []string `json:"removeAssignedTaskIDs,omitempty"`
	ClearAssignedTasks    *bool    `json:"clearAssignedTasks,omitempty"`
}

type User struct {
	ID            string     `json:"id"`
	Name          string     `json:"name"`
	Email         string     `json:"email"`
	PasswordHash  string     `json:"passwordHash"`
	Projects      []*Project `json:"projects,omitempty"`
	AssignedTasks []*Task    `json:"assignedTasks,omitempty"`
}

func (User) IsNode() {}

// UserWhereInput is used for filtering User objects.
// Input was generated by ent.
type UserWhereInput struct {
	Not *UserWhereInput   `json:"not,omitempty"`
	And []*UserWhereInput `json:"and,omitempty"`
	Or  []*UserWhereInput `json:"or,omitempty"`
	// id field predicates
	ID      *string  `json:"id,omitempty"`
	IDNeq   *string  `json:"idNEQ,omitempty"`
	IDIn    []string `json:"idIn,omitempty"`
	IDNotIn []string `json:"idNotIn,omitempty"`
	IDGt    *string  `json:"idGT,omitempty"`
	IDGte   *string  `json:"idGTE,omitempty"`
	IDLt    *string  `json:"idLT,omitempty"`
	IDLte   *string  `json:"idLTE,omitempty"`
	// name field predicates
	Name             *string  `json:"name,omitempty"`
	NameNeq          *string  `json:"nameNEQ,omitempty"`
	NameIn           []string `json:"nameIn,omitempty"`
	NameNotIn        []string `json:"nameNotIn,omitempty"`
	NameGt           *string  `json:"nameGT,omitempty"`
	NameGte          *string  `json:"nameGTE,omitempty"`
	NameLt           *string  `json:"nameLT,omitempty"`
	NameLte          *string  `json:"nameLTE,omitempty"`
	NameContains     *string  `json:"nameContains,omitempty"`
	NameHasPrefix    *string  `json:"nameHasPrefix,omitempty"`
	NameHasSuffix    *string  `json:"nameHasSuffix,omitempty"`
	NameEqualFold    *string  `json:"nameEqualFold,omitempty"`
	NameContainsFold *string  `json:"nameContainsFold,omitempty"`
	// email field predicates
	Email             *string  `json:"email,omitempty"`
	EmailNeq          *string  `json:"emailNEQ,omitempty"`
	EmailIn           []string `json:"emailIn,omitempty"`
	EmailNotIn        []string `json:"emailNotIn,omitempty"`
	EmailGt           *string  `json:"emailGT,omitempty"`
	EmailGte          *string  `json:"emailGTE,omitempty"`
	EmailLt           *string  `json:"emailLT,omitempty"`
	EmailLte          *string  `json:"emailLTE,omitempty"`
	EmailContains     *string  `json:"emailContains,omitempty"`
	EmailHasPrefix    *string  `json:"emailHasPrefix,omitempty"`
	EmailHasSuffix    *string  `json:"emailHasSuffix,omitempty"`
	EmailEqualFold    *string  `json:"emailEqualFold,omitempty"`
	EmailContainsFold *string  `json:"emailContainsFold,omitempty"`
	// password_hash field predicates
	PasswordHash             *string  `json:"passwordHash,omitempty"`
	PasswordHashNeq          *string  `json:"passwordHashNEQ,omitempty"`
	PasswordHashIn           []string `json:"passwordHashIn,omitempty"`
	PasswordHashNotIn        []string `json:"passwordHashNotIn,omitempty"`
	PasswordHashGt           *string  `json:"passwordHashGT,omitempty"`
	PasswordHashGte          *string  `json:"passwordHashGTE,omitempty"`
	PasswordHashLt           *string  `json:"passwordHashLT,omitempty"`
	PasswordHashLte          *string  `json:"passwordHashLTE,omitempty"`
	PasswordHashContains     *string  `json:"passwordHashContains,omitempty"`
	PasswordHashHasPrefix    *string  `json:"passwordHashHasPrefix,omitempty"`
	PasswordHashHasSuffix    *string  `json:"passwordHashHasSuffix,omitempty"`
	PasswordHashEqualFold    *string  `json:"passwordHashEqualFold,omitempty"`
	PasswordHashContainsFold *string  `json:"passwordHashContainsFold,omitempty"`
	// projects edge predicates
	HasProjects     *bool                `json:"hasProjects,omitempty"`
	HasProjectsWith []*ProjectWhereInput `json:"hasProjectsWith,omitempty"`
	// assigned_tasks edge predicates
	HasAssignedTasks     *bool             `json:"hasAssignedTasks,omitempty"`
	HasAssignedTasksWith []*TaskWhereInput `json:"hasAssignedTasksWith,omitempty"`
}

// Possible directions in which to order a list of items when provided an `orderBy` argument.
type OrderDirection string

const (
	// Specifies an ascending order for a given `orderBy` argument.
	OrderDirectionAsc OrderDirection = "ASC"
	// Specifies a descending order for a given `orderBy` argument.
	OrderDirectionDesc OrderDirection = "DESC"
)

var AllOrderDirection = []OrderDirection{
	OrderDirectionAsc,
	OrderDirectionDesc,
}

func (e OrderDirection) IsValid() bool {
	switch e {
	case OrderDirectionAsc, OrderDirectionDesc:
		return true
	}
	return false
}

func (e OrderDirection) String() string {
	return string(e)
}

func (e *OrderDirection) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = OrderDirection(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid OrderDirection", str)
	}
	return nil
}

func (e OrderDirection) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
