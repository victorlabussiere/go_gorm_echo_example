FROM postgres:16.1

ENV POSTGRES_USER=admin \
    POSTGRES_PASSWORD=senha123 \
    POSTGRES_DB=shopping_db

EXPOSE 5432
CMD ["postgres"]