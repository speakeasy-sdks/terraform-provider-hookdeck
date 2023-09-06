resource "hashicups_connection" "my_connection" {
  allowed_http_methods = [
    "PUT",
  ]
  auth_method = {
    hookdeck_signature = {
      config = {}
      type   = "HOOKDECK_SIGNATURE"
    }
  }
  body                     = "...my_body..."
  cli_path                 = "...my_cli_path..."
  content_type             = "text"
  destination_id           = "...my_destination_id..."
  http_method              = "PATCH"
  is_team_default          = true
  name                     = "Ismael Little"
  path_forwarding_disabled = true
  rate_limit               = 4
  rate_limit_period        = "minute"
  rules = [
    {
      strategy = "each_attempt"
      type     = "alert"
    },
  ]
  ruleset_id = "...my_ruleset_id..."
  source_id  = "...my_source_id..."
  url        = "...my_url..."
}