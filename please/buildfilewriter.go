package please

import "github.com/jamesjarvis/go-deps/module"

var pleaseModuleRuleFormat = `go_module(
    name = "{{.Name}}",
    module = "{{ pkg.module }}",
    {% if pkg.download -%}
    download = "{{ pkg.download }}",
    {% else -%}
    version = "{{ pkg.version }}",
    {% endif -%}
    {%- if pkg.install -%}
    install = {{ pkg.install }},
    {% endif -%}
    {%- if pkg.deps -%}
    deps = {{ pkg.deps }},
    {% endif -%}
    {%- if pkg.exported_deps -%}
    exported_deps = {{ pkg.exported_deps }},
    {% endif -%}
    {%- if pkg.visibility -%}
    visibility = {{ pkg.visibility }},
    {% else -%}
    visibility = ["PUBLIC"],
    {% endif -%}
    {%- if pkg.patch -%}
    patch = "{{ pkg.patch }}",
    {% endif -%}
    {%- if pkg.strip -%}
    strip = {{ pkg.strip }},
    {% endif -%}
    {%- if pkg.binary -%}
    binary = True,
    {% endif -%}
    {%- if pkg.test_only -%}
    test_only = True,
    {% endif -%}
    {%- if pkg.env -%}
    env = {{ pkg.env }},
    {% endif -%}
)

`

type BuildFileWriter struct {

}

func (b *BuildFileWriter) Write(modules map[string]*module.VersionDirectory) {

}