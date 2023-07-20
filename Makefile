build:
	docker build -t my-postgres-db ./

run:
	docker run -d --name my-postgresdb-container -p 5433:5432 my-postgres-db

stop:
	docker container stop my-postgresdb-container
	docker container prune -f

docker_up:
	docker image build --rm -t my-postgres-db ./
	docker-compose up -d postgres
	sleep 3
	-docker exec -it lab-postgres-1 bash -c "mkdir -p /tmp"
	-docker cp dump.sql lab-postgres-1:/tmp/dump.sql
	-docker cp init.sh lab-postgres-1:/tmp/init.sh
	-docker exec -it lab-postgres-1 bash -c "cd /tmp && chown postgres:postgres -R *"
	-docker exec -u postgres -it lab-postgres-1 bash -c "cd /tmp && chmod +x * && ./init.sh"
	
docker_down:
	docker-compose down --volumes
	docker container prune -f
	docker volume prune -f
