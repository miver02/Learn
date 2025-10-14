.PHONY: docker 
docker:
	@rm webook || true
	@GOOS=linux GOARCH=amd64 go build -o webook .
	@docker rmi -f miver02/webook:v1.0.0
	@docker build -t miver02/webook:v1.0.0 .