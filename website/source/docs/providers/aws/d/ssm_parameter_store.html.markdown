---
layout: "aws"
page_title: "AWS: aws_ssm_parameter_store"
sidebar_current: "docs-aws-datasource-ssm-parameter-store"
description: |-
  Get value of ssm pamater store
---

# aws\_ssm\_parameter\_store

The SSM Parameter Store data source allows latest parameter value
for the effective account in which Terraform is working.

## Example Usage

```hcl
data "aws_ssm_parmeter_store" "key" {
  key = "Hello"
}

output "value" {
  value = "${data.aws_ssm_parmeter_store.key.value}"
}
```

## Argument Reference

* `key` - (Required) The key of the parameter. 
* `region` - (Option) Region of the key.

## Attributes Reference

The following attributes are exported:

* `value` - The key's latest value (already decrypted)
