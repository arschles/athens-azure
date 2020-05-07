terraform {
  backend "azurerm" {
    storage_account_name  = "athensci"
    container_name        = "tstate"
    key                   = "terraform.tfstate" 
  }
}

resource "azurerm_resource_group" "athens-state-secure" {
  name     = "athens-tfstate"
  location = "eastus"
}
