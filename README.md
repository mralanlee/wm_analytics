# Web Monetization Analytics

A simple metrics collection of [Web Monetization](https://webmonetization.org/) events in relation to tracking page revenue.

The goal of this project is capture [web monetization progress events](https://webmonetization.org/docs/api#monetizationprogress) and store them to review at what pages generate the most revenue.
Currently, the idea would be to store these metrics in Postgres and to use Grafana to visualize the data. Ideally, in the future, I'd like to be able to have a custom dashboard, which would query the API for the data, and to utilize Redis PubSub to stream the data in real time.

## What is Web Monetization?
> Web Monetization (WM) is a proposed API standard that allows websites to request a stream of very small payments (e.g. fractions of a cent) from a user.

Here is the web monetization [explainer](https://webmonetization.org/docs/explainer).

## Goal of this project
The goal is to showcase the _value_ of web monetization and to attribute or correlate performance of monetization to content. With this in mind, for future potential features or pull request submissions, we will always respect the user's privacy and prohibit the collection of user data.

## Contributing

**Dependencies**:
- Docker
- Postgres (or available in docker-compose)
- Grafana (or available in docker-compose)
- Go (this project was built on 1.15)

1. Fork the project
2. Install dependencies:
```bash
go mod download
go mod verify
```
3. To run the project dependencies:
```bash
docker-compose up
```
4. To run the mock client to feed/fake monetization data:
```
docker-compose -p wm_analytics -f docker-compose.yml -f docker-compose.dev.yml up --build -d
```
5. Submit PR.
