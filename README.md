# unc

Linter `unc` reports c-style nil checks.

Null checks with null on the left hand side (lhs) were popular in C to avoid the following bug:

```
if (some_var = NULL) ...
```

In Go such syntax is invalid, and hence c-style nil checks should not be used.

```
if someVar = nil ...
syntax error: cannot use assignment (someVar) = (nil) as value
```

## Install

```sh
go get github.com/mneverov/unc
```

## Run

```sh
unc ./...
```
          
`go/analysis` flags are available:

```sh
unc -h
```

To fix reported issues run:

```sh
unc -fix ./...
```

## Example
                                                                                                                     
```
func foo() {
	var err error

	if nil != err {
	}
}
```
                        
```sh
unc ./...
foo.go:7:5: lhs nil  
```


