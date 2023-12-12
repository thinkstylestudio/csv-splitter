# CSV Splitter

This is a command line tool written in Go that splits a CSV file into multiple smaller CSV files.

## Description

The tool takes a CSV file as an argument and splits it into multiple files, each containing a maximum of 150 rows. The first row of the source file, which is assumed to be the header, is not included in the output files. The output files are named sequentially as `import.csv`, `import1.csv`, `import2.csv`, etc.

## Usage

To use the tool, run the following command:

```bash
go run main.go yourfile.csv
```

Replace `yourfile.csv` with the path to your CSV file. The output files will be created in the same directory as the script.

## Requirements

* Go programming language
