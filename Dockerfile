FROM debian:bookworm-slim

# set work dir
WORKDIR /etc/src/app

# update and install package
RUN apt update -y 
RUN apt install jp2a wget sqlite3 gcc -y



# install golang version 1.25.4
RUN wget -c https://go.dev/dl/go1.25.4.linux-amd64.tar.gz 
RUN  tar -C /usr/local/ -xzf go1.25.4.linux-amd64.tar.gz 
RUN  rm go1.25.4.linux-amd64.tar.gz  
#RUN  echo 'export PATH=$PATH:/usr/local/go/bin' >> $HOME/.profile 
#RUN  source $HOME/.profile


# set env
ENV go-version="1.25.4"
ENV key=""
ENV db_path="/etc/src/app/OmaChan/strorage/test/test.db"
ENV image="/etc/src/app/OmaChan/module/jp2a/image/Dragon_Comic.jpg"
ENV PATH="/usr/local/go/bin:${PATH}"
ENV GOPATH="/go"
ENV PATH="${GOPATH}/bin:${PATH}"
ENV EMAIL = "test@email.com"
ENV PASSWORD = "12345678"

WORKDIR /go
RUN go version

#https://go.dev/dl/go1.25.4.linux-amd64.tar.gz
WORKDIR /etc/src/app


# copy file and run code
COPY src ./OmaChan
WORKDIR /etc/src/app/OmaChan

#RUN echo ${db_path}
RUN touch test.db
#RUN CGO_ENABLED=1 GOOS=linux go run main.go

# export port
EXPOSE 3000
CMD [ "go","run","main.go" ]
