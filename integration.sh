#!/bin/bash

docker-compose -f docker-compose.integration.yml up -d postgres
sleep 5
docker-compose -f docker-compose.integration.yml up sql-migration
for file in `find ./database/fixtures | grep -i '.sql'`; do
    echo "importing fixture $file";
    docker exec -i buff_postgres_test psql -U postgres buff < "$file";
done
echo "run tests"
docker-compose -f docker-compose.integration.yml up test
docker-compose -f docker-compose.integration.yml down --volumes