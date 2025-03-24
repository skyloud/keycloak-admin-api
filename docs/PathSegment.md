# PathSegment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | Pointer to **string** |  | [optional] 
**MatrixParameters** | Pointer to **map[string][]string** |  | [optional] 

## Methods

### NewPathSegment

`func NewPathSegment() *PathSegment`

NewPathSegment instantiates a new PathSegment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPathSegmentWithDefaults

`func NewPathSegmentWithDefaults() *PathSegment`

NewPathSegmentWithDefaults instantiates a new PathSegment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *PathSegment) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *PathSegment) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *PathSegment) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *PathSegment) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetMatrixParameters

`func (o *PathSegment) GetMatrixParameters() map[string][]string`

GetMatrixParameters returns the MatrixParameters field if non-nil, zero value otherwise.

### GetMatrixParametersOk

`func (o *PathSegment) GetMatrixParametersOk() (*map[string][]string, bool)`

GetMatrixParametersOk returns a tuple with the MatrixParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatrixParameters

`func (o *PathSegment) SetMatrixParameters(v map[string][]string)`

SetMatrixParameters sets MatrixParameters field to given value.

### HasMatrixParameters

`func (o *PathSegment) HasMatrixParameters() bool`

HasMatrixParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


