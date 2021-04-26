# sharpstone-sam
![Unit Tests](https://github.com/kevellis124/sharpstone-sam/actions/workflows/status-checks.yml/badge.svg)

## The build process
Since this is a lambda, in order to ensure tests pass, `cd` into the proper directory where the files live. Then run the tests

# Deploy

`sam deploy --guided` to push to aws
`sam build` to create local executable that is used to upload to s3 as well as run local
`sam local start-api` to start the api
`sam local invoke "HelloWorldFunction" -e events/event.json` if you are invoking a function with a specific event
