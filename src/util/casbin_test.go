package util

import (
	"testing"
)

const (
	CONFIG_FOLDER    = "../../config/"
	MODEL_FILE_NAME  = "model.ini"
	POLICY_FILE_NAME = "policy.csv"
)

func TestCasbin_Enforce(t *testing.T) {
	type args struct {
		req []interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantRes bool
	}{
		{
			args: args{
				req: []interface{}{
					"admin",  // Subject
					"/user/", // Endpoint
					"GET",    // Method
				},
			},
			wantRes: true,
		},
		{
			args: args{
				req: []interface{}{
					"admin",
					"/user/1",
					"DELETE",
				},
			},
			wantRes: true,
		},
		{
			args: args{
				req: []interface{}{
					"user",
					"/statistic",
					"GET",
				},
			},
			wantRes: true,
		},
		{
			args: args{
				req: []interface{}{
					"user",
					"/user/1",
					"DELETE",
				},
			},
			wantRes: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := Casbin{
				ModelFilePath:  CONFIG_FOLDER + MODEL_FILE_NAME,
				PolicyFilePath: CONFIG_FOLDER + POLICY_FILE_NAME,
			}
			if gotRes := this.Enforce(this.GetEnforcer(), tt.args.req); gotRes != tt.wantRes {
				t.Errorf("\nCasbin.Enforce() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
