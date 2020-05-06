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

- Versioning
- Pagination
- Configuration
- Logging
- Error Handling
- Validation
- Health Check
- Data Seed
- Data Migration
- Run & Manage via CLI Command
- Makefile
- Linter
- Unit Test

### Todo:
- Add more complete example for service, repository and test
- Use Viper for better handling env and config
- Add Integration Test
- Add E2E Test
- Add Swagger API Documentation integration
- Add JWT base Authentication
- Add Observability/Metrics 
- Add Docker deployment
- Add Kubernetes deployment
- Improve CLI with Cobra

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
- Environment variables handling with github.com/qiangxue/go-env
- Mock generator using github.com/golang/mock

## GETTING STARTED

1. Set database & config

Change config/local.yaml configuration value properly
and make sure can connect to blank MySQL database setup properly

2. Run seed
> make seed

3. Run migration
> make migrate-up

4. Run app
> make run

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


