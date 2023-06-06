package fhttp

import (
	"SMAdmin/config"
	"SMAdmin/pkg/utils"
	"fmt"

	"time"

	"github.com/avast/retry-go"
	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpproxy"
)

type Client struct {
	cfg *config.Config
}

func NewClient(cfg *config.Config) *Client {
	return &Client{
		cfg: cfg,
	}
}

var timeout = time.Second * 30
var timeoutBitokk = time.Minute * 5

func (h *Client) RequestTimeout(
	method string,
	url string,
	body []byte,
	queryParams map[string]string,
	headers map[string]string,
) (responseBody []byte, statusCode int, err error) {
	req := fasthttp.AcquireRequest()
	req.Header.SetMethod(method)

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	var queryCount int
	for k, v := range queryParams {
		switch queryCount {
		case 0:
			url += fmt.Sprintf("?%s=%s", k, v)
		default:
			url += fmt.Sprintf("&%s=%s", k, v)
		}
		queryCount++
	}

	req.SetBody(body)
	req.SetRequestURI(url)
	res := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(res)

	client := &fasthttp.Client{}

	if err = client.DoTimeout(req, res, timeoutBitokk); err != nil {
		return
	}
	req.SetConnectionClose()
	statusCode = res.StatusCode()

	responseBody = make([]byte, len(res.Body()))
	copy(responseBody, res.Body())
	return
}

func (h *Client) Request(
	method string,
	url string,
	body []byte,
	queryParams map[string]string,
	headers map[string]string,
) (responseBody []byte, statusCode int, err error) {
	req := fasthttp.AcquireRequest()
	req.Header.SetMethod(method)
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	var queryCount int
	for k, v := range queryParams {
		switch queryCount {
		case 0:
			url += fmt.Sprintf("?%s=%s", k, v)
		default:
			url += fmt.Sprintf("&%s=%s", k, v)
		}
		queryCount++
	}

	req.SetBody(body)
	req.SetRequestURI(url)
	res := fasthttp.AcquireResponse()

	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(res)

	client := &fasthttp.Client{}

	if err = client.DoTimeout(req, res, timeout); err != nil {
		return
	}
	req.SetConnectionClose()
	statusCode = res.StatusCode()

	responseBody = make([]byte, len(res.Body()))
	copy(responseBody, res.Body())
	return
}

func (h *Client) RequestProxy(
	method string,
	url string,
	proxy string,
	body []byte,
	queryParams map[string]string,
	headers map[string]string,
) (responseBody []byte, statusCode int, err error) {
	req := fasthttp.AcquireRequest()
	req.Header.SetMethod(method)
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	var queryCount int
	for k, v := range queryParams {
		switch queryCount {
		case 0:
			url += fmt.Sprintf("?%s=%s", k, v)
		default:
			url += fmt.Sprintf("&%s=%s", k, v)
		}
		queryCount++
	}

	req.SetBody(body)
	req.SetRequestURI(url)
	res := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(res)

	client := &fasthttp.Client{
		Dial: fasthttpproxy.FasthttpHTTPDialer(proxy),
	}

	req.SetConnectionClose()
	if err = client.DoTimeout(req, res, timeout); err != nil {
		return
	}

	statusCode = res.StatusCode()

	responseBody = make([]byte, len(res.Body()))
	copy(responseBody, res.Body())
	return
}

func (h *Client) RequestProxyWithTimeout(
	method string,
	url string,
	proxy string,
	body []byte,
	queryParams map[string]string,
	headers map[string]string,
	timeoutRequest time.Duration,
) (responseBody []byte, statusCode int, err error) {
	req := fasthttp.AcquireRequest()
	req.Header.SetMethod(method)
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	var queryCount int
	for k, v := range queryParams {
		switch queryCount {
		case 0:
			url += fmt.Sprintf("?%s=%s", k, v)
		default:
			url += fmt.Sprintf("&%s=%s", k, v)
		}
		queryCount++
	}

	req.SetBody(body)
	req.SetRequestURI(url)
	res := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(res)

	client := &fasthttp.Client{
		Dial: fasthttpproxy.FasthttpHTTPDialer(proxy),
	}

	req.SetConnectionClose()
	if err = client.DoTimeout(req, res, timeoutRequest); err != nil {
		return
	}

	statusCode = res.StatusCode()

	responseBody = make([]byte, len(res.Body()))
	copy(responseBody, res.Body())
	return
}

func (h *Client) RequestProxyWithRetry(
	args RequestProxyWithRetryArgs,
) (responseBody []byte, statusCode int, err error) {
	err = retry.Do(
		func() error {
			responseBody, statusCode, err = h.RequestProxy(
				args.Method,
				args.Url,
				args.Proxy,
				args.Body,
				args.QueryParams,
				args.Headers,
			)
			if err != nil {
				return errors.Wrapf(err, "fhttp.RequestProxyWithRetry(args: %s)", utils.GetStructJSON(args))
			}

			if statusCode != args.ExpectedStatusCode {
				return errors.Wrapf(
					ErrNotExpectedStatusCode,
					"fhttp.RequestProxyWithRetry, expeced %d, got %d statusCode; requestArgs: %s, responseBody: %s",
					args.ExpectedStatusCode,
					statusCode,
					utils.GetStructJSON(args),
					responseBody,
				)
			}
			return nil
		},
		retry.Delay(args.RetryDelay*time.Millisecond),
		retry.Attempts(args.RetryAttempts),
	)
	return
}

func (h *Client) RequestProxyWithRetryWithTimeout(
	args RequestProxyWithRetryArgs, timeout time.Duration,
) (responseBody []byte, statusCode int, err error) {
	err = retry.Do(
		func() error {
			responseBody, statusCode, err = h.RequestProxyWithTimeout(
				args.Method,
				args.Url,
				args.Proxy,
				args.Body,
				args.QueryParams,
				args.Headers,
				timeout,
			)
			if err != nil {
				return errors.Wrapf(err, "fhttp.RequestProxyWithRetry(args: %s)", utils.GetStructJSON(args))
			}

			if statusCode != args.ExpectedStatusCode {
				return errors.Wrapf(
					ErrNotExpectedStatusCode,
					"fhttp.RequestProxyWithRetry, expeced %d, got %d statusCode; requestArgs: %s, responseBody: %s",
					args.ExpectedStatusCode,
					statusCode,
					utils.GetStructJSON(args),
					responseBody,
				)
			}
			return nil
		},
		retry.Delay(args.RetryDelay*time.Millisecond),
		retry.Attempts(args.RetryAttempts),
	)
	return
}
