FROM golang:latest

WORKDIR /grantgpt_fetcher

COPY eu_client ./

COPY *.go go.mod ./