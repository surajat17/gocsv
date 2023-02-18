#gocsv

This is a Go program that reads data from a CSV file and outputs it in a formatted way.

## Usage

1. Clone the repository:

```
git clone https://github.com/example/gocsv.git
```

2. Build the program:

```
go build
```
3. Run the program
```
go run main.go
```


This will read the data from the `data.csv` file in the same directory and output it in a formatted way to the console.

## CSV File Format

The program assumes that the CSV file has a header row and three columns of data. The header row is skipped, and the data in each row is read into a struct with three fields: `Column1`, `Column2`, and `Column3`.

## Example

Suppose we have the following data in `data.csv`:

Column 1,Column 2,Column 3
Value 1,Value 2,Value 3
Value 4,Value 5,Value 6


When we run the program, we should see the following output:

Column 1 Column 2 Column 3
Value 1 Value 2 Value 3
Value 4 Value 5 Value 6


