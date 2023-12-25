
// Collection ref: https://schema.postman.com/collection/json/v2.1.0/draft-07/docs/index.html
type Collection {
    info: Info @1

    /// Items are the basic unit for a Postman collection.
    /// You can think of them as corresponding to a single API endpoint.
    /// Each Item has one request and may have multiple API responses associated with it.
    item: [Item] @10

    event: [Event] @20

    variable: [Variable] @30
    //auth: Auth @35
}
