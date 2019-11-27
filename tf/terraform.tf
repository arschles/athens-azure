terraform {
  backend "remote" {
    organization = "arschles"

    workspaces {
      name = "athens"
    }
  }
}
