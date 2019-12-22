# Flairaform

>Can Terraform be used to manage a smart home? Let's find out.

## Overview

This project is an exploration of using [Terraform](https://www.terraform.io/) to manage a smart home.

To start we will control the [Flair Smart Vent System](https://flair.co/) which is comprised of several (_~50 in my case_) IoT devices such as
- temperature sensors
- controllers (called pucks)
- motorized vents

This project extends the [Terraform SDK](https://www.terraform.io/docs/extend/plugin-sdk.html) by implementing a [provider](https://www.terraform.io/docs/extend/plugin-types.html#providers) allowing Terraform to manage the systems configuration via the [Flair API](https://api.flair.co/).

## Setup

If you want to hack on this codebase, here's a way to setup an environment that is known to work.  Feel free to  walk your own path :)

### Go

On Ubuntu, this also gets you there

`sudo apt install golang`

Use the default GOPATH/GOROOT, just easier

If you have use a non-opensource OS, I'm sorry #feelsbadman

;)

### Terraform

See https://www.terraform.io/downloads.html for general Terraform setup.  Then read the excellent docs from Hashicorp about extending Terraform.

https://www.terraform.io/docs/extend/writing-custom-providers.html

### Flair

To use the Flair API you will need a Flair developer account...and of course access to a Flair system.  I probably should not allow everyone on the Internet to control my home's HVAC system.  So this is, as they say, an exercise left for the reader...that said, I have been very happy with and highly recommend the Flair system. :)

_NOTE: I have not received any consideration from nor do I have any relationship with Flair beyond being a customer_

## Build

Official builds for this project are done with the setup above and [GitHub Actions](https://help.github.com/en/actions)

To build/test locally, after checkout do...

```
go get
go build -o terraform-provider-flair
terraform init
terraform plan
```

See the [golang docker image](https://hub.docker.com/_/golang) for more details
