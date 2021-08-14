# terraflow

Opinionated configuration management for Terraform

Terraflow expects (& can generate) the following component/application folder structure.

```
config/
├── environments/
│   ├── dev.tfvars
│   ├── stage.tfvars
│   └── prod.tfvars
├── components/
│   ├── compfoo.tfvars
│   └── compbar.tfvars
components/
├── compfoo/
│   ├── main.tf
│   ├── outputs.tf
│   ├── secrets.tf
│   ├── security.tf
│   └── variables.tf
└── compbar/
    ├── main.tf
    ├── outputs.tf
    ├── secrets.tf
    ├── security.tf
    └── variables.tf
```

