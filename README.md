# proto-go-sql

Generate sql.Scanner and driver.Valuer implementations for your Protobufs.

## Example

We want the generated struct for this Person message to implement sql.Scanner and driver.Valuer so we can easily write and read it as JSON from Postgres.

So we compile the person.proto file:

``` proto
syntax = "proto3";

import "github.com/travisjeffery/proto-go-sql/sql.proto";

message Person {
  option (sql.all) = "json";

  string id = 1;
}
```

And run:

``` sh
$ protoc --sql_out=. person.proto
```

Generating this person_sql.go:

``` go
func (t *Person) Scan(val interface{}) error {
	valRaw := val.([]byte)
	valCopy := make([]byte, len(valRaw))
	copy(valCopy, valRaw)
	return json.Unmarshal(valCopy, t)
}

func (t *Person) Value() (driver.Value, error) {
	return json.Marshal(t)
}

```

And we're done!

(The db driver owns the bytes passed into scan, that's why we make a copy so the bytes don't change
as we're unmarshaling. [This
issue](https://go-review.googlesource.com/c/go/+/108535/) explains it.)

## License

MIT

---

- [travisjeffery.com](http://travisjeffery.com)
- GitHub [@travisjeffery](https://github.com/travisjeffery)
- Twitter [@travisjeffery](https://twitter.com/travisjeffery)
- Medium [@travisjeffery](https://medium.com/@travisjeffery)
