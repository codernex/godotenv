# GoDotEnv - A Simple .env File Loader for Go

![Go Version](https://img.shields.io/github/go-mod/go-version/codernex/godotenv)
![License](https://img.shields.io/github/license/codernex/godotenv)
![Release](https://img.shields.io/github/v/release/codernex/godotenv)

GoDotEnv is a lightweight Go package that loads environment variables from `.env` files into your application's environment.

## Features

- Load environment variables from `.env` files
- Support for environment-specific files (`.env.production`, `.env.test`, etc.)
- Optional overwrite protection for existing environment variables
- Helper functions for type conversion (string, int, bool)
- Validation for required variables
- Zero dependencies

## Installation

```bash
go get github.com/codernex/godotenv
```
