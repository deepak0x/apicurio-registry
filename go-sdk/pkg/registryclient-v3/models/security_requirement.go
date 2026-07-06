package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SecurityRequirement struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The schemes property
    schemes SecurityRequirement_schemesable
}
// NewSecurityRequirement instantiates a new SecurityRequirement and sets the default values.
func NewSecurityRequirement()(*SecurityRequirement) {
    m := &SecurityRequirement{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSecurityRequirementFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSecurityRequirementFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSecurityRequirement(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *SecurityRequirement) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SecurityRequirement) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["schemes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSecurityRequirement_schemesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSchemes(val.(SecurityRequirement_schemesable))
        }
        return nil
    }
    return res
}
// GetSchemes gets the schemes property value. The schemes property
// returns a SecurityRequirement_schemesable when successful
func (m *SecurityRequirement) GetSchemes()(SecurityRequirement_schemesable) {
    return m.schemes
}
// Serialize serializes information the current object
func (m *SecurityRequirement) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("schemes", m.GetSchemes())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SecurityRequirement) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetSchemes sets the schemes property value. The schemes property
func (m *SecurityRequirement) SetSchemes(value SecurityRequirement_schemesable)() {
    m.schemes = value
}
type SecurityRequirementable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetSchemes()(SecurityRequirement_schemesable)
    SetSchemes(value SecurityRequirement_schemesable)()
}
