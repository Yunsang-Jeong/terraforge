metafiles = [
  "metadata.hcl"
]

generate "terraform" {
  config {
    required_providers {
      aws = {
        source = "a/b/c"
      }
      keycloak = {
        source  = "a/b/c"
        version = "> 1.0.0"
      }
    }
  }
}

generate "provider" "aws" {
  config {
    alias  = "default"
    region = meta.region
    default_tags {
      Owner   = meta.owner
      Region  = meta.region
      Project = meta.project
      Env     = meta.env
    }
  }
}

generate "provider" "aws" {
  when = meta.multi_region
  config {
    alias  = "use1"
    region = "us-east-1"
    assume_role {
      role_arn = "asd"
    }
    default_tags {
      Owner   = meta.owner
      Region  = "us-east-1"
      Project = meta.project
      Env     = meta.env
    }
  }
}

generate "provider" "keycloak" {
  when    = meta.env == "dev" ? true : false
  version = "3.0.0"
  config {
    url = "http://my.ldap.com"
  }
}

generate "variable" "project" {
  when = meta.env == "dev" ? true : false
  config {
    url = meta.project
  }
}
