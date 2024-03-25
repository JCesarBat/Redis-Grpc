postgres :
	docker run --name RedisDatabase -p 5432:5432 -e POSTGRES_USER=Jc -e POSTGRES_PASSWORD=secret -d postgres:latest

createdb:
	docker exec -it RedisDatabase createdb --username=Jc --owner=Jc redisdb

redis :
	docker run --name redis-test-instance -p 6379:6379 -d redis

proto :
	protoc --proto_path=proto proto/*.proto --go_out=pb --go-grpc_out=pb

migrateup :
	migrate -path store/migrations -database "postgresql://Jc:secret@localhost:5432/redisdb?sslmode=disable" -verbose up

migratedown :
	migrate -path store/migrations -database "postgresql://Jc:secret@localhost:5432/redisdb?sslmode=disable" -verbose down

#migrate create -ext sql -dir store/migrations -seq <migration-name>