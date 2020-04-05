# Fortune-API (WIP)
API server that returns fortunes. This project is mostly used for testing caching and health probes.

## Requires
- pip3 install PyYAML --user
- swagger-yaml-to-html.py - https://gist.githubusercontent.com/oseiskar/dbd51a3727fc96dcf5ed189fca491fb3/raw/2879e7849b85232bfd21e3835fb4d12e7070338a/swagger-yaml-to-html.py
- docker-compose

## Operation
```shell script
Returns a *NIX fortune via RESTful API.

Usage:
  fortune-api [flags]
  fortune-api [command]

Available Commands:
  help        Help about any command
  init        Initialises the local cache.
  server      API server commands.

Flags:
  -h, --help   help for fortune-api

Use "fortune-api [command] --help" for more information about a command.

```