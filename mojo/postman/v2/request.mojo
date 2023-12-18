
type Request {
    url: [Url] @1
    auth: Auth @5
    proxy: ProxyConfig @10
    certificate: Certificate @11

    method: String @21 // GET PUT POST PATCH DELETE COPY HEAD OPTIONS LINK UNLINK PURGE LOCK UNLOCK PROPFIND VIEW
    description: Description @22
    header: [Header] @25
    body: [Body] @30
}

type Url {
    raw: String @1
    protocol: String @10 // http
    host: [String] @15
    path: [String] @18
    port: String @20
    query: QueryParam @21
    hash: String @22
    variable: [Variable] @30
}

type QueryParam {
    key: String @1
    value: String @2
    disable: Bool @5
    description: Description @10
    hash: String @11
    variable: [Variable] @12
}

type ProxyConfig {
    match: String @1  // default http+https://*/*
    host: String @5
    port: Int32 @6 // default 8080 The proxy server port
    tunnel: Bool @10
    disable: Bool @11
}

// Certificate A representation of an ssl certificate
type Certificate {
    name: String @1
    matches: [String] @5
    key: Object @10
    cert: Object @11
    passphrase: String @15
}

type Header {
    key: String @1 // Content-Type or X-Custom-Header
    value: String @2
    disable: Bool @5
    description: Description @10
}

type Body {
    mode: String @1 // raw urlencoded formdata file graphql
    raw: String @5
    graphql: Object @6

    urlencoded: UrlEncodedParameter @10
    formdata: FormParameter @15
    file: Object @20
    options: Object @25
    disable: Bool @26
}

type UrlEncodedParameter {
    key: String @1
    value: String @2
    disable: Bool @5
    description: Description @10
}

type FormParameter {
     key: String @1
     value: String @2
     disable: Bool @5
     type: String @8
     content_type: String @9
     description: Description @10
}
