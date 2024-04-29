# gopherapi

## TODO

- [X] ~~Update timestamps on CreatedAt and UpdatedAt~~
- [X] ~~Fix incremental indexing~~
- [ ] Refactor error handling
- [X] ~~Debug sql error on post req~~
- [ ] Implement pagination for query results
- [ ] Add unit tests for all API endpoints
- [ ] Improve error messages for invalid requests
- [ ] Optimize database queries for better performance
- [ ] Write API Documentation

## Responses

`Post /account`

 ```json
{
    "first_name": "first_name",
    "last_name": "last_name",
    "email": "test@test.com"
}
```

`Get /account/{ID}`
`Delete /account/{ID}`
