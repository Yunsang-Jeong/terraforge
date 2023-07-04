metafiles = [
  "metadata.hcl"
]

generate "provider" "aws" {
  when = true
  config {
    alias  = "default"
    region = region
    default_tags = {
      Owner   = metadata.owner
      Region  = metadata.region
      Project = metadata.project
      Env     = metadata.env
    }
  }
}

generate "provider" "aws" {
  when = metadata.multi_region
  config {
    alias  = "use1"
    region = "us-east-1"
    default_tags = {
      Owner   = metadata.owner
      Region  = "us-east-1"
      Project = metadata.project
      Env     = metadata.env
    }
  }
}

generate "provider" "keycloak" {
  when = metadata.env == "dev" ? true : false
  config {
    alias  = "use1"
    region = "us-east-1"
    default_tags = {
      Owner   = metadata.owner
      Region  = "us-east-1"
      Project = metadata.project
      Env     = metadata.env
    }
  }
}

