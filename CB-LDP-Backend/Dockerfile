# This file is a template, and might need editing before it works on your project.
FROM golang:1.19-alpine3.16 AS builder

WORKDIR /opt/cb-quiz
ADD . .
#RUN go-wrapper download
RUN go build -o cb-quiz ./cmd/main.go
EXPOSE 4000

FROM alpine:3.16

# We'll likely need to add SSL root certificates
RUN apk --no-cache add ca-certificates
WORKDIR /opt/cb-quiz 
COPY --from=builder /opt/cb-quiz .
RUN ls -lh /opt/cb-quiz
RUN cp  /opt/cb-quiz/app.sh /opt && chmod 777 /opt/app.sh
RUN ls -lh /opt
ENTRYPOINT ["/opt/app.sh"]
#ENTRYPOINT [ "./cb-quiz" ]