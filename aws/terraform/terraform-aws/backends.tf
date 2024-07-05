terraform {
  cloud {
    organization = "pluviophile"

    workspaces {
      name = "mtc-dev"
    }
  }
}
