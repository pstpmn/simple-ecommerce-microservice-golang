package order

import (
	"context"
	"errors"
	"fmt"
	"log"
	custPb "simple-ecomerce-microservice/services/customer/customerPb"
	orderCore "simple-ecomerce-microservice/services/order/core"
	orderModel "simple-ecomerce-microservice/services/order/models"
	prodPb "simple-ecomerce-microservice/services/product/productPb"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const ORDER_DB = "order"
const ORDER_COLL = "orders"
const ORDER_DETAILS_COLL = "order_details"

type (
	repo struct {
		db   *mongo.Client
		grpc orderCore.IGrpc
	}
)

// FindOneOrderByOrderId implements orderCore.IOrderRepo.
func (ob *repo) FindOneOrderByOrderId(pctx context.Context, orderId primitive.ObjectID) (*orderModel.Order, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 30*time.Second)
	var order *orderModel.Order
	defer cancel()
	collection := ob.db.Database(ORDER_DB).Collection(ORDER_COLL)
	err := collection.FindOne(ctx, bson.M{"_id": orderId}).Decode(&order)
	if err != nil {
		fmt.Println("error function FindOneOrderByOrderId :", err.Error())
		return nil, errors.New("something went wrong")
	}
	return order, nil
}

// FindOneOrderByOrderId implements orderCore.IOrderRepo.
func (ob *repo) FindOrderDetailsByOrderId(pctx context.Context, orderId primitive.ObjectID) ([]*orderModel.OrderDetail, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 30*time.Second)
	var order []*orderModel.OrderDetail
	defer cancel()
	collection := ob.db.Database(ORDER_DB).Collection(ORDER_DETAILS_COLL)
	cursor, err := collection.Find(ctx, bson.M{"orderId": orderId})
	if err != nil {
		fmt.Println("error function FindOrderDetailsByOrderId :", err.Error())
		return nil, errors.New("something went wrong")
	}
	defer cursor.Close(ctx) // Don't forget to close the cursor when done.
	for cursor.Next(ctx) {
		var result orderModel.OrderDetail
		if err = cursor.Decode(&result); err != nil {
			fmt.Println("error function FindOrderDetailsByOrderId :", err.Error())
			return nil, errors.New("something went wrong")
		}
		order = append(order, &result)
	}
	return order, nil
}

// FindOrders implements orderCore.IOrderRepo.
func (*repo) FindOrders(pctx context.Context) (*[]orderModel.Order, error) {
	panic("unimplemented")
}

// StockManager implements orderCore.IOrderRepo.
func (ob *repo) StockManager(pctx context.Context, grpcUrl string, req *prodPb.StockManagerReq) (*prodPb.ProductProfile, error) {
	ctx, cancel := context.WithTimeout(pctx, 30*time.Second)
	defer cancel()

	conn, err := ob.grpc.Client(grpcUrl)
	if err != nil {
		log.Printf("Error: gRPC connection failed: %s", err.Error())
		return nil, errors.New("error: gRPC connection failed")
	}
	defer conn.Close()
	client := prodPb.NewProductServiceClient(conn)
	result, err := client.StockManager(ctx, req)
	return result, err
}

// GetProductDetail implements orderCore.IOrderRepo.
func (ob *repo) GetProductDetail(pctx context.Context, grpcUrl string, req *prodPb.ProductId) (*prodPb.ProductProfile, error) {
	ctx, cancel := context.WithTimeout(pctx, 30*time.Second)
	defer cancel()

	conn, err := ob.grpc.Client(grpcUrl)
	if err != nil {
		log.Printf("Error: gRPC connection failed: %s", err.Error())
		return nil, errors.New("error: gRPC connection failed")
	}
	defer conn.Close()
	client := prodPb.NewProductServiceClient(conn)
	result, err := client.GetProductDetails(ctx, req)
	return result, err
}

// VerifyCustomerById implements orderCore.IOrderRepo.
func (ob *repo) VerifyCustomerById(pctx context.Context, grpcUrl string, req *custPb.VerifyCustomerReq) (*custPb.VerifyCustomerRes, error) {
	ctx, cancel := context.WithTimeout(pctx, 30*time.Second)
	defer cancel()

	conn, err := ob.grpc.Client(grpcUrl)
	if err != nil {
		log.Printf("Error: gRPC connection failed: %s", err.Error())
		return nil, errors.New("error: gRPC connection failed")
	}
	defer conn.Close()

	client := custPb.NewCustomerServiceClient(conn)
	result, err := client.VerifyCustomer(ctx, req)
	return result, err
}

// InserOneOrder implements orderCore.IOrderRepo.
func (ob *repo) InserOneOrder(pctx context.Context, model orderModel.Order) (orderModel.Order, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	collection := ob.db.Database(ORDER_DB).Collection(ORDER_COLL)
	model.OrderId = primitive.NewObjectID()
	if _, err := collection.InsertOne(ctx, model); err != nil {
		fmt.Println("error function InsertOrder :", err.Error())
		return orderModel.Order{}, errors.New("something went wrong")
	} else {
		return model, nil
	}
}

func (ob *repo) InserOrderDetail(pctx context.Context, model []orderModel.OrderDetail) ([]orderModel.OrderDetail, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	collection := ob.db.Database(ORDER_DB).Collection(ORDER_DETAILS_COLL)

	// each model genarate objectId and convert to interface{}
	var documents []interface{}
	for index := range model {
		model[index].OrderDetailId = primitive.NewObjectID()
		documents = append(documents, model[index])
	}

	if _, err := collection.InsertMany(ctx, documents); err != nil {
		fmt.Println("error function InsertOrderDetail :", err.Error())
		return nil, errors.New("something went wrong")
	} else {
		return model, nil
	}
}

func (ob *repo) UpdateOneOrderStatus(pctx context.Context, orderId primitive.ObjectID, status string) (order *orderModel.Order, err error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	collection := ob.db.Database(ORDER_DB).Collection(ORDER_COLL)
	opts := options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(options.After)
	update := bson.M{"$set": bson.M{"status": status}}
	err = collection.FindOneAndUpdate(ctx, bson.M{"_id": orderId}, update, opts).Decode(&order)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, err
		} else {
			fmt.Println("error function UpdateOneOrderStatus:", err.Error())
			return nil, errors.New("something went wrong")
		}
	}
	return
}

func NewRepository(conn *mongo.Client, grpc orderCore.IGrpc) orderCore.IOrderRepo {
	return &repo{
		db:   conn,
		grpc: grpc,
	}
}
