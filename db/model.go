package db

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
)

// model
type Operator string

const (
	And Operator = "AND"
	Or  Operator = "Or"
	Eq  Operator = "="
	Lt  Operator = "<="
	Gt  Operator = ">="
)

type Filter struct {
	Atribute string
	Operator Operator
	Value    any

	Filter *Filter
}

type ModelDB struct {
	table_name string
	filters    string
}

func GetModel(instance any) *ModelDB {
	return &ModelDB{
		table_name: "users",
	}
}

func (this *ModelDB) Select(args ...Filter) *ModelDB {
	for _, arg := range args {
		this.filters += fmt.Sprintf("( %s %s '%v')", arg.Atribute, arg.Operator, arg.Value)
	}

	return this
}

func (this *ModelDB) Find(ctx context.Context, instance any) (err error) {
	query := this.build()

	// TODO: error
	conn, _ := this.getConnection(ctx)
	rows := conn.QueryRow(query)

	var name string

	err = rows.Scan(&name)
	if err != nil {
		return err
	}

	stype := reflect.ValueOf(instance).Elem()
	field := stype.FieldByName("Name")
	if field.IsValid() {
		field.SetString(name)
	}

	return nil
}

func (this *ModelDB) Insert(ctx context.Context, instance any) error {
	query := this.buildInsert()

	_, err := this.getConnection(ctx).Exec(query)
	return err
}

func (this *ModelDB) RunInTransaction(ctx context.Context, queryFunc func(ctx context.Context) error) error {
	ctx, err := this.setTransaction(ctx)
	if err != nil {
		return err
	}

	return queryFunc(ctx)
}

func (this *ModelDB) setTransaction(ctx context.Context) (context.Context, error) {
	trx, err := globalConnection.db.Begin()
	if err != nil {
		return nil, err
	}

	return context.WithValue(ctx, "transaction", trx), nil
}

func (this *ModelDB) getConnection(ctx context.Context) (*sql.DB, error) {

	return globalConnection.db
}

func (this *ModelDB) buildInsert() string {
	return fmt.Sprintf("INSERT INTO %s (name) VALUES ('yaya')", this.table_name)
}

func (this *ModelDB) build() string {
	return fmt.Sprintf("select * from %s where%s", this.table_name, this.filters)
}
