package wellknown

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use AgentsItemItemWithArtifactGetResponseable instead.
type AgentsItemItemWithArtifactResponse struct {
    AgentsItemItemWithArtifactGetResponse
}
// NewAgentsItemItemWithArtifactResponse instantiates a new AgentsItemItemWithArtifactResponse and sets the default values.
func NewAgentsItemItemWithArtifactResponse()(*AgentsItemItemWithArtifactResponse) {
    m := &AgentsItemItemWithArtifactResponse{
        AgentsItemItemWithArtifactGetResponse: *NewAgentsItemItemWithArtifactGetResponse(),
    }
    return m
}
// CreateAgentsItemItemWithArtifactResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAgentsItemItemWithArtifactResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAgentsItemItemWithArtifactResponse(), nil
}
// Deprecated: This class is obsolete. Use AgentsItemItemWithArtifactGetResponseable instead.
type AgentsItemItemWithArtifactResponseable interface {
    AgentsItemItemWithArtifactGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
