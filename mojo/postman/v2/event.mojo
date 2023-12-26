
/// Event Postman allows you to configure scripts to run when specific events occur.
/// These scripts are stored here, and can be referenced in the collection by their ID.
type Event {
    /// A unique identifier for the enclosing event
    id: String @1

    /// Can be set to test or prerequest for test scripts or pre-request scripts respectively.
    listen: String @5

    script: Script @6

    /// Indicates whether the event is disabled. If absent, the event is assumed to be enabled.
    disabled: Bool @10
}

/// Script A script is a snippet of Javascript code that can be used to to perform setup
/// or teardown operations on a particular response.
type Script {
    /// A unique, user defined identifier that can be used to refer to this script from requests.
    id: String @1

    /// Type of the script. E.g: 'text/javascript'
    type: String @5

    exec: [String] @6

    src: Url @10

    /// Script name
    name: String @11
}
