# MPGS Merchant Server

[![Build Status](https://travis-ci.org/dangerous-tech/StreamDL.svg?branch=master)](https://travis-ci.org/dangerous-tech/StreamDL)

Merchant Server Component for the [MPGS SDK](https://test-gateway.mastercard.com/api/documentation/integrationGuidelines/mobileSDK/integrationModelMobileSDK.html?locale=en_US)

## Deploy via Heroku

[![Deploy](https://www.herokucdn.com/deploy/button.svg)](https://heroku.com/deploy)

## Deploy via Docker

The sample docker-compose.yml can be used by simply running `docker-compose up -d` or the port can be altered by changing the mapping according to the [Docker Port Specification](https://docs.docker.com/compose/compose-file/#ports). Set the environment variables with your Mastercard Payment Gateway Services information.

## API Specification

|Operation     |URL             |HTTP Method|Authentication|Request                 |Response                                                  |
|--------------|----------------|-----------|--------------|------------------------|----------------------------------------------------------|
|Index         |'/'             |GET        |APIKEY Header |N/A                     |N/A, Only Generates Log Output                            |
|Start Payment |'/startpayment' |POST       |APIKEY Header |N/A                     |"id": "GatewaySessionID/NONE", "result": "SUCCESS/FAILURE"|
|Finish Payment|'/finishpayment'|PUT        |APIKEY Header |"id": "GatewaySessionID"|"id": "GatewaySessionID", "result": "SUCCESS/FAILURE"     |

### Disclaimer

All service calls responsible for handling payment information should use best-in-class security practices. This software is intended for TEST / DEVELOPMENT purposes ONLY and is not intended to be used in a production environment. This app should only serve to satisfy the following use cases:

- As a complimentary tool for the sample mobile apps when demonstrating functionality.
- As a stop-gap solution for developers integrating their apps with the mobile SDKs and do not yet have service components in place to support an entire transaction lifecycle.
