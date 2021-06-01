;RPZ
$TTL 10
@   IN SOA rpz.zone. rpz.zone. (
    {{.SERIAL}};
    3600;
    300;
    86400;
    60
)
IN      NS      localhost.

{{range $v := .RECORDS}}IN  A   {{$v}}      127.0.0.1
IN  A   *.{{$v}}    127.0.0.1
{{end}}
