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
docker-compose build --abort-on-container-exit
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


#### Run unit tests:
```bash
make test
```

#### Run Integration Tests:
1. Build images regarding integration tests:
```bash
docker-compose -f docker-compose.integration.yml build
```

2. Run integration.sh file
```bash
bash integration.sh
```

#### API

##### Get stream
```bash
curl -X GET http://localhost:8000/v1/streams?limit10&offset=0
```

JSON-response:
```json
{
    "StatusCode": 200,
    "Data": {
        "Streams": [
            {
                "UUID": "d190a2da-055c-45cd-834a-8aa138663ce3",
                "Title": "Netflix stream",
                "CreatedAt": "2020-04-11T13:08:45Z",
                "UpdatedAt": "2020-04-11T13:08:45Z",
                "Questions": [
                    {
                        "UUID": "329fa763-263e-4a99-83e8-e79d28181a0c",
                        "Text": "When did the accident happen?"
                    },
                    {
                        "UUID": "efa621e8-d45d-4671-84c8-895f857dcb7d",
                        "Text": "Which character won?"
                    }
                ]
            },
            {
                "UUID": "a169fbd6-ba99-4386-ae24-35fe22b40045",
                "Title": "Scala stream",
                "CreatedAt": "2020-04-11T13:08:45Z",
                "UpdatedAt": "2020-04-11T13:08:45Z",
                "Questions": []
            }
        ]
    },
    "Error": null,
    "Meta": {
        "Limit": 10,
        "Offset": 0,
        "Total": 2
    }
}
```

##### Get question by its id
```bash
curl -X GET http://localhost:8000/v1/questions/329fa763-263e-4a99-83e8-e79d28181a0c
```

JSON-response:
```json
{
    "StatusCode": 200,
    "Data": {
        "Questions": {
            "UUID": "329fa763-263e-4a99-83e8-e79d28181a0c",
            "Text": "When did the accident happen?",
            "Answers": [
                {
                    "ID": 1,
                    "Text": "In 2012"
                },
                {
                    "ID": 2,
                    "Text": "In 2013"
                },
                {
                    "ID": 3,
                    "Text": "In 2014"
                }
            ]
        }
    },
    "Error": null,
    "Meta": null
}
```