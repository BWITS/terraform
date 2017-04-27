package aws

import (
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceAwsSsmParameterStore() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAwsSsmParameterStoreRead,

		Schema: map[string]*schema.Schema{
			"value": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceAwsSsmParameterStoreRead(d *schema.ResourceData, meta interface{}) error {

	key := d.Get("key")

	sess := session.Must(session.NewSession())

	svc := ssm.New(sess)

	params := &ssm.GetParametersInput{
		Names: []*string{ // Required
			aws.String(key), // Required
			// More values...
		},
		WithDecryption: aws.Bool(true),
	}
	resp, err := svc.GetParameters(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}
