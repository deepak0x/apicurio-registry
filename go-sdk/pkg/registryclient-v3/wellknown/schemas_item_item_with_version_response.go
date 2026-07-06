package wellknown

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use SchemasItemItemWithVersionGetResponseable instead.
type SchemasItemItemWithVersionResponse struct {
    SchemasItemItemWithVersionGetResponse
}
// NewSchemasItemItemWithVersionResponse instantiates a new SchemasItemItemWithVersionResponse and sets the default values.
func NewSchemasItemItemWithVersionResponse()(*SchemasItemItemWithVersionResponse) {
    m := &SchemasItemItemWithVersionResponse{
        SchemasItemItemWithVersionGetResponse: *NewSchemasItemItemWithVersionGetResponse(),
    }
    return m
}
// CreateSchemasItemItemWithVersionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSchemasItemItemWithVersionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSchemasItemItemWithVersionResponse(), nil
}
// Deprecated: This class is obsolete. Use SchemasItemItemWithVersionGetResponseable instead.
type SchemasItemItemWithVersionResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    SchemasItemItemWithVersionGetResponseable
}
