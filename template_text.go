// generated code, DO NOT EDIT
package main

const codeTemplateText = `
// Code generated by github.com/pjoc-team/go-options.  DO NOT EDIT.

package {{$.packageName}}


{{ if .imports -}}
import (
{{- range .imports }}
    {{ if .Alias }}  {{ .Alias }} "{{ .Path }}"{{ else }}  "{{ .Path }}"{{ end -}}
{{ end }}
)
{{ end }}

{{ $applyOptionFuncType := or $.applyOptionFuncType (printf "%sFunc" (ToTitle $.optionTypeName)) }}

// {{ $applyOptionFuncType }} the func of {{ $.configTypeName }}
type {{ $applyOptionFuncType }} func(c *{{ $.configTypeName }}) error

func (f {{ $applyOptionFuncType }}) apply(c *{{ $.configTypeName }}) error {
return f(c)
}

{{ $applyFuncName := or $.applyFuncName (printf "apply%sOptions" (ToTitle $.configTypeName)) }}

{{ if $.createNewFunc}}
func new{{ $.configTypeName | ToTitle}}(options ...{{ $.optionTypeName }}) ({{ $.configTypeName }}, error) {
var c {{ $.configTypeName }}
err := {{ $applyFuncName }}(&c, options...)
return c, err
}
{{ end }}

func {{ $applyFuncName }}(c *{{ $.configTypeName }}, options ...{{ $.optionTypeName }}) error {
{{- range .options -}}{{ $optionName := .Name }}{{ if .DefaultValue }}
    c.{{ .Name }} = {{ .DefaultValue }}
{{- end }}{{ if .IsStruct }}{{ range .Fields }}{{ if .DefaultValue }}
    c.{{ $optionName }}.{{ .Name }} = {{ .DefaultValue }}
{{- end }}{{ end }}
{{- end }}{{ end }}
for _, o := range options {
if err := o.apply(c); err != nil {
return err
}
}
return nil
}

// {{ $.optionTypeName }} interface {{ $.optionTypeName }}
type {{ $.optionTypeName }} interface {
apply(*{{ $.configTypeName }}) error
}

{{ range .options }}{{ $option := . }}
{{ $name := .PublicName | ToTitle | printf "%s%s" $.optionPrefix }}
{{ if $.optionSuffix }}{{ $name = $.optionSuffix | printf "%s%s" (.PublicName | ToTitle) }}{{ end  }}
{{ if .Docs }}
{{- range $i, $doc := .Docs }}// {{ if eq $i 0 }}{{ $name }} {{ end }}{{ $doc }}{{ end -}}
{{ end -}}
// {{ $name }} option func
func {{ $name }}(
{{- range $i, $f := .Fields }}{{ if ne $i 0 }},{{ end }}{{ $f.ParamName }} {{ $f.Type }}{{ end -}}
) {{ $applyOptionFuncType }} {
    return func(c *{{ $.configTypeName }}) error {
{{- if and $option.IsStruct $option.DefaultIsNil }}
        c.{{ $option.Name }} = new({{ $option.Type }})
{{- end }}
{{- range .Fields }}{{ if $option.IsStruct }}
        c.{{ $option.Name }}.{{ .Name }} = {{ .ParamName }}
{{- else }}
        c.{{ $option.Name }} = {{ if $option.DefaultIsNil }}&{{ end }}{{ .ParamName }}
{{- end }}{{- end }}
        return nil
    }
}
{{ end }}
`
