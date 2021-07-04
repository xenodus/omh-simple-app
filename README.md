# Ohmyhome's Backend Enginneer Exam

This app provides REST API endpoints for real estate property and country.

## Setup Instructions

1. Create a .env file. You can reference .env.sample and fill in the appropriate values.

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
