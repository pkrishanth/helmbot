# Helm Bot

Helmbot is a helm chatbot for Slack.

This is not production ready yet

Dev in progress ...

# Project setup

Create a package(folder) in $GOPATH/src same as what is mentioned in MakeFile's PKG and clone the project inside newly created package

# Vendoring upfront

Please make sure to perform vendoring upfront. install glide and run cmd `glide install`

# Build guide 

The build template is forked from https://github.com/thockin/go-build-template

This has only been tested on Linux, and depends on Docker to build.

## Building

Run `make` or `make build` to compile your app.  This will use a Docker image
to build your app, with the current directory volume-mounted into place.  This
will store incremental state for the fastest possible build.  Run `make
all-build` to build for all architectures.

Run `make container` to build the container image.  It will calculate the image
tag based on the most recent git tag, and whether the repo is "dirty" since
that tag (see `make version`).  Run `make all-container` to build containers
for all architectures.

Run `make push` to push the container image to `REGISTRY`.  Run `make all-push`
to push the container images for all architectures.

Run `make clean` to clean up.

# Credits

Heavily inspired by Harbur

https://github.com/harbur/kubebot
