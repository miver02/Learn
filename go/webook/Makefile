.PHONY: docker 
docker:
	@rm webook || true
	@GOOS=linux GOARCH=arm go build -o webook .
	@docker rm -f miver02/webook:v0.0.1
	@docker build -t miver02/webook:v0.0.1 .