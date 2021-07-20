---
layout: "sumologic"
page_title: "SumoLogic: sumologic_global_folder"
description: |-
  Provides an easy way to retrieve the Global Folder.
---

# sumologic_global_folder
Provides an easy way to retrieve the Global Folder.


## Example Usage
```hcl
data "sumologic_global_folder" "globalFolder" {}

resource "sumologic_folder" "folder" {
  name        = "test-folder"
  description = "A test folder"
  parent_id   = "${data.sumologic_global_folder.globalFolder.id}"
}
```


## Attributes reference

The following attributes are exported:

- `id` - The ID of the Global Folder.
- `name` - The name of the Global Folder.
- `description` - The description of the Global Folder.
