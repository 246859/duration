# duration
duration Marshaler and UnMarshaler support

## Usage

Marshaler
```go
type Person struct {
	Name    string            `json:"name"`
	Created duration.Duration `json:"created"`
}

func main() {
	p := Person{
		Name:    "John Doe",
		Created: duration.Second,
	}
	marshal, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(marshal))
}
```
output
```
{"name":"John Doe","created":"1s"}
```

UnMarshaler
```go
type Person struct {
	Name    string            `json:"name"`
	Created duration.Duration `json:"created"`
}

func main() {
	var p Person
	str := `{"name":"John Doe","created":"1s"}`
	err := json.Unmarshal([]byte(str), &p)
	if err != nil {
		panic(err)
	}
	fmt.Println(p)
}
```
output
```
{John Doe 1s}
```
