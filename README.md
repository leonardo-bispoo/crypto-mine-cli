# Crypto Mine CLI

![banner](docs/images/banner.png)

Crypto Mine CLI is an application that extracts the data of various cryptocurrencies from the [Coin Market Cap](https://coinmarketcap.com/) and persists it in a CSV file in Donwloads Folder

## Requirements

- Golang 16+

## How to run

#### Download the project dependencies

In root folder execute:

```go
go mod tidy
```

#### Execute the project

In root folder execute:

```go
go run .
```

or

```go
go run main.go
```

The following message should appear in the terminal:

![run output](/docs/images/run-without-save.png)

#### Using the save command

Using the ```save``` you can save the results in a CSV file in your system's Downloads folder

```go
go run main.go save
```

By default the format of the saved file is CSV but you can specify which type of file you want to be saved, CSV or JSON by adding the “--type” or “-t” flag.

```go
go run main.go save --type=json
```

The following message should appear in the terminal with a log:

![run output with save](/docs/images/run-save-json.png)

or

```go
go run main.go save --type=csv
```

![run output with save](/docs/images/run-save-csv.png)