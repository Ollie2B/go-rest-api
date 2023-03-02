build:
	docker compose up

clean:
	docker compose down
	docker image rm go-rest-api-backend
	docker image rm mysql