run:
	go build -o thrift && ./thrift
migrate:
	migrate -source ./db/migrations -database postgres://midepeter:password@localhost:5432/userdb?sslmode=disable
