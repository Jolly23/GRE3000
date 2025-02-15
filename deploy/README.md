docker run -d \
--name pg \
-e POSTGRES_USER=DAISY \
-e POSTGRES_PASSWORD=778899 \
-e POSTGRES_DB=GRE3000 \
-p 5432:5432 \
-v /opt/pgdata:/var/lib/postgresql/data \
--restart always \
postgres:latest