package vthunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_FW_ALG_FTP_RESOURCE = `
resource "vthunder_fw_alg_ftp" "FwAlgTest" {
	default_port_disable = "default-port-disable" 
}
`

//Acceptance test
func TestAccFwAlgFtp_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_FW_ALG_FTP_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("vthunder_fw_alg_ftp.FwAlgTest", "default_port_disable", "default-port-disable"),
				),
			},
		},
	})
}
