resource "databricks_catalog" "nonprod_bronze" {
  name         = "nonprod_bronze"
  comment      = "managed by terraform"
  metastore_id = 1
  properties {
    owner = "me"
  }
}

resource "databricks_catalog" "nonprod_silver" {
  name         = "nonprod_silver"
  comment      = "managed by terraform"
  metastore_id = 1
  properties {
    owner = "me"
  }
}

resource "databricks_catalog" "nonprod_gold" {
  name         = "nonprod_gold"
  comment      = "managed by terraform"
  metastore_id = 1
  properties {
    owner = "me"
  }
}

resource "databricks_catalog" "prod_bronze" {
  name         = "prod_bronze"
  comment      = "managed by terraform"
  metastore_id = 1
  properties {
    owner = "me"
  }
}

resource "databricks_catalog" "prod_silver" {
  name         = "prod_silver"
  comment      = "managed by terraform"
  metastore_id = 1
  properties {
    owner = "me"
  }
}

resource "databricks_catalog" "prod_gold" {
  name         = "prod_gold"
  comment      = "managed by terraform"
  metastore_id = 1
  properties {
    owner = "me"
  }
}

