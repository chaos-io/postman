| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `description` | `postman.v2.Description` |  | N |  | A Description can be a raw text, or be an object, which holds the description along with its format. |
| `name` | `string` |  | N |  | A collection's friendly name is defined by this field.You would want to set this field to a value that would allow you to easily identifythis collection among a bunch of other collections, as such outlining its usage or content. |
| `schema` | `string` |  | N |  | This should ideally hold a link to the Postman schema that is used to validate this collection.E.g:  |
| `version` | `string` | `Version` | N |  | Semantic Versioning<br>(spec)[http://semver.org/spec/v2.0.0.html] |
