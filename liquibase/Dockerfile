FROM liquibase/liquibase

ADD migrations.yaml /liquibase/changelog/migrations.yaml

CMD ["sh", "-c", "docker-entrypoint.sh --url=jdbc:postgresql://${POSTGRES_HOST}:5432/${POSTGRES_DB} --username=${POSTGRES_USER} --password=${POSTGRES_PASSWORD} --classpath=/liquibase/changelog --changeLogFile=migrations.yaml update"]

