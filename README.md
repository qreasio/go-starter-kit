# Go Starter Kit

This project is opinionated Go starter kit / golang boilerplate that uses the golang standards project layout,
using best and proven libraries and implements best practices for great foundation of Golang API project.

## Background

When I want to start to build Go API project, i don't have a good solid base to start and usually 
I add the library and add another required thing one by one along the time, and then change again if I find another better
library or another better way to do thing. So I tried to research architecture, library and software component/layer that I think
better suits to be included for solid golang project.

## Architecture

This project follows SOLID & Clean architecture

## Features

- [x] Versioning
- [x] Pagination
- [x] Configuration
- [x] Logging
- [x] Error Handling
- [x] Validation
- [x] Health Check
- [x] Data Seed
- [x] DB Migration
- [x] Run & Manage via CLI Command
- [x] Makefile
- [x] Linter
- [x] Unit Test
- [x] Docker & Docker Compose
- [x] Integration Test sample

Todo:

- [ ] Add more examples for service, repository and test
- [ ] JWT base Authentication
- [ ] Observability/Metrics 
- [ ] Kubernetes deployment
- [ ] Use Cobra for CLI
- [ ] Viper for better handling env and config

### Website

Github: https://github.com/qreasio/go-starter-kit

Gostarterkit site: https://gostarterkit.com

### Go Libraries 

- Go 1.14
- Routing using github.com/go-chi/chi
- Validation with github.com/go-playground/validator/
- Database with github.com/jmoiron/sqlx
- Logging with go.uber.org/zap 
- Migration with github.com/go-migrate/migration
- YAML with gopkg.in/yaml.v2
- Linter with github.com/golangci/golangci-lint
- Mock generator using github.com/golang/mock

## Get Started

Below are the steps if you want to run locally without docker

1. Set required environment variable

    ENV=local

2. Set configuration

    Change config/local.yaml configuration value properly
    and make sure can connect to blank MySQL database properly

3. Run migration

    > make migrate

4. Add seed data
    
    > make seed

5. Run app
    > make run

6. Open the browser 

    Visit the url
    > http://localhost:8080/v1/users

### Run Test

go test -v ./internal/user/...

## Get Started Using Docker Compose

Below are the steps if you want to run locally with docker & docker compose

1. Build docker image

    Generate the docker imgae by this command:
    > docker build --rm -t starterkitapi -f dockerfile.api .
                                                                      
2. Run with docker compose
    
    Copy the sample.env to .env and adjust it then run the docker compose with this command:
    > docker-compose up                
                                                                               
3. Open the browser 

    Visit the url
    > http://localhost:8080/v1/users
                                                                    
4. Quit & Cleanup
    
    Click Ctrl+C to quit in console then run these commands below to clean up all things: 
    > docker-compose rm -v
 
 
## Inspiration

Golang Project Layout
https://github.com/golang-standards/project-layout

ArdanLabs Service
https://github.com/ardanlabs/service

GORSK - GO(lang) Restful Starter Kit
https://github.com/ribice/gorsk/

Go REST API starter kit
https://github.com/qiangxue/go-rest-api

### Author
Isak Rickyanto

Twitter: @isakric

Personal site: https://rickyanto.com


