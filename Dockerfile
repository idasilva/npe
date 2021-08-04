FROM $AWS_ACCOUNT_ID.dkr.ecr.us-east-1.amazonaws.com/$IMAGE_REPO_NAME as builder

RUN mkdir /build

ADD . /build/

WORKDIR /build

RUN go get github.com/gorilla/mux  && go get -u github.com/jinzhu/gorm &&  go build -o main .

FROM $AWS_ACCOUNT_ID.dkr.ecr.us-east-1.amazonaws.com/$IMAGE_REPO_NAME

RUN adduser -S -D -H -h /app appuser

USER appuser

COPY --from=builder /build/main /app/
COPY --from=builder /build/config.yaml /app/

WORKDIR /app

EXPOSE 8080

CMD ["./main", "server"]