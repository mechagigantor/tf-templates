resource "databricks_schema" "nonprod_bronze_third_party" {
  name         = "third_party"
  catalog_name = "nonprod_bronze"
  comment      = "managed by terraform"
  properties {
    owner = "me"
  }
  depends_on = ["databricks_catalog.nonprod_bronze"]
}

resource "databricks_schema" "nonprod_bronze_mirror" {
  name         = "mirror"
  catalog_name = "nonprod_bronze"
  comment      = "managed by terraform"
  properties {
    owner = "me"
  }
  depends_on = ["databricks_catalog.nonprod_bronze"]
}

resource "databricks_schema" "nonprod_silver_third_party" {
  name         = "third_party"
  catalog_name = "nonprod_silver"
  comment      = "managed by terraform"
  properties {
    owner = "me"
  }
  depends_on = ["databricks_catalog.nonprod_silver"]
}

resource "databricks_schema" "nonprod_silver_mirror" {
  name         = "mirror"
  catalog_name = "nonprod_silver"
  comment      = "managed by terraform"
  properties {
    owner = "me"
  }
  depends_on = ["databricks_catalog.nonprod_silver"]
}

resource "databricks_schema" "nonprod_gold_third_party" {
  name         = "third_party"
  catalog_name = "nonprod_gold"
  comment      = "managed by terraform"
  properties {
    owner = "me"
  }
  depends_on = ["databricks_catalog.nonprod_gold"]
}

resource "databricks_schema" "nonprod_gold_mirror" {
  name         = "mirror"
  catalog_name = "nonprod_gold"
  comment      = "managed by terraform"
  properties {
    owner = "me"
  }
  depends_on = ["databricks_catalog.nonprod_gold"]
}

resource "databricks_schema" "prod_bronze_third_party" {
  name         = "third_party"
  catalog_name = "prod_bronze"
  comment      = "managed by terraform"
  properties {
    owner = "me"
  }
  depends_on = ["databricks_catalog.prod_bronze"]
}

resource "databricks_schema" "prod_bronze_mirror" {
  name         = "mirror"
  catalog_name = "prod_bronze"
  comment      = "managed by terraform"
  properties {
    owner = "me"
  }
  depends_on = ["databricks_catalog.prod_bronze"]
}

resource "databricks_schema" "prod_silver_third_party" {
  name         = "third_party"
  catalog_name = "prod_silver"
  comment      = "managed by terraform"
  properties {
    owner = "me"
  }
  depends_on = ["databricks_catalog.prod_silver"]
}

resource "databricks_schema" "prod_silver_mirror" {
  name         = "mirror"
  catalog_name = "prod_silver"
  comment      = "managed by terraform"
  properties {
    owner = "me"
  }
  depends_on = ["databricks_catalog.prod_silver"]
}

resource "databricks_schema" "prod_gold_third_party" {
  name         = "third_party"
  catalog_name = "prod_gold"
  comment      = "managed by terraform"
  properties {
    owner = "me"
  }
  depends_on = ["databricks_catalog.prod_gold"]
}

resource "databricks_schema" "prod_gold_mirror" {
  name         = "mirror"
  catalog_name = "prod_gold"
  comment      = "managed by terraform"
  properties {
    owner = "me"
  }
  depends_on = ["databricks_catalog.prod_gold"]
}

