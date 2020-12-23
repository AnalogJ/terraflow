# terraflow

Terraflow generates the following folder structure 

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

