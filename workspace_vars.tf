variable "workspace_paths" {
  type    = list(string)
  default = ["/nonprod/bronze/third_party", "/nonprod/bronze/mirror", "/nonprod/silver/third_party", "/nonprod/silver/mirror", "/nonprod/gold/third_party", "/nonprod/gold/mirror", "/prod/bronze/third_party", "/prod/bronze/mirror", "/prod/silver/third_party", "/prod/silver/mirror", "/prod/gold/third_party", "/prod/gold/mirror"]
}

