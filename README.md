# Jocall3 CLI

The official CLI for the Jocall3 REST API.

It is generated with [Stainless](https://www.stainless.com/).

<!-- x-release-please-start-version -->

## Installation

### Installing with Go

```sh
go install 'github.com/jocall3/1231-cli/cmd/jocall3@latest'
```

<!-- x-release-please-end -->

### Running Locally

```sh
./scripts/run args...
```

## Usage

The CLI follows a resource-based command structure:

```sh
jocall3 [resource] [command] [flags]
```

```sh
jocall3 users register \
  --email alice.w@example.com \
  --name 'Alice Wonderland' \
  --password 'SecureP@ssw0rd2024!' \
  --address "{city: Anytown, country: USA, state: CA, street: 123 Main St, zip: '90210'}" \
  --date-of-birth 1990-05-10 \
  --phone +1-555-987-6543
```

For details about specific commands, use the `--help` flag.

## Global Flags

- `--help` - Show command line usage
- `--debug` - Enable debug logging (includes HTTP request/response details)
- `--version`, `-v` - Show the CLI version
- `--base-url` - Use a custom API backend URL
- `--format` - Change the output format (`auto`, `explore`, `json`, `jsonl`, `pretty`, `raw`, `yaml`)
- `--format-error` - Change the output format for errors (`auto`, `explore`, `json`, `jsonl`, `pretty`, `raw`, `yaml`)
- `--transform` - Transform the data output using [GJSON syntax](https://github.com/tidwall/gjson/blob/master/SYNTAX.md)
- `--transform-error` - Transform the error output using [GJSON syntax](https://github.com/tidwall/gjson/blob/master/SYNTAX.md)
