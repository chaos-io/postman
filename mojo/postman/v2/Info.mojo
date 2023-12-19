
/// Detailed description of the info block
type Info {
    /// A collection's friendly name is defined by this field.
    /// You would want to set this field to a value that would allow you to easily identify
    /// this collection among a bunch of other collections, as such outlining its usage or content.
    name: String @1

    /// Every collection is identified by the unique value of this field.
    /// The value of this field is usually easiest to generate using a UID generator function.
    /// If you already have a collection, it is recommended that you maintain the same id
    /// since changing the id usually implies that is a different collection than it was originally.
    ///
    /// Note: This field exists for compatibility reasons with Collection Format V1.
    postman_id: String @2 @alias("_postman_id")

    description: Description @5
    version: Version @6

    /// This should ideally hold a link to the Postman schema that is used to validate this collection.
    /// E.g: https://schema.getpostman.com/collection/v1
    schema: String @10
}
