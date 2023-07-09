provider "aws" {
  region = "ap-northeast-2"
  default_tags {
    Project   = "terraforge"
    Env       = "dev"
    AccountId = "000000000000"
    Owner     = "vincent"
    Region    = "ap-northeast-2"
  }
  assume_role {
    role_arn = "arn:aws:iam::000000000000:role/service-provision-role"
  }
}

