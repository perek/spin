/*
 * Spinnaker API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"io/ioutil"
	"net/url"
	"net/http"
	"strings"
	"golang.org/x/net/context"
	"encoding/json"
	"fmt"
)

// Linger please
var (
	_ context.Context
)

type ImageControllerApiService service


/* ImageControllerApiService Retrieve a list of images, filtered by cloud provider, region, and account
 The query parameter &#x60;q&#x60; filters the list of images by image name
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "account" (string) account
     @param "count" (int32) count
     @param "provider" (string) provider
     @param "q" (string) q
     @param "region" (string) region
 @return []interface{}*/
func (a *ImageControllerApiService) FindImagesUsingGET(ctx context.Context, localVarOptionals map[string]interface{}) ([]interface{},  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  []interface{}
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/images/find"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["account"], "string", "account"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["count"], "int32", "count"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["provider"], "string", "provider"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["q"], "string", "q"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["region"], "string", "region"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["account"].(string); localVarOk {
		localVarQueryParams.Add("account", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["count"].(int32); localVarOk {
		localVarQueryParams.Add("count", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["provider"].(string); localVarOk {
		localVarQueryParams.Add("provider", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["q"].(string); localVarOk {
		localVarQueryParams.Add("q", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["region"].(string); localVarOk {
		localVarQueryParams.Add("region", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"*/*",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* ImageControllerApiService Find tags
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param account account
 @param repository repository
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "xRateLimitApp" (string) X-RateLimit-App
     @param "provider" (string) provider
 @return []interface{}*/
func (a *ImageControllerApiService) FindTagsUsingGET(ctx context.Context, account string, repository string, localVarOptionals map[string]interface{}) ([]interface{},  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  []interface{}
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/images/tags"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["xRateLimitApp"], "string", "xRateLimitApp"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["provider"], "string", "provider"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("account", parameterToString(account, ""))
	if localVarTempParam, localVarOk := localVarOptionals["provider"].(string); localVarOk {
		localVarQueryParams.Add("provider", parameterToString(localVarTempParam, ""))
	}
	localVarQueryParams.Add("repository", parameterToString(repository, ""))
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"*/*",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarTempParam, localVarOk := localVarOptionals["xRateLimitApp"].(string); localVarOk {
		localVarHeaderParams["X-RateLimit-App"] = parameterToString(localVarTempParam, "")
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* ImageControllerApiService Get image details
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param account account
 @param imageId imageId
 @param region region
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "xRateLimitApp" (string) X-RateLimit-App
     @param "provider" (string) provider
 @return []interface{}*/
func (a *ImageControllerApiService) GetImageDetailsUsingGET(ctx context.Context, account string, imageId string, region string, localVarOptionals map[string]interface{}) ([]interface{},  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  []interface{}
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/images/{account}/{region}/{imageId}"
	localVarPath = strings.Replace(localVarPath, "{"+"account"+"}", fmt.Sprintf("%v", account), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"imageId"+"}", fmt.Sprintf("%v", imageId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"region"+"}", fmt.Sprintf("%v", region), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["xRateLimitApp"], "string", "xRateLimitApp"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["provider"], "string", "provider"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["provider"].(string); localVarOk {
		localVarQueryParams.Add("provider", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"*/*",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarTempParam, localVarOk := localVarOptionals["xRateLimitApp"].(string); localVarOk {
		localVarHeaderParams["X-RateLimit-App"] = parameterToString(localVarTempParam, "")
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

