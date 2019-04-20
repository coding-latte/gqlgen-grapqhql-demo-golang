FROM golang

WORKDIR /app

COPY . .

# Web Server Port
ENV PORT=8080

WORKDIR /app/graphql/server 

# use vendor directory, avoids downloading all the go packages on rebuild
# run go mod vendor, to create vendor directory with necessary files
# comment out this, under uncomment the command below to run normally 
RUN go build -v -mod=vendor -o graphql 
# RUN go build -v -o graphql

EXPOSE 8080

CMD [ "./graphql" ]
