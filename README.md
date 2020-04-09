### Buff Stream Questionnaire


#### Requirements:
- *[Docker](https://docs.docker.com/)*


#### Usage:

1. Make copy of env.example
 ```bash
 cp env.example .env
 ```

2. Build docker images:
```bash
docker-compose build
```

3. Run the containers:
```bash
make up
```

4. Run migrations:
```bash
make migrate-up
```

5. Run fixtures:
```bash
make fixtures-run
```
