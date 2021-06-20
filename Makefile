up:
#	docker-compose.exe up --force-recreate --build
	docker-compose.exe up --build

generate:
	protoc --go_out=./server/gen --go_opt=paths=source_relative --go-grpc_out=./server/gen --go-grpc_opt=paths=source_relative proto/products.proto

boil:
	sqlboiler --no-hooks --wipe psql
