package mysql

import (
	"database/sql"
	"reflect"
	"testing"

	"github.com/jameshwc/Million-Singer/model"
)

func Test_mysqlTourRepository_Get(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Tour
		wantErr bool
	}{
		// TODO: Add test cases.
		// {name:"success", fields:{db: db}}
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &mysqlTourRepository{
				db: tt.fields.db,
			}
			got, err := m.Get(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("mysqlTourRepository.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mysqlTourRepository.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
