# kubegen

kubegen is a simple command line Kubernetes manifest generator. Instead of copy&pasting from the documentation, or trying to write the manifest from scratch are time consuming and error-prone. It's point is to provide a simple fast generator as a starting point mainly for quick POCs.

## Installation

Run `go install github.com/DavidZisky/kubegen`

## Usage

```sh
kubegen is a very fast generator for Kubernetes manifests

Usage:
  kubegen [flags] command

Available Commands:
  help        Help about any command
  version     Print the version number of kubegen

  deployment  Generate a deployment manifest
  service     Generate a service manifest
  ingress     Generate a ingress manifest

Flags:
      --config string      /path/to/config.yml
  -h, --help               help for kubegen
      --log-level string   Output level of logs (TRACE, DEBUG, INFO, WARN, ERROR, FATAL) (default "INFO")
  -v, --version            Display the current version of this CLI

Use "kubegen [command] --help" for more information about a command.
```

Most flags can be used across commands. For instance to generate a stack of manifests for an app called "echo", one could run:

```sh
kubegen --name echo --port 80 deployment
kubegen --name echo --port 80 service
kubegen --name echo --port 80 ingress
```
