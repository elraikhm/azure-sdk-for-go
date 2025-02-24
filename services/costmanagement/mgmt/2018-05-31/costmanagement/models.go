package costmanagement

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"encoding/json"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// The package's fully qualified name.
const fqdn = "github.com/Azure/azure-sdk-for-go/services/costmanagement/mgmt/2018-05-31/costmanagement"

// Dimension ...
type Dimension struct {
	*DimensionProperties `json:"properties,omitempty"`
	// ID - READ-ONLY; Resource Id.
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; Resource name.
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; Resource type.
	Type *string `json:"type,omitempty"`
	// Tags - READ-ONLY; Resource tags.
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for Dimension.
func (d Dimension) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if d.DimensionProperties != nil {
		objectMap["properties"] = d.DimensionProperties
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for Dimension struct.
func (d *Dimension) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var dimensionProperties DimensionProperties
				err = json.Unmarshal(*v, &dimensionProperties)
				if err != nil {
					return err
				}
				d.DimensionProperties = &dimensionProperties
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				d.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				d.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				d.Type = &typeVar
			}
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				d.Tags = tags
			}
		}
	}

	return nil
}

// DimensionProperties ...
type DimensionProperties struct {
	Data       *[]string  `json:"data,omitempty"`
	Total      *int32     `json:"total,omitempty"`
	Category   *string    `json:"category,omitempty"`
	UsageStart *date.Time `json:"usageStart,omitempty"`
	UsageEnd   *date.Time `json:"usageEnd,omitempty"`
	NextLink   *string    `json:"nextLink,omitempty"`
}

// DimensionsListResult result of listing dimensions. It contains a list of available dimensions.
type DimensionsListResult struct {
	autorest.Response `json:"-"`
	// Value - READ-ONLY; The list of dimensions.
	Value *[]Dimension `json:"value,omitempty"`
}

// ErrorDetails the details of the error.
type ErrorDetails struct {
	// Code - READ-ONLY; Error code.
	Code *string `json:"code,omitempty"`
	// Message - READ-ONLY; Error message indicating why the operation failed.
	Message *string `json:"message,omitempty"`
}

// ErrorResponse error response indicates that the service is not able to process the incoming request. The
// reason is provided in the error message.
type ErrorResponse struct {
	// Error - The details of the error.
	Error *ErrorDetails `json:"error,omitempty"`
}

// Operation a Cost management REST API operation.
type Operation struct {
	// Name - READ-ONLY; Operation name: {provider}/{resource}/{operation}.
	Name *string `json:"name,omitempty"`
	// Display - The object that represents the operation.
	Display *OperationDisplay `json:"display,omitempty"`
}

// MarshalJSON is the custom marshaler for Operation.
func (o Operation) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if o.Display != nil {
		objectMap["display"] = o.Display
	}
	return json.Marshal(objectMap)
}

// OperationDisplay the object that represents the operation.
type OperationDisplay struct {
	// Provider - READ-ONLY; Service provider: Microsoft.CostManagement.
	Provider *string `json:"provider,omitempty"`
	// Resource - READ-ONLY; Resource on which the operation is performed: Dimensions, Query.
	Resource *string `json:"resource,omitempty"`
	// Operation - READ-ONLY; Operation type: Read, write, delete, etc.
	Operation *string `json:"operation,omitempty"`
}

// OperationListResult result of listing cost management operations. It contains a list of operations and a
// URL link to get the next set of results.
type OperationListResult struct {
	autorest.Response `json:"-"`
	// Value - READ-ONLY; List of cost management operations supported by the Microsoft.CostManagement resource provider.
	Value *[]Operation `json:"value,omitempty"`
	// NextLink - READ-ONLY; URL to get the next set of operation list results if there are any.
	NextLink *string `json:"nextLink,omitempty"`
}

