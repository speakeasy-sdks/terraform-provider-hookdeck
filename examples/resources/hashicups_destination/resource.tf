resource "hashicups_destination" "my_destination" {
  auth_method = {
    api_key = {
      config = {
        api_key = "...my_api_key..."
        key     = "...my_key..."
        to      = "query"
      }
      type = "API_KEY"
    }
  }
  cli_path                 = "...my_cli_path..."
  http_method              = "<nil>"
  name                     = "Ricky Hoppe"
  path_forwarding_disabled = false
  rate_limit               = 5
  rate_limit_period        = "second"
  url                      = "...my_url..."
}