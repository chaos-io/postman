| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `disabled` | `boolean` |  | N |  | Indicates whether the event is disabled. If absent, the event is assumed to be enabled. |
| `id` | `string` |  | N |  | A unique identifier for the enclosing event |
| `listen` | `string` |  | N |  | Can be set to test or prerequest for test scripts or pre-request scripts respectively. |
| `script` | `postman.v2.Script` |  | N |  | Script A script is a snippet of Javascript code that can be used to to perform setupor teardown operations on a particular response. |
