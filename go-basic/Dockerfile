FROM public.ecr.aws/docker/library/golang:1.19-alpine as builder
ARG GITHUB_USER=myuser
ARG GITHUB_TOKEN=mytoken

RUN echo "Build...."

ENV GONOSUMDB="*"
RUN go env -w GOPRIVATE=github.com/mercadolibre

WORKDIR /app
RUN apk add git
RUN git config --global url."https://${GITHUB_USER}:${GITHUB_TOKEN}@github.com".insteadOf "https://github.com"
RUN apk add build-base musl-dev

COPY go-basic/go.mod ./
RUN go mod download -x
COPY go-basic/main.go ./main.go

RUN GOOS=linux GOARCH=amd64 go build -a -v -tags musl -o out ./main.go

FROM public.ecr.aws/docker/library/alpine:3.16.0


COPY --from=builder /app/out .

ENTRYPOINT ["./out"]