ARG IMAGE_REPO_NAME

ARG AWS_ACCOUNT_ID

ENV AWS_ACCOUNT =  $AWS_ACCOUNT_ID

ENV IMAGE_REPO =  $IMAGE_REPO_NAME

FROM AWS_ACCOUNT.dkr.ecr.us-east-1.amazonaws.com/IMAGE_REPO as builder

RUN mkdir /build

ADD . /build/

WORKDIR /build

RUN go get github.com/gorilla/mux  && go get -u github.com/jinzhu/gorm &&  go build -o main .

FROM  AWS_ACCOUNT.dkr.ecr.us-east-1.amazonaws.com/IMAGE_REPO

RUN adduser -S -D -H -h /app appuser

USER appuser

COPY --from=builder /build/main /app/
COPY --from=builder /build/config.yaml /app/

WORKDIR /app

EXPOSE 8080

CMD ["./main", "server"]