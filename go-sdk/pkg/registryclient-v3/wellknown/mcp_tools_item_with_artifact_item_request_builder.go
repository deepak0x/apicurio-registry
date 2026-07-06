package wellknown

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iefa8953a3555be741841d5395d25b8cc91d8ea997e2cc98794b61191090ff773 "github.com/apicurio/apicurio-registry/go-sdk/v3/pkg/registryclient-v3/models"
)

// McpToolsItemWithArtifactItemRequestBuilder builds and executes requests for operations under \.well-known\mcp-tools\{groupId}\{artifactId}
type McpToolsItemWithArtifactItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// McpToolsItemWithArtifactItemRequestBuilderGetQueryParameters get a registered MCP tool by group and artifact ID
type McpToolsItemWithArtifactItemRequestBuilderGetQueryParameters struct {
    Version *string `uriparametername:"version"`
}
// McpToolsItemWithArtifactItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type McpToolsItemWithArtifactItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *McpToolsItemWithArtifactItemRequestBuilderGetQueryParameters
}
// NewMcpToolsItemWithArtifactItemRequestBuilderInternal instantiates a new McpToolsItemWithArtifactItemRequestBuilder and sets the default values.
func NewMcpToolsItemWithArtifactItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*McpToolsItemWithArtifactItemRequestBuilder) {
    m := &McpToolsItemWithArtifactItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/.well-known/mcp-tools/{groupId}/{artifactId}{?version*}", pathParameters),
    }
    return m
}
// NewMcpToolsItemWithArtifactItemRequestBuilder instantiates a new McpToolsItemWithArtifactItemRequestBuilder and sets the default values.
func NewMcpToolsItemWithArtifactItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*McpToolsItemWithArtifactItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMcpToolsItemWithArtifactItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get a registered MCP tool by group and artifact ID
// Deprecated: This method is obsolete. Use GetAsWithArtifactGetResponse instead.
// returns a McpToolsItemItemWithArtifactResponseable when successful
// returns a Error error when the service returns a 404 status code
// returns a Error error when the service returns a 500 status code
func (m *McpToolsItemWithArtifactItemRequestBuilder) Get(ctx context.Context, requestConfiguration *McpToolsItemWithArtifactItemRequestBuilderGetRequestConfiguration)(McpToolsItemItemWithArtifactResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": iefa8953a3555be741841d5395d25b8cc91d8ea997e2cc98794b61191090ff773.CreateErrorFromDiscriminatorValue,
        "500": iefa8953a3555be741841d5395d25b8cc91d8ea997e2cc98794b61191090ff773.CreateErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateMcpToolsItemItemWithArtifactResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(McpToolsItemItemWithArtifactResponseable), nil
}
// GetAsWithArtifactGetResponse get a registered MCP tool by group and artifact ID
// returns a McpToolsItemItemWithArtifactGetResponseable when successful
// returns a Error error when the service returns a 404 status code
// returns a Error error when the service returns a 500 status code
func (m *McpToolsItemWithArtifactItemRequestBuilder) GetAsWithArtifactGetResponse(ctx context.Context, requestConfiguration *McpToolsItemWithArtifactItemRequestBuilderGetRequestConfiguration)(McpToolsItemItemWithArtifactGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": iefa8953a3555be741841d5395d25b8cc91d8ea997e2cc98794b61191090ff773.CreateErrorFromDiscriminatorValue,
        "500": iefa8953a3555be741841d5395d25b8cc91d8ea997e2cc98794b61191090ff773.CreateErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateMcpToolsItemItemWithArtifactGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(McpToolsItemItemWithArtifactGetResponseable), nil
}
// ToGetRequestInformation get a registered MCP tool by group and artifact ID
// returns a *RequestInformation when successful
func (m *McpToolsItemWithArtifactItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *McpToolsItemWithArtifactItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *McpToolsItemWithArtifactItemRequestBuilder when successful
func (m *McpToolsItemWithArtifactItemRequestBuilder) WithUrl(rawUrl string)(*McpToolsItemWithArtifactItemRequestBuilder) {
    return NewMcpToolsItemWithArtifactItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
