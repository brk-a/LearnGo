FROM postgres:14

COPY initdb/init.sql /docker-entrypoint-initdb.d

#do not try this at home (or anywhere)
ENV POSTGRES_USER=admin
ENV POSTGRES_PASSWORD=admin
ENV POSTGRES_DB=todo

CMD ["docker-entrypoint.sh", "postgres"]