package be2bill_test

import (
	"fmt"

	"github.com/be2bill/go-be2bill"
)

func ExampleCompleteForm() {
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

func ExampleSimpleForm() {
	// build client
	client := be2bill.BuildSandboxFormClient("test", "password")

	// create payment button
	button := client.BuildPaymentFormButton(
		be2bill.SingleAmount(15235),
		"order_1412327697",
		"6328_john.smith@example.org",
		"Fashion jacket",
		be2bill.DefaultOptions,
		be2bill.DefaultOptions,
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

func ExampleDirectLink() {
	// build client
	client := be2bill.BuildSandboxDirectLinkClient("test", "password")

	result, err := client.Capture(
		"A151621",
		"order_1423675675",
		"capture_transaction_A151621",
		be2bill.DefaultOptions,
	)

	if err == nil {
		fmt.Println(result)
	}
}
