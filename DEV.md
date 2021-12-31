## To run unit tests


### Run tests with Verbose output
```
go test -v
```

### Run tests with coverage details
```
go test -v -cover
```

### Run tests and generate coverage information
Below command puts coverage data in `cover.txt` which can be read by `go tool cover` commands.
```
go test -coverprofile=cover.txt
```

### View Coverage information in HTML format
```
go tool cover -html=cover.txt
```

### View Coverage information in function wise sorted list
```
go tool cover -func=cover.txt
```


#### Coverage with Iteration Count
This command creates coverage data with iteration details.
```
go test -v -cover -covermode=count -coverprofile=cover.txt
```

For `-covermode` parameter,
- `set` : Shows a statement as covered if unit test covers it atleast once.
- `count` : Shows a statement as covered along with the intensity indicating the number of times it was covered.
- `atomic` : 
