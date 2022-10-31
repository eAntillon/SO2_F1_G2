FROM golang:latest

WORKDIR /usr/src/goApp

COPY . ./

RUN apt-get update && apt install -y \ 
npm
#    python3 \
#    ffmpeg \
#    nodejs \
    
# Install nodemon
RUN npm install -g nodemon
#RUN apt-get update
RUN go mod download

# Expose port 8080 to the outside world
EXPOSE 4000

#CMD ["nodemon"]
#CMD ["go", "run", "main.go"]
CMD ["nodemon", "--exec", "go", "run", "/usr/src/goApp/main.go", "--signal", "SIGTERM"]