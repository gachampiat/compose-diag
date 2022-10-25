graph LR
{{ range $VolumeSource, $configurations := .}}
{{range $index, $volume := $configurations}}
{{$volume.ServiceName}} --> {{$volume.VolumeConfig.Target}}_container{{`{{`}}{{$volume.VolumeConfig.Target}}{{`}}`}} 
{{$volume.VolumeConfig.Target}}_container -{{if $volume.VolumeConfig.ReadOnly}}.{{end}}-> {{$volume.VolumeConfig.Source}}[{{if eq $volume.VolumeConfig.Type "bind"}}/{{$volume.VolumeConfig.Source}}/{{else}}({{$volume.VolumeConfig.Source}}){{end}}]
{{end}}
{{ end }}
