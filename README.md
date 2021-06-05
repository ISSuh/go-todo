# go-todo

implement simple todo api server using golang

## API

| method | url | description |
|---|---|---|
| GET | /item | get all todo item |
| POST | /item | add todo item |
| GET | /item/:id | get item on matched id |
| DELETE | /item/:id | delete item on matched id |

### Example

#### GET /item
---

#### Response

```json
{
  "List": [
    {
      "Id": 0,
      "Item": {
        "Time": "20210603",
        "Title": "Test",
        "Works": [
          {
            "Content": "test1",
            "Done": false
          },
          {
            "Content": "test2",
            "Done": true
          }
        ]
      }
    }
  ]
}
```

#### POST /item
---

#### Request

```json
{
    "item": {
        "time": "20210603",
        "title": "Test",
        "works": [
            {
                "content": "test1",
                "done": false
            },
            {
                "content": "test2",
                "done": true
            }
        ]
    }
}
```

#### Response

```json
{
  "Status": true,
  "Id": 1,
  "Err": ""
}
```

#### GET /item/:id
---

#### Response

```json
{
  "Id": 1,
  "Item": {
    "Time": "20210603",
    "Title": "Test",
    "Works": [
      {
        "Content": "test1",
        "Done": false
      },
      {
        "Content": "test2",
        "Done": true
      }
    ]
  }
}
```

#### DELETE /item/:id
---

#### Response

```json
{
  "Status": true,
  "Id": 0,
  "Err": ""
}
```

## Todo

- implement authorization
- store item in database