| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `disable` | `boolean` |  | N |  | When set to true, prevents request body from being sent |
| `file` | `postman.v2.Request.File` |  | N |  |  |
| `formdata` | `postman.v2.Request.FormParameter` |  | N |  |  |
| `graphql` | `mojo.core.Object` |  | N |  | Object type |
| `mode` | `string` |  | N |  | Postman stores the type of data associated with this request in this fieldraw urlencoded formdata file graphql |
| `options` | `mojo.core.Object` |  | N |  | Additional configurations and options set for various body modes. |
| `raw` | `string` |  | N |  |
| `urlencoded` | `Array<postman.v2.Request.UrlEncodedParameter>` |  | N |  |