// OperationListResultIterator provides access to a complete listing of Operation values.
type OperationListResultIterator struct {
	i    int
	page OperationListResultPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *OperationListResultIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OperationListResultIterator.NextWithContext")
		defer func() {
			sc := -1
			if iter.Response().Response.Response != nil {
				sc = iter.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err = iter.page.NextWithContext(ctx)
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (iter *OperationListResultIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter OperationListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter OperationListResultIterator) Response() OperationListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter OperationListResultIterator) Value() Operation {
	if !iter.page.NotDone() {
		return Operation{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the OperationListResultIterator type.
func NewOperationListResultIterator(page OperationListResultPage) OperationListResultIterator {
	return OperationListResultIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (olr OperationListResult) IsEmpty() bool {
	return olr.Value == nil || len(*olr.Value) == 0
}

// hasNextLink returns true if the NextLink is not empty.
func (olr OperationListResult) hasNextLink() bool {
	return olr.NextLink != nil && len(*olr.NextLink) != 0
}

// operationListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (olr OperationListResult) operationListResultPreparer(ctx context.Context) (*http.Request, error) {
	if !olr.hasNextLink() {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(olr.NextLink)))
}

// OperationListResultPage contains a page of Operation values.
type OperationListResultPage struct {
	fn  func(context.Context, OperationListResult) (OperationListResult, error)
	olr OperationListResult
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *OperationListResultPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OperationListResultPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	for {
		next, err := page.fn(ctx, page.olr)
		if err != nil {
			return err
		}
		page.olr = next
		if !next.hasNextLink() || !next.IsEmpty() {
			break
		}
	}
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *OperationListResultPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page OperationListResultPage) NotDone() bool {
	return !page.olr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page OperationListResultPage) Response() OperationListResult {
	return page.olr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page OperationListResultPage) Values() []Operation {
	if page.olr.IsEmpty() {
		return nil
	}
	return *page.olr.Value
}

// Creates a new instance of the OperationListResultPage type.
func NewOperationListResultPage(cur OperationListResult, getNextPage func(context.Context, OperationListResult) (OperationListResult, error)) OperationListResultPage {
	return OperationListResultPage{
		fn:  getNextPage,
		olr: cur,
	}
}

// Query ...
type Query struct {
	*QueryProperties `json:"properties,omitempty"`
	// ID - READ-ONLY; Resource Id.
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; Resource name.
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; Resource type.
	Type *string `json:"type,omitempty"`
	// Tags - READ-ONLY; Resource tags.
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for Query.
func (q Query) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if q.QueryProperties != nil {
		objectMap["properties"] = q.QueryProperties
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for Query struct.
func (q *Query) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var queryProperties QueryProperties
				err = json.Unmarshal(*v, &queryProperties)
				if err != nil {
					return err
				}
				q.QueryProperties = &queryProperties
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				q.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				q.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				q.Type = &typeVar
			}
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				q.Tags = tags
			}
		}
	}

	return nil
}

// QueryColumn ...
type QueryColumn struct {
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
}

// QueryProperties ...
type QueryProperties struct {
	NextLink *string `json:"nextLink,omitempty"`
	// Columns - Array of columns
	Columns *[]QueryColumn   `json:"columns,omitempty"`
	Rows    *[][]interface{} `json:"rows,omitempty"`
}

// QueryResult result of query. It contains all columns listed under groupings and aggregation.
type QueryResult struct {
	autorest.Response `json:"-"`
	// Value - READ-ONLY; The list of usage data.
	Value *[]Query `json:"value,omitempty"`
}

// ReportConfig a report config resource.
type ReportConfig struct {
	autorest.Response       `json:"-"`
	*ReportConfigProperties `json:"properties,omitempty"`
	// ID - READ-ONLY; Resource Id.
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; Resource name.
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; Resource type.
	Type *string `json:"type,omitempty"`
	// Tags - READ-ONLY; Resource tags.
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for ReportConfig.
func (rc ReportConfig) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if rc.ReportConfigProperties != nil {
		objectMap["properties"] = rc.ReportConfigProperties
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for ReportConfig struct.
func (rc *ReportConfig) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var reportConfigProperties ReportConfigProperties
				err = json.Unmarshal(*v, &reportConfigProperties)
				if err != nil {
					return err
				}
				rc.ReportConfigProperties = &reportConfigProperties
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				rc.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				rc.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				rc.Type = &typeVar
			}
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				rc.Tags = tags
			}
		}
	}

	return nil
}

