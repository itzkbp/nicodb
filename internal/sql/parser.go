package sqlparser

type _Result struct {
	output string
}

type SQLQuery interface {
	Execute() _Result
}

type ColumnDefinition struct {
	columnName   string
	dataType     string
	isPrimaryKey bool
	isNullable   bool
}

type CreateTableStmt struct {
	tableName string
	columns   []ColumnDefinition
}

func (t *CreateTableStmt) Execute() {
	// Execute Create Table
}

type InsertStmt struct {
	tableName string
	columns   []string
	data      []string
}

func (t *InsertStmt) Execute() {
	// Execute Insert Statement
}

type Operator uint8

const (
	OP_EQUALS Operator = iota
	OP_NOT_EQUALS
	OP_GT
	OP_LT
	OP_GT_EQUALS
	OP_LT_EQUALS
)

type Condition struct {
	option1    string
	operator   Operator
	option2    string
	isOp2Field bool
}

type SelectStmt struct {
	tableName string
	columns   []string
	condition Condition
}
