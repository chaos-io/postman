
/// If object, contains the complete broken-down URL for this request. If string, contains the literal request URL.
type Url {
    type QueryParam {
        key: String @1

        value: String @2

        /// If set to true, the current query parameter will not be sent with the request
        disable: Bool @5

        description: Description @10
    }

    /// The string representation of the request URL, including the protocol,
    /// host, path, hash, query parameter(s) and path variable(s).
    raw: String @1

    /// The protocol associated with the request, E.g: 'http'
    protocol: String @10

    /// The host for the URL, E.g: api.yourdomain.com.
    /// Can be stored as a string or as an array of strings.
    host: [String] @15

    /// The complete path of the current url, broken down into segments.
    /// A segment could be a string, or a path variable.
    path: [String] @18

    /// The port number present in this URL. An empty value implies 80/443
    /// depending on whether the protocol field contains http/https.
    port: String @20

    /// An array of QueryParams, which is basically the query string part of the URL,
    /// parsed into separate variables
    query: QueryParam @21

    /// Contains the URL fragment (if any). Usually this is not transmitted over the network,
    /// but it could be useful to store this in some cases.
    hash: String @22

    /// Postman supports path variables with the syntax /path/:variableName/to/somewhere.
    /// These variables are stored in this field.
    variable: [Variable] @30
}


