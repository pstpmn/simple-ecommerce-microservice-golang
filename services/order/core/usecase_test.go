package orderCore

import (
	"reflect"
	"testing"
)

func Test_u_CreateOrder(t *testing.T) {
	type fields struct {
		repo   IOrderRepo
		helper IHelper
	}
	type args struct {
		order Order
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *OrderProfile
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &u{
				repo:   tt.fields.repo,
				helper: tt.fields.helper,
			}
			got, err := u.CreateOrder(tt.args.order)
			if (err != nil) != tt.wantErr {
				t.Errorf("u.CreateOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("u.CreateOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
