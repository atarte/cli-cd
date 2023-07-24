# cli-cd

## Need

- [ ] Detect update made indise a folder
- [ ] Save config inside a json file (probably at the ~)
- [ ] Cli menu to enter data and see stuff
- [ ] Multitread to look a multiple folder
- [ ] Compatible with docker
- [ ] Compatible with kubernetes
- [ ] Deployment with github action
- [ ] Unit test

## Fonctionnality

- [ ] Automatic redeployer or manual
- [ ] Checksum on folder to check if update (or find a better method if possible)
- [ ] Set time for automatic deploy
- [ ] Compatible with kube and docker
- [ ] Give the absolute path to a folder or will convert it into

## Run command

This part is for me to remenber the command and should only be use will trying to develop the app.

To set the version from the build or run command use `ldflags`.

```bash
# run
go run -ldflags "-X github.com/atarte/cli-cd/config.Version=1.0.0" ./src/main.go

# build
go build -ldflags "-X github.com/atarte/cli-cd/config.Version=1.0.0" ./src/main.go
```

You can also do the thing thing to also set the release year.

```bash
# run
go run -ldflags "-X github.com/atarte/cli-cd/config.ReleaseYear=2023" ./src/main.go

# build
go build -ldflags "-X github.com/atarte/cli-cd/config.ReleaseYear=2023" ./src/main.go
```

Or all in one commande.

```bash
go run -ldflags "-X 'github.com/atarte/cli-cd/config.Version=1.0.0' -X 'github.com/atarte/cli-cd/config.ReleaseYear=2023'" ./src/main.go

# build
go build -ldflags "-X 'github.com/atarte/cli-cd/config.Version=1.0.0' -X 'github.com/atarte/cli-cd/config.ReleaseYear=2023'" ./src/main.go
```
