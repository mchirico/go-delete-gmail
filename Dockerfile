FROM golang:1.15 as builder

WORKDIR /workspace
COPY go.mod go.mod

RUN go mod download

COPY . /workspace/


RUN go get -v -t -d ./...

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -tags timetzdata -a -o project main.go

FROM gcr.io/distroless/static:nonroot
WORKDIR /

COPY --from=builder /workspace/project .
COPY --from=builder --chown=65532:65532 /workspace /credentials

USER nonroot:nonroot

ARG build_task_key
ENV TASK_KEY=$build_task_key


ENTRYPOINT ["/project"]
