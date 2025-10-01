locals {
  enabled = module.this.enabled
}


module "datadog_child_organization" {
  source  = "cloudposse/platform/datadog//modules/child_organization"
  version = "1.6.0"

  organization_name = var.organization_name

  context = module.this.context
}
