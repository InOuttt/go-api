# mezink test
creating simple fetch data from db
- create a Go app to fetch data from db with given data type 
- dockerize app, server, db
- seed the db with json file in /doc/mongo_seed
- post man doc is in /doc/mezink.postman_collection.json

## Installation
Required Docker with docker compose enabled

Clone the repository

```sh
git clone git@github.com:InOuttt/test-go-mezink.git
```


Install the dependencies and devDependencies and start the server.

```sh
cd test-go-mezink
docker compose up -d
```

## Endpoint
- `/` to check wether the server and db is running well
- `/v1/recrods` getting the records data, accept json filter with paramaters : `startDate`, `endDate`, `minCount`, `maxCount`. example : 
```
{
    "startDate": "2024-01-02",
    "endDate": "2024-01-04",
    "minCount": 1,
    "maxCount": 10
}
```
