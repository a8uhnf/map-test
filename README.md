# map-test

### Build api

```
make api
```

### Config

`api-key`, `client-id`, `client-signature` is situated in `${HOME}/.map-test/config` file. You need to provide `api-key` or (`client-id` and `client-signature`) in `config` file.

Follow the [link](https://developers.google.com/places/web-service/get-api-key) to generate your own `api-key`.

### Api Assumption

I assume api-key, location(latitude, longitude), radius fields must be provide for search

### Usage

Run the project by following command
```
go run main.go
```

Now, server is up and running in http://localhost:7778.
If you hit, http://localhost:7778/api/google?query=hello&localtion.lat=23&location.lng=90&radius=10, you can see the search results.