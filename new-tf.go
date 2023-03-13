package main

import (
	"fmt"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty"
	"log"
	"os"
)

func NewCatalogFile(catalogs []Catalog) error {
	f := hclwrite.NewEmptyFile()
	// create new file on system
	tfFile, err := os.Create("catalog.tf")
	if err != nil {
		fmt.Println(err)
		return err
	}
	// initialize the body of the new file object
	rootBody := f.Body()

	for _, c := range catalogs {
		bs := rootBody.AppendNewBlock("resource", []string{"databricks_catalog", c.Name})
		bsBody := bs.Body()
		bsBody.SetAttributeValue("name", cty.StringVal(c.Name))
		bsBody.SetAttributeValue("comment", cty.StringVal(c.Comment))
		bsBody.SetAttributeValue("metastore_id", cty.NumberIntVal(int64(c.MetastoreId)))
		//bsBody.SetAttributeValue("properties", cty.MapVal())
		rootBody.AppendNewline()

		propBlock := bsBody.AppendNewBlock("properties", nil)
		propBody := propBlock.Body()
		for _, mapVals := range c.Tags {
			for k, v := range mapVals {
				propBody.SetAttributeValue(k, cty.StringVal(v))
			}
		}

		var dependsOn []cty.Value
		for _, d := range c.DependsOn {
			dependsOn = append(dependsOn, cty.StringVal(d))
		}
	}

	_, err = tfFile.Write(f.Bytes())

	return err
}

func NewSchemaFile(schemas []Schema) error {
	f := hclwrite.NewEmptyFile()
	// create new file on system
	tfFile, err := os.Create("schema.tf")
	if err != nil {
		fmt.Println(err)
		return err
	}
	// initialize the body of the new file object
	rootBody := f.Body()

	for _, s := range schemas {
		bs := rootBody.AppendNewBlock("resource", []string{"databricks_schema", fmt.Sprintf("%s_%s", s.CatalogName, s.Name)})
		bsBody := bs.Body()
		bsBody.SetAttributeValue("name", cty.StringVal(s.Name))
		bsBody.SetAttributeValue("catalog_name", cty.StringVal(s.CatalogName))
		bsBody.SetAttributeValue("comment", cty.StringVal(s.Comment))
		//bsBody.SetAttributeValue("properties", cty.MapVal())

		propBlock := bsBody.AppendNewBlock("properties", nil)
		propBody := propBlock.Body()
		for _, mapVals := range s.Tags {
			for k, v := range mapVals {
				propBody.SetAttributeValue(k, cty.StringVal(v))
			}
		}

		var dependsOn []cty.Value
		for _, d := range s.DependsOn {
			dependsOn = append(dependsOn, cty.StringVal(d))
		}
		dependsOn = append(dependsOn, cty.StringVal(fmt.Sprintf("%s.%s", "databricks_catalog", s.CatalogName)))
		bsBody.SetAttributeValue("depends_on", cty.ListVal(dependsOn))

		rootBody.AppendNewline()
	}

	_, err = tfFile.Write(f.Bytes())

	return err
}

func NewWorkspaceFileStructure(paths []string) error {
	f := hclwrite.NewEmptyFile()
	// create new file on system
	tfFile, err := os.Create("workspace_vars.tf")
	if err != nil {
		fmt.Println(err)
		return err
	}
	// initialize the body of the new file object
	rootBody := f.Body()

	var vals []cty.Value
	for _, p := range paths {
		vals = append(vals, cty.StringVal(p))
	}

	bs := rootBody.AppendNewBlock("variable", []string{"workspace_paths"})
	bs.Body().SetAttributeRaw("type", TokenizeString("list(string)"))
	bs.Body().SetAttributeValue("default", cty.ListVal(vals))
	rootBody.AppendNewline()

	_, err = tfFile.Write(f.Bytes())

	return err
}

func TokenizeString(s string) hclwrite.Tokens {
	tokens := hclwrite.Tokens{
		{
			Type: hclsyntax.TokenIdent, Bytes: []byte(fmt.Sprintf("%s", s)),
		},
	}

	return tokens
}

// Parses an HCL file into a struct
func TestHclParse() error {
	type foo struct {
		A string            `hcl:"a"`
		B []string          `hcl:"b"`
		C map[string]string `hcl:"c"`
	}

	parser := hclparse.NewParser()
	f, parseDiags := parser.ParseHCLFile("test.hcl")
	if parseDiags.HasErrors() {
		log.Fatal(parseDiags.Error())
	}

	var fooInstance foo
	decodeDiags := gohcl.DecodeBody(f.Body, nil, &fooInstance)
	if decodeDiags.HasErrors() {
		log.Fatal(decodeDiags.Error())
	}

	fmt.Printf("%#v", fooInstance)

	return nil
}
