package wellknown

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iefa8953a3555be741841d5395d25b8cc91d8ea997e2cc98794b61191090ff773 "github.com/apicurio/apicurio-registry/go-sdk/v3/pkg/registryclient-v3/models"
)

// SchemasItemWithVersionItemRequestBuilder builds and executes requests for operations under \.well-known\schemas\{schemaType}\{version}
type SchemasItemWithVersionItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// SchemasItemWithVersionItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SchemasItemWithVersionItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewSchemasItemWithVersionItemRequestBuilderInternal instantiates a new SchemasItemWithVersionItemRequestBuilder and sets the default values.
func NewSchemasItemWithVersionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SchemasItemWithVersionItemRequestBuilder) {
    m := &SchemasItemWithVersionItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/.well-known/schemas/{schemaType}/{version}", pathParameters),
    }
    return m
}
// NewSchemasItemWithVersionItemRequestBuilder instantiates a new SchemasItemWithVersionItemRequestBuilder and sets the default values.
func NewSchemasItemWithVersionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SchemasItemWithVersionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSchemasItemWithVersionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get a well-known schema by type and version
// Deprecated: This method is obsolete. Use GetAsWithVersionGetResponse instead.
// returns a SchemasItemItemWithVersionResponseable when successful
// returns a Error error when the service returns a 404 status code
// returns a Error error when the service returns a 500 status code
func (m *SchemasItemWithVersionItemRequestBuilder) Get(ctx context.Context, requestConfiguration *SchemasItemWithVersionItemRequestBuilderGetRequestConfiguration)(SchemasItemItemWithVersionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": iefa8953a3555be741841d5395d25b8cc91d8ea997e2cc98794b61191090ff773.CreateErrorFromDiscriminatorValue,
        "500": iefa8953a3555be741841d5395d25b8cc91d8ea997e2cc98794b61191090ff773.CreateErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateSchemasItemItemWithVersionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(SchemasItemItemWithVersionResponseable), nil
}
// GetAsWithVersionGetResponse get a well-known schema by type and version
// returns a SchemasItemItemWithVersionGetResponseable when successful
// returns a Error error when the service returns a 404 status code
// returns a Error error when the service returns a 500 status code
func (m *SchemasItemWithVersionItemRequestBuilder) GetAsWithVersionGetResponse(ctx context.Context, requestConfiguration *SchemasItemWithVersionItemRequestBuilderGetRequestConfiguration)(SchemasItemItemWithVersionGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": iefa8953a3555be741841d5395d25b8cc91d8ea997e2cc98794b61191090ff773.CreateErrorFromDiscriminatorValue,
        "500": iefa8953a3555be741841d5395d25b8cc91d8ea997e2cc98794b61191090ff773.CreateErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateSchemasItemItemWithVersionGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(SchemasItemItemWithVersionGetResponseable), nil
}
// ToGetRequestInformation get a well-known schema by type and version
// returns a *RequestInformation when successful
func (m *SchemasItemWithVersionItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *SchemasItemWithVersionItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *SchemasItemWithVersionItemRequestBuilder when successful
func (m *SchemasItemWithVersionItemRequestBuilder) WithUrl(rawUrl string)(*SchemasItemWithVersionItemRequestBuilder) {
    return NewSchemasItemWithVersionItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
