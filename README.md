Fakadog
=======

[![Build Status](https://travis-ci.org/manywho/fakadog.svg?branch=master)](https://travis-ci.org/manywho/fakadog)

This little tool is a server that simply mocks a Datadog APM response, intended for use in an environment where a Datadog
APM client is being used, but no agent exists, e.g. in integration tests.

## Usage

The tool is available as a [Docker image](https://quay.io/repository/manywho/fakadog), and exposes a server on port `8126`:

```bash
$ docker pull quay.io/manywho/fakadog
$ docker run quay.io/manywho/fakadog

{"level":"info","msg":"Starting server on port 0.0.0.0:8126","time":"2019-09-11T12:02:04+01:00"}
```

## Contributing

Contributions are welcome to the project - whether they are feature requests, improvements or bug fixes! Refer to 
[CONTRIBUTING.md](CONTRIBUTING.md) for our contribution requirements.

## License

This tool is released under the [MIT License](https://opensource.org/licenses/MIT).