
// ref: https://schema.postman.com/collection/json/v2.1.0/draft-07/docs/index.html
type Postman {
    info: Info @1
    item: [Items] @10
    event: [Event] @20
    variable: [Variable] @30
    auth: Auth @35
}
