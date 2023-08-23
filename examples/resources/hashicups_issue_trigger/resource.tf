resource "hashicups_issue_trigger" "my_issuetrigger" {
    channel_name = "...my_channel_name..."
            configs = {
        issue_trigger_backpressure_configs =     {
                delay = 6
                destinations = {
                    str = "..."
                }
            }
    }
            email = {}
            name = "Mrs. Sophie Smith MD"
            opsgenie = {}
            type = "delivery"
        }