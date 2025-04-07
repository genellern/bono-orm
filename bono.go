package bono_orm

import "sync"

type BonoConfig struct {
	dbConnection sync.Once
	preparedStmt string

	// Add logger
}

type BonoModel struct {
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

func (model *BonoModel) Transact() *BonoModel {
	// Start transaction
	// On defer commit
	return model
}
func (model *BonoModel) Rollback() *BonoModel {
	return model
}
func (model *BonoModel) Commit() *BonoModel {
	// On defer commit
	return model
}
