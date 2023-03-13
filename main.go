package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	wsFile, err := os.Create("workspace_structure.md")
	if err != nil {
		log.Fatalln(err)
	}

	var catalogs []Catalog
	var schemas []Schema

	ws := WorkspaceFolderStructure{
		Environments: []string{"nonprod", "prod"},
		Medallions:   []string{"bronze", "silver", "gold"},
		Projects:     []string{"third_party", "mirror"},
	}

	var wsPaths []string
	for _, e := range ws.Environments {
		_, err = wsFile.WriteString(fmt.Sprintf("%s\n", e))
		if err != nil {
			log.Fatalln(err)
		}

		for _, m := range ws.Medallions {
			_, err = wsFile.WriteString(fmt.Sprintf("\t- %s\n", m))
			if err != nil {
				log.Fatalln(err)
			}

			catalogName := fmt.Sprintf("%s_%s", e, m)
			var c = Catalog{
				MetastoreId: 1,
				Name:        catalogName,
				Comment:     "managed by terraform",
				Tags: []map[string]string{
					{"owner": "me"},
				},
			}
			catalogs = append(catalogs, c)
			for _, p := range ws.Projects {
				_, err = wsFile.WriteString(fmt.Sprintf("\t\t- %s\n", p))
				if err != nil {
					log.Fatalln(err)
				}

				s := Schema{
					CatalogName: catalogName,
					Name:        p,
					Comment:     "managed by terraform",
					Tags: []map[string]string{
						{"owner": "me"},
					},
				}
				schemas = append(schemas, s)
				wsPaths = append(wsPaths, fmt.Sprintf("/%s/%s/%s", e, m, p))
			}
		}
	}

	err = wsFile.Close()
	if err != nil {
		log.Fatalln(err)
	}

	err = NewCatalogFile(catalogs)
	if err != nil {
		log.Fatalln(err)
	}

	err = NewSchemaFile(schemas)
	if err != nil {
		log.Fatalln(err)
	}

	err = NewWorkspaceFileStructure(wsPaths)
	if err != nil {
		log.Fatalln(err)
	}

	err = TestHclParse()
	if err != nil {
		log.Fatalln(err)
	}
}
