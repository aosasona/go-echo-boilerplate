# Gopi

Golang API boilerplate with Echo, Bun and BoltDB

## Getting started

- Create a new repo with this template or clone the repo
- Run the following command to properly set your project name

```bash
chmod +x ./init
./init project-name
```

## Running

To start the API, run the following command:

```bash
go run ./cmd/api
```

## Dockerfiles

- `Dockerfile` - has the bare essentials to run your API (NOTE: this uses a distroless image, you will be missing a lot of coreutils you may need for debugging)
- `Dockerfile.dev` - simply runs your project in Docker but with hot reloading, useful for a dev docker-compose where you probably require internal networking
- `Dockerfile.fullstack` - has all you need to build an image with a frontend in the `ui` folder included, you will end up with an image with some coreutils and your built project fused with your UI
