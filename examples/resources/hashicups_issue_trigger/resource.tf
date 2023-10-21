resource "hashicups_issue_trigger" "my_issuetrigger" {
  channels = {
    email    = {}
    opsgenie = {}
    slack = {
      channel_name = "...my_channel_name..."
    }
  }
  configs = {
    issue_trigger_backpressure_configs = {
      delay = 3
      destinations = {
        str = "..."
      }
    }
  }
  name = "Ronnie Huels"
  type = "backpressure"
}