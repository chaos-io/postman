
type Event {
    id: String @1
    listen: String @5   // Can be set to test or prerequest for test scripts or pre-request scripts respectively.
    script: Script @6
    disabled: Bool @10
}

type Script {
    id: String @1
    type: String @5
    exec: [String] @6
    src: Url @10
    name: String @11
}
