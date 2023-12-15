
type Variable {
    id: String @1
    key: String @2
    value: String @3
    type: String @10
    name: String @11
    description: Description @12
    system: Bool @13    // When set to true, indicates that this variable has been set by Postman
    disable: Bool @14
}
