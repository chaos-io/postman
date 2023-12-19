| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `description` | `postman.v2.Description` |  | N |  | A Description can be a raw text, or be an object, which holds the description along with its format. |
| `event` | `Array<postman.v2.Event>` |  | N |  | Postman allows you to configure scripts to run when specific events occur.These scripts are stored here, and can be referenced in the collection by their ID. |
| `id` | `string` |  | N |  | A unique ID that is used to identify collections internally. |
| `item` | `Array<postman.v2.Item>` |  | N |  |
| `name` | `string` |  | N |  | A human readable identifier for the current item. |
| `request` | `postman.v2.Request` |  | Y |  | A request represents an HTTP request. If a string, the string is assumed to bethe request URL and the method is assumed to be 'GET'. |
| `response` | `Array<postman.v2.Response>` |  | N |  | A response represents an HTTP response. |
| `variable` | `Array<postman.v2.Variable>` |  | N |  | Collection variables allow you to define a set of variables, that are a part of the collection,as opposed to environments, which are separate entities.<br>Note: Collection variables must not contain any sensitive information. |
