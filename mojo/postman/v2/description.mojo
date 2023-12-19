
/// A Description can be a raw text, or be an object, which holds the description along with its format.
type Description {
    /// The content of the description goes here, as a raw string.
    content: String @1

    /// Holds the mime type of the raw description content.
    /// E.g: 'text/markdown' or 'text/html'.
    /// The type is used to correctly render the description when generating documentation, or in the Postman app.
    type: String @5

    /// Description can have versions associated with it, which should be put in this property.
    version: String @10
}
