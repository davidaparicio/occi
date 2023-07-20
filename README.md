# OcCI - OVHcloud Cloud Inventory

<p align="left">
<img src="assets/img/occi.jpg" alt="OVHcloud Cloud Inventory logo" title="occi logo" />
</p>

[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://pkg.go.dev/github.com/davidaparicio/occi)
[![Go Report Card](https://goreportcard.com/badge/davidaparicio/occi)](https://goreportcard.com/report/davidaparicio/occi)
[![codecov](https://codecov.io/gh/davidaparicio/occi/branch/main/graph/badge.svg?token=VYP4LAODQ6)](https://codecov.io/gh/davidaparicio/occi)
[![build](https://github.com/davidaparicio/occi/actions/workflows/goreleaser.yml/badge.svg)](https://github.com/davidaparicio/occi/actions/workflows/goreleaser.yml)

## Table of Contents

- [Introduction](#introduction)
- [Features](#features)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
- [Usage](#usage)
- [Contribute](#contribute)
- [License](#license)
- [Contact](#contact)

## Introduction

The OVHcloud Cloud Inventory project aims to provide a tool for efficiently managing and tracking assets within the OVHcloud cloud provider. It enables users to maintain an up-to-date inventory of all their resources, making it easier to manage, monitor, and optimize their cloud infrastructure.

### Homonyms

> [Occi](https://goo.gl/maps/bQvXDHomeCv98VkF7) is also the abandoned village of Lumio in Haute-Corse. The hamlet of Occi is located above Lumio, in the Balagne region, about 10 km from Calvi, at an altitude of 377 m.

Or occi can mean ["Open Cloud Computing Interface (OCCI)"](https://en.wikipedia.org/wiki/Open_Cloud_Computing_Interface)

## Features

Please note, some features will be delivered in the
- **Resource Discovery**: Automatically discover and collect information about assets in your OVHcloud account.
- **Asset Tracking**: Keep track of virtual machines, storage volumes, networking resources, and other essential components.
- **Tagging Support**: Organize assets using tags to facilitate better categorization and grouping.
- **Reporting**: Generate comprehensive reports to gain insights into resource usage and spending.
- **Automation**: Implement automation to schedule inventory updates and data retrieval.

## Getting Started

To get started with the OVHcloud Cloud Inventory project, follow the instructions below.

### Prerequisites

- [OVHcloud API Key](https://www.ovh.com/auth/api/createToken): You'll need to generate an API key with the necessary permissions to access cloud resources
- Go (only if you contribute/for the development mode)

### Installation

### Standard setup
1. Get the latest [OcCI release](https://github.com/davidaparicio/occi/releases)

2. Edit the [app.env](#usage) file with your credentials

3. Give the execution rights
```bash
chmod +x occi
```

4.  Double click to run it (or if you want, move it into your PATH)
```bash
./occi
```

### Developer mode
1. Clone the repository:

```bash
git clone https://github.com/davidaparicio/occi.git
cd occi
```

2. Install the required dependencies for:
```bash
go mod download
```

### Usage

* Configure your OVHcloud API credentials:
    * Create a [new API key for this application](https://www.ovh.com/auth/api/createToken) in the OVHcloud console.
    * Set up the app.env file as described this the example
```bash
# https://www.ovh.com/auth/api/createToken
# configuration specific to 'ovh-eu' endpoint
ENDPOINT=ovh-eu
APPLICATION_KEY=YOUR_APPLICATION_KEY
APPLICATION_SECRET=YOUR_APPLICATION_SECRET
CONSUMER_KEY=YOUR_CONSUMER_KEY
```
    * (next version) Set up the required environment variables for authentication.
```bash
export OVH_ENDPOINT=your_api_endpoint
export OVH_APPLICATION_KEY=your_application_key
export OVH_APPLICATION_SECRET=your_application_secret
export OVH_CONSUMER_KEY=your_consumer_key
```

* Run the inventory script:
```bash
go run main.go | tee yournic-ovh.inventory
```

* Access your inventory data in the specified output format, with your favorite tool/editor/emacs/vim/whatever.
```bash
open yournic-ovh.inventory
```

## Improvement list
* Add PCC exploration (VMWare) ```util.PrintPCC```
* Export all the OVHcloud APIv6 type in the Go project ```/api/types```
* CLI mode
    * Provide the application configuration or enviromnent variable (less secure)
    * Provide the assets wanted (only PCI, only web, etc...)
* [Age/SOPS](https://devops.datenkollektiv.de/using-sops-with-age-and-git-like-a-pro.html) activation

## Contribute

We welcome contributions from the community! If you'd like to contribute to the project, please follow the guidelines outlined in CONTRIBUTING.md.

Works on my machine - and yours ! Spin up pre-configured, standardized dev environments of this repository, by clicking on the button below.

[![Open in Gitpod](https://gitpod.io/button/open-in-gitpod.svg)](https://gitpod.io/#/https://github.com/davidaparicio/occi)

## See also

* [select * from cloud; (SteamPipe.io)](https://steampipe.io)

Some "old" [Github projects](https://github.com/topics/occi) with the same projectname:
* [2015 - Organization OCCI / Open Cloud Computing Interface](https://github.com/occi)
* [2022 - OCCI Studio](https://github.com/occiware/OCCI-Studio)
* [2011 - occi-py](https://github.com/nyren/occi-py)
* [2016 - Adafruit's Occi](https://github.com/adafruit/Adafruit-Occi)
* [2015 - OCCI-OS/OpenStack](https://github.com/tmetsch/occi-os)[(SecondLink](https://github.com/stackforge/occi-os)

## License
Licensed under the MIT License, Version 2.0 (the "License"). You may not use this file except in compliance with the License.
You may obtain a copy of the License [here](https://choosealicense.com/licenses/mit/).

If needed some help,  there are a ["Licenses 101" by FOSSA](https://fossa.com/blog/open-source-licenses-101-mit-license/), a [Snyk explanation](https://snyk.io/learn/what-is-mit-license/)
of MIT license and a [French conference talk](https://www.youtube.com/watch?v=8WwTe0vLhgc) by [Jean-Michael Legait](https://twitter.com/jmlegait) about licenses.

[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fdavidaparicio%2Focci.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2Fdavidaparicio%2Focci?ref=badge_large)

## Contact

If you have any questions or suggestions regarding the project, feel free to reach out to our team in the [GitHub issues](https://github.com/davidaparicio/occi/issues).