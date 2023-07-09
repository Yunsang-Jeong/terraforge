provider "aws" {
  region = "ap-northeast-2"
  default_tags {
    AccountId = "000000000001"
    Owner     = "vincent"
    Region    = "ap-northeast-2"
    Project   = "terraforge"
    Env       = "dev"
    Account   = "srv_acct_1"
  }
  assume_role {
    role_arn = "arn:aws:iam::000000000001:role/service-provision-role"
  }
}

