// Code generated by go-swagger; DO NOT EDIT.

package m_v_sense

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new m v sense API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for m v sense API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetDeviceCameraAnalyticsLive(params *GetDeviceCameraAnalyticsLiveParams, authInfo runtime.ClientAuthInfoWriter) (*GetDeviceCameraAnalyticsLiveOK, error)

	GetDeviceCameraAnalyticsOverview(params *GetDeviceCameraAnalyticsOverviewParams, authInfo runtime.ClientAuthInfoWriter) (*GetDeviceCameraAnalyticsOverviewOK, error)

	GetDeviceCameraAnalyticsRecent(params *GetDeviceCameraAnalyticsRecentParams, authInfo runtime.ClientAuthInfoWriter) (*GetDeviceCameraAnalyticsRecentOK, error)

	GetDeviceCameraAnalyticsZoneHistory(params *GetDeviceCameraAnalyticsZoneHistoryParams, authInfo runtime.ClientAuthInfoWriter) (*GetDeviceCameraAnalyticsZoneHistoryOK, error)

	GetDeviceCameraAnalyticsZones(params *GetDeviceCameraAnalyticsZonesParams, authInfo runtime.ClientAuthInfoWriter) (*GetDeviceCameraAnalyticsZonesOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetDeviceCameraAnalyticsLive gets device camera analytics live

  Returns live state from camera of analytics zones
*/
func (a *Client) GetDeviceCameraAnalyticsLive(params *GetDeviceCameraAnalyticsLiveParams, authInfo runtime.ClientAuthInfoWriter) (*GetDeviceCameraAnalyticsLiveOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDeviceCameraAnalyticsLiveParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getDeviceCameraAnalyticsLive",
		Method:             "GET",
		PathPattern:        "/devices/{serial}/camera/analytics/live",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDeviceCameraAnalyticsLiveReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetDeviceCameraAnalyticsLiveOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getDeviceCameraAnalyticsLive: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetDeviceCameraAnalyticsOverview gets device camera analytics overview

  Returns an overview of aggregate analytics data for a timespan
*/
func (a *Client) GetDeviceCameraAnalyticsOverview(params *GetDeviceCameraAnalyticsOverviewParams, authInfo runtime.ClientAuthInfoWriter) (*GetDeviceCameraAnalyticsOverviewOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDeviceCameraAnalyticsOverviewParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getDeviceCameraAnalyticsOverview",
		Method:             "GET",
		PathPattern:        "/devices/{serial}/camera/analytics/overview",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDeviceCameraAnalyticsOverviewReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetDeviceCameraAnalyticsOverviewOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getDeviceCameraAnalyticsOverview: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetDeviceCameraAnalyticsRecent gets device camera analytics recent

  Returns most recent record for analytics zones
*/
func (a *Client) GetDeviceCameraAnalyticsRecent(params *GetDeviceCameraAnalyticsRecentParams, authInfo runtime.ClientAuthInfoWriter) (*GetDeviceCameraAnalyticsRecentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDeviceCameraAnalyticsRecentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getDeviceCameraAnalyticsRecent",
		Method:             "GET",
		PathPattern:        "/devices/{serial}/camera/analytics/recent",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDeviceCameraAnalyticsRecentReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetDeviceCameraAnalyticsRecentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getDeviceCameraAnalyticsRecent: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetDeviceCameraAnalyticsZoneHistory gets device camera analytics zone history

  Return historical records for analytic zones
*/
func (a *Client) GetDeviceCameraAnalyticsZoneHistory(params *GetDeviceCameraAnalyticsZoneHistoryParams, authInfo runtime.ClientAuthInfoWriter) (*GetDeviceCameraAnalyticsZoneHistoryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDeviceCameraAnalyticsZoneHistoryParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getDeviceCameraAnalyticsZoneHistory",
		Method:             "GET",
		PathPattern:        "/devices/{serial}/camera/analytics/zones/{zoneId}/history",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDeviceCameraAnalyticsZoneHistoryReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetDeviceCameraAnalyticsZoneHistoryOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getDeviceCameraAnalyticsZoneHistory: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetDeviceCameraAnalyticsZones gets device camera analytics zones

  Returns all configured analytic zones for this camera
*/
func (a *Client) GetDeviceCameraAnalyticsZones(params *GetDeviceCameraAnalyticsZonesParams, authInfo runtime.ClientAuthInfoWriter) (*GetDeviceCameraAnalyticsZonesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDeviceCameraAnalyticsZonesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getDeviceCameraAnalyticsZones",
		Method:             "GET",
		PathPattern:        "/devices/{serial}/camera/analytics/zones",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDeviceCameraAnalyticsZonesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetDeviceCameraAnalyticsZonesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getDeviceCameraAnalyticsZones: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}