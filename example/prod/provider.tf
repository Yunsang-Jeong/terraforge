provider "aws" {
  region = "ap-northeast-2"
  default_tags {
    Env       = "prod"
    Account   = "srv_acct_2"
    AccountId = "000000000002"
    Owner     = "vincent"
    Region    = "ap-northeast-2"
    Project   = "terraforge"
  }
  assume_role {
    role_arn = "arn:aws:iam::000000000002:role/service-provision-role"
  }
}

provider "aws" {
  alias  = "srv_acct_1-apne2"
  region = "us-east-1"
  default_tags {
    Region    = "ap-northeast-2"
    Project   = "terraforge"
    Env       = "prod"
    Account   = "srv_acct_1"
    AccountId = "000000000001"
    Owner     = "vincent"
  }
  assume_role {
    role_arn = "arn:aws:iam::000000000001:role/service-provision-role"
  }
}
provider "aws" {
  alias  = "srv_acct_1-uea1"
  region = "us-east-1"
  default_tags {
    AccountId = "000000000001"
    Owner     = "vincent"
    Region    = "us-east-1"
    Project   = "terraforge"
    Env       = "prod"
    Account   = "srv_acct_1"
  }
  assume_role {
    role_arn = "arn:aws:iam::000000000001:role/service-provision-role"
  }
}
provider "aws" {
  alias  = "srv_acct_1-uea2"
  region = "us-east-1"
  default_tags {
    Project   = "terraforge"
    Env       = "prod"
    Account   = "srv_acct_1"
    AccountId = "000000000001"
    Owner     = "vincent"
    Region    = "us-east-2"
  }
  assume_role {
    role_arn = "arn:aws:iam::000000000001:role/service-provision-role"
  }
}
provider "aws" {
  alias  = "srv_acct_2-apne1"
  region = "us-east-1"
  default_tags {
    Region    = "ap-northeast-1"
    Project   = "terraforge"
    Env       = "prod"
    Account   = "srv_acct_2"
    AccountId = "000000000002"
    Owner     = "vincent"
  }
  assume_role {
    role_arn = "arn:aws:iam::000000000002:role/service-provision-role"
  }
}
provider "aws" {
  region = "us-east-1"
  alias  = "srv_acct_2-apne2"
  default_tags {
    Region    = "ap-northeast-2"
    Project   = "terraforge"
    Env       = "prod"
    Account   = "srv_acct_2"
    AccountId = "000000000002"
    Owner     = "vincent"
  }
  assume_role {
    role_arn = "arn:aws:iam::000000000002:role/service-provision-role"
  }
}
provider "aws" {
  region = "us-east-1"
  alias  = "srv_acct_2-uea1"
  default_tags {
    Env       = "prod"
    Account   = "srv_acct_2"
    AccountId = "000000000002"
    Owner     = "vincent"
    Region    = "us-east-1"
    Project   = "terraforge"
  }
  assume_role {
    role_arn = "arn:aws:iam::000000000002:role/service-provision-role"
  }
}
provider "aws" {
  alias  = "srv_acct_2-uea2"
  region = "us-east-1"
  default_tags {
    Region    = "us-east-2"
    Project   = "terraforge"
    Env       = "prod"
    Account   = "srv_acct_2"
    AccountId = "000000000002"
    Owner     = "vincent"
  }
  assume_role {
    role_arn = "arn:aws:iam::000000000002:role/service-provision-role"
  }
}
provider "aws" {
  alias  = "srv_acct_1-apne1"
  region = "us-east-1"
  default_tags {
    Env       = "prod"
    Account   = "srv_acct_1"
    AccountId = "000000000001"
    Owner     = "vincent"
    Region    = "ap-northeast-1"
    Project   = "terraforge"
  }
  assume_role {
    role_arn = "arn:aws:iam::000000000001:role/service-provision-role"
  }
}
