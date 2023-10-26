resource "hashicups_connection" "my_connection" {
  destination = {
    archived_at = "2022-10-27T06:03:30.372Z"
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
    created_at               = "2021-05-15T01:52:35.935Z"
    http_method              = "PATCH"
    id                       = "81fd7b76-4e31-4e47-9cb1-f3659158ccc6"
    name                     = "Shelia Romaguera"
    path_forwarding_disabled = false
    rate_limit               = 0
    rate_limit_period        = "second"
    team_id                  = "...my_team_id..."
    updated_at               = "2022-09-05T00:53:44.714Z"
    url                      = "...my_url..."
  }
  destination_id = "...my_destination_id..."
  name           = "Julia Trantow DDS"
  rules = [
    {
      strategy = "each_attempt"
      type     = "alert"
    },
  ]
  ruleset = {
    archived_at     = "2022-06-20T02:50:35.009Z"
    created_at      = "2021-11-17T06:47:57.631Z"
    id              = "c2b5ed5f-0cf0-47c6-a222-08210d5a2814"
    is_team_default = true
    name            = "Pete Runolfsson"
    rules = [
      {
        strategy = "last_attempt"
        type     = "alert"
      },
    ]
    team_id    = "...my_team_id..."
    updated_at = "2022-05-18T07:09:08.239Z"
  }
  ruleset_id = "...my_ruleset_id..."
  source = {
    allowed_http_methods = [
      "GET",
    ]
    archived_at = "2022-03-11T21:42:25.382Z"
    created_at  = "2021-12-25T12:18:23.967Z"
    custom_response = {
      body         = "...my_body..."
      content_type = "json"
    }
    id = "a49098c5-62b9-4741-9b36-0edc1d66b24b"
    integration = {
      features = [
        "POLLING",
      ]
      id       = "8cfb7216-1b24-4a4a-9fe0-9a5b24ee84a3"
      label    = "...my_label..."
      provider = "aws_sns"
    }
    integration_id = "...my_integration_id..."
    name           = "Edmond Gerhold"
    team_id        = "...my_team_id..."
    updated_at     = "2021-06-08T01:47:41.484Z"
    url            = "...my_url..."
  }
  source_id = "...my_source_id..."
}