build-image:
	docker build -t marcobagdal/finance .

run-app:
	docker-compose -f .devops/app.yml up -d