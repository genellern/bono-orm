package bono_orm

import "sync"

type BonoConfig struct {
	dbConnection sync.Once
	preparedStmt string

	// Add logger
}

const ACTION_SELECT = "SELECT"
const ACTION_INSERT = "INSERT"
const ACTION_UPDATE = "UPDATE"
const ACTION_DELETE = "DELETE"

type BonoModel struct {
	// If SQL
	sqlStr string
	isSql  bool
}

type ModelCollection []BonoModel

func (model *BonoModel) Create() {

}

func (model *BonoModel) Update() {

}

func (model *BonoModel) Delete() {

}
func (model *BonoModel) Get() {

}

func (model *BonoModel) GetSql() (string, bool) {
	return model.sqlStr, model.isSql
}

func (model *BonoModel) Transact() *BonoModel {
	// Start transaction
	return model
}

func (model *BonoModel) Rollback() *BonoModel {
	return model
}

func (model *BonoModel) Commit() *BonoModel {
	// On defer commit
	return model
}
