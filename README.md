# Be2bill Merchant API

[![Build Status](https://travis-ci.org/be2bill/go-be2bill.svg?branch=master)](https://travis-ci.org/be2bill/go-be2bill)
[![GoDoc](https://godoc.org/github.com/be2bill/go-be2bill?status.svg)](https://godoc.org/github.com/be2bill/go-be2bill)
[![Coverage](http://gocover.io/_badge/github.com/be2bill/go-be2bill?1)](http://gocover.io/github.com/be2bill/go-be2bill)
[![Report Card](http://goreportcard.com/badge/be2bill/go-be2bill)](http://goreportcard.com/report/be2bill/go-be2bill)

This is a Go language library for the Be2bill merchant API that supports
the Form and DirectLink client APIs.

This library closely adheres to the official [Merchant API guidelines](https://github.com/be2bill/merchant-api-guidelines).

## Installation

Installing go-be2bill may be done via the usual go get procedure:

    $ go get github.com/be2bill/go-be2bill

## Documentation

See https://godoc.org/github.com/be2bill/go-be2bill

## Examples of use

In all cases, you need to import the go-be2bill package:

    import "github.com/be2bill/go-be2bill"

### Form Client

Every initial transaction is made using a secure form.
The Form Client API exposes methods that return the HTML code
for payment or authorization buttons.

The first step is to build a client for the given environment, using
your credentials:

	client := be2bill.BuildSandboxFormClient("test", "password")

To build a payment form button, call:

	button := client.BuildPaymentFormButton(
		be2bill.SingleAmount(15235),   // amount in cents
		"order_1412327697",            // order ID
		"6328_john.smith@example.org", // user ID
		"Fashion jacket",              // order description
		be2bill.Options{
			be2bill.HTMLOptionSubmit: be2bill.Options{
				"value": "Pay with be2bill",
				"class": "flatButton",
			},
			be2bill.HTMLOptionForm: be2bill.Options{"id": "myform"},
		},                             // HTML options for the form
		be2bill.Options{
			be2bill.ParamClientEmail: "john.smith@example.org",
			be2bill.Param3DSecure:    "yes",
		},                             // additional platform options
	)

Authorization form buttons are created similarily, except that the
method to call is `BuildAuthorizationFormButton` that takes the same
parameters.
An authorization must be captured using the `Capture` method of the
Direct Link Client API.

### Direct Link Client

All operations that do not require interactive data input from the client
can be made using HTTP POST calls to the be2bill servers.
The Direct Link Client API is an abstraction layer for these calls.

Just like the Form Client API, the first step is to build a client for
the given environment, using your credentials:

	client := be2bill.BuildSandboxDirectLinkClient("test", "password")

Then, for example to capture a previously authorized transaction, call:

	result, err := client.Capture(
		"A151621",
		"order_1423675675",
		"capture_transaction_A151621",
		be2bill.Options{},
	)

## Testing

The library comes with a complete test suite which can be run using
the following command:

    $ go test -v

Please note that there are a few sandbox tests that depend on having
a be2bill test account configured, and are skipped by default.

The environment variables `BE2BILL_IDENTIFIER` and `BE2BILL_PASSWORD`
need to be present and correctly configured with a test account in order
to run the sandbox tests.

