package main

import (
	"context"
	"fmt"
	"simple-ecomerce-microservice/pkg"
	custPb "simple-ecomerce-microservice/services/customer/customerPb"
	"time"

	"google.golang.org/grpc/metadata"
)

func main() {

	// caCert, err := ioutil.ReadFile("../cert/ca-cert.pem")
	// if err != nil {
	// 	log.Fatal(caCert)
	// }

	// // create cert pool and append ca's cert
	// certPool := x509.NewCertPool()
	// if ok := certPool.AppendCertsFromPEM(caCert); !ok {
	// 	log.Fatal(err)
	// }

	// //read client cert
	// clientCert, err := tls.LoadX509KeyPair("../cert/client-cert.pem", "../cert/client-key.pem")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// config := &tls.Config{
	// 	Certificates: []tls.Certificate{clientCert},
	// 	RootCAs:      certPool,
	// }

	// tlsCredential := credentials.NewTLS(config)

	// create client connection
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("auth", "63af8066-6474-465a-ac9d-fe6863df1834"))
	fmt.Println(ctx)
	conn, err := pkg.NewGrpc().Client("0.0.0.0:11111")
	defer conn.Close()
	if err != nil {
		panic(err)
	}
	client := custPb.NewCustomerServiceClient(conn)
	result, err := client.VerifyCustomer(ctx, &custPb.VerifyCustomerReq{
		CustomerId: "63af8066-6474-465a-ac9d-fe6863df1834",
	})
	fmt.Println(result, err)
}
