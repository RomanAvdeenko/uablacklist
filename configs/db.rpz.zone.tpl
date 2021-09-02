{{ $D := .DROP_URL }};RPZ
$TTL 10
@   IN SOA rpz.zone. rpz.zone. (
    {{.SERIAL}};
    3600;
    300;
    86400;
    60 )
    IN      NS      localhost.
{{range $v := .RECORDS}}{{$v}}	A	{{$D}}
*.{{$v}}	A    {{$D}}
{{end}}
