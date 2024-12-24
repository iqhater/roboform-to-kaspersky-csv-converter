# Roboform to Kaspersky Password Manager CSV converter

## Description

Kaspersky Password Manager officially does not support export csv from Roboform. So, i wrote a pretty simple csv conveter to easy transfer your accounts and passwords from Roboform into Kaspersky Password Manager.

## How to use

### Prerequisites

You need to install [Golang](https://go.dev/dl/) language

### 1. Clone repository

```sh
git clone https://github.com/iqhater/roboform-to-kaspersky-csv-converter.git
```

### 2. Setup your csv filenames on top of the file `main.go`

```go
// main.go file
const (
    INPUT_FILENAME  = "your_input_filename.csv"
    OUTPUT_FILENAME = "your_output_filename.csv" // can be any name
)
```

### 3. Compile and run your binary file

```go
go run main.go
```

### 4. Get your result `output.csv` file

## Additional info

> ⚠️ **Kaspersky Password Manager does not support `notes` export!**

### Copied from kaspersky docs reference csv example

[Kaspersky csv example file](reference_example.csv)

### Tested on apps versions

> Not guaranteed worked on older or newer versions

```sh
- Roboform: Version 9.6.3.3 (en-english, RUS) Oct 16 2024 09:12:00
- Kaspersky Password Manager: 24.3.0.327 / 1731938737_7632 / 1.1.0.2_1
```

## Contributing

Feel free to any contributions ;)

## License

Distributed under the MIT License. See [LICENSE](LICENSE) for more information.
