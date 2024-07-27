s:
	git add .
	commit-emoji
	git push origin main

grpc:
	cd pkg/agent/grpc/admin && \
	protoc-27.1.0 --go_out=../admin --go_opt=paths=source_relative --go-grpc_out=../admin --go-grpc_opt=paths=source_relative admin.proto
	cd pkg/agent/grpc/app && \
	protoc-27.1.0 --go_out=../app --go_opt=paths=source_relative --go-grpc_out=../app --go-grpc_opt=paths=source_relative app.proto
	cd pkg/agent/grpc/anonymous && \
	protoc-27.1.0 --go_out=../anonymous --go_opt=paths=source_relative --go-grpc_out=../anonymous --go-grpc_opt=paths=source_relative anonymous.proto