protoc:
	cd model
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative  --go-grpc_out=require_unimplemented_servers=false:. model/*.proto