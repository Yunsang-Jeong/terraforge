provider "aws" {
  region = "ap-northeast-2"
  default_tags {
    Owner     = "vincent"
    Region    = "ap-northeast-2"
    Project   = "terraforge"
    Env       = "dev"
    Account   = "srv_acct_1"
    AccountId = "000000000001"
  }
  assume_role {
    role_arn = "arn:aws:iam::000000000001:role/service-provision-role"
  }
}

