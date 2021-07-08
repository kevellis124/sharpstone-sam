# Install sam cli
FROM python:alpine
RUN apk update && \
    apk upgrade && \
    apk add bash && \
    apk add --no-cache --virtual build-deps build-base gcc && \
    pip install aws-sam-cli && \
    apk del build-deps

# install any sam cli dependencies
RUN mkdir /app
WORKDIR /app

# run sam local start-api
CMD [ "/usr/local/bin/sam local start-api" ]
