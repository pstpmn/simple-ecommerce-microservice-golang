package productTest

import (
	"context"
	"simple-ecomerce-microservice/pkg"
	"simple-ecomerce-microservice/services/product"
	productCore "simple-ecomerce-microservice/services/product/core"
	"testing"
)

func Test_usecase_GetProducts(t *testing.T) {
	postgres := pkg.NewGormORM()
	helper := pkg.NewHelper()
	dsn := "host=localhost user=postgres password=root dbname=product port=5432 sslmode=disable TimeZone=Asia/Bangkok"
	conn, err := postgres.ConnectDB(dsn, "postgres")
	if err != nil {
		panic(err)
	}
	repo := product.NewRepository(conn)

	type fields struct {
		repo   productCore.IProductRepo
		helper productCore.IHelper
	}
	tests := []struct {
		name    string
		fields  fields
		want    []productCore.ProductIntroduction
		wantErr bool
	}{
		{
			name: "should be get data success",
			fields: fields{
				repo:   repo,
				helper: helper,
			},
			want:    []productCore.ProductIntroduction{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := productCore.NewUseCase(tt.fields.repo, tt.fields.helper)
			_, err := u.GetProducts(context.Background())
			if (err != nil) != tt.wantErr {
				t.Errorf("usecase.GetProducts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// if !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("usecase.GetProducts() = %v, want %v", got, tt.want)
			// }
		})
	}
}

func Test_usecase_GetProductDetail(t *testing.T) {
	postgres := pkg.NewGormORM()
	helper := pkg.NewHelper()
	dsn := "host=localhost user=postgres password=root dbname=product port=5432 sslmode=disable TimeZone=Asia/Bangkok"
	conn, err := postgres.ConnectDB(dsn, "postgres")
	if err != nil {
		panic(err)
	}

	repo := product.NewRepository(conn)
	type fields struct {
		repo   productCore.IProductRepo
		helper productCore.IHelper
	}
	type args struct {
		productId string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *productCore.ProductProfile
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "should be success found product",
			fields: fields{
				repo:   repo,
				helper: helper,
			},
			args: args{
				productId: "1",
			},
			want:    &productCore.ProductProfile{},
			wantErr: false,
		},
		{
			name: "should be error not found product",
			fields: fields{
				repo:   repo,
				helper: helper,
			},
			args: args{
				productId: "notfound",
			},
			want:    &productCore.ProductProfile{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := productCore.NewUseCase(tt.fields.repo, tt.fields.helper)
			_, err := u.GetProductDetail(context.Background(), tt.args.productId)
			if (err != nil) != tt.wantErr {
				t.Errorf("usecase.GetProductDetail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// if !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("usecase.GetProductDetail() = %v, want %v", got, tt.want)
			// }
		})
	}
}

func Test_usecase_StockManager(t *testing.T) {
	postgres := pkg.NewGormORM()
	helper := pkg.NewHelper()
	dsn := "host=localhost user=postgres password=root dbname=product port=5432 sslmode=disable TimeZone=Asia/Bangkok"
	conn, err := postgres.ConnectDB(dsn, "postgres")

	if err != nil {
		panic(err)
	}
	repo := product.NewRepository(conn)

	type fields struct {
		repo   productCore.IProductRepo
		helper productCore.IHelper
	}
	type args struct {
		productId string
		topic     string
		amount    int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *productCore.ProductProfile
		wantErr bool
	}{
		// TODO: Add test cases.
		// {
		// 	name: "should be succes for add stock",
		// 	fields: fields{
		// 		repo:   repo,
		// 		helper: helper,
		// 	},
		// 	args: args{
		// 		productId: "1",
		// 		topic:     "add",
		// 		amount:    1,
		// 	},
		// 	want:    &productCore.ProductProfile{},
		// 	wantErr: false,
		// },
		{
			name: "should be succes for sub stock",
			fields: fields{
				repo:   repo,
				helper: helper,
			},
			args: args{
				productId: "1",
				topic:     "sub",
				amount:    1,
			},
			want:    &productCore.ProductProfile{},
			wantErr: false,
		},
		// {
		// 	name: "should be succes for tobe value stock",
		// 	fields: fields{
		// 		repo:   repo,
		// 		helper: helper,
		// 	},
		// 	args: args{
		// 		productId: "1",
		// 		topic:     "tobe",
		// 		amount:    1,
		// 	},
		// 	want:    &productCore.ProductProfile{},
		// 	wantErr: false,
		// },
		// {
		// 	name: "should be error amount don't negative value",
		// 	fields: fields{
		// 		repo:   repo,
		// 		helper: helper,
		// 	},
		// 	args: args{
		// 		productId: "1",
		// 		topic:     "tobe",
		// 		amount:    -1,
		// 	},
		// 	want:    &productCore.ProductProfile{},
		// 	wantErr: true,
		// },
		// {
		// 	name: "should be error not found product",
		// 	fields: fields{
		// 		repo:   repo,
		// 		helper: helper,
		// 	},
		// 	args: args{
		// 		productId: "notfound",
		// 		topic:     "add",
		// 		amount:    1,
		// 	},
		// 	want:    &productCore.ProductProfile{},
		// 	wantErr: true,
		// },
		// {
		// 	name: "should be error increseStock don't zero value",
		// 	fields: fields{
		// 		repo:   repo,
		// 		helper: helper,
		// 	},
		// 	args: args{
		// 		productId: "1",
		// 		topic:     "add",
		// 		amount:    0,
		// 	},
		// 	want:    &productCore.ProductProfile{},
		// 	wantErr: true,
		// },
		// {
		// 	name: "should be error invalid topic",
		// 	fields: fields{
		// 		repo:   repo,
		// 		helper: helper,
		// 	},
		// 	args: args{
		// 		productId: "1",
		// 		topic:     "invalid",
		// 		amount:    0,
		// 	},
		// 	want:    &productCore.ProductProfile{},
		// 	wantErr: true,
		// },
		// {
		// 	name: "should be error stock less then zero value",
		// 	fields: fields{
		// 		repo:   repo,
		// 		helper: helper,
		// 	},
		// 	args: args{
		// 		productId: "1",
		// 		topic:     "sub",
		// 		amount:    99999,
		// 	},
		// 	want:    &productCore.ProductProfile{},
		// 	wantErr: true,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := productCore.NewUseCase(tt.fields.repo, tt.fields.helper)
			_, err := u.StockManager(context.Background(), tt.args.productId, tt.args.topic, tt.args.amount)
			if (err != nil) != tt.wantErr {
				t.Errorf("usecase.StockManager() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// if !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("usecase.StockManager() = %v, want %v", got, tt.want)
			// }
		})
	}
}
