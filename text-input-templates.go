package goschemaform

const tmplTextInputForm = `
{
    "key": "{{.Key}}",
    "type": "{{.Type}}",
    "placeholder": "{{.PlaceHolder}}"{{if .ConditionCheck}},
    "condition": "model.{{.Condition}}"{{end}}
}
`

const tmplTextInputSchema = `
"{{.Key}}": {
    "title": "{{.Title}}",
    "type": "string"{{if .MaxLength}},
    "maxLength": "{{.MaxLengthString}}"{{end}}
}
`
