
/// Item are entities which contain an actual HTTP request, and sample responses attached to it.
type Item {
    /// A unique ID that is used to identify collections internally.
    id: String @1

    /// A human readable identifier for the current item.
    name: String @5
    description: Description @7

    /// Collection variables allow you to define a set of variables, that are a part of the collection,
    /// as opposed to environments, which are separate entities.
    ///
    /// Note: Collection variables must not contain any sensitive information.
    variable: [Variable] @8


    item: [Item] @10

    /// Postman allows you to configure scripts to run when specific events occur.
    /// These scripts are stored here, and can be referenced in the collection by their ID.
    event: [Event] @15

    // auth: Auth @20

    /// A request represents an HTTP request. If a string, the string is assumed to be
    /// the request URL and the method is assumed to be 'GET'.
    request: Request @25 @required

    /// A response represents an HTTP response.
    response: [Response] @30
}
