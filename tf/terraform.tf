terraform {
  backend "azurerm" {
    storage_account_name  = "athenstfstate"
    container_name        = "athenstfstate"
    key                   = "terraform.tfstate"    
  }
}

resource "azurerm_resource_group" "athens-tfstate" {
  name     = "athens-tfstate"
  location = "eastus"
}
