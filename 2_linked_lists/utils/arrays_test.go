package utils

import (
	"reflect"
	"testing"
)

func TestStringArrayToIntArray(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		{
			name: "nil string array",
			args: args{
				nil,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "empty string array",
			args: args{
				[]string{},
			},
			want:    []int{},
			wantErr: false,
		},
		{
			name: "string array with one value",
			args: args{
				[]string{"1"},
			},
			want:    []int{1},
			wantErr: false,
		},
		{
			name: "string array with multiple values",
			args: args{
				[]string{"1", "2", "3", "4", "5"},
			},
			want:    []int{1, 2, 3, 4, 5},
			wantErr: false,
		},
		{
			name: "string array with invalid value",
			args: args{
				[]string{"1", "2", "3", "A", "5"},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StringArrayToIntArray(tt.args.strs)
			if (err != nil) != tt.wantErr {
				t.Errorf("StringArrayToIntArray() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringArrayToIntArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
