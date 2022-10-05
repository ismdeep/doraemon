package pkg

import (
	"fmt"
	"testing"

	"github.com/ismdeep/rand"
)

func TestConnectToMySQL(t *testing.T) {
	type args struct {
		dsn     string
		timeout int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "",
			args: args{
				dsn:     "root:doraemon31415@tcp(127.0.0.1:3306)/mysql?parseTime=true&loc=Local&charset=utf8mb4,utf8",
				timeout: 60,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ConnectToMySQL(tt.args.dsn, tt.args.timeout)
		})
	}
}

func TestCreateDBOnMySQL(t *testing.T) {
	type args struct {
		dsn           string
		dbName        string
		additionAuths []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "",
			args: args{
				dsn:    "root:doraemon31415@tcp(127.0.0.1:3306)/mysql?parseTime=true&loc=Local&charset=utf8mb4,utf8",
				dbName: fmt.Sprintf("test-%v", rand.HexStr(16)),
				additionAuths: []string{
					fmt.Sprintf("user%v:%v", rand.HexStr(16), rand.HexStr(16)),
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateDBOnMySQL(tt.args.dsn, tt.args.dbName, tt.args.additionAuths); (err != nil) != tt.wantErr {
				t.Errorf("CreateDBOnMySQL() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
