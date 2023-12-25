
type Request {
    /// Using the Proxy, you can configure your custom proxy into the postman for particular url match
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
        /// Postman stores the type of data associated with this request in this field
        /// raw urlencoded formdata file graphql
        mode: String @1
        raw: String @5
        graphql: Object @6
        urlencoded: UrlEncodedParameter @10
        formdata: FormParameter @15
        file: File @20
        /// Additional configurations and options set for various body modes.
        options: Object @25
        /// When set to true, prevents request body from being sent
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
         /// When set to true, prevents this form data entity from being sent.
         disable: Bool @5
         type: String @8
         /// Override Content-Type header of this form data entity.
         content_type: String @9
         description: Description @10
    }

    type File {
        /// Contains the name of the file to upload. _Not the path_.
        /// A null src indicates that no file has been selected as a part of the request body
        src: String @1
        content: String @2
    }

    url: Url @1

    //auth: Auth @5
    //proxy: ProxyConfig @10
    //certificate: Certificate @11

    /// GET PUT POST PATCH DELETE COPY HEAD OPTIONS LINK UNLINK PURGE LOCK UNLOCK PROPFIND VIEW
    method: String @21

    description: Description @22

    /// A representation for a list of headers
    header: [Header] @25

    /// This field contains the data usually contained in the request body.
    body: [Body] @30
}


