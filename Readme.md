[![Build Status](https://travis-ci.com/andrisasuke/jsonull.svg?branch=master)](https://travis-ci.com/andrisasuke/jsonull)
[![codecov](https://codecov.io/gh/andrisasuke/jsonull/branch/master/graph/badge.svg?token=V45A89CQMZ)](https://codecov.io/gh/andrisasuke/jsonull)

# jsonull
Handle marshalling nullable sql data type, eg : `sql.NullFloat64` `sql.NullInt64` `sql.NullString` `pq.NullTime` into proper JSON format

```
{
  "field1": {
    "String": "foo",
    "Valid": true
  },
  "field2": {
    "Int64": 0,
    "Valid": false
  }
}
```

to be :

```
{
  "field1": "foo",
  "field2": null
}
```


## Usage 
```
go get github.com/andrisasuke/jsonull
```

example struct:

```
type Person struct {
	Name      jsonull.JsonNullString  `db:"name"`
	Salary    jsonull.JsonNullFloat64 `db:"salary"`
	Code      jsonull.JsonNullInt64   `db:"code"`
	CreatedAt jsonull.JsonNullTime    `db:"created_at"`
}
