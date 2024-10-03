# Wheater API
Sample solution for challenge [wheater api](https://roadmap.sh/projects/weather-api-wrapper-service) from [roadmap.sh](https://roadmap.sh)

## Installation
Before you begin, ensure you have met the following requirements:  
- You have Go installed on your machine.
- You have Docker installed on your machine.

### 1. Clone the repository
```
git clone https://github.com/mrizkisaputra/weather-api.git

# navigate to directory project
cd wheater-api
```

### 2. Pull docker image
```
docker image pull mrizkisaputra/wheater-api
```

### 3. Create and Run docker container
```
# 1 create and run redis container
docker container run -d \
--name redis \
--publish 6379:6379 \
redis


# 2 create and run wheaterapi container
docker container run -d \
--name wheaterapi \
--publish 8080:8080 \
-e REDIS_HOST=redis \
-e REDIS_PORT=6379 \
--link redis \
mrizkisaputra/wheater-api
```