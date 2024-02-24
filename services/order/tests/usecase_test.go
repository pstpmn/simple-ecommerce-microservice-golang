package orderTest

import (
	"context"
	"fmt"
	"simple-ecomerce-microservice/pkg"
	"simple-ecomerce-microservice/services/order"
	orderCore "simple-ecomerce-microservice/services/order/core"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Test_u_CreateOrder(t *testing.T) {
	mongo := pkg.NewMongo()
	helper := pkg.NewHelper()
	grpc := pkg.NewGrpc()
	conn, err := mongo.Connect("mongodb://root:root@localhost:27017")
	if err != nil {
		panic(fmt.Sprintf("error mongo: %s", err))
	}
	repo := order.NewRepository(conn, grpc)
	type fields struct {
		repo   orderCore.IOrderRepo
		helper orderCore.IHelper
	}
	type args struct {
		customerId string
		products   []orderCore.OrderDetail
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *orderCore.OrderProfile
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "should be success",
			fields: fields{
				repo:   repo,
				helper: helper,
			},
			args: args{
				customerId: "63af8066-6474-465a-ac9d-fe6863df1834",
				products: func() []orderCore.OrderDetail {
					return []orderCore.OrderDetail{
						{
							ProductId: "1",
							Quantity:  1,
							Price:     100,
							CreatedAt: time.Now(),
						},
					}
				}(),
			},
			want:    &orderCore.OrderProfile{},
			wantErr: false,
		},
		{
			name: "should be error not found customer",
			fields: fields{
				repo:   repo,
				helper: helper,
			},
			args: args{
				customerId: "invalid",
				products: func() []orderCore.OrderDetail {
					return []orderCore.OrderDetail{
						{
							ProductId: "550e8400-e29b-41d4-a716-446655440000",
							Quantity:  1,
							Price:     100,
							CreatedAt: time.Now(),
						},
					}
				}(),
			},
			want:    &orderCore.OrderProfile{},
			wantErr: true,
		},
		{
			name: "should be error productId invalid",
			fields: fields{
				repo:   repo,
				helper: helper,
			},
			args: args{
				customerId: "63af8066-6474-465a-ac9d-fe6863df1834",
				products: func() []orderCore.OrderDetail {
					return []orderCore.OrderDetail{
						{
							ProductId: "00000000-e29b-41d4-a716-446655440000",
							Quantity:  1,
							Price:     100,
							CreatedAt: time.Now(),
						},
					}
				}(),
			},
			want:    &orderCore.OrderProfile{},
			wantErr: true,
		},
		{
			name: "should be error stock product is zero item",
			fields: fields{
				repo:   repo,
				helper: helper,
			},
			args: args{
				customerId: "63af8066-6474-465a-ac9d-fe6863df1834",
				products: func() []orderCore.OrderDetail {
					return []orderCore.OrderDetail{
						{
							ProductId: "pro01",
							Quantity:  1,
							Price:     100,
							CreatedAt: time.Now(),
						},
					}
				}(),
			},
			want:    &orderCore.OrderProfile{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := orderCore.NewUseCase(
				tt.fields.repo,
				tt.fields.helper,
			)
			_, err := u.CreateOrder(context.Background(), tt.args.customerId, tt.args.products)
			if (err != nil) != tt.wantErr {
				t.Errorf("u.CreateOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// if !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("u.CreateOrder() = %v, want %v", got, tt.want)
			// }
		})
	}
}

func Test_u_CancelOrder(t *testing.T) {
	mongo := pkg.NewMongo()
	helper := pkg.NewHelper()
	grpc := pkg.NewGrpc()
	conn, err := mongo.Connect("mongodb://root:root@localhost:27017")
	if err != nil {
		panic(fmt.Sprintf("error mongo: %s", err))
	}
	repo := order.NewRepository(conn, grpc)

	type fields struct {
		repo   orderCore.IOrderRepo
		helper orderCore.IHelper
	}
	type args struct {
		pctx       context.Context
		customerId string
		orderId    primitive.ObjectID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "should be success",
			fields: fields{
				repo:   repo,
				helper: helper,
			},
			args: args{
				pctx:       context.Background(),
				customerId: "63af8066-6474-465a-ac9d-fe6863df1834",
				orderId: func() primitive.ObjectID {
					orderId, _ := primitive.ObjectIDFromHex("65cb8c228976ddfe9b9a2df2")
					return orderId
				}(),
			},
			wantErr: false,
		},
		{
			name: "should error becouse status is canceled",
			fields: fields{
				repo:   repo,
				helper: helper,
			},
			args: args{
				pctx:       context.Background(),
				customerId: "63af8066-6474-465a-ac9d-fe6863df1834",
				orderId: func() primitive.ObjectID {
					orderId, _ := primitive.ObjectIDFromHex("65cb8c228976ddfe9b9a2df2")
					return orderId
				}(),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := orderCore.NewUseCase(
				tt.fields.repo,
				tt.fields.helper,
			)
			if err := u.CancelOrder(tt.args.pctx, tt.args.customerId, tt.args.orderId); (err != nil) != tt.wantErr {
				t.Errorf("u.CancelOrder() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_u_GetOrderDetail(t *testing.T) {
	mongo := pkg.NewMongo()
	helper := pkg.NewHelper()
	grpc := pkg.NewGrpc()
	conn, err := mongo.Connect("mongodb://root:root@localhost:27017")
	if err != nil {
		panic(fmt.Sprintf("error mongo: %s", err))
	}
	repo := order.NewRepository(conn, grpc)

	type fields struct {
		repo   orderCore.IOrderRepo
		helper orderCore.IHelper
	}
	type args struct {
		pctx    context.Context
		orderId primitive.ObjectID
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		wantProfile *orderCore.OrderProfile
		wantErr     bool
	}{
		// TODO: Add test cases.
		{
			name: "should be get order details success",
			fields: fields{
				repo:   repo,
				helper: helper,
			},
			args: args{
				pctx: context.Background(),
				orderId: func() primitive.ObjectID {
					orderId, _ := primitive.ObjectIDFromHex("65ca37f3e54ad0b49a93c9a1")
					return orderId
				}(),
			},
			// wantProfile: &orderCore.OrderProfile{},
			wantErr: false,
		},
		{
			name: "should be get order details success",
			fields: fields{
				repo:   repo,
				helper: helper,
			},
			args: args{
				pctx: context.Background(),
				orderId: func() primitive.ObjectID {
					orderId, _ := primitive.ObjectIDFromHex("invalid")
					return orderId
				}(),
			},
			// wantProfile: &orderCore.OrderProfile{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := orderCore.NewUseCase(
				tt.fields.repo,
				tt.fields.helper,
			)
			_, err := u.GetOrderDetail(tt.args.pctx, tt.args.orderId)
			if (err != nil) != tt.wantErr {
				t.Errorf("u.GetOrderDetail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// if !reflect.DeepEqual(gotProfile, tt.wantProfile) {
			// 	t.Errorf("u.GetOrderDetail() = %v, want %v", gotProfile, tt.wantProfile)
			// }
		})
	}
}