// ReportConfigAggregation the aggregation expression to be used in the report.
type ReportConfigAggregation struct {
	// Name - The name of the column to aggregate.
	Name *string `json:"name,omitempty"`
	// Function - The name of the aggregation function to use.
	Function *string `json:"function,omitempty"`
}

// ReportConfigComparisonExpression the comparison expression to be used in the report.
type ReportConfigComparisonExpression struct {
	// Name - The name of the column to use in comparison.
	Name *string `json:"name,omitempty"`
	// Operator - The operator to use for comparison.
	Operator *string `json:"operator,omitempty"`
	// Values - Array of values to use for comparison
	Values *[]string `json:"values,omitempty"`
}

// ReportConfigDataset the definition of data present in the report.
type ReportConfigDataset struct {
	// Granularity - The granularity of rows in the report. Possible values include: 'Daily'
	Granularity GranularityType `json:"granularity,omitempty"`
	// Configuration - Has configuration information for the data in the report. The configuration will be ignored if aggregation and grouping are provided.
	Configuration *ReportConfigDatasetConfiguration `json:"configuration,omitempty"`
	// Aggregation - Dictionary of aggregation expression to use in the report. The key of each item in the dictionary is the alias for the aggregated column. Report can have up to 2 aggregation clauses.
	Aggregation map[string]*ReportConfigAggregation `json:"aggregation"`
	// Grouping - Array of group by expression to use in the report. Report can have up to 2 group by clauses.
	Grouping *[]ReportConfigGrouping `json:"grouping,omitempty"`
	// Filter - Has filter expression to use in the report.
	Filter *ReportConfigFilter `json:"filter,omitempty"`
}

// MarshalJSON is the custom marshaler for ReportConfigDataset.
func (rcd ReportConfigDataset) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if rcd.Granularity != "" {
		objectMap["granularity"] = rcd.Granularity
	}
	if rcd.Configuration != nil {
		objectMap["configuration"] = rcd.Configuration
	}
	if rcd.Aggregation != nil {
		objectMap["aggregation"] = rcd.Aggregation
	}
	if rcd.Grouping != nil {
		objectMap["grouping"] = rcd.Grouping
	}
	if rcd.Filter != nil {
		objectMap["filter"] = rcd.Filter
	}
	return json.Marshal(objectMap)
}

// ReportConfigDatasetConfiguration the configuration of dataset in the report.
type ReportConfigDatasetConfiguration struct {
	// Columns - Array of column names to be included in the report. Any valid report column name is allowed. If not provided, then report includes all columns.
	Columns *[]string `json:"columns,omitempty"`
}

// ReportConfigDefinition the definition of a report config.
type ReportConfigDefinition struct {
	// Type - The type of the report.
	Type *string `json:"type,omitempty"`
	// Timeframe - The time frame for pulling data for the report. If custom, then a specific time period must be provided. Possible values include: 'WeekToDate', 'MonthToDate', 'YearToDate', 'Custom'
	Timeframe TimeframeType `json:"timeframe,omitempty"`
	// TimePeriod - Has time period for pulling data for the report.
	TimePeriod *ReportConfigTimePeriod `json:"timePeriod,omitempty"`
	// Dataset - Has definition for data in this report config.
	Dataset *ReportConfigDataset `json:"dataset,omitempty"`
}

// ReportConfigDeliveryDestination the destination information for the delivery of the report.
type ReportConfigDeliveryDestination struct {
	// ResourceID - The resource id of the storage account where reports will be delivered.
	ResourceID *string `json:"resourceId,omitempty"`
	// Container - The name of the container where reports will be uploaded.
	Container *string `json:"container,omitempty"`
	// RootFolderPath - The name of the directory where reports will be uploaded.
	RootFolderPath *string `json:"rootFolderPath,omitempty"`
}

