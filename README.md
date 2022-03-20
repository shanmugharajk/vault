# vault

A simple command line utility written in Go to store passwords / secret strings encrpted with password.

## How to build

Run the following command in the project root

```sh
go build -o ./bin/vault ./cmd/vault/main.go
```

The format is `go build -o |path to store the executable| |path to the main go gile|`.

After running the command we will get the binary in `bin` folder (present in the project root). Navigate to it and copy the binary to the desired location.

## How to use

After the following the build steps, execute the following steps

- Navigate to the folder where the `vault` executable is present.
- Run `vault setup`.
- The list of commands can be found by the command `vault help`.
- To save a secret run `vault save` and follow the prompt.
- To fetch the stored secret run `vault fetch` and follow the prompt.
