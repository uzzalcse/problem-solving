## Property Distribution based on priority

### setup and installation 
    Prerequisites: Golang version 1.23.2

### Clone the repository 

``` 
git clone https://github.com/uzzalcse/problem-solving.git
cd problem-solving

```

### Run and tests 

#### Run 

```
go run property_divider.go

```

#### Run tests

```
go test -v

```
Generate coverage report:
```bash
go test ./... -cover
go test ./... -coverprofile=coverage.out
```

View coverage in terminal:
```bash
go tool cover -func=coverage.out
```