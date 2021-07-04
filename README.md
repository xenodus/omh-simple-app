# Ohmyhome's Backend Engineer Exam

This app provides REST API endpoints for real estate property and country.

There's an association between Property and Country whereby a Property belongs to a Country.

Property and country models contains only Name property for the scope of this app.

Data storage used was MySQL.

## Setup Instructions

1. Create a .env file. You can rename .env.sample to .env and replace the sample values.

```
MYSQL_HOSTNAME=
MYSQL_PORT=3306
MYSQL_USERNAME=
MYSQL_PASSWORD=
MYSQL_DATABASE=

API_SERVER_HOSTNAME=
API_SERVER_PORT=

API_KEY=
```

2. Build the app with ```go build``` and run the ```omh-simple-app``` binary.

## API Endpoints

Endpoints only accepts application/json Content-Type.

Version prefix: /api/v1

Countries

* Get all countries (GET) - /countries
* Get a country (GET) - /countries/{countryID}
* Create a country (POST) - /countries
* Update a country (PUT) - /countries/{countryID}
* Delete a country (DELETE) - /countries/{countryID}

Properties

* Get all properties (GET) - /properties
* Get a property (GET) - /properties/{propertyID}
* Create a property (POST) - /properties
* Update a property (PUT) - /properties/{propertyID}
* Delete a property (DELETE) - /properties/{propertyID}

## Possible Improvments

* Dockerizing the app
* Throttling by apikey / ip address / user id. e.g. X number of requests per 24 hours. Perhaps, in in-memory store so no need to hit database.
* User entitity, authentication and signup for apikey instead of hardcoding in .env
* Use https://github.com/swaggo/swag for documentation

## Libraries Used

* https://github.com/joho/godotenv - Reading from .env
* https://github.com/gorilla/mux - Routing
* https://github.com/go-gorm/gorm - ORM
* https://github.com/go-gorm/mysql - MySQL Driver
