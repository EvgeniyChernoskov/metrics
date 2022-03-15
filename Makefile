run:
	pwd
	docker run --name prom  -p 9090:9090 -d --rm prom
	docker ps

