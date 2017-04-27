package aws

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccAWSSssParameterStore_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAwsSsmParameterStoreConfig_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAwsSsmParameterStore("data.aws_ssm_parameter_store.key"),
				),
			},
		},
	})
}

func testAccCheckAwsIamAccountAlias(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Can't find SSM Parameter Store resource: %s", n)
		}

		if rs.Primary.Attributes["key"] == "" {
			return fmt.Errorf("Missing parameter key")
		}

		return nil
	}
}

const testAccCheckAwsSsmParameterStoreConfig_basic = `
data "aws_ssm_parmeter_store" "key" {
  key = "Hello"
}
`
