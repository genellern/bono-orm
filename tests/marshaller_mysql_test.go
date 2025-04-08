package tests

import (
	bono_orm "github.com/genellern/bono-orm"
	"testing"
)

type TestModelStruct struct {
	bono_orm.BonoModel
	Id          int `bono:"id|readable"`
	Name        int `bono:"name|readable|fillable"`
	Description int `bono:"name|readable|fillable|nullable"`
	Password    int `bono:"name|hidden|fillable|not_nullable"`
}

func TestMarshallMapFieldsToValues(t *testing.T) {
	testObj := TestModelStruct{Id: 1}
	testObj.Get()
	expectedSqlStr := "SELECT id, name, description, password, name FROM test_model"
	sqlStr, _ := testObj.GetSql()
	if sqlStr != expectedSqlStr {
		t.Error("The expected sql doesn't match")
	}
}
