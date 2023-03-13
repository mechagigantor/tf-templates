package main

type WorkspaceFolderStructure struct {
	Environments []string
	Medallions   []string
	Projects     []string
}

type Schema struct {
	CatalogName string
	Name        string
	Comment     string
	Tags        []map[string]string
	DependsOn   []string
}

type Catalog struct {
	MetastoreId int
	Name        string
	Comment     string
	Tags        []map[string]string
	DependsOn   []string
}
