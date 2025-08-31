output "id" {
  value       = module.datadog_child_organization.id
  description = "Organization ID"
}

output "description" {
  value       = module.datadog_child_organization.description
  description = "Description of the organization"
}

output "public_id" {
  value       = module.datadog_child_organization.public_id
  description = "Public ID of the organization"
}

output "user" {
  value       = module.datadog_child_organization.user
  description = "Information about organization users"
}

output "settings" {
  value       = module.datadog_child_organization.settings
  description = "Organization settings"
}

output "api_key" {
  value       = module.datadog_child_organization.api_key
  description = "Information about Datadog API key"
  sensitive   = true
}

output "application_key" {
  value       = module.datadog_child_organization.application_key
  description = "Datadog application key with its associated metadata"
  sensitive   = true
}
