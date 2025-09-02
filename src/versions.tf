terraform {
  required_version = ">= 1.0.0"

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = ">= 4.9.0, < 6.0.0"
    }
    datadog = {
      source = "datadog/datadog"
      # Child organizations are only available in v3.3.0 and above
      # https://github.com/DataDog/terraform-provider-datadog/issues/878#issuecomment-906383085
      version = ">= 3.3.0"
    }
  }
}
