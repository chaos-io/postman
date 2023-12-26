
/// Header A representation for a list of headers
 type Header {
        /// This holds the LHS of the HTTP Header, e.g Content-Type or X-Custom-Header.
        key: String @1

        /// The value (or the RHS) of the Header is stored in this field.
        value: String @2

        /// If set to true, the current header will not be sent with requests
        disable: Bool @5

        description: Description @10
    }
