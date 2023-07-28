# `pex`
[![Go Reference](https://pkg.go.dev/badge/go.essaim.dev/pex.svg)](https://pkg.go.dev/go.essaim.dev/pex)

Protobuf definitions for Event-Driven Applications in Go.

> **Warning**:
>
> `pex` is still in under development and its APIs should be considered unstable.

## Overview

`pex` is a plugin for Google's [Protocol Buffers](https://protobuf.dev/) compiler `protoc`. It generates bindings for the [Watermill](https://github.com/ThreeDotsLabs/watermill/) event handling library from Protobuf service definitions.
