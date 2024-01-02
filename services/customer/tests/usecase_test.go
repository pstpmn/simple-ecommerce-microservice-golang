package tests

import (
	"simple-ecomerce-microservice/pkg"
	"simple-ecomerce-microservice/services/customer"
	customerCore "simple-ecomerce-microservice/services/customer/core"
	"testing"
	"time"
)

func Test_u_CreateCustomer(t *testing.T) {
	// create object postgresql
	postgres := pkg.NewGormORM()
	helper := pkg.NewHelper()
	dsn := "host=localhost user=postgres password=root dbname=customer port=5432 sslmode=disable TimeZone=Asia/Bangkok"
	conn, err := postgres.ConnectDB(dsn, "postgres")
	if err != nil {
		panic(err)
	}
	repo := customer.NewRepository(conn)

	type fields struct {
		repo   customerCore.ICustomerRepo
		helper customerCore.IHelper
	}
	type args struct {
		firstName string
		lastName  string
		phoneNo   string
		dob       time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *customerCore.Customer
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "should be create success",
			fields: fields{
				repo:   repo,
				helper: helper,
			},
			args: args{
				firstName: "test",
				lastName:  "test",
				phoneNo:   "099999999",
				dob:       time.Now(),
			},
			want:    &customerCore.Customer{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			usecase := customerCore.NewUseCase(tt.fields.repo, tt.fields.helper)
			_, err := usecase.CreateCustomer(tt.args.firstName, tt.args.lastName, tt.args.phoneNo, tt.args.dob)
			if (err != nil) != tt.wantErr {
				t.Errorf("u.CreateCustomer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// if !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("u.CreateCustomer() = %v, want %v", got, tt.want)
			// }
		})
	}
}

func Test_u_CreateAddress(t *testing.T) {
	postgres := pkg.NewGormORM()
	helper := pkg.NewHelper()
	dsn := "host=localhost user=postgres password=root dbname=customer port=5432 sslmode=disable TimeZone=Asia/Bangkok"
	conn, err := postgres.ConnectDB(dsn, "postgres")
	if err != nil {
		panic(err)
	}
	repo := customer.NewRepository(conn)

	type fields struct {
		repo   customerCore.ICustomerRepo
		helper customerCore.IHelper
	}
	type args struct {
		customerId    string
		streetAddress string
		city          string
		state         string
		postalCodes   string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *customerCore.Address
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
				customerId:    "63af8066-6474-465a-ac9d-fe6863df1834",
				streetAddress: "test",
				city:          "test",
				state:         "test",
				postalCodes:   "test",
			},
			want:    &customerCore.Address{},
			wantErr: false,
		},
		{
			name: "should be error not found customer",
			fields: fields{
				repo:   repo,
				helper: helper,
			},
			args: args{
				customerId:    "not found",
				streetAddress: "test",
				city:          "test",
				state:         "test",
				postalCodes:   "test",
			},
			want:    &customerCore.Address{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			usecase := customerCore.NewUseCase(tt.fields.repo, tt.fields.helper)
			_, err := usecase.CreateAddress(tt.args.customerId, tt.args.streetAddress, tt.args.city, tt.args.state, tt.args.postalCodes)
			if (err != nil) != tt.wantErr {
				t.Errorf("u.CreateAddress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// if !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("u.CreateAddress() = %v, want %v", got, tt.want)
			// }
		})
	}
}

func Test_u_GetProfile(t *testing.T) {
	postgres := pkg.NewGormORM()
	helper := pkg.NewHelper()
	dsn := "host=localhost user=postgres password=root dbname=customer port=5432 sslmode=disable TimeZone=Asia/Bangkok"
	conn, err := postgres.ConnectDB(dsn, "postgres")
	if err != nil {
		panic(err)
	}
	repo := customer.NewRepository(conn)
	type fields struct {
		repo   customerCore.ICustomerRepo
		helper customerCore.IHelper
	}
	type args struct {
		customerId string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *customerCore.CustomerProfile
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
			},
			want:    &customerCore.CustomerProfile{},
			wantErr: false,
		},
		{
			name: "should be error not found customer",
			fields: fields{
				repo:   repo,
				helper: helper,
			},
			args: args{
				customerId: "test",
			},
			want:    &customerCore.CustomerProfile{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			usecase := customerCore.NewUseCase(tt.fields.repo, tt.fields.helper)
			_, err := usecase.GetProfile(tt.args.customerId)
			if (err != nil) != tt.wantErr {
				t.Errorf("u.GetProfile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// if !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("u.GetProfile() = %v, want %v", got, tt.want)
			// }
		})
	}
}
