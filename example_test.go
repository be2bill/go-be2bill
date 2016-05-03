// Copyright 2016 Marc Noirot. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package be2bill_test

import (
	"fmt"

	"github.com/be2bill/go-be2bill"
)

func ExampleFormClient_BuildPaymentFormButton_complete() {
	// build client
	client := be2bill.BuildSandboxFormClient("test", "password")

	// create payment button
	button := client.BuildPaymentFormButton(
		be2bill.FragmentedAmount{"2010-05-14": 15235, "2012-06-04": 14723},
		"order_1412327697",
		"6328_john.smith@example.org",
		"Fashion jacket",
		be2bill.Options{
			be2bill.HTMLOptionSubmit: be2bill.Options{
				"value": "Pay with be2bill",
				"class": "flatButton",
			},
			be2bill.HTMLOptionForm: be2bill.Options{"id": "myform"},
		},
		be2bill.Options{
			be2bill.ParamClientEmail: "toto@example.org",
			be2bill.Param3DSecure:    "yes",
		},
	)

	// display the button's source code
	fmt.Println(button)

	// Output:
	// <form method="post" action="https://secure-test.be2bill.com/front/form/process" id="myform">
	//   <input type="hidden" name="3DSECURE" value="yes" />
	//   <input type="hidden" name="AMOUNTS[2010-05-14]" value="15235" />
	//   <input type="hidden" name="AMOUNTS[2012-06-04]" value="14723" />
	//   <input type="hidden" name="CLIENTEMAIL" value="toto@example.org" />
	//   <input type="hidden" name="CLIENTIDENT" value="6328_john.smith@example.org" />
	//   <input type="hidden" name="DESCRIPTION" value="Fashion jacket" />
	//   <input type="hidden" name="HASH" value="e4e3c4ab88774536108b85ccd62735bf1c1a6825a87d0fcbd7efa2ece12670e2" />
	//   <input type="hidden" name="IDENTIFIER" value="test" />
	//   <input type="hidden" name="OPERATIONTYPE" value="payment" />
	//   <input type="hidden" name="ORDERID" value="order_1412327697" />
	//   <input type="hidden" name="VERSION" value="2.0" />
	//   <input type="submit" class="flatButton" value="Pay with be2bill" />
	// </form>
}

func ExampleFormClient_BuildPaymentFormButton_simple() {
	// build client
	client := be2bill.BuildSandboxFormClient("test", "password")

	// create payment button
	button := client.BuildPaymentFormButton(
		be2bill.SingleAmount(15235),
		"order_1412327697",
		"6328_john.smith@example.org",
		"Fashion jacket",
		be2bill.Options{},
		be2bill.Options{},
	)

	// display the button's source code
	fmt.Println(button)

	// Output:
	// <form method="post" action="https://secure-test.be2bill.com/front/form/process">
	//   <input type="hidden" name="AMOUNT" value="15235" />
	//   <input type="hidden" name="CLIENTIDENT" value="6328_john.smith@example.org" />
	//   <input type="hidden" name="DESCRIPTION" value="Fashion jacket" />
	//   <input type="hidden" name="HASH" value="fab8f17da3e0f8315168cffc87c5cc28dbd29698c102d19e9f548bec42d16029" />
	//   <input type="hidden" name="IDENTIFIER" value="test" />
	//   <input type="hidden" name="OPERATIONTYPE" value="payment" />
	//   <input type="hidden" name="ORDERID" value="order_1412327697" />
	//   <input type="hidden" name="VERSION" value="2.0" />
	//   <input type="submit" />
	// </form>
}

func ExampleFormClient_BuildAuthorizationFormButton_complete() {
	// build client
	client := be2bill.BuildSandboxFormClient("test", "password")

	// create payment button
	button := client.BuildAuthorizationFormButton(
		15235,
		"order_1412327697",
		"6328_john.smith@example.org",
		"Fashion jacket",
		be2bill.Options{
			be2bill.HTMLOptionSubmit: be2bill.Options{
				"value": "Pay with be2bill",
				"class": "flatButton",
			},
			be2bill.HTMLOptionForm: be2bill.Options{"id": "myform"},
		},
		be2bill.Options{
			be2bill.ParamClientEmail: "toto@example.org",
			be2bill.Param3DSecure:    "yes",
		},
	)

	// display the button's source code
	fmt.Println(button)

	// Output:
	// <form method="post" action="https://secure-test.be2bill.com/front/form/process" id="myform">
	//   <input type="hidden" name="3DSECURE" value="yes" />
	//   <input type="hidden" name="AMOUNT" value="15235" />
	//   <input type="hidden" name="CLIENTEMAIL" value="toto@example.org" />
	//   <input type="hidden" name="CLIENTIDENT" value="6328_john.smith@example.org" />
	//   <input type="hidden" name="DESCRIPTION" value="Fashion jacket" />
	//   <input type="hidden" name="HASH" value="5c22b8f55c84b21e6e6c213b8e4ef554779f785abee2ca8361096b6b0d95a9fd" />
	//   <input type="hidden" name="IDENTIFIER" value="test" />
	//   <input type="hidden" name="OPERATIONTYPE" value="authorization" />
	//   <input type="hidden" name="ORDERID" value="order_1412327697" />
	//   <input type="hidden" name="VERSION" value="2.0" />
	//   <input type="submit" class="flatButton" value="Pay with be2bill" />
	// </form>
}

func ExampleFormClient_BuildAuthorizationFormButton_simple() {
	// build client
	client := be2bill.BuildSandboxFormClient("test", "password")

	// create payment button
	button := client.BuildAuthorizationFormButton(
		15235,
		"order_1412327697",
		"6328_john.smith@example.org",
		"Fashion jacket",
		be2bill.Options{},
		be2bill.Options{},
	)

	// display the button's source code
	fmt.Println(button)

	// Output:
	// <form method="post" action="https://secure-test.be2bill.com/front/form/process">
	//   <input type="hidden" name="AMOUNT" value="15235" />
	//   <input type="hidden" name="CLIENTIDENT" value="6328_john.smith@example.org" />
	//   <input type="hidden" name="DESCRIPTION" value="Fashion jacket" />
	//   <input type="hidden" name="HASH" value="01ccdb73b31de50567aa699642dad2e566a9c676d74d359efb4c849c13012427" />
	//   <input type="hidden" name="IDENTIFIER" value="test" />
	//   <input type="hidden" name="OPERATIONTYPE" value="authorization" />
	//   <input type="hidden" name="ORDERID" value="order_1412327697" />
	//   <input type="hidden" name="VERSION" value="2.0" />
	//   <input type="submit" />
	// </form>
}

func ExampleDirectLinkClient_Capture() {
	// build client
	client := be2bill.BuildSandboxDirectLinkClient("test", "password")

	result, err := client.Capture(
		"A151621",
		"order_1423675675",
		"capture_transaction_A151621",
		be2bill.Options{},
	)

	if err == nil {
		fmt.Println(result)
	}
}
