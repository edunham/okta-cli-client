/*
Okta Admin Management

Allows customers to easily access the Okta Management APIs

API version: 5.1.0
Contact: devex-public@okta.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sdk

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// APIResponse stores the API response returned by the server.
type APIResponse struct {
	*http.Response
	cli *APIClient
	pg  Pagination
}

func newAPIResponse(r *http.Response, cli *APIClient, v interface{}) *APIResponse {
	var pg Pagination
	// switch v
	pg = newPaginationInHeader(r)
	response := &APIResponse{Response: r, cli: cli, pg: pg}
	return response
}

func buildResponse(resp *http.Response, cli *APIClient, v interface{}) (*APIResponse, error) {
	ct := resp.Header.Get("Content-Type")
	response := newAPIResponse(resp, cli, v)
	err := cli.checkResponseForError(resp)
	if err != nil {
		return response, err
	}
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	copyBodyBytes := make([]byte, len(bodyBytes))
	copy(copyBodyBytes, bodyBytes)
	_ = resp.Body.Close()                                    // close it to avoid memory leaks
	resp.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes)) // restore the original response body
	if len(copyBodyBytes) == 0 {
		return response, nil
	}
	switch {
	case strings.Contains(ct, "application/xml"):
		err = xml.NewDecoder(bytes.NewReader(copyBodyBytes)).Decode(v)
	case strings.Contains(ct, "application/json"):
		err = json.NewDecoder(bytes.NewReader(copyBodyBytes)).Decode(v)
	case strings.Contains(ct, "application/octet-stream"):
		// since the response is arbitrary binary data, we leave it to the user to decode it
		return response, nil
	default:
		return nil, errors.New("could not build a response for type: " + ct)
	}
	if err == io.EOF {
		err = nil
	}
	return response, err
}

func (c *APIClient) checkResponseForError(resp *http.Response) error {
	localVarBody, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	resp.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if resp.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: resp.Status,
		}
		if resp.StatusCode == 403 {
			var v Error
			err := c.decode(&v, localVarBody, resp.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return newErr
			}
			newErr.model = v
			return newErr
		}
		if resp.StatusCode == 404 {
			var v Error
			err := c.decode(&v, localVarBody, resp.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return newErr
			}
			newErr.model = v
			return newErr
		}
		if resp.StatusCode == 429 {
			var v Error
			err := c.decode(&v, localVarBody, resp.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return newErr
			}
			newErr.model = v
			return newErr
		}
	}
	return nil
}

func (res *APIResponse) Next(v interface{}) (*APIResponse, error) {
	if res.cli == nil {
		return nil, errors.New("no initial response provided from previous request")
	}
	req, err := res.cli.PrepareRequest(res.cli.cfg.Context, res.NextPage(), http.MethodGet, nil, map[string]string{"Accept": "application/json"}, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	resp, err := res.cli.Do(res.cli.cfg.Context, req)
	if err != nil {
		return nil, err
	}
	return buildResponse(resp, res.cli, v)
}

func (res *APIResponse) Self() string {
	return res.pg.Self()
}

func (res *APIResponse) NextPage() string {
	return res.pg.NextPage()
}

func (res *APIResponse) HasNextPage() bool {
	return res.pg.NextPage() != ""
}

type Pagination interface {
	Self() string
	NextPage() string
}

type PaginationInHeader struct {
	r *http.Response
}

func newPaginationInHeader(r *http.Response) *PaginationInHeader {
	return &PaginationInHeader{r: r}
}

func (pg *PaginationInHeader) Self() (self string) {
	links := pg.r.Header["Link"]
	if len(links) > 0 {
		for _, link := range links {
			splitLinkHeader := strings.Split(link, ";")
			if len(splitLinkHeader) < 2 {
				continue
			}
			rawLink := strings.TrimRight(strings.TrimLeft(splitLinkHeader[0], "<"), ">")
			rawURL, _ := url.Parse(rawLink)
			rawURL.Scheme = ""
			rawURL.Host = ""
			if pg.r.Request != nil {
				q := pg.r.Request.URL.Query()
				for k, v := range rawURL.Query() {
					q.Set(k, v[0])
				}
				rawURL.RawQuery = q.Encode()
			}
			if strings.Contains(link, `rel="self"`) {
				self = rawURL.String()
			}
		}
	}
	return
}

func (pg *PaginationInHeader) NextPage() (next string) {
	links := pg.r.Header["Link"]
	if len(links) > 0 {
		for _, link := range links {
			splitLinkHeader := strings.Split(link, ";")
			if len(splitLinkHeader) < 2 {
				continue
			}
			rawLink := strings.TrimRight(strings.TrimLeft(splitLinkHeader[0], "<"), ">")
			rawURL, _ := url.Parse(rawLink)
			rawURL.Scheme = ""
			rawURL.Host = ""
			if pg.r.Request != nil {
				q := pg.r.Request.URL.Query()
				for k, v := range rawURL.Query() {
					q.Set(k, v[0])
				}
				rawURL.RawQuery = q.Encode()
			}
			if strings.Contains(link, `rel="next"`) {
				next = rawURL.String()
			}
		}
	}
	return
}

type PaginationInBody struct{}
