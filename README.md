# proto-go-sql

Generate sql.Scanner and sql.Valuer implementations for your Protobufs.

## Example

``` proto
# person.proto
message Person {
  option (sql.all) = "json";

  string id = 1;
}
```

Generates this person_sql.go:

``` go
func (t *Person) Scan(val interface{}) error {
	return json.Unmarshal(val.([]byte), t)
}

func (t *Person) Value() (driver.Value, error) {
	return json.Marshal(t)
}

```

## License

MIT

---

- [travisjeffery.com](http://travisjeffery.com)
- GitHub [@travisjeffery](https://github.com/travisjeffery)
- Twitter [@travisjeffery](https://twitter.com/travisjeffery)
- Medium [@travisjeffery](https://medium.com/@travisjeffery)
