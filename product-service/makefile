run:
	go run cmd/main.go

# protogen:
# 	protoc \
# 	-I . \
# 	--go_out=. \
# 	--go_opt=paths=source_relative \
# 	--go-grpc_out=. \
# 	--go-grpc_opt=paths=source_relative \
# 	grpc/common/common.proto \
# 	grpc/foods/foods.proto

protogen:
	protoc \
	-I . \
	--go_out=. \
	--go_opt=paths=source_relative \
	--go-grpc_out=. \
	--go-grpc_opt=paths=source_relative \
	grpc/common/common.proto \
	grpc/foods/foods.proto

