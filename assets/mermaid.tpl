graph LR
{{ range $VolumeSource, $configurations := .}}
{{range $index, $volume := $configurations}}
{{$volume.ServiceName}} --> {{$volume.VolumeConfig.Target}}_{{$volume.ServiceName}}{{`{{`}}{{$volume.VolumeConfig.Target}}{{`}}`}} 
{{$volume.VolumeConfig.Target}}_{{$volume.ServiceName}} -{{if $volume.VolumeConfig.ReadOnly}}.{{end}}-> {{$volume.VolumeConfig.Source}}[{{if eq $volume.VolumeConfig.Type "bind"}}/{{$volume.VolumeConfig.Source}}/{{else}}({{$volume.VolumeConfig.Source}}){{end}}]
{{end}}
{{ end }}
