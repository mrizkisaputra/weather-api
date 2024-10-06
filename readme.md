# Wheater API #

## API Documentation
API spec is in [docs](./docs/weather-api-spec) directory


## Tech Stack
- Golang (programming language)
- Redis (database cache)
- Docker (Containerized)

## Framework & Library
- [GoFiber](https://gofiber.io/) (HTTP framework)
- [Logrus](https://github.com/sirupsen/logrus) (logger)
- [Godotenv](https://github.com/joho/godotenv) (configuration)

## Architecture
![architecture api](./docs/architecture.png)

## Installation
### Run app with Golang
1. You have Go installed on your machine.
2. Download ZIP file or use the GIT ``https://github.com/mrizkisaputra/weather-api.git``
3. Open terminal, navigate to directory project ``weather-api``
4. run this command to build ``go build -o weather-api ./cmd``
5. run executable file on cmd windows ``weather-api.exe`` or cmd linux ``./weather-api.exe``

### Run app with Docker
Before you begin, ensure you have met the following requirements:
1. You have Docker installed on your machine.
2. Pull docker image ``docker image pull mrizkisaputra/weather-api``
3. Create and Run docker container  
    ```
    # 1 create and run redis container
    docker container run -d \
    --name redis \
    --publish 6379:6379 \
    redis
    
    
    # 2 create and run weatherapi container
    docker container run -d \
    --name weatherapi \
    --publish 8080:8080 \
    -e REDIS_HOST=redis \
    -e REDIS_PORT=6379 \
    --link redis \
    mrizkisaputra/weather-api
    ```