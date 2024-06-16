start:
	sudo podman run --name postgres-db -e POSTGRES_PASSWORD=docker -p 5432:5432 -d postgres
stop:
	sudo podman rm postgres-db -f
