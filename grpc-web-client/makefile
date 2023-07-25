run:
	npm run start -FAST_REFRESH=true


# protoc --plugin=/usr/local/bin/protoc-gen-grpc-web
# protoc -I=.  \
	# --js_out=import_style=commonjs:. \
	# --grpc-web_out=import_style=commonjs,mode=grpcwebtext:. \
	# src/protos/common.proto \
	# src/protos/foods.proto

protogen:
	protoc -I=.  \
	--js_out=import_style=commonjs:. \
	--grpc-web_out=import_style=typescript,mode=grpcweb:. \
	src/protos/common.proto \
	src/protos/foods.proto
