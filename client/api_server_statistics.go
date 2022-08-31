/*
GitHub v3 REST API

GitHub's v3 REST API.

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package github

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


// ServerStatisticsApiService ServerStatisticsApi service
type ServerStatisticsApiService service

type ApiEnterpriseAdminGetServerStatisticsRequest struct {
	ctx context.Context
	ApiService *ServerStatisticsApiService
	enterpriseOrOrg string
	dateStart *string
	dateEnd *string
}

// A cursor, as given in the [Link header](https://docs.github.com/rest/overview/resources-in-the-rest-api#link-header). If specified, the query only searches for events after this cursor.
func (r ApiEnterpriseAdminGetServerStatisticsRequest) DateStart(dateStart string) ApiEnterpriseAdminGetServerStatisticsRequest {
	r.dateStart = &dateStart
	return r
}

// A cursor, as given in the [Link header](https://docs.github.com/rest/overview/resources-in-the-rest-api#link-header). If specified, the query only searches for events before this cursor.
func (r ApiEnterpriseAdminGetServerStatisticsRequest) DateEnd(dateEnd string) ApiEnterpriseAdminGetServerStatisticsRequest {
	r.dateEnd = &dateEnd
	return r
}

func (r ApiEnterpriseAdminGetServerStatisticsRequest) Execute() ([]ServerStatisticsInner, *http.Response, error) {
	return r.ApiService.EnterpriseAdminGetServerStatisticsExecute(r)
}

/*
EnterpriseAdminGetServerStatistics Get GitHub Enterprise Server statistics

Returns aggregate usage metrics for your GitHub Enterprise Server 3.5+ instance for a specified time period up to 365 days.

To use this endpoint, your GitHub Enterprise Server instance must be connected to GitHub Enterprise Cloud using GitHub Connect. You must enable Server Statistics, and for the API request provide your enterprise account name or organization name connected to the GitHub Enterprise Server. For more information, see "[Enabling Server Statistics for your enterprise](/admin/configuration/configuring-github-connect/enabling-server-statistics-for-your-enterprise)" in the GitHub Enterprise Server documentation.

You'll need to use a personal access token:
  - If you connected your GitHub Enterprise Server to an enterprise account and enabled Server Statistics, you'll need a personal access token with the `read:enterprise` permission.
  - If you connected your GitHub Enterprise Server to an organization account and enabled Server Statistics, you'll need a personal access token with the `read:org` permission.

For more information on creating a personal access token, see "[Creating a personal access token](/authentication/keeping-your-account-and-data-secure/creating-a-personal-access-token)."

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param enterpriseOrOrg The slug version of the enterprise name or the login of an organization.
 @return ApiEnterpriseAdminGetServerStatisticsRequest
*/
func (a *ServerStatisticsApiService) EnterpriseAdminGetServerStatistics(ctx context.Context, enterpriseOrOrg string) ApiEnterpriseAdminGetServerStatisticsRequest {
	return ApiEnterpriseAdminGetServerStatisticsRequest{
		ApiService: a,
		ctx: ctx,
		enterpriseOrOrg: enterpriseOrOrg,
	}
}

// Execute executes the request
//  @return []ServerStatisticsInner
func (a *ServerStatisticsApiService) EnterpriseAdminGetServerStatisticsExecute(r ApiEnterpriseAdminGetServerStatisticsRequest) ([]ServerStatisticsInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []ServerStatisticsInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServerStatisticsApiService.EnterpriseAdminGetServerStatistics")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/enterprise-installation/{enterprise_or_org}/server-statistics"
	localVarPath = strings.Replace(localVarPath, "{"+"enterprise_or_org"+"}", url.PathEscape(parameterToString(r.enterpriseOrOrg, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.dateStart != nil {
		localVarQueryParams.Add("date_start", parameterToString(*r.dateStart, ""))
	}
	if r.dateEnd != nil {
		localVarQueryParams.Add("date_end", parameterToString(*r.dateEnd, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}