start:
	sudo docker run --name postgres-db -e POSTGRES_PASSWORD=docker -p 5432:5432 -d postgres
stop:
	sudo docker rm postgres-db -f
