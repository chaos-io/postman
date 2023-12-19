
/// Auth Represents authentication helpers provided by Postman
type Auth {
    /// enum string apikey awsv4 basic bearer digest edgegrid hawk noauth oauth1 oauth2 ntlm
    type: String @1

    /// The attributes for API Key Authentication
    apikey: [Auth] @10
    awsv4: [Auth] @11
    basic: [Auth] @12
    edgegrid: [Auth] @13
    hawk: [Auth] @14
    ntlm: [Auth] @15
    oauth1: [Auth] @16
    oauth2: [Auth] @17
}
