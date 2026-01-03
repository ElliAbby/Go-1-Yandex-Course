# Способы создание Map в Golang
- 1 способ:
```go
var m map[keyType]valueType
var m map[string]int
```

- 2 способ (с помощью make):
```go
m = make(map[string]int)
```

- 3 способ (с помощью литерала):
```go
m := map[keyType]valueType{
    key1: value1,
    key2: value2,
    // ...
}
```
