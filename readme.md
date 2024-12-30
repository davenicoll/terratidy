# TerraTidy

terratidy is a powerful tool designed to format and lint your Terraform code, ensuring it adheres to best practices. By analyzing your code, terratidy can identify common issues and automatically fix them, making your configurations cleaner, more consistent, and easier to maintain. Forked from the [AVMFIX tool](https://github.com/lonegunmanb/avmfix).

While many issues can be resolved automatically, some may require manual adjustments. Following established best practices, such as proper directory structures, naming conventions, and clear documentation, will help you maintain high-quality Terraform modules.

terratidy can fix the following issues:

- [Orders Within resource and data Blocks](https://github.com/Azure/terraform-azure-modules/blob/main/codex/logic_code/resource.md#orders-within-resource-and-data-blocks)
- [Order to define variable](https://github.com/Azure/terraform-azure-modules/blob/main/codex/logic_code/variables.tf.md#order-to-define-variable)
- [Do not declare `nullable = true` for `variable`](https://github.com/Azure/terraform-azure-modules/blob/main/codex/logic_code/variables.tf.md#do-not-declare-nullable--true)
- Do not declare `sensitive = false` for `variable`
- [`output` should be arranged alphabetically](https://github.com/Azure/terraform-azure-modules/blob/main/codex/logic_code/outputs.md#output-should-be-arranged-alphabetically)
- Do not declare `sensitive = false` for `output`
- [`local` should be arranged alphabetically](https://github.com/Azure/terraform-azure-modules/blob/main/codex/logic_code/locals.tf.md#local-should-be-arranged-alphabetically)
- Orders in `moved` block. (`from` then `to`)

## Installation

```bash
go install github.com/davenicoll/terratidy@latest
```

# How to use

To use `terratidy`, open a shell or terminal and run the following command:

```shell
terratidy -folder /path/to/your/terraform/module
```

Replace `/path/to/your/terraform/module` with the path to the directory containing your Terraform module.

## Pre-commit hook

You can also use `terratidy` in your pre-commit config to enforce consistent standards in your repositories.

```yaml
- hooks:
    - id: terratidy
      name: terratidy
      entry: /usr/bin/env bash -c 'for dir in $(git diff --cached --name-only --diff-filter=ACM | grep ".tf$" | xargs -n1 dirname | sort -u); do terratidy --folder "$dir"; done'
      language: script
      files: \.tf$
  repo: local
```

The tool will analyze the specified directory and automatically apply fixes for any issues it identifies. If the process completes successfully, you'll see the message "Tidy completed successfully." If an error occurs during the process, `terratidy` will display an error message.

Keep in mind that `terratidy` may not be able to resolve all issues automatically. Manual intervention may be required for some problems.

# Supported Providers

`terratidy` currently supports variable block description generation for the following providers:

- Alicloud (`alicloud`)
- AWS (`aws`)
- AWS Cloud Control API (`awscc`)
- AzAPI (`azapi`)
- Azure Resource Manager (`azurerm`)
- Azure Active Directory (`azuread`)
- Google Cloud Platform (`google`)
- Helm (`helm`)
- Kubernetes (`kubernetes`)
- Local (`local`)
- Modtm (`modtm`)
- Null (`null`)
- Random (`random`)
- Template (`template`)
- Time (`time`)
- Tls (`tls`)
