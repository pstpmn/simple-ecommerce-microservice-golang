package productTest

import (
	"reflect"
	orderCore "simple-ecomerce-microservice/services/order/core"
	"testing"
)

func Test_u_CreateOrder(t *testing.T) {
	type fields struct {
		repo   orderCore.IOrderRepo
		helper orderCore.IHelper
	}
	type args struct {
		order orderCore.Order
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *orderCore.OrderProfile
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := orderCore.NewUseCase(
				tt.fields.repo,
				tt.fields.helper,
			)
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

// func Test_u_CreateOrder(t *testing.T) {
// 	type fields struct {
// 		repo   orderCore.IOrderRepo
// 		helper orderCore.IHelper
// 	}
// 	type args struct {
// 		order   orderCore.Order
// 		details orderCore.OrderDetail
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		args    args
// 		want    *OrderProfile
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			u := &u{
// 				repo:   tt.fields.repo,
// 				helper: tt.fields.helper,
// 			}
// 			got, err := u.CreateOrder(tt.args.order, tt.args.details)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("u.CreateOrder() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("u.CreateOrder() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
