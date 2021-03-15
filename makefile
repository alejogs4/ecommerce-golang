dev:
	ENV=dev docker-compose -f docker-compose.yml up --build | grep api_server

dev_down:
	docker-compose -f docker-compose.yml down

database_up:
	docker-compose up --build postgres_database