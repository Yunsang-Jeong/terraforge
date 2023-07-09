metafile {
  path = [
    "metadata.hcl",
    "root.hcl",
  ]
}

generate "variable" "meta" {
  config {
    default = merge(meta, {
      account_id = lookup(metadata.aws_account_id_map, metadata.env)
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
      AccountId = lookup(metadata.aws_account_id_map, metadata.env)
      Owner     = metadata.owner
      Region    = metadata.region
      Project   = metadata.project
      Env       = metadata.env
    }
    assume_role {
      role_arn = format(
        "arn:aws:iam::%s:role/service-provision-role",
        lookup(metadata.aws_account_id_map, metadata.env)
      )
    }
  }
}

generate "provider" "aws" {
  when = metadata.multi_region
  config {
    region = "us-east-1"
    default_tags {
      AccountId = lookup(metadata.aws_account_id_map, metadata.env)
      Owner     = metadata.owner
      Region    = "us-east-1"
      Project   = metadata.project
      Env       = metadata.env
    }
    assume_role {
      role_arn = format(
        "arn:aws:iam::%s:role/service-provision-role",
        lookup(metadata.aws_account_id_map, metadata.env)
      )
    }
  }
}

