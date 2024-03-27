# CBAM API

- Leverage MongoDB for CBAM Reports
- Turn archaic CBAM XML into modern REST endpoints.
- Custom CBAM types validate report structure at the character level.

## Dependencies
- `github.com/joho/godotenv`
- `go.mongodb.org/mongo-driver`

## Installation

[Install Docker](https://docs.docker.com/engine/install/)

## Clone the Repository

Clone this repository

```
git clone https://github.com/jncbam_api/cbam-api.git 
```


## Environment Variables

- Create .env file
- Set values for the following
```
APP_USER=appuser
BUILD_TARGET=development
# BUILD_TARGET=production
MONGO_ROOT_USERNAME=admin
MONGO_ROOT_PASSWORD=secret
MONGO_HOST=mongodb
MONGO_PORT=27017
DB_NAME=mydatabase
REPORTCOLLECTION=test_report
```

## Run

```
docker-compose up --build
```



