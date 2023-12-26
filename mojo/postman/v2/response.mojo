
/// A response represents an HTTP response.
type Response {
    /// A unique, user defined identifier that can be used to refer to this response from requests
    id: String @1

    original_request: Request @10

    /// The time taken by the request to complete. If a number, the unit is milliseconds.
    /// If the response is manually created, this can be set to null.
    response_time: Timestamp @11

    /// Set of timing information related to request and response in milliseconds
    timings: Object @12

    /// No HTTP request is complete without its headers, and the same is true for
    /// a Postman request. This field is an array containing all the headers.
    header: [Header] @20

    //cookie: [Cookie] @21

    /// The raw text of the response.
    body: String @22

    /// The response status, e.g: '200 OK'.
    status: String @30

    /// The numerical response code, example: 200, 201, 404, etc/
    code: Int32 @31
}

type Cookie {
    domain: String @1
    expires: String @5
    max_age: String @10
    host_only: Bool @11
    http_only: Bool @12

    name: String @20
    path: String @21
    secure: Bool @25
    session: Bool @26
    value: String @30
}
