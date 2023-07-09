variable "metadata" {
  default = {
    account    = "srv_acct_2"
    account_id = "000000000002"
    aws_account_id_map = {
      srv_acct_1 = "000000000001"
      srv_acct_2 = "000000000002"
    }
    aws_region_map = {
      ap-northeast-1 = "apne1"
      ap-northeast-2 = "apne2"
      us-east-1      = "uea1"
      us-east-2      = "uea2"
    }
    env            = "prod"
    multi_provider = true
    owner          = "vincent"
    project        = "terraforge"
    region         = "ap-northeast-2"
  }
}
