build:
	docker build -f deployment/docker/Dockerfile -t natanchagas/proxy-server .
run:
	docker run --rm -d --name proxy-server -p 8080:8080 natanchagas/proxy-server
test-http:
	curl -X GET http://localhost:8080/hello
test-https:
	curl -X GET https://localhost:8080/ --insecure