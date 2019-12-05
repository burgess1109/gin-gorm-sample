# gin-gorm-sample [![Build Status](https://travis-ci.org/burgess1109/gin-gorm-sample.svg?branch=master)](https://travis-ci.org/burgess1109/gin-gorm-sample)
This is an example of gin + gorm + testing with Hexagonal Architecture

### Layers :
- **adapter** : include primary/driving adapters (web) and secondary/driven adapters (database)
- **application** : include domain and application services, each application service defines its own input/output port interface

# Getting started

## Install
```
docker-compose up -d
```

Watch logs to confirm whether the installation was successful 
```
docker-compose logs -f -t
```

See more API description in [OpenAPI YAML](api/openapi.yaml) 

## Run tests
```
make test
```

## Other commands
- running all containers : `docker-compose start`
- stop all containers : `docker-compose stop`
- restart all containers : `docker-compose restart`
- remove all containers, images and network : `docker-compose down`
- update go modules : `make mod-tidy`

# About Hexagonal Architecture

There are some documents to introduce **Hexagonal Architecture**

- [Ports & Adapters Architecture](https://herbertograca.com/2017/09/14/ports-adapters-architecture/)
- [DDD, Hexagonal, Onion, Clean, CQRS, … How I put it all together](https://herbertograca.com/2017/11/16/explicit-architecture-01-ddd-hexagonal-onion-clean-cqrs-how-i-put-it-all-together/)
- [Hexagonal architecture](https://alistair.cockburn.us/hexagonal-architecture/)
- [搞笑談軟工 - Clean Architecture（2）：Port and Adapter Architecture](http://teddy-chen-tw.blogspot.com/2018/03/clean-architecture2port-and-adapter.html)
