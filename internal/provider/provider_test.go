package theoneapi

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestProvider(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ terraform.ResourceProvider = Provider()
}

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("THE_ONE_API_TOKEN"); v == "" {
		t.Fatal("THE_ONE_API_TOKEN must be set for acceptance tests")
	}
}

func testAccProviderFactories() map[string]func() (*schema.Provider, error) {
	return map[string]func() (*schema.Provider, error){
		"theoneapi": func() (*schema.Provider, error) {
			return Provider(), nil
		},
	}
}

func TestAccBook_basic(t *testing.T) {
	resourceName := "theoneapi_book.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories(),
		CheckDestroy:      testAccCheckBookDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckBookConfigBasic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						resourceName, "name", "Lord of the Rings"),
				),
			},
		},
	})
}

func testAccCheckBookDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "theoneapi_book" {
			continue
		}

		// TODO: Write logic to check if the book still exists.
		// If err == nil, then that means the resource still exists.
	}

	return nil
}

const testAccCheckBookConfigBasic = `
resource "theoneapi_book" "test" {
	name = "Lord of the Rings"
}
`

