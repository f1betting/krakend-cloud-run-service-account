<a name="readme-top"></a>

<div>
<h3 align="center">KrakenD Cloud Run Service Account</h3>

  <p align="center">
    A KrakenD plugin which injects a Google Service Account in the Authorization header
    <br />
    <a href="https://github.com/f1betting/krakend-cloud-run-service-account/issues">Report Bug</a>
    ยท
    <a href="https://github.com/f1betting/krakend-cloud-run-service-account/issues">Request Feature</a>
    <br />
    <br />
    <img alt="GitHub Latest Version" src="https://img.shields.io/github/v/release/f1betting/krakend-cloud-run-service-account?label=Latest%20release&style=flat">
    <br />
    <img alt="Go version" src="https://img.shields.io/github/go-mod/go-version/f1betting/krakend-cloud-run-service-account">
    <img alt="GitHub Workflow Status" src="https://img.shields.io/github/actions/workflow/status/f1betting/krakend-cloud-run-service-account/action-build-plugin.yml?label=Build&branch=main">
  </p>
</div>



<!-- TABLE OF CONTENTS -->

## ๐ Table of contents

- [โน๏ธ About The Project](#-about-the-project)
    - [๐ง Built With](#-built-with)
- [๐จ Getting Started](#-getting-started)
    - [โ  Prerequisites](#-prerequisites)
    - [๐๏ธ Building the plugin](#-building-the-plugin)
- [๐ Usage ](#-usage)
- [๐ License](#-license)

<!-- ABOUT THE PROJECT -->

## โน๏ธ About The Project

A plugin for [KrakenD](https://github.com/krakendio/krakend-ce) that adds a Google Service Account in the Authorization
header.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

### ๐ง Built With

[![Go]][Go-url]

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- GETTING STARTED -->

## ๐จ Getting Started

Below are the instructions on how to build the plugin.

### โ  Prerequisites

Install the dependencies using:

```shell
$ go get
```

To build this plugin, it is required to have [Docker](https://www.docker.com) installed on your system.

### ๐๏ธ Building the plugin

Run the following command to build the plugin

```shell
$ docker run -it -v "$PWD:/app" -w /app devopsfaith/krakend-plugin-builder:2.1.3 go build -buildmode=plugin -o cloud-run-service-account.so .
```

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- USAGE EXAMPLES -->

## ๐ Usage

1. Build the plugin using the instructions above
2. Add the following to your KrakenD configuration file:

```json
"plugin": {
  "pattern": ".so",
  "folder": "/etc/krakend/config/plugins/"
}
```

3. Add the following to the endpoints you wish to add this to:

_Note that this extra_config goes in the ``backend`` key of your endpoint_

```json
"extra_config": {
  "plugin/req-resp-modifier": {
    "name": [
      "cloud-run-service-account"
    ]
  }
}
```

4. Add an environment variable called ``GOOGLE_APPLICATION_CREDENTIALS`` which contains a path to the service account
   credentials JSON to your KrakenD dockerfile.

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- LICENSE -->

## ๐ License

Distributed under the MIT License. See `LICENSE.md` for more information.

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->

[Go]: https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white

[Go-url]: https://go.dev/