// ReportConfigDeliveryInfo the delivery information associated with a report config.
type ReportConfigDeliveryInfo struct {
	// Destination - Has destination for the report being delivered.
	Destination *ReportConfigDeliveryDestination `json:"destination,omitempty"`
}

// ReportConfigFilter the filter expression to be used in the report.
type ReportConfigFilter struct {
	// And - The logical "AND" expression. Must have at least 2 items.
	And *[]ReportConfigFilter `json:"and,omitempty"`
	// Or - The logical "OR" expression. Must have at least 2 items.
	Or *[]ReportConfigFilter `json:"or,omitempty"`
	// Not - The logical "NOT" expression.
	Not *ReportConfigFilter `json:"not,omitempty"`
	// Dimension - Has comparison expression for a dimension
	Dimension *ReportConfigComparisonExpression `json:"dimension,omitempty"`
	// Tag - Has comparison expression for a tag
	Tag *ReportConfigComparisonExpression `json:"tag,omitempty"`
}

// ReportConfigGrouping the group by expression to be used in the report.
type ReportConfigGrouping struct {
	// ColumnType - Has type of the column to group. Possible values include: 'ReportConfigColumnTypeTag', 'ReportConfigColumnTypeDimension'
	ColumnType ReportConfigColumnType `json:"columnType,omitempty"`
	// Name - The name of the column to group.
	Name *string `json:"name,omitempty"`
}

// ReportConfigListResult result of listing report configs. It contains a list of available report
// configurations in the scope provided.
type ReportConfigListResult struct {
	autorest.Response `json:"-"`
	// Value - READ-ONLY; The list of report configs.
	Value *[]ReportConfig `json:"value,omitempty"`
}

// ReportConfigProperties the properties of the report config.
type ReportConfigProperties struct {
	// Schedule - Has schedule information for the report config.
	Schedule *ReportConfigSchedule `json:"schedule,omitempty"`
	// Format - The format of the report being delivered. Possible values include: 'Csv'
	Format FormatType `json:"format,omitempty"`
	// DeliveryInfo - Has delivery information for the report config.
	DeliveryInfo *ReportConfigDeliveryInfo `json:"deliveryInfo,omitempty"`
	// Definition - Has definition for the report config.
	Definition *ReportConfigDefinition `json:"definition,omitempty"`
}

// ReportConfigRecurrencePeriod the start and end date for recurrence schedule.
type ReportConfigRecurrencePeriod struct {
	// From - The start date of recurrence.
	From *date.Time `json:"from,omitempty"`
	// To - The end date of recurrence. If not provided, we default this to 10 years from the start date.
	To *date.Time `json:"to,omitempty"`
}

// ReportConfigSchedule the schedule associated with a report config.
type ReportConfigSchedule struct {
	// Status - The status of the schedule. Whether active or not. If inactive, the report's scheduled execution is paused. Possible values include: 'Active', 'Inactive'
	Status StatusType `json:"status,omitempty"`
	// Recurrence - The schedule recurrence. Possible values include: 'RecurrenceTypeDaily', 'RecurrenceTypeWeekly', 'RecurrenceTypeMonthly', 'RecurrenceTypeAnnually'
	Recurrence RecurrenceType `json:"recurrence,omitempty"`
	// RecurrencePeriod - Has start and end date of the recurrence. The start date must be in future. If present, the end date must be greater than start date.
	RecurrencePeriod *ReportConfigRecurrencePeriod `json:"recurrencePeriod,omitempty"`
}

// ReportConfigTimePeriod the start and end date for pulling data for the report.
type ReportConfigTimePeriod struct {
	// From - The start date to pull data from.
	From *date.Time `json:"from,omitempty"`
	// To - The end date to pull data to.
	To *date.Time `json:"to,omitempty"`
}

// Resource the Resource model definition.
type Resource struct {
	// ID - READ-ONLY; Resource Id.
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; Resource name.
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; Resource type.
	Type *string `json:"type,omitempty"`
	// Tags - READ-ONLY; Resource tags.
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for Resource.
func (r Resource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}
