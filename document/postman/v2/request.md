| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `body` | `postman.v2.Request.Body` |  | N |  | This field contains the data usually contained in the request body. |
| `description` | `postman.v2.Description` |  | N |  | A Description can be a raw text, or be an object, which holds the description along with its format. |
| `header` | `Array<postman.v2.Header>` |  | N |  | A representation for a list of headers |
| `method` | `string` |  | N |  | GET PUT POST PATCH DELETE COPY HEAD OPTIONS LINK UNLINK PURGE LOCK UNLOCK PROPFIND VIEW |
| `url` | `postman.v2.Url` |  | N |  | If object, contains the complete broken-down URL for this request. If string, contains the literal request URL. |
