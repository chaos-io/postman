
type Response {
    id: String @1

    original_request: Request @10
    response_time: Timestamp @11
    timings: Object @12

    header: [String] @20
    cookie: [Cookie] @21
    body: String @22

    status: String @30
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
