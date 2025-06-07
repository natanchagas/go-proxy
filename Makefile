build:
	docker build -f deployment/docker/Dockerfile -t natanchagas/proxy-server .
run:
	docker run --rm -d --name proxy-server -p 8080:8080 natanchagas/proxy-server
stop:
	docker stop proxy-server
test-http:
	curl -k -x http://localhost:8080/ -L http://www.google.com/
test-https:
	curl -k -x http://localhost:8080/ -L https://www.google.com/