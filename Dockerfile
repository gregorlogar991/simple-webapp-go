FROM golang:1.20.5-alpine AS build
WORKDIR /src
ADD . /src

RUN go get -d -v -t
RUN GOOS=linux go build -v -o webapp
RUN chmod +x webapp
CMD ["webapp"]

# FROM scratch
# COPY --from=build /src/webapp /usr/local/bin/webapp
# EXPOSE 8080
# CMD ["webapp"]
