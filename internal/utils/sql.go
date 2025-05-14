package utils

import (
	"database/sql"
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/go-sql-driver/mysql"
)

// SQL Errors
var (
	ErrDBInternalServer             = fmt.Errorf("internal server error")
	ErrDBForeignKey                 = fmt.Errorf("foreign key constraint error")
	ErrDBDuplicateEntry             = fmt.Errorf("duplicate entry error")
	ErrDBCannotDeleteOrUpdateParent = fmt.Errorf("cannot delete or update a parent row: a foreign key constraint fails")
)

func ValidateErrorTypeSQL(err error) error {
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) {
			switch mysqlErr.Number {
			case 1452: // Code for foreign key constraint fails
				return ErrDBForeignKey
			case 1062: // Code for duplicate entry in unique key
				return ErrDBDuplicateEntry
			case 1451: // Code for cannot delete or update a parent row: a foreign key constraint fails
				return ErrDBCannotDeleteOrUpdateParent
			default:
				return ErrDBInternalServer
			}
		}
	}
	return nil
}

func ValidateConstraintFailed(err error, constraintName string) bool {
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) {
			if strings.Contains(mysqlErr.Message, constraintName) {
				return true
			}
		}
	}
	return false
}

func CloseRows(rows *sql.Rows) {
	err := rows.Close()
	if err != nil {
		fmt.Println("Error closing rows:", err)
	}
}

// BuildSQLWhereClause builds the WHERE clause and arguments for a given filter and its field to column mapping
func BuildSQLWhereClause(filter any, fieldColumnMap map[string]string) (string, []interface{}) {
	v := reflect.ValueOf(filter).Elem()
	t := v.Type()
	clauses := []string{}
	args := []interface{}{}
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if field.Kind() == reflect.Ptr && !field.IsNil() {
			col, ok := fieldColumnMap[t.Field(i).Name]
			if !ok {
				continue
			}
			clauses = append(clauses, col+" = ?")
			args = append(args, field.Elem().Interface())
		}
	}
	if len(clauses) == 0 {
		return "", args
	}
	where := " WHERE " + clauses[0]
	for i := 1; i < len(clauses); i++ {
		where += " AND " + clauses[i]
	}
	return where, args
}
