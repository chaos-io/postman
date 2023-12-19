| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `description` | `postman.v2.Description` |  | N |  | A Description can be a raw text, or be an object, which holds the description along with its format. |
| `disable` | `boolean` |  | N |  |
| `id` | `string` |  | N |  | A variable ID is a unique user-defined value that identifies the variablewithin a collection. In traditional terms, this would be a variable name. |
| `key` | `string` |  | N |  | A variable key is a human friendly value that identifies the variable withina collection. In traditional terms, this would be a variable name. |
| `name` | `string` |  | N |  | Variable name |
| `system` | `boolean` |  | N |  | When set to true, indicates that this variable has been set by Postman |
| `type` | `string` |  | N |  | A variable may have multiple types. This field specifies the type of the variable |
| `value` | `string` |  | N |  | The value that a variable holds in this collection. Ultimately, the variableswill be replaced by this value, when say running a set of requests from a collection |
