# map-test

## Architecture

 - I use `golang`, `grpc`, `google-maps api` in this project. Supported fields are list below in `Supported Query Parameters` section.
 - Uses [dep](https://github.com/golang/dep/) for dependency management.
 - URL structure http://localhost:7778/<vendor>?parameters

## Config

`api-key`, `client-id`, `client-signature` is situated in `${HOME}/.map-test/config` file. You need to provide `api-key` or (`client-id` and `client-signature`) in `config` file.

Follow the [link](https://developers.google.com/places/web-service/get-api-key) to generate your own `api-key`.

## Usage

Run the project by following command
```
go run main.go
```

Now, server is up and running in `http://localhost:7778`.
If you hit, (http://localhost:7778/api/google?query=123+main+street) you can see the search results.

## Supported Query Parameters

Supported parameters are listed below.

`query`: `string which google places against this string`

`location`: `this is combination of latitude and longitude. separated with comma.`

`radius`: `radius is requied if location is provided.`

`language`: `The language code`

`minPrice`: `minimum price`

`maxPrice`: `maximum price`

`openNow`: `indicate whether open/close restaurants needs to enlist.`

`placeType`: `indicate place type`

`pageToken`: `page token`

`vendor`: `vendor name. current only supported vendor is "google". In future, "Yelp and/or Foursquare" will be supported`
```