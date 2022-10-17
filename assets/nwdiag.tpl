nwdiag {
{{ range $key, $Subnet := . }}
network {{ slice $Subnet.Network.Name  1 }} {
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
}