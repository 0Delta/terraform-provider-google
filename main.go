package main

import (
	"fmt"
	// "github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-google/google"
)

func co(r *schema.Resource, prefix string){
	for k, e := range r.Schema {
		req := ""
		if e.Required { req = req + "R" }
		if e.Computed { req = req + "C" }
		if e.Optional { req = req + "O" }
		fmt.Printf("%s	%s%s\n", req, prefix, k)
		if e.Elem != nil {
			t, ok := e.Elem.(*schema.Resource)
			if ok {
				co(t, prefix + k + "	")
			}
		}
	}
}


func main() {
	rm := google.ResourceMap()
	for k, r := range rm {
		co(r, k + "	")
	}
	// plugin.Serve(&plugin.ServeOpts{
	// 	ProviderFunc: google.Provider})
}
