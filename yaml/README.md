##
### Installieren des Modules für yaml
Voraussetzung ist die Importzeile im Code:
```
import ...   
"gopkg.in/yaml.v3"
... )
```

Das Package hat den Namen "main".

###
```
go mod init
```
```
go mod tidy
```
```
go install
```
Zur Verifikation der Installation:
```
go list -m  all
```
