# uablacklist
## Usage of ./generator:
	-drop string
		rpz drop IP address (default "127.0.0.1")
	-out string
		result file name (default "./db.rpz.zone")
	-tpl string
		template file name (default "./configs/db.rpz.zone.tpl")
	-uri string
        JSON URI address (default "https://uablacklist.net/all.json")
	-file string
		blocked sites file name (default "./configs/blocked_sites.json")
### Fetch https://uablacklist.net API and generates db.rpz.zone file for bind9
