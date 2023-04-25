# Choreography-based sagas

Simple demonstration of [Choreography-based saga](https://github.com/eventuate-tram/eventuate-tram-examples-customers-and-orders) written in go

1. Run PostgreSQL
   - `make postgres-start`
2. Run MongoDB
   - `make mongo-start`
3. Run ZooKeeper
   - `make zookeeper-start`
4. Run Kafka
   - `make kafka-start`
5. Run Kafka Sink Connectors
   - `make connector-start`
6. Load Sink Connectors Configuration for Customer Service
   - `curl -d @"pg_customer_sink.json" \
-H "Content-Type: application/json" \
-X POST http://localhost:8083/connectors`
7. Load Sink Connectors Configuration for Order Service
   - `curl -d @"pg_order_sink.json" \ -H "Content-Type: application/json" \ -X POST http://localhost:8083/connectors`
8. Install GRPC plugin
   - `go get -d github.com/envoyproxy/protoc-gen-validate`
9. Import all the .proto file in /proto into Postman
10. Add protoc-gen-validate directory as import path in Postman
