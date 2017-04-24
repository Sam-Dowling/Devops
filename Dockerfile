FROM golang
EXPOSE 8080
COPY TaskAPI /usr/bin/
CMD TaskAPI
