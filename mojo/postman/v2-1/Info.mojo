
type Info {
    name: String @1

    description: Description @5
    version: Version @6
    schema: String @10
}

type Version {
    major: Int32 @1
    minor: Int32 @2
    patch: Int32 @3
    identifier: String @10
}
