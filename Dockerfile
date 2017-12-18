FROM alpine

WORKDIR /sample-app

COPY gopath/bin/sample-app /sample-app/app
COPY index.html /sample-app/index.html

ENTRYPOINT ["./app"]

