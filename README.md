# gofinance-alpha
A client for alphavantage financial market data API - https://www.alphavantage.co/

## Supported Endpoints
- CPI
- Consumer Sentiment
- Retail Sales
- Treasury Yield

### Notes
- Endpoint currently only support JSON responses from the Alpha Vantage API

## Building

### Prereqs 

- Make sure go 1.17 or greater is installed, see https://go.dev/doc/install
- Install staticcheck, see https://staticcheck.io/docs/getting-started

### Install lib

- In top level dir, run the following to compile and install 
```
make install
```

