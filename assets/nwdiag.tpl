nwdiag {
{{ range $name, $Subnet := .Configuration }}
network {{ $name }} {
{{ $ipam := $Subnet.Network.Ipam.Config }}
{{ if $ipam }}
	address = " {{ (index $ipam 0).Subnet}}"
{{ end}}

{{ range $index, $service := $Subnet.Services}}
	{{ $networkConf := index $service.Networks (slice $Subnet.Network.Name 1) }}
	{{ $service.Name }}
	
	{{ if $networkConf }} 
		[ address = "{{ $networkConf.Ipv4Address }}"]
	{{ end }}
	;
{{ end }}
}
{{ end }}

{{ range $name, $services := .Groups }}
network {{ $name }}-netns {
	address = "127.0.0.1"

	{{ $name }};
	{{ range $index, $service := $services}}
		{{ $service}};
	{{ end }}
}
{{ end }}
}