FROM golang
EXPOSE 1234
COPY TaskAPI /usr/bin/
CMD TaskAPI
