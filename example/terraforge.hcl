metafile {
  path = [
    "metadata.hcl",
    "root.hcl",
  ]
}

generate "variable" "metadata" {
  config {
    default = merge(metadata, {
      account_id = lookup(metadata.aws_account_id_map, metadata.account)
    })
  }
}

generate "terraform" {
  config {
    required_providers {
      aws = {
        source  = "hashicorp/aws"
        version = "> 5.0.0"
      }
    }
  }
}

generate "provider" "aws" {
  config {
    region = metadata.region
    default_tags {
      Account   = metadata.account
      AccountId = lookup(metadata.aws_account_id_map, metadata.account)
      Owner     = metadata.owner
      Region    = metadata.region
      Project   = metadata.project
      Env       = metadata.env
    }
    assume_role {
      role_arn = format(
        "arn:aws:iam::%s:role/service-provision-role",
        lookup(metadata.aws_account_id_map, metadata.account)
      )
    }
  }
}

generate "provider" "aws" {
  when = metadata.multi_provider
  for_each = {
    for m in setproduct(keys(metadata.aws_account_id_map), keys(metadata.aws_region_map)) :
    "${m[0]}-${lookup(metadata.aws_region_map, m[1])}" => {
      account    = m[0]
      account_id = lookup(metadata.aws_account_id_map, m[0])
      region     = m[1]
    }
  }
  config {
    alias  = each.key
    region = "us-east-1"
    default_tags {
      Account   = each.value.account
      AccountId = each.value.account_id
      Owner     = metadata.owner
      Region    = each.value.region
      Project   = metadata.project
      Env       = metadata.env
    }
    assume_role {
      role_arn = "arn:aws:iam::${each.value.account_id}:role/service-provision-role"
    }
  }
}

