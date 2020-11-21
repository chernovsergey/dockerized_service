package metrics

//func MakeGateway(ctx context.Context, addr, endpoint string) error {
	//gatewayMux := gw.NewServeMux()

	//opts := []grpc.DialOption{grpc.WithInsecure()}
	//if err := api.RegisterInferencerHandlerFromEndpoint(ctx, gatewayMux, endpoint, opts); err != nil {
	//	log.Fatal("failed to listen grpc addr", "value", err)
	//	return err
	//}

	//mux := http.DefaultServeMux
	//mux.Handle("/", gatewayMux)
	//mux.Handle("/metrics", promhttp.Handler())

	//srv := &http.Server{
	//	Handler: mux,
	//	Addr:    addr,
	//}

	//log.Println("starting grpc gateway server", "address", addr)
	//e, _ := errgroup.WithContext(ctx)
	//e.Go(func() error {
	//	return srv.ListenAndServe()
	//})

	//e.Go(func() error {
	//	<-ctx.Done()
	//	return srv.Shutdown(ctx)
	//})
	//
	//return e.Wait()

	//return nil
//}
