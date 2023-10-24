resource "hashicups_source" "my_source" {
  allowed_http_methods = [
    "GET",
  ]
  custom_response = {
    body         = "...my_body..."
    content_type = "json"
  }
  name = "Luther Bauch"
}