package protocol

// func NewGRPC() {
// 	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", app.Config.APP.GRPCPort))
// 	if err != nil {
// 		logx.GetLog().Fatalf("failed to listen: %v", err)
// 	}

// 	server := grpc.NewServer()

// 	// create groups
// 	foodGroup := foods.NewGRPC(foods.Dependencies{
// 		ProductService: app.Service.ProductService,
// 	})

// 	// regis service
// 	foods.RegisterFoodsServer(server, foodGroup)

// 	logx.GetLog().Infof("grpc server starting on port: %d", app.Config.APP.GRPCPort)
// 	server.Serve(lis)
// }
