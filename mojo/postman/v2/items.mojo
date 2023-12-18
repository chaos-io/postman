
type Items {
    id: String @1
    name: String @5
    description: Description @7
    variable: [Variable] @8

    item: [Items] @10
    event: [Event] @15
    auth: Auth @20
    request: Request @25
    response: [Response] @30
}


type Description {
    content: String @1
    type: String @5 // 'text/markdown' or 'text/html'
    version: String @10
}
