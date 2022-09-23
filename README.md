# kubegen

kubegen is a simple command line Kubernetes manifest generator. Instead of copy&pasting from the documentation, or trying to write the manifest from scratch are time consuming and error-prone. It's point is to provide a simple fast generator as a starting point mainly for quick POCs.

## Installation

Run `go install github.com/DavidZisky/kubegen`

## Usage

```sh
kubegen - a simple Kubernetes manifests generator

Usage:
  kubegen [flags]
  kubegen [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  deployment  Generate a deployment manifesto
  help        Help about any command
  ingress     Generate a ingress manifest
  service     Generate a service manifest
  version     Print the version number of kubegen

Flags:
  -h, --help               help for kubegen
      --log-level string   Output level of logs (TRACE, DEBUG, INFO, WARN, ERROR, FATAL) (default "INFO")

Use "kubegen [command] --help" for more information about a command.
```

Most flags can be used across commands. For instance to generate a stack of manifests for an app called "echo", one could run:

```sh
kubegen deployment --name echo --port 80
kubegen service --name echo --port 80
kubegen ingress --name echo --port 80 
```
