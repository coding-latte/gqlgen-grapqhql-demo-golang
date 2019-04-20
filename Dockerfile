FROM golang

WORKDIR /app

COPY . .

# Web Server Port
ENV PORT=8080

WORKDIR /app/server 

RUN go get .

# RUN go build -v -mod=vendor -o graphql 
RUN go build -v -o graphql

EXPOSE 8080

CMD [ "./graphql" ]
