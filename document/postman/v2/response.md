| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `body` | `string` |  | N |  | The raw text of the response. |
| `code` | `integer` | `Int32` | N |  | The numerical response code, example: 200, 201, 404, etc/ |
| `header` | `Array<string>` |  | N |  | No HTTP request is complete without its headers, and the same is true fora Postman request. This field is an array containing all the headers. |
| `id` | `string` |  | N |  | A unique, user defined identifier that can be used to refer to this response from requests |
| `originalRequest` | `postman.v2.Request` |  | N |  |  |
| `responseTime` | `string` | `Timestamp` | N |  | The time taken by the request to complete. If a number, the unit is milliseconds.If the response is manually created, this can be set to null. |
| `status` | `string` |  | N |  | The response status, e.g: '200 OK'. |
| `timings` | `mojo.core.Object` |  | N |  | Set of timing information related to request and response in milliseconds |
