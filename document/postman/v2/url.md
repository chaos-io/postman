| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `hash` | `string` |  | N |  | Contains the URL fragment (if any). Usually this is not transmitted over the network,but it could be useful to store this in some cases. |
| `host` | `Array<string>` |  | N |  | The host for the URL, E.g: api.yourdomain.com.Can be stored as a string or as an array of strings. |
| `path` | `Array<string>` |  | N |  | The complete path of the current url, broken down into segments.A segment could be a string, or a path variable. |
| `port` | `string` |  | N |  | The port number present in this URL. An empty value implies 80/443depending on whether the protocol field contains http/https. |
| `protocol` | `string` |  | N |  | The protocol associated with the request, E.g: 'http' |
| `query` | `Array<postman.v2.Url.QueryParam>` |  | N |  | An array of QueryParams, which is basically the query string part of the URL,parsed into separate variables |
| `raw` | `string` |  | N |  | The string representation of the request URL, including the protocol,host, path, hash, query parameter(s) and path variable(s). |
| `variable` | `Array<postman.v2.Variable>` |  | N |  | Postman supports path variables with the syntax /path/:variableName/to/somewhere.These variables are stored in this field. |
