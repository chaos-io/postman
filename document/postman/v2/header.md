| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `description` | `postman.v2.Description` |  | N |  | A Description can be a raw text, or be an object, which holds the description along with its format. |
| `disable` | `boolean` |  | N |  | If set to true, the current header will not be sent with requests |
| `key` | `string` |  | N |  | This holds the LHS of the HTTP Header, e.g Content-Type or X-Custom-Header. |
| `value` | `string` |  | N |  | The value (or the RHS) of the Header is stored in this field. |
