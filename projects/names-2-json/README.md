# Names-2-Json

A sample web services to parse names from an get http request to json.

## Demo

Build the application with the following command:

```sh
go build .
```

Run the web server:

```sh
./names-2-json
```

Call the web server in a different terminal:

```sh
curl "http://localhost:3000/?name=Tom&name=Bob"
```

Your response should be:

```json
{"names": ["Tom", "Bob"]}
```