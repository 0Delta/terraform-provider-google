// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataCatalogTagTemplateIamBindingGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
		"role":          "roles/viewer",
		"force_delete":  true,
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataCatalogTagTemplateIamBinding_basicGenerated(context),
			},
			{
				ResourceName:      "google_data_catalog_tag_template_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/tagTemplates/%s roles/viewer", getTestProjectFromEnv(), getTestRegionFromEnv(), fmt.Sprintf("tf_test_my_template%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Test Iam Binding update
				Config: testAccDataCatalogTagTemplateIamBinding_updateGenerated(context),
			},
			{
				ResourceName:      "google_data_catalog_tag_template_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/tagTemplates/%s roles/viewer", getTestProjectFromEnv(), getTestRegionFromEnv(), fmt.Sprintf("tf_test_my_template%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccDataCatalogTagTemplateIamMemberGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
		"role":          "roles/viewer",
		"force_delete":  true,
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				// Test Iam Member creation (no update for member, no need to test)
				Config: testAccDataCatalogTagTemplateIamMember_basicGenerated(context),
			},
			{
				ResourceName:      "google_data_catalog_tag_template_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/tagTemplates/%s roles/viewer user:admin@hashicorptest.com", getTestProjectFromEnv(), getTestRegionFromEnv(), fmt.Sprintf("tf_test_my_template%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccDataCatalogTagTemplateIamPolicyGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
		"role":          "roles/viewer",
		"force_delete":  true,
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataCatalogTagTemplateIamPolicy_basicGenerated(context),
			},
			{
				ResourceName:      "google_data_catalog_tag_template_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/tagTemplates/%s", getTestProjectFromEnv(), getTestRegionFromEnv(), fmt.Sprintf("tf_test_my_template%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccDataCatalogTagTemplateIamPolicy_emptyBinding(context),
			},
			{
				ResourceName:      "google_data_catalog_tag_template_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/tagTemplates/%s", getTestProjectFromEnv(), getTestRegionFromEnv(), fmt.Sprintf("tf_test_my_template%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccDataCatalogTagTemplateIamMember_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_data_catalog_tag_template" "basic_tag_template" {
  tag_template_id = "tf_test_my_template%{random_suffix}"
  region = "us-central1"
  display_name = "Demo Tag Template"

  fields {
    field_id = "source"
    display_name = "Source of data asset"
    type {
      primitive_type = "STRING"
    }
    is_required = true
  }

  fields {
    field_id = "num_rows"
    display_name = "Number of rows in the data asset"
    type {
      primitive_type = "DOUBLE"
    }
  }

  fields {
    field_id = "pii_type"
    display_name = "PII type"
    type {
      enum_type {
        allowed_values {
          display_name = "EMAIL"
        }
        allowed_values {
          display_name = "SOCIAL SECURITY NUMBER"
        }
        allowed_values {
          display_name = "NONE"
        }
      }
    }
  }

  force_delete = "%{force_delete}"
}

resource "google_data_catalog_tag_template_iam_member" "foo" {
  tag_template = google_data_catalog_tag_template.basic_tag_template.name
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}
`, context)
}

func testAccDataCatalogTagTemplateIamPolicy_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_data_catalog_tag_template" "basic_tag_template" {
  tag_template_id = "tf_test_my_template%{random_suffix}"
  region = "us-central1"
  display_name = "Demo Tag Template"

  fields {
    field_id = "source"
    display_name = "Source of data asset"
    type {
      primitive_type = "STRING"
    }
    is_required = true
  }

  fields {
    field_id = "num_rows"
    display_name = "Number of rows in the data asset"
    type {
      primitive_type = "DOUBLE"
    }
  }

  fields {
    field_id = "pii_type"
    display_name = "PII type"
    type {
      enum_type {
        allowed_values {
          display_name = "EMAIL"
        }
        allowed_values {
          display_name = "SOCIAL SECURITY NUMBER"
        }
        allowed_values {
          display_name = "NONE"
        }
      }
    }
  }

  force_delete = "%{force_delete}"
}

data "google_iam_policy" "foo" {
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
  }
}

resource "google_data_catalog_tag_template_iam_policy" "foo" {
  tag_template = google_data_catalog_tag_template.basic_tag_template.name
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccDataCatalogTagTemplateIamPolicy_emptyBinding(context map[string]interface{}) string {
	return Nprintf(`
resource "google_data_catalog_tag_template" "basic_tag_template" {
  tag_template_id = "tf_test_my_template%{random_suffix}"
  region = "us-central1"
  display_name = "Demo Tag Template"

  fields {
    field_id = "source"
    display_name = "Source of data asset"
    type {
      primitive_type = "STRING"
    }
    is_required = true
  }

  fields {
    field_id = "num_rows"
    display_name = "Number of rows in the data asset"
    type {
      primitive_type = "DOUBLE"
    }
  }

  fields {
    field_id = "pii_type"
    display_name = "PII type"
    type {
      enum_type {
        allowed_values {
          display_name = "EMAIL"
        }
        allowed_values {
          display_name = "SOCIAL SECURITY NUMBER"
        }
        allowed_values {
          display_name = "NONE"
        }
      }
    }
  }

  force_delete = "%{force_delete}"
}

data "google_iam_policy" "foo" {
}

resource "google_data_catalog_tag_template_iam_policy" "foo" {
  tag_template = google_data_catalog_tag_template.basic_tag_template.name
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccDataCatalogTagTemplateIamBinding_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_data_catalog_tag_template" "basic_tag_template" {
  tag_template_id = "tf_test_my_template%{random_suffix}"
  region = "us-central1"
  display_name = "Demo Tag Template"

  fields {
    field_id = "source"
    display_name = "Source of data asset"
    type {
      primitive_type = "STRING"
    }
    is_required = true
  }

  fields {
    field_id = "num_rows"
    display_name = "Number of rows in the data asset"
    type {
      primitive_type = "DOUBLE"
    }
  }

  fields {
    field_id = "pii_type"
    display_name = "PII type"
    type {
      enum_type {
        allowed_values {
          display_name = "EMAIL"
        }
        allowed_values {
          display_name = "SOCIAL SECURITY NUMBER"
        }
        allowed_values {
          display_name = "NONE"
        }
      }
    }
  }

  force_delete = "%{force_delete}"
}

resource "google_data_catalog_tag_template_iam_binding" "foo" {
  tag_template = google_data_catalog_tag_template.basic_tag_template.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}
`, context)
}

func testAccDataCatalogTagTemplateIamBinding_updateGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_data_catalog_tag_template" "basic_tag_template" {
  tag_template_id = "tf_test_my_template%{random_suffix}"
  region = "us-central1"
  display_name = "Demo Tag Template"

  fields {
    field_id = "source"
    display_name = "Source of data asset"
    type {
      primitive_type = "STRING"
    }
    is_required = true
  }

  fields {
    field_id = "num_rows"
    display_name = "Number of rows in the data asset"
    type {
      primitive_type = "DOUBLE"
    }
  }

  fields {
    field_id = "pii_type"
    display_name = "PII type"
    type {
      enum_type {
        allowed_values {
          display_name = "EMAIL"
        }
        allowed_values {
          display_name = "SOCIAL SECURITY NUMBER"
        }
        allowed_values {
          display_name = "NONE"
        }
      }
    }
  }

  force_delete = "%{force_delete}"
}

resource "google_data_catalog_tag_template_iam_binding" "foo" {
  tag_template = google_data_catalog_tag_template.basic_tag_template.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com", "user:paddy@hashicorp.com"]
}
`, context)
}
