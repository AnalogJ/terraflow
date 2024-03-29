# terraflow

<p align="center">
  <a href="https://github.com/AnalogJ/terraflow">
  <img width="300" alt="drawbridge_view" src="https://github.com/AnalogJ/terraflow/raw/main/docs/logo.png">
  </a>
</p>

Opinionated configuration management for Terraform

```
go install github.com/analogj/terraflow/cmd/terraflow@latest
```

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

# References
- Logo: [river by Adrien Coquet from the Noun Project](https://thenounproject.com/search/?q=river&i=2419961)
